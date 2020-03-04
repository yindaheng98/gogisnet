package client

import (
	"context"
	"github.com/yindaheng98/gogisnet/protocol"
	"github.com/yindaheng98/gogistry"
	gogistryProto "github.com/yindaheng98/gogistry/protocol"
	"github.com/yindaheng98/gogistry/registrant"
	"time"
)

type Client struct {
	c2sRegistrant *registrant.Registrant
	Events        *events
}

func New(info protocol.ClientInfo, option Option) *Client {
	RegistrantInfo := protocol.C2SInfo{
		ClientInfo:         info,
		ResponseSendOption: option.ResponseSendOption,
	}
	c2sRegistrant := gogistery.NewRegistrant(RegistrantInfo,
		option.RegistryN, option.CandidateList, option.RetryNController, option.RequestProto)
	c := &Client{c2sRegistrant, nil}
	c.initEvent()
	return c
}

func (c *Client) GetClientInfo() protocol.ClientInfo {
	return c.c2sRegistrant.Info.(protocol.C2SInfo).ClientInfo
}
func (c *Client) GetC2SInfo() protocol.C2SInfo {
	return c.c2sRegistrant.Info.(protocol.C2SInfo)
}

func (c *Client) GetS2CConnections() []protocol.S2CInfo {
	connections := c.c2sRegistrant.GetConnections()
	s2cInfos := make([]protocol.S2CInfo, len(connections))
	for i, c := range connections {
		s2cInfos[i] = c.(protocol.S2CInfo)
	}
	return s2cInfos
}

func (c *Client) Run(ctx context.Context) {
	c.c2sRegistrant.Run(ctx)
}

func (c *Client) SetWatchdogTimeDelta(t time.Duration) {
	c.c2sRegistrant.WatchdogTimeDelta = t
}

func (c *Client) SetCandidateBlacklist(blacklist []protocol.S2CInfo) {
	<-c.c2sRegistrant.CandidateBlacklist
	CandidateBlacklist := make([]gogistryProto.RegistryInfo, len(blacklist))
	for i, c := range blacklist {
		CandidateBlacklist[i] = c
	}
	c.c2sRegistrant.CandidateBlacklist <- CandidateBlacklist
}
