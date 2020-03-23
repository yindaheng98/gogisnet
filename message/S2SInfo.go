package message

import (
	"fmt"
	"github.com/yindaheng98/gogistry/protocol"
)

//Message that send from server to server.
//Implementation of "github.com/yindaheng98/gogistry/protocol".RegistryInfo
//Implementation of "github.com/yindaheng98/gogistry/protocol".RegistrantInfo
type S2SInfo struct {
	ServerInfo
	ResponseSendOption protocol.ResponseSendOption
	RequestSendOption  protocol.RequestSendOption
	GraphQueryAddr     string
	Candidates         []S2SInfo
	S2CInfo            S2CInfo
}

func (info S2SInfo) GetRegistrantID() string {
	return info.ServerInfo.GetServerID()
}
func (info S2SInfo) GetRegistryID() string {
	return info.ServerInfo.GetServerID()
}

func (info S2SInfo) GetRequestSendOption() protocol.RequestSendOption { //此服务端接收何种请求
	return info.RequestSendOption
}
func (info S2SInfo) GetResponseSendOption() protocol.ResponseSendOption { //此服务端接收何种请求
	return info.ResponseSendOption
}

func (info S2SInfo) GetCandidates() []protocol.RegistryInfo {
	r := make([]protocol.RegistryInfo, len(info.Candidates))
	for i, c := range info.Candidates {
		r[i] = c
	}
	return r
}
func (info S2SInfo) String() string {
	return fmt.Sprintf("S2SInfo{%s,RequestSendOption:%s,ResponseSendOption:%s,Candidates:%s}",
		info.ServerInfo.String(), info.RequestSendOption.String(), info.ResponseSendOption.String(), info.Candidates)
}
