package server

import (
	"google.golang.org/grpc"
	"net"
)

type GRPCServerOption struct {
	Listener   net.Listener
	InitOption []grpc.ServerOption
	BufferLen  uint64
}

func DefaultOption(ListenNetwork, ListenAddr string) (option GRPCServerOption, err error) {
	option = GRPCServerOption{
		Listener:   nil,
		InitOption: nil,
		BufferLen:  100,
	}
	Listener, err := net.Listen(ListenNetwork, ListenAddr)
	if err != nil {
		option.Listener = Listener
	}
	return
}
