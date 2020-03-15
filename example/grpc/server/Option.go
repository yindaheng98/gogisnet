package server

import (
	pb "github.com/yindaheng98/gogisnet/example/grpc/protocol/protobuf"
	"github.com/yindaheng98/gogisnet/example/grpc/protocol/registrant"
	"github.com/yindaheng98/gogisnet/example/grpc/protocol/registry"
	"github.com/yindaheng98/gogisnet/message"
	"github.com/yindaheng98/gogisnet/server"
)

type Option struct {
	ServiceOption server.Option
	GRPCOption    GRPCOption
	initServer    message.S2SInfo
}

type GRPCOption struct {
	S2SRegistryOption   registry.GRPCRegistryOption
	S2CRegistryOption   registry.GRPCRegistryOption
	S2SRegistrantOption registrant.GRPCClientOption
}

func DefaultOption(S2SBoardCastAddr, S2CBoardCastAddr string, initServer *pb.S2SInfo) (option Option, err error) {
	option = Option{ //初始化
		GRPCOption: GRPCOption{
			S2SRegistryOption:   registry.DefaultOption(),
			S2CRegistryOption:   registry.DefaultOption(),
			S2SRegistrantOption: registrant.DefaultOption(),
		},
	}
	if initServer == nil { //构造初始轮询服务器
		initServer = &pb.S2SInfo{
			ServerInfo: &pb.ServerInfo{
				ServerID:    "Undefined",
				ServiceType: "Undefined",
			},
			ResponseSendOption: &pb.ResponseSendOption{},
			RequestSendOption:  &pb.RequestSendOption{Addr: "Undefined"},
			Candidates:         nil,
			S2CInfo: &pb.S2CInfo{
				ServerInfo: &pb.ServerInfo{
					ServerID:    "Undefined",
					ServiceType: "Undefined",
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
	option.initServer = *init
	option.ServiceOption = server.DefaultOption(*init, nil, nil, nil)
	option.ServiceOption.S2SRegistryOption.RequestSendOption = &pb.RequestSendOption{Addr: S2SBoardCastAddr}
	option.ServiceOption.S2CRegistryOption.RequestSendOption = &pb.RequestSendOption{Addr: S2CBoardCastAddr}
	option.ServiceOption.S2SRegistrantOption.ResponseSendOption = &pb.ResponseSendOption{}
	return
}
