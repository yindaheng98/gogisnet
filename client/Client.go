package client

import (
	"github.com/yindaheng98/gogisnet/protocol"
	"github.com/yindaheng98/gogistry"
	"github.com/yindaheng98/gogistry/registrant"
)

type Client struct {
	*registrant.Registrant
}

func New(info protocol.ClientInfo, option Option) *Client {
	ClientRegistrantInfo := protocol.ClientRegistrantInfo{
		ClientInfo:         info,
		ResponseSendOption: option.ResponseSendOption,
	}
	Registrant := gogistery.NewRegistrant(ClientRegistrantInfo,
		option.RegistryN, option.CandidateList, option.RetryNController, option.RequestProto)
	return &Client{Registrant}
}
