package protocol

import (
	"context"
	"fmt"
	"github.com/yindaheng98/gogisnet/example/grpc/protocol/client"
	pb "github.com/yindaheng98/gogisnet/example/grpc/protocol/protobuf"
	"github.com/yindaheng98/gogisnet/example/grpc/protocol/server"
	"github.com/yindaheng98/gogistry/protocol"
	"google.golang.org/grpc"
	"net"
	"sync"
	"testing"
	"time"
)

func CreateS2SResponseProtocol(t *testing.T, ctx context.Context, port uint16) protocol.ResponseProtocol {
	s := server.NewS2SServiceServer(server.DefaultOption())
	go func() {
		errChan := make(chan error, 1)
		go func() {
			t.Log(fmt.Sprintf("Server:%d listen on %s.", port, GetAddr(port)))
			listener, err := net.Listen("tcp", GetAddr(port))
			if err != nil {
				errChan <- err
				return
			}
			errChan <- s.Serve(listener)
		}()
		select {
		case err := <-errChan:
			t.Log(fmt.Sprintf("Server:%d stopped itself. err:%s.", port, err))
		case <-ctx.Done():
			s.GracefulStop()
			t.Log(fmt.Sprintf("Server:%d stopped by context.", port))
		}
	}()
	return s.NewResponseProtocol()
}

func CreateS2SRequestProtocol() protocol.RequestProtocol {
	c := client.NewS2SClient(client.DefaultOption())
	return c.NewRequestProtocol()
}

func NewS2SInfo(i uint64) *pb.S2SInfo {
	return &pb.S2SInfo{
		ServerInfo: &pb.ServerInfo{
			ServerID:       fmt.Sprintf("ServerID:%d", i),
			ServiceType:    fmt.Sprintf("ServiceType:%d", i),
			AdditionalInfo: fmt.Sprintf("AdditionalInfo:%d", i),
		},
		ResponseSendOption: &pb.ResponseSendOption{Option: nil},
		RequestSendOption: &pb.RequestSendOption{
			Option: nil,
			Addr:   "Unknown",
		},
		Candidates: nil,
		S2CInfo: &pb.S2CInfo{
			ServerInfo: &pb.ServerInfo{
				ServerID:       fmt.Sprintf("ServerID:%d", i),
				ServiceType:    fmt.Sprintf("ServiceType:%d", i),
				AdditionalInfo: fmt.Sprintf("AdditionalInfo:%d", i),
			},
			RequestSendOption: &pb.RequestSendOption{
				Option: nil,
				Addr:   "Unknown",
			},
			Candidates: nil,
		},
	}
}

func S2SRequestProtocolTest(p protocol.RequestProtocol, ctx context.Context, port uint16, i uint64) {
	S2SRequest := pb.S2SRequest{
		S2SInfo:    NewS2SInfo(i),
		Disconnect: false,
	}
	request, err := S2SRequest.Unpack()
	if err != nil {
		fmt.Println(err)
		return
	}
	RequestProtocolTest(p, ctx, port, *request, func(s string) {
		fmt.Println(fmt.Sprintf("S2SRequestProtocolTest-%02d->", i) + s)
	})
}

func S2SResponseProtocolTest(p protocol.ResponseProtocol, ctx context.Context, i uint64) {
	S2SResponse := pb.S2SResponse{
		S2SInfo: NewS2SInfo(i),
		Timeout: 0,
		Reject:  false,
	}
	response, err := S2SResponse.Unpack()
	if err != nil {
		fmt.Println(err)
		return
	}
	ResponseProtocolTest(p, ctx, *response, func(s string) {
		fmt.Println(fmt.Sprintf("S2SResponseProtocolTest-%02d->", i) + s)
	})
}

func S2SOneTurnTest(ctx context.Context, port uint16,
	reqp protocol.RequestProtocol, resp protocol.ResponseProtocol) {
	wg := new(sync.WaitGroup)
	wg.Add(REPEATN * 2)
	for i := uint64(1); i <= REPEATN; i++ {
		go func(i uint64) {
			defer wg.Done()
			S2SRequestProtocolTest(reqp, ctx, port, i)
		}(i)
		go func(i uint64) {
			defer wg.Done()
			S2SResponseProtocolTest(resp, ctx, i)
		}(i)
	}
	wg.Wait()
}

func TestS2S(t *testing.T) {
	grpc.WithInsecure()

	reqps := make([]protocol.RequestProtocol, REQPN)
	for i := 0; i < REQPN; i++ {
		reqps[i] = CreateS2SRequestProtocol()
	}

	resps := make([]protocol.ResponseProtocol, RESPN)
	ctx := context.Background()
	serverCtx, cancelServer := context.WithCancel(ctx)
	defer cancelServer()
	ports := make([]uint16, RESPN)
	for j := 0; j < RESPN; j++ {
		ports[j] = uint16(8000 + j)
		resps[j] = CreateS2SResponseProtocol(t, serverCtx, ports[j])
	}

	testCtx, cancelTest := context.WithCancel(ctx)
	wg := new(sync.WaitGroup)
	wg.Add(REQPN * RESPN)
	defer cancelTest()
	for i := 0; i < REQPN; i++ {
		for j := 0; j < RESPN; j++ {
			go func(i, j int) {
				defer wg.Done()
				S2SOneTurnTest(testCtx, ports[j], reqps[i], resps[j])
			}(i, j)
		}
	}
	wg.Wait()
	fmt.Println("Test finished.")
	cancelTest()
	cancelServer()
	time.Sleep(1e9)
}
