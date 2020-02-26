package emitters

import (
	"github.com/yindaheng98/go-utility/Emitter"
	"github.com/yindaheng98/gogisnet/protocol"
)

type C2SInfoEmitter struct {
	Emitter.Emitter
}

func NewSyncC2SInfoEmitter() *C2SInfoEmitter {
	return &C2SInfoEmitter{Emitter.NewSyncEmitter()}
}
func NewAsyncC2SInfoEmitter() *C2SInfoEmitter {
	return &C2SInfoEmitter{Emitter.NewAsyncEmitter()}
}

func (e *C2SInfoEmitter) AddHandler(handler func(info protocol.C2SInfo)) {
	e.Emitter.AddHandler(func(i interface{}) {
		handler(i.(protocol.C2SInfo))
	})
}

func (e *C2SInfoEmitter) Emit(info protocol.C2SInfo) {
	e.Emitter.Emit(info)
}
