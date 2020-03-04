package protocol

import (
	"fmt"
	"testing"
)

func TestConvertC2SRequest(t *testing.T) {
	pbC2SReq := &C2SRequest{
		C2SInfo: &C2SInfo{
			ClientInfo: &ClientInfo{
				ID:             "Here's ID",
				Type:           "Here's Type",
				AdditionalInfo: "Here's AdditionalInfo",
			},
			ResponseSendOption: "Here's ResponseSendOption",
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
			ServerInfo: &ServerInfo{
				ID:             fmt.Sprintf("Here's S2C ID %d", i),
				Type:           fmt.Sprintf("Here's S2C Type %d", i),
				Addr:           fmt.Sprintf("Here's S2C Addr %d", i),
				AdditionalInfo: fmt.Sprintf("Here's S2C AdditionalInfo %d", i),
			},
			RequestSendOption: fmt.Sprintf("Here's S2C RequestSendOption %d", i),
			Candidates:        nil,
		}
	}
	return Candidates
}

func TestConvertS2CResponse(t *testing.T) {
	pbS2CRes := &S2CResponse{
		S2CInfo: &S2CInfo{
			ServerInfo: &ServerInfo{
				ID:             "Here's ID",
				Type:           "Here's Type",
				Addr:           "Here's Addr",
				AdditionalInfo: "Here's AdditionalInfo",
			},
			RequestSendOption: "Here's RequestSendOption",
			Candidates:        S2CInfoList(3),
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
			ServerInfo: &ServerInfo{
				ID:             fmt.Sprintf("Here's S2S ID %d", i),
				Type:           fmt.Sprintf("Here's S2S Type %d", i),
				Addr:           fmt.Sprintf("Here's S2S Addr %d", i),
				AdditionalInfo: fmt.Sprintf("Here's S2S AdditionalInfo %d", i),
			},
			ResponseSendOption: fmt.Sprintf("Here's S2S ResponseSendOption %d", i),
			RequestSendOption:  fmt.Sprintf("Here's S2S RequestSendOption %d", i),
			Candidates:         nil,
			S2CInfo: &S2CInfo{
				ServerInfo: &ServerInfo{
					ID:             "Here's S2C ID",
					Type:           "Here's S2C Type",
					Addr:           "Here's S2C addr",
					AdditionalInfo: "Here's S2C AdditionalInfo",
				},
				RequestSendOption: "Here's S2C RequestSendOption",
				Candidates:        S2CInfoList(m),
			},
		}
	}
	return Candidates
}

func TestConvertS2SRequest(t *testing.T) {
	pbS2SReq := &S2SRequest{
		S2SInfo: &S2SInfo{
			ServerInfo: &ServerInfo{
				ID:             "Here's ID",
				Type:           "Here's Type",
				Addr:           "Here's addr",
				AdditionalInfo: "Here's AdditionalInfo",
			},
			RequestSendOption:  "Here's S2C RequestSendOption",
			ResponseSendOption: "Here's S2C ResponseSendOption",
			S2CInfo: &S2CInfo{
				ServerInfo: &ServerInfo{
					ID:             "Here's S2C ID",
					Type:           "Here's S2C Type",
					Addr:           "Here's S2C addr",
					AdditionalInfo: "Here's S2C AdditionalInfo",
				},
				RequestSendOption: "Here's S2C RequestSendOption",
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
			ServerInfo: &ServerInfo{
				ID:             "Here's ID",
				Type:           "Here's Type",
				Addr:           "Here's addr",
				AdditionalInfo: "Here's AdditionalInfo",
			},
			RequestSendOption:  "Here's S2C RequestSendOption",
			ResponseSendOption: "Here's S2C ResponseSendOption",
			S2CInfo: &S2CInfo{
				ServerInfo: &ServerInfo{
					ID:             "Here's S2C ID",
					Type:           "Here's S2C Type",
					Addr:           "Here's S2C addr",
					AdditionalInfo: "Here's S2C AdditionalInfo",
				},
				RequestSendOption: "Here's S2C RequestSendOption",
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
