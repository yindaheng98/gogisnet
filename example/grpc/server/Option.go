package server

import (
	grpcServiceClient "github.com/yindaheng98/gogisnet/example/grpc/protocol/client"
	pb "github.com/yindaheng98/gogisnet/example/grpc/protocol/protobuf"
	grpcServiceServer "github.com/yindaheng98/gogisnet/example/grpc/protocol/server"
	"github.com/yindaheng98/gogisnet/server"
)

type GRPCServerNetworkOption struct { //GRPC服务器网络设置
	ListenNetwork string //监听哪种网络
	ListenAddr    string //监听哪个地址
	BoardCastAddr string //向外广播哪个地址
}

type Option struct {
	ServiceOption server.Option
	GRPCOption    GRPCOption
}

type GRPCOption struct {
	S2SServerOption grpcServiceServer.GRPCServerOption
	S2CServerOption grpcServiceServer.GRPCServerOption
	S2SClientOption grpcServiceClient.GRPCClientOption
}

func DefaultOption(S2SNetworkOption, S2CNetworkOption GRPCServerNetworkOption, initServer *pb.S2SInfo) (option Option, err error) {
	option = Option{GRPCOption: GRPCOption{S2SClientOption: grpcServiceClient.DefaultOption()}} //初始化
	var e error
	option.GRPCOption.S2SServerOption, e = grpcServiceServer.DefaultOption(S2SNetworkOption.ListenNetwork, S2SNetworkOption.ListenAddr)
	if e != nil { //S2S服务器GRPC设置
		err = e
	}
	option.GRPCOption.S2CServerOption, e = grpcServiceServer.DefaultOption(S2CNetworkOption.ListenNetwork, S2CNetworkOption.ListenAddr)
	if e != nil { //S2C服务器GRPC设置
		err = e
	}

	if initServer == nil { //构造初始轮询服务器
		initServer = &pb.S2SInfo{
			ServerInfo: &pb.ServerInfo{
				ServerID:       "Undefined",
				ServiceType:    "Undefined",
				AdditionalInfo: "",
			},
			ResponseSendOption: &pb.ResponseSendOption{},
			RequestSendOption:  &pb.RequestSendOption{Addr: "Undefined"},
			Candidates:         nil,
			S2CInfo: &pb.S2CInfo{
				ServerInfo: &pb.ServerInfo{
					ServerID:       "Undefined",
					ServiceType:    "Undefined",
					AdditionalInfo: "",
				},
				RequestSendOption: &pb.RequestSendOption{Addr: "Undefined"},
				Candidates:        nil,
			},
		}
	}
	init, e := initServer.Unpack() //解包初始轮询服务器
	if e == nil {                  //构造服务设置
		option.ServiceOption = server.DefaultOption(*init, nil, nil, nil)
		option.ServiceOption.S2SRegistryOption.RequestSendOption = &pb.RequestSendOption{Addr: S2SNetworkOption.BoardCastAddr}
		option.ServiceOption.S2CRegistryOption.RequestSendOption = &pb.RequestSendOption{Addr: S2CNetworkOption.BoardCastAddr}
		return
	}
	err = e
	return
}
