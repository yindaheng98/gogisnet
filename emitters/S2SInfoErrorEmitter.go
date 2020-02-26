package emitters

import (
	"github.com/yindaheng98/go-utility/Emitter"
	"github.com/yindaheng98/gogisnet/protocol"
)

type S2SInfoErrorEmitter struct {
	Emitter.IndefiniteEmitter
}

func NewSyncS2SInfoErrorEmitter() *S2SInfoErrorEmitter {
	return &S2SInfoErrorEmitter{Emitter.NewSyncIndefiniteEmitter()}
}
func NewAsyncS2SInfoErrorEmitter() *S2SInfoErrorEmitter {
	return &S2SInfoErrorEmitter{Emitter.NewAsyncIndefiniteEmitter()}
}

func (e *S2SInfoErrorEmitter) AddHandler(handler func(info protocol.S2SInfo, err error)) {
	e.IndefiniteEmitter.AddHandler(func(args ...interface{}) {
		if args[1] == nil {
			handler(args[0].(protocol.S2SInfo), nil)
		} else {
			handler(args[0].(protocol.S2SInfo), args[1].(error))
		}
	})
}

func (e *S2SInfoErrorEmitter) Emit(info protocol.S2SInfo, err error) {
	e.IndefiniteEmitter.Emit(info, err)
}
