package registrant

import (
	"context"
	pb "github.com/yindaheng98/gogisnet/example/grpc/protocol/protobuf"
	"github.com/yindaheng98/gogistry/protocol"
	"google.golang.org/grpc"
	"time"
)

type S2CRegistrant struct {
	pool       *ConnectionPool
	CallOption []grpc.CallOption
}

func NewS2CRegistrant(option GRPCRegistrantOption) *S2CRegistrant {
	return &S2CRegistrant{NewConnectionPool(option.DialOption, option.MaxDialHoldDuration),
		option.CallOption}
}

func (c *S2CRegistrant) getClient(addr string) (client pb.S2CRegistryClient, err error) {
	conn, err := c.pool.GetClientConn(addr)
	if err != nil {
		return
	}
	client = pb.NewS2CRegistryClient(conn)
	return
}

func (c *S2CRegistrant) NewRequestProtocol() C2SRequestProtocol {
	return C2SRequestProtocol{clients: c, CallOption: c.CallOption}
}

type C2SRequestProtocol struct {
	clients    *S2CRegistrant
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
	clients    *S2CRegistrant
	CallOption []grpc.CallOption
}

func (c *S2CRegistrant) NewC2SPINGer() *C2SPINGer {
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
