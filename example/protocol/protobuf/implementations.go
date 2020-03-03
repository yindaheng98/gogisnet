package protocol

type ServerInfoPointer struct {
	*ServerInfo
}

func (info ServerInfoPointer) GetServerID() string {
	return info.ID
}
func (info ServerInfoPointer) GetServiceType() string {
	return info.Type
}

type ClientInfoPointer struct {
	*ClientInfo
}

func (info ClientInfoPointer) GetClientID() string {
	return info.ID
}
func (info ClientInfoPointer) GetServiceType() string {
	return info.Type
}
