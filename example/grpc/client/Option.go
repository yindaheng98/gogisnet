package client

import (
	"github.com/yindaheng98/gogisnet/client"
	grpcClient "github.com/yindaheng98/gogisnet/example/grpc/protocol/client"
	"github.com/yindaheng98/gogisnet/protocol"
)

type Option struct {
	ServiceOption client.Option
	GRPCOption    grpcClient.GRPCClientOption
}

func DefaultOption(initServer protocol.S2CInfo) Option {
	return Option{
		ServiceOption: client.DefaultOption(initServer, nil),
		GRPCOption:    grpcClient.DefaultOption(),
	}
}
