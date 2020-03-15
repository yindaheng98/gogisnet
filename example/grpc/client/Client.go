package client

import (
	"github.com/yindaheng98/gogisnet/client"
	pb "github.com/yindaheng98/gogisnet/example/grpc/protocol/protobuf"
	"github.com/yindaheng98/gogisnet/example/grpc/protocol/registrant"
	"github.com/yindaheng98/gogistry/example/CandidateList"
)

type Client struct {
	*client.Client
}

func New(ClientInfo *pb.ClientInfo, option Option) *Client {
	ServiceOption, GRPCOption := option.ServiceOption, option.GRPCOption
	S2CRegistrant := registrant.NewS2CRegistrant(GRPCOption)
	ServiceOption.RequestProto = S2CRegistrant.NewRequestProtocol()
	ServiceOption.CandidateList = CandidateList.NewPingerCandidateList(3, S2CRegistrant.NewC2SPINGer(), 1e9, option.initServer, 1e9, 10)
	return &Client{client.New(ClientInfo, ServiceOption)}
}
