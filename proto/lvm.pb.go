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

// Code generated by protoc-gen-go.
// source: lvm.proto
// DO NOT EDIT!

/*
Package lvm is a generated protocol buffer package.

It is generated from these files:
	lvm.proto

It has these top-level messages:
	LogicalVolume
	ListLVRequest
	ListLVReply
	CreateLVRequest
	CreateLVReply
	RemoveLVRequest
	RemoveLVReply
	CloneLVRequest
	CloneLVReply
*/
package lvm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type LogicalVolume_Attributes_Type int32

const (
	LogicalVolume_Attributes_MALFORMED_TYPE               LogicalVolume_Attributes_Type = 0
	LogicalVolume_Attributes_MIRRORED                     LogicalVolume_Attributes_Type = 1
	LogicalVolume_Attributes_MIRRORED_WITHOUT_SYNC        LogicalVolume_Attributes_Type = 2
	LogicalVolume_Attributes_ORIGIN                       LogicalVolume_Attributes_Type = 3
	LogicalVolume_Attributes_ORIGIN_WITH_MERGING_SNAPSHOT LogicalVolume_Attributes_Type = 4
	LogicalVolume_Attributes_RAID                         LogicalVolume_Attributes_Type = 5
	LogicalVolume_Attributes_RAID_WITHOUT_SYNC            LogicalVolume_Attributes_Type = 6
	LogicalVolume_Attributes_SNAPSHOT                     LogicalVolume_Attributes_Type = 7
	LogicalVolume_Attributes_MERGING_SNAPSHOT             LogicalVolume_Attributes_Type = 8
	LogicalVolume_Attributes_PV_MOVE                      LogicalVolume_Attributes_Type = 9
	LogicalVolume_Attributes_VIRTUAL_MIRROR               LogicalVolume_Attributes_Type = 10
	LogicalVolume_Attributes_VIRTUAL_RAID_IMAGE           LogicalVolume_Attributes_Type = 11
	LogicalVolume_Attributes_RAID_IMAGE_OUT_OF_SYNC       LogicalVolume_Attributes_Type = 12
	LogicalVolume_Attributes_MIRROR_LOG                   LogicalVolume_Attributes_Type = 13
	LogicalVolume_Attributes_UNDER_CONVERSION             LogicalVolume_Attributes_Type = 14
	LogicalVolume_Attributes_THIN                         LogicalVolume_Attributes_Type = 15
	LogicalVolume_Attributes_THIN_POOL                    LogicalVolume_Attributes_Type = 16
	LogicalVolume_Attributes_THIN_POOL_DATA               LogicalVolume_Attributes_Type = 17
	LogicalVolume_Attributes_RAID_OR_THIN_POOL_METADATA   LogicalVolume_Attributes_Type = 18
)

var LogicalVolume_Attributes_Type_name = map[int32]string{
	0:  "MALFORMED_TYPE",
	1:  "MIRRORED",
	2:  "MIRRORED_WITHOUT_SYNC",
	3:  "ORIGIN",
	4:  "ORIGIN_WITH_MERGING_SNAPSHOT",
	5:  "RAID",
	6:  "RAID_WITHOUT_SYNC",
	7:  "SNAPSHOT",
	8:  "MERGING_SNAPSHOT",
	9:  "PV_MOVE",
	10: "VIRTUAL_MIRROR",
	11: "VIRTUAL_RAID_IMAGE",
	12: "RAID_IMAGE_OUT_OF_SYNC",
	13: "MIRROR_LOG",
	14: "UNDER_CONVERSION",
	15: "THIN",
	16: "THIN_POOL",
	17: "THIN_POOL_DATA",
	18: "RAID_OR_THIN_POOL_METADATA",
}
var LogicalVolume_Attributes_Type_value = map[string]int32{
	"MALFORMED_TYPE":               0,
	"MIRRORED":                     1,
	"MIRRORED_WITHOUT_SYNC":        2,
	"ORIGIN":                       3,
	"ORIGIN_WITH_MERGING_SNAPSHOT": 4,
	"RAID":                       5,
	"RAID_WITHOUT_SYNC":          6,
	"SNAPSHOT":                   7,
	"MERGING_SNAPSHOT":           8,
	"PV_MOVE":                    9,
	"VIRTUAL_MIRROR":             10,
	"VIRTUAL_RAID_IMAGE":         11,
	"RAID_IMAGE_OUT_OF_SYNC":     12,
	"MIRROR_LOG":                 13,
	"UNDER_CONVERSION":           14,
	"THIN":                       15,
	"THIN_POOL":                  16,
	"THIN_POOL_DATA":             17,
	"RAID_OR_THIN_POOL_METADATA": 18,
}

func (x LogicalVolume_Attributes_Type) String() string {
	return proto.EnumName(LogicalVolume_Attributes_Type_name, int32(x))
}
func (LogicalVolume_Attributes_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0, 0}
}

type LogicalVolume_Attributes_Permissions int32

const (
	LogicalVolume_Attributes_MALFORMED_PERMISSIONS LogicalVolume_Attributes_Permissions = 0
	LogicalVolume_Attributes_WRITEABLE             LogicalVolume_Attributes_Permissions = 1
	LogicalVolume_Attributes_READ_ONLY             LogicalVolume_Attributes_Permissions = 2
	LogicalVolume_Attributes_READ_ONLY_ACTIVATION  LogicalVolume_Attributes_Permissions = 3
)

var LogicalVolume_Attributes_Permissions_name = map[int32]string{
	0: "MALFORMED_PERMISSIONS",
	1: "WRITEABLE",
	2: "READ_ONLY",
	3: "READ_ONLY_ACTIVATION",
}
var LogicalVolume_Attributes_Permissions_value = map[string]int32{
	"MALFORMED_PERMISSIONS": 0,
	"WRITEABLE":             1,
	"READ_ONLY":             2,
	"READ_ONLY_ACTIVATION":  3,
}

func (x LogicalVolume_Attributes_Permissions) String() string {
	return proto.EnumName(LogicalVolume_Attributes_Permissions_name, int32(x))
}
func (LogicalVolume_Attributes_Permissions) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0, 1}
}

type LogicalVolume_Attributes_Allocation int32

const (
	LogicalVolume_Attributes_MALFORMED_ALLOCATION LogicalVolume_Attributes_Allocation = 0
	LogicalVolume_Attributes_ANYWHERE             LogicalVolume_Attributes_Allocation = 1
	LogicalVolume_Attributes_CONTIGUOUS           LogicalVolume_Attributes_Allocation = 2
	LogicalVolume_Attributes_INHERITED            LogicalVolume_Attributes_Allocation = 3
	LogicalVolume_Attributes_CLING                LogicalVolume_Attributes_Allocation = 4
	LogicalVolume_Attributes_NORMAL               LogicalVolume_Attributes_Allocation = 5
	LogicalVolume_Attributes_ANYWHERE_LOCKED      LogicalVolume_Attributes_Allocation = 6
	LogicalVolume_Attributes_CONTIGUOUS_LOCKED    LogicalVolume_Attributes_Allocation = 7
	LogicalVolume_Attributes_INHERITED_LOCKED     LogicalVolume_Attributes_Allocation = 8
	LogicalVolume_Attributes_CLING_LOCKED         LogicalVolume_Attributes_Allocation = 9
	LogicalVolume_Attributes_NORMAL_LOCKED        LogicalVolume_Attributes_Allocation = 10
)

var LogicalVolume_Attributes_Allocation_name = map[int32]string{
	0:  "MALFORMED_ALLOCATION",
	1:  "ANYWHERE",
	2:  "CONTIGUOUS",
	3:  "INHERITED",
	4:  "CLING",
	5:  "NORMAL",
	6:  "ANYWHERE_LOCKED",
	7:  "CONTIGUOUS_LOCKED",
	8:  "INHERITED_LOCKED",
	9:  "CLING_LOCKED",
	10: "NORMAL_LOCKED",
}
var LogicalVolume_Attributes_Allocation_value = map[string]int32{
	"MALFORMED_ALLOCATION": 0,
	"ANYWHERE":             1,
	"CONTIGUOUS":           2,
	"INHERITED":            3,
	"CLING":                4,
	"NORMAL":               5,
	"ANYWHERE_LOCKED":      6,
	"CONTIGUOUS_LOCKED":    7,
	"INHERITED_LOCKED":     8,
	"CLING_LOCKED":         9,
	"NORMAL_LOCKED":        10,
}

func (x LogicalVolume_Attributes_Allocation) String() string {
	return proto.EnumName(LogicalVolume_Attributes_Allocation_name, int32(x))
}
func (LogicalVolume_Attributes_Allocation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0, 2}
}

type LogicalVolume_Attributes_State int32

const (
	LogicalVolume_Attributes_MALFORMED_STATE                           LogicalVolume_Attributes_State = 0
	LogicalVolume_Attributes_ACTIVE                                    LogicalVolume_Attributes_State = 1
	LogicalVolume_Attributes_SUSPENDED                                 LogicalVolume_Attributes_State = 2
	LogicalVolume_Attributes_INVALID_SNAPSHOT                          LogicalVolume_Attributes_State = 3
	LogicalVolume_Attributes_INVALID_SUSPENDED_SNAPSHOT                LogicalVolume_Attributes_State = 4
	LogicalVolume_Attributes_SNAPSHOT_MERGE_FAILED                     LogicalVolume_Attributes_State = 5
	LogicalVolume_Attributes_SUSPENDED_SNAPSHOT_MERGE_FAILED           LogicalVolume_Attributes_State = 6
	LogicalVolume_Attributes_MAPPED_DEVICE_PRESENT_WITHOUT_TABLES      LogicalVolume_Attributes_State = 7
	LogicalVolume_Attributes_MAPPED_DEVICE_PRESENT_WITH_INACTIVE_TABLE LogicalVolume_Attributes_State = 8
)

var LogicalVolume_Attributes_State_name = map[int32]string{
	0: "MALFORMED_STATE",
	1: "ACTIVE",
	2: "SUSPENDED",
	3: "INVALID_SNAPSHOT",
	4: "INVALID_SUSPENDED_SNAPSHOT",
	5: "SNAPSHOT_MERGE_FAILED",
	6: "SUSPENDED_SNAPSHOT_MERGE_FAILED",
	7: "MAPPED_DEVICE_PRESENT_WITHOUT_TABLES",
	8: "MAPPED_DEVICE_PRESENT_WITH_INACTIVE_TABLE",
}
var LogicalVolume_Attributes_State_value = map[string]int32{
	"MALFORMED_STATE":                           0,
	"ACTIVE":                                    1,
	"SUSPENDED":                                 2,
	"INVALID_SNAPSHOT":                          3,
	"INVALID_SUSPENDED_SNAPSHOT":                4,
	"SNAPSHOT_MERGE_FAILED":                     5,
	"SUSPENDED_SNAPSHOT_MERGE_FAILED":           6,
	"MAPPED_DEVICE_PRESENT_WITHOUT_TABLES":      7,
	"MAPPED_DEVICE_PRESENT_WITH_INACTIVE_TABLE": 8,
}

func (x LogicalVolume_Attributes_State) String() string {
	return proto.EnumName(LogicalVolume_Attributes_State_name, int32(x))
}
func (LogicalVolume_Attributes_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0, 3}
}

type LogicalVolume_Attributes_TargetType int32

const (
	LogicalVolume_Attributes_MALFORMED_TARGET LogicalVolume_Attributes_TargetType = 0
	LogicalVolume_Attributes_MIRROR_TARGET    LogicalVolume_Attributes_TargetType = 1
	LogicalVolume_Attributes_RAID_TARGET      LogicalVolume_Attributes_TargetType = 2
	LogicalVolume_Attributes_SNAPSHOT_TARGET  LogicalVolume_Attributes_TargetType = 3
	LogicalVolume_Attributes_THIN_TARGET      LogicalVolume_Attributes_TargetType = 4
	LogicalVolume_Attributes_UNKNOWN_TARGET   LogicalVolume_Attributes_TargetType = 5
	LogicalVolume_Attributes_VIRTUAL_TARGET   LogicalVolume_Attributes_TargetType = 6
)

var LogicalVolume_Attributes_TargetType_name = map[int32]string{
	0: "MALFORMED_TARGET",
	1: "MIRROR_TARGET",
	2: "RAID_TARGET",
	3: "SNAPSHOT_TARGET",
	4: "THIN_TARGET",
	5: "UNKNOWN_TARGET",
	6: "VIRTUAL_TARGET",
}
var LogicalVolume_Attributes_TargetType_value = map[string]int32{
	"MALFORMED_TARGET": 0,
	"MIRROR_TARGET":    1,
	"RAID_TARGET":      2,
	"SNAPSHOT_TARGET":  3,
	"THIN_TARGET":      4,
	"UNKNOWN_TARGET":   5,
	"VIRTUAL_TARGET":   6,
}

func (x LogicalVolume_Attributes_TargetType) String() string {
	return proto.EnumName(LogicalVolume_Attributes_TargetType_name, int32(x))
}
func (LogicalVolume_Attributes_TargetType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0, 4}
}

type LogicalVolume_Attributes_Health int32

const (
	LogicalVolume_Attributes_MALFORMED_HEALTH LogicalVolume_Attributes_Health = 0
	LogicalVolume_Attributes_OK               LogicalVolume_Attributes_Health = 1
	LogicalVolume_Attributes_PARTIAL          LogicalVolume_Attributes_Health = 2
	LogicalVolume_Attributes_REFRESH_NEEDED   LogicalVolume_Attributes_Health = 3
	LogicalVolume_Attributes_MISMATCHES_EXIST LogicalVolume_Attributes_Health = 4
	LogicalVolume_Attributes_WRITEMOSTLY      LogicalVolume_Attributes_Health = 5
)

var LogicalVolume_Attributes_Health_name = map[int32]string{
	0: "MALFORMED_HEALTH",
	1: "OK",
	2: "PARTIAL",
	3: "REFRESH_NEEDED",
	4: "MISMATCHES_EXIST",
	5: "WRITEMOSTLY",
}
var LogicalVolume_Attributes_Health_value = map[string]int32{
	"MALFORMED_HEALTH": 0,
	"OK":               1,
	"PARTIAL":          2,
	"REFRESH_NEEDED":   3,
	"MISMATCHES_EXIST": 4,
	"WRITEMOSTLY":      5,
}

func (x LogicalVolume_Attributes_Health) String() string {
	return proto.EnumName(LogicalVolume_Attributes_Health_name, int32(x))
}
func (LogicalVolume_Attributes_Health) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0, 5}
}

type LogicalVolume struct {
	Name                 string                    `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Size                 uint64                    `protobuf:"varint,2,opt,name=size" json:"size,omitempty"`
	Uuid                 string                    `protobuf:"bytes,3,opt,name=uuid" json:"uuid,omitempty"`
	Attributes           *LogicalVolume_Attributes `protobuf:"bytes,4,opt,name=attributes" json:"attributes,omitempty"`
	CopyPercent          string                    `protobuf:"bytes,5,opt,name=copy_percent,json=copyPercent" json:"copy_percent,omitempty"`
	ActualDevMajorNumber uint32                    `protobuf:"varint,6,opt,name=actual_dev_major_number,json=actualDevMajorNumber" json:"actual_dev_major_number,omitempty"`
	ActualDevMinorNumber uint32                    `protobuf:"varint,7,opt,name=actual_dev_minor_number,json=actualDevMinorNumber" json:"actual_dev_minor_number,omitempty"`
	Tags                 []string                  `protobuf:"bytes,8,rep,name=tags" json:"tags,omitempty"`
}

func (m *LogicalVolume) Reset()                    { *m = LogicalVolume{} }
func (m *LogicalVolume) String() string            { return proto.CompactTextString(m) }
func (*LogicalVolume) ProtoMessage()               {}
func (*LogicalVolume) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LogicalVolume) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LogicalVolume) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *LogicalVolume) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *LogicalVolume) GetAttributes() *LogicalVolume_Attributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *LogicalVolume) GetCopyPercent() string {
	if m != nil {
		return m.CopyPercent
	}
	return ""
}

func (m *LogicalVolume) GetActualDevMajorNumber() uint32 {
	if m != nil {
		return m.ActualDevMajorNumber
	}
	return 0
}

func (m *LogicalVolume) GetActualDevMinorNumber() uint32 {
	if m != nil {
		return m.ActualDevMinorNumber
	}
	return 0
}

func (m *LogicalVolume) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type LogicalVolume_Attributes struct {
	Type              LogicalVolume_Attributes_Type        `protobuf:"varint,1,opt,name=type,enum=lvm.LogicalVolume_Attributes_Type" json:"type,omitempty"`
	Permissions       LogicalVolume_Attributes_Permissions `protobuf:"varint,2,opt,name=permissions,enum=lvm.LogicalVolume_Attributes_Permissions" json:"permissions,omitempty"`
	Allocation        LogicalVolume_Attributes_Allocation  `protobuf:"varint,3,opt,name=allocation,enum=lvm.LogicalVolume_Attributes_Allocation" json:"allocation,omitempty"`
	FixedMinor        bool                                 `protobuf:"varint,4,opt,name=fixed_minor,json=fixedMinor" json:"fixed_minor,omitempty"`
	State             LogicalVolume_Attributes_State       `protobuf:"varint,5,opt,name=state,enum=lvm.LogicalVolume_Attributes_State" json:"state,omitempty"`
	Open              bool                                 `protobuf:"varint,6,opt,name=open" json:"open,omitempty"`
	TargetType        LogicalVolume_Attributes_TargetType  `protobuf:"varint,7,opt,name=target_type,json=targetType,enum=lvm.LogicalVolume_Attributes_TargetType" json:"target_type,omitempty"`
	Zeroing           bool                                 `protobuf:"varint,8,opt,name=zeroing" json:"zeroing,omitempty"`
	Health            LogicalVolume_Attributes_Health      `protobuf:"varint,9,opt,name=health,enum=lvm.LogicalVolume_Attributes_Health" json:"health,omitempty"`
	ActivationSkipped bool                                 `protobuf:"varint,10,opt,name=activation_skipped,json=activationSkipped" json:"activation_skipped,omitempty"`
}

func (m *LogicalVolume_Attributes) Reset()                    { *m = LogicalVolume_Attributes{} }
func (m *LogicalVolume_Attributes) String() string            { return proto.CompactTextString(m) }
func (*LogicalVolume_Attributes) ProtoMessage()               {}
func (*LogicalVolume_Attributes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *LogicalVolume_Attributes) GetType() LogicalVolume_Attributes_Type {
	if m != nil {
		return m.Type
	}
	return LogicalVolume_Attributes_MALFORMED_TYPE
}

func (m *LogicalVolume_Attributes) GetPermissions() LogicalVolume_Attributes_Permissions {
	if m != nil {
		return m.Permissions
	}
	return LogicalVolume_Attributes_MALFORMED_PERMISSIONS
}

func (m *LogicalVolume_Attributes) GetAllocation() LogicalVolume_Attributes_Allocation {
	if m != nil {
		return m.Allocation
	}
	return LogicalVolume_Attributes_MALFORMED_ALLOCATION
}

func (m *LogicalVolume_Attributes) GetFixedMinor() bool {
	if m != nil {
		return m.FixedMinor
	}
	return false
}

func (m *LogicalVolume_Attributes) GetState() LogicalVolume_Attributes_State {
	if m != nil {
		return m.State
	}
	return LogicalVolume_Attributes_MALFORMED_STATE
}

func (m *LogicalVolume_Attributes) GetOpen() bool {
	if m != nil {
		return m.Open
	}
	return false
}

func (m *LogicalVolume_Attributes) GetTargetType() LogicalVolume_Attributes_TargetType {
	if m != nil {
		return m.TargetType
	}
	return LogicalVolume_Attributes_MALFORMED_TARGET
}

func (m *LogicalVolume_Attributes) GetZeroing() bool {
	if m != nil {
		return m.Zeroing
	}
	return false
}

func (m *LogicalVolume_Attributes) GetHealth() LogicalVolume_Attributes_Health {
	if m != nil {
		return m.Health
	}
	return LogicalVolume_Attributes_MALFORMED_HEALTH
}

func (m *LogicalVolume_Attributes) GetActivationSkipped() bool {
	if m != nil {
		return m.ActivationSkipped
	}
	return false
}

type ListLVRequest struct {
	VolumeGroup string `protobuf:"bytes,1,opt,name=volume_group,json=volumeGroup" json:"volume_group,omitempty"`
}

func (m *ListLVRequest) Reset()                    { *m = ListLVRequest{} }
func (m *ListLVRequest) String() string            { return proto.CompactTextString(m) }
func (*ListLVRequest) ProtoMessage()               {}
func (*ListLVRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ListLVRequest) GetVolumeGroup() string {
	if m != nil {
		return m.VolumeGroup
	}
	return ""
}

type ListLVReply struct {
	Volumes []*LogicalVolume `protobuf:"bytes,1,rep,name=volumes" json:"volumes,omitempty"`
}

func (m *ListLVReply) Reset()                    { *m = ListLVReply{} }
func (m *ListLVReply) String() string            { return proto.CompactTextString(m) }
func (*ListLVReply) ProtoMessage()               {}
func (*ListLVReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ListLVReply) GetVolumes() []*LogicalVolume {
	if m != nil {
		return m.Volumes
	}
	return nil
}

type CreateLVRequest struct {
	VolumeGroup string   `protobuf:"bytes,1,opt,name=volume_group,json=volumeGroup" json:"volume_group,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Size        uint64   `protobuf:"varint,3,opt,name=size" json:"size,omitempty"`
	Mirrors     uint32   `protobuf:"varint,4,opt,name=mirrors" json:"mirrors,omitempty"`
	Tags        []string `protobuf:"bytes,5,rep,name=tags" json:"tags,omitempty"`
}

func (m *CreateLVRequest) Reset()                    { *m = CreateLVRequest{} }
func (m *CreateLVRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateLVRequest) ProtoMessage()               {}
func (*CreateLVRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CreateLVRequest) GetVolumeGroup() string {
	if m != nil {
		return m.VolumeGroup
	}
	return ""
}

func (m *CreateLVRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateLVRequest) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *CreateLVRequest) GetMirrors() uint32 {
	if m != nil {
		return m.Mirrors
	}
	return 0
}

func (m *CreateLVRequest) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type CreateLVReply struct {
	CommandOutput string `protobuf:"bytes,1,opt,name=command_output,json=commandOutput" json:"command_output,omitempty"`
}

func (m *CreateLVReply) Reset()                    { *m = CreateLVReply{} }
func (m *CreateLVReply) String() string            { return proto.CompactTextString(m) }
func (*CreateLVReply) ProtoMessage()               {}
func (*CreateLVReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CreateLVReply) GetCommandOutput() string {
	if m != nil {
		return m.CommandOutput
	}
	return ""
}

type RemoveLVRequest struct {
	VolumeGroup string `protobuf:"bytes,1,opt,name=volume_group,json=volumeGroup" json:"volume_group,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *RemoveLVRequest) Reset()                    { *m = RemoveLVRequest{} }
func (m *RemoveLVRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveLVRequest) ProtoMessage()               {}
func (*RemoveLVRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RemoveLVRequest) GetVolumeGroup() string {
	if m != nil {
		return m.VolumeGroup
	}
	return ""
}

func (m *RemoveLVRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type RemoveLVReply struct {
	CommandOutput string `protobuf:"bytes,1,opt,name=command_output,json=commandOutput" json:"command_output,omitempty"`
}

func (m *RemoveLVReply) Reset()                    { *m = RemoveLVReply{} }
func (m *RemoveLVReply) String() string            { return proto.CompactTextString(m) }
func (*RemoveLVReply) ProtoMessage()               {}
func (*RemoveLVReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *RemoveLVReply) GetCommandOutput() string {
	if m != nil {
		return m.CommandOutput
	}
	return ""
}

type CloneLVRequest struct {
	SourceName string `protobuf:"bytes,1,opt,name=source_name,json=sourceName" json:"source_name,omitempty"`
	DestName   string `protobuf:"bytes,2,opt,name=dest_name,json=destName" json:"dest_name,omitempty"`
}

func (m *CloneLVRequest) Reset()                    { *m = CloneLVRequest{} }
func (m *CloneLVRequest) String() string            { return proto.CompactTextString(m) }
func (*CloneLVRequest) ProtoMessage()               {}
func (*CloneLVRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *CloneLVRequest) GetSourceName() string {
	if m != nil {
		return m.SourceName
	}
	return ""
}

func (m *CloneLVRequest) GetDestName() string {
	if m != nil {
		return m.DestName
	}
	return ""
}

type CloneLVReply struct {
	CommandOutput string `protobuf:"bytes,1,opt,name=command_output,json=commandOutput" json:"command_output,omitempty"`
}

func (m *CloneLVReply) Reset()                    { *m = CloneLVReply{} }
func (m *CloneLVReply) String() string            { return proto.CompactTextString(m) }
func (*CloneLVReply) ProtoMessage()               {}
func (*CloneLVReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *CloneLVReply) GetCommandOutput() string {
	if m != nil {
		return m.CommandOutput
	}
	return ""
}

func init() {
	proto.RegisterType((*LogicalVolume)(nil), "lvm.LogicalVolume")
	proto.RegisterType((*LogicalVolume_Attributes)(nil), "lvm.LogicalVolume.Attributes")
	proto.RegisterType((*ListLVRequest)(nil), "lvm.ListLVRequest")
	proto.RegisterType((*ListLVReply)(nil), "lvm.ListLVReply")
	proto.RegisterType((*CreateLVRequest)(nil), "lvm.CreateLVRequest")
	proto.RegisterType((*CreateLVReply)(nil), "lvm.CreateLVReply")
	proto.RegisterType((*RemoveLVRequest)(nil), "lvm.RemoveLVRequest")
	proto.RegisterType((*RemoveLVReply)(nil), "lvm.RemoveLVReply")
	proto.RegisterType((*CloneLVRequest)(nil), "lvm.CloneLVRequest")
	proto.RegisterType((*CloneLVReply)(nil), "lvm.CloneLVReply")
	proto.RegisterEnum("lvm.LogicalVolume_Attributes_Type", LogicalVolume_Attributes_Type_name, LogicalVolume_Attributes_Type_value)
	proto.RegisterEnum("lvm.LogicalVolume_Attributes_Permissions", LogicalVolume_Attributes_Permissions_name, LogicalVolume_Attributes_Permissions_value)
	proto.RegisterEnum("lvm.LogicalVolume_Attributes_Allocation", LogicalVolume_Attributes_Allocation_name, LogicalVolume_Attributes_Allocation_value)
	proto.RegisterEnum("lvm.LogicalVolume_Attributes_State", LogicalVolume_Attributes_State_name, LogicalVolume_Attributes_State_value)
	proto.RegisterEnum("lvm.LogicalVolume_Attributes_TargetType", LogicalVolume_Attributes_TargetType_name, LogicalVolume_Attributes_TargetType_value)
	proto.RegisterEnum("lvm.LogicalVolume_Attributes_Health", LogicalVolume_Attributes_Health_name, LogicalVolume_Attributes_Health_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

func init() { proto.RegisterFile("lvm.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0x4d, 0x6f, 0xdb, 0x46,
	0x13, 0x36, 0xf5, 0xad, 0xa1, 0x25, 0xaf, 0x37, 0x4e, 0x5e, 0xc6, 0x6f, 0xdb, 0xb8, 0x4c, 0x0a,
	0x38, 0x40, 0x63, 0x14, 0x0a, 0x12, 0xa0, 0x68, 0x7b, 0x60, 0xc5, 0xb5, 0x44, 0x98, 0x1f, 0xc2,
	0x92, 0x96, 0xeb, 0xd3, 0x96, 0x91, 0x59, 0x87, 0xad, 0x24, 0xaa, 0x14, 0x25, 0xd4, 0xf9, 0x01,
	0x3d, 0xf4, 0x5c, 0xf4, 0x7f, 0xf5, 0xde, 0x63, 0x7f, 0x47, 0x51, 0xec, 0xf2, 0x4b, 0x72, 0x5a,
	0x23, 0x41, 0x6f, 0xbb, 0xcf, 0xcc, 0x33, 0xf3, 0xec, 0xcc, 0x2c, 0xb9, 0xd0, 0x9e, 0xae, 0x67,
	0x27, 0x8b, 0x38, 0x4a, 0x22, 0x5c, 0x9d, 0xae, 0x67, 0xea, 0x6f, 0x08, 0x3a, 0x66, 0x74, 0x1d,
	0x4e, 0xfc, 0xe9, 0x38, 0x9a, 0xae, 0x66, 0x01, 0xc6, 0x50, 0x9b, 0xfb, 0xb3, 0x40, 0x91, 0x8e,
	0xa4, 0xe3, 0x36, 0x15, 0x6b, 0x8e, 0x2d, 0xc3, 0x37, 0x81, 0x52, 0x39, 0x92, 0x8e, 0x6b, 0x54,
	0xac, 0x39, 0xb6, 0x5a, 0x85, 0x57, 0x4a, 0x35, 0xf5, 0xe3, 0x6b, 0xfc, 0x15, 0x80, 0x9f, 0x24,
	0x71, 0xf8, 0x6a, 0x95, 0x04, 0x4b, 0xa5, 0x76, 0x24, 0x1d, 0xcb, 0xbd, 0x0f, 0x4f, 0x78, 0xca,
	0xad, 0x1c, 0x27, 0x5a, 0xe1, 0x44, 0x37, 0x08, 0xf8, 0x63, 0xd8, 0x9d, 0x44, 0x8b, 0x1b, 0xb6,
	0x08, 0xe2, 0x49, 0x30, 0x4f, 0x94, 0xba, 0x08, 0x2d, 0x73, 0x6c, 0x94, 0x42, 0xf8, 0x05, 0xfc,
	0xcf, 0x9f, 0x24, 0x2b, 0x7f, 0xca, 0xae, 0x82, 0x35, 0x9b, 0xf9, 0xdf, 0x47, 0x31, 0x9b, 0xaf,
	0x66, 0xaf, 0x82, 0x58, 0x69, 0x1c, 0x49, 0xc7, 0x1d, 0x7a, 0x90, 0x9a, 0xf5, 0x60, 0x6d, 0x71,
	0xa3, 0x2d, 0x6c, 0xb7, 0x69, 0xe1, 0xbc, 0xa4, 0x35, 0x6f, 0xd3, 0xb8, 0x31, 0xa3, 0x61, 0xa8,
	0x25, 0xfe, 0xf5, 0x52, 0x69, 0x1d, 0x55, 0xf9, 0x19, 0xf9, 0xfa, 0xf0, 0xcf, 0x0e, 0x40, 0xa9,
	0x1f, 0xbf, 0x84, 0x5a, 0x72, 0xb3, 0x48, 0xcb, 0xd5, 0xed, 0xa9, 0x77, 0x1e, 0xf6, 0xc4, 0xbb,
	0x59, 0x04, 0x54, 0xf8, 0xe3, 0x33, 0x90, 0x17, 0x41, 0x3c, 0x0b, 0x97, 0xcb, 0x30, 0x9a, 0x2f,
	0x45, 0x65, 0xbb, 0xbd, 0xa7, 0x77, 0xd3, 0x47, 0x25, 0x81, 0x6e, 0xb2, 0xf1, 0x10, 0xc0, 0x9f,
	0x4e, 0xa3, 0x89, 0x9f, 0x84, 0xd1, 0x5c, 0x74, 0xa4, 0xdb, 0x3b, 0xbe, 0x3b, 0x96, 0x56, 0xf8,
	0xd3, 0x0d, 0x2e, 0x7e, 0x04, 0xf2, 0x77, 0xe1, 0x4f, 0xc1, 0x55, 0x5a, 0x23, 0xd1, 0xc2, 0x16,
	0x05, 0x01, 0x89, 0xc2, 0xe0, 0xcf, 0xa1, 0xbe, 0x4c, 0xfc, 0x24, 0x10, 0xcd, 0xe9, 0xf6, 0x1e,
	0xdf, 0x9d, 0xc5, 0xe5, 0xae, 0x34, 0x65, 0xf0, 0x6a, 0x46, 0x8b, 0x60, 0x2e, 0x1a, 0xd5, 0xa2,
	0x62, 0x8d, 0x0d, 0x90, 0x13, 0x3f, 0xbe, 0x0e, 0x12, 0x26, 0xaa, 0xd8, 0x7c, 0x17, 0xe9, 0x9e,
	0x20, 0x88, 0x5a, 0x42, 0x52, 0xac, 0xb1, 0x02, 0xcd, 0x37, 0x41, 0x1c, 0x85, 0xf3, 0x6b, 0xa5,
	0x25, 0x32, 0xe4, 0x5b, 0xfc, 0x25, 0x34, 0x5e, 0x07, 0xfe, 0x34, 0x79, 0xad, 0xb4, 0x45, 0xfc,
	0x27, 0x77, 0xc7, 0x1f, 0x0a, 0x5f, 0x9a, 0x71, 0xf0, 0x33, 0xc0, 0xfe, 0x24, 0x09, 0xd7, 0xa2,
	0x40, 0x6c, 0xf9, 0x43, 0xb8, 0x58, 0x04, 0x57, 0x0a, 0x88, 0x14, 0xfb, 0xa5, 0xc5, 0x4d, 0x0d,
	0xea, 0x5f, 0x15, 0xa8, 0x09, 0x3d, 0x18, 0xba, 0x96, 0x66, 0x9e, 0x3a, 0xd4, 0x22, 0x3a, 0xf3,
	0x2e, 0x47, 0x04, 0xed, 0xe0, 0x5d, 0x68, 0x59, 0x06, 0xa5, 0x0e, 0x25, 0x3a, 0x92, 0xf0, 0x43,
	0xb8, 0x9f, 0xef, 0xd8, 0x85, 0xe1, 0x0d, 0x9d, 0x73, 0x8f, 0xb9, 0x97, 0x76, 0x1f, 0x55, 0x30,
	0x40, 0xc3, 0xa1, 0xc6, 0xc0, 0xb0, 0x51, 0x15, 0x1f, 0xc1, 0x07, 0xe9, 0x5a, 0x38, 0x31, 0x8b,
	0xd0, 0x81, 0x61, 0x0f, 0x98, 0x6b, 0x6b, 0x23, 0x77, 0xe8, 0x78, 0xa8, 0x86, 0x5b, 0x50, 0xa3,
	0x9a, 0xa1, 0xa3, 0x3a, 0xbe, 0x0f, 0xfb, 0x7c, 0xb5, 0x1d, 0xae, 0xc1, 0xf3, 0x16, 0xee, 0x4d,
	0x7c, 0x00, 0xe8, 0xad, 0x20, 0x2d, 0x2c, 0x43, 0x73, 0x34, 0x66, 0x96, 0x33, 0x26, 0xa8, 0xcd,
	0xc5, 0x8f, 0x0d, 0xea, 0x9d, 0x6b, 0x26, 0x4b, 0x25, 0x22, 0xc0, 0x0f, 0x00, 0xe7, 0x98, 0xc8,
	0x61, 0x58, 0xda, 0x80, 0x20, 0x19, 0x1f, 0xc2, 0x83, 0x72, 0xcf, 0x78, 0x56, 0xe7, 0x34, 0x4d,
	0xbc, 0x8b, 0xbb, 0x00, 0x29, 0x9f, 0x99, 0xce, 0x00, 0x75, 0x78, 0xea, 0x73, 0x5b, 0x27, 0x94,
	0xf5, 0x1d, 0x7b, 0x4c, 0xa8, 0x6b, 0x38, 0x36, 0xea, 0x72, 0xfd, 0xde, 0xd0, 0xb0, 0xd1, 0x1e,
	0xee, 0x40, 0x9b, 0xaf, 0xd8, 0xc8, 0x71, 0x4c, 0x84, 0xb8, 0x8c, 0x62, 0xcb, 0x74, 0xcd, 0xd3,
	0xd0, 0x3e, 0xfe, 0x08, 0x0e, 0x45, 0x3a, 0x87, 0xb2, 0xd2, 0x66, 0x11, 0x4f, 0x13, 0x76, 0xac,
	0x7e, 0x0b, 0xf2, 0xc6, 0x45, 0x11, 0x45, 0x2e, 0xda, 0x30, 0x22, 0xd4, 0x32, 0x5c, 0x9e, 0xd5,
	0x45, 0x3b, 0x3c, 0xd9, 0x05, 0x35, 0x3c, 0xa2, 0x7d, 0x6d, 0x12, 0x24, 0xf1, 0x2d, 0x25, 0x9a,
	0xce, 0x1c, 0xdb, 0xbc, 0x44, 0x15, 0xac, 0xc0, 0x41, 0xb1, 0x65, 0x5a, 0xdf, 0x33, 0xc6, 0x9a,
	0xc7, 0xe5, 0x56, 0xd5, 0xdf, 0x25, 0x80, 0xf2, 0xfe, 0x70, 0xc7, 0x32, 0x83, 0x66, 0x9a, 0x4e,
	0x3f, 0x75, 0x14, 0xed, 0xd6, 0xec, 0xcb, 0x8b, 0x21, 0xa1, 0x3c, 0x7e, 0x17, 0xa0, 0xef, 0xd8,
	0x9e, 0x31, 0x38, 0x77, 0xce, 0x5d, 0x54, 0xe1, 0xf9, 0x0c, 0x7b, 0x48, 0xb8, 0x02, 0x1d, 0x55,
	0x71, 0x1b, 0xea, 0x7d, 0xd3, 0xb0, 0x07, 0xa8, 0xc6, 0xbb, 0x6f, 0x3b, 0xd4, 0xd2, 0x4c, 0x54,
	0xc7, 0xf7, 0x60, 0x2f, 0x8f, 0xc1, 0x4c, 0xa7, 0x7f, 0x46, 0x74, 0xd4, 0xe0, 0x6d, 0x2e, 0x43,
	0xe5, 0xb0, 0x68, 0x6c, 0x11, 0x31, 0x47, 0x5b, 0x18, 0xc1, 0xae, 0x08, 0x9c, 0x23, 0x6d, 0xbc,
	0x0f, 0x9d, 0x34, 0x7e, 0x0e, 0x81, 0xfa, 0x73, 0x05, 0xea, 0xe2, 0xb6, 0xf2, 0x84, 0xe5, 0x71,
	0x5c, 0x4f, 0xf3, 0xf8, 0xe0, 0x02, 0x34, 0x44, 0x09, 0xb2, 0x3a, 0xb9, 0xe7, 0xee, 0x88, 0xd8,
	0x3a, 0xd1, 0x51, 0x25, 0x4d, 0x3a, 0xd6, 0x4c, 0x43, 0x2f, 0xa7, 0xa9, 0xca, 0xbb, 0x54, 0xa0,
	0xb9, 0xf3, 0xe6, 0xc8, 0x3e, 0x84, 0xfb, 0xf9, 0x4e, 0x4c, 0x34, 0x61, 0xa7, 0x9a, 0x61, 0x12,
	0x3e, 0xc3, 0x8f, 0xe1, 0xd1, 0xdb, 0x94, 0x6d, 0xa7, 0x06, 0x3e, 0x86, 0x27, 0x96, 0x36, 0x1a,
	0x11, 0x9d, 0xe9, 0x64, 0x6c, 0xf4, 0x09, 0x1b, 0x51, 0xe2, 0x12, 0xdb, 0x2b, 0x26, 0xdf, 0xe3,
	0x5d, 0x75, 0x51, 0x13, 0x3f, 0x83, 0xa7, 0xff, 0xee, 0xc9, 0x0c, 0x3b, 0x3d, 0x57, 0xea, 0x8f,
	0x5a, 0xea, 0xaf, 0x12, 0x40, 0xf9, 0x85, 0x11, 0x77, 0xa5, 0xbc, 0xc5, 0x1a, 0x1d, 0x10, 0x0f,
	0xed, 0xf0, 0x02, 0x66, 0x63, 0x9d, 0x41, 0x12, 0xde, 0x03, 0x59, 0x8c, 0x65, 0x06, 0x54, 0x78,
	0x1d, 0x0b, 0xf1, 0x19, 0x58, 0xe5, 0x5e, 0x62, 0x68, 0x33, 0xa0, 0xc6, 0x27, 0xfc, 0xdc, 0x3e,
	0xb3, 0x9d, 0x8b, 0x02, 0xab, 0x6f, 0x5e, 0xbe, 0x0c, 0x6b, 0xa8, 0x73, 0x68, 0xa4, 0xdf, 0xa5,
	0x6d, 0x45, 0x43, 0xa2, 0x99, 0xde, 0x10, 0xed, 0xe0, 0x06, 0x54, 0x9c, 0x33, 0x24, 0x89, 0x5b,
	0xac, 0x51, 0xcf, 0xd0, 0x4c, 0x54, 0xe1, 0x81, 0x28, 0x39, 0xa5, 0xc4, 0x1d, 0x32, 0x9b, 0x10,
	0x5d, 0x8c, 0x19, 0xa7, 0x1b, 0xae, 0xa5, 0x79, 0xfd, 0x21, 0x71, 0x19, 0xf9, 0xc6, 0x70, 0xb9,
	0x8c, 0x3d, 0x90, 0xc5, 0x55, 0xb0, 0x1c, 0xd7, 0x33, 0x2f, 0x51, 0x5d, 0xed, 0x41, 0xc7, 0x0c,
	0x97, 0x89, 0x39, 0xa6, 0xc1, 0x8f, 0xab, 0x60, 0x99, 0xf0, 0x9f, 0xf3, 0x5a, 0x7c, 0x2a, 0xd9,
	0x75, 0x1c, 0xad, 0x16, 0xd9, 0xfb, 0x40, 0x4e, 0xb1, 0x01, 0x87, 0xd4, 0x2f, 0x40, 0xce, 0x39,
	0x8b, 0xe9, 0x0d, 0xfe, 0x14, 0x9a, 0xa9, 0x75, 0xa9, 0x48, 0x47, 0xd5, 0x63, 0xb9, 0x87, 0xdf,
	0xfe, 0xee, 0xd2, 0xdc, 0x45, 0xfd, 0x45, 0x82, 0xbd, 0x7e, 0x1c, 0xf8, 0x49, 0xf0, 0x3e, 0x39,
	0x8b, 0xe7, 0x4a, 0xe5, 0x1f, 0x9e, 0x2b, 0xd5, 0x8d, 0xe7, 0x8a, 0x02, 0xcd, 0x59, 0x18, 0xc7,
	0x51, 0x9c, 0xbe, 0x4b, 0x3a, 0x34, 0xdf, 0x16, 0x3f, 0xf9, 0x7a, 0xf9, 0x93, 0x57, 0x5f, 0x42,
	0xa7, 0xd4, 0xc2, 0xcf, 0xf2, 0x09, 0x74, 0x27, 0xd1, 0x6c, 0xe6, 0xcf, 0xaf, 0x58, 0xb4, 0x4a,
	0x16, 0xab, 0x24, 0xd3, 0xd2, 0xc9, 0x50, 0x47, 0x80, 0xea, 0x10, 0xf6, 0x68, 0x30, 0x8b, 0xd6,
	0xff, 0xf9, 0x0c, 0x5c, 0x41, 0x19, 0xe9, 0x3d, 0x14, 0xd8, 0xd0, 0xed, 0x4f, 0xa3, 0xf9, 0x86,
	0x80, 0x47, 0x20, 0x2f, 0xa3, 0x55, 0x3c, 0x09, 0xd8, 0xc6, 0xbb, 0x0e, 0x52, 0xc8, 0xe6, 0xe5,
	0xfa, 0x3f, 0xb4, 0xaf, 0x82, 0x65, 0xc2, 0x36, 0x34, 0xb4, 0x38, 0xc0, 0x8d, 0xea, 0x0b, 0xd8,
	0x2d, 0xe2, 0xbd, 0xbb, 0x8c, 0xde, 0x1f, 0x12, 0x54, 0xcd, 0xb1, 0x85, 0x3f, 0x83, 0x46, 0x3a,
	0x12, 0x38, 0x6b, 0xfe, 0xe6, 0x4c, 0x1d, 0xa2, 0x2d, 0x6c, 0x31, 0xbd, 0x51, 0x77, 0xf0, 0x4b,
	0x68, 0xe5, 0xa5, 0xc7, 0x07, 0xc2, 0x7e, 0x6b, 0x2a, 0x0e, 0xf1, 0x2d, 0xb4, 0xe0, 0xe5, 0x05,
	0xcb, 0x78, 0xb7, 0x3a, 0x91, 0xf1, 0xb6, 0xaa, 0xaa, 0xee, 0xe0, 0xe7, 0xd0, 0xcc, 0x0e, 0x88,
	0xef, 0xa5, 0x81, 0xb7, 0xca, 0x77, 0xb8, 0xbf, 0x0d, 0x0a, 0xd2, 0xab, 0x86, 0x78, 0x42, 0x3f,
	0xff, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x17, 0x21, 0xdd, 0x7b, 0x4f, 0x0b, 0x00, 0x00,
}
