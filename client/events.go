package client

import (
	"github.com/yindaheng98/go-utility/Emitter"
	"github.com/yindaheng98/gogisnet/emitters"
	"github.com/yindaheng98/gogisnet/message"
	gogistryEmitters "github.com/yindaheng98/gogistry/emitters"
	gogistryProto "github.com/yindaheng98/gogistry/protocol"
)

type events struct {
	NewConnection    *emitters.S2CInfoEmitter
	UpdateConnection *emitters.S2CInfoEmitter
	Disconnection    *emitters.S2CInfoErrorEmitter
	Error            *Emitter.ErrorEmitter
	Retry            *gogistryEmitters.TobeSendRequestErrorEmitter
}

func (c *Client) newEvent() *events {
	oldEvents := c.c2sRegistrant.Events
	newEvents := &events{
		NewConnection:    emitters.NewAsyncS2CInfoEmitter(),
		UpdateConnection: emitters.NewAsyncS2CInfoEmitter(),
		Disconnection:    emitters.NewAsyncS2CInfoErrorEmitter(),
		Error:            Emitter.NewAsyncErrorEmitter(),
		Retry:            gogistryEmitters.NewAsyncTobeSendRequestErrorEmitter(),
	}
	oldEvents.NewConnection.AddHandler(func(info gogistryProto.RegistryInfo) {
		newEvents.NewConnection.Emit(info.(message.S2CInfo))
	})
	oldEvents.UpdateConnection.AddHandler(func(info gogistryProto.RegistryInfo) {
		newEvents.UpdateConnection.Emit(info.(message.S2CInfo))
	})
	oldEvents.Disconnection.AddHandler(func(info gogistryProto.RegistryInfo, err error) {
		newEvents.Disconnection.Emit(info.(message.S2CInfo), err)
	})
	oldEvents.Retry.AddHandler(func(request gogistryProto.TobeSendRequest, err error) {
		newEvents.Retry.Emit(request, err)
	})
	oldEvents.Error.AddHandler(func(err error) {
		newEvents.Error.Emit(err)
	})
	oldEvents.EnableAll()
	return newEvents
}
func (c *Client) initEvent() {
	c.Events = c.newEvent()
}
