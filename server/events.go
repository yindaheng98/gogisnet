package server

import (
	"github.com/yindaheng98/go-utility/Emitter"
	"github.com/yindaheng98/gogisnet/emitters"
	"github.com/yindaheng98/gogisnet/protocol"
	gogistryProto "github.com/yindaheng98/gogistry/protocol"
	gogistryEmitters "github.com/yindaheng98/gogistry/util/emitters"
)

func cascades(list [][]Emitter.Emitter) {
	for _, pair := range list {
		Emitter.Cascade(pair[0], pair[1])
	}
}

func indefiniteCascades(list [][]Emitter.IndefiniteEmitter) {
	for _, pair := range list {
		Emitter.IndefiniteCascade(pair[0], pair[1])
	}
}

type s2sRegistryEvents struct {
	NewConnection     *emitters.S2SInfoEmitter
	UpdateConnection  *emitters.S2SInfoEmitter
	ConnectionTimeout *emitters.S2SInfoEmitter
	Disconnection     *emitters.S2SInfoEmitter
	Error             *Emitter.ErrorEmitter
}

func (s *Server) newS2SRegistryEvents() *s2sRegistryEvents {
	oldEvents := s.s2sRegistry.Events
	newEvents := &s2sRegistryEvents{
		NewConnection:     emitters.NewAsyncS2SInfoEmitter(),
		UpdateConnection:  emitters.NewAsyncS2SInfoEmitter(),
		ConnectionTimeout: emitters.NewAsyncS2SInfoEmitter(),
		Disconnection:     emitters.NewAsyncS2SInfoEmitter(),
		Error:             Emitter.NewAsyncErrorEmitter(),
	}
	cascades([][]Emitter.Emitter{
		{oldEvents.NewConnection.Emitter, newEvents.NewConnection.Emitter},
		{oldEvents.UpdateConnection.Emitter, newEvents.UpdateConnection.Emitter},
		{oldEvents.ConnectionTimeout.Emitter, newEvents.ConnectionTimeout.Emitter},
		{oldEvents.Disconnection.Emitter, newEvents.Disconnection.Emitter},
		{oldEvents.Error.Emitter, newEvents.Error.Emitter},
	})
	oldEvents.EnableAll()
	return newEvents
}

type s2sRegistrantEvents struct {
	NewConnection    *emitters.S2SInfoEmitter
	UpdateConnection *emitters.S2SInfoEmitter
	Disconnection    *emitters.S2SInfoErrorEmitter
	Error            *Emitter.ErrorEmitter
	Retry            *gogistryEmitters.TobeSendRequestErrorEmitter
}

func (s *Server) newS2SRegistrantEvents() *s2sRegistrantEvents {
	oldEvents := s.s2sRegistrant.Events
	newEvents := &s2sRegistrantEvents{
		NewConnection:    emitters.NewAsyncS2SInfoEmitter(),
		UpdateConnection: emitters.NewAsyncS2SInfoEmitter(),
		Disconnection:    emitters.NewAsyncS2SInfoErrorEmitter(),
		Error:            Emitter.NewAsyncErrorEmitter(),
		Retry:            gogistryEmitters.NewAsyncTobeSendRequestErrorEmitter(),
	}
	cascades([][]Emitter.Emitter{
		{oldEvents.NewConnection.Emitter, newEvents.NewConnection.Emitter},
		{oldEvents.UpdateConnection.Emitter, newEvents.UpdateConnection.Emitter},
		{oldEvents.Error.Emitter, newEvents.Error.Emitter},
	})
	indefiniteCascades([][]Emitter.IndefiniteEmitter{
		{oldEvents.Disconnection.IndefiniteEmitter, newEvents.Disconnection.IndefiniteEmitter},
		{oldEvents.Retry.IndefiniteEmitter, newEvents.Retry.IndefiniteEmitter},
	})
	oldEvents.EnableAll()
	return newEvents
}

type s2cRegistryEvents struct {
	NewConnection     *emitters.C2SInfoEmitter
	UpdateConnection  *emitters.C2SInfoEmitter
	ConnectionTimeout *emitters.C2SInfoEmitter
	Disconnection     *emitters.C2SInfoEmitter
	Error             *Emitter.ErrorEmitter
}

func (s *Server) newS2CRegistryEvents() *s2cRegistryEvents {
	oldEvents := s.s2cRegistry.Events
	newEvents := &s2cRegistryEvents{
		NewConnection:     emitters.NewAsyncC2SInfoEmitter(),
		UpdateConnection:  emitters.NewAsyncC2SInfoEmitter(),
		ConnectionTimeout: emitters.NewAsyncC2SInfoEmitter(),
		Disconnection:     emitters.NewAsyncC2SInfoEmitter(),
		Error:             Emitter.NewAsyncErrorEmitter(),
	}
	cascades([][]Emitter.Emitter{
		{oldEvents.NewConnection.Emitter, newEvents.NewConnection.Emitter},
		{oldEvents.UpdateConnection.Emitter, newEvents.UpdateConnection.Emitter},
		{oldEvents.ConnectionTimeout.Emitter, newEvents.ConnectionTimeout.Emitter},
		{oldEvents.Disconnection.Emitter, newEvents.Disconnection.Emitter},
		{oldEvents.Error.Emitter, newEvents.Error.Emitter},
	})
	oldEvents.EnableAll()
	return newEvents
}

type events struct {
	ServerNewConnection *emitters.ServerInfoEmitter
	ServerDisconnection *emitters.ServerInfoEmitter
	ClientNewConnection *emitters.ClientInfoEmitter
	ClientDisconnection *emitters.ClientInfoEmitter
	S2SRegistryEvent    *s2sRegistryEvents
	S2SRegistrantEvent  *s2sRegistrantEvents
	S2CRegistryEvent    *s2cRegistryEvents
}

func (s *Server) newEvents() *events {
	newEvents := &events{
		ServerNewConnection: emitters.NewAsyncServerInfoEmitter(),
		ServerDisconnection: emitters.NewAsyncServerInfoEmitter(),
		ClientNewConnection: emitters.NewAsyncClientInfoEmitter(),
		ClientDisconnection: emitters.NewAsyncClientInfoEmitter(),
		S2SRegistryEvent:    s.newS2SRegistryEvents(),
		S2SRegistrantEvent:  s.newS2SRegistrantEvents(),
		S2CRegistryEvent:    s.newS2CRegistryEvents(),
	}
	s.s2sRegistry.Events.NewConnection.AddHandler(func(info gogistryProto.RegistrantInfo) {
		newEvents.ServerNewConnection.Emit(info.(protocol.S2SInfo).ServerInfo)
	})
	s.s2sRegistrant.Events.NewConnection.AddHandler(func(info gogistryProto.RegistryInfo) {
		newEvents.ServerNewConnection.Emit(info.(protocol.S2SInfo).ServerInfo)
	})
	s.s2sRegistry.Events.Disconnection.AddHandler(func(info gogistryProto.RegistrantInfo) {
		newEvents.ServerDisconnection.Emit(info.(protocol.S2SInfo).ServerInfo)
	})
	s.s2sRegistrant.Events.Disconnection.AddHandler(func(info gogistryProto.RegistryInfo, err error) {
		newEvents.ServerDisconnection.Emit(info.(protocol.S2SInfo).ServerInfo)
	})
	s.s2cRegistry.Events.NewConnection.AddHandler(func(info gogistryProto.RegistrantInfo) {
		newEvents.ClientNewConnection.Emit(info.(protocol.C2SInfo).ClientInfo)
	})
	s.s2cRegistry.Events.Disconnection.AddHandler(func(info gogistryProto.RegistrantInfo) {
		newEvents.ClientDisconnection.Emit(info.(protocol.C2SInfo).ClientInfo)
	})
	return newEvents
}
func (s *Server) initEvents() {
	s.Events = s.newEvents()
}
