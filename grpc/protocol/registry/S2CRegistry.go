package registry

import (
	"context"
	pb "github.com/yindaheng98/gogisnet/grpc/protocol/protobuf"
	"google.golang.org/grpc"
)

type S2CRegistryServer struct {
	*grpc.Server
	service *s2cRegistryService
}

func NewS2CRegistryServer(option GRPCRegistryOption) *S2CRegistryServer {
	Server := grpc.NewServer(option.ServerOption...)
	Service := newS2CService(option.BufferSize)
	pb.RegisterS2CRegistryServer(Server, Service)
	return &S2CRegistryServer{Server, Service}
}

func (s *S2CRegistryServer) NewResponseProtocol() ResponseProtocol {
	return s.service.newResponseProtocol()
}

type s2cRegistryService struct {
	chanpairService
}

func newS2CService(bufn uint64) *s2cRegistryService {
	return &s2cRegistryService{newChanpairService(bufn)}
}

func (s *s2cRegistryService) Poll(ctx context.Context, S2SRequest *pb.C2SRequest) (*pb.S2CResponse, error) {
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
