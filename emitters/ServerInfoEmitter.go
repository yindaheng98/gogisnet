package emitters

import (
	"github.com/yindaheng98/go-utility/Emitter"
	"github.com/yindaheng98/gogisnet/message"
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

func (e *ServerInfoEmitter) AddHandler(handler func(info message.ServerInfo)) {
	e.Emitter.AddHandler(func(i interface{}) {
		handler(i.(message.ServerInfo))
	})
}

func (e *ServerInfoEmitter) Emit(info message.ServerInfo) {
	e.Emitter.Emit(info)
}
