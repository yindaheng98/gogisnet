package registrant

import "net"

func parseAddr(addr string, defaultPort int) (*net.TCPAddr, error) {
	if Addr, err := net.ResolveTCPAddr("", addr); err == nil {
		return Addr, nil
	}
	if IP, err := net.ResolveIPAddr("", addr); err == nil {
		return &net.TCPAddr{
			IP:   IP.IP,
			Port: defaultPort,
			Zone: IP.Zone,
		}, nil
	} else {
		return nil, err
	}
}

//Parse the addr for S2S port (default port is 4241).
func ParseS2SAddr(addr string) (*net.TCPAddr, error) {
	return parseAddr(addr, 4241)
}

//Parse the addr for S2S port (default port is 4240).
func ParseS2CAddr(addr string) (*net.TCPAddr, error) {
	return parseAddr(addr, 4240)
}
