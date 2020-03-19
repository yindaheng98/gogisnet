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
	S2SRegistryOption   RegistryOption
	S2SRegistrantOption RegistrantOption
	S2CRegistryOption   RegistryOption
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

	//BoardCastAddr is the IP Address that will send with messages to other server, so other server can easily find the current server
	BoardCastAddr string

	//MaxRegistrants defined how much registrants this registry can connect at most
	MaxRegistrants uint64

	//TimeoutController is the TimeoutController used in gogistry registry
	TimeoutController registry.TimeoutController
}

//PutOption can convert a RegistryOption into a server.RegistryOption
func (o RegistryOption) PutOption(op *server.RegistryOption) {
	op.RequestSendOption = &pb.RequestSendOption{Addr: o.BoardCastAddr}
	op.TimeoutController = o.TimeoutController
	op.MaxRegistrants = o.MaxRegistrants
}

//RegistryOption is the options for a gogistry registrant
type RegistrantOption struct { //服务端面向服务端的收发设置

	//RegistryN defined how much registries this registrant can connect at most
	RegistryN uint64

	//RetryNController is the RetryNController used in gogistry registrant
	RetryNController gogistrant.WaitTimeoutRetryNController

	CandidateListOption registrant.PingerCandidateListOption
}
