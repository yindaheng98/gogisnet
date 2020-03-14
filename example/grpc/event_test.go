package grpc

import (
	"fmt"
	"github.com/yindaheng98/gogisnet/example/grpc/client"
	"github.com/yindaheng98/gogisnet/example/grpc/server"
	"github.com/yindaheng98/gogisnet/protocol"
	gogistryProto "github.com/yindaheng98/gogistry/protocol"
)

func PutServerEvent(s *server.Server, logger func(string)) {

	check := func() string {
		C2SConnections := s.GetC2SConnections()
		S2SConnections := s.GetS2SConnections()
		return fmt.Sprintf("\nS2C:%d,%s\nS2S:%d,%s\n", len(C2SConnections), C2SConnections, len(S2SConnections), S2SConnections)
	}
	s.Events.S2SRegistrantEvent.NewConnection.AddHandler(func(info protocol.S2SInfo) {
		logger("-->ServerNewConnection-->" + info.GetServerID() + check())
	})
	s.Events.S2SRegistrantEvent.NewConnection.Enable()
	s.Events.S2SRegistrantEvent.Disconnection.AddHandler(func(info protocol.S2SInfo, err error) {
		logger("-->ServerDisconnection-->" + info.GetServerID() + check())
	})
	s.Events.S2SRegistrantEvent.Disconnection.Enable()

	s.Events.S2SRegistryEvent.NewConnection.AddHandler(func(info protocol.S2SInfo) {
		logger("<--ServerNewConnection<--" + info.GetServerID())
	})
	s.Events.S2SRegistryEvent.NewConnection.Enable()
	s.Events.S2SRegistryEvent.Disconnection.AddHandler(func(info protocol.S2SInfo) {
		logger("<--ServerDisconnection<--" + info.GetServerID())
	})
	s.Events.S2SRegistryEvent.Disconnection.Enable()

	s.Events.ClientNewConnection.AddHandler(func(info protocol.ClientInfo) {
		logger("-->ClientNewConnection-->" + info.GetClientID() + check())
	})
	s.Events.ClientNewConnection.Enable()
	s.Events.ClientDisconnection.AddHandler(func(info protocol.ClientInfo) {
		logger("-->ClientDisconnection-->" + info.GetClientID() + check())
	})
	s.Events.ClientDisconnection.Enable()

}

func PutClientEvent(c *client.Client, logger func(string)) {
	check := func() string {
		S2CConnections := c.GetS2CConnections()
		return fmt.Sprintf("\nC2S:%d,%s", len(S2CConnections), S2CConnections)
	}
	c.Events.NewConnection.AddHandler(func(info protocol.S2CInfo) {
		logger("--NewConnection-->" + info.GetServerID() + check())
	})
	c.Events.NewConnection.Enable()
	c.Events.Disconnection.AddHandler(func(info protocol.S2CInfo, err error) {
		logger("--Disconnection-->" + info.GetServerID() + fmt.Sprintf(",ERROR: %s", err) + check())
	})
	c.Events.Disconnection.Enable()
	c.Events.UpdateConnection.AddHandler(func(info protocol.S2CInfo) {
		logger("--UpdateConnection-->" + info.GetServerID())
	})
	c.Events.UpdateConnection.Enable()
	c.Events.Retry.AddHandler(func(request gogistryProto.TobeSendRequest, err error) {
		logger("--Retry-->" + request.Option.String() + fmt.Sprintf(",ERROR: %s", err))
	})
	c.Events.Retry.Enable()
}
