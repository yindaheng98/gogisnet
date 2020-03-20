package registry

import (
	"context"
	pb "github.com/yindaheng98/gogisnet/grpc/protocol/protobuf"
	"google.golang.org/grpc"
)

type S2SRegistryServer struct {
	*grpc.Server
	service *s2sRegistryService
}

func NewS2SRegistryServer(option GRPCRegistryOption) *S2SRegistryServer {
	Server := grpc.NewServer(option.ServerOption...)
	Service := newS2SService(option.BufferSize)
	pb.RegisterS2SRegistryServer(Server, Service)
	return &S2SRegistryServer{Server, Service}
}

func (s *S2SRegistryServer) NewResponseProtocol() ResponseProtocol {
	return s.service.newResponseProtocol()
}

type s2sRegistryService struct {
	chanpairService
}

func newS2SService(bufn uint64) *s2sRegistryService {
	return &s2sRegistryService{newChanpairService(bufn)}
}

func (s *s2sRegistryService) Poll(ctx context.Context, S2SRequest *pb.S2SRequest) (*pb.S2SResponse, error) {
	request, err := S2SRequest.Unpack()
	if err != nil {
		return nil, err
	}
	response, err := s.chanpairService.Poll(ctx, *request)
	if err != nil {
		return nil, err
	}
	return pb.S2SResponsePack(response)
}
