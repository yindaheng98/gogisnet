package server

import (
	grpcServiceClient "github.com/yindaheng98/gogisnet/example/grpc/protocol/client"
	pb "github.com/yindaheng98/gogisnet/example/grpc/protocol/protobuf"
	grpcServiceServer "github.com/yindaheng98/gogisnet/example/grpc/protocol/server"
	"github.com/yindaheng98/gogisnet/server"
)

type Option struct {
	ServiceOption server.Option
	GRPCOption    GRPCOption
}

type GRPCOption struct {
	S2SServerOption grpcServiceServer.GRPCServerOption
	S2CServerOption grpcServiceServer.GRPCServerOption
	S2SClientOption grpcServiceClient.GRPCClientOption
}

func DefaultOption(initServer pb.S2SInfo) (option Option, err error) {
	init, err := initServer.Unpack()
	if err != nil {
		return
	}
	return Option{
		ServiceOption: server.DefaultOption(*init, nil, nil, nil),
		GRPCOption: GRPCOption{
			S2SServerOption: grpcServiceServer.DefaultOption(),
			S2CServerOption: grpcServiceServer.DefaultOption(),
			S2SClientOption: grpcServiceClient.DefaultOption(),
		},
	}, nil
}
