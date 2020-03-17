package client

import (
	"github.com/yindaheng98/gogisnet/client"
	pb "github.com/yindaheng98/gogisnet/grpc/protocol/protobuf"
	"github.com/yindaheng98/gogisnet/grpc/protocol/registrant"
)

//Option is the options for gRPC gogisnet client
type Option struct {
	ServiceOption       client.Option
	GRPCOption          registrant.GRPCRegistrantOption
	CandidateListOption registrant.CandidateListOption
	initServer          *pb.S2CInfo
}

//DefaultOption returns a default Option
func DefaultOption(initServer *pb.S2CInfo) (option Option, err error) {
	init, err := initServer.Unpack()
	if err != nil {
		return
	}
	option = Option{
		ServiceOption:       client.DefaultOption(*init, nil),
		GRPCOption:          registrant.DefaultOption(),
		CandidateListOption: registrant.DefaultCandidateListOption(),
		initServer:          initServer,
	}
	option.ServiceOption.ResponseSendOption = &pb.ResponseSendOption{}
	return
}
