package server

import (
	"github.com/yindaheng98/gogisnet/protocol"
	gogistry "github.com/yindaheng98/gogistry"
	"github.com/yindaheng98/gogistry/registrant"
	"github.com/yindaheng98/gogistry/registry"
	"sync"
)

type Server struct {
	s2cRegistry   *registry.Registry
	s2sRegistry   *registry.Registry
	s2sRegistrant *registrant.Registrant
}

func New(info protocol.ServerInfo, option Option) *Server {
	S2SRegistryOption, S2SRegistrantOption, S2CRegistryOption :=
		option.S2SRegistryOption, option.S2SRegistrantOption, option.S2CRegistryOption
	s2cInfo := protocol.S2CInfo{
		ServerInfo:        info,
		RequestSendOption: S2CRegistryOption.RequestSendOption,
		Candidates:        []protocol.S2CInfo{}}
	s2sInfo := protocol.S2SInfo{
		ServerInfo:         info,
		ResponseSendOption: S2SRegistrantOption.ResponseSendOption,
		RequestSendOption:  S2SRegistryOption.RequestSendOption,
		Candidates:         []protocol.S2SInfo{},
		S2CInfo:            s2cInfo}
	s := &Server{
		s2sRegistry: gogistry.NewRegistry(s2sInfo,
			S2SRegistryOption.MaxRegistrants,
			S2SRegistryOption.TimeoutController,
			S2SRegistryOption.ResponseProto),
		s2sRegistrant: gogistry.NewRegistrant(s2sInfo,
			S2SRegistrantOption.RegistryN,
			S2SRegistrantOption.CandidateList,
			S2SRegistrantOption.RetryNController,
			S2SRegistrantOption.RequestProto),
		s2cRegistry: gogistry.NewRegistry(s2cInfo,
			S2CRegistryOption.MaxRegistrants,
			S2CRegistryOption.TimeoutController,
			S2CRegistryOption.ResponseProto),
	}
	s.initConnectionLoader()
	return s
}

func (s *Server) Run() {
	wg := new(sync.WaitGroup)
	wg.Add(3)
	go func() {
		s.s2sRegistry.Run()
		wg.Done()
	}()
	go func() {
		s.s2sRegistrant.Run()
		wg.Done()
	}()
	go func() {
		s.s2cRegistry.Run()
		wg.Done()
	}()
	wg.Wait()
}

func (s *Server) Stop() {
	wg := new(sync.WaitGroup)
	wg.Add(3)
	go func() {
		s.s2sRegistry.Stop()
		wg.Done()
	}()
	go func() {
		s.s2sRegistrant.Stop()
		wg.Done()
	}()
	go func() {
		s.s2cRegistry.Stop()
		wg.Done()
	}()
	wg.Wait()
}
