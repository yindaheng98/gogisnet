package client

import (
	"github.com/yindaheng98/gogisnet/grpc/option"
	pb "github.com/yindaheng98/gogisnet/grpc/protocol/protobuf"
	"github.com/yindaheng98/gogisnet/grpc/protocol/registrant"
)

//Option is the options for gRPC gogisnet client
type Option struct {
	ServiceOption option.RegistrantOption
	GRPCOption    registrant.GRPCRegistrantOption

	//InitServer is the information about the first server that the client should connect
	InitServer *pb.S2CInfo
}

//DefaultOption returns a default Option
func DefaultOption() Option {
	ServiceOption := option.DefaultRegistrantOption()
	ServiceOption.RegistryN = 1
	return Option{
		ServiceOption: ServiceOption,
		GRPCOption:    registrant.DefaultOption(),
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
