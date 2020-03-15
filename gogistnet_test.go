package gogisnet

import (
	"context"
	"fmt"
	protobuf "github.com/yindaheng98/gogisnet/example/grpc/protocol/protobuf"
	"github.com/yindaheng98/gogisnet/message"
	exampleProto "github.com/yindaheng98/gogistry/example/protocol"
	"sync"
	"testing"
	"time"
)

var initS2CRegistry message.S2CInfo
var initS2SRegistry message.S2SInfo

func InitS2SRegistry() {
	ServerInfo := &protobuf.ServerInfo{
		ServerID:       "UNDEF",
		ServiceType:    "UNDEF",
		AdditionalInfo: []byte("UNDEF"),
	}
	initS2SRegistry = message.S2SInfo{
		ServerInfo:         ServerInfo,
		ResponseSendOption: exampleProto.ResponseSendOption{Timestamp: time.Now()},
		RequestSendOption:  exampleProto.RequestSendOption{RequestAddr: "UNDEF", Timestamp: time.Now()},
		Candidates:         []message.S2SInfo{},
		S2CInfo: message.S2CInfo{
			ServerInfo:        ServerInfo,
			RequestSendOption: exampleProto.ResponseSendOption{Timestamp: time.Now()},
			Candidates:        []message.S2CInfo{},
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
		&protobuf.ServerInfo{
			ServerID:    fmt.Sprintf("SERVER_%d", id),
			ServiceType: Type,
		}, option)

	check := func() string {
		C2SConnections := server.GetC2SConnections()
		S2SConnections := server.GetS2SConnections()
		return fmt.Sprintf("\nS2C:%d,%s\nS2S:%d,%s\n", len(C2SConnections), C2SConnections, len(S2SConnections), S2SConnections)
	}
	server.Events.ServerNewConnection.AddHandler(func(info message.ServerInfo) {
		t.Log(s + fmt.Sprintf("%s---->ServerNewConnection--->%s", server.GetServerInfo(), info) + check())
	})
	server.Events.ServerNewConnection.Enable()
	server.Events.ServerDisconnection.AddHandler(func(info message.ServerInfo) {
		t.Log(s + fmt.Sprintf("%s---->ServerDisconnection--->%s", server.GetServerInfo(), info) + check())
	})
	server.Events.ServerDisconnection.Enable()
	server.Events.ClientNewConnection.AddHandler(func(info message.ClientInfo) {
		t.Log(s + fmt.Sprintf("%s---->ClientNewConnection--->%s", server.GetServerInfo(), info) + check())
	})
	server.Events.ClientNewConnection.Enable()
	server.Events.ClientDisconnection.AddHandler(func(info message.ClientInfo) {
		t.Log(s + fmt.Sprintf("%s---->ClientDisconnection--->%s", server.GetServerInfo(), info) + check())
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
	}()
}

func ClientTest(t *testing.T, ctx context.Context, id int, Type string, wg *sync.WaitGroup) {
	s := fmt.Sprintf("ClientTest(%d)----", id)
	option := DefaultClientOption(initS2CRegistry, exampleProto.NewChanNetRequestProtocol())
	option.ResponseSendOption = exampleProto.ResponseSendOption{Timestamp: time.Now()}
	client := NewClient(
		&protobuf.ClientInfo{
			ClientID:    fmt.Sprintf("CLIENT_%d", id),
			ServiceType: Type,
		}, option)

	check := func() string {
		S2CConnections := client.GetS2CConnections()
		return fmt.Sprintf("\nC2S:%d,%s", len(S2CConnections), S2CConnections)
	}
	client.Events.NewConnection.AddHandler(func(info message.S2CInfo) {
		t.Log(s + fmt.Sprintf("%s---->NewConnection---->%s", client.GetClientInfo(), info.ServerInfo) + check())
	})
	client.Events.NewConnection.Enable()
	client.Events.Disconnection.AddHandler(func(info message.S2CInfo, err error) {
		t.Log(s + fmt.Sprintf("%s---->Disconnection---->%s, error:%s", client.GetClientInfo(), info.ServerInfo, err) + check())
	})
	client.Events.Disconnection.Enable()
	go func() {
		go func() {
			defer wg.Done()
			t.Log(s + fmt.Sprintf("%s is going to start.", client.GetClientInfo()))
			client.Run(ctx)
			fmt.Println(s + fmt.Sprintf("%s stopped itself.", client.GetClientInfo()))
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
