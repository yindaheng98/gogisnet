package protocol

type ClientInfo interface {
	GetClientID() string
	GetServiceType() string
	String() string
}
