package client

import (
	"github.com/yindaheng98/gogisnet/message"
	"github.com/yindaheng98/gogisnet/server"
	"github.com/yindaheng98/gogistry/example/CandidateList"
	"github.com/yindaheng98/gogistry/example/RetryNController"
	"github.com/yindaheng98/gogistry/protocol"
)

type Option server.RegistrantOption

func DefaultOption(initS2CInfo message.S2CInfo, RequestProto protocol.RequestProtocol) Option {
	return Option{
		RegistryN:        1,
		CandidateList:    CandidateList.NewSimpleCandidateList(3, initS2CInfo),
		RetryNController: RetryNController.DefaultLinearRetryNController(),
		RequestProto:     RequestProto,
	}
}
