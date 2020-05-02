package server

import (
	"context"
	"github.com/yindaheng98/gogisnet/message"
	"sync"
	"time"
)

//Get the topology graph of the whole gogisnet.
func (s *Server) GetGraph(ctx context.Context, QueryTimeout time.Duration) message.Graph {
	graph := message.Graph{
		Vertexes: map[string]message.Vertex{},
		Clients:  map[string]message.C2SInfo{},
	}
	if s.GraphQueryProtocol == nil {
		return graph
	}

	graphChan := make(chan *message.Graph, 1)
	travelledChan := make(chan map[string]struct{}, 1)
	graphChan <- &graph
	GraphQueryInfo := s.GetGraphQueryInfo()
	travelledChan <- map[string]struct{}{GraphQueryInfo.GetVertex().GetServerID(): {}}
	constructGraph(ctx, GraphQueryInfo, graphChan, travelledChan, s.GraphQueryProtocol, QueryTimeout)
	return *<-graphChan
}

//以深度优先遍历构造Graph
func constructGraph(ctx context.Context, info message.GraphQueryInfo,
	graphChan chan *message.Graph, travelledChan chan map[string]struct{},
	proto message.GraphQueryProtocol, QueryTimeout time.Duration) {
	select {
	case <-ctx.Done():
		return //超时
	case graph := <-graphChan: //等待其他进程完成graph处理
		Vertex := info.GetVertex()                    //构造Vertex
		graph.Vertexes[Vertex.GetServerID()] = Vertex //给graph的Vertex表赋值
		for _, c := range info.Clients {
			graph.Clients[c.GetClientID()] = c //更新graph中的客户端列表
		}
		graphChan <- graph //结束后放回graph
	}

	select {
	case <-ctx.Done():
		return //超时
	case travelled := <-travelledChan: //深度优先遍历
		wg := new(sync.WaitGroup)                                    //用于等待线程完成
		for _, s := range append(info.Indegree, info.Outdegree...) { //深度优先遍历
			if _, ok := travelled[s.GetServerID()]; !ok { //如果没有遍历过
				wg.Add(1)
				travelled[s.GetServerID()] = struct{}{} //就占位遍历标记
				go func(s message.S2SInfo) {            //然后递归进行遍历
					defer wg.Done() //结束时报告线程结束
					queryCtx, _ := context.WithTimeout(ctx, QueryTimeout)
					i, err := proto.Query(queryCtx, s) //获取信息
					if err == nil && i != nil {        //如果成功
						constructGraph(ctx, *i, graphChan, travelledChan, proto, QueryTimeout) //就递归
					}
				}(s)
			}
		}
		travelledChan <- travelled //结束后放回遍历标记
		wg.Wait()                  //等待完成
	}

}

//Get the GraphQueryInfo of this server.
func (s *Server) GetGraphQueryInfo() message.GraphQueryInfo {
	return message.GraphQueryInfo{
		S2SInfo:   s.GetS2SInfo(),
		Indegree:  s.GetS2SIndegreeConnections(),
		Outdegree: s.GetS2SOutdegreeConnections(),
		Clients:   s.GetC2SConnections(),
	}
}
