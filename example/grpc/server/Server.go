package server

import (
	"context"
	pb "github.com/yindaheng98/gogisnet/example/grpc/protocol/protobuf"
	grpcServiceClient "github.com/yindaheng98/gogisnet/example/grpc/protocol/registrant"
	grpcServiceServer "github.com/yindaheng98/gogisnet/example/grpc/protocol/registry"
	"github.com/yindaheng98/gogisnet/server"
	"github.com/yindaheng98/gogistry/example/CandidateList"
	"net"
)

type Server struct {
	*server.Server
	grpcS2SServer *grpcServiceServer.S2SServiceServer
	grpcS2CServer *grpcServiceServer.S2CServiceServer
}

func New(ServerInfo *pb.ServerInfo, option Option) *Server {
	ServiceOption, GRPCOption := option.ServiceOption, option.GRPCOption

	//注册中心初始化
	S2SServiceServer := grpcServiceServer.NewS2SServiceServer(GRPCOption.S2SServerOption)
	S2CServiceServer := grpcServiceServer.NewS2CServiceServer(GRPCOption.S2CServerOption)
	ServiceOption.S2SRegistryOption.ResponseProto = S2SServiceServer.NewResponseProtocol()
	ServiceOption.S2CRegistryOption.ResponseProto = S2CServiceServer.NewResponseProtocol()

	//注册器初始化
	S2SClient := grpcServiceClient.NewS2SClient(GRPCOption.S2SClientOption)
	RequestProtocol := S2SClient.NewRequestProtocol()
	ServiceOption.S2SRegistrantOption.RequestProto = RequestProtocol
	ServiceOption.S2SRegistrantOption.CandidateList = CandidateList.NewPingerCandidateList(3, S2SClient.NewS2SPINGer(), 1e9, option.initServer, 1e9, 10)

	return &Server{
		Server:        server.New(ServerInfo, ServiceOption),
		grpcS2SServer: S2SServiceServer,
		grpcS2CServer: S2CServiceServer,
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
		S2SErrChan <- s.grpcS2SServer.Serve(Listener)
	}()

	S2CErrChan := make(chan error, 1)
	go func() {
		Listener, err := net.Listen(option.S2CListenNetwork, option.S2CListenAddr)
		if err != nil {
			S2CErrChan <- err
		}
		S2CErrChan <- s.grpcS2CServer.Serve(Listener)
	}()

	select {
	case err = <-S2SErrChan:
	case err = <-S2CErrChan:
	case <-ctx.Done():
	}
	s.grpcS2SServer.GracefulStop()
	s.grpcS2CServer.GracefulStop()
	return
}
