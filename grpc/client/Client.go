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
	ServiceOption, GRPCOption := option.ServiceOption, option.GRPCOption
	S2CRegistrant := registrant.NewS2CRegistrant(GRPCOption)
	ServiceOption.RequestProto = S2CRegistrant.NewRequestProtocol()
	ServiceOption.CandidateList = S2CRegistrant.NewCandidateList(option.initServer, option.CandidateListOption)
	return &Client{client.New(ClientInfo, ServiceOption)}
}
