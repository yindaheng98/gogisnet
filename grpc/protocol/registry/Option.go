package registry

import (
	"google.golang.org/grpc"
)

//GRPCRegistryOption is the options for gRPC server in Registry
type GRPCRegistryOption struct {
	//ServerOption is the option for initialize gRPC server.
	ServerOption []grpc.ServerOption `yaml:"-"`
	BufferSize   uint64              `yaml:"BufferSize" usage:"Size of the receiving buffer."`
}

func DefaultOption() GRPCRegistryOption {
	return GRPCRegistryOption{
		ServerOption: nil,
		BufferSize:   100,
	}
}
