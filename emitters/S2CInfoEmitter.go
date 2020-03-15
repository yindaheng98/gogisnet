package emitters

import (
	"github.com/yindaheng98/go-utility/Emitter"
	"github.com/yindaheng98/gogisnet/message"
)

type S2CInfoEmitter struct {
	Emitter.Emitter
}

func NewSyncS2CInfoEmitter() *S2CInfoEmitter {
	return &S2CInfoEmitter{Emitter.NewSyncEmitter()}
}
func NewAsyncS2CInfoEmitter() *S2CInfoEmitter {
	return &S2CInfoEmitter{Emitter.NewAsyncEmitter()}
}

func (e *S2CInfoEmitter) AddHandler(handler func(info message.S2CInfo)) {
	e.Emitter.AddHandler(func(i interface{}) {
		handler(i.(message.S2CInfo))
	})
}

func (e *S2CInfoEmitter) Emit(info message.S2CInfo) {
	e.Emitter.Emit(info)
}
