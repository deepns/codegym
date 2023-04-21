// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: reminders/reminders.proto

package reminders

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ReminderService_GetReminders_FullMethodName   = "/ReminderService/GetReminders"
	ReminderService_GetReminder_FullMethodName    = "/ReminderService/GetReminder"
	ReminderService_CreateReminder_FullMethodName = "/ReminderService/CreateReminder"
)

// ReminderServiceClient is the client API for ReminderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReminderServiceClient interface {
	// GetReminders returns a list of reminders.
	GetReminders(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetRemindersResponse, error)
	// GetReminder returns a specific reminder by ID.
	GetReminder(ctx context.Context, in *GetReminderRequest, opts ...grpc.CallOption) (*Reminder, error)
	// CreateReminder creates a new reminder.
	CreateReminder(ctx context.Context, in *Reminder, opts ...grpc.CallOption) (*CreateReminderResponse, error)
}

type reminderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReminderServiceClient(cc grpc.ClientConnInterface) ReminderServiceClient {
	return &reminderServiceClient{cc}
}

func (c *reminderServiceClient) GetReminders(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetRemindersResponse, error) {
	out := new(GetRemindersResponse)
	err := c.cc.Invoke(ctx, ReminderService_GetReminders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reminderServiceClient) GetReminder(ctx context.Context, in *GetReminderRequest, opts ...grpc.CallOption) (*Reminder, error) {
	out := new(Reminder)
	err := c.cc.Invoke(ctx, ReminderService_GetReminder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reminderServiceClient) CreateReminder(ctx context.Context, in *Reminder, opts ...grpc.CallOption) (*CreateReminderResponse, error) {
	out := new(CreateReminderResponse)
	err := c.cc.Invoke(ctx, ReminderService_CreateReminder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReminderServiceServer is the server API for ReminderService service.
// All implementations must embed UnimplementedReminderServiceServer
// for forward compatibility
type ReminderServiceServer interface {
	// GetReminders returns a list of reminders.
	GetReminders(context.Context, *Empty) (*GetRemindersResponse, error)
	// GetReminder returns a specific reminder by ID.
	GetReminder(context.Context, *GetReminderRequest) (*Reminder, error)
	// CreateReminder creates a new reminder.
	CreateReminder(context.Context, *Reminder) (*CreateReminderResponse, error)
	mustEmbedUnimplementedReminderServiceServer()
}

// UnimplementedReminderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReminderServiceServer struct {
}

func (UnimplementedReminderServiceServer) GetReminders(context.Context, *Empty) (*GetRemindersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReminders not implemented")
}
func (UnimplementedReminderServiceServer) GetReminder(context.Context, *GetReminderRequest) (*Reminder, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReminder not implemented")
}
func (UnimplementedReminderServiceServer) CreateReminder(context.Context, *Reminder) (*CreateReminderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReminder not implemented")
}
func (UnimplementedReminderServiceServer) mustEmbedUnimplementedReminderServiceServer() {}

// UnsafeReminderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReminderServiceServer will
// result in compilation errors.
type UnsafeReminderServiceServer interface {
	mustEmbedUnimplementedReminderServiceServer()
}

func RegisterReminderServiceServer(s grpc.ServiceRegistrar, srv ReminderServiceServer) {
	s.RegisterService(&ReminderService_ServiceDesc, srv)
}

func _ReminderService_GetReminders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReminderServiceServer).GetReminders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReminderService_GetReminders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReminderServiceServer).GetReminders(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReminderService_GetReminder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReminderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReminderServiceServer).GetReminder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReminderService_GetReminder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReminderServiceServer).GetReminder(ctx, req.(*GetReminderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReminderService_CreateReminder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Reminder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReminderServiceServer).CreateReminder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReminderService_CreateReminder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReminderServiceServer).CreateReminder(ctx, req.(*Reminder))
	}
	return interceptor(ctx, in, info, handler)
}

// ReminderService_ServiceDesc is the grpc.ServiceDesc for ReminderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReminderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ReminderService",
	HandlerType: (*ReminderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetReminders",
			Handler:    _ReminderService_GetReminders_Handler,
		},
		{
			MethodName: "GetReminder",
			Handler:    _ReminderService_GetReminder_Handler,
		},
		{
			MethodName: "CreateReminder",
			Handler:    _ReminderService_CreateReminder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reminders/reminders.proto",
}
