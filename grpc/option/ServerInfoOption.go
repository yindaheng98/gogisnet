package option

import (
	pb "github.com/yindaheng98/gogisnet/grpc/protocol/protobuf"
)

type ServerInfoOption struct {
	ServerID       string            `yaml:"ServerID" usage:"Unique ID of the server. Set to 'undefined' to generate a unique ID automatically."`
	ServiceType    string            `yaml:"ServiceType" usage:"Type of the server. Must be same as the type of the server you want to connect."`
	AdditionalInfo map[string][]byte `yaml:"AdditionalInfo" usage:"The additional information you want to attach to this server."`
}

func DefaultServerInfoOption() ServerInfoOption {
	return ServerInfoOption{
		ServerID:       "undefined",
		ServiceType:    "undefined",
		AdditionalInfo: map[string][]byte{},
	}
}
func (o ServerInfoOption) PutOption(op *pb.ServerInfo) {
	op.ServerID = o.ServerID
	op.ServiceType = o.ServiceType
	op.AdditionalInfo = o.AdditionalInfo
}
