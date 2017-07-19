// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: payload.proto

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ExportPayload_Status int32

const (
	ExportPayload_NONE      ExportPayload_Status = 0
	ExportPayload_SUCCESS   ExportPayload_Status = 1
	ExportPayload_DUPLICATE ExportPayload_Status = 2
	ExportPayload_FAILED    ExportPayload_Status = 3
)

var ExportPayload_Status_name = map[int32]string{
	0: "NONE",
	1: "SUCCESS",
	2: "DUPLICATE",
	3: "FAILED",
}
var ExportPayload_Status_value = map[string]int32{
	"NONE":      0,
	"SUCCESS":   1,
	"DUPLICATE": 2,
	"FAILED":    3,
}

func (x ExportPayload_Status) String() string {
	return proto.EnumName(ExportPayload_Status_name, int32(x))
}
func (ExportPayload_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorPayload, []int{1, 0}
}

type Payload struct {
	Data []byte `protobuf:"bytes,1,opt,name=Data,json=data,proto3" json:"Data,omitempty"`
}

func (m *Payload) Reset()                    { *m = Payload{} }
func (m *Payload) String() string            { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()               {}
func (*Payload) Descriptor() ([]byte, []int) { return fileDescriptorPayload, []int{0} }

func (m *Payload) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// BackupPayload is used both as a request and a response.
// When used in request, groups represents the list of groups that need to be backed up.
// When used in response, groups represent the list of groups that were backed up.
type ExportPayload struct {
	ReqId   uint64               `protobuf:"varint,1,opt,name=req_id,json=reqId,proto3" json:"req_id,omitempty"`
	GroupId uint32               `protobuf:"varint,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	Status  ExportPayload_Status `protobuf:"varint,3,opt,name=status,proto3,enum=protos.ExportPayload_Status" json:"status,omitempty"`
}

func (m *ExportPayload) Reset()                    { *m = ExportPayload{} }
func (m *ExportPayload) String() string            { return proto.CompactTextString(m) }
func (*ExportPayload) ProtoMessage()               {}
func (*ExportPayload) Descriptor() ([]byte, []int) { return fileDescriptorPayload, []int{1} }

func (m *ExportPayload) GetReqId() uint64 {
	if m != nil {
		return m.ReqId
	}
	return 0
}

func (m *ExportPayload) GetGroupId() uint32 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *ExportPayload) GetStatus() ExportPayload_Status {
	if m != nil {
		return m.Status
	}
	return ExportPayload_NONE
}

func init() {
	proto.RegisterType((*Payload)(nil), "protos.Payload")
	proto.RegisterType((*ExportPayload)(nil), "protos.ExportPayload")
	proto.RegisterEnum("protos.ExportPayload_Status", ExportPayload_Status_name, ExportPayload_Status_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Worker service

type WorkerClient interface {
	// Connection testing RPC.
	Echo(ctx context.Context, in *Payload, opts ...grpc.CallOption) (*Payload, error)
	// Data serving RPCs.
	AssignUids(ctx context.Context, in *Num, opts ...grpc.CallOption) (*AssignedIds, error)
	Mutate(ctx context.Context, in *Mutations, opts ...grpc.CallOption) (*Payload, error)
	ServeTask(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Result, error)
	PredicateAndSchemaData(ctx context.Context, opts ...grpc.CallOption) (Worker_PredicateAndSchemaDataClient, error)
	Sort(ctx context.Context, in *SortMessage, opts ...grpc.CallOption) (*SortResult, error)
	Schema(ctx context.Context, in *SchemaRequest, opts ...grpc.CallOption) (*SchemaResult, error)
	// RAFT serving RPCs.
	RaftMessage(ctx context.Context, in *Payload, opts ...grpc.CallOption) (*Payload, error)
	// The stream takes the same sort of Payload values that RaftMessage takes.
	// Never returns.
	RaftMessageStream(ctx context.Context, opts ...grpc.CallOption) (Worker_RaftMessageStreamClient, error)
	JoinCluster(ctx context.Context, in *RaftContext, opts ...grpc.CallOption) (*Payload, error)
	UpdateMembership(ctx context.Context, in *MembershipUpdate, opts ...grpc.CallOption) (*MembershipUpdate, error)
	Export(ctx context.Context, in *ExportPayload, opts ...grpc.CallOption) (*ExportPayload, error)
}

type workerClient struct {
	cc *grpc.ClientConn
}

func NewWorkerClient(cc *grpc.ClientConn) WorkerClient {
	return &workerClient{cc}
}

func (c *workerClient) Echo(ctx context.Context, in *Payload, opts ...grpc.CallOption) (*Payload, error) {
	out := new(Payload)
	err := grpc.Invoke(ctx, "/protos.Worker/Echo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) AssignUids(ctx context.Context, in *Num, opts ...grpc.CallOption) (*AssignedIds, error) {
	out := new(AssignedIds)
	err := grpc.Invoke(ctx, "/protos.Worker/AssignUids", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) Mutate(ctx context.Context, in *Mutations, opts ...grpc.CallOption) (*Payload, error) {
	out := new(Payload)
	err := grpc.Invoke(ctx, "/protos.Worker/Mutate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) ServeTask(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := grpc.Invoke(ctx, "/protos.Worker/ServeTask", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) PredicateAndSchemaData(ctx context.Context, opts ...grpc.CallOption) (Worker_PredicateAndSchemaDataClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Worker_serviceDesc.Streams[0], c.cc, "/protos.Worker/PredicateAndSchemaData", opts...)
	if err != nil {
		return nil, err
	}
	x := &workerPredicateAndSchemaDataClient{stream}
	return x, nil
}

type Worker_PredicateAndSchemaDataClient interface {
	Send(*GroupKeys) error
	Recv() (*KV, error)
	grpc.ClientStream
}

type workerPredicateAndSchemaDataClient struct {
	grpc.ClientStream
}

func (x *workerPredicateAndSchemaDataClient) Send(m *GroupKeys) error {
	return x.ClientStream.SendMsg(m)
}

func (x *workerPredicateAndSchemaDataClient) Recv() (*KV, error) {
	m := new(KV)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *workerClient) Sort(ctx context.Context, in *SortMessage, opts ...grpc.CallOption) (*SortResult, error) {
	out := new(SortResult)
	err := grpc.Invoke(ctx, "/protos.Worker/Sort", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) Schema(ctx context.Context, in *SchemaRequest, opts ...grpc.CallOption) (*SchemaResult, error) {
	out := new(SchemaResult)
	err := grpc.Invoke(ctx, "/protos.Worker/Schema", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) RaftMessage(ctx context.Context, in *Payload, opts ...grpc.CallOption) (*Payload, error) {
	out := new(Payload)
	err := grpc.Invoke(ctx, "/protos.Worker/RaftMessage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) RaftMessageStream(ctx context.Context, opts ...grpc.CallOption) (Worker_RaftMessageStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Worker_serviceDesc.Streams[1], c.cc, "/protos.Worker/RaftMessageStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &workerRaftMessageStreamClient{stream}
	return x, nil
}

type Worker_RaftMessageStreamClient interface {
	Send(*Payload) error
	CloseAndRecv() (*Payload, error)
	grpc.ClientStream
}

type workerRaftMessageStreamClient struct {
	grpc.ClientStream
}

func (x *workerRaftMessageStreamClient) Send(m *Payload) error {
	return x.ClientStream.SendMsg(m)
}

func (x *workerRaftMessageStreamClient) CloseAndRecv() (*Payload, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Payload)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *workerClient) JoinCluster(ctx context.Context, in *RaftContext, opts ...grpc.CallOption) (*Payload, error) {
	out := new(Payload)
	err := grpc.Invoke(ctx, "/protos.Worker/JoinCluster", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) UpdateMembership(ctx context.Context, in *MembershipUpdate, opts ...grpc.CallOption) (*MembershipUpdate, error) {
	out := new(MembershipUpdate)
	err := grpc.Invoke(ctx, "/protos.Worker/UpdateMembership", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) Export(ctx context.Context, in *ExportPayload, opts ...grpc.CallOption) (*ExportPayload, error) {
	out := new(ExportPayload)
	err := grpc.Invoke(ctx, "/protos.Worker/Export", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Worker service

type WorkerServer interface {
	// Connection testing RPC.
	Echo(context.Context, *Payload) (*Payload, error)
	// Data serving RPCs.
	AssignUids(context.Context, *Num) (*AssignedIds, error)
	Mutate(context.Context, *Mutations) (*Payload, error)
	ServeTask(context.Context, *Query) (*Result, error)
	PredicateAndSchemaData(Worker_PredicateAndSchemaDataServer) error
	Sort(context.Context, *SortMessage) (*SortResult, error)
	Schema(context.Context, *SchemaRequest) (*SchemaResult, error)
	// RAFT serving RPCs.
	RaftMessage(context.Context, *Payload) (*Payload, error)
	// The stream takes the same sort of Payload values that RaftMessage takes.
	// Never returns.
	RaftMessageStream(Worker_RaftMessageStreamServer) error
	JoinCluster(context.Context, *RaftContext) (*Payload, error)
	UpdateMembership(context.Context, *MembershipUpdate) (*MembershipUpdate, error)
	Export(context.Context, *ExportPayload) (*ExportPayload, error)
}

func RegisterWorkerServer(s *grpc.Server, srv WorkerServer) {
	s.RegisterService(&_Worker_serviceDesc, srv)
}

func _Worker_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Payload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Worker/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).Echo(ctx, req.(*Payload))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_AssignUids_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Num)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).AssignUids(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Worker/AssignUids",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).AssignUids(ctx, req.(*Num))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_Mutate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Mutations)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).Mutate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Worker/Mutate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).Mutate(ctx, req.(*Mutations))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_ServeTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).ServeTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Worker/ServeTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).ServeTask(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_PredicateAndSchemaData_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WorkerServer).PredicateAndSchemaData(&workerPredicateAndSchemaDataServer{stream})
}

type Worker_PredicateAndSchemaDataServer interface {
	Send(*KV) error
	Recv() (*GroupKeys, error)
	grpc.ServerStream
}

type workerPredicateAndSchemaDataServer struct {
	grpc.ServerStream
}

func (x *workerPredicateAndSchemaDataServer) Send(m *KV) error {
	return x.ServerStream.SendMsg(m)
}

func (x *workerPredicateAndSchemaDataServer) Recv() (*GroupKeys, error) {
	m := new(GroupKeys)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Worker_Sort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SortMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).Sort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Worker/Sort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).Sort(ctx, req.(*SortMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_Schema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).Schema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Worker/Schema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).Schema(ctx, req.(*SchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_RaftMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Payload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).RaftMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Worker/RaftMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).RaftMessage(ctx, req.(*Payload))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_RaftMessageStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WorkerServer).RaftMessageStream(&workerRaftMessageStreamServer{stream})
}

type Worker_RaftMessageStreamServer interface {
	SendAndClose(*Payload) error
	Recv() (*Payload, error)
	grpc.ServerStream
}

type workerRaftMessageStreamServer struct {
	grpc.ServerStream
}

func (x *workerRaftMessageStreamServer) SendAndClose(m *Payload) error {
	return x.ServerStream.SendMsg(m)
}

func (x *workerRaftMessageStreamServer) Recv() (*Payload, error) {
	m := new(Payload)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Worker_JoinCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RaftContext)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).JoinCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Worker/JoinCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).JoinCluster(ctx, req.(*RaftContext))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_UpdateMembership_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MembershipUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).UpdateMembership(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Worker/UpdateMembership",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).UpdateMembership(ctx, req.(*MembershipUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_Export_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportPayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).Export(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Worker/Export",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).Export(ctx, req.(*ExportPayload))
	}
	return interceptor(ctx, in, info, handler)
}

var _Worker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Worker",
	HandlerType: (*WorkerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _Worker_Echo_Handler,
		},
		{
			MethodName: "AssignUids",
			Handler:    _Worker_AssignUids_Handler,
		},
		{
			MethodName: "Mutate",
			Handler:    _Worker_Mutate_Handler,
		},
		{
			MethodName: "ServeTask",
			Handler:    _Worker_ServeTask_Handler,
		},
		{
			MethodName: "Sort",
			Handler:    _Worker_Sort_Handler,
		},
		{
			MethodName: "Schema",
			Handler:    _Worker_Schema_Handler,
		},
		{
			MethodName: "RaftMessage",
			Handler:    _Worker_RaftMessage_Handler,
		},
		{
			MethodName: "JoinCluster",
			Handler:    _Worker_JoinCluster_Handler,
		},
		{
			MethodName: "UpdateMembership",
			Handler:    _Worker_UpdateMembership_Handler,
		},
		{
			MethodName: "Export",
			Handler:    _Worker_Export_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PredicateAndSchemaData",
			Handler:       _Worker_PredicateAndSchemaData_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "RaftMessageStream",
			Handler:       _Worker_RaftMessageStream_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "payload.proto",
}

func (m *Payload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Payload) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPayload(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
	}
	return i, nil
}

func (m *ExportPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExportPayload) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ReqId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintPayload(dAtA, i, uint64(m.ReqId))
	}
	if m.GroupId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintPayload(dAtA, i, uint64(m.GroupId))
	}
	if m.Status != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintPayload(dAtA, i, uint64(m.Status))
	}
	return i, nil
}

func encodeFixed64Payload(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Payload(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintPayload(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Payload) Size() (n int) {
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovPayload(uint64(l))
	}
	return n
}

func (m *ExportPayload) Size() (n int) {
	var l int
	_ = l
	if m.ReqId != 0 {
		n += 1 + sovPayload(uint64(m.ReqId))
	}
	if m.GroupId != 0 {
		n += 1 + sovPayload(uint64(m.GroupId))
	}
	if m.Status != 0 {
		n += 1 + sovPayload(uint64(m.Status))
	}
	return n
}

func sovPayload(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPayload(x uint64) (n int) {
	return sovPayload(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Payload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPayload
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Payload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Payload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPayload
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPayload(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPayload
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ExportPayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPayload
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ExportPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExportPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReqId", wireType)
			}
			m.ReqId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReqId |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupId", wireType)
			}
			m.GroupId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GroupId |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= (ExportPayload_Status(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPayload(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPayload
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPayload(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPayload
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPayload
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPayload
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthPayload
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPayload
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipPayload(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthPayload = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPayload   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("payload.proto", fileDescriptorPayload) }

var fileDescriptorPayload = []byte{
	// 539 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xdb, 0x6e, 0xd3, 0x4c,
	0x10, 0xf6, 0xb6, 0xae, 0xdb, 0x4e, 0x9a, 0xfe, 0xee, 0xf4, 0x2f, 0x2a, 0x16, 0x44, 0x95, 0xaf,
	0x22, 0x84, 0xa2, 0xb6, 0x80, 0x38, 0x48, 0x5c, 0x84, 0xd4, 0x40, 0xe8, 0x81, 0x62, 0x37, 0x70,
	0x89, 0xb6, 0xf5, 0x90, 0x58, 0x49, 0xbc, 0xce, 0xee, 0x1a, 0x35, 0x6f, 0xd2, 0xf7, 0xe0, 0x25,
	0xb8, 0xe4, 0x11, 0x50, 0x78, 0x11, 0x14, 0x3b, 0x4e, 0xda, 0x2a, 0x48, 0xbd, 0xb2, 0xbf, 0xc3,
	0xcc, 0x67, 0xcf, 0xec, 0x42, 0x39, 0xe1, 0xc3, 0x9e, 0xe0, 0x61, 0x2d, 0x91, 0x42, 0x0b, 0xb4,
	0xb2, 0x87, 0x72, 0x36, 0xdb, 0x92, 0x27, 0x1d, 0x49, 0x2a, 0x11, 0xb1, 0xa2, 0x5c, 0x74, 0xd6,
	0xd4, 0x45, 0x87, 0xfa, 0x7c, 0x82, 0x40, 0x73, 0xd5, 0xcd, 0xdf, 0xdd, 0x87, 0xb0, 0x7c, 0x9a,
	0xf7, 0x41, 0x04, 0xf3, 0x80, 0x6b, 0xbe, 0xcd, 0x76, 0x58, 0x75, 0xcd, 0x37, 0x43, 0xae, 0xb9,
	0xfb, 0x83, 0x41, 0xd9, 0xbb, 0x4c, 0x84, 0xd4, 0x85, 0x6b, 0x0b, 0x2c, 0x49, 0x83, 0xaf, 0x51,
	0x98, 0xf9, 0x4c, 0x7f, 0x49, 0xd2, 0xa0, 0x19, 0xe2, 0x7d, 0x58, 0x69, 0x4b, 0x91, 0x26, 0x63,
	0x61, 0x61, 0x87, 0x55, 0xcb, 0xfe, 0x72, 0x86, 0x9b, 0x21, 0x3e, 0x05, 0x4b, 0x69, 0xae, 0x53,
	0xb5, 0xbd, 0xb8, 0xc3, 0xaa, 0xeb, 0xfb, 0x0f, 0xf2, 0x68, 0x55, 0xbb, 0xd1, 0xb8, 0x16, 0x64,
	0x1e, 0x7f, 0xe2, 0x75, 0x5f, 0x81, 0x95, 0x33, 0xb8, 0x02, 0xe6, 0xc9, 0xc7, 0x13, 0xcf, 0x36,
	0xb0, 0x04, 0xcb, 0x41, 0xab, 0xd1, 0xf0, 0x82, 0xc0, 0x66, 0x58, 0x86, 0xd5, 0x83, 0xd6, 0xe9,
	0x51, 0xb3, 0x51, 0x3f, 0xf3, 0xec, 0x05, 0x04, 0xb0, 0xde, 0xd6, 0x9b, 0x47, 0xde, 0x81, 0xbd,
	0xb8, 0x7f, 0xb5, 0x04, 0xd6, 0x17, 0x21, 0xbb, 0x24, 0xf1, 0x11, 0x98, 0xde, 0x45, 0x47, 0xe0,
	0x7f, 0x45, 0xe8, 0x24, 0xce, 0xb9, 0x4d, 0xb8, 0x06, 0xee, 0x02, 0xd4, 0x95, 0x8a, 0xda, 0x71,
	0x2b, 0x0a, 0x15, 0x96, 0x0a, 0xc3, 0x49, 0xda, 0x77, 0x36, 0x0b, 0x90, 0x1b, 0x28, 0x6c, 0x86,
	0xca, 0x35, 0xb0, 0x06, 0xd6, 0x71, 0xaa, 0xb9, 0x26, 0xdc, 0x28, 0x0c, 0x19, 0x8e, 0x44, 0xac,
	0xe6, 0x25, 0x3c, 0x86, 0xd5, 0x80, 0xe4, 0x77, 0x3a, 0xe3, 0xaa, 0x8b, 0xe5, 0x42, 0xff, 0x94,
	0x92, 0x1c, 0x3a, 0xeb, 0x05, 0xf4, 0x49, 0xa5, 0x3d, 0xed, 0x1a, 0xf8, 0x1a, 0xee, 0x9d, 0x4a,
	0x0a, 0xa3, 0x0b, 0xae, 0xa9, 0x1e, 0x87, 0x41, 0xb6, 0xc3, 0xf1, 0x8a, 0x66, 0x69, 0xef, 0xc6,
	0x33, 0x3e, 0xa4, 0xa1, 0x72, 0xa0, 0xa0, 0x0e, 0x3f, 0xbb, 0x46, 0x95, 0xed, 0x32, 0xdc, 0x03,
	0x33, 0x10, 0x52, 0xe3, 0xf4, 0xdb, 0xc7, 0xe8, 0x98, 0x94, 0xe2, 0x6d, 0x72, 0xf0, 0x3a, 0x39,
	0x4d, 0x7c, 0x0e, 0x56, 0x9e, 0x82, 0x5b, 0x53, 0x3d, 0xc3, 0x3e, 0x0d, 0x52, 0x52, 0xda, 0xf9,
	0xff, 0x36, 0x3d, 0x29, 0xdc, 0x83, 0x92, 0xcf, 0xbf, 0x15, 0xdd, 0xef, 0x34, 0xed, 0x97, 0xb0,
	0x71, 0xad, 0x24, 0xd0, 0x92, 0x78, 0xff, 0x2e, 0x85, 0x55, 0x86, 0xcf, 0xa0, 0xf4, 0x41, 0x44,
	0x71, 0xa3, 0x97, 0x2a, 0x4d, 0x72, 0xf6, 0x83, 0xe3, 0x7e, 0x0d, 0x11, 0x6b, 0xba, 0xd4, 0xf3,
	0x12, 0xdf, 0x83, 0xdd, 0x4a, 0x42, 0xae, 0xe9, 0x98, 0xfa, 0xe7, 0x24, 0x55, 0x27, 0x4a, 0x70,
	0x7b, 0xba, 0xb7, 0x29, 0x97, 0x7b, 0x9c, 0x7f, 0x2a, 0xae, 0x81, 0x2f, 0xc0, 0xca, 0x0f, 0xef,
	0x6c, 0x4e, 0x37, 0x0e, 0xb3, 0x33, 0x9f, 0x76, 0x8d, 0x37, 0xf6, 0xcf, 0x51, 0x85, 0xfd, 0x1a,
	0x55, 0xd8, 0xef, 0x51, 0x85, 0x5d, 0xfd, 0xa9, 0x18, 0xe7, 0xf9, 0xc5, 0x7d, 0xf2, 0x37, 0x00,
	0x00, 0xff, 0xff, 0x5b, 0x12, 0x77, 0x23, 0xd0, 0x03, 0x00, 0x00,
}
