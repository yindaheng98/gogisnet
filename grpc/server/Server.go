package server

import (
	"context"
	"github.com/yindaheng98/gogisnet/grpc/protocol/graph"
	pb "github.com/yindaheng98/gogisnet/grpc/protocol/protobuf"
	"github.com/yindaheng98/gogisnet/grpc/protocol/registrant"
	"github.com/yindaheng98/gogisnet/grpc/protocol/registry"
	"github.com/yindaheng98/gogisnet/server"
	"net"
)

type Server struct {
	*server.Server
	grpcS2SRegistry  *registry.S2SRegistryServer
	grpcS2CRegistry  *registry.S2CRegistryServer
	graphQueryServer *graph.GraphQueryServer
}

func New(ServerInfo *pb.ServerInfo, option Option) *Server {
	ServiceOption, GRPCOption := option.ServiceOption, option.GRPCOption

	//注册中心初始化
	S2SRegistryServer := registry.NewS2SRegistryServer(GRPCOption.S2SRegistryOption)
	S2CRegistryServer := registry.NewS2CRegistryServer(GRPCOption.S2CRegistryOption)
	ServiceOption.S2SRegistryOption.ResponseProto = S2SRegistryServer.NewResponseProtocol()
	ServiceOption.S2CRegistryOption.ResponseProto = S2CRegistryServer.NewResponseProtocol()

	//注册器初始化
	S2SRegistrant := registrant.NewS2SRegistrant(GRPCOption.S2SRegistrantOption)
	RequestProtocol := S2SRegistrant.NewRequestProtocol()
	ServiceOption.S2SRegistrantOption.RequestProto = RequestProtocol
	ServiceOption.S2SRegistrantOption.CandidateList = S2SRegistrant.NewCandidateList(option.initServer, GRPCOption.CandidateListOption)

	GraphQueryOption := GRPCOption.GraphQueryOption
	ServerInfo.GraphQueryAddr = GraphQueryOption.GraphQueryServerOption.BoardCastAddr
	s := server.New(ServerInfo, ServiceOption)
	s.GraphQueryProtocol = graph.NewGraphQueryClient(GraphQueryOption.GraphQueryClientOption).NewGraphQueryProtocol()
	return &Server{
		Server:           s,
		grpcS2SRegistry:  S2SRegistryServer,
		grpcS2CRegistry:  S2CRegistryServer,
		graphQueryServer: graph.NewGraphQueryServer(s, GraphQueryOption.GraphQueryServerOption),
	}
}

func (s *Server) Run(ctx context.Context, option ListenerOption) (err error) {
	go s.Server.Run(ctx)

	S2SErrChan := make(chan error, 1)
	go func() {
		Listener, err := net.Listen(option.S2SListenNetwork, option.S2SListenAddr)
		if err != nil {
			S2SErrChan <- err
			return
		}
		S2SErrChan <- s.grpcS2SRegistry.Serve(Listener)
	}()

	S2CErrChan := make(chan error, 1)
	go func() {
		Listener, err := net.Listen(option.S2CListenNetwork, option.S2CListenAddr)
		if err != nil {
			S2CErrChan <- err
			return
		}
		S2CErrChan <- s.grpcS2CRegistry.Serve(Listener)
	}()

	GQErrChan := make(chan error, 1)
	go func() {
		Listener, err := net.Listen(option.GraphQueryListenNetwork, option.GraphQueryListenAddr)
		if err != nil {
			GQErrChan <- err
			return
		}
		GQErrChan <- s.graphQueryServer.Serve(Listener)
	}()

	select {
	case err = <-S2SErrChan:
	case err = <-S2CErrChan:
	case err = <-GQErrChan:
	case <-ctx.Done():
	}
	s.grpcS2SRegistry.GracefulStop()
	s.grpcS2CRegistry.GracefulStop()
	s.graphQueryServer.GracefulStop()
	return
}
