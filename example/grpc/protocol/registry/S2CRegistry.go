package registry

import (
	"context"
	pb "github.com/yindaheng98/gogisnet/example/grpc/protocol/protobuf"
	"google.golang.org/grpc"
)

type S2CRegistryServer struct {
	*grpc.Server
	service *s2cServiceServer
}

func NewS2CRegistryServer(option GRPCRegistryOption) *S2CRegistryServer {
	Server := grpc.NewServer(option.InitOption...)
	Service := newS2CService(option.BufferLen)
	pb.RegisterS2CRegistryServer(Server, Service)
	return &S2CRegistryServer{Server, Service}
}

func (s *S2CRegistryServer) NewResponseProtocol() ResponseProtocol {
	return s.service.newResponseProtocol()
}

type s2cServiceServer struct {
	chanpairService
}

func newS2CService(bufn uint64) *s2cServiceServer {
	return &s2cServiceServer{newChanpairService(bufn)}
}

func (s *s2cServiceServer) Poll(ctx context.Context, S2SRequest *pb.C2SRequest) (*pb.S2CResponse, error) {
	request, err := S2SRequest.Unpack()
	if err != nil {
		return nil, err
	}
	response, err := s.chanpairService.Poll(ctx, *request)
	if err != nil {
		return nil, err
	}
	return pb.S2CResponsePack(response)
}
