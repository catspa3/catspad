package rpchandlers

import (
	"github.com/catspa3/catspad/app/appmessage"
	"github.com/catspa3/catspad/app/rpc/rpccontext"
	"github.com/catspa3/catspad/infrastructure/network/netadapter/router"
)

// HandleGetHeaders handles the respectively named RPC command
func HandleGetHeaders(context *rpccontext.Context, _ *router.Router, request appmessage.Message) (appmessage.Message, error) {
	response := &appmessage.GetHeadersResponseMessage{}
	response.Error = appmessage.RPCErrorf("not implemented")
	return response, nil
}
