package emitters

import (
	"github.com/yindaheng98/go-utility/Emitter"
	"github.com/yindaheng98/gogisnet/protocol"
)

type ClientInfoEmitter struct {
	Emitter.Emitter
}

func NewSyncClientInfoEmitter() *ClientInfoEmitter {
	return &ClientInfoEmitter{Emitter.NewSyncEmitter()}
}
func NewAsyncClientInfoEmitter() *ClientInfoEmitter {
	return &ClientInfoEmitter{Emitter.NewAsyncEmitter()}
}

func (e *ClientInfoEmitter) AddHandler(handler func(info protocol.ClientInfo)) {
	e.Emitter.AddHandler(func(i interface{}) {
		handler(i.(protocol.ClientInfo))
	})
}

func (e *ClientInfoEmitter) Emit(info protocol.ClientInfo) {
	e.Emitter.Emit(info)
}
