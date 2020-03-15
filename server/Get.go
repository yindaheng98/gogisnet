package server

import (
	"github.com/yindaheng98/gogisnet/message"
)

func (s *Server) GetServerInfo() message.ServerInfo {
	return s.GetS2SInfo().ServerInfo
}
func (s *Server) GetS2SInfo() message.S2SInfo {
	return s.s2sRegistrant.Info.(message.S2SInfo)
}
func (s *Server) GetS2SConnections() []message.S2SInfo { //获取所有连接的服务器信息
	return append(s.GetS2SIndegreeConnections(), s.GetS2SOutdegreeConnections()...)
}
func (s *Server) GetS2SIndegreeConnections() []message.S2SInfo { //获取入度连接的服务器信息
	RegistrantInfos := s.s2sRegistry.GetConnections() //入度连接即s2sRegistry的连接
	S2SInfos := make([]message.S2SInfo, len(RegistrantInfos))
	for i, RegistrantInfo := range RegistrantInfos {
		S2SInfos[i] = RegistrantInfo.(message.S2SInfo)
	}
	return S2SInfos
}
func (s *Server) GetS2SOutdegreeConnections() []message.S2SInfo { //获取出度连接的服务器信息
	RegistrantInfos := s.s2sRegistrant.GetConnections() //出度连接即s2sRegistrant的连接
	S2SInfos := make([]message.S2SInfo, len(RegistrantInfos))
	for i, RegistrantInfo := range RegistrantInfos {
		S2SInfos[i] = RegistrantInfo.(message.S2SInfo)
		S2SInfos[i].Candidates = []message.S2SInfo{} //清空候选服务器列表，避免信息量过大
	}
	return S2SInfos

}
func (s *Server) GetS2CInfo() message.S2CInfo {
	return s.s2cRegistry.Info.(message.S2CInfo)
}
func (s *Server) GetC2SConnections() []message.C2SInfo { //获取所有连接的客户端信息
	RegistrantInfos := s.s2cRegistry.GetConnections()
	C2SInfos := make([]message.C2SInfo, len(RegistrantInfos))
	for i, RegistrantInfo := range RegistrantInfos {
		C2SInfos[i] = RegistrantInfo.(message.C2SInfo)
	}
	return C2SInfos
}
