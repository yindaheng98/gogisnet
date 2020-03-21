package option

import (
	pb "github.com/yindaheng98/gogisnet/grpc/protocol/protobuf"
)

type ServerInfoOption struct {
	ServerID                string `yaml:"ServerID" usage:"Unique ID of the server."`
	ServiceType             string `yaml:"ServiceType" usage:"Type of the server. Must be same as the type of the server you want to connect."`
	GraphQueryBroadCastAddr string `yaml:"GraphQueryBroadCastAddr" usage:"Broad cast address of GraphQuery service of the server."`
	AdditionalInfo          string `yaml:"AdditionalInfo" usage:"The additional information you want to attach to this server."`
}

func DefaultServerInfoOption() ServerInfoOption {
	return ServerInfoOption{
		ServerID:                "SERVER-" + RandomString(64),
		ServiceType:             "undefined",
		GraphQueryBroadCastAddr: GetIP() + ":4242",
		AdditionalInfo:          "",
	}
}

func (o ServerInfoOption) PutOption(op *pb.ServerInfo) {
	op.ServerID = o.ServerID
	op.ServiceType = o.ServiceType
	op.GraphQueryBroadCastAddr = o.GraphQueryBroadCastAddr
	op.AdditionalInfo = []byte(o.AdditionalInfo)
}
