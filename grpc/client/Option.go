package client

import (
	pb "github.com/yindaheng98/gogisnet/grpc/protocol/protobuf"
	"github.com/yindaheng98/gogisnet/grpc/protocol/registrant"
	"github.com/yindaheng98/gogisnet/grpc/server"
	"github.com/yindaheng98/gogistry/example/RetryNController"
)

//Option is the options for gRPC gogisnet client
type Option struct {
	ServiceOption server.RegistrantOption
	GRPCOption    registrant.GRPCRegistrantOption

	//InitServer is the information about the first server that the client should connect
	InitServer *pb.S2CInfo
}

//DefaultOption returns a default Option
func DefaultOption() (option Option) {
	return Option{
		ServiceOption: server.RegistrantOption{
			RegistryN:           1,
			RetryNController:    RetryNController.DefaultLinearRetryNController(),
			CandidateListOption: registrant.DefaultPingerCandidateListOption(),
		},
		GRPCOption: registrant.DefaultOption(),
		InitServer: &pb.S2CInfo{
			ServerInfo: &pb.ServerInfo{
				ServerID:    "Undefined",
				ServiceType: "Undefined",
			},
			RequestSendOption: &pb.RequestSendOption{Addr: "Undefined"},
			Candidates:        nil,
		},
	}
}
