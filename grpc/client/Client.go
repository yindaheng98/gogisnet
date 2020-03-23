package client

import (
	"github.com/yindaheng98/gogisnet/client"
	"github.com/yindaheng98/gogisnet/grpc/option"
	pb "github.com/yindaheng98/gogisnet/grpc/protocol/protobuf"
	"github.com/yindaheng98/gogisnet/grpc/protocol/registrant"
	"github.com/yindaheng98/gogisnet/server"
)

type Client struct {
	*client.Client
}

func New(ClientInfoOption option.ClientInfoOption, opt Option) *Client {
	grpcS2CRegistrant := registrant.NewS2CRegistrant(opt.GRPCOption) //注册器gRPC初始化

	InitServer := &pb.S2CInfo{}
	opt.InitServerOption.PutOption(InitServer)
	InitServerMessage, err := InitServer.Unpack()
	if err != nil {
		panic(err)
	}
	ServiceOption := client.DefaultOption(*InitServerMessage, grpcS2CRegistrant.NewRequestProtocol()) //服务设置初始化
	opt.ServiceOption.PutOption((*server.RegistrantOption)(&ServiceOption))                           //服务设置修改
	ServiceOption.CandidateList = grpcS2CRegistrant.NewPingerCandidateList(InitServer, opt.ServiceOption.CandidateListOption)

	if ClientInfoOption.ClientID == "undefined" {
		ClientInfoOption.ClientID = "CLIENT-" + option.RandomString(64)
	}
	ClientInfo := &pb.ClientInfo{}
	ClientInfoOption.PutOption(ClientInfo)
	return &Client{client.New(ClientInfo, ServiceOption)} //客户端初始化
}
