package option

import pb "github.com/yindaheng98/gogisnet/grpc/protocol/protobuf"

type ClientInfoOption struct {
	ClientID       string `yaml:"ClientID" usage:"Unique ID of the client."`
	ServiceType    string `yaml:"ServiceType" usage:"Type of the client. Must be same as the type of the server you want to connect."`
	AdditionalInfo string `yaml:"AdditionalInfo" usage:"The additional information you want to attach to this client."`
}

func DefaultClientInfoOption() ClientInfoOption {
	return ClientInfoOption{
		ClientID:       "CLIENT-" + RandomString(64),
		ServiceType:    "undefined",
		AdditionalInfo: "",
	}
}
func (o ClientInfoOption) PutOption(op *pb.ClientInfo) {
	op.ClientID = o.ClientID
	op.ServiceType = o.ServiceType
	op.AdditionalInfo = []byte(o.AdditionalInfo)
}
