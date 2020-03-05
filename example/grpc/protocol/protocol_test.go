package protocol

import (
	"context"
	"fmt"
	pb "github.com/yindaheng98/gogisnet/example/grpc/protocol/protobuf"
	"github.com/yindaheng98/gogistry/protocol"
)

func GetAddr(port uint16) string {
	return fmt.Sprintf("[::1]:%d", port)
}

func RequestProtocolTest(p protocol.RequestProtocol, ctx context.Context, port uint16,
	request protocol.Request, logger func(string)) {
	requestChan := make(chan protocol.TobeSendRequest, 1)
	responseChan := make(chan protocol.ReceivedResponse, 1)
	go p.Request(ctx, requestChan, responseChan)
	requestChan <- protocol.TobeSendRequest{
		Request: request,
		Option: &pb.RequestSendOption{
			Option: nil,
			Addr:   GetAddr(port),
		},
	}
	logger(fmt.Sprintf("We have sent %s a request %s.", GetAddr(port), request))
	select {
	case res := <-responseChan:
		response, err := res.Response, res.Error
		if err != nil {
			logger(fmt.Sprintf("But an error occurred in response: %s.", err))
		} else {
			logger(fmt.Sprintf("And received a response: %s", response))
		}
	case <-ctx.Done():
		logger("Exited by context.")
	}
}

func ResponseProtocolTest(p protocol.ResponseProtocol, ctx context.Context,
	response protocol.Response, logger func(string)) {
	responseChan := make(chan protocol.TobeSendResponse, 1)
	requestChan := make(chan protocol.ReceivedRequest, 1)
	go p.Response(ctx, requestChan, responseChan)
	select {
	case req := <-requestChan:
		request, err := req.Request, req.Error
		if err != nil {
			logger(fmt.Sprintf("An error occurred in request: %s.", err))
			return
		} else {
			logger(fmt.Sprintf("We have received a request: %s", request))
		}
		responseChan <- protocol.TobeSendResponse{
			Response: response,
			Option: &pb.ResponseSendOption{
				Option: nil,
			},
		}
		logger(fmt.Sprintf("And we have sent back a response: %s", response))
	case <-ctx.Done():
		logger("Exited by context.")
	}
}
