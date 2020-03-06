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

func DefaultOption(S2SBoardCastAddr, S2CBoardCastAddr string, initServer *pb.S2SInfo) (option Option, err error) {
	option = Option{ //初始化
		GRPCOption: GRPCOption{
			S2SServerOption: grpcServiceServer.DefaultOption(),
			S2CServerOption: grpcServiceServer.DefaultOption(),
			S2SClientOption: grpcServiceClient.DefaultOption(),
		},
	}
	if initServer == nil { //构造初始轮询服务器
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
	init, err := initServer.Unpack() //解包初始轮询服务器
	if err != nil {                  //构造服务设置
		return
	}
	option.ServiceOption = server.DefaultOption(*init, nil, nil, nil)
	option.ServiceOption.S2SRegistryOption.RequestSendOption = &pb.RequestSendOption{Addr: S2SBoardCastAddr}
	option.ServiceOption.S2CRegistryOption.RequestSendOption = &pb.RequestSendOption{Addr: S2CBoardCastAddr}
	return
}
