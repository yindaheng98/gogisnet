package server

import (
	"github.com/yindaheng98/gogisnet/protocol"
	gogistry "github.com/yindaheng98/gogistry"
	gogistryProto "github.com/yindaheng98/gogistry/protocol"
	"github.com/yindaheng98/gogistry/registrant"
	"github.com/yindaheng98/gogistry/registry"
	"sync"
	"time"
)

type Server struct {
	s2cRegistry              *registry.Registry
	s2sRegistry              *registry.Registry
	s2sRegistrant            *registrant.Registrant
	s2sBlacklist             chan []gogistryProto.RegistrantInfo
	s2sRegistrantStopChan    chan bool
	s2sRegistrantStoppedChan chan bool
	Events                   *events
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
	s.s2sBlacklist = make(chan []gogistryProto.RegistrantInfo, 1)
	s.s2sBlacklist <- []gogistryProto.RegistrantInfo{}
	s.s2sRegistrantStopChan = make(chan bool, 1)
	close(s.s2sRegistrantStopChan)
	s.s2sRegistrantStoppedChan = make(chan bool, 1)
	close(s.s2sRegistrantStoppedChan)
	s.initConnectionLoader()
	s.initEvents()
	return s
}

func (s *Server) Run() {
	s.s2sRegistrantStopChan = make(chan bool, 1)
	s.s2sRegistrantStoppedChan = make(chan bool, 1)
	wg := new(sync.WaitGroup)
	wg.Add(3)
	go func() {
		defer wg.Done()
		s.s2sRegistry.Run()
	}()
	go func() {
		defer wg.Done()
		for {
			stoppedChan := make(chan bool, 1)
			go func() {
				s.s2sRegistrant.Run()
				stoppedChan <- true
				close(stoppedChan)
			}()
			select {
			case <-stoppedChan: //s2sRegistrant正常退出则继续循环
			case <-s.s2sRegistrantStopChan: //server要求停止
				s.s2sRegistrant.Stop() //就让s2sRegistrant停止
				<-stoppedChan          //并等待s2sRegistrant停止
				s.s2sRegistrantStoppedChan <- true
				close(s.s2sRegistrantStoppedChan)
				return //然后退出
			}
		}
	}()
	go func() {
		defer wg.Done()
		s.s2cRegistry.Run()
	}()
	wg.Wait()
}

func (s *Server) Stop() {
	wg := new(sync.WaitGroup)
	wg.Add(3)
	go func() {
		defer wg.Done()
		s.s2sRegistry.Stop()
	}()
	go func() {
		defer wg.Done()
		s.s2sRegistrantStopChan <- true
		close(s.s2sRegistrantStopChan)
		<-s.s2sRegistrantStoppedChan
	}()
	go func() {
		defer wg.Done()
		s.s2cRegistry.Stop()
	}()
	wg.Wait()
}

func (s *Server) SetS2SWatchdogTimeDelta(t time.Duration) {
	s.s2sRegistrant.WatchdogTimeDelta = t
}

func (s *Server) SetS2SCandidateBlacklist(blacklist []protocol.ServerInfo) {
	<-s.s2sBlacklist
	CandidateBlacklist := make([]gogistryProto.RegistrantInfo, len(blacklist))
	for i, c := range blacklist {
		CandidateBlacklist[i] = protocol.S2SInfo{ServerInfo: c}
	}
	s.s2sBlacklist <- CandidateBlacklist
}
