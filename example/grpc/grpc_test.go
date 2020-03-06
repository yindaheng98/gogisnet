package grpc

import (
	"context"
	"fmt"
	"github.com/yindaheng98/gogisnet/example/grpc/client"
	pb "github.com/yindaheng98/gogisnet/example/grpc/protocol/protobuf"
	"github.com/yindaheng98/gogisnet/example/grpc/server"
	"sync"
	"testing"
	"time"
)

func GetAddr(port uint16) string {
	return fmt.Sprintf("[::1]:%d", port)
}

var initS2SServer *pb.S2SInfo

func ServerTest(ctx context.Context, S2SPort, S2CPort uint16) (err error) {
	ServerID := fmt.Sprintf("Server-%d/%d", S2SPort, S2CPort)
	ServerInfo := &pb.ServerInfo{
		ServerID:       ServerID,
		ServiceType:    "Hello World Service",
		AdditionalInfo: "",
	}
	S2SBoardCastAddr, S2CBoardCastAddr := GetAddr(S2SPort), GetAddr(S2CPort)
	option, err := server.DefaultOption(S2SBoardCastAddr, S2CBoardCastAddr, initS2SServer)
	if err != nil {
		return
	}
	s := server.New(ServerInfo, option)
	if S2SServer, err := pb.S2SInfoPack(s.GetS2SInfo()); err == nil {
		initS2SServer = S2SServer
	}
	if S2CServer, err := pb.S2CInfoPack(s.GetS2CInfo()); err == nil {
		initS2CServer = S2CServer
	}
	PutServerEvent(s, func(s string) {
		fmt.Println(ServerID + s)
	})
	listenerOption := server.DefaultListenerOption()
	listenerOption.S2SListenAddr = S2SBoardCastAddr
	listenerOption.S2CListenAddr = S2CBoardCastAddr
	errChan := make(chan error, 1)
	go func() {
		fmt.Println(ServerID + " started.")
		errChan <- s.Run(ctx, listenerOption)
	}()
	select {
	case err = <-errChan:
		fmt.Println(ServerID + fmt.Sprintf(" exited by error: %s.", err))
	case <-ctx.Done():
		fmt.Println(ServerID + " exited by context.")
	}
	return
}

var initS2CServer *pb.S2CInfo

func ClientTest(ctx context.Context, id uint16) (err error) {
	ClientID := fmt.Sprintf("Client-%02d", id)
	ClientInfo := &pb.ClientInfo{
		ClientID:       ClientID,
		ServiceType:    "Hello World Service",
		AdditionalInfo: "",
	}
	option, err := client.DefaultOption(initS2CServer)
	if err != nil {
		return
	}
	c := client.New(ClientInfo, option)
	c.SetWatchdogTimeDelta(3e9)
	okChan := make(chan bool, 1)
	go func() {
		fmt.Println(ClientID + " started.")
		c.Run(ctx)
		okChan <- true
	}()
	select {
	case <-okChan:
		fmt.Println(ClientID + " exited itself.")
	case <-ctx.Done():
		fmt.Println(ClientID + " exited by context.")
	}
	return
}

const SERVERN = 10
const CLIENTN = 60

func Test(t *testing.T) {
	ctx := context.Background()

	serverWG := new(sync.WaitGroup)
	serverWG.Add(SERVERN)
	serverCtx, cancelServer := context.WithCancel(ctx)
	for i := uint16(0); i < SERVERN; i++ {
		go func(i uint16) {
			defer serverWG.Done()
			err := ServerTest(serverCtx, 4240+i*2, 4240+i*2+1)
			if err != nil {
				t.Log(err)
			}
		}(i)
	}
	time.Sleep(2e9)

	clientWG := new(sync.WaitGroup)
	clientWG.Add(CLIENTN)
	clientCtx, cancelClient := context.WithCancel(ctx)
	for i := uint16(0); i < CLIENTN; i++ {
		go func(i uint16) {
			defer clientWG.Done()
			err := ClientTest(clientCtx, i)
			if err != nil {
				t.Log(err)
			}
		}(i)
	}

	time.Sleep(10e9)
	cancelServer()
	time.Sleep(10e9)
	cancelClient()
	serverWG.Wait()
	clientWG.Wait()
	time.Sleep(1e9)
}
