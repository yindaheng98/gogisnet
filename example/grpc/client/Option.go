package client

import (
	"github.com/yindaheng98/gogisnet/client"
	pb "github.com/yindaheng98/gogisnet/example/grpc/protocol/protobuf"
	grpcServiceClient "github.com/yindaheng98/gogisnet/example/grpc/protocol/registrant"
	"github.com/yindaheng98/gogisnet/protocol"
)

type Option struct {
	ServiceOption client.Option
	GRPCOption    grpcServiceClient.GRPCClientOption
	initServer    protocol.S2CInfo
}

func DefaultOption(initServer *pb.S2CInfo) (option Option, err error) {
	init, err := initServer.Unpack()
	if err != nil {
		return
	}
	option = Option{
		ServiceOption: client.DefaultOption(*init, nil),
		GRPCOption:    grpcServiceClient.DefaultOption(),
		initServer:    *init,
	}
	option.ServiceOption.ResponseSendOption = &pb.ResponseSendOption{}
	return
}
