package server

import (
	"context"
	pb "github.com/yindaheng98/gogisnet/example/grpc/protocol/protobuf"
	"github.com/yindaheng98/gogisnet/example/grpc/protocol/registrant"
	"github.com/yindaheng98/gogisnet/example/grpc/protocol/registry"
	"github.com/yindaheng98/gogisnet/server"
	"github.com/yindaheng98/gogistry/example/CandidateList"
	"net"
)

type Server struct {
	*server.Server
	grpcS2SRegistry *registry.S2SRegistryServer
	grpcS2CRegistry *registry.S2CRegistryServer
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
	ServiceOption.S2SRegistrantOption.CandidateList = CandidateList.NewPingerCandidateList(3, S2SRegistrant.NewS2SPINGer(), 1e9, option.initServer, 1e9, 10)

	return &Server{
		Server:          server.New(ServerInfo, ServiceOption),
		grpcS2SRegistry: S2SRegistryServer,
		grpcS2CRegistry: S2CRegistryServer,
	}
}

func (s *Server) Run(ctx context.Context, option ListenerOption) (err error) {
	go s.Server.Run(ctx)

	S2SErrChan := make(chan error, 1)
	go func() {
		Listener, err := net.Listen(option.S2SListenNetwork, option.S2SListenAddr)
		if err != nil {
			S2SErrChan <- err
		}
		S2SErrChan <- s.grpcS2SRegistry.Serve(Listener)
	}()

	S2CErrChan := make(chan error, 1)
	go func() {
		Listener, err := net.Listen(option.S2CListenNetwork, option.S2CListenAddr)
		if err != nil {
			S2CErrChan <- err
		}
		S2CErrChan <- s.grpcS2CRegistry.Serve(Listener)
	}()

	select {
	case err = <-S2SErrChan:
	case err = <-S2CErrChan:
	case <-ctx.Done():
	}
	s.grpcS2SRegistry.GracefulStop()
	s.grpcS2CRegistry.GracefulStop()
	return
}
