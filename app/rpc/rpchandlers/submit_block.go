package rpchandlers

import (
	"encoding/json"
	"github.com/catspa3/catspad/app/appmessage"
	"github.com/catspa3/catspad/app/protocol/protocolerrors"
	"github.com/catspa3/catspad/app/rpc/rpccontext"
	"github.com/catspa3/catspad/domain/consensus/ruleerrors"
	"github.com/catspa3/catspad/domain/consensus/utils/consensushashing"
	"github.com/catspa3/catspad/infrastructure/network/netadapter/router"
	"github.com/catspa3/catspad/util/storedb"
	"github.com/pkg/errors"
	"github.com/bmatsuo/lmdb-go/lmdb"

	"fmt"
)

// HandleSubmitBlock handles the respectively named RPC command
func HandleSubmitBlock(context *rpccontext.Context, _ *router.Router, request appmessage.Message) (appmessage.Message, error) {
	submitBlockRequest := request.(*appmessage.SubmitBlockRequestMessage)

	var err error
	isSynced := false
	// The node is considered synced if it has peers and consensus state is nearly synced
	if context.ProtocolManager.Context().HasPeers() {
		isSynced, err = context.ProtocolManager.Context().IsNearlySynced()
		if err != nil {
			return nil, err
		}
	}

	if !context.Config.AllowSubmitBlockWhenNotSynced && !isSynced {
		return &appmessage.SubmitBlockResponseMessage{
			Error:        appmessage.RPCErrorf("Block not submitted - node is not synced"),
			RejectReason: appmessage.RejectReasonIsInIBD,
		}, nil
	}

	domainBlock, err := appmessage.RPCBlockToDomainBlock(submitBlockRequest.Block)
	if err != nil {
		return &appmessage.SubmitBlockResponseMessage{
			Error:        appmessage.RPCErrorf("Could not parse block: %s", err),
			RejectReason: appmessage.RejectReasonBlockInvalid,
		}, nil
	}

	if !submitBlockRequest.AllowNonDAABlocks {
		virtualDAAScore, err := context.Domain.Consensus().GetVirtualDAAScore()
		if err != nil {
			return nil, err
		}
		// A simple heuristic check which signals that the mined block is out of date
		// and should not be accepted unless user explicitly requests
		daaWindowSize := uint64(context.Config.NetParams().DifficultyAdjustmentWindowSize)
		if virtualDAAScore > daaWindowSize && domainBlock.Header.DAAScore() < virtualDAAScore-daaWindowSize {
			return &appmessage.SubmitBlockResponseMessage{
				Error: appmessage.RPCErrorf("Block rejected. Reason: block DAA score %d is too far "+
					"behind virtual's DAA score %d", domainBlock.Header.DAAScore(), virtualDAAScore),
				RejectReason: appmessage.RejectReasonBlockInvalid,
			}, nil
		}
	}

	err = context.ProtocolManager.AddBlock(domainBlock)
	if err != nil {
		isProtocolOrRuleError := errors.As(err, &ruleerrors.RuleError{}) || errors.As(err, &protocolerrors.ProtocolError{})
		if !isProtocolOrRuleError {
			return nil, err
		}

		jsonBytes, _ := json.MarshalIndent(submitBlockRequest.Block.Header, "", "    ")
		if jsonBytes != nil {
			log.Warnf("The RPC submitted block triggered a rule/protocol error (%s), printing "+
				"the full header for debug purposes: \n%s", err, string(jsonBytes))
		}

		return &appmessage.SubmitBlockResponseMessage{
			Error:        appmessage.RPCErrorf("Block rejected. Reason: %s", err),
			RejectReason: appmessage.RejectReasonBlockInvalid,
		}, nil
	}

	hash := consensushashing.BlockHash(domainBlock)
//	log.Infof("Accepted block %s via submitBlock", consensushashing.BlockHash(domainBlock))
	log.Infof("Accepted block %s via submitBlock", hash)

	if storedb.Store {
		rpcBlock := appmessage.DomainBlockToRPCBlock(domainBlock)
		err = context.PopulateBlockWithVerboseData(rpcBlock, domainBlock.Header, domainBlock, true)
		if err != nil {
			if errors.Is(err, rpccontext.ErrBuildBlockVerboseDataInvalidBlock) {
				errorMessage := &appmessage.GetBlockResponseMessage{}
				errorMessage.Error = appmessage.RPCErrorf("Block %s is invalid", hash)
				return errorMessage, nil
			}
			fmt.Println("PopulateBlockWithVerboseData error:", err)
			return nil, err
		}

		for _, txid := range rpcBlock.VerboseData.TransactionIDs {
			fmt.Printf("%+v\n", txid)
			err = storedb.Env.Update(func(txn *lmdb.Txn) error {
				err := txn.Put(storedb.Dbi, []byte(txid), []byte(hash.String()), 0)
				return err
			})
			if err != nil {
				fmt.Println(err)
			}
		}
	}


	response := appmessage.NewSubmitBlockResponseMessage()
	return response, nil
}
