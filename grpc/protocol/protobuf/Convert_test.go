package protocol

import (
	"fmt"
	"testing"
)

func NewClientInfo() *ClientInfo {
	return &ClientInfo{
		ClientID:       "Here's ID",
		ServiceType:    "Here's Type",
		AdditionalInfo: []byte("Here's AdditionalInfo"),
	}
}

func NewServerInfo(s string, i uint) *ServerInfo {
	return &ServerInfo{
		ServerID:       fmt.Sprintf("Here's %s ID %d", s, i),
		ServiceType:    fmt.Sprintf("Here's %s Type %d", s, i),
		AdditionalInfo: []byte(fmt.Sprintf("Here's %s AdditionalInfo %d", s, i)),
	}
}

func TestConvertC2SRequest(t *testing.T) {
	pbC2SReq := &C2SRequest{
		C2SInfo: &C2SInfo{
			ClientInfo:         NewClientInfo(),
			ResponseSendOption: nil,
		},
		Disconnect: true,
	}
	if C2SReq, err := pbC2SReq.Unpack(); err == nil {
		t.Log("Convert success: " + C2SReq.String())
		if pbC2SReq, err = C2SRequestPack(*C2SReq); err == nil {
			t.Log("Convert back success: ", pbC2SReq)
		} else {
			t.Log("Convert back failed: ", err.Error())
		}
	} else {
		t.Log("Convert failed: ", err.Error())
	}
}

func S2CInfoList(n uint) []*S2CInfo {
	Candidates := make([]*S2CInfo, n)
	for i := uint(1); i <= n; i++ {
		Candidates[i-1] = &S2CInfo{
			ServerInfo:        NewServerInfo("S2C", i),
			RequestSendOption: nil,
			Candidates:        nil,
		}
	}
	return Candidates
}

func TestConvertS2CResponse(t *testing.T) {
	pbS2CRes := &S2CResponse{
		S2CInfo: &S2CInfo{
			ServerInfo: NewServerInfo("", 0),
			RequestSendOption: &RequestSendOption{
				Option: nil,
				Addr:   "Here's S2C Addr",
			},
			Candidates: S2CInfoList(3),
		},
		Reject:  true,
		Timeout: 1e9,
	}
	if C2SReq, err := pbS2CRes.Unpack(); err == nil {
		t.Log("Convert success: " + C2SReq.String())
		if pbS2CRes, err = S2CResponsePack(*C2SReq); err == nil {
			t.Log("Convert back success: ", pbS2CRes)
		} else {
			t.Log("Convert back failed: ", err.Error())
		}
	} else {
		t.Log("Convert failed: ", err.Error())
	}
}

func S2SInfoList(n uint, m uint) []*S2SInfo {
	Candidates := make([]*S2SInfo, n)
	for i := uint(1); i <= n; i++ {
		Candidates[i-1] = &S2SInfo{
			ServerInfo:         NewServerInfo("", 0),
			ResponseSendOption: nil,
			RequestSendOption: &RequestSendOption{
				Option: nil,
				Addr:   fmt.Sprintf("Here's S2S Addr %d", i),
			},
			Candidates: nil,
			S2CInfo: &S2CInfo{
				ServerInfo: &ServerInfo{
					ServerID:       "Here's S2C ID",
					ServiceType:    "Here's S2C Type",
					AdditionalInfo: []byte("Here's S2C AdditionalInfo"),
				},
				RequestSendOption: &RequestSendOption{
					Option: nil,
					Addr:   fmt.Sprintf("Here's S2S Addr %d", i),
				},
				Candidates: S2CInfoList(m),
			},
		}
	}
	return Candidates
}

func TestConvertS2SRequest(t *testing.T) {
	pbS2SReq := &S2SRequest{
		S2SInfo: &S2SInfo{
			ServerInfo: NewServerInfo("", 0),
			RequestSendOption: &RequestSendOption{
				Option: nil,
				Addr:   "Here's S2S Addr",
			},
			ResponseSendOption: &ResponseSendOption{
				Option: nil,
			},
			S2CInfo: &S2CInfo{
				ServerInfo: &ServerInfo{
					ServerID:       "Here's S2C ID",
					ServiceType:    "Here's S2C Type",
					AdditionalInfo: []byte("Here's S2C AdditionalInfo"),
				},
				RequestSendOption: nil,
				Candidates:        nil,
			},
			Candidates: S2SInfoList(3, 3),
		},
		Disconnect: true,
	}
	if S2SReq, err := pbS2SReq.Unpack(); err == nil {
		t.Log("Convert success: " + S2SReq.String())
		if pbS2SReq, err = S2SRequestPack(*S2SReq); err == nil {
			t.Log("Convert back success: ", pbS2SReq)
		} else {
			t.Log("Convert back failed: ", err.Error())
		}
	} else {
		t.Log("Convert failed: ", err.Error())
	}
}

func TestConvertS2SResponse(t *testing.T) {
	pbS2SRes := &S2SResponse{
		S2SInfo: &S2SInfo{
			ServerInfo:         NewServerInfo("", 0),
			RequestSendOption:  nil,
			ResponseSendOption: nil,
			S2CInfo: &S2CInfo{
				ServerInfo:        NewServerInfo("S2C", 0),
				RequestSendOption: nil,
				Candidates:        nil,
			},
			Candidates: S2SInfoList(3, 3),
		},
		Reject:  true,
		Timeout: 1e9,
	}
	if S2SRes, err := pbS2SRes.Unpack(); err == nil {
		t.Log("Convert success: " + S2SRes.String())
		if pbS2SRes, err = S2SResponsePack(*S2SRes); err == nil {
			t.Log("Convert back success: ", pbS2SRes)
		} else {
			t.Log("Convert back failed: ", err.Error())
		}
	} else {
		t.Log("Convert failed: ", err.Error())
	}
}

func C2SInfoList(n uint) []*C2SInfo {
	Candidates := make([]*C2SInfo, n)
	for i := uint(1); i <= n; i++ {
		Candidates[i-1] = &C2SInfo{
			ClientInfo:         NewClientInfo(),
			ResponseSendOption: &ResponseSendOption{},
		}
	}
	return Candidates
}

func TestGraphQueryInfo(t *testing.T) {
	pbGraphQueryInfo := &GraphQueryInfo{
		S2SInfo:   S2SInfoList(1, 1)[0],
		Indegree:  S2SInfoList(3, 3),
		Outdegree: S2SInfoList(3, 3),
		Clients:   C2SInfoList(8),
	}
	if GraphQueryInfo, err := pbGraphQueryInfo.Unpack(); err == nil {
		t.Log(fmt.Sprintf("Convert success: %s", GraphQueryInfo))
		if pbGraphQueryInfo, err = GraphQueryInfoPack(*GraphQueryInfo); err == nil {
			t.Log("Convert back success: ", pbGraphQueryInfo)
		} else {
			t.Log("Convert back failed: ", err.Error())
		}
	} else {
		t.Log("Convert failed: ", err.Error())
	}

}
