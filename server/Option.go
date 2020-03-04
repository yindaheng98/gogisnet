package server

import (
	"github.com/yindaheng98/gogistry/protocol"
	"github.com/yindaheng98/gogistry/registrant"
	"github.com/yindaheng98/gogistry/registry"
	"github.com/yindaheng98/gogistry/util/CandidateList"
	"github.com/yindaheng98/gogistry/util/RetryNController"
	"github.com/yindaheng98/gogistry/util/TimeoutController"
)

type Option struct {
	S2SRegistryOption   S2SRegistryOption
	S2SRegistrantOption S2SRegistrantOption
	S2CRegistryOption   S2CRegistryOption
}

type S2SRegistryOption struct { //服务端面向服务端的收发设置
	RequestSendOption protocol.RequestSendOption
	MaxRegistrants    uint
	TimeoutController registry.TimeoutController
	ResponseProto     protocol.ResponseProtocol
}

type S2SRegistrantOption struct { //服务端面向服务端的收发设置
	ResponseSendOption protocol.ResponseSendOption
	RegistryN          uint64
	CandidateList      registrant.RegistryCandidateList
	RetryNController   registrant.RetryNController
	RequestProto       protocol.RequestProtocol
}

type S2CRegistryOption struct { //服务端面向客户端的接收设置
	RequestSendOption protocol.RequestSendOption
	MaxRegistrants    uint
	TimeoutController registry.TimeoutController
	ResponseProto     protocol.ResponseProtocol
}

func DefaultOption(initS2SRegistry protocol.RegistryInfo,
	S2SResponseProto protocol.ResponseProtocol,
	S2SRequestProto protocol.RequestProtocol,
	S2CResponseProto protocol.ResponseProtocol) Option {
	return Option{
		S2SRegistryOption: S2SRegistryOption{
			MaxRegistrants:    4,
			TimeoutController: TimeoutController.NewLogTimeoutController(1e9, 10e9, 2),
			ResponseProto:     S2SResponseProto,
		},
		S2SRegistrantOption: S2SRegistrantOption{
			RegistryN:        4,
			CandidateList:    CandidateList.NewSimpleCandidateList(4, initS2SRegistry, 1e9, 3),
			RetryNController: RetryNController.SimpleRetryNController{},
			RequestProto:     S2SRequestProto,
		},
		S2CRegistryOption: S2CRegistryOption{
			MaxRegistrants:    8,
			TimeoutController: TimeoutController.NewLogTimeoutController(1e9, 10e9, 2),
			ResponseProto:     S2CResponseProto,
		},
	}
}
