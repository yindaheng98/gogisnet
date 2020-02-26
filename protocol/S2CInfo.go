package protocol

import (
	"fmt"
	"github.com/yindaheng98/gogistry/protocol"
)

//服务器面向客户端的信息
type S2CInfo struct {
	ServerInfo
	RequestSendOption protocol.RequestSendOption
	Candidates        []S2CInfo
}

func (info S2CInfo) GetRegistryID() string {
	return info.ServerInfo.GetServerID()
}
func (info S2CInfo) GetRequestSendOption() protocol.RequestSendOption { //此服务端接收何种请求
	return info.RequestSendOption
}
func (info S2CInfo) GetCandidates() []protocol.RegistryInfo {
	r := make([]protocol.RegistryInfo, len(info.Candidates))
	for i, c := range info.Candidates {
		r[i] = c
	}
	return r
}
func (info S2CInfo) String() string {
	return fmt.Sprintf("S2CInfo{%s,RequestSendOption:%s,Candidates:%s}",
		info.ServerInfo, info.RequestSendOption, info.Candidates)
}
