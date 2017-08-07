package grpc

import (
	"log"

	pb"github.com/job-center/pkg/grpc/types"
	"github.com/job-center/server/util"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type RpcClient struct {
	Address string
}

func (cl *RpcClient) RpcClient() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(cl.Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewRpcFlowClient(conn)

	// Contact the server and print out its response.
	r, err := c.TestClinet(context.Background(), &pb.RpcFlowRequest{Request: "hello grpc server"})
	if err != nil {
		util.Logger.SetFatal(err)
	}
	util.Logger.SetInfo(r.Response)
}
