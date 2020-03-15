package message

//Carry on the information of client.
type ClientInfo interface {
	GetClientID() string
	GetServiceType() string
	String() string
}
