package gogisnet

import (
	"context"
	"fmt"
	protobuf "github.com/yindaheng98/gogisnet/example/grpc/protocol/protobuf"
	"github.com/yindaheng98/gogisnet/protocol"
	exampleProto "github.com/yindaheng98/gogistry/example/protocol"
	gogistryProto "github.com/yindaheng98/gogistry/protocol"
	"sync"
	"testing"
	"time"
)

var initS2CRegistry gogistryProto.RegistryInfo
var initS2SRegistry gogistryProto.RegistryInfo

func InitS2SRegistry() {
	ServerInfo := protobuf.ServerInfoPointer{
		ServerInfo: &protobuf.ServerInfo{
			ID:   "UNDEF",
			Type: "UNDEF",
		},
	}
	initS2SRegistry = protocol.S2SInfo{
		ServerInfo:         ServerInfo,
		ResponseSendOption: exampleProto.ResponseSendOption{Timestamp: time.Now()},
		RequestSendOption:  exampleProto.RequestSendOption{RequestAddr: "UNDEF", Timestamp: time.Now()},
		Candidates:         []protocol.S2SInfo{},
		S2CInfo: protocol.S2CInfo{
			ServerInfo:        ServerInfo,
			RequestSendOption: exampleProto.ResponseSendOption{Timestamp: time.Now()},
			Candidates:        []protocol.S2CInfo{},
		},
	}
}

func ServerTest(t *testing.T, ctx context.Context, id int, Type string, wg *sync.WaitGroup) {
	s := fmt.Sprintf("ServerTest(%d)----", id)
	option := DefaultServerOption(initS2SRegistry,
		exampleProto.NewChanNetResponseProtocol(),
		exampleProto.NewChanNetRequestProtocol(),
		exampleProto.NewChanNetResponseProtocol())
	option.S2SRegistryOption.RequestSendOption = exampleProto.RequestSendOption{
		RequestAddr: option.S2SRegistryOption.ResponseProto.(exampleProto.ChanNetResponseProtocol).GetAddr(),
		Timestamp:   time.Now()}
	option.S2SRegistrantOption.ResponseSendOption = exampleProto.ResponseSendOption{Timestamp: time.Now()}
	option.S2CRegistryOption.RequestSendOption = exampleProto.RequestSendOption{
		RequestAddr: option.S2CRegistryOption.ResponseProto.(exampleProto.ChanNetResponseProtocol).GetAddr(),
		Timestamp:   time.Now()}
	server := NewServer(
		protobuf.ServerInfoPointer{
			ServerInfo: &protobuf.ServerInfo{
				ID:   fmt.Sprintf("SERVER_%d", id),
				Type: Type,
			},
		}, option)
	server.Events.ServerNewConnection.AddHandler(func(info protocol.ServerInfo) {
		t.Log(s + fmt.Sprintf("%s---->ServerNewConnection--->%s", server.GetServerInfo(), info))
	})
	server.Events.ServerNewConnection.Enable()
	server.Events.ServerDisconnection.AddHandler(func(info protocol.ServerInfo) {
		t.Log(s + fmt.Sprintf("%s---->ServerDisconnection--->%s", server.GetServerInfo(), info))
	})
	server.Events.ServerDisconnection.Enable()
	server.Events.ClientNewConnection.AddHandler(func(info protocol.ClientInfo) {
		t.Log(s + fmt.Sprintf("%s---->ClientNewConnection--->%s", server.GetServerInfo(), info))
	})
	server.Events.ClientNewConnection.Enable()
	server.Events.ClientDisconnection.AddHandler(func(info protocol.ClientInfo) {
		t.Log(s + fmt.Sprintf("%s---->ClientDisconnection--->%s", server.GetServerInfo(), info))
	})
	server.Events.ClientDisconnection.Enable()
	initS2SRegistry = server.GetS2SInfo()
	initS2CRegistry = server.GetS2CInfo()
	go func() {
		go func() {
			defer wg.Done()
			fmt.Println(s + fmt.Sprintf("%s is going to start.", server.GetServerInfo()))
			server.Run(ctx)
			t.Log(s + fmt.Sprintf("%s stopped itself.", server.GetServerInfo()))
		}()
		go func() {
			for i := 0; ; i++ {
				select {
				case <-time.After(1e9):
					C2SConnections := server.GetC2SConnections()
					S2SConnections := server.GetS2SConnections()
					t.Log(s + fmt.Sprintf("Check(%d)-S2C:%d\nS2S:%d,%s", i, len(C2SConnections), len(S2SConnections), S2SConnections))
				case <-ctx.Done():
					return
				}
			}
		}()
	}()
}

func ClientTest(t *testing.T, ctx context.Context, id int, Type string, wg *sync.WaitGroup) {
	s := fmt.Sprintf("ClientTest(%d)----", id)
	option := DefaultClientOption(initS2CRegistry, exampleProto.NewChanNetRequestProtocol())
	option.ResponseSendOption = exampleProto.ResponseSendOption{Timestamp: time.Now()}
	client := NewClient(
		protobuf.ClientInfoPointer{
			ClientInfo: &protobuf.ClientInfo{
				ID:   fmt.Sprintf("CLIENT_%d", id),
				Type: Type,
			},
		}, option)
	client.Events.NewConnection.AddHandler(func(info protocol.S2CInfo) {
		t.Log(s + fmt.Sprintf("%s---->NewConnection---->%s", client.GetClientInfo(), info.ServerInfo))
	})
	client.Events.NewConnection.Enable()
	client.Events.Disconnection.AddHandler(func(info protocol.S2CInfo, err error) {
		t.Log(s + fmt.Sprintf("%s---->Disconnection---->%s, error:%s", client.GetClientInfo(), info.ServerInfo, err))
	})
	client.Events.Disconnection.Enable()
	go func() {
		go func() {
			defer wg.Done()
			t.Log(s + fmt.Sprintf("%s is going to start.", client.GetClientInfo()))
			client.Run(ctx)
			fmt.Println(s + fmt.Sprintf("%s stopped itself.", client.GetClientInfo()))
		}()
		go func() {
			for i := 0; ; i++ {
				select {
				case <-time.After(1e9):
					t.Log(s + fmt.Sprintf("Check(%d)-C2S:%d", i, len(client.GetS2CConnections())))
				case <-ctx.Done():
					return
				}
			}
		}()
	}()
}

const SERVERN = 10
const CLIENTN = 60

func TestServerClient(t *testing.T) {
	ctx := context.Background()
	InitS2SRegistry()
	Type := "TYPE_0"
	serverWG := new(sync.WaitGroup)
	serverWG.Add(SERVERN)
	serverCtx, cancelServer := context.WithCancel(ctx)
	for i := 0; i < SERVERN; i++ {
		ServerTest(t, serverCtx, i, Type, serverWG)
	}
	time.Sleep(2e9)
	clientWG := new(sync.WaitGroup)
	clientWG.Add(CLIENTN)
	clientCtx, cancelClient := context.WithCancel(ctx)
	for i := 0; i < CLIENTN; i++ {
		ClientTest(t, clientCtx, i, Type, clientWG)
		time.Sleep(0.5e9)
	}
	time.Sleep(10e9)
	cancelServer()
	cancelClient()
	serverWG.Wait()
	clientWG.Wait()
}
