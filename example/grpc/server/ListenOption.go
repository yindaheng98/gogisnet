package server

type ListenerOption struct {
	S2SListenNetwork, S2SListenAddr               string
	S2CListenNetwork, S2CListenAddr               string
	GraphQueryListenNetwork, GraphQueryListenAddr string
}

func DefaultListenerOption() ListenerOption {
	return ListenerOption{
		S2SListenNetwork: "tcp", S2SListenAddr: "localhost:4241",
		S2CListenNetwork: "tcp", S2CListenAddr: "localhost:4240",
		GraphQueryListenNetwork: "tcp", GraphQueryListenAddr: "localhost:4242",
	}
}
