// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/services/feed_item_service.proto

package services // import "google.golang.org/genproto/googleapis/ads/googleads/v0/services"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import resources "google.golang.org/genproto/googleapis/ads/googleads/v0/resources"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import field_mask "google.golang.org/genproto/protobuf/field_mask"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Request message for [FeedItemService.GetFeedItem][google.ads.googleads.v0.services.FeedItemService.GetFeedItem].
type GetFeedItemRequest struct {
	// The resource name of the feed item to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFeedItemRequest) Reset()         { *m = GetFeedItemRequest{} }
func (m *GetFeedItemRequest) String() string { return proto.CompactTextString(m) }
func (*GetFeedItemRequest) ProtoMessage()    {}
func (*GetFeedItemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_feed_item_service_b421090e91948349, []int{0}
}
func (m *GetFeedItemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeedItemRequest.Unmarshal(m, b)
}
func (m *GetFeedItemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeedItemRequest.Marshal(b, m, deterministic)
}
func (dst *GetFeedItemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeedItemRequest.Merge(dst, src)
}
func (m *GetFeedItemRequest) XXX_Size() int {
	return xxx_messageInfo_GetFeedItemRequest.Size(m)
}
func (m *GetFeedItemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeedItemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeedItemRequest proto.InternalMessageInfo

func (m *GetFeedItemRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [FeedItemService.MutateFeedItems][google.ads.googleads.v0.services.FeedItemService.MutateFeedItems].
type MutateFeedItemsRequest struct {
	// The ID of the customer whose feed items are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The list of operations to perform on individual feed items.
	Operations           []*FeedItemOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MutateFeedItemsRequest) Reset()         { *m = MutateFeedItemsRequest{} }
func (m *MutateFeedItemsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateFeedItemsRequest) ProtoMessage()    {}
func (*MutateFeedItemsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_feed_item_service_b421090e91948349, []int{1}
}
func (m *MutateFeedItemsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateFeedItemsRequest.Unmarshal(m, b)
}
func (m *MutateFeedItemsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateFeedItemsRequest.Marshal(b, m, deterministic)
}
func (dst *MutateFeedItemsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateFeedItemsRequest.Merge(dst, src)
}
func (m *MutateFeedItemsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateFeedItemsRequest.Size(m)
}
func (m *MutateFeedItemsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateFeedItemsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateFeedItemsRequest proto.InternalMessageInfo

func (m *MutateFeedItemsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateFeedItemsRequest) GetOperations() []*FeedItemOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

// A single operation (create, update, remove) on an feed item.
type FeedItemOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*FeedItemOperation_Create
	//	*FeedItemOperation_Update
	//	*FeedItemOperation_Remove
	Operation            isFeedItemOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *FeedItemOperation) Reset()         { *m = FeedItemOperation{} }
func (m *FeedItemOperation) String() string { return proto.CompactTextString(m) }
func (*FeedItemOperation) ProtoMessage()    {}
func (*FeedItemOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_feed_item_service_b421090e91948349, []int{2}
}
func (m *FeedItemOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedItemOperation.Unmarshal(m, b)
}
func (m *FeedItemOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedItemOperation.Marshal(b, m, deterministic)
}
func (dst *FeedItemOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedItemOperation.Merge(dst, src)
}
func (m *FeedItemOperation) XXX_Size() int {
	return xxx_messageInfo_FeedItemOperation.Size(m)
}
func (m *FeedItemOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedItemOperation.DiscardUnknown(m)
}

var xxx_messageInfo_FeedItemOperation proto.InternalMessageInfo

func (m *FeedItemOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isFeedItemOperation_Operation interface {
	isFeedItemOperation_Operation()
}

type FeedItemOperation_Create struct {
	Create *resources.FeedItem `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type FeedItemOperation_Update struct {
	Update *resources.FeedItem `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type FeedItemOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*FeedItemOperation_Create) isFeedItemOperation_Operation() {}

func (*FeedItemOperation_Update) isFeedItemOperation_Operation() {}

func (*FeedItemOperation_Remove) isFeedItemOperation_Operation() {}

func (m *FeedItemOperation) GetOperation() isFeedItemOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *FeedItemOperation) GetCreate() *resources.FeedItem {
	if x, ok := m.GetOperation().(*FeedItemOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *FeedItemOperation) GetUpdate() *resources.FeedItem {
	if x, ok := m.GetOperation().(*FeedItemOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *FeedItemOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*FeedItemOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*FeedItemOperation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _FeedItemOperation_OneofMarshaler, _FeedItemOperation_OneofUnmarshaler, _FeedItemOperation_OneofSizer, []interface{}{
		(*FeedItemOperation_Create)(nil),
		(*FeedItemOperation_Update)(nil),
		(*FeedItemOperation_Remove)(nil),
	}
}

func _FeedItemOperation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*FeedItemOperation)
	// operation
	switch x := m.Operation.(type) {
	case *FeedItemOperation_Create:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Create); err != nil {
			return err
		}
	case *FeedItemOperation_Update:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Update); err != nil {
			return err
		}
	case *FeedItemOperation_Remove:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Remove)
	case nil:
	default:
		return fmt.Errorf("FeedItemOperation.Operation has unexpected type %T", x)
	}
	return nil
}

func _FeedItemOperation_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*FeedItemOperation)
	switch tag {
	case 1: // operation.create
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(resources.FeedItem)
		err := b.DecodeMessage(msg)
		m.Operation = &FeedItemOperation_Create{msg}
		return true, err
	case 2: // operation.update
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(resources.FeedItem)
		err := b.DecodeMessage(msg)
		m.Operation = &FeedItemOperation_Update{msg}
		return true, err
	case 3: // operation.remove
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Operation = &FeedItemOperation_Remove{x}
		return true, err
	default:
		return false, nil
	}
}

func _FeedItemOperation_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*FeedItemOperation)
	// operation
	switch x := m.Operation.(type) {
	case *FeedItemOperation_Create:
		s := proto.Size(x.Create)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *FeedItemOperation_Update:
		s := proto.Size(x.Update)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *FeedItemOperation_Remove:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Remove)))
		n += len(x.Remove)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Response message for an feed item mutate.
type MutateFeedItemsResponse struct {
	// All results for the mutate.
	Results              []*MutateFeedItemResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *MutateFeedItemsResponse) Reset()         { *m = MutateFeedItemsResponse{} }
func (m *MutateFeedItemsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateFeedItemsResponse) ProtoMessage()    {}
func (*MutateFeedItemsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_feed_item_service_b421090e91948349, []int{3}
}
func (m *MutateFeedItemsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateFeedItemsResponse.Unmarshal(m, b)
}
func (m *MutateFeedItemsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateFeedItemsResponse.Marshal(b, m, deterministic)
}
func (dst *MutateFeedItemsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateFeedItemsResponse.Merge(dst, src)
}
func (m *MutateFeedItemsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateFeedItemsResponse.Size(m)
}
func (m *MutateFeedItemsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateFeedItemsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateFeedItemsResponse proto.InternalMessageInfo

func (m *MutateFeedItemsResponse) GetResults() []*MutateFeedItemResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the feed item mutate.
type MutateFeedItemResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateFeedItemResult) Reset()         { *m = MutateFeedItemResult{} }
func (m *MutateFeedItemResult) String() string { return proto.CompactTextString(m) }
func (*MutateFeedItemResult) ProtoMessage()    {}
func (*MutateFeedItemResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_feed_item_service_b421090e91948349, []int{4}
}
func (m *MutateFeedItemResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateFeedItemResult.Unmarshal(m, b)
}
func (m *MutateFeedItemResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateFeedItemResult.Marshal(b, m, deterministic)
}
func (dst *MutateFeedItemResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateFeedItemResult.Merge(dst, src)
}
func (m *MutateFeedItemResult) XXX_Size() int {
	return xxx_messageInfo_MutateFeedItemResult.Size(m)
}
func (m *MutateFeedItemResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateFeedItemResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateFeedItemResult proto.InternalMessageInfo

func (m *MutateFeedItemResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetFeedItemRequest)(nil), "google.ads.googleads.v0.services.GetFeedItemRequest")
	proto.RegisterType((*MutateFeedItemsRequest)(nil), "google.ads.googleads.v0.services.MutateFeedItemsRequest")
	proto.RegisterType((*FeedItemOperation)(nil), "google.ads.googleads.v0.services.FeedItemOperation")
	proto.RegisterType((*MutateFeedItemsResponse)(nil), "google.ads.googleads.v0.services.MutateFeedItemsResponse")
	proto.RegisterType((*MutateFeedItemResult)(nil), "google.ads.googleads.v0.services.MutateFeedItemResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FeedItemServiceClient is the client API for FeedItemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FeedItemServiceClient interface {
	// Returns the requested feed item in full detail.
	GetFeedItem(ctx context.Context, in *GetFeedItemRequest, opts ...grpc.CallOption) (*resources.FeedItem, error)
	// Creates, updates, or removes feed items. Operation statuses are
	// returned.
	MutateFeedItems(ctx context.Context, in *MutateFeedItemsRequest, opts ...grpc.CallOption) (*MutateFeedItemsResponse, error)
}

type feedItemServiceClient struct {
	cc *grpc.ClientConn
}

func NewFeedItemServiceClient(cc *grpc.ClientConn) FeedItemServiceClient {
	return &feedItemServiceClient{cc}
}

func (c *feedItemServiceClient) GetFeedItem(ctx context.Context, in *GetFeedItemRequest, opts ...grpc.CallOption) (*resources.FeedItem, error) {
	out := new(resources.FeedItem)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.FeedItemService/GetFeedItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedItemServiceClient) MutateFeedItems(ctx context.Context, in *MutateFeedItemsRequest, opts ...grpc.CallOption) (*MutateFeedItemsResponse, error) {
	out := new(MutateFeedItemsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.FeedItemService/MutateFeedItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedItemServiceServer is the server API for FeedItemService service.
type FeedItemServiceServer interface {
	// Returns the requested feed item in full detail.
	GetFeedItem(context.Context, *GetFeedItemRequest) (*resources.FeedItem, error)
	// Creates, updates, or removes feed items. Operation statuses are
	// returned.
	MutateFeedItems(context.Context, *MutateFeedItemsRequest) (*MutateFeedItemsResponse, error)
}

func RegisterFeedItemServiceServer(s *grpc.Server, srv FeedItemServiceServer) {
	s.RegisterService(&_FeedItemService_serviceDesc, srv)
}

func _FeedItemService_GetFeedItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeedItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedItemServiceServer).GetFeedItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.FeedItemService/GetFeedItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedItemServiceServer).GetFeedItem(ctx, req.(*GetFeedItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedItemService_MutateFeedItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateFeedItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedItemServiceServer).MutateFeedItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.FeedItemService/MutateFeedItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedItemServiceServer).MutateFeedItems(ctx, req.(*MutateFeedItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FeedItemService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v0.services.FeedItemService",
	HandlerType: (*FeedItemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFeedItem",
			Handler:    _FeedItemService_GetFeedItem_Handler,
		},
		{
			MethodName: "MutateFeedItems",
			Handler:    _FeedItemService_MutateFeedItems_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v0/services/feed_item_service.proto",
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/services/feed_item_service.proto", fileDescriptor_feed_item_service_b421090e91948349)
}

var fileDescriptor_feed_item_service_b421090e91948349 = []byte{
	// 591 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xc1, 0x6a, 0xdb, 0x4a,
	0x14, 0x7d, 0x92, 0x1f, 0x2e, 0x19, 0xb5, 0x84, 0x0e, 0xa1, 0x15, 0xa6, 0x50, 0xa3, 0x76, 0x61,
	0x1c, 0x3a, 0x72, 0xed, 0x52, 0xe2, 0x98, 0x2c, 0x6c, 0x68, 0x9c, 0x2c, 0xd2, 0x1a, 0x05, 0xb2,
	0x28, 0x06, 0x33, 0xb1, 0xae, 0x85, 0xb0, 0xa5, 0x51, 0x35, 0x23, 0x6f, 0x42, 0x36, 0x85, 0x6e,
	0xbb, 0xe9, 0x1f, 0x14, 0xba, 0xe9, 0x5f, 0x74, 0xdb, 0x6d, 0xfb, 0x09, 0xfd, 0x8d, 0x42, 0x91,
	0x46, 0xa3, 0xd8, 0x71, 0x8d, 0x9b, 0xec, 0xae, 0x46, 0xe7, 0x9c, 0x7b, 0x74, 0xe6, 0x5e, 0xa1,
	0x3d, 0x8f, 0x31, 0x6f, 0x06, 0x36, 0x75, 0xb9, 0x2d, 0xcb, 0xb4, 0x9a, 0x37, 0x6c, 0x0e, 0xf1,
	0xdc, 0x1f, 0x03, 0xb7, 0x27, 0x00, 0xee, 0xc8, 0x17, 0x10, 0x8c, 0xf2, 0x23, 0x12, 0xc5, 0x4c,
	0x30, 0x5c, 0x95, 0x70, 0x42, 0x5d, 0x4e, 0x0a, 0x26, 0x99, 0x37, 0x88, 0x62, 0x56, 0x9e, 0xaf,
	0xd3, 0x8e, 0x81, 0xb3, 0x24, 0x5e, 0x12, 0x97, 0xa2, 0x95, 0x47, 0x8a, 0x12, 0xf9, 0x36, 0x0d,
	0x43, 0x26, 0xa8, 0xf0, 0x59, 0xc8, 0xf3, 0xb7, 0x79, 0x4b, 0x3b, 0x7b, 0x3a, 0x4f, 0x26, 0xf6,
	0xc4, 0x87, 0x99, 0x3b, 0x0a, 0x28, 0x9f, 0x4a, 0x84, 0xd5, 0x46, 0xb8, 0x0f, 0xe2, 0x10, 0xc0,
	0x3d, 0x16, 0x10, 0x38, 0xf0, 0x2e, 0x01, 0x2e, 0xf0, 0x13, 0x74, 0x4f, 0xb5, 0x1c, 0x85, 0x34,
	0x00, 0x53, 0xab, 0x6a, 0xb5, 0x2d, 0xe7, 0xae, 0x3a, 0x7c, 0x4d, 0x03, 0xb0, 0x3e, 0x6a, 0xe8,
	0xc1, 0x49, 0x22, 0xa8, 0x00, 0x45, 0xe7, 0x8a, 0xff, 0x18, 0x19, 0xe3, 0x84, 0x0b, 0x16, 0x40,
	0x3c, 0xf2, 0xdd, 0x9c, 0x8d, 0xd4, 0xd1, 0xb1, 0x8b, 0x4f, 0x11, 0x62, 0x11, 0xc4, 0xd2, 0xac,
	0xa9, 0x57, 0x4b, 0x35, 0xa3, 0xd9, 0x22, 0x9b, 0x02, 0x22, 0xaa, 0xd1, 0x1b, 0xc5, 0x75, 0x16,
	0x64, 0xac, 0x0f, 0x3a, 0xba, 0xbf, 0x82, 0xc0, 0x1d, 0x64, 0x24, 0x91, 0x4b, 0x05, 0x64, 0x9f,
	0x6d, 0xfe, 0x5f, 0xd5, 0x6a, 0x46, 0xb3, 0xa2, 0x7a, 0xa9, 0x64, 0xc8, 0x61, 0x9a, 0xcc, 0x09,
	0xe5, 0x53, 0x07, 0x49, 0x78, 0x5a, 0xe3, 0x57, 0xa8, 0x3c, 0x8e, 0x81, 0x0a, 0x99, 0x80, 0xd1,
	0xdc, 0x5d, 0xeb, 0xb1, 0xb8, 0xa2, 0xc2, 0xe4, 0xd1, 0x7f, 0x4e, 0x4e, 0x4e, 0x65, 0xa4, 0xa8,
	0xa9, 0xdf, 0x4a, 0x46, 0x92, 0xb1, 0x89, 0xca, 0x31, 0x04, 0x6c, 0x0e, 0x66, 0x29, 0x4d, 0x34,
	0x7d, 0x23, 0x9f, 0x7b, 0x06, 0xda, 0x2a, 0x82, 0xb0, 0xa6, 0xe8, 0xe1, 0xca, 0xbd, 0xf0, 0x88,
	0x85, 0x1c, 0xf0, 0x00, 0xdd, 0x89, 0x81, 0x27, 0x33, 0xa1, 0x42, 0x7f, 0xb9, 0x39, 0xf4, 0x65,
	0x2d, 0x27, 0xa3, 0x3b, 0x4a, 0xc6, 0xea, 0xa0, 0x9d, 0xbf, 0x01, 0xfe, 0x69, 0x84, 0x9a, 0xbf,
	0x75, 0xb4, 0xad, 0x78, 0xa7, 0xb2, 0x1f, 0xfe, 0xa2, 0x21, 0x63, 0x61, 0x24, 0xf1, 0x8b, 0xcd,
	0x0e, 0x57, 0x27, 0xb8, 0x72, 0x93, 0x84, 0xad, 0xd6, 0xfb, 0x1f, 0xbf, 0x3e, 0xe9, 0xcf, 0xf0,
	0x6e, 0xba, 0x6b, 0x17, 0x4b, 0xb6, 0x0f, 0xd4, 0xd0, 0x72, 0xbb, 0x9e, 0x2d, 0x5f, 0x96, 0xa7,
	0x5d, 0xbf, 0xc4, 0xdf, 0x34, 0xb4, 0x7d, 0x2d, 0x66, 0xbc, 0x77, 0xd3, 0x34, 0xd5, 0xc6, 0x54,
	0xda, 0xb7, 0x60, 0xca, 0x3b, 0xb5, 0xda, 0x99, 0xfb, 0x96, 0x45, 0x52, 0xf7, 0x57, 0x76, 0x2f,
	0x16, 0x36, 0xf0, 0xa0, 0x7e, 0x79, 0x65, 0x7e, 0x3f, 0xc8, 0x84, 0xf6, 0xb5, 0x7a, 0xef, 0xa7,
	0x86, 0x9e, 0x8e, 0x59, 0xb0, 0xb1, 0x77, 0x6f, 0xe7, 0xda, 0x2d, 0x0d, 0xd2, 0xb5, 0x19, 0x68,
	0x6f, 0x8f, 0x72, 0xa6, 0xc7, 0x66, 0x34, 0xf4, 0x08, 0x8b, 0x3d, 0xdb, 0x83, 0x30, 0x5b, 0x2a,
	0xf5, 0x07, 0x8b, 0x7c, 0xbe, 0xfe, 0x67, 0xd9, 0x51, 0xc5, 0x67, 0xbd, 0xd4, 0xef, 0x76, 0xbf,
	0xea, 0xd5, 0xbe, 0x14, 0xec, 0xba, 0x9c, 0xc8, 0x32, 0xad, 0xce, 0x1a, 0x24, 0x6f, 0xcc, 0xbf,
	0x2b, 0xc8, 0xb0, 0xeb, 0xf2, 0x61, 0x01, 0x19, 0x9e, 0x35, 0x86, 0x0a, 0x72, 0x5e, 0xce, 0x0c,
	0xb4, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0xcc, 0x47, 0xea, 0xc7, 0xac, 0x05, 0x00, 0x00,
}
