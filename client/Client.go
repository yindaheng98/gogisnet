package client

import (
	"context"
	"github.com/yindaheng98/gogisnet/message"
	"github.com/yindaheng98/gogistry"
	gogistryProto "github.com/yindaheng98/gogistry/protocol"
	"github.com/yindaheng98/gogistry/registrant"
	"time"
)

type Client struct {
	c2sRegistrant *registrant.Registrant
	Events        *events
}

func New(info message.ClientInfo, option Option) *Client {
	RegistrantInfo := message.C2SInfo{
		ClientInfo:         info,
		ResponseSendOption: option.ResponseSendOption,
	}
	c2sRegistrant := gogistery.NewRegistrant(RegistrantInfo,
		option.RegistryN, option.CandidateList, option.RetryNController, option.RequestProto)
	c := &Client{c2sRegistrant, nil}
	c.initEvent()
	return c
}

func (c *Client) GetClientInfo() message.ClientInfo {
	return c.c2sRegistrant.Info.(message.C2SInfo).ClientInfo
}
func (c *Client) GetC2SInfo() message.C2SInfo {
	return c.c2sRegistrant.Info.(message.C2SInfo)
}

func (c *Client) GetS2CConnections() []message.S2CInfo {
	connections := c.c2sRegistrant.GetConnections()
	s2cInfos := make([]message.S2CInfo, len(connections))
	for i, c := range connections {
		s2cInfos[i] = c.(message.S2CInfo)
	}
	return s2cInfos
}

func (c *Client) Run(ctx context.Context) {
	c.c2sRegistrant.Run(ctx)
}

func (c *Client) SetWatchdogTimeDelta(t time.Duration) {
	c.c2sRegistrant.WatchdogTimeDelta = t
}

func (c *Client) SetCandidateBlacklist(blacklist []message.S2CInfo) {
	<-c.c2sRegistrant.CandidateBlacklist
	CandidateBlacklist := make([]gogistryProto.RegistryInfo, len(blacklist))
	for i, c := range blacklist {
		CandidateBlacklist[i] = c
	}
	c.c2sRegistrant.CandidateBlacklist <- CandidateBlacklist
}
