// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/services/customer_service.proto

package services // import "google.golang.org/genproto/googleapis/ads/googleads/v0/services"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/wrappers"
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

// Request message for [CustomerService.GetCustomer][google.ads.googleads.v0.services.CustomerService.GetCustomer].
type GetCustomerRequest struct {
	// The resource name of the customer to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCustomerRequest) Reset()         { *m = GetCustomerRequest{} }
func (m *GetCustomerRequest) String() string { return proto.CompactTextString(m) }
func (*GetCustomerRequest) ProtoMessage()    {}
func (*GetCustomerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_customer_service_90b708bb54f49834, []int{0}
}
func (m *GetCustomerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCustomerRequest.Unmarshal(m, b)
}
func (m *GetCustomerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCustomerRequest.Marshal(b, m, deterministic)
}
func (dst *GetCustomerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCustomerRequest.Merge(dst, src)
}
func (m *GetCustomerRequest) XXX_Size() int {
	return xxx_messageInfo_GetCustomerRequest.Size(m)
}
func (m *GetCustomerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCustomerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCustomerRequest proto.InternalMessageInfo

func (m *GetCustomerRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [CustomerService.MutateCustomer][google.ads.googleads.v0.services.CustomerService.MutateCustomer].
type MutateCustomerRequest struct {
	// The ID of the customer being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The operation to perform on the customer
	Operation            *CustomerOperation `protobuf:"bytes,4,opt,name=operation,proto3" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *MutateCustomerRequest) Reset()         { *m = MutateCustomerRequest{} }
func (m *MutateCustomerRequest) String() string { return proto.CompactTextString(m) }
func (*MutateCustomerRequest) ProtoMessage()    {}
func (*MutateCustomerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_customer_service_90b708bb54f49834, []int{1}
}
func (m *MutateCustomerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCustomerRequest.Unmarshal(m, b)
}
func (m *MutateCustomerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCustomerRequest.Marshal(b, m, deterministic)
}
func (dst *MutateCustomerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCustomerRequest.Merge(dst, src)
}
func (m *MutateCustomerRequest) XXX_Size() int {
	return xxx_messageInfo_MutateCustomerRequest.Size(m)
}
func (m *MutateCustomerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCustomerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCustomerRequest proto.InternalMessageInfo

func (m *MutateCustomerRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateCustomerRequest) GetOperation() *CustomerOperation {
	if m != nil {
		return m.Operation
	}
	return nil
}

// Request message for [CustomerService.CreateCustomerClient][google.ads.googleads.v0.services.CustomerService.CreateCustomerClient].
type CreateCustomerClientRequest struct {
	// The ID of the Manager under whom client customer is being created.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The new client customer to create. The resource name on this customer
	// will be ignored.
	CustomerClient       *resources.Customer `protobuf:"bytes,2,opt,name=customer_client,json=customerClient,proto3" json:"customer_client,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *CreateCustomerClientRequest) Reset()         { *m = CreateCustomerClientRequest{} }
func (m *CreateCustomerClientRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCustomerClientRequest) ProtoMessage()    {}
func (*CreateCustomerClientRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_customer_service_90b708bb54f49834, []int{2}
}
func (m *CreateCustomerClientRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCustomerClientRequest.Unmarshal(m, b)
}
func (m *CreateCustomerClientRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCustomerClientRequest.Marshal(b, m, deterministic)
}
func (dst *CreateCustomerClientRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCustomerClientRequest.Merge(dst, src)
}
func (m *CreateCustomerClientRequest) XXX_Size() int {
	return xxx_messageInfo_CreateCustomerClientRequest.Size(m)
}
func (m *CreateCustomerClientRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCustomerClientRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCustomerClientRequest proto.InternalMessageInfo

func (m *CreateCustomerClientRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *CreateCustomerClientRequest) GetCustomerClient() *resources.Customer {
	if m != nil {
		return m.CustomerClient
	}
	return nil
}

// A single update on a customer.
type CustomerOperation struct {
	// Mutate operation. Only updates are supported for customer.
	Update *resources.Customer `protobuf:"bytes,1,opt,name=update,proto3" json:"update,omitempty"`
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask           *field_mask.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CustomerOperation) Reset()         { *m = CustomerOperation{} }
func (m *CustomerOperation) String() string { return proto.CompactTextString(m) }
func (*CustomerOperation) ProtoMessage()    {}
func (*CustomerOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_customer_service_90b708bb54f49834, []int{3}
}
func (m *CustomerOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerOperation.Unmarshal(m, b)
}
func (m *CustomerOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerOperation.Marshal(b, m, deterministic)
}
func (dst *CustomerOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerOperation.Merge(dst, src)
}
func (m *CustomerOperation) XXX_Size() int {
	return xxx_messageInfo_CustomerOperation.Size(m)
}
func (m *CustomerOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerOperation.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerOperation proto.InternalMessageInfo

func (m *CustomerOperation) GetUpdate() *resources.Customer {
	if m != nil {
		return m.Update
	}
	return nil
}

func (m *CustomerOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

// Response message for CreateCustomerClient mutate.
type CreateCustomerClientResponse struct {
	// The resource name of the newly created customer client.
	ResourceName         string   `protobuf:"bytes,2,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCustomerClientResponse) Reset()         { *m = CreateCustomerClientResponse{} }
func (m *CreateCustomerClientResponse) String() string { return proto.CompactTextString(m) }
func (*CreateCustomerClientResponse) ProtoMessage()    {}
func (*CreateCustomerClientResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_customer_service_90b708bb54f49834, []int{4}
}
func (m *CreateCustomerClientResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCustomerClientResponse.Unmarshal(m, b)
}
func (m *CreateCustomerClientResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCustomerClientResponse.Marshal(b, m, deterministic)
}
func (dst *CreateCustomerClientResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCustomerClientResponse.Merge(dst, src)
}
func (m *CreateCustomerClientResponse) XXX_Size() int {
	return xxx_messageInfo_CreateCustomerClientResponse.Size(m)
}
func (m *CreateCustomerClientResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCustomerClientResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCustomerClientResponse proto.InternalMessageInfo

func (m *CreateCustomerClientResponse) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Response message for customer mutate.
type MutateCustomerResponse struct {
	// Result for the mutate.
	Result               *MutateCustomerResult `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MutateCustomerResponse) Reset()         { *m = MutateCustomerResponse{} }
func (m *MutateCustomerResponse) String() string { return proto.CompactTextString(m) }
func (*MutateCustomerResponse) ProtoMessage()    {}
func (*MutateCustomerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_customer_service_90b708bb54f49834, []int{5}
}
func (m *MutateCustomerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCustomerResponse.Unmarshal(m, b)
}
func (m *MutateCustomerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCustomerResponse.Marshal(b, m, deterministic)
}
func (dst *MutateCustomerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCustomerResponse.Merge(dst, src)
}
func (m *MutateCustomerResponse) XXX_Size() int {
	return xxx_messageInfo_MutateCustomerResponse.Size(m)
}
func (m *MutateCustomerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCustomerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCustomerResponse proto.InternalMessageInfo

func (m *MutateCustomerResponse) GetResult() *MutateCustomerResult {
	if m != nil {
		return m.Result
	}
	return nil
}

// The result for the customer mutate.
type MutateCustomerResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateCustomerResult) Reset()         { *m = MutateCustomerResult{} }
func (m *MutateCustomerResult) String() string { return proto.CompactTextString(m) }
func (*MutateCustomerResult) ProtoMessage()    {}
func (*MutateCustomerResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_customer_service_90b708bb54f49834, []int{6}
}
func (m *MutateCustomerResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCustomerResult.Unmarshal(m, b)
}
func (m *MutateCustomerResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCustomerResult.Marshal(b, m, deterministic)
}
func (dst *MutateCustomerResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCustomerResult.Merge(dst, src)
}
func (m *MutateCustomerResult) XXX_Size() int {
	return xxx_messageInfo_MutateCustomerResult.Size(m)
}
func (m *MutateCustomerResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCustomerResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCustomerResult proto.InternalMessageInfo

func (m *MutateCustomerResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [CustomerService.ListAccessibleCustomers][google.ads.googleads.v0.services.CustomerService.ListAccessibleCustomers].
type ListAccessibleCustomersRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAccessibleCustomersRequest) Reset()         { *m = ListAccessibleCustomersRequest{} }
func (m *ListAccessibleCustomersRequest) String() string { return proto.CompactTextString(m) }
func (*ListAccessibleCustomersRequest) ProtoMessage()    {}
func (*ListAccessibleCustomersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_customer_service_90b708bb54f49834, []int{7}
}
func (m *ListAccessibleCustomersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAccessibleCustomersRequest.Unmarshal(m, b)
}
func (m *ListAccessibleCustomersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAccessibleCustomersRequest.Marshal(b, m, deterministic)
}
func (dst *ListAccessibleCustomersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAccessibleCustomersRequest.Merge(dst, src)
}
func (m *ListAccessibleCustomersRequest) XXX_Size() int {
	return xxx_messageInfo_ListAccessibleCustomersRequest.Size(m)
}
func (m *ListAccessibleCustomersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAccessibleCustomersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListAccessibleCustomersRequest proto.InternalMessageInfo

// Response message for [CustomerService.ListAccessibleCustomers][google.ads.googleads.v0.services.CustomerService.ListAccessibleCustomers].
type ListAccessibleCustomersResponse struct {
	// Resource name of customers directly accessible by the
	// user authenticating the call.
	ResourceNames        []string `protobuf:"bytes,1,rep,name=resource_names,json=resourceNames,proto3" json:"resource_names,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAccessibleCustomersResponse) Reset()         { *m = ListAccessibleCustomersResponse{} }
func (m *ListAccessibleCustomersResponse) String() string { return proto.CompactTextString(m) }
func (*ListAccessibleCustomersResponse) ProtoMessage()    {}
func (*ListAccessibleCustomersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_customer_service_90b708bb54f49834, []int{8}
}
func (m *ListAccessibleCustomersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAccessibleCustomersResponse.Unmarshal(m, b)
}
func (m *ListAccessibleCustomersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAccessibleCustomersResponse.Marshal(b, m, deterministic)
}
func (dst *ListAccessibleCustomersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAccessibleCustomersResponse.Merge(dst, src)
}
func (m *ListAccessibleCustomersResponse) XXX_Size() int {
	return xxx_messageInfo_ListAccessibleCustomersResponse.Size(m)
}
func (m *ListAccessibleCustomersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAccessibleCustomersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListAccessibleCustomersResponse proto.InternalMessageInfo

func (m *ListAccessibleCustomersResponse) GetResourceNames() []string {
	if m != nil {
		return m.ResourceNames
	}
	return nil
}

func init() {
	proto.RegisterType((*GetCustomerRequest)(nil), "google.ads.googleads.v0.services.GetCustomerRequest")
	proto.RegisterType((*MutateCustomerRequest)(nil), "google.ads.googleads.v0.services.MutateCustomerRequest")
	proto.RegisterType((*CreateCustomerClientRequest)(nil), "google.ads.googleads.v0.services.CreateCustomerClientRequest")
	proto.RegisterType((*CustomerOperation)(nil), "google.ads.googleads.v0.services.CustomerOperation")
	proto.RegisterType((*CreateCustomerClientResponse)(nil), "google.ads.googleads.v0.services.CreateCustomerClientResponse")
	proto.RegisterType((*MutateCustomerResponse)(nil), "google.ads.googleads.v0.services.MutateCustomerResponse")
	proto.RegisterType((*MutateCustomerResult)(nil), "google.ads.googleads.v0.services.MutateCustomerResult")
	proto.RegisterType((*ListAccessibleCustomersRequest)(nil), "google.ads.googleads.v0.services.ListAccessibleCustomersRequest")
	proto.RegisterType((*ListAccessibleCustomersResponse)(nil), "google.ads.googleads.v0.services.ListAccessibleCustomersResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CustomerServiceClient is the client API for CustomerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomerServiceClient interface {
	// Returns the requested customer in full detail.
	GetCustomer(ctx context.Context, in *GetCustomerRequest, opts ...grpc.CallOption) (*resources.Customer, error)
	// Updates a customer. Operation statuses are returned.
	MutateCustomer(ctx context.Context, in *MutateCustomerRequest, opts ...grpc.CallOption) (*MutateCustomerResponse, error)
	// Returns resource names of customers directly accessible by the
	// user authenticating the call.
	ListAccessibleCustomers(ctx context.Context, in *ListAccessibleCustomersRequest, opts ...grpc.CallOption) (*ListAccessibleCustomersResponse, error)
	// Creates a new client under manager. The new client customer is returned.
	CreateCustomerClient(ctx context.Context, in *CreateCustomerClientRequest, opts ...grpc.CallOption) (*CreateCustomerClientResponse, error)
}

type customerServiceClient struct {
	cc *grpc.ClientConn
}

func NewCustomerServiceClient(cc *grpc.ClientConn) CustomerServiceClient {
	return &customerServiceClient{cc}
}

func (c *customerServiceClient) GetCustomer(ctx context.Context, in *GetCustomerRequest, opts ...grpc.CallOption) (*resources.Customer, error) {
	out := new(resources.Customer)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.CustomerService/GetCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) MutateCustomer(ctx context.Context, in *MutateCustomerRequest, opts ...grpc.CallOption) (*MutateCustomerResponse, error) {
	out := new(MutateCustomerResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.CustomerService/MutateCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) ListAccessibleCustomers(ctx context.Context, in *ListAccessibleCustomersRequest, opts ...grpc.CallOption) (*ListAccessibleCustomersResponse, error) {
	out := new(ListAccessibleCustomersResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.CustomerService/ListAccessibleCustomers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) CreateCustomerClient(ctx context.Context, in *CreateCustomerClientRequest, opts ...grpc.CallOption) (*CreateCustomerClientResponse, error) {
	out := new(CreateCustomerClientResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.CustomerService/CreateCustomerClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerServiceServer is the server API for CustomerService service.
type CustomerServiceServer interface {
	// Returns the requested customer in full detail.
	GetCustomer(context.Context, *GetCustomerRequest) (*resources.Customer, error)
	// Updates a customer. Operation statuses are returned.
	MutateCustomer(context.Context, *MutateCustomerRequest) (*MutateCustomerResponse, error)
	// Returns resource names of customers directly accessible by the
	// user authenticating the call.
	ListAccessibleCustomers(context.Context, *ListAccessibleCustomersRequest) (*ListAccessibleCustomersResponse, error)
	// Creates a new client under manager. The new client customer is returned.
	CreateCustomerClient(context.Context, *CreateCustomerClientRequest) (*CreateCustomerClientResponse, error)
}

func RegisterCustomerServiceServer(s *grpc.Server, srv CustomerServiceServer) {
	s.RegisterService(&_CustomerService_serviceDesc, srv)
}

func _CustomerService_GetCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).GetCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.CustomerService/GetCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).GetCustomer(ctx, req.(*GetCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_MutateCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).MutateCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.CustomerService/MutateCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).MutateCustomer(ctx, req.(*MutateCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_ListAccessibleCustomers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAccessibleCustomersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).ListAccessibleCustomers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.CustomerService/ListAccessibleCustomers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).ListAccessibleCustomers(ctx, req.(*ListAccessibleCustomersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_CreateCustomerClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCustomerClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).CreateCustomerClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.CustomerService/CreateCustomerClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).CreateCustomerClient(ctx, req.(*CreateCustomerClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v0.services.CustomerService",
	HandlerType: (*CustomerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCustomer",
			Handler:    _CustomerService_GetCustomer_Handler,
		},
		{
			MethodName: "MutateCustomer",
			Handler:    _CustomerService_MutateCustomer_Handler,
		},
		{
			MethodName: "ListAccessibleCustomers",
			Handler:    _CustomerService_ListAccessibleCustomers_Handler,
		},
		{
			MethodName: "CreateCustomerClient",
			Handler:    _CustomerService_CreateCustomerClient_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v0/services/customer_service.proto",
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/services/customer_service.proto", fileDescriptor_customer_service_90b708bb54f49834)
}

var fileDescriptor_customer_service_90b708bb54f49834 = []byte{
	// 697 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x4f, 0x4f, 0xd4, 0x4e,
	0x18, 0xce, 0x2c, 0xbf, 0x6c, 0xc2, 0xbb, 0x3f, 0x20, 0x4e, 0x50, 0x37, 0x2b, 0x81, 0xb5, 0x42,
	0xc0, 0x35, 0xb6, 0x1b, 0x30, 0xa2, 0x4b, 0xd6, 0xb8, 0x6c, 0x22, 0x98, 0x08, 0x62, 0x35, 0x1c,
	0xcc, 0x26, 0xa4, 0xb4, 0xc3, 0xda, 0xd0, 0x76, 0x6a, 0xa7, 0xc5, 0x03, 0xe1, 0xe2, 0x95, 0xa3,
	0x7a, 0xf0, 0xea, 0xd1, 0xb3, 0x9f, 0xc2, 0xc4, 0x93, 0x7c, 0x04, 0x3f, 0x88, 0xe9, 0x74, 0xa6,
	0xec, 0xbf, 0xb2, 0x2c, 0xde, 0xde, 0x9d, 0x7d, 0x9f, 0xe7, 0x7d, 0xde, 0xbf, 0x85, 0xd5, 0x36,
	0xa5, 0x6d, 0x87, 0x68, 0x86, 0xc5, 0xb4, 0xc4, 0x8c, 0xad, 0xa3, 0xaa, 0xc6, 0x48, 0x70, 0x64,
	0x9b, 0x84, 0x69, 0x66, 0xc4, 0x42, 0xea, 0x92, 0x60, 0x4f, 0xbc, 0xa8, 0x7e, 0x40, 0x43, 0x8a,
	0xcb, 0x89, 0xb7, 0x6a, 0x58, 0x4c, 0x4d, 0x81, 0xea, 0x51, 0x55, 0x95, 0xc0, 0x52, 0x35, 0x8b,
	0x3a, 0x20, 0x8c, 0x46, 0x41, 0x27, 0x77, 0xc2, 0x59, 0x9a, 0x91, 0x08, 0xdf, 0xd6, 0x0c, 0xcf,
	0xa3, 0xa1, 0x11, 0xda, 0xd4, 0x63, 0xe2, 0x5f, 0x11, 0x51, 0xe3, 0xbf, 0xf6, 0xa3, 0x03, 0xed,
	0xc0, 0x26, 0x8e, 0xb5, 0xe7, 0x1a, 0xec, 0x50, 0x78, 0xcc, 0xf6, 0x7a, 0x7c, 0x08, 0x0c, 0xdf,
	0x27, 0x81, 0x60, 0x50, 0x1e, 0x03, 0xde, 0x20, 0x61, 0x53, 0x04, 0xd5, 0xc9, 0xfb, 0x88, 0xb0,
	0x10, 0xdf, 0x81, 0x09, 0xa9, 0x68, 0xcf, 0x33, 0x5c, 0x52, 0x44, 0x65, 0xb4, 0x34, 0xae, 0xff,
	0x2f, 0x1f, 0xb7, 0x0d, 0x97, 0x28, 0xa7, 0x08, 0xae, 0x6f, 0x45, 0xa1, 0x11, 0x92, 0x5e, 0xf8,
	0x1c, 0x14, 0xd2, 0x12, 0xd9, 0x96, 0x00, 0x83, 0x7c, 0x7a, 0x6e, 0xe1, 0x57, 0x30, 0x4e, 0x7d,
	0x12, 0xf0, 0x5c, 0x8a, 0xff, 0x95, 0xd1, 0x52, 0x61, 0x79, 0x45, 0x1d, 0x56, 0x3d, 0x55, 0x86,
	0x79, 0x29, 0xa1, 0xfa, 0x39, 0x8b, 0xf2, 0x19, 0xc1, 0xad, 0x66, 0x40, 0x3a, 0xd4, 0x34, 0x1d,
	0x9b, 0x78, 0xe1, 0xa5, 0x35, 0xbd, 0x81, 0xa9, 0xd4, 0xc1, 0xe4, 0xd0, 0x62, 0x8e, 0x2b, 0xbb,
	0x97, 0xa9, 0x2c, 0xed, 0x5a, 0x2a, 0x4d, 0x9f, 0x34, 0xbb, 0xa2, 0x2b, 0x5f, 0x10, 0x5c, 0xeb,
	0xd3, 0x8d, 0x9b, 0x90, 0x8f, 0x7c, 0xcb, 0x08, 0x93, 0xc2, 0x8e, 0x18, 0x42, 0x40, 0xf1, 0x1a,
	0x14, 0x12, 0x8b, 0xf7, 0x5b, 0x88, 0x2d, 0x49, 0x26, 0xd9, 0x70, 0xf5, 0x59, 0x3c, 0x12, 0x5b,
	0x06, 0x3b, 0xd4, 0x21, 0x71, 0x8f, 0x6d, 0xa5, 0x09, 0x33, 0x83, 0xab, 0xc5, 0x7c, 0xea, 0x31,
	0xd2, 0x3f, 0x01, 0xb9, 0x01, 0x13, 0xf0, 0x0e, 0x6e, 0xf4, 0x0e, 0x80, 0x80, 0x6f, 0x43, 0x3e,
	0x20, 0x2c, 0x72, 0x64, 0x0d, 0x1f, 0x0e, 0xef, 0x6e, 0x1f, 0x53, 0xe4, 0x84, 0xba, 0x60, 0x51,
	0xd6, 0x60, 0x7a, 0xd0, 0xff, 0x97, 0x1b, 0xd4, 0x32, 0xcc, 0xbe, 0xb0, 0x59, 0xd8, 0x30, 0x4d,
	0xc2, 0x98, 0xbd, 0xef, 0xa4, 0x24, 0x4c, 0x0c, 0x87, 0xb2, 0x09, 0x73, 0x99, 0x1e, 0x22, 0xa3,
	0x05, 0x98, 0xec, 0x8a, 0xc4, 0x8a, 0xa8, 0x3c, 0xb6, 0x34, 0xae, 0x4f, 0x74, 0x86, 0x62, 0xcb,
	0xa7, 0x79, 0x98, 0x92, 0xe0, 0xd7, 0x49, 0x6a, 0xf8, 0x2b, 0x82, 0x42, 0xc7, 0x92, 0xe1, 0x07,
	0xc3, 0x8b, 0xd1, 0xbf, 0x93, 0xa5, 0x51, 0x66, 0x44, 0x59, 0xfc, 0xf8, 0xfb, 0xcf, 0xa7, 0xdc,
	0x6d, 0x3c, 0x17, 0x1f, 0x97, 0xe3, 0x2e, 0xe1, 0x75, 0x39, 0xa1, 0x4c, 0xab, 0x9c, 0xe0, 0x1f,
	0x08, 0x26, 0xbb, 0x2b, 0x8b, 0x57, 0x47, 0xef, 0x55, 0xa2, 0xf0, 0xd1, 0x15, 0x9a, 0xcc, 0x8b,
	0xab, 0x68, 0x5c, 0xee, 0x5d, 0x65, 0x3e, 0x96, 0x7b, 0xae, 0xef, 0xb8, 0x63, 0x63, 0xeb, 0x95,
	0x93, 0x9a, 0xcb, 0xd1, 0x35, 0x54, 0xc1, 0xbf, 0x10, 0xdc, 0xcc, 0xe8, 0x18, 0x7e, 0x3a, 0x5c,
	0xc6, 0xc5, 0xe3, 0x50, 0x6a, 0xfc, 0x03, 0x83, 0xc8, 0xe8, 0x3e, 0xcf, 0x68, 0x11, 0x2f, 0x74,
	0x65, 0x54, 0x73, 0x32, 0x34, 0x9f, 0x21, 0x98, 0x1e, 0xb4, 0x8f, 0xb8, 0x7e, 0x89, 0xb3, 0x98,
	0x7d, 0xf5, 0x4a, 0x4f, 0xae, 0x0a, 0x17, 0x69, 0xd4, 0x79, 0x1a, 0xab, 0xca, 0xf2, 0xc5, 0x8d,
	0x31, 0x07, 0x70, 0xd4, 0x50, 0x65, 0xfd, 0x0c, 0xc1, 0xbc, 0x49, 0xdd, 0xa1, 0x22, 0xd6, 0xa7,
	0x7b, 0x76, 0x66, 0x27, 0xbe, 0x5e, 0x3b, 0xe8, 0xed, 0xa6, 0x40, 0xb6, 0xa9, 0x63, 0x78, 0x6d,
	0x95, 0x06, 0x6d, 0xad, 0x4d, 0x3c, 0x7e, 0xdb, 0xe4, 0x07, 0xd4, 0xb7, 0x59, 0xf6, 0xa7, 0x7a,
	0x4d, 0x1a, 0xdf, 0x72, 0x63, 0x1b, 0x8d, 0xc6, 0xf7, 0x5c, 0x79, 0x23, 0x21, 0x6c, 0x58, 0x4c,
	0x4d, 0xcc, 0xd8, 0xda, 0xad, 0xaa, 0x22, 0x30, 0xfb, 0x29, 0x5d, 0x5a, 0x0d, 0x8b, 0xb5, 0x52,
	0x97, 0xd6, 0x6e, 0xb5, 0x25, 0x5d, 0xf6, 0xf3, 0x5c, 0xc0, 0xca, 0xdf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x2a, 0x1f, 0x64, 0x81, 0x2a, 0x08, 0x00, 0x00,
}
