package emitters

import (
	"github.com/yindaheng98/go-utility/Emitter"
	"github.com/yindaheng98/gogisnet/protocol"
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

func (e *S2CInfoErrorEmitter) AddHandler(handler func(info protocol.S2CInfo, err error)) {
	e.IndefiniteEmitter.AddHandler(func(args ...interface{}) {
		handler(args[0].(protocol.S2CInfo), args[1].(error))
	})
}

func (e *S2CInfoErrorEmitter) Emit(info protocol.S2CInfo, err error) {
	e.IndefiniteEmitter.Emit(info, err)
}
