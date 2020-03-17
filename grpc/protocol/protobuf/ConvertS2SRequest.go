package protocol

import (
	"errors"
	"fmt"
	"github.com/yindaheng98/gogisnet/message"
	gogistryProto "github.com/yindaheng98/gogistry/protocol"
)

func (request *S2SRequest) Unpack() (*gogistryProto.Request, error) {
	if request == nil {
		return nil, errors.New("S2SRequest is nil")
	}
	S2SInfo, err := request.S2SInfo.Unpack()
	if err != nil {
		return nil, err
	}
	return &gogistryProto.Request{
		RegistrantInfo: *S2SInfo,
		Disconnect:     request.Disconnect,
	}, nil
}

func S2SRequestPack(request gogistryProto.Request) (i *S2SRequest, e error) {
	i, e = nil, nil
	defer func() {
		if r := recover(); r != nil {
			i, e = nil, errors.New(fmt.Sprint(r))
		}
	}()
	S2SInfo, err := S2SInfoPack(request.RegistrantInfo.(message.S2SInfo))
	if err != nil {
		return nil, err
	}
	return &S2SRequest{
		S2SInfo:    S2SInfo,
		Disconnect: request.Disconnect,
	}, nil
}
