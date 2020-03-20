package option

import (
	pb "github.com/yindaheng98/gogisnet/grpc/protocol/protobuf"
	"github.com/yindaheng98/gogisnet/grpc/protocol/registrant"
	"github.com/yindaheng98/gogisnet/server"
	"github.com/yindaheng98/gogistry/example/RetryNController"
	gogistrant "github.com/yindaheng98/gogistry/registrant"
)

//RegistryOption is the options for a gogistry registrant
type RegistrantOption struct { //服务端面向服务端的收发设置
	RegistryN           uint64                                 `yaml:"RegistryN" usage:"RegistryN defined how much registries this registrant can connect at most."`
	RetryNController    gogistrant.WaitTimeoutRetryNController `yaml:"RetryNController" usage:"RetryNController is the RetryNController used in gogistry registrant."`
	CandidateListOption registrant.PingerCandidateListOption   `yaml:"CandidateListOption" usage:"The option for PingerCandidateList used in S2CRegistrant and S2SRegistrant."`
}

//DefaultRegistrantOption returns a default RegistrantOption
func DefaultRegistrantOption() RegistrantOption {
	return RegistrantOption{
		RegistryN:           4,
		RetryNController:    RetryNController.DefaultLinearRetryNController(),
		CandidateListOption: registrant.DefaultPingerCandidateListOption(),
	}
}

func (o RegistrantOption) PutOption(op *server.RegistrantOption) {
	op.RegistryN = o.RegistryN
	op.RetryNController = o.RetryNController
	op.ResponseSendOption = &pb.ResponseSendOption{}
}
