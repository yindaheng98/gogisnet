package protocol

import (
	"context"
	"fmt"
	pb "github.com/yindaheng98/gogisnet/example/grpc/protocol/protobuf"
	"github.com/yindaheng98/gogisnet/example/grpc/protocol/registrant"
	"github.com/yindaheng98/gogisnet/example/grpc/protocol/registry"
	"github.com/yindaheng98/gogistry/protocol"
	"google.golang.org/grpc"
	"net"
	"sync"
	"testing"
	"time"
)

func CreateS2CResponseProtocol(t *testing.T, ctx context.Context, port uint16) protocol.ResponseProtocol {
	s := registry.NewS2CServiceServer(registry.DefaultOption())
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

func CreateS2CRequestProtocol() protocol.RequestProtocol {
	c := registrant.NewS2CClient(registrant.DefaultOption())
	return c.NewRequestProtocol()
}

func S2CRequestProtocolTest(p protocol.RequestProtocol, ctx context.Context, port uint16, i uint64) {
	C2SRequest := pb.C2SRequest{
		C2SInfo: &pb.C2SInfo{
			ClientInfo: &pb.ClientInfo{
				ClientID:       fmt.Sprintf("ClientID:%d", i),
				ServiceType:    fmt.Sprintf("ServiceType:%d", i),
				AdditionalInfo: []byte(fmt.Sprintf("AdditionalInfo:%d", i)),
			},
			ResponseSendOption: &pb.ResponseSendOption{Option: nil},
		},
		Disconnect: false,
	}
	request, err := C2SRequest.Unpack()
	if err != nil {
		fmt.Println(err)
		return
	}
	RequestProtocolTest(p, ctx, port, *request, func(s string) {
		fmt.Println(fmt.Sprintf("S2CRequestProtocolTest-%02d->", i) + s)
	})
}

func S2CResponseProtocolTest(p protocol.ResponseProtocol, ctx context.Context, i uint64) {
	S2CResponse := pb.S2CResponse{
		S2CInfo: &pb.S2CInfo{
			ServerInfo: &pb.ServerInfo{
				ServerID:       fmt.Sprintf("ServerID:%d", i),
				ServiceType:    fmt.Sprintf("ServiceType:%d", i),
				AdditionalInfo: []byte(fmt.Sprintf("AdditionalInfo:%d", i)),
			},
			RequestSendOption: &pb.RequestSendOption{
				Option: nil,
				Addr:   "Unknown",
			},
			Candidates: nil,
		},
		Timeout: 0,
		Reject:  false,
	}
	response, err := S2CResponse.Unpack()
	if err != nil {
		fmt.Println(err)
		return
	}
	ResponseProtocolTest(p, ctx, *response, func(s string) {
		fmt.Println(fmt.Sprintf("S2CResponseProtocolTest-%02d->", i) + s)
	})
}

const REPEATN = 10

func S2COneTurnTest(ctx context.Context, port uint16,
	reqp protocol.RequestProtocol, resp protocol.ResponseProtocol) {
	wg := new(sync.WaitGroup)
	wg.Add(REPEATN * 2)
	for i := uint64(1); i <= REPEATN; i++ {
		go func(i uint64) {
			defer wg.Done()
			S2CRequestProtocolTest(reqp, ctx, port, i)
		}(i)
		go func(i uint64) {
			defer wg.Done()
			S2CResponseProtocolTest(resp, ctx, i)
		}(i)
	}
	wg.Wait()
}

const REQPN = 10
const RESPN = 10

func TestS2C(t *testing.T) {
	grpc.WithInsecure()

	reqps := make([]protocol.RequestProtocol, REQPN)
	for i := 0; i < REQPN; i++ {
		reqps[i] = CreateS2CRequestProtocol()
	}

	resps := make([]protocol.ResponseProtocol, RESPN)
	ctx := context.Background()
	serverCtx, cancelServer := context.WithCancel(ctx)
	defer cancelServer()
	ports := make([]uint16, RESPN)
	for j := 0; j < RESPN; j++ {
		ports[j] = uint16(8000 + j)
		resps[j] = CreateS2CResponseProtocol(t, serverCtx, ports[j])
	}

	testCtx, cancelTest := context.WithCancel(ctx)
	wg := new(sync.WaitGroup)
	wg.Add(REQPN * RESPN)
	defer cancelTest()
	for i := 0; i < REQPN; i++ {
		for j := 0; j < RESPN; j++ {
			go func(i, j int) {
				defer wg.Done()
				S2COneTurnTest(testCtx, ports[j], reqps[i], resps[j])
			}(i, j)
		}
	}
	wg.Wait()
	fmt.Println("Test finished.")
	cancelTest()
	cancelServer()
	time.Sleep(1e9)
}
