package graph

import (
	"github.com/yindaheng98/gogisnet/grpc/protocol/registrant"
	"google.golang.org/grpc"
)

//GraphQueryOption contains the options for GraphQueryServer and GraphQueryClient
type GraphQueryOption struct {
	GraphQueryClientOption GraphQueryClientOption `yaml:"GraphQueryClientOption" usage:"The options for GraphQueryClient."`

	//GraphQueryServerOption is The options for GraphQueryServer.
	GraphQueryServerOption GraphQueryServerOption `yaml:"-"`
}

//GraphQueryClientOption contains the options for GraphQueryClient
type GraphQueryClientOption registrant.GRPCRegistrantOption

//GraphQueryServerOption contains the options for GraphQueryServer
type GraphQueryServerOption []grpc.ServerOption

//DefaultOption returns a default GraphQueryOption
func DefaultOption() GraphQueryOption {
	return GraphQueryOption{
		GraphQueryClientOption: GraphQueryClientOption(registrant.DefaultOption()),
		GraphQueryServerOption: []grpc.ServerOption{},
	}
}
