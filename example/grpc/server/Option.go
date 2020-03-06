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

func DefaultOption(S2SListenAddr, S2CListenAddr string, initServer *pb.S2SInfo) (Option, error) {
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
	ServiceOption := server.DefaultOption(*init, nil, nil, nil)
	ServiceOption.S2SRegistryOption.RequestSendOption = &pb.RequestSendOption{Addr: S2SListenAddr}
	ServiceOption.S2CRegistryOption.RequestSendOption = &pb.RequestSendOption{Addr: S2CListenAddr}
	return Option{
		ServiceOption: ServiceOption,
		GRPCOption: GRPCOption{
			S2SServerOption: grpcServiceServer.DefaultOption(),
			S2CServerOption: grpcServiceServer.DefaultOption(),
			S2SClientOption: grpcServiceClient.DefaultOption(),
		},
	}, nil
}
