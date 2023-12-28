package client

import (
	"context"
	"github.com/catspa3/catspad/cmd/kaspawallet/daemon/server"
	"time"

	"github.com/pkg/errors"

	"github.com/catspa3/catspad/cmd/kaspawallet/daemon/pb"
	"google.golang.org/grpc"
)

// Connect connects to the kaspawalletd server, and returns the client instance
func Connect(address string) (pb.KaspawalletdClient, func(), error) {
	// Connection is local, so 1 second timeout is sufficient
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(server.MaxDaemonSendMsgSize)))
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, nil, errors.New("catspawallet daemon is not running, start it with `catspawallet start-daemon`")
		}
		return nil, nil, err
	}

	return pb.NewKaspawalletdClient(conn), func() {
		conn.Close()
	}, nil
}
