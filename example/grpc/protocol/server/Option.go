package server

import "google.golang.org/grpc"

type GRPCServerOption struct {
	InitOption []grpc.ServerOption
	BufferLen  uint64
}
