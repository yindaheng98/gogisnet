package client

import (
	"context"
	pb "github.com/yindaheng98/gogisnet/example/grpc/protocol/protobuf"
	"github.com/yindaheng98/gogistry/protocol"
	"google.golang.org/grpc"
	"time"
)

type S2CClient struct {
	connections *connections
	CallOption  []grpc.CallOption
}

func NewS2CClient(option GRPCClientOption) *S2CClient {
	return &S2CClient{newConnections(option.DialOption, option.MaxDialHoldDuration),
		option.CallOption}
}

func (c *S2CClient) getClient(addr string) (client pb.S2CServiceClient, err error) {
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
	defer func() { recover() }()

	C2SRequest, err := pb.C2SRequestPack(request) //封装请求
	if err != nil {
		responseChan <- protocol.ReceivedResponse{Error: err}
		return
	}

	client, err := p.clients.getClient(option.(*pb.RequestSendOption).Addr) //从请求中取出地址得到一个客户端
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

type C2SPINGer struct {
	clients    *S2CClient
	CallOption []grpc.CallOption
}

func (c *S2CClient) NewC2SPINGer() *C2SPINGer {
	return &C2SPINGer{clients: c, CallOption: c.CallOption}
}
func (p C2SPINGer) PING(ctx context.Context, info protocol.RegistryInfo) (ok bool) {
	ok = false
	client, err := p.clients.getClient(info.GetRequestSendOption().(*pb.RequestSendOption).Addr)
	if err != nil {
		return
	}
	_, err = client.PING(ctx, &pb.Timestamp{Timestamp: uint64(time.Now().UnixNano())}, p.CallOption...)
	if err != nil {
		return
	}
	return true
}
