package server

import (
	"context"
	"github.com/yindaheng98/gogisnet/grpc/option"
	"github.com/yindaheng98/gogisnet/grpc/protocol/graph"
	pb "github.com/yindaheng98/gogisnet/grpc/protocol/protobuf"
	"github.com/yindaheng98/gogisnet/grpc/protocol/registrant"
	"github.com/yindaheng98/gogisnet/grpc/protocol/registry"
	"github.com/yindaheng98/gogisnet/server"
	"net"
)

//Server is gRPC-implemented gogisnet server
type Server struct {
	*server.Server
	grpcS2SRegistry  *registry.S2SRegistryServer
	grpcS2CRegistry  *registry.S2CRegistryServer
	graphQueryServer *graph.GraphQueryServer
}

//New returns a instance of Server
func New(ServerInfoOption option.ServerInfoOption, opt Option) *Server {
	GRPCOption := opt.GRPCOption
	grpcS2SRegistry := registry.NewS2SRegistryServer(GRPCOption.S2SRegistryOption)   //注册中心gRPC初始化
	grpcS2CRegistry := registry.NewS2CRegistryServer(GRPCOption.S2CRegistryOption)   //注册中心gRPC初始化
	grpcS2SRegistrant := registrant.NewS2SRegistrant(GRPCOption.S2SRegistrantOption) //注册器gRPC初始化

	//设置协议
	InitServer := &pb.S2SInfo{}
	opt.InitServerOption.PutOption(InitServer)
	InitServerMessage, err := InitServer.Unpack()
	if err != nil {
		panic(err)
	}
	ServiceOption := server.DefaultOption(*InitServerMessage,
		grpcS2SRegistry.NewResponseProtocol(),
		grpcS2SRegistrant.NewRequestProtocol(),
		grpcS2CRegistry.NewResponseProtocol()) //服务设置初始化

	//生成注册中心设置
	opt.ServiceOption.S2CRegistryOption.PutOption(&ServiceOption.S2CRegistryOption) //服务设置修改
	opt.ServiceOption.S2SRegistryOption.PutOption(&ServiceOption.S2SRegistryOption) //服务设置修改

	//生成注册器设置
	S2SRegistrantOption := opt.ServiceOption.S2SRegistrantOption
	S2SRegistrantOption.PutOption(&ServiceOption.S2SRegistrantOption) //服务设置修改
	ServiceOption.S2SRegistrantOption.CandidateList =
		grpcS2SRegistrant.NewPingerCandidateList(InitServer, S2SRegistrantOption.CandidateListOption)
	ServiceOption.GraphQuerySendOption = &pb.GraphQuerySendOption{Addr: opt.ServiceOption.GraphQueryOption.BoardCastAddr}

	//初始化服务器
	if ServerInfoOption.ServerID == "undefined" {
		ServerInfoOption.ServerID = "SERVER-" + option.RandomString(64)
	}
	ServerInfo := &pb.ServerInfo{}
	ServerInfoOption.PutOption(ServerInfo)
	s := server.New(ServerInfo, ServiceOption)

	//初始化GraphQuery
	GraphQueryOption := GRPCOption.GraphQueryOption
	s.GraphQueryProtocol = graph.NewGraphQueryClient(GraphQueryOption.GraphQueryClientOption).NewGraphQueryProtocol()
	return &Server{
		Server:           s,
		grpcS2SRegistry:  grpcS2SRegistry,
		grpcS2CRegistry:  grpcS2CRegistry,
		graphQueryServer: graph.NewGraphQueryServer(s, GraphQueryOption.GraphQueryServerOption),
	}
}

//Run the server until <-ctx.Done() is done
func (s *Server) Run(ctx context.Context, option option.ListenerOption) (err error) {
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
