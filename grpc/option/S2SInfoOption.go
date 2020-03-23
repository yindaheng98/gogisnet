package option

import pb "github.com/yindaheng98/gogisnet/grpc/protocol/protobuf"

type S2SInfoOption struct {
	ServerInfoOption ServerInfoOption `yaml:"ServerInfoOption" usage:"Information about this server."`
	BoardCastAddr    string           `yaml:"BoardCastAddr" usage:"Broad cast address of the service."`
	GraphQueryAddr   string           `yaml:"GraphQueryAddr" usage:"Broad cast address of GraphQuery service of the server."`
	S2CInfoOption    S2CInfoOption    `yaml:"S2CInfoOption" usage:"S2CInfo of this server."`
}

func DefaultS2SInfoOption() S2SInfoOption {
	return S2SInfoOption{
		ServerInfoOption: DefaultServerInfoOption(),
		BoardCastAddr:    "undefined",
		GraphQueryAddr:   GetIP() + ":4242",
		S2CInfoOption:    DefaultS2CInfoOption(),
	}
}

func (o S2SInfoOption) PutOption(op *pb.S2SInfo) {
	if op.ServerInfo == nil {
		op.ServerInfo = &pb.ServerInfo{}
	}
	o.ServerInfoOption.PutOption(op.ServerInfo)

	if op.S2CInfo == nil {
		op.S2CInfo = &pb.S2CInfo{}
	}
	o.S2CInfoOption.PutOption(op.S2CInfo)

	if op.RequestSendOption == nil {
		op.RequestSendOption = &pb.RequestSendOption{}
	}
	op.RequestSendOption.Addr = o.BoardCastAddr
	op.GraphQueryAddr = o.GraphQueryAddr
}
