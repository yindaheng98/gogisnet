package gogisnet

import (
	"github.com/yindaheng98/gogisnet/client"
	"github.com/yindaheng98/gogisnet/message"
	"github.com/yindaheng98/gogisnet/server"
	gogistryProto "github.com/yindaheng98/gogistry/protocol"
)

func DefaultClientOption(initRegistry message.S2CInfo, RequestProto gogistryProto.RequestProtocol) client.Option {
	return client.DefaultOption(initRegistry, RequestProto)
}

func NewClient(info message.ClientInfo, option client.Option) *client.Client {
	return client.New(info, option)
}

func DefaultServerOption(initS2SRegistry message.S2SInfo,
	S2SResponseProto gogistryProto.ResponseProtocol,
	S2SRequestProto gogistryProto.RequestProtocol,
	S2CResponseProto gogistryProto.ResponseProtocol) server.Option {
	return server.DefaultOption(initS2SRegistry, S2SResponseProto, S2SRequestProto, S2CResponseProto)
}

func NewServer(info message.ServerInfo, option server.Option) *server.Server {
	return server.New(info, option)
}
