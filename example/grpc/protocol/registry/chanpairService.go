package registry

import (
	"context"
	"errors"
	pb "github.com/yindaheng98/gogisnet/example/grpc/protocol/protobuf"
	"github.com/yindaheng98/gogistry/protocol"
	"time"
)

type chanpair struct {
	ctx          context.Context
	requestChan  chan protocol.Request
	responseChan chan protocol.Response
}

type chanpairService struct {
	pairChan chan chanpair
}

func newChanpairService(bufn uint64) chanpairService {
	return chanpairService{make(chan chanpair, bufn)}
}

func (chanpairService) PING(context.Context, *pb.Timestamp) (*pb.Timestamp, error) {
	now := time.Now().UnixNano()
	return &pb.Timestamp{Timestamp: uint64(now)}, nil
}

//用于接收外界发来的请求然后发回响应
func (s *chanpairService) Poll(ctx context.Context, request protocol.Request) (protocol.Response, error) {
	pair := chanpair{
		ctx:          ctx,
		requestChan:  make(chan protocol.Request, 1),
		responseChan: make(chan protocol.Response, 1),
	}
	pair.requestChan <- request
	s.pairChan <- pair
	select {
	case response := <-pair.responseChan:
		return response, nil
	case <-ctx.Done():
		return protocol.Response{}, errors.New("poll exited")
	}
}

//用于内部接收外界发来的请求并将内部生成的响应发回外部
func (s *chanpairService) response(ctx context.Context, requestChan chan<- protocol.ReceivedRequest, responseChan <-chan protocol.TobeSendResponse) {
	select {
	case <-ctx.Done():
		return
	case pair := <-s.pairChan:
		req := <-pair.requestChan
		defer func() { recover() }()
		requestChan <- protocol.ReceivedRequest{Request: req}
		select {
		case <-pair.ctx.Done():
			return
		case res := <-responseChan:
			pair.responseChan <- res.Response
		}
	}

}

//返回一个自身服务器对应的响应协议
func (s *chanpairService) newResponseProtocol() ResponseProtocol {
	return ResponseProtocol{s}
}

//响应协议
type ResponseProtocol struct {
	service *chanpairService
}

func (p ResponseProtocol) Response(ctx context.Context, requestChan chan<- protocol.ReceivedRequest, responseChan <-chan protocol.TobeSendResponse) {
	p.service.response(ctx, requestChan, responseChan)
}
