package server

import "github.com/yindaheng98/gogisnet/protocol"

func (s *Server) GetServerInfo() protocol.ServerInfo {
	return s.GetS2SInfo().ServerInfo
}
func (s *Server) GetS2SInfo() protocol.S2SInfo {
	return s.s2sRegistrant.Info.(protocol.S2SInfo)
}
func (s *Server) GetS2SConnections() []protocol.S2SInfo { //获取所有连接的服务器信息
	return append(s.GetS2SIndegreeConnections(), s.GetS2SOutdegreeConnections()...)
}
func (s *Server) GetS2SIndegreeConnections() []protocol.S2SInfo { //获取入度连接的服务器信息
	RegistrantInfos := s.s2sRegistry.GetConnections() //入度连接即s2sRegistry的连接
	S2SInfos := make([]protocol.S2SInfo, len(RegistrantInfos))
	for i, RegistrantInfo := range RegistrantInfos {
		S2SInfos[i] = RegistrantInfo.(protocol.S2SInfo)
	}
	return S2SInfos
}
func (s *Server) GetS2SOutdegreeConnections() []protocol.S2SInfo { //获取出度连接的服务器信息
	RegistrantInfos := s.s2sRegistrant.GetConnections() //出度连接即s2sRegistrant的连接
	S2SInfos := make([]protocol.S2SInfo, len(RegistrantInfos))
	for i, RegistrantInfo := range RegistrantInfos {
		S2SInfos[i] = RegistrantInfo.(protocol.S2SInfo)
		S2SInfos[i].Candidates = []protocol.S2SInfo{} //清空候选服务器列表，避免信息量过大
	}
	return S2SInfos

}
func (s *Server) GetS2CInfo() protocol.S2CInfo {
	return s.s2cRegistry.Info.(protocol.S2CInfo)
}
func (s *Server) GetC2SConnections() []protocol.C2SInfo { //获取所有连接的客户端信息
	RegistrantInfos := s.s2cRegistry.GetConnections()
	C2SInfos := make([]protocol.C2SInfo, len(RegistrantInfos))
	for i, RegistrantInfo := range RegistrantInfos {
		C2SInfos[i] = RegistrantInfo.(protocol.C2SInfo)
	}
	return C2SInfos
}
