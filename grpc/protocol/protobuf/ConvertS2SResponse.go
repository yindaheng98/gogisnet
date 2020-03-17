package protocol

import (
	"errors"
	"fmt"
	"github.com/yindaheng98/gogisnet/message"
	gogistryProto "github.com/yindaheng98/gogistry/protocol"
	"time"
)

func (response *S2SResponse) Unpack() (*gogistryProto.Response, error) {
	if response == nil {
		return nil, errors.New("S2SResponse is nil")
	}
	S2SInfo, err := response.S2SInfo.Unpack()
	if err != nil {
		return nil, err
	}
	return &gogistryProto.Response{
		RegistryInfo: *S2SInfo,
		Timeout:      time.Duration(response.Timeout),
		Reject:       response.Reject,
	}, nil
}

func S2SResponsePack(response gogistryProto.Response) (i *S2SResponse, e error) {
	i, e = nil, nil
	defer func() {
		if r := recover(); r != nil {
			i, e = nil, errors.New(fmt.Sprint(r))
		}
	}()
	S2SInfo, err := S2SInfoPack(response.RegistryInfo.(message.S2SInfo))
	if err != nil {
		return nil, err
	}
	return &S2SResponse{
		S2SInfo: S2SInfo,
		Timeout: int64(response.Timeout),
		Reject:  response.Reject,
	}, nil
}
