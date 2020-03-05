package protocol

import (
	"errors"
	"fmt"
	"github.com/yindaheng98/gogisnet/protocol"
)

func (info *S2CInfo) Unpack() (*protocol.S2CInfo, error) {
	if info == nil {
		return nil, errors.New("S2CInfo is nil")
	}
	var Candidates []protocol.S2CInfo
	if info.Candidates != nil {
		for _, c := range info.Candidates {
			Candidate, _ := c.Unpack()
			if Candidate != nil {
				Candidates = append(Candidates, *Candidate)
			}
		}
	}
	return &protocol.S2CInfo{
		ServerInfo:        info.ServerInfo,
		RequestSendOption: info.RequestSendOption,
		Candidates:        Candidates,
	}, nil
}

func S2CInfoPack(info protocol.S2CInfo) (i *S2CInfo, e error) {
	i, e = nil, nil
	var Candidates []*S2CInfo
	if info.Candidates != nil {
		for _, c := range info.Candidates {
			Candidate, _ := S2CInfoPack(c)
			if Candidate != nil {
				Candidates = append(Candidates, Candidate)
			}
		}
	}
	defer func() {
		if r := recover(); r != nil {
			i, e = nil, errors.New(fmt.Sprint(r))
		}
	}()
	return &S2CInfo{
		ServerInfo:        info.ServerInfo.(*ServerInfo),
		RequestSendOption: info.RequestSendOption.(*RequestSendOption),
		Candidates:        Candidates,
	}, nil
}
