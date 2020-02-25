package server

import (
	"github.com/yindaheng98/gogisnet/protocol"
	gogistryProto "github.com/yindaheng98/gogistry/protocol"
)

func (s *Server) initConnectionLoader() {
	loadFromS2SRegistry := func(gogistryProto.RegistrantInfo) { s.loadFromS2SRegistry() }
	loadFromS2SRegistrant := func(gogistryProto.RegistryInfo) { s.loadFromS2SRegistrant() }
	loadFromS2SRegistrant2 := func(gogistryProto.TobeSendRequest, error) { s.loadFromS2SRegistrant() }
	s.s2sRegistry.Events.NewConnection.AddHandler(loadFromS2SRegistry)
	s.s2sRegistry.Events.NewConnection.Enable()
	s.s2sRegistry.Events.Disconnection.AddHandler(loadFromS2SRegistry)
	s.s2sRegistry.Events.Disconnection.Enable()
	s.s2sRegistrant.Events.NewConnection.AddHandler(loadFromS2SRegistrant)
	s.s2sRegistrant.Events.NewConnection.Enable()
	s.s2sRegistrant.Events.Disconnection.AddHandler(loadFromS2SRegistrant2)
	s.s2sRegistrant.Events.Disconnection.Enable()
}

func (s *Server) loadFromS2SRegistry() { //s2sRegistry的变动会影响s2sRegistrant的黑名单和s2cRegistry的候选连接列表
	s.loadBlacklist()
	s.loadCandidates()
}

func (s *Server) loadFromS2SRegistrant() { //s2sRegistrant的变动只会影响s2cRegistry的候选连接列表
	s.loadCandidates()
}

func (s *Server) loadBlacklist() { //将s2sRegistry的已连接项作为s2sRegistrant的黑名单
	<-s.s2sRegistrant.CandidateBlacklist                                //阻塞黑名单
	connections := s.s2sRegistry.GetConnections()                       //从面向服务器的注册中心中取出已连接的服务器
	blacklist := make([]gogistryProto.RegistryInfo, len(connections)+1) //构造黑名单
	blacklist[0] = s.s2sRegistry.Info                                   //黑名单首先要包含自己，不能让自己的注册器连自己的注册中心
	for i, connection := range connections {
		blacklist[i+1] = connection.(protocol.S2SInfo) //所有s2sRegistry已连接的服务器都不能再由s2sRegistrant连接
	}
	s.s2sRegistrant.CandidateBlacklist <- blacklist //放回黑名单
}

func (s *Server) loadCandidates() { //将s2sRegistry和s2sRegistrant的已连接项作为s2cRegistry的候选连接列表
	registrantConnections := s.s2sRegistrant.GetConnections()
	registryConnections := s.s2sRegistry.GetConnections()
	candidates := make([]protocol.S2CInfo, len(registrantConnections)+len(registryConnections))
	for j, connection := range registrantConnections {
		candidates[j] = connection.(protocol.S2SInfo).S2CInfo
	}
	l := len(registrantConnections)
	for j, connection := range registryConnections {
		candidates[j+l] = connection.(protocol.S2SInfo).S2CInfo
	}
	//将s2sRegistrant和s2sRegistry两边的连接全部集中到一个里面
	info := s.s2cRegistry.Info.(protocol.S2CInfo)
	info.Candidates = candidates
	s.s2cRegistry.Info = info //然后再赋值
}
