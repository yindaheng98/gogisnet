package emitters

import (
	"github.com/yindaheng98/go-utility/Emitter"
	"github.com/yindaheng98/gogisnet/protocol"
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

func (e *S2CInfoEmitter) AddHandler(handler func(info protocol.S2CInfo)) {
	e.Emitter.AddHandler(func(i interface{}) {
		handler(i.(protocol.S2CInfo))
	})
}

func (e *S2CInfoEmitter) Emit(info protocol.S2CInfo) {
	e.Emitter.Emit(info)
}
