// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/openbank/openbank/v1/profile/all.proto

package profile

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	account "github.com/openbank/openbank/v1/account"
	types "github.com/openbank/openbank/v1/types"
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

// GetProfileRequest is the request message for retrieving customer profile.
type GetProfileRequest struct {
	ProfileID            string   `protobuf:"bytes,1,opt,name=ProfileID,json=profile_id,proto3" json:"ProfileID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProfileRequest) Reset()         { *m = GetProfileRequest{} }
func (m *GetProfileRequest) String() string { return proto.CompactTextString(m) }
func (*GetProfileRequest) ProtoMessage()    {}
func (*GetProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d361c7b766adec27, []int{0}
}

func (m *GetProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProfileRequest.Unmarshal(m, b)
}
func (m *GetProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProfileRequest.Marshal(b, m, deterministic)
}
func (m *GetProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProfileRequest.Merge(m, src)
}
func (m *GetProfileRequest) XXX_Size() int {
	return xxx_messageInfo_GetProfileRequest.Size(m)
}
func (m *GetProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetProfileRequest proto.InternalMessageInfo

func (m *GetProfileRequest) GetProfileID() string {
	if m != nil {
		return m.ProfileID
	}
	return ""
}

// GetProfileResponse is the response message for retrieving customer profile.
type GetProfileResponse struct {
	Profile              *types.Profile     `protobuf:"bytes,1,opt,name=Profile,json=profile,proto3" json:"Profile,omitempty"`
	Accounts             []*account.Account `protobuf:"bytes,2,rep,name=Accounts,json=accounts,proto3" json:"Accounts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetProfileResponse) Reset()         { *m = GetProfileResponse{} }
func (m *GetProfileResponse) String() string { return proto.CompactTextString(m) }
func (*GetProfileResponse) ProtoMessage()    {}
func (*GetProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d361c7b766adec27, []int{1}
}

func (m *GetProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProfileResponse.Unmarshal(m, b)
}
func (m *GetProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProfileResponse.Marshal(b, m, deterministic)
}
func (m *GetProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProfileResponse.Merge(m, src)
}
func (m *GetProfileResponse) XXX_Size() int {
	return xxx_messageInfo_GetProfileResponse.Size(m)
}
func (m *GetProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetProfileResponse proto.InternalMessageInfo

func (m *GetProfileResponse) GetProfile() *types.Profile {
	if m != nil {
		return m.Profile
	}
	return nil
}

func (m *GetProfileResponse) GetAccounts() []*account.Account {
	if m != nil {
		return m.Accounts
	}
	return nil
}

// GetProfileCardsRequest is the request message for retrieving customer cards.
type GetProfileCardsRequest struct {
	ProfileID            string   `protobuf:"bytes,1,opt,name=ProfileID,json=profile_id,proto3" json:"ProfileID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProfileCardsRequest) Reset()         { *m = GetProfileCardsRequest{} }
func (m *GetProfileCardsRequest) String() string { return proto.CompactTextString(m) }
func (*GetProfileCardsRequest) ProtoMessage()    {}
func (*GetProfileCardsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d361c7b766adec27, []int{2}
}

func (m *GetProfileCardsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProfileCardsRequest.Unmarshal(m, b)
}
func (m *GetProfileCardsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProfileCardsRequest.Marshal(b, m, deterministic)
}
func (m *GetProfileCardsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProfileCardsRequest.Merge(m, src)
}
func (m *GetProfileCardsRequest) XXX_Size() int {
	return xxx_messageInfo_GetProfileCardsRequest.Size(m)
}
func (m *GetProfileCardsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProfileCardsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetProfileCardsRequest proto.InternalMessageInfo

func (m *GetProfileCardsRequest) GetProfileID() string {
	if m != nil {
		return m.ProfileID
	}
	return ""
}

// GetProfileCardsResponse is the response message for retrieving customer cards.
type GetProfileCardsResponse struct {
	Cards                []*ProfileCard `protobuf:"bytes,1,rep,name=Cards,json=cards,proto3" json:"Cards,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetProfileCardsResponse) Reset()         { *m = GetProfileCardsResponse{} }
func (m *GetProfileCardsResponse) String() string { return proto.CompactTextString(m) }
func (*GetProfileCardsResponse) ProtoMessage()    {}
func (*GetProfileCardsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d361c7b766adec27, []int{3}
}

func (m *GetProfileCardsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProfileCardsResponse.Unmarshal(m, b)
}
func (m *GetProfileCardsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProfileCardsResponse.Marshal(b, m, deterministic)
}
func (m *GetProfileCardsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProfileCardsResponse.Merge(m, src)
}
func (m *GetProfileCardsResponse) XXX_Size() int {
	return xxx_messageInfo_GetProfileCardsResponse.Size(m)
}
func (m *GetProfileCardsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProfileCardsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetProfileCardsResponse proto.InternalMessageInfo

func (m *GetProfileCardsResponse) GetCards() []*ProfileCard {
	if m != nil {
		return m.Cards
	}
	return nil
}

// ProfileCard holds details about a customer's credit card.
type ProfileCard struct {
	// CardToken is the card number.
	CardToken string `protobuf:"bytes,1,opt,name=CardToken,json=card_token,proto3" json:"CardToken,omitempty"`
	// Category is the card type.
	Category string `protobuf:"bytes,2,opt,name=Category,json=category,proto3" json:"Category,omitempty"`
	// LastFour is the last four digits of the card.
	LastFour string `protobuf:"bytes,3,opt,name=LastFour,json=last_four,proto3" json:"LastFour,omitempty"`
	// OwnerName is the name of the card's owner.
	OwnerName string `protobuf:"bytes,4,opt,name=OwnerName,json=owner_name,proto3" json:"OwnerName,omitempty"`
	// Type is the type of the account associated with the card.
	Type                 string   `protobuf:"bytes,5,opt,name=Type,json=type,proto3" json:"Type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProfileCard) Reset()         { *m = ProfileCard{} }
func (m *ProfileCard) String() string { return proto.CompactTextString(m) }
func (*ProfileCard) ProtoMessage()    {}
func (*ProfileCard) Descriptor() ([]byte, []int) {
	return fileDescriptor_d361c7b766adec27, []int{4}
}

func (m *ProfileCard) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProfileCard.Unmarshal(m, b)
}
func (m *ProfileCard) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProfileCard.Marshal(b, m, deterministic)
}
func (m *ProfileCard) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProfileCard.Merge(m, src)
}
func (m *ProfileCard) XXX_Size() int {
	return xxx_messageInfo_ProfileCard.Size(m)
}
func (m *ProfileCard) XXX_DiscardUnknown() {
	xxx_messageInfo_ProfileCard.DiscardUnknown(m)
}

var xxx_messageInfo_ProfileCard proto.InternalMessageInfo

func (m *ProfileCard) GetCardToken() string {
	if m != nil {
		return m.CardToken
	}
	return ""
}

func (m *ProfileCard) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *ProfileCard) GetLastFour() string {
	if m != nil {
		return m.LastFour
	}
	return ""
}

func (m *ProfileCard) GetOwnerName() string {
	if m != nil {
		return m.OwnerName
	}
	return ""
}

func (m *ProfileCard) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func init() {
	proto.RegisterType((*GetProfileRequest)(nil), "profile.GetProfileRequest")
	proto.RegisterType((*GetProfileResponse)(nil), "profile.GetProfileResponse")
	proto.RegisterType((*GetProfileCardsRequest)(nil), "profile.GetProfileCardsRequest")
	proto.RegisterType((*GetProfileCardsResponse)(nil), "profile.GetProfileCardsResponse")
	proto.RegisterType((*ProfileCard)(nil), "profile.ProfileCard")
}

func init() {
	proto.RegisterFile("github.com/openbank/openbank/v1/profile/all.proto", fileDescriptor_d361c7b766adec27)
}

var fileDescriptor_d361c7b766adec27 = []byte{
	// 1119 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0x5b, 0x6f, 0x1b, 0xc5,
	0x17, 0xdf, 0x5d, 0x27, 0xae, 0x33, 0x51, 0xab, 0x76, 0xfe, 0x7f, 0x81, 0x65, 0x50, 0x19, 0x82,
	0x44, 0xd3, 0xdb, 0xf8, 0xd2, 0x14, 0x41, 0x24, 0x84, 0x9c, 0x94, 0x84, 0x84, 0x96, 0x1a, 0x13,
	0x2a, 0xda, 0x97, 0x68, 0xbc, 0x7b, 0x6c, 0x0f, 0x59, 0xcf, 0x6c, 0x67, 0x66, 0xe3, 0x9a, 0x27,
	0xe0, 0x29, 0x2f, 0x48, 0x51, 0x78, 0xac, 0xc4, 0xe7, 0x28, 0x42, 0x7c, 0x07, 0x24, 0x5e, 0x78,
	0x40, 0xe2, 0x05, 0x89, 0xaf, 0xc0, 0x13, 0x42, 0xb3, 0x17, 0x5f, 0xe4, 0x94, 0x88, 0xf2, 0x64,
	0xeb, 0x5c, 0x7e, 0xe7, 0x77, 0x7e, 0xe7, 0x9c, 0xdd, 0x45, 0xf5, 0x1e, 0x37, 0xfd, 0xb8, 0x43,
	0x7d, 0x39, 0xa8, 0xca, 0x08, 0x44, 0x87, 0x89, 0x83, 0xc9, 0x9f, 0xc3, 0x7a, 0x35, 0x52, 0xb2,
	0xcb, 0x43, 0xa8, 0xb2, 0x30, 0xa4, 0x91, 0x92, 0x46, 0xe2, 0x73, 0x99, 0xa9, 0xf2, 0x6a, 0x4f,
	0xca, 0x9e, 0x75, 0x45, 0xbc, 0xca, 0x84, 0x90, 0x86, 0x19, 0x2e, 0x85, 0x4e, 0xc3, 0x2a, 0x37,
	0x92, 0x1f, 0xff, 0x66, 0x0f, 0xc4, 0x4d, 0x3d, 0x64, 0xbd, 0x1e, 0xa8, 0xaa, 0x8c, 0x92, 0x88,
	0x53, 0xa2, 0xcf, 0xe4, 0xc1, 0x7c, 0x5f, 0xc6, 0xc2, 0x4c, 0x78, 0x54, 0xaa, 0x67, 0xa5, 0x98,
	0x51, 0x04, 0x7a, 0x92, 0xb0, 0xb2, 0x85, 0x2e, 0x6d, 0x83, 0x69, 0xa5, 0xec, 0xdb, 0xf0, 0x38,
	0x06, 0x6d, 0xf0, 0x55, 0xb4, 0x94, 0x59, 0x76, 0xee, 0x94, 0x5d, 0xe2, 0xae, 0x2e, 0x6d, 0xa0,
	0x92, 0x53, 0x76, 0x56, 0x9d, 0x9a, 0xd3, 0x72, 0xda, 0x28, 0x6b, 0x76, 0x9f, 0x07, 0xeb, 0xc5,
	0x92, 0x73, 0xd1, 0x29, 0x3b, 0x2b, 0xdf, 0xb8, 0x08, 0x4f, 0x03, 0xe9, 0x48, 0x0a, 0x0d, 0x78,
	0x0d, 0x9d, 0xcb, 0x4c, 0x09, 0xce, 0x72, 0xe3, 0x02, 0x4d, 0x18, 0xd0, 0xcc, 0x3a, 0x83, 0x9b,
	0x8b, 0x88, 0xdf, 0x41, 0xa5, 0x66, 0xda, 0x9a, 0x2e, 0x7b, 0xa4, 0xb0, 0xba, 0xdc, 0xb8, 0x44,
	0xb3, 0x5e, 0x35, 0xcd, 0x3c, 0x33, 0x99, 0xa5, 0xdc, 0x3b, 0xe6, 0xf3, 0x21, 0x7a, 0x69, 0x42,
	0x67, 0x93, 0xa9, 0x40, 0xff, 0x87, 0xe6, 0x1e, 0xa2, 0x97, 0xe7, 0xc0, 0xb2, 0x06, 0xdf, 0x42,
	0x8b, 0x89, 0xa1, 0xec, 0x26, 0x3c, 0xff, 0x4f, 0xb3, 0x74, 0x3a, 0x15, 0x3d, 0x83, 0xbf, 0xe8,
	0xdb, 0xf0, 0x31, 0xf4, 0xaf, 0x2e, 0x5a, 0x9e, 0x0a, 0xb5, 0xec, 0xec, 0xef, 0x9e, 0x3c, 0x00,
	0x71, 0x1a, 0x3b, 0x9b, 0xbd, 0x6f, 0xac, 0x17, 0xbf, 0x89, 0x4a, 0x9b, 0xcc, 0x40, 0x4f, 0xaa,
	0x51, 0xd9, 0x9b, 0x8b, 0x2c, 0xf9, 0x99, 0x0f, 0x5f, 0x41, 0xa5, 0xbb, 0x4c, 0x9b, 0x2d, 0x19,
	0xab, 0x72, 0x61, 0x2e, 0x6e, 0x29, 0x64, 0xda, 0xec, 0x77, 0x65, 0xac, 0x6c, 0xed, 0xfb, 0x43,
	0x01, 0xea, 0x23, 0x36, 0x80, 0xf2, 0xc2, 0x7c, 0x6d, 0x69, 0x9d, 0xfb, 0x82, 0x0d, 0x00, 0x5f,
	0x46, 0x0b, 0x7b, 0xa3, 0x08, 0xca, 0x8b, 0x73, 0x51, 0x0b, 0x76, 0xbe, 0x79, 0x7b, 0x8d, 0x1f,
	0x8a, 0xe8, 0x42, 0xd6, 0xde, 0x27, 0xa0, 0x0e, 0xb9, 0x0f, 0xf8, 0x0f, 0x0f, 0xa1, 0x89, 0x9a,
	0xb8, 0x32, 0x56, 0x6c, 0x6e, 0x0f, 0x2b, 0xaf, 0x9c, 0xea, 0x4b, 0x95, 0x5f, 0x79, 0xea, 0x9d,
	0x34, 0xff, 0x72, 0xc7, 0xfb, 0x85, 0xcf, 0x7f, 0x1c, 0x83, 0x1a, 0x91, 0xfc, 0x10, 0xaf, 0xb7,
	0xc1, 0xc4, 0x4a, 0x68, 0x62, 0xfa, 0x90, 0x1b, 0x09, 0x13, 0x01, 0x61, 0x5a, 0x4b, 0x9f, 0x33,
	0x03, 0x01, 0xc9, 0xb7, 0x66, 0x77, 0x0f, 0x15, 0x1a, 0xb5, 0x1a, 0xbe, 0x87, 0x2e, 0x67, 0xa5,
	0x09, 0x3c, 0x01, 0x3f, 0xb6, 0x31, 0x3a, 0xf6, 0x7d, 0xd0, 0xba, 0x1b, 0x87, 0xe1, 0x88, 0xe2,
	0xeb, 0xe8, 0x6a, 0xe5, 0xca, 0x1b, 0xd5, 0x00, 0xba, 0x5c, 0xf0, 0xf4, 0x6e, 0x33, 0xec, 0x79,
	0x86, 0xbb, 0x75, 0x54, 0x58, 0xab, 0xad, 0xe1, 0x6b, 0x68, 0x35, 0xe5, 0x02, 0x01, 0x19, 0xf6,
	0x41, 0x24, 0x8c, 0x14, 0x68, 0x19, 0x2b, 0x1f, 0x08, 0xd7, 0x44, 0x48, 0x43, 0xba, 0x32, 0x16,
	0x01, 0xed, 0x50, 0x74, 0x03, 0x15, 0xef, 0x37, 0x63, 0xd3, 0x6f, 0xe0, 0x15, 0x44, 0xfa, 0xc6,
	0x44, 0x7a, 0xbd, 0x5a, 0x65, 0xb1, 0xe9, 0xd3, 0x8e, 0x38, 0xa0, 0x46, 0xe6, 0xf5, 0xa8, 0x02,
	0x16, 0x7c, 0xfd, 0xf3, 0xef, 0xdf, 0x7a, 0xe7, 0xf1, 0xf2, 0xd4, 0x43, 0xe9, 0xc8, 0x73, 0x8e,
	0xbd, 0x64, 0x18, 0xf8, 0xfb, 0x02, 0xba, 0x30, 0xbb, 0xb7, 0xf8, 0xb5, 0x53, 0x14, 0x9d, 0xbe,
	0x8e, 0x0a, 0x79, 0x7e, 0x40, 0xa6, 0xfb, 0x6f, 0xde, 0x49, 0xf3, 0x47, 0x6f, 0xa2, 0xfb, 0x72,
	0xaa, 0x7b, 0xb2, 0xd6, 0x95, 0x5e, 0xae, 0x3a, 0x13, 0x84, 0x29, 0xc5, 0x46, 0x44, 0x76, 0xc9,
	0x14, 0xcc, 0xb4, 0xfa, 0x43, 0x6e, 0xfa, 0x89, 0x16, 0xd9, 0x18, 0x48, 0x57, 0xaa, 0x99, 0x69,
	0x75, 0x98, 0x86, 0x80, 0x48, 0x31, 0x36, 0xf0, 0x00, 0x84, 0xe1, 0x5d, 0x0e, 0xea, 0x05, 0xb4,
	0xdd, 0x7d, 0x98, 0x0e, 0xb9, 0x7d, 0xe6, 0x90, 0x6b, 0x88, 0x56, 0x6e, 0x9c, 0x31, 0xe4, 0x19,
	0x4d, 0x5e, 0x68, 0x6c, 0x18, 0x5f, 0x9c, 0x7e, 0x97, 0x58, 0xfd, 0x26, 0xb3, 0xab, 0x14, 0x8e,
	0x3c, 0x67, 0xe3, 0x59, 0xf1, 0xa4, 0x79, 0x52, 0x44, 0x85, 0x06, 0xad, 0xe1, 0xed, 0xf1, 0x73,
	0x82, 0x34, 0x5b, 0x3b, 0xb8, 0xde, 0x52, 0xf2, 0x90, 0x07, 0xa0, 0xc9, 0x66, 0xfb, 0xd3, 0x3b,
	0x44, 0x46, 0xa0, 0xd2, 0x77, 0x87, 0xd5, 0xcb, 0x8a, 0x90, 0x07, 0xe7, 0x62, 0xd0, 0xc6, 0x62,
	0x9d, 0xd6, 0x68, 0xed, 0x9a, 0xeb, 0x35, 0x2e, 0xb2, 0x28, 0x0a, 0xb9, 0x9f, 0x24, 0x54, 0x3f,
	0xd7, 0x52, 0xac, 0xcf, 0x59, 0xda, 0xfb, 0x56, 0xe0, 0x1a, 0xfe, 0x0c, 0x3d, 0x38, 0x4d, 0xe0,
	0x54, 0xbf, 0x8e, 0x0c, 0x46, 0x56, 0xe4, 0x01, 0x0b, 0xbb, 0x52, 0x0d, 0x98, 0xb1, 0x5a, 0x4a,
	0x45, 0x02, 0x09, 0xa9, 0xf2, 0x03, 0x66, 0xfc, 0x74, 0xc6, 0xf0, 0x24, 0x02, 0xdf, 0xba, 0xb3,
	0x5c, 0xda, 0xbe, 0x6b, 0x0b, 0xd4, 0xf1, 0xfb, 0x68, 0xf3, 0xf9, 0x05, 0xc6, 0x40, 0xbe, 0x14,
	0x86, 0xf1, 0xec, 0x9a, 0x63, 0x0d, 0xea, 0x8a, 0x26, 0xbe, 0x82, 0x64, 0x1f, 0x58, 0xa8, 0x69,
	0xbb, 0x65, 0xd1, 0x6e, 0xe1, 0x1d, 0xb4, 0x3d, 0x8f, 0x66, 0xe3, 0x27, 0x50, 0x7d, 0x76, 0x08,
	0x24, 0x02, 0x35, 0xe0, 0x5a, 0x73, 0x2b, 0x96, 0xb4, 0x0b, 0x08, 0x5a, 0xcf, 0xec, 0x0e, 0x6d,
	0xff, 0xfb, 0x0d, 0x6b, 0x6f, 0xa1, 0xc2, 0xed, 0x5a, 0x0d, 0xbf, 0x87, 0xde, 0x9d, 0x4d, 0x61,
	0x82, 0xc4, 0x62, 0xac, 0x00, 0x28, 0x25, 0x15, 0x91, 0xbe, 0x1f, 0xab, 0x74, 0xbd, 0x2d, 0xa2,
	0x06, 0x75, 0x08, 0x8a, 0x68, 0x1e, 0x00, 0x7d, 0xf4, 0x9d, 0x87, 0x9e, 0x7a, 0xe3, 0x85, 0x3a,
	0xf6, 0x4a, 0x05, 0xfc, 0x95, 0xdb, 0xcc, 0x48, 0xca, 0xf1, 0x39, 0xe4, 0x2c, 0xb4, 0xa5, 0xa1,
	0x40, 0x1b, 0xc5, 0x93, 0x02, 0xb6, 0xa3, 0xd8, 0xf4, 0xad, 0x36, 0x7e, 0x72, 0x69, 0x56, 0x00,
	0x4d, 0xc9, 0x9e, 0xbd, 0xb5, 0x89, 0x83, 0xa7, 0x97, 0x95, 0xa0, 0x76, 0x65, 0x18, 0xca, 0x61,
	0x2a, 0x81, 0xad, 0x2a, 0x15, 0xff, 0x22, 0x8d, 0xd8, 0x94, 0x01, 0x90, 0x6e, 0x28, 0x87, 0x74,
	0x75, 0xa1, 0x51, 0x4a, 0xbe, 0x3a, 0x62, 0xd3, 0x5f, 0x5f, 0x4a, 0x3e, 0x26, 0xec, 0x7b, 0x68,
	0xe3, 0x31, 0xba, 0x7d, 0xf6, 0xca, 0xe3, 0x4b, 0x0f, 0x38, 0x0c, 0xc7, 0xd4, 0x03, 0x66, 0x18,
	0x7a, 0x1b, 0xbd, 0xfe, 0x4f, 0x69, 0x43, 0xc5, 0x0d, 0xe0, 0xff, 0xdd, 0x63, 0x82, 0xf5, 0x60,
	0x26, 0xf3, 0x03, 0xb7, 0xe5, 0x3c, 0xca, 0xbf, 0x16, 0xbe, 0x74, 0x9d, 0x23, 0xd7, 0x39, 0x76,
	0x9d, 0x67, 0xae, 0xf3, 0x8b, 0xeb, 0xfc, 0xe9, 0x3a, 0x3f, 0x79, 0x4e, 0xa7, 0x98, 0x7c, 0xde,
	0xdc, 0xfa, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xc6, 0xcc, 0xf4, 0x44, 0xcc, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProfileServiceClient is the client API for ProfileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProfileServiceClient interface {
	// GetProfile retrieves user profile.
	GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*GetProfileResponse, error)
	// GetProfileCard retrieves cards.
	GetProfileCard(ctx context.Context, in *GetProfileCardsRequest, opts ...grpc.CallOption) (*GetProfileCardsResponse, error)
}

type profileServiceClient struct {
	cc *grpc.ClientConn
}

func NewProfileServiceClient(cc *grpc.ClientConn) ProfileServiceClient {
	return &profileServiceClient{cc}
}

func (c *profileServiceClient) GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*GetProfileResponse, error) {
	out := new(GetProfileResponse)
	err := c.cc.Invoke(ctx, "/profile.ProfileService/GetProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) GetProfileCard(ctx context.Context, in *GetProfileCardsRequest, opts ...grpc.CallOption) (*GetProfileCardsResponse, error) {
	out := new(GetProfileCardsResponse)
	err := c.cc.Invoke(ctx, "/profile.ProfileService/GetProfileCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfileServiceServer is the server API for ProfileService service.
type ProfileServiceServer interface {
	// GetProfile retrieves user profile.
	GetProfile(context.Context, *GetProfileRequest) (*GetProfileResponse, error)
	// GetProfileCard retrieves cards.
	GetProfileCard(context.Context, *GetProfileCardsRequest) (*GetProfileCardsResponse, error)
}

// UnimplementedProfileServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProfileServiceServer struct {
}

func (*UnimplementedProfileServiceServer) GetProfile(ctx context.Context, req *GetProfileRequest) (*GetProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
func (*UnimplementedProfileServiceServer) GetProfileCard(ctx context.Context, req *GetProfileCardsRequest) (*GetProfileCardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfileCard not implemented")
}

func RegisterProfileServiceServer(s *grpc.Server, srv ProfileServiceServer) {
	s.RegisterService(&_ProfileService_serviceDesc, srv)
}

func _ProfileService_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.ProfileService/GetProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).GetProfile(ctx, req.(*GetProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_GetProfileCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileCardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).GetProfileCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.ProfileService/GetProfileCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).GetProfileCard(ctx, req.(*GetProfileCardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProfileService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "profile.ProfileService",
	HandlerType: (*ProfileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProfile",
			Handler:    _ProfileService_GetProfile_Handler,
		},
		{
			MethodName: "GetProfileCard",
			Handler:    _ProfileService_GetProfileCard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/openbank/openbank/v1/profile/all.proto",
}
