package lvm

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Client API for LVM service

type LVMClient interface {
	ListLV(ctx context.Context, in *ListLVRequest, opts ...grpc.CallOption) (*ListLVReply, error)
	CreateLV(ctx context.Context, in *CreateLVRequest, opts ...grpc.CallOption) (*CreateLVReply, error)
	RemoveLV(ctx context.Context, in *RemoveLVRequest, opts ...grpc.CallOption) (*RemoveLVReply, error)
	CloneLV(ctx context.Context, in *CloneLVRequest, opts ...grpc.CallOption) (*CloneLVReply, error)
}

type lVMClient struct {
	cc *grpc.ClientConn
}

func NewLVMClient(cc *grpc.ClientConn) LVMClient {
	return &lVMClient{cc}
}

func (c *lVMClient) ListLV(ctx context.Context, in *ListLVRequest, opts ...grpc.CallOption) (*ListLVReply, error) {
	out := new(ListLVReply)
	err := grpc.Invoke(ctx, "/lvm.LVM/ListLV", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lVMClient) CreateLV(ctx context.Context, in *CreateLVRequest, opts ...grpc.CallOption) (*CreateLVReply, error) {
	out := new(CreateLVReply)
	err := grpc.Invoke(ctx, "/lvm.LVM/CreateLV", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lVMClient) RemoveLV(ctx context.Context, in *RemoveLVRequest, opts ...grpc.CallOption) (*RemoveLVReply, error) {
	out := new(RemoveLVReply)
	err := grpc.Invoke(ctx, "/lvm.LVM/RemoveLV", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lVMClient) CloneLV(ctx context.Context, in *CloneLVRequest, opts ...grpc.CallOption) (*CloneLVReply, error) {
	out := new(CloneLVReply)
	err := grpc.Invoke(ctx, "/lvm.LVM/CloneLV", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
