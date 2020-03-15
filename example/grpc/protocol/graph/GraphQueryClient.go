package graph

import (
	"context"
	pb "github.com/yindaheng98/gogisnet/example/grpc/protocol/protobuf"
	"github.com/yindaheng98/gogisnet/example/grpc/protocol/registrant"
	"github.com/yindaheng98/gogisnet/message"
	"google.golang.org/grpc"
)

//GraphQueryClient is a grpc client for GraphQuery service.
type GraphQueryClient struct {
	pool *registrant.ConnectionPool

	//CallOption is the options when call the RPC.
	CallOption []grpc.CallOption
}

//NewGraphQueryClient returns the pointer to a GraphQueryClient
func NewGraphQueryClient(option GraphQueryClientOption) *GraphQueryClient {
	return &GraphQueryClient{registrant.NewConnectionPool(option.DialOption, option.MaxDialHoldDuration),
		option.CallOption}
}

func (c *GraphQueryClient) getClient(addr string) (client pb.GraphQueryClient, err error) {
	conn, err := c.pool.GetClientConn(addr)
	if err != nil {
		return
	}
	client = pb.NewGraphQueryClient(conn)
	return
}

//Generate a GraphQueryProtocol used in server.
func (c *GraphQueryClient) NewGraphQueryProtocol() GraphQueryProtocol {
	return GraphQueryProtocol{clients: c, CallOption: c.CallOption}
}

//Implementation of message.GraphQueryProtocol, based on GraphQueryClient.
//About message.GraphQueryProtocol: https://godoc.org/github.com/yindaheng98/gogisnet/message#GraphQueryProtocol
type GraphQueryProtocol struct {
	clients    *GraphQueryClient
	CallOption []grpc.CallOption
}

//Implementation of message.GraphQueryProtocol.Query
func (p GraphQueryProtocol) Query(ctx context.Context, info message.ServerInfo) (ginfo *message.GraphQueryInfo, err error) {
	client, err := p.clients.getClient(info.(*pb.ServerInfo).GraphQueryAddr)
	if err != nil {
		return
	}
	gi, err := client.Query(ctx, &pb.Empty{}, p.CallOption...)
	if err != nil {
		return
	}
	ginfo, err = gi.Unpack()
	return
}
