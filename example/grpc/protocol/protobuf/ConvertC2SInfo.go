package protocol

import (
	"errors"
	"fmt"
	"github.com/yindaheng98/gogisnet/protocol"
)

func (info *C2SInfo) Unpack() (*protocol.C2SInfo, error) {
	if info == nil {
		return nil, errors.New("C2SInfo is nil")
	}
	if info.ClientInfo == nil {
		return nil, errors.New("ClientInfo is nil")
	}
	return &protocol.C2SInfo{
		ClientInfo:         info.ClientInfo,
		ResponseSendOption: SendOption{Option: info.ResponseSendOption},
	}, nil
}

func C2SInfoPack(info protocol.C2SInfo) (i *C2SInfo, e error) {
	if info.ClientInfo == nil {
		return nil, errors.New("ClientInfo is nil")
	}
	defer func() {
		if r := recover(); r != nil {
			i, e = nil, errors.New(fmt.Sprint(r))
		}
	}()
	return &C2SInfo{
		ClientInfo:         info.ClientInfo.(*ClientInfo),
		ResponseSendOption: info.ResponseSendOption.(SendOption).Option,
	}, nil
}
