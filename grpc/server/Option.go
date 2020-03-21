package server

import (
	"github.com/yindaheng98/gogisnet/grpc/option"
	"github.com/yindaheng98/gogisnet/grpc/protocol/graph"
	"github.com/yindaheng98/gogisnet/grpc/protocol/registrant"
	"github.com/yindaheng98/gogisnet/grpc/protocol/registry"
)

//Option contains all the options for gRPC gogisnet server
type Option struct {
	ServiceOption    ServiceOption        `yaml:"ServiceOption" usage:"Option for gogisnet service."`
	GRPCOption       GRPCOption           `yaml:"GRPCOption" usage:"Option for gRPC server in registry and gRPC client in registrant."`
	InitServerOption option.S2SInfoOption `yaml:"InitServerOption" usage:"Information about the first server that the client should connect."`
}

//DefaultOption returns a default Option
func DefaultOption() Option {
	return Option{ //初始化
		ServiceOption:    defaultServiceOption(),
		GRPCOption:       defaultGRPCOption(),
		InitServerOption: option.DefaultS2SInfoOption(),
	}
}

//Option for Gogisnet service
type ServiceOption struct {
	S2SRegistryOption   option.RegistryOption   `yaml:"S2SRegistryOption" usage:"Option for S2SRegistry in gogisnet server."`
	S2SRegistrantOption option.RegistrantOption `yaml:"S2SRegistrantOption" usage:"Option for S2SRegistrant in gogisnet server."`
	S2CRegistryOption   option.RegistryOption   `yaml:"S2CRegistryOption" usage:"Option for S2CRegistry in gogisnet server."`
}

//DefaultServiceOption returns a default ServiceOption
func defaultServiceOption() ServiceOption {
	ip := option.GetIP()
	S2SRegistryOption := option.DefaultRegistryOption()
	S2SRegistryOption.BoardCastAddr = ip + ":4241"
	S2CRegistryOption := option.DefaultRegistryOption()
	S2CRegistryOption.BoardCastAddr = ip + ":4240"
	S2CRegistryOption.MaxRegistrants = 16
	return ServiceOption{
		S2SRegistryOption:   S2SRegistryOption,
		S2SRegistrantOption: option.DefaultRegistrantOption(),
		S2CRegistryOption:   S2CRegistryOption,
	}
}

//GRPCOption is the gRPC options for gRPC gogisnet server
type GRPCOption struct {
	S2SRegistryOption   registry.GRPCRegistryOption     `yaml:"S2SRegistryOption" usage:"Option for gRPC server in S2SRegistry."`
	S2CRegistryOption   registry.GRPCRegistryOption     `yaml:"S2CRegistryOption" usage:"Option for gRPC server in S2CRegistry."`
	S2SRegistrantOption registrant.GRPCRegistrantOption `yaml:"S2SRegistrantOption" usage:"Option for gRPC client in S2SRegistrant."`
	GraphQueryOption    graph.GraphQueryOption          `yaml:"GraphQueryOption" usage:"Option for gRPC server in GraphQuery service."`
}

//DefaultGRPCOption returns a default GRPCOption
func defaultGRPCOption() GRPCOption {
	return GRPCOption{
		S2SRegistryOption:   registry.DefaultOption(),
		S2CRegistryOption:   registry.DefaultOption(),
		S2SRegistrantOption: registrant.DefaultOption(),
		GraphQueryOption:    graph.DefaultOption(),
	}
}
