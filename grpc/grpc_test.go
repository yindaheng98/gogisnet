package grpc

import (
	"context"
	"fmt"
	pb "github.com/yindaheng98/gogisnet/grpc/protocol/protobuf"
	"sync"
	"testing"
	"time"
)

func GetAddr(port uint16) string {
	return fmt.Sprintf("[::1]:%d", port)
}

var initS2SServer *pb.S2SInfo

func ServerTest(t *testing.T, ctx context.Context, S2SPort, S2CPort, GQPort uint16) (err error) {
	ServerID := fmt.Sprintf("Server-%d/%d", S2SPort, S2CPort)
	ServerInfo := &pb.ServerInfo{
		ServerID:    ServerID,
		ServiceType: "Hello World Service",
	}
	S2SBoardCastAddr, S2CBoardCastAddr, GQBoardCastAddr := GetAddr(S2SPort), GetAddr(S2CPort), GetAddr(GQPort)
	option := DefaultServerOption()
	option.ServiceOption.S2SRegistryOption.BoardCastAddr = S2SBoardCastAddr
	option.ServiceOption.S2CRegistryOption.BoardCastAddr = S2CBoardCastAddr
	option.InitServer.S2CInfo.ServerInfo.GraphQueryBroadCastAddr = GQBoardCastAddr
	if initS2SServer != nil {
		option.InitServer = initS2SServer
	}
	s := NewServer(ServerInfo, option)
	if S2SServer, err := pb.S2SInfoPack(s.GetS2SInfo()); err == nil {
		initS2SServer = S2SServer
	}
	if S2CServer, err := pb.S2CInfoPack(s.GetS2CInfo()); err == nil {
		initS2CServer = S2CServer
	}
	PutServerEvent(s, func(s string) { t.Log(ServerID + s) })
	listenerOption := DefaultServerListenerOption()
	listenerOption.S2SListenAddr = S2SBoardCastAddr
	listenerOption.S2CListenAddr = S2CBoardCastAddr
	listenerOption.GraphQueryListenAddr = GQBoardCastAddr
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
		ClientID:    ClientID,
		ServiceType: "Hello World Service",
	}
	option := DefaultClientOption()
	option.InitServer = initS2CServer
	c := NewClient(ClientInfo, option)
	PutClientEvent(c, func(s string) { fmt.Println(ClientID + s) })
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

const SERVERN = 5
const CLIENTN = 30

func Test(t *testing.T) {
	ctx := context.Background()

	serverWG := new(sync.WaitGroup)
	serverWG.Add(SERVERN)
	serverCtx, cancelServer := context.WithCancel(ctx)
	for i := uint16(0); i < SERVERN; i++ {
		go func(i uint16) {
			defer serverWG.Done()
			err := ServerTest(t, serverCtx, 4240+i*3, 4240+i*3+1, 4240+i*3+2)
			if err != nil {
				t.Log(err)
			}
		}(i)
		time.Sleep(0.2e9)
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

	time.Sleep(100e9)
	cancelServer()
	time.Sleep(1e9)
	cancelClient()
	serverWG.Wait()
	clientWG.Wait()
	time.Sleep(1e9)
}
