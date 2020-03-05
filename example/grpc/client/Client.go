package client

import (
	"github.com/yindaheng98/gogisnet/client"
	grpcClient "github.com/yindaheng98/gogisnet/example/grpc/protocol/client"
	pb "github.com/yindaheng98/gogisnet/example/grpc/protocol/protobuf"
)

type Client struct {
	*client.Client
}

func New(ClientInfo *pb.ClientInfo, option Option) *Client {
	ServiceOption, GRPCOption := option.ServiceOption, option.GRPCOption
	S2CClient := grpcClient.NewS2CClient(GRPCOption)
	ServiceOption.RequestProto = S2CClient.NewRequestProtocol()
	return &Client{client.New(ClientInfo, ServiceOption)}
}
