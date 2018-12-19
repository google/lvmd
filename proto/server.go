package lvm

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Server API for LVM service

type LVMServer interface {
	ListLV(context.Context, *ListLVRequest) (*ListLVReply, error)
	CreateLV(context.Context, *CreateLVRequest) (*CreateLVReply, error)
	RemoveLV(context.Context, *RemoveLVRequest) (*RemoveLVReply, error)
	CloneLV(context.Context, *CloneLVRequest) (*CloneLVReply, error)
}

func RegisterLVMServer(s *grpc.Server, srv LVMServer) {
	s.RegisterService(&_LVM_serviceDesc, srv)
}

func _LVM_ListLV_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLVRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LVMServer).ListLV(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lvm.LVM/ListLV",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LVMServer).ListLV(ctx, req.(*ListLVRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LVM_CreateLV_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLVRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LVMServer).CreateLV(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lvm.LVM/CreateLV",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LVMServer).CreateLV(ctx, req.(*CreateLVRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LVM_RemoveLV_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveLVRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LVMServer).RemoveLV(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lvm.LVM/RemoveLV",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LVMServer).RemoveLV(ctx, req.(*RemoveLVRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LVM_CloneLV_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloneLVRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LVMServer).CloneLV(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lvm.LVM/CloneLV",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LVMServer).CloneLV(ctx, req.(*CloneLVRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LVM_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lvm.LVM",
	HandlerType: (*LVMServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListLV",
			Handler:    _LVM_ListLV_Handler,
		},
		{
			MethodName: "CreateLV",
			Handler:    _LVM_CreateLV_Handler,
		},
		{
			MethodName: "RemoveLV",
			Handler:    _LVM_RemoveLV_Handler,
		},
		{
			MethodName: "CloneLV",
			Handler:    _LVM_CloneLV_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lvm.proto",
}
