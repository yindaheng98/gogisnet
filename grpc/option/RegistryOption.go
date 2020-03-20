package option

import (
	pb "github.com/yindaheng98/gogisnet/grpc/protocol/protobuf"
	"github.com/yindaheng98/gogisnet/server"
	"github.com/yindaheng98/gogistry/example/TimeoutController"
	"github.com/yindaheng98/gogistry/registry"
)

//RegistryOption is the options for a gogistry registry
type RegistryOption struct { //服务端面向服务端的收发设置
	BoardCastAddr     string                     `yaml:"BoardCastAddr" usage:"The IP Address that will send with messages to other server, so other server can easily find the current server."`
	MaxRegistrants    uint64                     `yaml:"MaxRegistrants" usage:"MaxRegistrants defined how much registrants this registry can connect at most."`
	TimeoutController registry.TimeoutController `yaml:"TimeoutController" usage:"TimeoutController is the TimeoutController used in gogistry registry."`
}

//DefaultRegistrantOption returns a default RegistrantOption
func DefaultRegistryOption() RegistryOption {
	return RegistryOption{
		BoardCastAddr:     "localhost:4242",
		MaxRegistrants:    4,
		TimeoutController: TimeoutController.DefaultLogTimeoutController(),
	}
}

//PutOption can convert a RegistryOption into a server.RegistryOption
func (o RegistryOption) PutOption(op *server.RegistryOption) {
	op.RequestSendOption = &pb.RequestSendOption{Addr: o.BoardCastAddr}
	op.TimeoutController = o.TimeoutController
	op.MaxRegistrants = o.MaxRegistrants
}
