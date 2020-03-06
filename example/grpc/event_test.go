package grpc

import (
	"github.com/yindaheng98/gogisnet/example/grpc/server"
	"github.com/yindaheng98/gogisnet/protocol"
)

func PutServerEvent(s *server.Server, logger func(string)) {
	s.Events.ClientNewConnection.AddHandler(func(info protocol.ClientInfo) {
		logger("-->ClientNewConnection-->" + info.String())
	})
	s.Events.ClientDisconnection.AddHandler(func(info protocol.ClientInfo) {
		logger("-->ClientDisconnection-->" + info.String())
	})
}
