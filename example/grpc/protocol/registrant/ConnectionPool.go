package registrant

import (
	"google.golang.org/grpc"
	"sync"
	"time"
)

type ConnectionPool struct {
	connections         map[string]*grpc.ClientConn
	connectionsMu       *sync.RWMutex
	DialOption          []grpc.DialOption
	MaxDialHoldDuration time.Duration
}

func NewConnectionPool(DialOption []grpc.DialOption, MaxDialHoldDuration time.Duration) *ConnectionPool {
	return &ConnectionPool{make(map[string]*grpc.ClientConn), new(sync.RWMutex),
		DialOption, MaxDialHoldDuration}
}

func (c *ConnectionPool) GetClientConn(addr string) (connection *grpc.ClientConn, err error) {
	c.connectionsMu.RLock()
	connection, ok := c.connections[addr]
	c.connectionsMu.RUnlock()
	if !ok { //如果没有
		connection, err = grpc.Dial(addr, c.DialOption...) //就新建
		if err == nil {
			c.connectionsMu.Lock()
			c.connections[addr] = connection
			go func() { //过一段时间后删除
				time.Sleep(c.MaxDialHoldDuration)
				c.connectionsMu.Lock()
				defer c.connectionsMu.Unlock()
				delete(c.connections, addr)
			}()
			c.connectionsMu.Unlock()
		}
	}
	return
}
