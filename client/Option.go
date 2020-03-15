package client

import (
	"github.com/yindaheng98/gogisnet/message"
	"github.com/yindaheng98/gogistry/example/CandidateList"
	"github.com/yindaheng98/gogistry/example/RetryNController"
	"github.com/yindaheng98/gogistry/protocol"
	"github.com/yindaheng98/gogistry/registrant"
)

type Option struct {
	ResponseSendOption protocol.ResponseSendOption
	RegistryN          uint64
	CandidateList      registrant.RegistryCandidateList
	RetryNController   registrant.WaitTimeoutRetryNController
	RequestProto       protocol.RequestProtocol
}

func DefaultOption(initS2CInfo message.S2CInfo, RequestProto protocol.RequestProtocol) Option {
	return Option{
		RegistryN:        1,
		CandidateList:    CandidateList.NewSimpleCandidateList(3, initS2CInfo, 1e9, 10),
		RetryNController: RetryNController.SimpleRetryNController{},
		RequestProto:     RequestProto,
	}
}
