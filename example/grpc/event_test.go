package grpc

import (
	"fmt"
	"github.com/yindaheng98/gogisnet/example/grpc/client"
	"github.com/yindaheng98/gogisnet/example/grpc/server"
	"github.com/yindaheng98/gogisnet/protocol"
	gogistryProto "github.com/yindaheng98/gogistry/protocol"
)

func PutServerEvent(s *server.Server, logger func(string)) {

	s.Events.S2SRegistrantEvent.NewConnection.AddHandler(func(info protocol.S2SInfo) {
		logger("-->ServerNewConnection-->" + info.GetServerID())
	})
	s.Events.S2SRegistrantEvent.NewConnection.Enable()
	s.Events.S2SRegistrantEvent.Disconnection.AddHandler(func(info protocol.S2SInfo, err error) {
		logger("-->ServerDisconnection-->" + info.GetServerID())
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
		logger("--NewConnection-->" + info.GetServerID())
	})
	c.Events.NewConnection.Enable()
	c.Events.Disconnection.AddHandler(func(info protocol.S2CInfo, err error) {
		logger("--Disconnection-->" + info.GetServerID() + fmt.Sprintf(",ERROR: %s", err))
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
