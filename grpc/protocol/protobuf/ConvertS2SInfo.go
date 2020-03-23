package protocol

import (
	"errors"
	"fmt"
	"github.com/yindaheng98/gogisnet/message"
)

func (info *S2SInfo) Unpack() (*message.S2SInfo, error) {
	if info == nil {
		return nil, errors.New("S2SInfo is nil")
	}
	var Candidates []message.S2SInfo
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
	return &message.S2SInfo{
		ServerInfo:         info.ServerInfo,
		ResponseSendOption: info.ResponseSendOption,
		RequestSendOption:  info.RequestSendOption,
		GraphQueryAddr:     info.GraphQueryAddr,
		Candidates:         Candidates,
		S2CInfo:            *S2CInfo,
	}, nil
}

func S2SInfoPack(info message.S2SInfo) (i *S2SInfo, e error) {
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
	var rso *ResponseSendOption
	if info.ResponseSendOption != nil {
		rso = info.ResponseSendOption.(*ResponseSendOption)
	}
	defer func() {
		if r := recover(); r != nil {
			i, e = nil, errors.New(fmt.Sprint(r))
		}
	}()
	return &S2SInfo{
		ServerInfo:         info.ServerInfo.(*ServerInfo),
		ResponseSendOption: rso,
		RequestSendOption:  info.RequestSendOption.(*RequestSendOption),
		GraphQueryAddr:     info.GraphQueryAddr,
		Candidates:         Candidates,
		S2CInfo:            S2CInfo,
	}, nil
}
