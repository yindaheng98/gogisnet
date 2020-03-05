package client

import (
	"context"
	pb "github.com/yindaheng98/gogisnet/example/grpc/protocol/protobuf"
	"github.com/yindaheng98/gogistry/protocol"
	"google.golang.org/grpc"
)

type S2SClient struct {
	connections *connections
	CallOption  []grpc.CallOption
}

func NewS2SClient(option GRPCClientOption) *S2SClient {
	return &S2SClient{newConnections(option.DialOption, option.MaxDialHoldDuration),
		option.CallOption}
}

func (c *S2SClient) GetClient(addr string) (client pb.S2SServiceClient, err error) {
	conn, err := c.connections.GetClientConn(addr)
	if err != nil {
		return
	}
	client = pb.NewS2SServiceClient(conn)
	return
}

func (c *S2SClient) NewRequestProtocol() S2SRequestProtocol {
	return S2SRequestProtocol{clients: c, CallOption: c.CallOption}
}

type S2SRequestProtocol struct {
	clients    *S2SClient
	CallOption []grpc.CallOption
}

func (p S2SRequestProtocol) Request(ctx context.Context, requestChan <-chan protocol.TobeSendRequest, responseChan chan<- protocol.ReceivedResponse) {
	var tobeSendRequest protocol.TobeSendRequest
	select {
	case tobeSendRequest = <-requestChan: //先取请求
	case <-ctx.Done():
		return
	}
	request, option := tobeSendRequest.Request, tobeSendRequest.Option

	S2SRequest, err := pb.S2SRequestPack(request) //封装请求
	if err != nil {
		responseChan <- protocol.ReceivedResponse{Error: err}
		return
	}

	client, err := p.clients.GetClient(option.(*pb.RequestSendOption).Addr) //从请求中取出地址得到一个客户端
	if err != nil {
		responseChan <- protocol.ReceivedResponse{Error: err}
		return
	}

	S2SResponse, err := client.Poll(ctx, S2SRequest, p.CallOption...) //发出请求
	if err != nil {
		responseChan <- protocol.ReceivedResponse{Error: err}
		return
	}

	response, err := S2SResponse.Unpack() //解包响应
	if err != nil {
		responseChan <- protocol.ReceivedResponse{Error: err}
		return
	}

	responseChan <- protocol.ReceivedResponse{Response: *response}
}
