package client

import (
	"github.com/yindaheng98/gogisnet/client"
	pb "github.com/yindaheng98/gogisnet/grpc/protocol/protobuf"
	"github.com/yindaheng98/gogisnet/grpc/protocol/registrant"
	"github.com/yindaheng98/gogisnet/message"
)

type Option struct {
	ServiceOption client.Option
	GRPCOption    registrant.GRPCRegistrantOption
	initServer    message.S2CInfo
}

func DefaultOption(initServer *pb.S2CInfo) (option Option, err error) {
	init, err := initServer.Unpack()
	if err != nil {
		return
	}
	option = Option{
		ServiceOption: client.DefaultOption(*init, nil),
		GRPCOption:    registrant.DefaultOption(),
		initServer:    *init,
	}
	option.ServiceOption.ResponseSendOption = &pb.ResponseSendOption{}
	return
}
