package protocol

import (
	"errors"
	"fmt"
	"github.com/yindaheng98/gogisnet/message"
)

func (info *C2SInfo) Unpack() (*message.C2SInfo, error) {
	if info == nil {
		return nil, errors.New("C2SInfo is nil")
	}
	if info.ClientInfo == nil {
		return nil, errors.New("ClientInfo is nil")
	}
	return &message.C2SInfo{
		ClientInfo:         info.ClientInfo,
		ResponseSendOption: info.ResponseSendOption,
	}, nil
}

func C2SInfoPack(info message.C2SInfo) (i *C2SInfo, e error) {
	if info.ClientInfo == nil {
		return nil, errors.New("ClientInfo is nil")
	}
	var rso *ResponseSendOption
	if info.ResponseSendOption != nil {
		rso = info.ResponseSendOption.(*ResponseSendOption)
	}
	defer func() {
		if r := recover(); r != nil {
			i, e = nil, errors.New(fmt.Sprint(r))
		}
	}()
	return &C2SInfo{
		ClientInfo:         info.ClientInfo.(*ClientInfo),
		ResponseSendOption: rso,
	}, nil
}
