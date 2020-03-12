package registry

import (
	"google.golang.org/grpc"
)

type GRPCServerOption struct {
	InitOption []grpc.ServerOption
	BufferLen  uint64
}

func DefaultOption() GRPCServerOption {
	return GRPCServerOption{
		InitOption: nil,
		BufferLen:  100,
	}
}
