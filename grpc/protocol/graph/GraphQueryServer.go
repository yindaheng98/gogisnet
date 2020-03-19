package graph

import (
	"context"
	pb "github.com/yindaheng98/gogisnet/grpc/protocol/protobuf"
	"github.com/yindaheng98/gogisnet/server"
	"google.golang.org/grpc"
)

//A grpc server for GraphQuery service.
//About GraphQueryProtocol: https://godoc.org/github.com/yindaheng98/gogisnet/message#GraphQueryProtocol
type GraphQueryServer struct {
	*grpc.Server
	service *graphQueryService
}

//NewGraphQueryServer returns the pointer to a GraphQueryServer.
func NewGraphQueryServer(s *server.Server, Option GraphQueryServerOption) *GraphQueryServer {
	Server := grpc.NewServer(Option...)
	Service := &graphQueryService{s}
	pb.RegisterGraphQueryServer(Server, Service)
	return &GraphQueryServer{Server, Service}
}

type graphQueryService struct {
	server *server.Server
}

func (s *graphQueryService) Query(context.Context, *pb.Empty) (*pb.GraphQueryInfo, error) {
	return pb.GraphQueryInfoPack(s.server.GetGraphQueryInfo())
}
