package gogisnet

import (
	"fmt"
	"github.com/yindaheng98/gogisnet/protocol"
	exampleProto "github.com/yindaheng98/gogistry/example/protocol"
	gogistryProto "github.com/yindaheng98/gogistry/protocol"
	"sync"
	"testing"
	"time"
)

type TestServerInfo struct {
	ID   string
	Type string
}

func (info TestServerInfo) GetServerID() string {
	return info.ID
}
func (info TestServerInfo) GetServiceType() string {
	return info.Type
}
func (info TestServerInfo) String() string {
	return fmt.Sprintf("TestServerInfo{ID:%s,Type:%s}", info.ID, info.Type)
}

type TestClientInfo struct {
	ID   string
	Type string
}

func (info TestClientInfo) GetClientID() string {
	return info.ID
}
func (info TestClientInfo) GetServiceType() string {
	return info.Type
}
func (info TestClientInfo) String() string {
	return fmt.Sprintf("TestClientInfo{ID:%s,Type:%s}", info.ID, info.Type)
}

var initS2CRegistry gogistryProto.RegistryInfo
var initS2SRegistry gogistryProto.RegistryInfo

func InitS2SRegistry() {
	ServerInfo := TestServerInfo{
		ID:   "UNDEF",
		Type: "UNDEF",
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

func ServerTest(t *testing.T, id int, Type string, wg *sync.WaitGroup) {
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
	server := NewServer(TestServerInfo{ID: fmt.Sprintf("SERVER_%d", id), Type: Type}, option)
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
		fmt.Println(s + fmt.Sprintf("%s is going to start.", server.GetServerInfo()))
		stoppedChan := make(chan bool, 1)
		go func() {
			server.Run()
			stoppedChan <- true
			close(stoppedChan)
		}()
		go func() {
			for i := 0; ; i++ {
				select {
				case <-time.After(1e9):
					C2SConnections := server.GetC2SConnections()
					S2SConnections := server.GetS2SConnections()
					t.Log(s + fmt.Sprintf("Check(%d)-S2C:%d\nS2S:%d,%s", i, len(C2SConnections), len(S2SConnections), S2SConnections))
				case <-stoppedChan:
					return
				}
			}
		}()
		select {
		case <-stoppedChan:
			t.Log(s + fmt.Sprintf("%s stopped itself.", server.GetServerInfo()))
		case <-time.After(40e9):
			server.Stop()
			t.Log(s + fmt.Sprintf("%s stopped manully.", server.GetServerInfo()))
		}
		wg.Done()
	}()
}

func ClientTest(t *testing.T, id int, Type string, wg *sync.WaitGroup) {
	s := fmt.Sprintf("ClientTest(%d)----", id)
	option := DefaultClientOption(initS2CRegistry, exampleProto.NewChanNetRequestProtocol())
	option.ResponseSendOption = exampleProto.ResponseSendOption{Timestamp: time.Now()}
	client := NewClient(TestClientInfo{ID: fmt.Sprintf("CLIENT_%d", id), Type: Type}, option)
	client.Events.NewConnection.AddHandler(func(info protocol.S2CInfo) {
		t.Log(s + fmt.Sprintf("%s---->NewConnection---->%s", client.GetClientInfo(), info.ServerInfo))
	})
	client.Events.NewConnection.Enable()
	client.Events.Disconnection.AddHandler(func(info protocol.S2CInfo, err error) {
		t.Log(s + fmt.Sprintf("%s---->Disconnection---->%s, error:%s", client.GetClientInfo(), info.ServerInfo, err))
	})
	client.Events.Disconnection.Enable()
	go func() {
		t.Log(s + fmt.Sprintf("%s is going to start.", client.GetClientInfo()))
		stoppedChan := make(chan bool, 1)
		go func() {
			client.Run()
			stoppedChan <- true
			close(stoppedChan)
		}()
		go func() {
			for i := 0; ; i++ {
				select {
				case <-time.After(1e9):
					t.Log(s + fmt.Sprintf("Check(%d)-C2S:%d", i, len(client.GetS2CConnections())))
				case <-stoppedChan:
					return
				}
			}
		}()
		select {
		case <-stoppedChan:
			fmt.Println(s + fmt.Sprintf("%s stopped itself.", client.GetClientInfo()))
		case <-time.After(16e9):
			client.Stop()
			fmt.Println(s + fmt.Sprintf("%s stopped manully.", client.GetClientInfo()))
		}
		wg.Done()
	}()
}

const SERVERN = 10
const CLIENTN = 60

func TestServerClient(t *testing.T) {
	InitS2SRegistry()
	Type := "TYPE_0"
	serverWG := new(sync.WaitGroup)
	serverWG.Add(SERVERN)
	for i := 0; i < SERVERN; i++ {
		ServerTest(t, i, Type, serverWG)
	}
	time.Sleep(2e9)
	clientWG := new(sync.WaitGroup)
	clientWG.Add(CLIENTN)
	for i := 0; i < CLIENTN; i++ {
		ClientTest(t, i, Type, clientWG)
		time.Sleep(0.5e9)
	}
	serverWG.Wait()
	clientWG.Wait()
}
