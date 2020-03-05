package server

import (
	"context"
	pb "github.com/yindaheng98/gogisnet/example/grpc/protocol/protobuf"
	"google.golang.org/grpc"
)

type S2CServiceServer struct {
	*grpc.Server
	service *s2cService
}

func NewS2CServiceServer(option GRPCServerOption) *S2CServiceServer {
	Server := grpc.NewServer(option.InitOption...)
	Service := newS2CService(option.BufferLen)
	pb.RegisterS2CServiceServer(Server, Service)
	return &S2CServiceServer{Server, Service}
}

func (s *S2CServiceServer) NewResponseProtocol() ResponseProtocol {
	return s.service.newResponseProtocol()
}

type s2cService struct {
	chanpairService
}

func newS2CService(bufn uint64) *s2cService {
	return &s2cService{newChanpairService(bufn)}
}

func (s *s2cService) Poll(ctx context.Context, S2SRequest *pb.C2SRequest) (*pb.S2CResponse, error) {
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
