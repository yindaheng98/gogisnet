package message

//Carry on the information of client.
type ServerInfo interface {
	GetServerID() string
	GetServiceType() string //记录服务类型，注册中心和注册器的服务类型必须一致
	String() string
}
