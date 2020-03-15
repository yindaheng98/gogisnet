package emitters

import (
	"github.com/yindaheng98/go-utility/Emitter"
	"github.com/yindaheng98/gogisnet/message"
)

type S2CInfoErrorEmitter struct {
	Emitter.IndefiniteEmitter
}

func NewSyncS2CInfoErrorEmitter() *S2CInfoErrorEmitter {
	return &S2CInfoErrorEmitter{Emitter.NewSyncIndefiniteEmitter()}
}
func NewAsyncS2CInfoErrorEmitter() *S2CInfoErrorEmitter {
	return &S2CInfoErrorEmitter{Emitter.NewAsyncIndefiniteEmitter()}
}

func (e *S2CInfoErrorEmitter) AddHandler(handler func(info message.S2CInfo, err error)) {
	e.IndefiniteEmitter.AddHandler(func(args ...interface{}) {
		if args[1] == nil {
			handler(args[0].(message.S2CInfo), nil)
		} else {
			handler(args[0].(message.S2CInfo), args[1].(error))
		}
	})
}

func (e *S2CInfoErrorEmitter) Emit(info message.S2CInfo, err error) {
	e.IndefiniteEmitter.Emit(info, err)
}
