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

func DefaultOption(initServer *pb.S2SInfo) (Option, error) {
	if initServer == nil {
		initServer = &pb.S2SInfo{
			ServerInfo: &pb.ServerInfo{
				ServerID:       "Undefined",
				ServiceType:    "Undefined",
				AdditionalInfo: "",
			},
			ResponseSendOption: &pb.ResponseSendOption{},
			RequestSendOption:  &pb.RequestSendOption{Addr: "Undefined"},
			Candidates:         nil,
			S2CInfo: &pb.S2CInfo{
				ServerInfo: &pb.ServerInfo{
					ServerID:       "Undefined",
					ServiceType:    "Undefined",
					AdditionalInfo: "",
				},
				RequestSendOption: &pb.RequestSendOption{Addr: "Undefined"},
				Candidates:        nil,
			},
		}
	}
	init, err := initServer.Unpack()
	if err != nil {
		return Option{}, err
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
