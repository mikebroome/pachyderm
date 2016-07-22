// Code generated by protoc-gen-go.
// source: server/pfs/db/persist/persist.proto
// DO NOT EDIT!

/*
Package persist is a generated protocol buffer package.

It is generated from these files:
	server/pfs/db/persist/persist.proto

It has these top-level messages:
	Clock
	ClockID
	BranchClock
	Repo
	Branch
	Diff
	Commit
*/
package persist

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "go.pedge.io/pb/go/google/protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Clock struct {
	// a document either has these two fields
	Branch string `protobuf:"bytes,1,opt,name=branch" json:"branch,omitempty"`
	Clock  uint64 `protobuf:"varint,2,opt,name=clock" json:"clock,omitempty"`
}

func (m *Clock) Reset()                    { *m = Clock{} }
func (m *Clock) String() string            { return proto.CompactTextString(m) }
func (*Clock) ProtoMessage()               {}
func (*Clock) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ClockID struct {
	ID     string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Repo   string `protobuf:"bytes,2,opt,name=repo" json:"repo,omitempty"`
	Branch string `protobuf:"bytes,3,opt,name=branch" json:"branch,omitempty"`
	Clock  uint64 `protobuf:"varint,4,opt,name=clock" json:"clock,omitempty"`
}

func (m *ClockID) Reset()                    { *m = ClockID{} }
func (m *ClockID) String() string            { return proto.CompactTextString(m) }
func (*ClockID) ProtoMessage()               {}
func (*ClockID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type BranchClock struct {
	Clocks []*Clock `protobuf:"bytes,1,rep,name=clocks" json:"clocks,omitempty"`
}

func (m *BranchClock) Reset()                    { *m = BranchClock{} }
func (m *BranchClock) String() string            { return proto.CompactTextString(m) }
func (*BranchClock) ProtoMessage()               {}
func (*BranchClock) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *BranchClock) GetClocks() []*Clock {
	if m != nil {
		return m.Clocks
	}
	return nil
}

type Repo struct {
	Name    string                     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Created *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=created" json:"created,omitempty"`
}

func (m *Repo) Reset()                    { *m = Repo{} }
func (m *Repo) String() string            { return proto.CompactTextString(m) }
func (*Repo) ProtoMessage()               {}
func (*Repo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Repo) GetCreated() *google_protobuf.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

type Branch struct {
	ID   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Repo string `protobuf:"bytes,2,opt,name=repo" json:"repo,omitempty"`
	Name string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
}

func (m *Branch) Reset()                    { *m = Branch{} }
func (m *Branch) String() string            { return proto.CompactTextString(m) }
func (*Branch) ProtoMessage()               {}
func (*Branch) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type Diff struct {
	ID       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	CommitID string `protobuf:"bytes,2,opt,name=commit_id,json=commitId" json:"commit_id,omitempty"`
	Path     string `protobuf:"bytes,3,opt,name=path" json:"path,omitempty"`
	// appends can either be blockrefs, or a special
	// constant "delete", which means deleting the file
	BlockRefs []string `protobuf:"bytes,4,rep,name=block_refs,json=blockRefs" json:"block_refs,omitempty"`
	Delete    bool     `protobuf:"varint,5,opt,name=delete" json:"delete,omitempty"`
	Size      uint64   `protobuf:"varint,6,opt,name=size" json:"size,omitempty"`
}

func (m *Diff) Reset()                    { *m = Diff{} }
func (m *Diff) String() string            { return proto.CompactTextString(m) }
func (*Diff) ProtoMessage()               {}
func (*Diff) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type Commit struct {
	ID           string         `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Repo         string         `protobuf:"bytes,2,opt,name=repo" json:"repo,omitempty"`
	BranchClocks []*BranchClock `protobuf:"bytes,3,rep,name=branch_clocks,json=branchClocks" json:"branch_clocks,omitempty"`
	// true means append
	// false means removal
	Modifications map[string]bool            `protobuf:"bytes,4,rep,name=modifications" json:"modifications,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	Started       *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=started" json:"started,omitempty"`
	Finished      *google_protobuf.Timestamp `protobuf:"bytes,6,opt,name=finished" json:"finished,omitempty"`
	Cancelled     bool                       `protobuf:"varint,7,opt,name=cancelled" json:"cancelled,omitempty"`
	Provenance    []string                   `protobuf:"bytes,8,rep,name=provenance" json:"provenance,omitempty"`
}

func (m *Commit) Reset()                    { *m = Commit{} }
func (m *Commit) String() string            { return proto.CompactTextString(m) }
func (*Commit) ProtoMessage()               {}
func (*Commit) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Commit) GetBranchClocks() []*BranchClock {
	if m != nil {
		return m.BranchClocks
	}
	return nil
}

func (m *Commit) GetModifications() map[string]bool {
	if m != nil {
		return m.Modifications
	}
	return nil
}

func (m *Commit) GetStarted() *google_protobuf.Timestamp {
	if m != nil {
		return m.Started
	}
	return nil
}

func (m *Commit) GetFinished() *google_protobuf.Timestamp {
	if m != nil {
		return m.Finished
	}
	return nil
}

func init() {
	proto.RegisterType((*Clock)(nil), "Clock")
	proto.RegisterType((*ClockID)(nil), "ClockID")
	proto.RegisterType((*BranchClock)(nil), "BranchClock")
	proto.RegisterType((*Repo)(nil), "Repo")
	proto.RegisterType((*Branch)(nil), "Branch")
	proto.RegisterType((*Diff)(nil), "Diff")
	proto.RegisterType((*Commit)(nil), "Commit")
}

func init() { proto.RegisterFile("server/pfs/db/persist/persist.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 474 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x52, 0xc9, 0x6a, 0xdc, 0x40,
	0x10, 0x65, 0x2c, 0x8d, 0x46, 0x53, 0x63, 0x87, 0xd0, 0x84, 0x20, 0x26, 0x2b, 0xca, 0xc5, 0x97,
	0x48, 0xc4, 0x59, 0x08, 0x39, 0x0d, 0xb1, 0x73, 0xf0, 0x21, 0x10, 0x9a, 0xdc, 0x72, 0x18, 0xb4,
	0x94, 0xec, 0xc6, 0xda, 0xe8, 0x6e, 0x0f, 0x38, 0x3f, 0x11, 0xc8, 0x17, 0xa7, 0xbb, 0x5a, 0xca,
	0x28, 0x38, 0x60, 0x9f, 0x54, 0xf5, 0xba, 0x5e, 0x2d, 0xef, 0x09, 0x5e, 0x29, 0x94, 0x3b, 0x94,
	0x69, 0x5f, 0xa9, 0xb4, 0xcc, 0xd3, 0x1e, 0xa5, 0x12, 0x4a, 0x8f, 0xdf, 0xa4, 0x97, 0x9d, 0xee,
	0xd6, 0x2f, 0x2e, 0xba, 0xee, 0xa2, 0xc6, 0x94, 0xb2, 0xfc, 0xba, 0x4a, 0xb5, 0x68, 0x50, 0xe9,
	0xac, 0xe9, 0x5d, 0x41, 0xfc, 0x1e, 0xe6, 0xa7, 0x75, 0x57, 0x5c, 0xb1, 0xc7, 0x10, 0xe4, 0x32,
	0x6b, 0x8b, 0xcb, 0x68, 0xf6, 0x72, 0x76, 0xbc, 0xe4, 0x43, 0xc6, 0x1e, 0xc1, 0xbc, 0xb0, 0x05,
	0xd1, 0x81, 0x81, 0x7d, 0xee, 0x92, 0xf8, 0x07, 0x2c, 0x88, 0x76, 0x7e, 0xc6, 0x1e, 0xc0, 0x81,
	0x28, 0x07, 0x92, 0x89, 0x18, 0x03, 0x5f, 0x62, 0xdf, 0x51, 0xfd, 0x92, 0x53, 0x3c, 0x69, 0xee,
	0xfd, 0xbf, 0xb9, 0x3f, 0x6d, 0xfe, 0x1a, 0x56, 0x9f, 0xe9, 0xdd, 0x6d, 0xf6, 0x1c, 0x02, 0xc2,
	0x95, 0x19, 0xe2, 0x1d, 0xaf, 0x4e, 0x82, 0x84, 0x70, 0x3e, 0xa0, 0xf1, 0x37, 0xf0, 0xb9, 0x1d,
	0x62, 0x06, 0xb7, 0x59, 0x83, 0xc3, 0x2a, 0x14, 0xb3, 0x77, 0xb0, 0x28, 0x24, 0x66, 0x1a, 0x4b,
	0xda, 0x67, 0x75, 0xb2, 0x4e, 0x9c, 0x22, 0xc9, 0xa8, 0x48, 0xf2, 0x7d, 0x54, 0x84, 0x8f, 0xa5,
	0xf1, 0x06, 0x02, 0xb7, 0xc0, 0xbd, 0x8e, 0x1b, 0xe7, 0x7a, 0xfb, 0xb9, 0xf1, 0xef, 0x19, 0xf8,
	0x67, 0xa2, 0xaa, 0x6e, 0x35, 0x78, 0x02, 0xcb, 0xa2, 0x6b, 0x1a, 0xa1, 0xb7, 0xa2, 0x1c, 0xba,
	0x84, 0x0e, 0x38, 0xa7, 0xee, 0x7d, 0xa6, 0x47, 0x91, 0x28, 0x66, 0xcf, 0x00, 0x72, 0x7b, 0xe7,
	0x56, 0x62, 0xa5, 0x8c, 0x4e, 0x9e, 0x79, 0x59, 0x12, 0xc2, 0x0d, 0x60, 0x95, 0x2d, 0xb1, 0x46,
	0x8d, 0xd1, 0xdc, 0x90, 0x42, 0x3e, 0x64, 0xb6, 0x95, 0x12, 0x3f, 0x31, 0x0a, 0x48, 0x58, 0x8a,
	0xe3, 0x5f, 0x1e, 0x04, 0xa7, 0x34, 0xeb, 0x5e, 0x77, 0xbd, 0x81, 0x23, 0x67, 0xd3, 0x76, 0x90,
	0xdf, 0x23, 0xf9, 0x0f, 0x93, 0x89, 0x39, 0xfc, 0x30, 0xdf, 0x27, 0x8a, 0x6d, 0xe0, 0xa8, 0xe9,
	0x4a, 0x51, 0x89, 0x22, 0xd3, 0xa2, 0x6b, 0xdd, 0xbe, 0x56, 0x74, 0x37, 0x36, 0xf9, 0x3a, 0x7d,
	0xfc, 0xd2, 0x6a, 0x79, 0xc3, 0xff, 0x25, 0x58, 0xc3, 0x8c, 0x19, 0xd2, 0x1a, 0x36, 0xbf, 0xdb,
	0xb0, 0xa1, 0x94, 0x7d, 0x80, 0xb0, 0x12, 0xad, 0x50, 0x97, 0x86, 0x16, 0xdc, 0x49, 0xfb, 0x5b,
	0xcb, 0x9e, 0x1a, 0x37, 0xcc, 0xfa, 0x58, 0xd7, 0x86, 0xb8, 0x20, 0x01, 0xf7, 0x80, 0xf9, 0xf1,
	0xc0, 0xb0, 0x77, 0xd8, 0x5a, 0x24, 0x0a, 0x49, 0xfa, 0x09, 0xb2, 0xde, 0x00, 0xbb, 0x7d, 0x10,
	0x7b, 0x08, 0xde, 0x15, 0xde, 0x0c, 0xda, 0xda, 0xd0, 0xfe, 0xe5, 0xbb, 0xac, 0xbe, 0x46, 0x52,
	0x37, 0xe4, 0x2e, 0xf9, 0x74, 0xf0, 0x71, 0x96, 0x07, 0xb4, 0xdd, 0xdb, 0x3f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xff, 0x04, 0x91, 0x2c, 0xcc, 0x03, 0x00, 0x00,
}
