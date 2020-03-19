package server

import (
	"github.com/yindaheng98/gogisnet/grpc/protocol/graph"
	pb "github.com/yindaheng98/gogisnet/grpc/protocol/protobuf"
	"github.com/yindaheng98/gogisnet/grpc/protocol/registrant"
	"github.com/yindaheng98/gogisnet/grpc/protocol/registry"
)

//Option contains all the options for gRPC gogisnet server
type Option struct {
	ServiceOption ServiceOption
	GRPCOption    GRPCOption

	//InitServer is the information about the first server that the client should connect
	InitServer *pb.S2SInfo
}

//GRPCOption is the gRPC options for gRPC gogisnet server
type GRPCOption struct {
	S2SRegistryOption   registry.GRPCRegistryOption
	S2CRegistryOption   registry.GRPCRegistryOption
	S2SRegistrantOption registrant.GRPCRegistrantOption
	GraphQueryOption    graph.GraphQueryOption
}

//DefaultOption returns a default Option
func DefaultOption() (option Option) {
	option = Option{ //初始化
		ServiceOption: DefaultServiceOption(),
		GRPCOption: GRPCOption{
			S2SRegistryOption:   registry.DefaultOption(),
			S2CRegistryOption:   registry.DefaultOption(),
			S2SRegistrantOption: registrant.DefaultOption(),
			GraphQueryOption:    graph.DefaultOption(),
		},
	}
	option.InitServer = &pb.S2SInfo{
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
	return
}
