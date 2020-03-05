package client

import (
	"context"
	pb "github.com/yindaheng98/gogisnet/example/grpc/protocol/protobuf"
	"github.com/yindaheng98/gogistry/protocol"
	"google.golang.org/grpc"
)

type S2CClient struct {
	connections *connections
	CallOption  []grpc.CallOption
}

func NewS2CClient(option GRPCClientOption) *S2CClient {
	return &S2CClient{newConnections(option.DialOption, option.MaxDialHoldDuration),
		option.CallOption}
}

func (c *S2CClient) GetClient(addr string) (client pb.S2CServiceClient, err error) {
	conn, err := c.connections.GetClientConn(addr)
	if err != nil {
		return
	}
	client = pb.NewS2CServiceClient(conn)
	return
}

func (c *S2CClient) NewRequestProtocol() C2SRequestProtocol {
	return C2SRequestProtocol{clients: c, CallOption: c.CallOption}
}

type C2SRequestProtocol struct {
	clients    *S2CClient
	CallOption []grpc.CallOption
}

func (p C2SRequestProtocol) Request(ctx context.Context, requestChan <-chan protocol.TobeSendRequest, responseChan chan<- protocol.ReceivedResponse) {
	var tobeSendRequest protocol.TobeSendRequest
	select {
	case tobeSendRequest = <-requestChan: //先取请求
	case <-ctx.Done():
		return
	}
	request, option := tobeSendRequest.Request, tobeSendRequest.Option

	C2SRequest, err := pb.C2SRequestPack(request) //封装请求
	if err != nil {
		responseChan <- protocol.ReceivedResponse{Error: err}
		return
	}

	client, err := p.clients.GetClient(option.(*pb.RequestSendOption).Addr) //从请求中取出地址得到一个客户端
	if err != nil {
		responseChan <- protocol.ReceivedResponse{Error: err}
		return
	}

	C2SResponse, err := client.Poll(ctx, C2SRequest, p.CallOption...) //发出请求
	if err != nil {
		responseChan <- protocol.ReceivedResponse{Error: err}
		return
	}

	response, err := C2SResponse.Unpack() //解包响应
	if err != nil {
		responseChan <- protocol.ReceivedResponse{Error: err}
		return
	}

	responseChan <- protocol.ReceivedResponse{Response: *response}
}
