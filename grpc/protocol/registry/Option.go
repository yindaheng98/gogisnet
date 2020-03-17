package registry

import (
	"google.golang.org/grpc"
)

type GRPCRegistryOption struct {
	InitOption []grpc.ServerOption
	BufferLen  uint64
}

func DefaultOption() GRPCRegistryOption {
	return GRPCRegistryOption{
		InitOption: nil,
		BufferLen:  100,
	}
}
