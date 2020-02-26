package emitters

import (
	"github.com/yindaheng98/go-utility/Emitter"
	"github.com/yindaheng98/gogisnet/protocol"
)

type S2SInfoEmitter struct {
	Emitter.Emitter
}

func NewSyncS2SInfoEmitter() *S2SInfoEmitter {
	return &S2SInfoEmitter{Emitter.NewSyncEmitter()}
}
func NewAsyncS2SInfoEmitter() *S2SInfoEmitter {
	return &S2SInfoEmitter{Emitter.NewAsyncEmitter()}
}

func (e *S2SInfoEmitter) AddHandler(handler func(info protocol.S2SInfo)) {
	e.Emitter.AddHandler(func(i interface{}) {
		handler(i.(protocol.S2SInfo))
	})
}

func (e *S2SInfoEmitter) Emit(info protocol.S2SInfo) {
	e.Emitter.Emit(info)
}
