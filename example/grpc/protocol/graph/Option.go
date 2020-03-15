package graph

import (
	"github.com/yindaheng98/gogisnet/example/grpc/protocol/registrant"
	"google.golang.org/grpc"
)

//GraphQueryOption contains the options for GraphQueryServer and GraphQueryClient
type GraphQueryOption struct {

	//GraphQueryClientOption contains the options for GraphQueryClient
	GraphQueryClientOption GraphQueryClientOption

	//GraphQueryServerOption contains the options for GraphQueryServer
	GraphQueryServerOption GraphQueryServerOption
}

//GraphQueryClientOption contains the options for GraphQueryClient
type GraphQueryClientOption registrant.GRPCRegistrantOption

//GraphQueryServerOption contains the options for GraphQueryServer
type GraphQueryServerOption struct {
	ServerOption  []grpc.ServerOption
	BoardCastAddr string
}

//DefaultOption returns a default GraphQueryOption
func DefaultOption(GraphQueryBoardCastAddr string) GraphQueryOption {
	return GraphQueryOption{
		GraphQueryClientOption: GraphQueryClientOption(registrant.DefaultOption()),
		GraphQueryServerOption: GraphQueryServerOption{BoardCastAddr: GraphQueryBoardCastAddr},
	}
}
