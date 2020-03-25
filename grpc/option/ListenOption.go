package option

type ListenerOption struct {
	S2SListenNetwork        string `yaml:"S2SListenNetwork" usage:"Network type that the S2SRegistry listen on (e.g. tcp, udp)."`
	S2SListenAddr           string `yaml:"S2SListenAddr" usage:"Network address and port that the S2SRegistry listen on."`
	S2CListenNetwork        string `yaml:"S2CListenNetwork" usage:"Network type that the S2CRegistry listen on (e.g. tcp, udp)."`
	S2CListenAddr           string `yaml:"S2CListenAddr" usage:"Network address and port that the S2CRegistry listen on."`
	GraphQueryListenNetwork string `yaml:"GraphQueryListenNetwork" usage:"Network type that the GraphQueryRegistry listen on (e.g. tcp, udp)."`
	GraphQueryListenAddr    string `yaml:"GraphQueryListenAddr" usage:"Network address and port that the GraphQuery server listen on."`
}

func DefaultListenerOption() ListenerOption {
	return ListenerOption{
		S2SListenNetwork: "tcp", S2SListenAddr: "0.0.0.0:4241",
		S2CListenNetwork: "tcp", S2CListenAddr: "0.0.0.0:4240",
		GraphQueryListenNetwork: "tcp", GraphQueryListenAddr: "0.0.0.0:4242",
	}
}
