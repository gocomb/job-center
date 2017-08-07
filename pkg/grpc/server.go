package grpc

import (
	"log"
	"net"

	pb "github.com/job-center/pkg/grpc/types"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type RpcFlowServer struct {
	Port string
}

// SayHello implements helloworld.GreeterServer
func (sc *RpcFlowServer) TestClinet(ctx context.Context, in *pb.RpcFlowRequest) (*pb.RpcFlowResponse, error) {
	return &pb.RpcFlowResponse{Response: "receive client: " + in.Request}, nil
}

func (sc *RpcFlowServer) ServerConn() {
	lis, err := net.Listen("tcp", sc.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterRpcFlowServer(s, &RpcFlowServer{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
