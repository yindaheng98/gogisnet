package client

import (
	"github.com/yindaheng98/gogisnet/client"
	pb "github.com/yindaheng98/gogisnet/example/grpc/protocol/protobuf"
	grpcServiceClient "github.com/yindaheng98/gogisnet/example/grpc/protocol/registrant"
	"github.com/yindaheng98/gogistry/example/CandidateList"
)

type Client struct {
	*client.Client
}

func New(ClientInfo *pb.ClientInfo, option Option) *Client {
	ServiceOption, GRPCOption := option.ServiceOption, option.GRPCOption
	S2CClient := grpcServiceClient.NewS2CClient(GRPCOption)
	ServiceOption.RequestProto = S2CClient.NewRequestProtocol()
	ServiceOption.CandidateList = CandidateList.NewPingerCandidateList(3, S2CClient.NewC2SPINGer(), 1e9, option.initServer, 1e9, 10)
	return &Client{client.New(ClientInfo, ServiceOption)}
}
