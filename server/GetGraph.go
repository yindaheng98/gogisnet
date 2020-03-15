package server

import (
	"context"
	"github.com/yindaheng98/gogisnet/message"
	"sync"
)

//Get the topology graph of the whole gogisnet.
func (s *Server) GetGraph(ctx context.Context) *message.Graph {
	if s.GraphQueryProtocol == nil {
		return nil
	}

	graphChan := make(chan *message.Graph, 1)
	graphChan <- &message.Graph{
		Vertexes: map[string]message.Vertex{},
		Clients:  map[string]message.ClientInfo{},
	}
	constructGraph(ctx, s.GetGraphQueryInfo(), graphChan, s.GraphQueryProtocol)
	return <-graphChan
}

//以深度优先遍历构造Graph
func constructGraph(ctx context.Context, info message.GraphQueryInfo, graphChan chan *message.Graph, proto message.GraphQueryProtocol) {
	var graph *message.Graph
	select {
	case graph = <-graphChan: //等待其他进程完成graph处理
		defer func() { graphChan <- graph }() //结束后放回graph
	case <-ctx.Done():
		return //超时
	}

	//用info.Clients构造Clients
	Clients := make([]string, len(info.Clients))
	for i, c := range info.Clients {
		Clients[i] = c.GetClientID()
		graph.Clients[c.GetClientID()] = c //同时更新graph中的客户端列表
	}

	//用info.Outdegree构造EdgeTo
	EdgeTo := make([]string, len(info.Outdegree))
	for i, s := range info.Outdegree {
		EdgeTo[i] = s.GetServerID()
	}

	//构造Vertex
	Vertex := message.Vertex{
		ServerInfo: info.ServerInfo,
		EdgeTo:     EdgeTo,
		Clients:    Clients,
	}
	graph.Vertexes[Vertex.GetServerID()] = Vertex //给graph的Vertex表赋值

	//深度优先遍历
	wg := new(sync.WaitGroup)                                    //用于等待线程完成
	for _, s := range append(info.Indegree, info.Outdegree...) { //遍历
		if _, ok := graph.Vertexes[s.GetServerID()]; !ok { //如果没有遍历过
			wg.Add(1)
			go func() { //就递归进行遍历
				defer wg.Done() //结束时报告线程结束
				constructGraph(ctx, proto.Query(ctx, s), graphChan, proto)
			}()
		}
	}
	wg.Wait() //等待完成
}

//Get the GraphQueryInfo of this server.
func (s *Server) GetGraphQueryInfo() message.GraphQueryInfo {
	S2SIndegree := s.GetS2SIndegreeConnections()
	Indegree := make([]message.ServerInfo, len(S2SIndegree))
	for i, info := range S2SIndegree {
		Indegree[i] = info.ServerInfo
	}

	S2SOutdegree := s.GetS2SOutdegreeConnections()
	Outdegree := make([]message.ServerInfo, len(S2SOutdegree))
	for i, info := range S2SOutdegree {
		Outdegree[i] = info.ServerInfo
	}

	C2Ss := s.GetC2SConnections()
	Clients := make([]message.ClientInfo, len(C2Ss))
	for i, info := range C2Ss {
		Clients[i] = info.ClientInfo
	}

	return message.GraphQueryInfo{
		ServerInfo: s.GetServerInfo(),
		Indegree:   Indegree,
		Outdegree:  Outdegree,
		Clients:    Clients,
	}
}
