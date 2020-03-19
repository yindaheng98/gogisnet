package grpc

import (
	"github.com/yindaheng98/gogisnet/grpc/client"
	pb "github.com/yindaheng98/gogisnet/grpc/protocol/protobuf"
	"github.com/yindaheng98/gogisnet/grpc/server"
)

//DefaultServerOption returns the default option for server initialization.
func DefaultServerOption() server.Option {
	return server.DefaultOption()
}

//DefaultServerListenerOption returns the default option for the port listen when server running.
func DefaultServerListenerOption() server.ListenerOption {
	return server.DefaultListenerOption()
}

//DefaultServerOption returns the default option for client initialization.
func DefaultClientOption() client.Option {
	return client.DefaultOption()
}

//NewServer initialize a server and returns its pointer.
func NewServer(ServerInfo *pb.ServerInfo, option server.Option) *server.Server {
	return server.New(ServerInfo, option)
}

//NewServer initialize a client and returns its pointer.
func NewClient(ClientInfo *pb.ClientInfo, option client.Option) *client.Client {
	return client.New(ClientInfo, option)
}
