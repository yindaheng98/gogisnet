package client

import (
	"github.com/yindaheng98/gogisnet/grpc/option"
	"github.com/yindaheng98/gogisnet/grpc/protocol/registrant"
)

//Option is the options for gRPC gogisnet client
type Option struct {
	ServiceOption    option.RegistrantOption         `yaml:"ServiceOption" usage:"Option for gogisnet service."`
	GRPCOption       registrant.GRPCRegistrantOption `yaml:"GRPCOption" usage:"Option for gRPC server in registry and gRPC client in registrant."`
	InitServerOption option.S2CInfoOption            `yaml:"InitServerOption" usage:"The information about the first server that the client should connect."`
}

//DefaultOption returns a default Option
func DefaultOption() Option {
	ServiceOption := option.DefaultRegistrantOption()
	ServiceOption.RegistryN = 1
	return Option{
		ServiceOption:    ServiceOption,
		GRPCOption:       registrant.DefaultOption(),
		InitServerOption: option.DefaultS2CInfoOption(),
	}
}
