package client

import (
	"github.com/yindaheng98/gogisnet/client"
	grpcServiceClient "github.com/yindaheng98/gogisnet/example/grpc/protocol/client"
	pb "github.com/yindaheng98/gogisnet/example/grpc/protocol/protobuf"
)

type Option struct {
	ServiceOption client.Option
	GRPCOption    grpcServiceClient.GRPCClientOption
}

func DefaultOption(initServer *pb.S2CInfo) (option Option, err error) {
	init, err := initServer.Unpack()
	if err != nil {
		return
	}
	return Option{
		ServiceOption: client.DefaultOption(*init, nil),
		GRPCOption:    grpcServiceClient.DefaultOption(),
	}, nil
}
