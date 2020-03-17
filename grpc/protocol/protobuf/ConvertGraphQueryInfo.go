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
	Indegree := make([]message.ServerInfo, len(info.Indegree))
	for i, v := range info.Indegree {
		Indegree[i] = v
	}
	Outdegree := make([]message.ServerInfo, len(info.Outdegree))
	for i, v := range info.Outdegree {
		Outdegree[i] = v
	}
	Clients := make([]message.ClientInfo, len(info.Clients))
	for i, v := range info.Clients {
		Clients[i] = v
	}
	return &message.GraphQueryInfo{
		ServerInfo: info.ServerInfo,
		Indegree:   Indegree,
		Outdegree:  Outdegree,
		Clients:    Clients,
	}, nil
}

//Convert a message.GraphQueryInfo into a protobuf GraphQueryInfo
func GraphQueryInfoPack(info message.GraphQueryInfo) (*GraphQueryInfo, error) {
	if info.ServerInfo == nil {
		return nil, errors.New("ServerInfo is nil")
	}
	Indegree := make([]*ServerInfo, len(info.Indegree))
	for i, v := range info.Indegree {
		if v != nil {
			Indegree[i] = v.(*ServerInfo)
		}
	}
	Outdegree := make([]*ServerInfo, len(info.Outdegree))
	for i, v := range info.Outdegree {
		if v != nil {
			Outdegree[i] = v.(*ServerInfo)
		}
	}
	Clients := make([]*ClientInfo, len(info.Clients))
	for i, v := range info.Clients {
		if v != nil {
			Clients[i] = v.(*ClientInfo)
		}
	}
	return &GraphQueryInfo{
		ServerInfo: info.ServerInfo.(*ServerInfo),
		Indegree:   Indegree,
		Outdegree:  Outdegree,
		Clients:    Clients,
	}, nil
}
