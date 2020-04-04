package option

import pb "github.com/yindaheng98/gogisnet/grpc/protocol/protobuf"

type ClientInfoOption struct {
	ClientID       string            `yaml:"ClientID" usage:"Unique ID of the client. Set to 'undefined' to generate a unique ID automatically."`
	ServiceType    string            `yaml:"ServiceType" usage:"Type of the client. Must be same as the type of the server you want to connect."`
	AdditionalInfo map[string][]byte `yaml:"AdditionalInfo" usage:"The additional information you want to attach to this client."`
}

func DefaultClientInfoOption() ClientInfoOption {
	return ClientInfoOption{
		ClientID:       "undefined",
		ServiceType:    "undefined",
		AdditionalInfo: map[string][]byte{},
	}
}
func (o ClientInfoOption) PutOption(op *pb.ClientInfo) {
	op.ClientID = o.ClientID
	op.ServiceType = o.ServiceType
	op.AdditionalInfo = o.AdditionalInfo
}
