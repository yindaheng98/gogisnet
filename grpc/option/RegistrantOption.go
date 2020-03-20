package option

import (
	pb "github.com/yindaheng98/gogisnet/grpc/protocol/protobuf"
	"github.com/yindaheng98/gogisnet/grpc/protocol/registrant"
	"github.com/yindaheng98/gogisnet/server"
	"github.com/yindaheng98/gogistry/example/RetryNController"
	"time"
)

//RegistryOption is the options for a gogistry registrant
type RegistrantOption struct { //服务端面向服务端的收发设置
	RegistryN              uint64                               `yaml:"RegistryN" usage:"RegistryN defined how much registries this registrant can connect at most."`
	RetryNControllerOption RetryNControllerOption               `yaml:"RetryNControllerOption" usage:"RetryNControllerOption is the option for RetryNController used in gogistry registrant."`
	CandidateListOption    registrant.PingerCandidateListOption `yaml:"CandidateListOption" usage:"The option for PingerCandidateList used in S2CRegistrant and S2SRegistrant."`
}

//DefaultRegistrantOption returns a default RegistrantOption
func DefaultRegistrantOption() RegistrantOption {
	return RegistrantOption{
		RegistryN:              4,
		RetryNControllerOption: defaultRetryNControllerOption(),
		CandidateListOption:    registrant.DefaultPingerCandidateListOption(),
	}
}

func (o RegistrantOption) PutOption(op *server.RegistrantOption) {
	op.RegistryN = o.RegistryN
	o.RetryNControllerOption.PutOption(op.RetryNController.(*RetryNController.LinearRetryNController))
	op.ResponseSendOption = &pb.ResponseSendOption{}
}

//The option for RetryNController used in gogistry registrant
type RetryNControllerOption struct {
	K_RetryN   uint64        `yaml:"K_RetryN" usage:"nextRetryN=lastRetryN * K_RetryN + B_RetryN."`
	B_RetryN   uint64        `yaml:"B_RetryN" usage:"nextRetryN=lastRetryN * K_RetryN + B_RetryN."`
	K_SendTime time.Duration `yaml:"K_SendTime" usage:"nextSendTimeLimit=lastSendTime * K_SendTime + B_SendTime"`
	B_SendTime time.Duration `yaml:"B_SendTime" usage:"nextSendTimeLimit=lastSendTime * K_SendTime + B_SendTime"`
}

func defaultRetryNControllerOption() RetryNControllerOption {
	op := RetryNController.DefaultLinearRetryNController()
	return RetryNControllerOption{
		K_RetryN:   op.K_RetryN,
		B_RetryN:   op.B_RetryN,
		K_SendTime: op.K_SendTime,
		B_SendTime: op.B_SendTime,
	}
}

func (o RetryNControllerOption) PutOption(op *RetryNController.LinearRetryNController) {
	op.K_RetryN = o.K_RetryN
	op.B_RetryN = o.B_RetryN
	op.K_SendTime = o.K_SendTime
	op.B_SendTime = o.B_SendTime
}
