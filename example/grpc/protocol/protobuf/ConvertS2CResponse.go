package protocol

import (
	"errors"
	"fmt"
	"github.com/yindaheng98/gogisnet/protocol"
	gogistryProto "github.com/yindaheng98/gogistry/protocol"
	"time"
)

func (response *S2CResponse) Unpack() (*gogistryProto.Response, error) {
	if response == nil {
		return nil, errors.New("S2SResponse is nil")
	}
	S2CInfo, err := response.S2CInfo.Unpack()
	if err != nil {
		return nil, err
	}
	return &gogistryProto.Response{
		RegistryInfo: *S2CInfo,
		Timeout:      time.Duration(response.Timeout),
		Reject:       response.Reject,
	}, nil
}
func S2CResponsePack(response gogistryProto.Response) (i *S2CResponse, e error) {
	i, e = nil, nil
	defer func() {
		if r := recover(); r != nil {
			i, e = nil, errors.New(fmt.Sprint(r))
		}
	}()
	S2CInfo, err := S2CInfoPack(response.RegistryInfo.(protocol.S2CInfo))
	if err != nil {
		return nil, err
	}
	return &S2CResponse{
		S2CInfo: S2CInfo,
		Timeout: int64(response.Timeout),
		Reject:  response.Reject,
	}, nil
}
