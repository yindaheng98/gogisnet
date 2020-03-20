package client

import (
	"github.com/yindaheng98/gogisnet/client"
	pb "github.com/yindaheng98/gogisnet/grpc/protocol/protobuf"
	"github.com/yindaheng98/gogisnet/grpc/protocol/registrant"
)

type Client struct {
	*client.Client
}

func New(ClientInfo *pb.ClientInfo, option Option) *Client {
	InitServer, err := option.InitServer.Unpack()
	if err != nil {
		panic(err)
	}
	GRPCOption := option.GRPCOption

	grpcS2CRegistrant := registrant.NewS2CRegistrant(GRPCOption) //注册器gRPC初始化

	ServiceOption := client.DefaultOption(*InitServer, grpcS2CRegistrant.NewRequestProtocol())

	ServiceOption.CandidateList = grpcS2CRegistrant.NewPingerCandidateList(option.InitServer, option.ServiceOption.CandidateListOption)

	return &Client{client.New(ClientInfo, ServiceOption)}
}
