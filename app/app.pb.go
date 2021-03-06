// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app.proto

package app

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

// Next Tag: 5
type AppState struct {
	// nonce should be unique for each app session among the same signers
	Nonce uint64 `protobuf:"varint,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// for each nonce, new state has higher sequence number
	SeqNum uint64 `protobuf:"varint,2,opt,name=seq_num,json=seqNum,proto3" json:"seq_num,omitempty"`
	// app specific state
	State []byte `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	// on-chain response (settle, action) timeout
	Timeout              uint64   `protobuf:"varint,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppState) Reset()         { *m = AppState{} }
func (m *AppState) String() string { return proto.CompactTextString(m) }
func (*AppState) ProtoMessage()    {}
func (*AppState) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f9056a14b86d47, []int{0}
}

func (m *AppState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppState.Unmarshal(m, b)
}
func (m *AppState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppState.Marshal(b, m, deterministic)
}
func (m *AppState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppState.Merge(m, src)
}
func (m *AppState) XXX_Size() int {
	return xxx_messageInfo_AppState.Size(m)
}
func (m *AppState) XXX_DiscardUnknown() {
	xxx_messageInfo_AppState.DiscardUnknown(m)
}

var xxx_messageInfo_AppState proto.InternalMessageInfo

func (m *AppState) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *AppState) GetSeqNum() uint64 {
	if m != nil {
		return m.SeqNum
	}
	return 0
}

func (m *AppState) GetState() []byte {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *AppState) GetTimeout() uint64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

// Next Tag: 3
type StateProof struct {
	// serialized AppState
	AppState             []byte   `protobuf:"bytes,1,opt,name=app_state,json=appState,proto3" json:"app_state,omitempty"`
	Sigs                 [][]byte `protobuf:"bytes,2,rep,name=sigs,proto3" json:"sigs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StateProof) Reset()         { *m = StateProof{} }
func (m *StateProof) String() string { return proto.CompactTextString(m) }
func (*StateProof) ProtoMessage()    {}
func (*StateProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f9056a14b86d47, []int{1}
}

func (m *StateProof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StateProof.Unmarshal(m, b)
}
func (m *StateProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StateProof.Marshal(b, m, deterministic)
}
func (m *StateProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateProof.Merge(m, src)
}
func (m *StateProof) XXX_Size() int {
	return xxx_messageInfo_StateProof.Size(m)
}
func (m *StateProof) XXX_DiscardUnknown() {
	xxx_messageInfo_StateProof.DiscardUnknown(m)
}

var xxx_messageInfo_StateProof proto.InternalMessageInfo

func (m *StateProof) GetAppState() []byte {
	if m != nil {
		return m.AppState
	}
	return nil
}

func (m *StateProof) GetSigs() [][]byte {
	if m != nil {
		return m.Sigs
	}
	return nil
}

// used for multi-session app
// Next Tag: 3
type SessionQuery struct {
	// session ID
	Session []byte `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
	// query related to the specified session
	Query                []byte   `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionQuery) Reset()         { *m = SessionQuery{} }
func (m *SessionQuery) String() string { return proto.CompactTextString(m) }
func (*SessionQuery) ProtoMessage()    {}
func (*SessionQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f9056a14b86d47, []int{2}
}

func (m *SessionQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionQuery.Unmarshal(m, b)
}
func (m *SessionQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionQuery.Marshal(b, m, deterministic)
}
func (m *SessionQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionQuery.Merge(m, src)
}
func (m *SessionQuery) XXX_Size() int {
	return xxx_messageInfo_SessionQuery.Size(m)
}
func (m *SessionQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionQuery.DiscardUnknown(m)
}

var xxx_messageInfo_SessionQuery proto.InternalMessageInfo

func (m *SessionQuery) GetSession() []byte {
	if m != nil {
		return m.Session
	}
	return nil
}

func (m *SessionQuery) GetQuery() []byte {
	if m != nil {
		return m.Query
	}
	return nil
}

var E_Soltype = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         1003,
	Name:          "app.soltype",
	Tag:           "bytes,1003,opt,name=soltype",
	Filename:      "app.proto",
}

func init() {
	proto.RegisterType((*AppState)(nil), "app.AppState")
	proto.RegisterType((*StateProof)(nil), "app.StateProof")
	proto.RegisterType((*SessionQuery)(nil), "app.SessionQuery")
	proto.RegisterExtension(E_Soltype)
}

func init() { proto.RegisterFile("app.proto", fileDescriptor_e0f9056a14b86d47) }

var fileDescriptor_e0f9056a14b86d47 = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xbd, 0x6a, 0xfb, 0x30,
	0x14, 0xc5, 0x71, 0xe2, 0xc4, 0x89, 0xf0, 0x24, 0xfe, 0x83, 0xf9, 0x97, 0x80, 0x1b, 0x42, 0xf1,
	0x52, 0x1b, 0x9a, 0xa5, 0x14, 0x1a, 0x68, 0x0a, 0x1d, 0x3a, 0xf4, 0xc3, 0xd9, 0xba, 0x04, 0xdb,
	0xb9, 0x71, 0x45, 0x6d, 0x5d, 0x45, 0x1f, 0x94, 0xec, 0x7d, 0xca, 0xbc, 0x42, 0x5f, 0xa2, 0x58,
	0xc2, 0x43, 0x36, 0x9d, 0xa3, 0x73, 0x0f, 0xbf, 0x7b, 0xc9, 0xb4, 0x10, 0x22, 0x15, 0x12, 0x35,
	0xd2, 0x61, 0x21, 0xc4, 0xff, 0xb8, 0x46, 0xac, 0x1b, 0xc8, 0xac, 0x55, 0x9a, 0x7d, 0xb6, 0x03,
	0x55, 0x49, 0x26, 0x34, 0x4a, 0x17, 0x9b, 0xff, 0x78, 0x64, 0xf2, 0x20, 0xc4, 0x46, 0x17, 0x1a,
	0xe8, 0x8c, 0x8c, 0x38, 0xf2, 0x0a, 0x22, 0x2f, 0xf6, 0x12, 0x7f, 0x1d, 0x9c, 0x56, 0xbe, 0x61,
	0x5c, 0xe7, 0xce, 0xa5, 0x31, 0x09, 0x14, 0x1c, 0xb6, 0xdc, 0xb4, 0xd1, 0xe0, 0x3c, 0x30, 0x56,
	0x70, 0x78, 0x31, 0x2d, 0xfd, 0x47, 0x46, 0xaa, 0x6b, 0x8a, 0x86, 0xb1, 0x97, 0x84, 0xb9, 0x13,
	0xf4, 0x92, 0x04, 0x9a, 0xb5, 0x80, 0x46, 0x47, 0xfe, 0xf9, 0x5c, 0xef, 0xcf, 0xef, 0x09, 0xb1,
	0x08, 0x6f, 0x12, 0x71, 0x4f, 0x2f, 0xec, 0x22, 0x5b, 0x57, 0xe5, 0xd9, 0xaa, 0x49, 0xd1, 0x43,
	0x52, 0xe2, 0x2b, 0x56, 0xab, 0x68, 0x10, 0x0f, 0x93, 0x30, 0xb7, 0xef, 0xf9, 0x33, 0x09, 0x37,
	0xa0, 0x14, 0x43, 0xfe, 0x6e, 0x40, 0x1e, 0xe9, 0xa2, 0x23, 0xb5, 0xda, 0x8d, 0xaf, 0xc9, 0x69,
	0x15, 0x94, 0x47, 0x0d, 0x6a, 0x79, 0x93, 0xf7, 0x5f, 0x1d, 0xed, 0xa1, 0x8b, 0xdb, 0x6d, 0xc2,
	0xdc, 0x89, 0xbb, 0x5b, 0x12, 0x28, 0x6c, 0xf4, 0x51, 0x00, 0x9d, 0xa5, 0xee, 0x7e, 0x69, 0x7f,
	0xbf, 0xf4, 0x89, 0x41, 0xb3, 0x7b, 0x15, 0x9a, 0x21, 0x57, 0xd1, 0x6f, 0x10, 0x7b, 0xc9, 0x34,
	0xef, 0xe3, 0xeb, 0xab, 0x8f, 0x45, 0xcd, 0xf4, 0xa7, 0x29, 0xd3, 0x0a, 0xdb, 0xac, 0x82, 0x06,
	0xe4, 0x35, 0x07, 0xfd, 0x8d, 0xf2, 0x2b, 0xab, 0xf1, 0xb1, 0xd3, 0x59, 0x21, 0x44, 0x39, 0xb6,
	0x75, 0xcb, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xab, 0x9c, 0x67, 0xba, 0xae, 0x01, 0x00, 0x00,
}
