package server

import (
	"context"
	"github.com/yindaheng98/gogisnet/message"
	gogistryProto "github.com/yindaheng98/gogistry/protocol"
)

func (s *Server) initConnectionLoader() {
	//一旦有连接变化，立马更新黑名单和候选连接列表

	blacklist := <-s.s2sRegistrant.CandidateBlacklist
	s.s2sRegistrant.CandidateBlacklist <- append(blacklist, s.GetS2SInfo()) //首先服务器不可以自己连自己

	s.s2sRegistry.Events.NewConnection.AddHandler( //有服务端上线
		func(gogistryProto.RegistrantInfo) {
			s.loadS2SRegistrantBlacklist() //s2sRegistry的变动会影响s2sRegistrant的黑名单
			s.loadCandidates()             //s2sRegistry的变动会影响s2sRegistry和s2cRegistry的候选连接列表
		})
	s.s2sRegistry.Events.NewConnection.Enable()

	s.s2sRegistry.Events.Disconnection.AddHandler( //有服务端下线
		func(info gogistryProto.RegistrantInfo) {
			s2sInfo := info.(message.S2SInfo)
			s.deleteS2SRegistrantBlacklist(s2sInfo) //s2sRegistry的变动会影响s2sRegistrant的黑名单
			s.deleteCandidate(s2sInfo)              //s2sRegistry的变动会影响s2sRegistry和s2cRegistry的候选连接列表

			//还会影响s2sRegistrant的候选连接列表：s2sRegistry断连的放到s2sRegistrant中尝试连接
			s.s2sRegistrant.AddCandidates(context.Background(), []gogistryProto.RegistryInfo{info.(message.S2SInfo)})
		})
	s.s2sRegistry.Events.Disconnection.Enable()

	s.s2sRegistrant.Events.NewConnection.AddHandler( //有服务端上线
		func(gogistryProto.RegistryInfo) {
			s.loadCandidates() //s2sRegistrant的变动会影响s2sRegistry和s2cRegistry的候选连接列表
		})
	s.s2sRegistrant.Events.NewConnection.Enable()

	s.s2sRegistrant.Events.Disconnection.AddHandler( //有服务端掉线
		func(info gogistryProto.RegistryInfo, _ error) {
			s2sInfo := info.(message.S2SInfo)
			s.deleteCandidate(s2sInfo) //s2sRegistrant的变动会影响s2sRegistry和s2cRegistry的候选连接列表
		})
	s.s2sRegistrant.Events.Disconnection.Enable()

}

func (s *Server) loadS2SRegistrantBlacklist() { //将s2sRegistry的已连接项作为s2sRegistrant的黑名单
	<-s.s2sRegistrant.CandidateBlacklist                                //阻塞黑名单
	connections := s.s2sRegistry.GetConnections()                       //从面向服务器的注册中心中取出已连接的服务器
	blacklist := make([]gogistryProto.RegistryInfo, len(connections)+1) //构造黑名单
	blacklist[0] = s.s2sRegistry.Info                                   //黑名单首先要包含自己，不能让自己的注册器连自己的注册中心
	for i, connection := range connections {
		blacklist[i+1] = connection.(message.S2SInfo) //所有s2sRegistry已连接的服务器都不能再由s2sRegistrant连接
	}
	s.s2sRegistrant.CandidateBlacklist <- blacklist //放回黑名单
}

func (s *Server) deleteS2SRegistrantBlacklist(info message.S2SInfo) { //从s2sRegistrant的黑名单中删除一项

	s2sBlacklist := <-s.s2sBlacklist //取出服务器定义的黑名单
	for _, binfo := range s2sBlacklist {
		if binfo.(message.S2SInfo).GetRegistryID() == info.GetRegistryID() { //如果在黑名单里
			s.s2sBlacklist <- s2sBlacklist
			return //就不删除黑名单
		}
	}
	s.s2sBlacklist <- s2sBlacklist

	blacklist := <-s.s2sRegistrant.CandidateBlacklist //阻塞注册器黑名单
	for i, binfo := range blacklist {
		if binfo.(message.S2SInfo).GetRegistryID() == info.GetRegistryID() { //找到匹配的S2SInfo，删除之
			if i+1 < len(blacklist) {
				blacklist = append(blacklist[0:i], blacklist[i+1:]...)
			} else {
				blacklist = blacklist[0:i]
			}
		}
	}
	s.s2sRegistrant.CandidateBlacklist <- blacklist //放回注册器黑名单
}

func (s *Server) loadCandidates() { //将s2sRegistry和s2sRegistrant的已连接项作为s2sRegistry和s2cRegistry的候选连接列表

	s2sRegistryCandidates := s.GetS2SConnections() //获取s2sRegistry和s2sRegistrant的已连接项作为s2sRegistry候选连接列表
	s2sInfo := s.s2sRegistry.Info.(message.S2SInfo)
	s2sInfo.Candidates = s2sRegistryCandidates
	s.s2sRegistry.Info = s2sInfo

	s2cRegistryCandidates := make([]message.S2CInfo, len(s2sRegistryCandidates))
	for j, candidate := range s2sRegistryCandidates { //再逐个获取S2CInfo作为s2cRegistry的候选连接列表
		s2cRegistryCandidates[j] = candidate.S2CInfo
	}
	s2cInfo := s.s2cRegistry.Info.(message.S2CInfo)
	s2cInfo.Candidates = s2cRegistryCandidates
	s.s2cRegistry.Info = s2cInfo //然后再赋值
}

func (s *Server) deleteCandidate(info message.S2SInfo) { //从s2sRegistry和s2cRegistry的候选连接列表中删除一项
	s.deleteS2SCandidate(info)
	s.deleteS2CCandidate(info.S2CInfo)
}

func (s *Server) deleteS2CCandidate(info message.S2CInfo) { //从s2cRegistry的候选连接列表中删除一项
	s2cInfo := s.s2cRegistry.Info.(message.S2CInfo)               //取出注册中心信息
	defer func() { s.s2cRegistry.Info = s2cInfo }()               //放回注册中心信息
	s2cRegistryCandidates := s2cInfo.Candidates                   //取出候选注册中心信息
	defer func() { s2cInfo.Candidates = s2cRegistryCandidates }() //放回候选注册中心信息

	for i, c := range s2cRegistryCandidates { //遍历
		if c.GetRegistryID() == info.GetRegistryID() { //找到匹配的S2CInfo，删除之
			if i+1 < len(s2cRegistryCandidates) {
				s2cRegistryCandidates = append(s2cRegistryCandidates[0:i], s2cRegistryCandidates[i+1:]...)
			} else {
				s2cRegistryCandidates = s2cRegistryCandidates[0:i]
			}
			return
		}
	}
}

func (s *Server) deleteS2SCandidate(info message.S2SInfo) { //从s2sRegistry的候选连接列表中删除一项
	s2sInfo := s.s2sRegistry.Info.(message.S2SInfo)               //取出注册中心信息
	defer func() { s.s2sRegistry.Info = s2sInfo }()               //放回注册中心信息
	s2sRegistryCandidates := s2sInfo.Candidates                   //取出候选注册中心信息
	defer func() { s2sInfo.Candidates = s2sRegistryCandidates }() //放回候选注册中心信息

	for i, c := range s2sRegistryCandidates { //遍历
		if c.GetRegistryID() == info.GetRegistryID() { //找到匹配的S2SInfo，删除之
			if i+1 < len(s2sRegistryCandidates) {
				s2sRegistryCandidates = append(s2sRegistryCandidates[0:i], s2sRegistryCandidates[i+1:]...)
			} else {
				s2sRegistryCandidates = s2sRegistryCandidates[0:i]
			}
			return
		}
	}
}
