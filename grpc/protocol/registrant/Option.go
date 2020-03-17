package registrant

import (
	"google.golang.org/grpc"
	"time"
)

//GRPCRegistrantOption is the option for S2CRegistrant and S2SRegistrant
type GRPCRegistrantOption struct {

	//Option when dial gRPC connection
	DialOption []grpc.DialOption

	//Option when call gRPC functions
	CallOption []grpc.CallOption

	//How many ms can a gRPC connection hold at most
	MaxDialHoldDuration time.Duration //保持连接的最大时长
}

//Option for CandidateList in S2CRegistrant and S2SRegistrant
type CandidateListOption struct {

	//Default timeout
	InitTimeout time.Duration

	//Default retryN
	InitRetryN uint64

	//How many elements can the CandidateList store at most
	Size uint64

	//How many ms can a PING hold at most before received a response
	MaxPingTimeout time.Duration
}

//DefaultOption returns a default GRPCRegistrantOption
func DefaultOption() GRPCRegistrantOption {
	return GRPCRegistrantOption{
		DialOption:          []grpc.DialOption{grpc.WithInsecure()},
		CallOption:          nil,
		MaxDialHoldDuration: 1e9,
	}
}

//DefaultCandidateListOption returns a default CandidateListOption
func DefaultCandidateListOption() CandidateListOption {
	return CandidateListOption{
		InitTimeout:    1e9,
		InitRetryN:     10,
		Size:           8,
		MaxPingTimeout: 1e9,
	}
}
