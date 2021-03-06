// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package user

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type User struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	PhotoUrl             string   `protobuf:"bytes,4,opt,name=photoUrl,proto3" json:"photoUrl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPhotoUrl() string {
	if m != nil {
		return m.PhotoUrl
	}
	return ""
}

type GetUserRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (m *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(m, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetUserResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserResponse) Reset()         { *m = GetUserResponse{} }
func (m *GetUserResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserResponse) ProtoMessage()    {}
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *GetUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserResponse.Unmarshal(m, b)
}
func (m *GetUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserResponse.Marshal(b, m, deterministic)
}
func (m *GetUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserResponse.Merge(m, src)
}
func (m *GetUserResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserResponse.Size(m)
}
func (m *GetUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserResponse proto.InternalMessageInfo

func (m *GetUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type GetUsersRequest struct {
	Limit                int32    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               int32    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Id                   int32    `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUsersRequest) Reset()         { *m = GetUsersRequest{} }
func (m *GetUsersRequest) String() string { return proto.CompactTextString(m) }
func (*GetUsersRequest) ProtoMessage()    {}
func (*GetUsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *GetUsersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUsersRequest.Unmarshal(m, b)
}
func (m *GetUsersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUsersRequest.Marshal(b, m, deterministic)
}
func (m *GetUsersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUsersRequest.Merge(m, src)
}
func (m *GetUsersRequest) XXX_Size() int {
	return xxx_messageInfo_GetUsersRequest.Size(m)
}
func (m *GetUsersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUsersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUsersRequest proto.InternalMessageInfo

func (m *GetUsersRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *GetUsersRequest) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *GetUsersRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetUsersRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetUsersResponse struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUsersResponse) Reset()         { *m = GetUsersResponse{} }
func (m *GetUsersResponse) String() string { return proto.CompactTextString(m) }
func (*GetUsersResponse) ProtoMessage()    {}
func (*GetUsersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *GetUsersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUsersResponse.Unmarshal(m, b)
}
func (m *GetUsersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUsersResponse.Marshal(b, m, deterministic)
}
func (m *GetUsersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUsersResponse.Merge(m, src)
}
func (m *GetUsersResponse) XXX_Size() int {
	return xxx_messageInfo_GetUsersResponse.Size(m)
}
func (m *GetUsersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUsersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUsersResponse proto.InternalMessageInfo

func (m *GetUsersResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type LoginRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{6}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *LoginResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type CreateUserRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	PhotoUrl             string   `protobuf:"bytes,3,opt,name=photoUrl,proto3" json:"photoUrl,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{7}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateUserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CreateUserRequest) GetPhotoUrl() string {
	if m != nil {
		return m.PhotoUrl
	}
	return ""
}

func (m *CreateUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type CreateUserResponse struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserResponse) Reset()         { *m = CreateUserResponse{} }
func (m *CreateUserResponse) String() string { return proto.CompactTextString(m) }
func (*CreateUserResponse) ProtoMessage()    {}
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{8}
}

func (m *CreateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResponse.Unmarshal(m, b)
}
func (m *CreateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResponse.Marshal(b, m, deterministic)
}
func (m *CreateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResponse.Merge(m, src)
}
func (m *CreateUserResponse) XXX_Size() int {
	return xxx_messageInfo_CreateUserResponse.Size(m)
}
func (m *CreateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResponse proto.InternalMessageInfo

func (m *CreateUserResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CreateUserResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type UpdateUserRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	PhotoUrl             string   `protobuf:"bytes,4,opt,name=photo_url,json=photoUrl,proto3" json:"photo_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserRequest) Reset()         { *m = UpdateUserRequest{} }
func (m *UpdateUserRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUserRequest) ProtoMessage()    {}
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{9}
}

func (m *UpdateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserRequest.Unmarshal(m, b)
}
func (m *UpdateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserRequest.Marshal(b, m, deterministic)
}
func (m *UpdateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserRequest.Merge(m, src)
}
func (m *UpdateUserRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateUserRequest.Size(m)
}
func (m *UpdateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserRequest proto.InternalMessageInfo

func (m *UpdateUserRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateUserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateUserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UpdateUserRequest) GetPhotoUrl() string {
	if m != nil {
		return m.PhotoUrl
	}
	return ""
}

type UpdateUserResponse struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserResponse) Reset()         { *m = UpdateUserResponse{} }
func (m *UpdateUserResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateUserResponse) ProtoMessage()    {}
func (*UpdateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{10}
}

func (m *UpdateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserResponse.Unmarshal(m, b)
}
func (m *UpdateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserResponse.Marshal(b, m, deterministic)
}
func (m *UpdateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserResponse.Merge(m, src)
}
func (m *UpdateUserResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateUserResponse.Size(m)
}
func (m *UpdateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserResponse proto.InternalMessageInfo

func (m *UpdateUserResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*User)(nil), "user.User")
	proto.RegisterType((*GetUserRequest)(nil), "user.GetUserRequest")
	proto.RegisterType((*GetUserResponse)(nil), "user.GetUserResponse")
	proto.RegisterType((*GetUsersRequest)(nil), "user.GetUsersRequest")
	proto.RegisterType((*GetUsersResponse)(nil), "user.GetUsersResponse")
	proto.RegisterType((*LoginRequest)(nil), "user.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "user.LoginResponse")
	proto.RegisterType((*CreateUserRequest)(nil), "user.CreateUserRequest")
	proto.RegisterType((*CreateUserResponse)(nil), "user.CreateUserResponse")
	proto.RegisterType((*UpdateUserRequest)(nil), "user.UpdateUserRequest")
	proto.RegisterType((*UpdateUserResponse)(nil), "user.UpdateUserResponse")
}

func init() {
	proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf)
}

var fileDescriptor_116e343673f7ffaf = []byte{
	// 507 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x94, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xe5, 0xd8, 0x2e, 0xed, 0x04, 0x42, 0x33, 0x0d, 0xc1, 0x18, 0x84, 0xa2, 0x15, 0x07,
	0xc4, 0xa1, 0x56, 0x0b, 0x5c, 0x7a, 0x42, 0x42, 0x82, 0x4b, 0xc5, 0x21, 0x28, 0x17, 0x84, 0x84,
	0x4c, 0xb3, 0x0d, 0x06, 0xc7, 0x6b, 0xec, 0x4d, 0x39, 0x20, 0x2e, 0xbc, 0x02, 0xef, 0xc0, 0x0b,
	0xf1, 0x0a, 0x3c, 0x08, 0xeb, 0xd9, 0xb5, 0x63, 0xaf, 0xc3, 0x01, 0x6e, 0x9e, 0xd9, 0xd9, 0x6f,
	0xfe, 0x9d, 0xf9, 0x65, 0x80, 0x4d, 0xc9, 0x8b, 0xe3, 0xbc, 0x10, 0x52, 0xa0, 0x57, 0x7d, 0x87,
	0xf7, 0x56, 0x42, 0xac, 0x52, 0x1e, 0xc5, 0x79, 0x12, 0xc5, 0x59, 0x26, 0x64, 0x2c, 0x13, 0x91,
	0x95, 0xba, 0x86, 0xbd, 0x05, 0x6f, 0xa1, 0xaa, 0x70, 0x04, 0x83, 0x64, 0x19, 0x38, 0x33, 0xe7,
	0xa1, 0x3f, 0x57, 0x5f, 0x88, 0xe0, 0x65, 0xf1, 0x9a, 0x07, 0x03, 0x95, 0x39, 0x98, 0xd3, 0x37,
	0x4e, 0xc0, 0xe7, 0xeb, 0x38, 0x49, 0x03, 0x97, 0x92, 0x3a, 0xc0, 0x10, 0xf6, 0xf3, 0x0f, 0x0a,
	0xb5, 0x28, 0xd2, 0xc0, 0xa3, 0x83, 0x26, 0x66, 0x33, 0x18, 0xbd, 0xe4, 0xb2, 0x6a, 0x30, 0xe7,
	0x9f, 0x37, 0xbc, 0x94, 0x76, 0x1f, 0x76, 0x02, 0x37, 0x9b, 0x8a, 0x32, 0x57, 0xba, 0x38, 0xde,
	0x07, 0x12, 0x4e, 0x45, 0xc3, 0x53, 0x38, 0xa6, 0x17, 0x51, 0x05, 0xe5, 0xd9, 0x45, 0x73, 0xa5,
	0xac, 0xa9, 0x4a, 0x59, 0x9a, 0xac, 0x13, 0x69, 0xc0, 0x3a, 0xc0, 0x29, 0xec, 0x89, 0xcb, 0xcb,
	0x92, 0x4b, 0x7a, 0x85, 0x3f, 0x37, 0x91, 0xd1, 0xe0, 0xf6, 0xde, 0xea, 0x6d, 0xdf, 0xca, 0x9e,
	0xc0, 0xe1, 0xb6, 0x89, 0x11, 0x36, 0x03, 0xbf, 0x12, 0x50, 0xaa, 0x2e, 0xae, 0xa5, 0x4c, 0x1f,
	0xb0, 0x67, 0x70, 0xfd, 0x5c, 0xac, 0x92, 0xac, 0xa5, 0x4b, 0x4f, 0xcc, 0xb1, 0x27, 0x16, 0x97,
	0xe5, 0x17, 0x51, 0x2c, 0xcd, 0x7c, 0x9b, 0x98, 0x3d, 0x85, 0x1b, 0x86, 0x60, 0x9a, 0xda, 0x8b,
	0x51, 0x48, 0x29, 0x3e, 0xf1, 0xcc, 0xdc, 0xd4, 0x01, 0xdb, 0xc0, 0xf8, 0x79, 0xc1, 0x63, 0xc9,
	0xdb, 0xb3, 0xae, 0xdf, 0xe5, 0xec, 0xda, 0xe1, 0xe0, 0x6f, 0x3b, 0x74, 0xbb, 0x3b, 0xec, 0xa8,
	0xf5, 0x2c, 0xb5, 0x67, 0x80, 0xed, 0xb6, 0xff, 0x24, 0xf9, 0x23, 0x8c, 0x17, 0xf9, 0xd2, 0x92,
	0xfc, 0xff, 0x36, 0xbc, 0x0b, 0x07, 0x24, 0xf9, 0xdd, 0x66, 0x87, 0x0f, 0x1f, 0x00, 0xb6, 0x7b,
	0xed, 0xd6, 0x79, 0xfa, 0xd3, 0x85, 0x61, 0x55, 0xf0, 0x9a, 0x17, 0x57, 0xc9, 0x05, 0xc7, 0x57,
	0x70, 0xcd, 0x78, 0x00, 0x27, 0x7a, 0xd7, 0x5d, 0x33, 0x87, 0xb7, 0xac, 0xac, 0xe6, 0xb2, 0xe9,
	0xf7, 0x5f, 0xbf, 0x7f, 0x0c, 0x0e, 0x71, 0x14, 0x5d, 0x9d, 0x44, 0x64, 0x8c, 0xe8, 0x6b, 0xb2,
	0xfc, 0x86, 0xe7, 0xb0, 0x5f, 0x7b, 0x0a, 0xbb, 0x57, 0x6b, 0x23, 0x87, 0x53, 0x3b, 0x6d, 0x90,
	0x63, 0x42, 0x0e, 0xf1, 0xa0, 0x41, 0xe2, 0x0b, 0xf0, 0xc9, 0x29, 0x88, 0xfa, 0x4e, 0xdb, 0x78,
	0xe1, 0x51, 0x27, 0x67, 0x20, 0x13, 0x82, 0x8c, 0x18, 0x41, 0xd2, 0xea, 0xe8, 0xcc, 0x79, 0x84,
	0x0b, 0x80, 0xed, 0x0e, 0xf1, 0xb6, 0xbe, 0xd8, 0x33, 0x53, 0x18, 0xf4, 0x0f, 0x76, 0x61, 0x49,
	0x5b, 0x85, 0x7d, 0x03, 0xb0, 0x1d, 0x79, 0x8d, 0xed, 0x2d, 0xbc, 0xc6, 0xf6, 0xb7, 0xc3, 0xee,
	0x10, 0xf6, 0x28, 0xb4, 0xa6, 0xa8, 0xd8, 0xef, 0xf7, 0xe8, 0xdf, 0xf5, 0xf8, 0x4f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x75, 0x39, 0xc0, 0xc9, 0xed, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error) {
	out := new(GetUsersResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/GetUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) GetUser(ctx context.Context, req *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (*UnimplementedUserServiceServer) GetUsers(ctx context.Context, req *GetUsersRequest) (*GetUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (*UnimplementedUserServiceServer) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedUserServiceServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedUserServiceServer) UpdateUser(ctx context.Context, req *UpdateUserRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUsers(ctx, req.(*GetUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _UserService_GetUsers_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UserService_Login_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserService_UpdateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
