package server

import (
	"github.com/yindaheng98/gogisnet/grpc/protocol/graph"
	pb "github.com/yindaheng98/gogisnet/grpc/protocol/protobuf"
	"github.com/yindaheng98/gogisnet/grpc/protocol/registrant"
	"github.com/yindaheng98/gogisnet/grpc/protocol/registry"
	"github.com/yindaheng98/gogisnet/server"
)

//Option is the options for gRPC gogisnet server
type Option struct {
	ServiceOption server.Option
	GRPCOption    GRPCOption
	initServer    *pb.S2SInfo
}

//Option is the gRPC options for gRPC gogisnet server
type GRPCOption struct {
	S2SRegistryOption   registry.GRPCRegistryOption
	S2CRegistryOption   registry.GRPCRegistryOption
	S2SRegistrantOption registrant.GRPCRegistrantOption
	CandidateListOption registrant.CandidateListOption
	GraphQueryOption    graph.GraphQueryOption
}

//DefaultOption returns a default Option
func DefaultOption(S2SBoardCastAddr, S2CBoardCastAddr, GraphQueryBoardCastAddr string, initServer *pb.S2SInfo) (option Option, err error) {
	option = Option{ //初始化
		GRPCOption: GRPCOption{
			S2SRegistryOption:   registry.DefaultOption(),
			S2CRegistryOption:   registry.DefaultOption(),
			S2SRegistrantOption: registrant.DefaultOption(),
			CandidateListOption: registrant.DefaultCandidateListOption(),
			GraphQueryOption:    graph.DefaultOption(GraphQueryBoardCastAddr),
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
	option.initServer = initServer
	init, err := initServer.Unpack() //解包初始轮询服务器
	if err != nil {                  //构造服务设置
		return
	}
	option.ServiceOption = server.DefaultOption(*init, nil, nil, nil)
	option.ServiceOption.S2SRegistryOption.RequestSendOption = &pb.RequestSendOption{Addr: S2SBoardCastAddr}
	option.ServiceOption.S2CRegistryOption.RequestSendOption = &pb.RequestSendOption{Addr: S2CBoardCastAddr}
	option.ServiceOption.S2SRegistrantOption.ResponseSendOption = &pb.ResponseSendOption{}
	return
}
