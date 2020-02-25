package protocol

import (
	"fmt"
	"github.com/yindaheng98/gogistry/protocol"
)

type ClientInfo interface {
	GetClientID() string
	GetServiceType() string
	String() string
}

type ClientRegistrantInfo struct {
	ClientInfo
	ResponseSendOption protocol.ResponseSendOption
}

func (info ClientRegistrantInfo) GetResponseSendOption() protocol.ResponseSendOption { //此服务端接收何种请求
	return info.ResponseSendOption
}
func (info ClientRegistrantInfo) GetRegistrantID() string {
	return info.ClientInfo.GetClientID()
}
func (info ClientRegistrantInfo) String() string {
	return fmt.Sprintf("ServerRegistrantInfo{%s}", info.ClientInfo.String())
}
