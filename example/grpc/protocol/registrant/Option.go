package registrant

import (
	"google.golang.org/grpc"
	"time"
)

type GRPCClientOption struct {
	DialOption          []grpc.DialOption
	CallOption          []grpc.CallOption
	MaxDialHoldDuration time.Duration //保持连接的最大时长
}

func DefaultOption() GRPCClientOption {
	return GRPCClientOption{
		DialOption:          []grpc.DialOption{grpc.WithInsecure()},
		CallOption:          nil,
		MaxDialHoldDuration: 1e9,
	}
}
