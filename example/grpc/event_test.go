package grpc

import (
	"github.com/yindaheng98/gogisnet/example/grpc/client"
	"github.com/yindaheng98/gogisnet/example/grpc/server"
	"github.com/yindaheng98/gogisnet/protocol"
	gogistryProto "github.com/yindaheng98/gogistry/protocol"
)

func PutServerEvent(s *server.Server, logger func(string)) {
	s.Events.ServerNewConnection.AddHandler(func(info protocol.ServerInfo) {
		logger("-->ServerNewConnection-->" + info.GetServerID())
	})
	s.Events.ServerNewConnection.Enable()
	s.Events.ServerDisconnection.AddHandler(func(info protocol.ServerInfo) {
		logger("-->ServerDisconnection-->" + info.GetServerID())
	})
	s.Events.ServerDisconnection.Enable()
	s.Events.ClientNewConnection.AddHandler(func(info protocol.ClientInfo) {
		logger("-->ClientNewConnection-->" + info.GetClientID())
	})
	s.Events.ClientNewConnection.Enable()
	s.Events.ClientDisconnection.AddHandler(func(info protocol.ClientInfo) {
		logger("-->ClientDisconnection-->" + info.GetClientID())
	})
	s.Events.ClientDisconnection.Enable()
}

func PutClientEvent(c *client.Client, logger func(string)) {
	c.Events.NewConnection.AddHandler(func(info protocol.S2CInfo) {
		logger("--NewConnection-->" + info.String())
	})
	c.Events.NewConnection.Enable()
	c.Events.Disconnection.AddHandler(func(info protocol.S2CInfo, err error) {
		logger("--Disconnection-->" + info.GetServerID() + ",ERROR: " + err.Error())
	})
	c.Events.Disconnection.Enable()
	c.Events.UpdateConnection.AddHandler(func(info protocol.S2CInfo) {
		logger("--UpdateConnection-->" + info.GetServerID())
	})
	c.Events.UpdateConnection.Enable()
	c.Events.Retry.AddHandler(func(request gogistryProto.TobeSendRequest, err error) {
		logger("--Retry-->" + request.Request.RegistrantInfo.GetRegistrantID() + ",ERROR: " + err.Error())
	})
	c.Events.Retry.Enable()
}
