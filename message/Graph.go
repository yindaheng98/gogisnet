package message

import "context"

//Graph describes the topology structure of the gogistnet.
//Topology structure of a gogistnet should be able to access from any server in the gogistnet.
type Graph struct {

	//Vertexes contains all the server in the gogistnet.
	//The key of the map is the unique id of the server (ServerInfo.GetServerID())
	Vertexes map[string]Vertex

	//Clients contains all the clients in the gogistnet.
	//The key of the map is the unique id of the client (ClientInfo.GetClientID())
	Clients map[string]ClientInfo
}

//Vertex describes consists of the information of a server,
//the edges connect from the server to other servers,
//and its connecting clients.
type Vertex struct {

	//Information of a server.
	ServerInfo

	//The edges connect from this server to other servers
	//(The unique id of the server who receive the request and send back response).
	EdgeTo []string

	//Unique id of the clients connecting to the server.
	Clients []string
}

//GraphQueryInfo is used to transmit vertex information among servers to generate topology graph.
type GraphQueryInfo struct {

	//Who is this server?
	ServerInfo ServerInfo

	//Which server is connecting to this server?
	Indegree []ServerInfo

	//Which server is this server connecting to?
	Outdegree []ServerInfo

	//Which client is connecting to this server?
	Clients []ClientInfo
}

//GraphQueryProtocol is used to query vertex information among servers to generate topology graph.
type GraphQueryProtocol interface {
	Query(context.Context, ServerInfo) (*GraphQueryInfo, error)
}
