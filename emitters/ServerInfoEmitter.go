package emitters

import (
	"github.com/yindaheng98/go-utility/Emitter"
	"github.com/yindaheng98/gogisnet/protocol"
)

type ServerInfoEmitter struct {
	Emitter.Emitter
}

func NewSyncServerInfoEmitter() *ServerInfoEmitter {
	return &ServerInfoEmitter{Emitter.NewSyncEmitter()}
}
func NewAsyncServerInfoEmitter() *ServerInfoEmitter {
	return &ServerInfoEmitter{Emitter.NewAsyncEmitter()}
}

func (e *ServerInfoEmitter) AddHandler(handler func(info protocol.ServerInfo)) {
	e.Emitter.AddHandler(func(i interface{}) {
		handler(i.(protocol.ServerInfo))
	})
}

func (e *ServerInfoEmitter) Emit(info protocol.ServerInfo) {
	e.Emitter.Emit(info)
}
