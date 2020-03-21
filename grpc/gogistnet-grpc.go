package grpc

import (
	"github.com/yindaheng98/gogisnet/grpc/client"
	"github.com/yindaheng98/gogisnet/grpc/option"
	"github.com/yindaheng98/gogisnet/grpc/server"
)

//DefaultServerOption returns the default option for server initialization.
func DefaultServerOption() server.Option {
	return server.DefaultOption()
}

//DefaultServerListenerOption returns the default option for the port listen when server running.
func DefaultServerListenerOption() option.ListenerOption {
	return option.DefaultListenerOption()
}

//DefaultServerOption returns the default option for client initialization.
func DefaultClientOption() client.Option {
	return client.DefaultOption()
}

//NewServer initialize a server and returns its pointer.
func NewServer(ServerInfoOption option.ServerInfoOption, option server.Option) *server.Server {
	return server.New(ServerInfoOption, option)
}

//NewServer initialize a client and returns its pointer.
func NewClient(ClientInfoOption option.ClientInfoOption, option client.Option) *client.Client {
	return client.New(ClientInfoOption, option)
}
