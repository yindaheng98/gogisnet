// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package protocol

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RequestSendOption struct {
	Option               *any.Any `protobuf:"bytes,1,opt,name=Option,proto3" json:"Option,omitempty"`
	Addr                 string   `protobuf:"bytes,2,opt,name=Addr,proto3" json:"Addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestSendOption) Reset()         { *m = RequestSendOption{} }
func (m *RequestSendOption) String() string { return proto.CompactTextString(m) }
func (*RequestSendOption) ProtoMessage()    {}
func (*RequestSendOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *RequestSendOption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestSendOption.Unmarshal(m, b)
}
func (m *RequestSendOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestSendOption.Marshal(b, m, deterministic)
}
func (m *RequestSendOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestSendOption.Merge(m, src)
}
func (m *RequestSendOption) XXX_Size() int {
	return xxx_messageInfo_RequestSendOption.Size(m)
}
func (m *RequestSendOption) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestSendOption.DiscardUnknown(m)
}

var xxx_messageInfo_RequestSendOption proto.InternalMessageInfo

func (m *RequestSendOption) GetOption() *any.Any {
	if m != nil {
		return m.Option
	}
	return nil
}

func (m *RequestSendOption) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

type ResponseSendOption struct {
	Option               *any.Any `protobuf:"bytes,1,opt,name=Option,proto3" json:"Option,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseSendOption) Reset()         { *m = ResponseSendOption{} }
func (m *ResponseSendOption) String() string { return proto.CompactTextString(m) }
func (*ResponseSendOption) ProtoMessage()    {}
func (*ResponseSendOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *ResponseSendOption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseSendOption.Unmarshal(m, b)
}
func (m *ResponseSendOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseSendOption.Marshal(b, m, deterministic)
}
func (m *ResponseSendOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseSendOption.Merge(m, src)
}
func (m *ResponseSendOption) XXX_Size() int {
	return xxx_messageInfo_ResponseSendOption.Size(m)
}
func (m *ResponseSendOption) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseSendOption.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseSendOption proto.InternalMessageInfo

func (m *ResponseSendOption) GetOption() *any.Any {
	if m != nil {
		return m.Option
	}
	return nil
}

type GraphQuerySendOption struct {
	Option               *any.Any `protobuf:"bytes,1,opt,name=Option,proto3" json:"Option,omitempty"`
	Addr                 string   `protobuf:"bytes,2,opt,name=Addr,proto3" json:"Addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GraphQuerySendOption) Reset()         { *m = GraphQuerySendOption{} }
func (m *GraphQuerySendOption) String() string { return proto.CompactTextString(m) }
func (*GraphQuerySendOption) ProtoMessage()    {}
func (*GraphQuerySendOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *GraphQuerySendOption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GraphQuerySendOption.Unmarshal(m, b)
}
func (m *GraphQuerySendOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GraphQuerySendOption.Marshal(b, m, deterministic)
}
func (m *GraphQuerySendOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GraphQuerySendOption.Merge(m, src)
}
func (m *GraphQuerySendOption) XXX_Size() int {
	return xxx_messageInfo_GraphQuerySendOption.Size(m)
}
func (m *GraphQuerySendOption) XXX_DiscardUnknown() {
	xxx_messageInfo_GraphQuerySendOption.DiscardUnknown(m)
}

var xxx_messageInfo_GraphQuerySendOption proto.InternalMessageInfo

func (m *GraphQuerySendOption) GetOption() *any.Any {
	if m != nil {
		return m.Option
	}
	return nil
}

func (m *GraphQuerySendOption) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

type Timestamp struct {
	Timestamp            uint64   `protobuf:"varint,1,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Timestamp) Reset()         { *m = Timestamp{} }
func (m *Timestamp) String() string { return proto.CompactTextString(m) }
func (*Timestamp) ProtoMessage()    {}
func (*Timestamp) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *Timestamp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Timestamp.Unmarshal(m, b)
}
func (m *Timestamp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Timestamp.Marshal(b, m, deterministic)
}
func (m *Timestamp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Timestamp.Merge(m, src)
}
func (m *Timestamp) XXX_Size() int {
	return xxx_messageInfo_Timestamp.Size(m)
}
func (m *Timestamp) XXX_DiscardUnknown() {
	xxx_messageInfo_Timestamp.DiscardUnknown(m)
}

var xxx_messageInfo_Timestamp proto.InternalMessageInfo

func (m *Timestamp) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type ServerInfo struct {
	ServerID             string            `protobuf:"bytes,1,opt,name=ServerID,proto3" json:"ServerID,omitempty"`
	ServiceType          string            `protobuf:"bytes,2,opt,name=ServiceType,proto3" json:"ServiceType,omitempty"`
	AdditionalInfo       map[string]string `protobuf:"bytes,3,rep,name=AdditionalInfo,proto3" json:"AdditionalInfo,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ServerInfo) Reset()         { *m = ServerInfo{} }
func (m *ServerInfo) String() string { return proto.CompactTextString(m) }
func (*ServerInfo) ProtoMessage()    {}
func (*ServerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4}
}

func (m *ServerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerInfo.Unmarshal(m, b)
}
func (m *ServerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerInfo.Marshal(b, m, deterministic)
}
func (m *ServerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerInfo.Merge(m, src)
}
func (m *ServerInfo) XXX_Size() int {
	return xxx_messageInfo_ServerInfo.Size(m)
}
func (m *ServerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ServerInfo proto.InternalMessageInfo

func (m *ServerInfo) GetServerID() string {
	if m != nil {
		return m.ServerID
	}
	return ""
}

func (m *ServerInfo) GetServiceType() string {
	if m != nil {
		return m.ServiceType
	}
	return ""
}

func (m *ServerInfo) GetAdditionalInfo() map[string]string {
	if m != nil {
		return m.AdditionalInfo
	}
	return nil
}

type S2CInfo struct {
	ServerInfo           *ServerInfo        `protobuf:"bytes,1,opt,name=ServerInfo,proto3" json:"ServerInfo,omitempty"`
	RequestSendOption    *RequestSendOption `protobuf:"bytes,2,opt,name=RequestSendOption,proto3" json:"RequestSendOption,omitempty"`
	Candidates           []*S2CInfo         `protobuf:"bytes,3,rep,name=Candidates,proto3" json:"Candidates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *S2CInfo) Reset()         { *m = S2CInfo{} }
func (m *S2CInfo) String() string { return proto.CompactTextString(m) }
func (*S2CInfo) ProtoMessage()    {}
func (*S2CInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{5}
}

func (m *S2CInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2CInfo.Unmarshal(m, b)
}
func (m *S2CInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2CInfo.Marshal(b, m, deterministic)
}
func (m *S2CInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2CInfo.Merge(m, src)
}
func (m *S2CInfo) XXX_Size() int {
	return xxx_messageInfo_S2CInfo.Size(m)
}
func (m *S2CInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_S2CInfo.DiscardUnknown(m)
}

var xxx_messageInfo_S2CInfo proto.InternalMessageInfo

func (m *S2CInfo) GetServerInfo() *ServerInfo {
	if m != nil {
		return m.ServerInfo
	}
	return nil
}

func (m *S2CInfo) GetRequestSendOption() *RequestSendOption {
	if m != nil {
		return m.RequestSendOption
	}
	return nil
}

func (m *S2CInfo) GetCandidates() []*S2CInfo {
	if m != nil {
		return m.Candidates
	}
	return nil
}

type S2SInfo struct {
	ServerInfo           *ServerInfo           `protobuf:"bytes,1,opt,name=ServerInfo,proto3" json:"ServerInfo,omitempty"`
	ResponseSendOption   *ResponseSendOption   `protobuf:"bytes,2,opt,name=ResponseSendOption,proto3" json:"ResponseSendOption,omitempty"`
	RequestSendOption    *RequestSendOption    `protobuf:"bytes,3,opt,name=RequestSendOption,proto3" json:"RequestSendOption,omitempty"`
	GraphQuerySendOption *GraphQuerySendOption `protobuf:"bytes,4,opt,name=GraphQuerySendOption,proto3" json:"GraphQuerySendOption,omitempty"`
	Candidates           []*S2SInfo            `protobuf:"bytes,5,rep,name=Candidates,proto3" json:"Candidates,omitempty"`
	S2CInfo              *S2CInfo              `protobuf:"bytes,6,opt,name=S2CInfo,proto3" json:"S2CInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *S2SInfo) Reset()         { *m = S2SInfo{} }
func (m *S2SInfo) String() string { return proto.CompactTextString(m) }
func (*S2SInfo) ProtoMessage()    {}
func (*S2SInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{6}
}

func (m *S2SInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2SInfo.Unmarshal(m, b)
}
func (m *S2SInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2SInfo.Marshal(b, m, deterministic)
}
func (m *S2SInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2SInfo.Merge(m, src)
}
func (m *S2SInfo) XXX_Size() int {
	return xxx_messageInfo_S2SInfo.Size(m)
}
func (m *S2SInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_S2SInfo.DiscardUnknown(m)
}

var xxx_messageInfo_S2SInfo proto.InternalMessageInfo

func (m *S2SInfo) GetServerInfo() *ServerInfo {
	if m != nil {
		return m.ServerInfo
	}
	return nil
}

func (m *S2SInfo) GetResponseSendOption() *ResponseSendOption {
	if m != nil {
		return m.ResponseSendOption
	}
	return nil
}

func (m *S2SInfo) GetRequestSendOption() *RequestSendOption {
	if m != nil {
		return m.RequestSendOption
	}
	return nil
}

func (m *S2SInfo) GetGraphQuerySendOption() *GraphQuerySendOption {
	if m != nil {
		return m.GraphQuerySendOption
	}
	return nil
}

func (m *S2SInfo) GetCandidates() []*S2SInfo {
	if m != nil {
		return m.Candidates
	}
	return nil
}

func (m *S2SInfo) GetS2CInfo() *S2CInfo {
	if m != nil {
		return m.S2CInfo
	}
	return nil
}

type S2SRequest struct {
	S2SInfo              *S2SInfo `protobuf:"bytes,1,opt,name=S2SInfo,proto3" json:"S2SInfo,omitempty"`
	Disconnect           bool     `protobuf:"varint,2,opt,name=Disconnect,proto3" json:"Disconnect,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *S2SRequest) Reset()         { *m = S2SRequest{} }
func (m *S2SRequest) String() string { return proto.CompactTextString(m) }
func (*S2SRequest) ProtoMessage()    {}
func (*S2SRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{7}
}

func (m *S2SRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2SRequest.Unmarshal(m, b)
}
func (m *S2SRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2SRequest.Marshal(b, m, deterministic)
}
func (m *S2SRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2SRequest.Merge(m, src)
}
func (m *S2SRequest) XXX_Size() int {
	return xxx_messageInfo_S2SRequest.Size(m)
}
func (m *S2SRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_S2SRequest.DiscardUnknown(m)
}

var xxx_messageInfo_S2SRequest proto.InternalMessageInfo

func (m *S2SRequest) GetS2SInfo() *S2SInfo {
	if m != nil {
		return m.S2SInfo
	}
	return nil
}

func (m *S2SRequest) GetDisconnect() bool {
	if m != nil {
		return m.Disconnect
	}
	return false
}

type S2SResponse struct {
	S2SInfo              *S2SInfo `protobuf:"bytes,1,opt,name=S2SInfo,proto3" json:"S2SInfo,omitempty"`
	Timeout              int64    `protobuf:"varint,2,opt,name=Timeout,proto3" json:"Timeout,omitempty"`
	Reject               bool     `protobuf:"varint,3,opt,name=Reject,proto3" json:"Reject,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *S2SResponse) Reset()         { *m = S2SResponse{} }
func (m *S2SResponse) String() string { return proto.CompactTextString(m) }
func (*S2SResponse) ProtoMessage()    {}
func (*S2SResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{8}
}

func (m *S2SResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2SResponse.Unmarshal(m, b)
}
func (m *S2SResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2SResponse.Marshal(b, m, deterministic)
}
func (m *S2SResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2SResponse.Merge(m, src)
}
func (m *S2SResponse) XXX_Size() int {
	return xxx_messageInfo_S2SResponse.Size(m)
}
func (m *S2SResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_S2SResponse.DiscardUnknown(m)
}

var xxx_messageInfo_S2SResponse proto.InternalMessageInfo

func (m *S2SResponse) GetS2SInfo() *S2SInfo {
	if m != nil {
		return m.S2SInfo
	}
	return nil
}

func (m *S2SResponse) GetTimeout() int64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *S2SResponse) GetReject() bool {
	if m != nil {
		return m.Reject
	}
	return false
}

type ClientInfo struct {
	ClientID             string            `protobuf:"bytes,1,opt,name=ClientID,proto3" json:"ClientID,omitempty"`
	ServiceType          string            `protobuf:"bytes,2,opt,name=ServiceType,proto3" json:"ServiceType,omitempty"`
	AdditionalInfo       map[string]string `protobuf:"bytes,3,rep,name=AdditionalInfo,proto3" json:"AdditionalInfo,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ClientInfo) Reset()         { *m = ClientInfo{} }
func (m *ClientInfo) String() string { return proto.CompactTextString(m) }
func (*ClientInfo) ProtoMessage()    {}
func (*ClientInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{9}
}

func (m *ClientInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientInfo.Unmarshal(m, b)
}
func (m *ClientInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientInfo.Marshal(b, m, deterministic)
}
func (m *ClientInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientInfo.Merge(m, src)
}
func (m *ClientInfo) XXX_Size() int {
	return xxx_messageInfo_ClientInfo.Size(m)
}
func (m *ClientInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ClientInfo proto.InternalMessageInfo

func (m *ClientInfo) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *ClientInfo) GetServiceType() string {
	if m != nil {
		return m.ServiceType
	}
	return ""
}

func (m *ClientInfo) GetAdditionalInfo() map[string]string {
	if m != nil {
		return m.AdditionalInfo
	}
	return nil
}

type C2SInfo struct {
	ClientInfo           *ClientInfo         `protobuf:"bytes,1,opt,name=ClientInfo,proto3" json:"ClientInfo,omitempty"`
	ResponseSendOption   *ResponseSendOption `protobuf:"bytes,2,opt,name=ResponseSendOption,proto3" json:"ResponseSendOption,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *C2SInfo) Reset()         { *m = C2SInfo{} }
func (m *C2SInfo) String() string { return proto.CompactTextString(m) }
func (*C2SInfo) ProtoMessage()    {}
func (*C2SInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{10}
}

func (m *C2SInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2SInfo.Unmarshal(m, b)
}
func (m *C2SInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2SInfo.Marshal(b, m, deterministic)
}
func (m *C2SInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2SInfo.Merge(m, src)
}
func (m *C2SInfo) XXX_Size() int {
	return xxx_messageInfo_C2SInfo.Size(m)
}
func (m *C2SInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_C2SInfo.DiscardUnknown(m)
}

var xxx_messageInfo_C2SInfo proto.InternalMessageInfo

func (m *C2SInfo) GetClientInfo() *ClientInfo {
	if m != nil {
		return m.ClientInfo
	}
	return nil
}

func (m *C2SInfo) GetResponseSendOption() *ResponseSendOption {
	if m != nil {
		return m.ResponseSendOption
	}
	return nil
}

type C2SRequest struct {
	C2SInfo              *C2SInfo `protobuf:"bytes,1,opt,name=C2SInfo,proto3" json:"C2SInfo,omitempty"`
	Disconnect           bool     `protobuf:"varint,2,opt,name=Disconnect,proto3" json:"Disconnect,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2SRequest) Reset()         { *m = C2SRequest{} }
func (m *C2SRequest) String() string { return proto.CompactTextString(m) }
func (*C2SRequest) ProtoMessage()    {}
func (*C2SRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{11}
}

func (m *C2SRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2SRequest.Unmarshal(m, b)
}
func (m *C2SRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2SRequest.Marshal(b, m, deterministic)
}
func (m *C2SRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2SRequest.Merge(m, src)
}
func (m *C2SRequest) XXX_Size() int {
	return xxx_messageInfo_C2SRequest.Size(m)
}
func (m *C2SRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_C2SRequest.DiscardUnknown(m)
}

var xxx_messageInfo_C2SRequest proto.InternalMessageInfo

func (m *C2SRequest) GetC2SInfo() *C2SInfo {
	if m != nil {
		return m.C2SInfo
	}
	return nil
}

func (m *C2SRequest) GetDisconnect() bool {
	if m != nil {
		return m.Disconnect
	}
	return false
}

type S2CResponse struct {
	S2CInfo              *S2CInfo `protobuf:"bytes,1,opt,name=S2CInfo,proto3" json:"S2CInfo,omitempty"`
	Timeout              int64    `protobuf:"varint,2,opt,name=Timeout,proto3" json:"Timeout,omitempty"`
	Reject               bool     `protobuf:"varint,3,opt,name=Reject,proto3" json:"Reject,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *S2CResponse) Reset()         { *m = S2CResponse{} }
func (m *S2CResponse) String() string { return proto.CompactTextString(m) }
func (*S2CResponse) ProtoMessage()    {}
func (*S2CResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{12}
}

func (m *S2CResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2CResponse.Unmarshal(m, b)
}
func (m *S2CResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2CResponse.Marshal(b, m, deterministic)
}
func (m *S2CResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2CResponse.Merge(m, src)
}
func (m *S2CResponse) XXX_Size() int {
	return xxx_messageInfo_S2CResponse.Size(m)
}
func (m *S2CResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_S2CResponse.DiscardUnknown(m)
}

var xxx_messageInfo_S2CResponse proto.InternalMessageInfo

func (m *S2CResponse) GetS2CInfo() *S2CInfo {
	if m != nil {
		return m.S2CInfo
	}
	return nil
}

func (m *S2CResponse) GetTimeout() int64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *S2CResponse) GetReject() bool {
	if m != nil {
		return m.Reject
	}
	return false
}

func init() {
	proto.RegisterType((*RequestSendOption)(nil), "protocol.RequestSendOption")
	proto.RegisterType((*ResponseSendOption)(nil), "protocol.ResponseSendOption")
	proto.RegisterType((*GraphQuerySendOption)(nil), "protocol.GraphQuerySendOption")
	proto.RegisterType((*Timestamp)(nil), "protocol.Timestamp")
	proto.RegisterType((*ServerInfo)(nil), "protocol.ServerInfo")
	proto.RegisterMapType((map[string]string)(nil), "protocol.ServerInfo.AdditionalInfoEntry")
	proto.RegisterType((*S2CInfo)(nil), "protocol.S2CInfo")
	proto.RegisterType((*S2SInfo)(nil), "protocol.S2SInfo")
	proto.RegisterType((*S2SRequest)(nil), "protocol.S2SRequest")
	proto.RegisterType((*S2SResponse)(nil), "protocol.S2SResponse")
	proto.RegisterType((*ClientInfo)(nil), "protocol.ClientInfo")
	proto.RegisterMapType((map[string]string)(nil), "protocol.ClientInfo.AdditionalInfoEntry")
	proto.RegisterType((*C2SInfo)(nil), "protocol.C2SInfo")
	proto.RegisterType((*C2SRequest)(nil), "protocol.C2SRequest")
	proto.RegisterType((*S2CResponse)(nil), "protocol.S2CResponse")
}

func init() {
	proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626)
}

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 608 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x96, 0xeb, 0x34, 0x6d, 0x26, 0x02, 0xd1, 0x6d, 0x40, 0x26, 0x54, 0x55, 0xe4, 0x53, 0x10,
	0xc8, 0x05, 0x97, 0x03, 0xe2, 0x16, 0xb6, 0xa8, 0x8a, 0x84, 0x20, 0xac, 0x8b, 0x04, 0x47, 0x37,
	0xde, 0x06, 0x83, 0xbb, 0x6b, 0x6c, 0x27, 0x92, 0xdf, 0x01, 0x1e, 0x89, 0x67, 0xe1, 0x55, 0x90,
	0xd7, 0xeb, 0xbf, 0x64, 0x83, 0x94, 0x10, 0x71, 0x8a, 0x67, 0x3c, 0x9e, 0x6f, 0xe6, 0x9b, 0xef,
	0x53, 0xe0, 0x4e, 0x4c, 0xa3, 0x85, 0x3f, 0xa5, 0x56, 0x18, 0xf1, 0x84, 0xa3, 0x43, 0xf1, 0x33,
	0xe5, 0x41, 0xff, 0xe1, 0x8c, 0xf3, 0x59, 0x40, 0xcf, 0x44, 0xe2, 0x7a, 0x7e, 0x73, 0xe6, 0xb2,
	0x34, 0x2f, 0x32, 0x3f, 0xc2, 0x11, 0xa1, 0xdf, 0xe7, 0x34, 0x4e, 0x1c, 0xca, 0xbc, 0xf7, 0x61,
	0xe2, 0x73, 0x86, 0x9e, 0x42, 0x3b, 0x7f, 0x32, 0xb4, 0x81, 0x36, 0xec, 0xda, 0x3d, 0x2b, 0x6f,
	0x60, 0x15, 0x0d, 0xac, 0x11, 0x4b, 0x89, 0xac, 0x41, 0x08, 0x5a, 0x23, 0xcf, 0x8b, 0x8c, 0xbd,
	0x81, 0x36, 0xec, 0x10, 0xf1, 0x6c, 0xbe, 0x06, 0x44, 0x68, 0x1c, 0x72, 0x16, 0xd3, 0x6d, 0xfb,
	0x9a, 0x9f, 0xa0, 0x77, 0x19, 0xb9, 0xe1, 0x97, 0x0f, 0x73, 0x1a, 0xa5, 0x3b, 0x9d, 0xee, 0x31,
	0x74, 0xae, 0xfc, 0x5b, 0x1a, 0x27, 0xee, 0x6d, 0x88, 0x4e, 0x6a, 0x81, 0xe8, 0xd8, 0x22, 0x55,
	0xc2, 0xfc, 0xad, 0x01, 0x38, 0x34, 0x5a, 0xd0, 0x68, 0xcc, 0x6e, 0x38, 0xea, 0xc3, 0xa1, 0x8c,
	0x2e, 0x44, 0x6d, 0x87, 0x94, 0x31, 0x1a, 0x40, 0xd7, 0xc9, 0x0f, 0x70, 0x95, 0x86, 0x54, 0x02,
	0xd6, 0x53, 0x68, 0x02, 0x77, 0x47, 0x9e, 0xe7, 0x67, 0x73, 0xb9, 0x41, 0xd6, 0xcf, 0xd0, 0x07,
	0xfa, 0xb0, 0x6b, 0x0f, 0xad, 0xe2, 0x54, 0x56, 0x85, 0x65, 0x35, 0x4b, 0xdf, 0xb0, 0x24, 0x4a,
	0xc9, 0xd2, 0xf7, 0xfd, 0x11, 0x1c, 0x2b, 0xca, 0xd0, 0x3d, 0xd0, 0xbf, 0xd1, 0x54, 0x4e, 0x98,
	0x3d, 0xa2, 0x1e, 0xec, 0x2f, 0xdc, 0x60, 0x5e, 0x8c, 0x95, 0x07, 0xaf, 0xf6, 0x5e, 0x6a, 0xe6,
	0x2f, 0x0d, 0x0e, 0x1c, 0x1b, 0x8b, 0xf5, 0x5e, 0xd4, 0x97, 0x2d, 0xe9, 0x55, 0x0c, 0x47, 0xea,
	0xa4, 0x8c, 0x15, 0x1a, 0x12, 0x38, 0x5d, 0xfb, 0x51, 0xf5, 0xf1, 0x4a, 0x09, 0x51, 0x28, 0xef,
	0x39, 0x00, 0x76, 0x99, 0xe7, 0x7b, 0x6e, 0x42, 0x63, 0xc9, 0xce, 0x51, 0x6d, 0x80, 0x7c, 0x4e,
	0x52, 0x2b, 0x32, 0x7f, 0xe8, 0xd9, 0xfc, 0xce, 0x3f, 0xcc, 0xff, 0x56, 0x25, 0x56, 0xb9, 0xc0,
	0x49, 0x7d, 0x81, 0xe5, 0x1a, 0xa2, 0x12, 0xb9, 0x92, 0x0d, 0x7d, 0x2b, 0x36, 0x88, 0xda, 0x01,
	0x46, 0x4b, 0x74, 0x3b, 0xad, 0xba, 0xa9, 0xaa, 0x88, 0xda, 0x3d, 0x4d, 0x86, 0xf7, 0x57, 0x19,
	0x76, 0x96, 0x19, 0x46, 0x4f, 0x4a, 0x81, 0x18, 0x6d, 0x81, 0xac, 0xb8, 0x48, 0x51, 0x61, 0x7e,
	0x06, 0x70, 0x6c, 0x47, 0xee, 0x92, 0x7f, 0xea, 0xd4, 0xae, 0xa1, 0x80, 0x2a, 0xaf, 0x77, 0x0a,
	0x70, 0xe1, 0xc7, 0x53, 0xce, 0x18, 0x9d, 0x26, 0x82, 0xff, 0x43, 0x52, 0xcb, 0x98, 0x01, 0x74,
	0x45, 0xeb, 0x9c, 0xf2, 0xcd, 0x7a, 0x1b, 0x70, 0x90, 0x99, 0x9a, 0xcf, 0xf3, 0xc6, 0x3a, 0x29,
	0x42, 0xf4, 0x00, 0xda, 0x84, 0x7e, 0xcd, 0x10, 0x75, 0x81, 0x28, 0x23, 0xe1, 0x7c, 0x1c, 0xf8,
	0x94, 0x25, 0x85, 0xf3, 0x65, 0x54, 0x3a, 0xbf, 0x88, 0x77, 0xe3, 0xfc, 0x0a, 0xeb, 0x7f, 0x39,
	0xff, 0xa7, 0x06, 0x07, 0xb8, 0x72, 0x4e, 0x35, 0xc0, 0xaa, 0x73, 0xaa, 0x77, 0xa4, 0x4e, 0xca,
	0x4e, 0x9d, 0x93, 0x49, 0x07, 0x37, 0xa4, 0x83, 0xd7, 0x9d, 0x17, 0x17, 0xe7, 0xc5, 0x9b, 0x48,
	0x07, 0x37, 0xa5, 0x83, 0xd7, 0x49, 0xa7, 0xa9, 0xe8, 0xcd, 0xa5, 0x63, 0x27, 0x52, 0xa8, 0x33,
	0x3f, 0xce, 0x6e, 0x72, 0x0e, 0xad, 0x09, 0x0f, 0x02, 0xd4, 0x6b, 0xe8, 0x53, 0xee, 0xd9, 0xbf,
	0xbf, 0x94, 0x95, 0x23, 0x3e, 0x83, 0xd6, 0x64, 0xfc, 0xee, 0x12, 0x1d, 0x57, 0xaf, 0xcb, 0x7f,
	0xa5, 0xbe, 0x2a, 0x99, 0xa3, 0xe2, 0xbf, 0xa1, 0xe2, 0x35, 0xa8, 0x78, 0x7b, 0xd4, 0xeb, 0xb6,
	0xc8, 0x9d, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x76, 0x49, 0x07, 0xf4, 0x7d, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// S2SRegistryClient is the client API for S2SRegistry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type S2SRegistryClient interface {
	Poll(ctx context.Context, in *S2SRequest, opts ...grpc.CallOption) (*S2SResponse, error)
	PING(ctx context.Context, in *Timestamp, opts ...grpc.CallOption) (*Timestamp, error)
}

type s2SRegistryClient struct {
	cc grpc.ClientConnInterface
}

func NewS2SRegistryClient(cc grpc.ClientConnInterface) S2SRegistryClient {
	return &s2SRegistryClient{cc}
}

func (c *s2SRegistryClient) Poll(ctx context.Context, in *S2SRequest, opts ...grpc.CallOption) (*S2SResponse, error) {
	out := new(S2SResponse)
	err := c.cc.Invoke(ctx, "/protocol.S2SRegistry/Poll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s2SRegistryClient) PING(ctx context.Context, in *Timestamp, opts ...grpc.CallOption) (*Timestamp, error) {
	out := new(Timestamp)
	err := c.cc.Invoke(ctx, "/protocol.S2SRegistry/PING", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// S2SRegistryServer is the server API for S2SRegistry service.
type S2SRegistryServer interface {
	Poll(context.Context, *S2SRequest) (*S2SResponse, error)
	PING(context.Context, *Timestamp) (*Timestamp, error)
}

// UnimplementedS2SRegistryServer can be embedded to have forward compatible implementations.
type UnimplementedS2SRegistryServer struct {
}

func (*UnimplementedS2SRegistryServer) Poll(ctx context.Context, req *S2SRequest) (*S2SResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Poll not implemented")
}
func (*UnimplementedS2SRegistryServer) PING(ctx context.Context, req *Timestamp) (*Timestamp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PING not implemented")
}

func RegisterS2SRegistryServer(s *grpc.Server, srv S2SRegistryServer) {
	s.RegisterService(&_S2SRegistry_serviceDesc, srv)
}

func _S2SRegistry_Poll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(S2SRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S2SRegistryServer).Poll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.S2SRegistry/Poll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S2SRegistryServer).Poll(ctx, req.(*S2SRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _S2SRegistry_PING_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Timestamp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S2SRegistryServer).PING(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.S2SRegistry/PING",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S2SRegistryServer).PING(ctx, req.(*Timestamp))
	}
	return interceptor(ctx, in, info, handler)
}

var _S2SRegistry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.S2SRegistry",
	HandlerType: (*S2SRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Poll",
			Handler:    _S2SRegistry_Poll_Handler,
		},
		{
			MethodName: "PING",
			Handler:    _S2SRegistry_PING_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

// S2CRegistryClient is the client API for S2CRegistry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type S2CRegistryClient interface {
	Poll(ctx context.Context, in *C2SRequest, opts ...grpc.CallOption) (*S2CResponse, error)
	PING(ctx context.Context, in *Timestamp, opts ...grpc.CallOption) (*Timestamp, error)
}

type s2CRegistryClient struct {
	cc grpc.ClientConnInterface
}

func NewS2CRegistryClient(cc grpc.ClientConnInterface) S2CRegistryClient {
	return &s2CRegistryClient{cc}
}

func (c *s2CRegistryClient) Poll(ctx context.Context, in *C2SRequest, opts ...grpc.CallOption) (*S2CResponse, error) {
	out := new(S2CResponse)
	err := c.cc.Invoke(ctx, "/protocol.S2CRegistry/Poll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s2CRegistryClient) PING(ctx context.Context, in *Timestamp, opts ...grpc.CallOption) (*Timestamp, error) {
	out := new(Timestamp)
	err := c.cc.Invoke(ctx, "/protocol.S2CRegistry/PING", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// S2CRegistryServer is the server API for S2CRegistry service.
type S2CRegistryServer interface {
	Poll(context.Context, *C2SRequest) (*S2CResponse, error)
	PING(context.Context, *Timestamp) (*Timestamp, error)
}

// UnimplementedS2CRegistryServer can be embedded to have forward compatible implementations.
type UnimplementedS2CRegistryServer struct {
}

func (*UnimplementedS2CRegistryServer) Poll(ctx context.Context, req *C2SRequest) (*S2CResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Poll not implemented")
}
func (*UnimplementedS2CRegistryServer) PING(ctx context.Context, req *Timestamp) (*Timestamp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PING not implemented")
}

func RegisterS2CRegistryServer(s *grpc.Server, srv S2CRegistryServer) {
	s.RegisterService(&_S2CRegistry_serviceDesc, srv)
}

func _S2CRegistry_Poll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(C2SRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S2CRegistryServer).Poll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.S2CRegistry/Poll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S2CRegistryServer).Poll(ctx, req.(*C2SRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _S2CRegistry_PING_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Timestamp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S2CRegistryServer).PING(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.S2CRegistry/PING",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S2CRegistryServer).PING(ctx, req.(*Timestamp))
	}
	return interceptor(ctx, in, info, handler)
}

var _S2CRegistry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.S2CRegistry",
	HandlerType: (*S2CRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Poll",
			Handler:    _S2CRegistry_Poll_Handler,
		},
		{
			MethodName: "PING",
			Handler:    _S2CRegistry_PING_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
