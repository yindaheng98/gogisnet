package option

import (
	pb "github.com/yindaheng98/gogisnet/grpc/protocol/protobuf"
	"github.com/yindaheng98/gogisnet/server"
	"github.com/yindaheng98/gogistry/example/TimeoutController"
	"time"
)

//RegistryOption is the options for a gogistry registry
type RegistryOption struct { //服务端面向服务端的收发设置
	BoardCastAddr           string                  `yaml:"BoardCastAddr" usage:"The IP Address that will send with messages to other server, so other server can easily find the current server."`
	MaxRegistrants          uint64                  `yaml:"MaxRegistrants" usage:"MaxRegistrants defined how much registrants this registry can connect at most."`
	TimeoutControllerOption TimeoutControllerOption `yaml:"TimeoutControllerOption" usage:"TimeoutControllerOption is the option for TimeoutController used in gogistry registry."`
}

//DefaultRegistrantOption returns a default RegistrantOption
func DefaultRegistryOption() RegistryOption {
	return RegistryOption{
		BoardCastAddr:           "localhost:4242",
		MaxRegistrants:          4,
		TimeoutControllerOption: defaultTimeoutControllerOption(),
	}
}

//PutOption can convert a RegistryOption into a server.RegistryOption
func (o RegistryOption) PutOption(op *server.RegistryOption) {
	op.RequestSendOption = &pb.RequestSendOption{Addr: o.BoardCastAddr}
	o.TimeoutControllerOption.PutOption(op.TimeoutController.(*TimeoutController.LogTimeoutController))
	op.MaxRegistrants = o.MaxRegistrants
}

//The option for TimeoutController used in gogistry registry
type TimeoutControllerOption struct {
	MinimumTime    time.Duration //最小Timeout
	MaximumTime    time.Duration //最大Timeout
	IncreaseFactor float64       //从最小到最大的增长系数
}

func (o TimeoutControllerOption) PutOption(op *TimeoutController.LogTimeoutController) {
	op.IncreaseFactor = o.IncreaseFactor
	op.MaximumTime = o.MaximumTime
	op.MinimumTime = o.MinimumTime
}

func defaultTimeoutControllerOption() TimeoutControllerOption {
	op := TimeoutController.DefaultLogTimeoutController()
	return TimeoutControllerOption{
		MinimumTime:    op.MinimumTime,
		MaximumTime:    op.MaximumTime,
		IncreaseFactor: op.IncreaseFactor,
	}
}
