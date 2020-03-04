package protocol

import (
	"errors"
	"fmt"
	"github.com/yindaheng98/gogisnet/protocol"
	gogistryProto "github.com/yindaheng98/gogistry/protocol"
)

func (request *C2SRequest) Unpack() (*gogistryProto.Request, error) {
	if request == nil {
		return nil, errors.New("C2SRequest is nil")
	}
	C2SInfo, err := request.C2SInfo.Unpack()
	if err != nil {
		return nil, err
	}
	return &gogistryProto.Request{
		RegistrantInfo: *C2SInfo,
		Disconnect:     request.Disconnect,
	}, nil
}

func C2SRequestPack(request gogistryProto.Request) (i *C2SRequest, e error) {
	i, e = nil, nil
	defer func() {
		if r := recover(); r != nil {
			i, e = nil, errors.New(fmt.Sprint(r))
		}
	}()
	C2SInfo, err := C2SInfoPack(request.RegistrantInfo.(protocol.C2SInfo))
	if err != nil {
		return nil, err
	}
	return &C2SRequest{
		C2SInfo:    C2SInfo,
		Disconnect: request.Disconnect,
	}, nil
}
