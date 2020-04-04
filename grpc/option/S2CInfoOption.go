package option

import pb "github.com/yindaheng98/gogisnet/grpc/protocol/protobuf"

type S2CInfoOption struct {
	ServerInfoOption ServerInfoOption `yaml:"ServerInfoOption" usage:"Information about this server."`
	BoardCastAddr    string           `yaml:"BoardCastAddr" usage:"Broad cast address of the service."`
}

func DefaultS2CInfoOption() S2CInfoOption {
	return S2CInfoOption{
		ServerInfoOption: ServerInfoOption{
			ServerID:       "undefined",
			ServiceType:    "undefined",
			AdditionalInfo: map[string][]byte{},
		},
		BoardCastAddr: "undefined",
	}
}

func (o S2CInfoOption) PutOption(op *pb.S2CInfo) {
	if op.ServerInfo == nil {
		op.ServerInfo = &pb.ServerInfo{}
	}
	o.ServerInfoOption.PutOption(op.ServerInfo)

	if op.RequestSendOption == nil {
		op.RequestSendOption = &pb.RequestSendOption{}
	}
	op.RequestSendOption = &pb.RequestSendOption{Addr: o.BoardCastAddr}
}
