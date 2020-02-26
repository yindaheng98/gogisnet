package protocol

import (
	"fmt"
	"github.com/yindaheng98/gogistry/protocol"
)

//客户端面向服务器的信息
type C2SInfo struct {
	ClientInfo
	ResponseSendOption protocol.ResponseSendOption
}

func (info C2SInfo) GetResponseSendOption() protocol.ResponseSendOption { //此服务端接收何种请求
	return info.ResponseSendOption
}
func (info C2SInfo) GetRegistrantID() string {
	return info.ClientInfo.GetClientID()
}
func (info C2SInfo) String() string {
	return fmt.Sprintf("C2SInfo{%s}", info.ClientInfo.String())
}
