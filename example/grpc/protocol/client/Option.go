package client

import (
	"google.golang.org/grpc"
	"time"
)

type GRPCClientOption struct {
	DialOption          []grpc.DialOption
	CallOption          []grpc.CallOption
	MaxDialHoldDuration time.Duration //保持连接的最大时长
}
