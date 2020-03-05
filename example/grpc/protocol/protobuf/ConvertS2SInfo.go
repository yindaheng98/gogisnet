package protocol

import (
	"errors"
	"fmt"
	"github.com/yindaheng98/gogisnet/protocol"
)

func (info *S2SInfo) Unpack() (*protocol.S2SInfo, error) {
	if info == nil {
		return nil, errors.New("S2SInfo is nil")
	}
	var Candidates []protocol.S2SInfo
	if info.Candidates != nil {
		for _, c := range info.Candidates {
			Candidate, _ := c.Unpack()
			if Candidate != nil {
				Candidates = append(Candidates, *Candidate)
			}
		}
	}
	S2CInfo, err := info.S2CInfo.Unpack()
	if err != nil {
		return nil, err
	}
	return &protocol.S2SInfo{
		ServerInfo:         info.ServerInfo,
		ResponseSendOption: SendOption{Option: info.ResponseSendOption},
		RequestSendOption:  SendOption{Option: info.RequestSendOption},
		Candidates:         Candidates,
		S2CInfo:            *S2CInfo,
	}, nil
}

func S2SInfoPack(info protocol.S2SInfo) (i *S2SInfo, e error) {
	i, e = nil, nil
	var Candidates []*S2SInfo
	if info.Candidates != nil {
		for _, c := range info.Candidates {
			Candidate, _ := S2SInfoPack(c)
			if Candidate != nil {
				Candidates = append(Candidates, Candidate)
			}
		}
	}
	S2CInfo, err := S2CInfoPack(info.S2CInfo)
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			i, e = nil, errors.New(fmt.Sprint(r))
		}
	}()
	return &S2SInfo{
		ServerInfo:         info.ServerInfo.(*ServerInfo),
		ResponseSendOption: info.ResponseSendOption.(SendOption).Option,
		RequestSendOption:  info.RequestSendOption.(SendOption).Option,
		Candidates:         Candidates,
		S2CInfo:            S2CInfo,
	}, nil
}
