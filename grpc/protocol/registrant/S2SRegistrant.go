package registrant

import (
	"context"
	pb "github.com/yindaheng98/gogisnet/grpc/protocol/protobuf"
	"github.com/yindaheng98/gogistry/example/CandidateList"
	"github.com/yindaheng98/gogistry/protocol"
	"google.golang.org/grpc"
	"time"
)

//S2CRegistrant defines the gRPC client for S2SRegistrant
type S2SRegistrant struct {
	pool *ConnectionPool

	//CallOption is the option when call Poll or PING in gRPC
	CallOption []grpc.CallOption
}

//NewS2SRegistrant implements a S2SRegistrant and returns its pointer.
func NewS2SRegistrant(option GRPCRegistrantOption) *S2SRegistrant {
	return &S2SRegistrant{NewConnectionPool(option.DialOption, option.MaxDialHoldDuration),
		option.CallOption}
}

//NewCandidateList implements a RequestProtocol using current S2SRegistrant and returns it.
func (c *S2SRegistrant) NewRequestProtocol() S2SRequestProtocol {
	return S2SRequestProtocol{clients: c, CallOption: c.CallOption}
}

//NewCandidateList implements a CandidateList.PingerCandidateList using current S2SRegistrant and returns its pointer.
func (c *S2SRegistrant) NewPingerCandidateList(initServer *pb.S2SInfo, option PingerCandidateListOption) *CandidateList.PingerCandidateList {
	InitServer, _ := initServer.Unpack()
	l := CandidateList.NewPingerCandidateList(option.Size, s2sPINGer{clients: c}, option.MaxPingTimeout, InitServer)
	l.DefaultRetryN = option.DefaultRetryN
	l.DefaultTimeout = option.DefaultTimeout
	return l
}

func (c *S2SRegistrant) getClient(addr string) (client pb.S2SRegistryClient, err error) {
	conn, err := c.pool.GetClientConn(addr)
	if err != nil {
		return
	}
	client = pb.NewS2SRegistryClient(conn)
	return
}

type S2SRequestProtocol struct {
	clients    *S2SRegistrant
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
	defer func() { recover() }()

	S2SRequest, err := pb.S2SRequestPack(request) //封装请求
	if err != nil {
		responseChan <- protocol.ReceivedResponse{Error: err}
		return
	}

	client, err := p.clients.getClient(option.(*pb.RequestSendOption).Addr) //从请求中取出地址得到一个客户端
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

type s2sPINGer struct {
	clients *S2SRegistrant
}

func (p s2sPINGer) PING(ctx context.Context, info protocol.RegistryInfo) (ok bool) {
	ok = false
	client, err := p.clients.getClient(info.GetRequestSendOption().(*pb.RequestSendOption).Addr)
	if err != nil {
		return
	}
	_, err = client.PING(ctx, &pb.Timestamp{Timestamp: uint64(time.Now().UnixNano())}, p.clients.CallOption...)
	if err != nil {
		return
	}
	return true
}
