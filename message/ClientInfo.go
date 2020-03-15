package message

type ClientInfo interface {
	GetClientID() string
	GetServiceType() string
	String() string
}
