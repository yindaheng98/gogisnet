package client

import (
	"github.com/yindaheng98/gogistry/protocol"
	"github.com/yindaheng98/gogistry/registrant"
	"github.com/yindaheng98/gogistry/util/CandidateList"
	"github.com/yindaheng98/gogistry/util/RetryNController"
)

type Option struct {
	ResponseSendOption protocol.ResponseSendOption
	RegistryN          uint
	CandidateList      registrant.RegistryCandidateList
	RetryNController   registrant.RetryNController
	RequestProto       protocol.RequestProtocol
}

func DefaultOption(initRegistry protocol.RegistryInfo, RequestProto protocol.RequestProtocol) Option {
	return Option{
		RegistryN:        1,
		CandidateList:    CandidateList.NewSimpleCandidateList(3, initRegistry, 1e9, 10),
		RetryNController: RetryNController.SimpleRetryNController{},
		RequestProto:     RequestProto,
	}
}
