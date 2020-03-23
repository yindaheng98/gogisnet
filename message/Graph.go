package message

import "context"

//Graph describes the topology structure of the gogistnet.
//Topology structure of a gogistnet should be able to access from any server in the gogistnet.
type Graph struct {

	//Vertexes contains all the server in the gogistnet.
	//The key of the map is the unique id of the server (S2SInfo.GetServerID())
	Vertexes map[string]Vertex

	//Clients contains all the clients in the gogistnet.
	//The key of the map is the unique id of the client (C2SInfo.GetClientID())
	Clients map[string]C2SInfo
}

//Vertex describes consists of the information of a server,
//the edges connect from the server to other servers,
//and its connecting clients.
type Vertex struct {

	//Information of a server.
	S2SInfo

	//The edges connect from this server to other servers
	//(The unique id of the server who receive the request and send back response).
	EdgeTo []string

	//Unique id of the clients connecting to the server.
	Clients []string
}

//GraphQueryInfo is used to transmit vertex information among servers to generate topology graph.
type GraphQueryInfo struct {

	//Who is this server?
	S2SInfo S2SInfo

	//Which server is connecting to this server?
	Indegree []S2SInfo

	//Which server is this server connecting to?
	Outdegree []S2SInfo

	//Which client is connecting to this server?
	Clients []C2SInfo
}

//GetVertex can construct a Vertex from a GraphQueryInfo.
func (info GraphQueryInfo) GetVertex() Vertex {
	Clients := make([]string, len(info.Clients)) //用info.Clients构造Clients
	for i, c := range info.Clients {
		Clients[i] = c.GetClientID()
	}
	EdgeTo := make([]string, len(info.Outdegree)) //用info.Outdegree构造EdgeTo
	for i, s := range info.Outdegree {
		EdgeTo[i] = s.GetServerID()
	}

	//构造Vertex
	return Vertex{
		S2SInfo: info.S2SInfo,
		EdgeTo:  EdgeTo,
		Clients: Clients,
	}
}

//GraphQueryProtocol is used to query vertex information among servers to generate topology graph.
type GraphQueryProtocol interface {
	Query(context.Context, S2SInfo) (*GraphQueryInfo, error)
}

//GraphQuerySendOption is the option for sending messages in GraphQuery service.
type GraphQuerySendOption interface{}
