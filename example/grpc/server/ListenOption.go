package server

type ListenerOption struct {
	S2SListenNetwork, S2SListenAddr string
	S2CListenNetwork, S2CListenAddr string
}

func DefaultListenerOption() ListenerOption {
	return ListenerOption{
		S2SListenNetwork: "tcp", S2SListenAddr: "locolhost:4241",
		S2CListenNetwork: "tcp", S2CListenAddr: "locolhost:4240",
	}
}
