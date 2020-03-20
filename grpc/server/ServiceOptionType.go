package server

import (
	pb "github.com/yindaheng98/gogisnet/grpc/protocol/protobuf"
	"github.com/yindaheng98/gogisnet/grpc/protocol/registrant"
	"github.com/yindaheng98/gogisnet/server"
	"github.com/yindaheng98/gogistry/example/RetryNController"
	"github.com/yindaheng98/gogistry/example/TimeoutController"
	gogistrant "github.com/yindaheng98/gogistry/registrant"
	"github.com/yindaheng98/gogistry/registry"
)

//Option for Gogisnet service
type ServiceOption struct {
	S2SRegistryOption   RegistryOption   `yaml:"S2SRegistryOption" usage:"Option for S2SRegistry in gogisnet server."`
	S2SRegistrantOption RegistrantOption `yaml:"S2SRegistrantOption" usage:"Option for S2SRegistrant in gogisnet server."`
	S2CRegistryOption   RegistryOption   `yaml:"S2CRegistryOption" usage:"Option for S2CRegistry in gogisnet server."`
}

//DefaultServiceOption returns a default ServiceOption
func DefaultServiceOption() ServiceOption {
	ip := GetIP()
	return ServiceOption{
		S2SRegistryOption: RegistryOption{
			BoardCastAddr:     ip + "4241",
			MaxRegistrants:    4,
			TimeoutController: TimeoutController.DefaultLogTimeoutController(),
		},
		S2SRegistrantOption: RegistrantOption{
			RegistryN:           4,
			RetryNController:    RetryNController.DefaultLinearRetryNController(),
			CandidateListOption: registrant.DefaultPingerCandidateListOption(),
		},
		S2CRegistryOption: RegistryOption{
			BoardCastAddr:     ip + ":4240",
			MaxRegistrants:    16,
			TimeoutController: TimeoutController.DefaultLogTimeoutController(),
		},
	}
}

//RegistryOption is the options for a gogistry registry
type RegistryOption struct { //服务端面向服务端的收发设置
	BoardCastAddr     string                     `yaml:"BoardCastAddr" usage:"The IP Address that will send with messages to other server, so other server can easily find the current server."`
	MaxRegistrants    uint64                     `yaml:"MaxRegistrants" usage:"MaxRegistrants defined how much registrants this registry can connect at most."`
	TimeoutController registry.TimeoutController `yaml:"TimeoutController" usage:"TimeoutController is the TimeoutController used in gogistry registry."`
}

//PutOption can convert a RegistryOption into a server.RegistryOption
func (o RegistryOption) PutOption(op *server.RegistryOption) {
	op.RequestSendOption = &pb.RequestSendOption{Addr: o.BoardCastAddr}
	op.TimeoutController = o.TimeoutController
	op.MaxRegistrants = o.MaxRegistrants
}

//RegistryOption is the options for a gogistry registrant
type RegistrantOption struct { //服务端面向服务端的收发设置
	RegistryN           uint64                                 `yaml:"RegistryN" usage:"RegistryN defined how much registries this registrant can connect at most."`
	RetryNController    gogistrant.WaitTimeoutRetryNController `yaml:"RetryNController" usage:"RetryNController is the RetryNController used in gogistry registrant."`
	CandidateListOption registrant.PingerCandidateListOption   `yaml:"CandidateListOption" usage:"The option for PingerCandidateList used in S2CRegistrant and S2SRegistrant."`
}
