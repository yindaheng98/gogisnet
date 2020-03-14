package server

import (
	"context"
	"github.com/yindaheng98/gogisnet/protocol"
	gogistryProto "github.com/yindaheng98/gogistry/protocol"
)

func (s *Server) initConnectionLoader() {
	s.s2sRegistry.Events.NewConnection.AddHandler( //有服务端上线
		func(gogistryProto.RegistrantInfo) {
			s.loadFromS2SRegistry() //影响s2sRegistrant的黑名单和s2cRegistry的候选连接列表
		})
	s.s2sRegistry.Events.NewConnection.Enable()
	s.s2sRegistry.Events.Disconnection.AddHandler( //有服务端下线
		func(info gogistryProto.RegistrantInfo) {
			s.loadFromS2SRegistry() //影响s2sRegistrant的黑名单和s2cRegistry的候选连接列表
			//还会影响s2sRegistrant的候选连接列表
			s.s2sRegistrant.AddCandidates(context.Background(), []gogistryProto.RegistryInfo{info.(protocol.S2SInfo)})
		})
	s.s2sRegistry.Events.Disconnection.Enable()
	s.s2sRegistrant.Events.NewConnection.AddHandler( //有服务端上线
		func(gogistryProto.RegistryInfo) {
			s.loadFromS2SRegistrant() //影响s2cRegistry的候选连接列表
		})
	s.s2sRegistrant.Events.NewConnection.Enable()
	s.s2sRegistrant.Events.Disconnection.AddHandler( //有服务端掉线
		func(gogistryProto.RegistryInfo, error) {
			s.loadFromS2SRegistrant() //影响s2cRegistry的候选连接列表
		})
	s.s2sRegistrant.Events.Disconnection.Enable()
	//一旦有连接变化，立马更新黑名单和候选连接列表
}

func (s *Server) loadFromS2SRegistry() { //s2sRegistry的变动会影响s2sRegistrant的黑名单和s2cRegistry的候选连接列表
	s.loadBlacklist()
	s.loadCandidates()
}

func (s *Server) loadFromS2SRegistrant() { //s2sRegistrant的变动只会影响s2cRegistry的候选连接列表
	s.loadCandidates()
}

func (s *Server) loadBlacklist() { //将s2sRegistry的已连接项作为s2sRegistrant的黑名单
	<-s.s2sRegistrant.CandidateBlacklist                                   //阻塞黑名单
	s2sBlacklist := <-s.s2sBlacklist                                       //取出服务器定义的黑名单
	connections := append(s.s2sRegistry.GetConnections(), s2sBlacklist...) //从面向服务器的注册中心中取出已连接的服务器
	blacklist := make([]gogistryProto.RegistryInfo, len(connections)+1)    //构造黑名单
	blacklist[0] = s.s2sRegistry.Info                                      //黑名单首先要包含自己，不能让自己的注册器连自己的注册中心
	for i, connection := range connections {
		blacklist[i+1] = connection.(protocol.S2SInfo) //所有s2sRegistry已连接的服务器都不能再由s2sRegistrant连接
	}
	s.s2sBlacklist <- s2sBlacklist                  //放回黑名单
	s.s2sRegistrant.CandidateBlacklist <- blacklist //放回黑名单
}

func (s *Server) loadCandidates() { //将s2sRegistry和s2sRegistrant的已连接项作为s2sRegistry和s2cRegistry的候选连接列表
	s2sRegistryCandidates := s.GetS2SConnections() //获取s2sRegistry和s2sRegistrant的已连接项作为s2sRegistry候选连接列表
	s2cRegistryCandidates := make([]protocol.S2CInfo, len(s2sRegistryCandidates))
	for j, candidate := range s2sRegistryCandidates { //再逐个获取S2CInfo作为s2cRegistry的候选连接列表
		s2cRegistryCandidates[j] = candidate.S2CInfo
	}
	s2cInfo := s.s2cRegistry.Info.(protocol.S2CInfo)
	s2cInfo.Candidates = s2cRegistryCandidates
	s.s2cRegistry.Info = s2cInfo //然后再赋值
	s2sInfo := s.s2sRegistry.Info.(protocol.S2SInfo)
	s2sInfo.Candidates = s2sRegistryCandidates
	s.s2sRegistry.Info = s2sInfo
}
