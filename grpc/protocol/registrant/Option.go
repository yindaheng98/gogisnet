package registrant

import (
	"google.golang.org/grpc"
	"time"
)

//GRPCRegistrantOption is the option for S2CRegistrant and S2SRegistrant
type GRPCRegistrantOption struct {

	//DialOption is the option when dial gRPC connection.
	DialOption []grpc.DialOption `yaml:"-"`

	//CallOption is the option when call gRPC functions.
	CallOption          []grpc.CallOption `yaml:"-"`
	MaxDialHoldDuration time.Duration     `yaml:"MaxDialHoldDuration" usage:"How many ms can a gRPC connection hold at most."`
}

//DefaultOption returns a default GRPCRegistrantOption
func DefaultOption() GRPCRegistrantOption {
	return GRPCRegistrantOption{
		DialOption:          []grpc.DialOption{grpc.WithInsecure()},
		CallOption:          nil,
		MaxDialHoldDuration: 1e9,
	}
}

//PingerCandidateListOption is the option for PingerCandidateList used in S2CRegistrant and S2SRegistrant
type PingerCandidateListOption struct {

	//Default timeout
	DefaultTimeout time.Duration

	//Default retryN
	DefaultRetryN uint64

	//How many elements can the CandidateList store at most
	Size uint64

	//How many ms can a PING hold at most before received a response
	MaxPingTimeout time.Duration
}

func DefaultPingerCandidateListOption() PingerCandidateListOption {
	return PingerCandidateListOption{
		DefaultTimeout: 1e9,
		DefaultRetryN:  10,
		Size:           8,
		MaxPingTimeout: 1e9,
	}
}
