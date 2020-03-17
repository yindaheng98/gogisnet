package server

import (
	"github.com/yindaheng98/gogisnet/message"
	"github.com/yindaheng98/gogistry/example/CandidateList"
	"github.com/yindaheng98/gogistry/example/RetryNController"
	"github.com/yindaheng98/gogistry/example/TimeoutController"
	"github.com/yindaheng98/gogistry/protocol"
	"github.com/yindaheng98/gogistry/registrant"
	"github.com/yindaheng98/gogistry/registry"
)

type Option struct {
	S2SRegistryOption   RegistryOption
	S2SRegistrantOption RegistrantOption
	S2CRegistryOption   RegistryOption
}

type RegistryOption struct { //服务端面向服务端的收发设置
	RequestSendOption protocol.RequestSendOption
	MaxRegistrants    uint
	TimeoutController registry.TimeoutController
	ResponseProto     protocol.ResponseProtocol
}

type RegistrantOption struct { //服务端面向服务端的收发设置
	ResponseSendOption protocol.ResponseSendOption
	RegistryN          uint64
	CandidateList      registrant.RegistryCandidateList
	RetryNController   registrant.WaitTimeoutRetryNController
	RequestProto       protocol.RequestProtocol
}

func DefaultOption(initS2SInfo message.S2SInfo,
	S2SResponseProto protocol.ResponseProtocol,
	S2SRequestProto protocol.RequestProtocol,
	S2CResponseProto protocol.ResponseProtocol) Option {
	return Option{
		S2SRegistryOption: RegistryOption{
			MaxRegistrants:    4,
			TimeoutController: TimeoutController.NewLogTimeoutController(1e9, 10e9, 2),
			ResponseProto:     S2SResponseProto,
		},
		S2SRegistrantOption: RegistrantOption{
			RegistryN:        4,
			CandidateList:    CandidateList.NewSimpleCandidateList(4, initS2SInfo, 1e9, 3),
			RetryNController: RetryNController.SimpleRetryNController{},
			RequestProto:     S2SRequestProto,
		},
		S2CRegistryOption: RegistryOption{
			MaxRegistrants:    8,
			TimeoutController: TimeoutController.NewLogTimeoutController(1e9, 10e9, 2),
			ResponseProto:     S2CResponseProto,
		},
	}
}
