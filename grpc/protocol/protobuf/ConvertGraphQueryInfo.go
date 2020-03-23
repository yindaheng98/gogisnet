package protocol

import (
	"errors"
	"github.com/yindaheng98/gogisnet/message"
)

//Convert a protobuf GraphQueryInfo into a message.GraphQueryInfo
func (info *GraphQueryInfo) Unpack() (*message.GraphQueryInfo, error) {
	if info == nil {
		return nil, errors.New("GraphQueryInfo is nil")
	}
	S2SInfo, err := info.S2SInfo.Unpack()
	if err != nil {
		return nil, err
	}
	var Indegree []message.S2SInfo
	for _, v := range info.Indegree {
		vp, _ := v.Unpack()
		if vp != nil {
			Indegree = append(Indegree, *vp)
		}
	}
	var Outdegree []message.S2SInfo
	for _, v := range info.Outdegree {
		vp, _ := v.Unpack()
		if vp != nil {
			Outdegree = append(Outdegree, *vp)
		}
	}
	var Clients []message.C2SInfo
	for _, v := range info.Clients {
		vp, _ := v.Unpack()
		if vp != nil {
			Clients = append(Clients, *vp)
		}
	}
	return &message.GraphQueryInfo{
		S2SInfo:   *S2SInfo,
		Indegree:  Indegree,
		Outdegree: Outdegree,
		Clients:   Clients,
	}, nil
}

//Convert a message.GraphQueryInfo into a protobuf GraphQueryInfo
func GraphQueryInfoPack(info message.GraphQueryInfo) (*GraphQueryInfo, error) {
	s2sInfo, err := S2SInfoPack(info.S2SInfo)
	if err != nil {
		return nil, err
	}
	var Indegree []*S2SInfo
	for _, v := range info.Indegree {
		vp, _ := S2SInfoPack(v)
		if vp != nil {
			Indegree = append(Indegree, vp)
		}
	}
	var Outdegree []*S2SInfo
	for _, v := range info.Outdegree {
		vp, _ := S2SInfoPack(v)
		if vp != nil {
			Outdegree = append(Outdegree, vp)
		}
	}
	var Clients []*C2SInfo
	for _, v := range info.Clients {
		vp, _ := C2SInfoPack(v)
		if vp != nil {
			Clients = append(Clients, vp)
		}
	}
	return &GraphQueryInfo{
		S2SInfo:   s2sInfo,
		Indegree:  Indegree,
		Outdegree: Outdegree,
		Clients:   Clients,
	}, nil
}
