/*

Copyright 2017 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

*/

package server

import (
	"github.com/google/lvmd/commands"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	pb "github.com/google/lvmd/proto"
)

type Server struct{}

func NewServer() Server {
	return Server{}
}

func (s Server) ListLV(ctx context.Context, in *pb.ListLVRequest) (*pb.ListLVReply, error) {
	lvs, err := commands.ListLV(ctx, in.VolumeGroup)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "failed to list LVs: %v", err)
	}

	pblvs := make([]*pb.LogicalVolume, len(lvs))
	for i, v := range lvs {
		pblvs[i] = v.ToProto()
	}
	return &pb.ListLVReply{Volumes: pblvs}, nil
}

func (s Server) CreateLV(ctx context.Context, in *pb.CreateLVRequest) (*pb.CreateLVReply, error) {
	log, err := commands.CreateLV(ctx, in.VolumeGroup, in.Name, in.Size, in.Mirrors, in.Tags)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "failed to create lv: %v", err)
	}
	return &pb.CreateLVReply{CommandOutput: log}, nil
}

func (s Server) RemoveLV(ctx context.Context, in *pb.RemoveLVRequest) (*pb.RemoveLVReply, error) {
	log, err := commands.RemoveLV(ctx, in.VolumeGroup, in.Name)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "failed to remove lv: %v", err)
	}
	return &pb.RemoveLVReply{CommandOutput: log}, nil
}

func (s Server) CloneLV(ctx context.Context, in *pb.CloneLVRequest) (*pb.CloneLVReply, error) {
	log, err := commands.CloneLV(ctx, in.SourceName, in.DestName)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "failed to clone lv: %v", err)
	}
	return &pb.CloneLVReply{CommandOutput: log}, nil
}

func (s Server) ListVG(ctx context.Context, in *pb.ListVGRequest) (*pb.ListVGReply, error) {
	vgs, err := commands.ListVG(ctx)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "failed to list LVs: %v", err)
	}

	pbvgs := make([]*pb.VolumeGroup, len(vgs))
	for i, v := range vgs {
		pbvgs[i] = v.ToProto()
	}
	return &pb.ListVGReply{VolumeGroups: pbvgs}, nil
}

func (s Server) CreateVG(ctx context.Context, in *pb.CreateVGRequest) (*pb.CreateVGReply, error) {
	log, err := commands.CreateVG(ctx, in.Name, in.PhysicalVolume, in.Tags)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "failed to create vg: %v", err)
	}
	return &pb.CreateVGReply{CommandOutput: log}, nil
}

func (s Server) RemoveVG(ctx context.Context, in *pb.CreateVGRequest) (*pb.RemoveVGReply, error) {
	log, err := commands.RemoveVG(ctx, in.Name)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "failed to remove vg: %v", err)
	}
	return &pb.RemoveVGReply{CommandOutput: log}, nil
}

func (s Server) AddTagLV(ctx context.Context, in *pb.AddTagLVRequest) (*pb.AddTagLVReply, error) {
	log, err := commands.AddTagLV(ctx, in.VolumeGroup, in.Name, in.Tags)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "failed to add tags to lv: %v", err)
	}
	return &pb.AddTagLVReply{CommandOutput: log}, nil
}

func (s Server) RemoveTagLV(ctx context.Context, in *pb.RemoveTagLVRequest) (*pb.RemoveTagLVReply, error) {
	log, err := commands.RemoveTagLV(ctx, in.VolumeGroup, in.Name, in.Tags)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "failed to remove tags from lv: %v", err)
	}
	return &pb.RemoveTagLVReply{CommandOutput: log}, nil
}
