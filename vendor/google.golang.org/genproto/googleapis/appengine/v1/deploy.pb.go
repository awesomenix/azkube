// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/appengine/v1/deploy.proto

package appengine // import "google.golang.org/genproto/googleapis/appengine/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Code and application artifacts used to deploy a version to App Engine.
type Deployment struct {
	// Manifest of the files stored in Google Cloud Storage that are included
	// as part of this version. All files must be readable using the
	// credentials supplied with this call.
	Files map[string]*FileInfo `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// A Docker image that App Engine uses to run the version.
	// Only applicable for instances in App Engine flexible environment.
	Container *ContainerInfo `protobuf:"bytes,2,opt,name=container,proto3" json:"container,omitempty"`
	// The zip file for this deployment, if this is a zip deployment.
	Zip                  *ZipInfo `protobuf:"bytes,3,opt,name=zip,proto3" json:"zip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Deployment) Reset()         { *m = Deployment{} }
func (m *Deployment) String() string { return proto.CompactTextString(m) }
func (*Deployment) ProtoMessage()    {}
func (*Deployment) Descriptor() ([]byte, []int) {
	return fileDescriptor_deploy_ed99bebbe4262e1c, []int{0}
}
func (m *Deployment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Deployment.Unmarshal(m, b)
}
func (m *Deployment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Deployment.Marshal(b, m, deterministic)
}
func (dst *Deployment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Deployment.Merge(dst, src)
}
func (m *Deployment) XXX_Size() int {
	return xxx_messageInfo_Deployment.Size(m)
}
func (m *Deployment) XXX_DiscardUnknown() {
	xxx_messageInfo_Deployment.DiscardUnknown(m)
}

var xxx_messageInfo_Deployment proto.InternalMessageInfo

func (m *Deployment) GetFiles() map[string]*FileInfo {
	if m != nil {
		return m.Files
	}
	return nil
}

func (m *Deployment) GetContainer() *ContainerInfo {
	if m != nil {
		return m.Container
	}
	return nil
}

func (m *Deployment) GetZip() *ZipInfo {
	if m != nil {
		return m.Zip
	}
	return nil
}

// Single source file that is part of the version to be deployed. Each source
// file that is deployed must be specified separately.
type FileInfo struct {
	// URL source to use to fetch this file. Must be a URL to a resource in
	// Google Cloud Storage in the form
	// 'http(s)://storage.googleapis.com/\<bucket\>/\<object\>'.
	SourceUrl string `protobuf:"bytes,1,opt,name=source_url,json=sourceUrl,proto3" json:"source_url,omitempty"`
	// The SHA1 hash of the file, in hex.
	Sha1Sum string `protobuf:"bytes,2,opt,name=sha1_sum,json=sha1Sum,proto3" json:"sha1_sum,omitempty"`
	// The MIME type of the file.
	//
	// Defaults to the value from Google Cloud Storage.
	MimeType             string   `protobuf:"bytes,3,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileInfo) Reset()         { *m = FileInfo{} }
func (m *FileInfo) String() string { return proto.CompactTextString(m) }
func (*FileInfo) ProtoMessage()    {}
func (*FileInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_deploy_ed99bebbe4262e1c, []int{1}
}
func (m *FileInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileInfo.Unmarshal(m, b)
}
func (m *FileInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileInfo.Marshal(b, m, deterministic)
}
func (dst *FileInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileInfo.Merge(dst, src)
}
func (m *FileInfo) XXX_Size() int {
	return xxx_messageInfo_FileInfo.Size(m)
}
func (m *FileInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FileInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FileInfo proto.InternalMessageInfo

func (m *FileInfo) GetSourceUrl() string {
	if m != nil {
		return m.SourceUrl
	}
	return ""
}

func (m *FileInfo) GetSha1Sum() string {
	if m != nil {
		return m.Sha1Sum
	}
	return ""
}

func (m *FileInfo) GetMimeType() string {
	if m != nil {
		return m.MimeType
	}
	return ""
}

// Docker image that is used to start a VM container for the version you
// deploy.
type ContainerInfo struct {
	// URI to the hosted container image in a Docker repository. The URI must be
	// fully qualified and include a tag or digest.
	// Examples: "gcr.io/my-project/image:tag" or "gcr.io/my-project/image@digest"
	Image                string   `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContainerInfo) Reset()         { *m = ContainerInfo{} }
func (m *ContainerInfo) String() string { return proto.CompactTextString(m) }
func (*ContainerInfo) ProtoMessage()    {}
func (*ContainerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_deploy_ed99bebbe4262e1c, []int{2}
}
func (m *ContainerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainerInfo.Unmarshal(m, b)
}
func (m *ContainerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainerInfo.Marshal(b, m, deterministic)
}
func (dst *ContainerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerInfo.Merge(dst, src)
}
func (m *ContainerInfo) XXX_Size() int {
	return xxx_messageInfo_ContainerInfo.Size(m)
}
func (m *ContainerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerInfo proto.InternalMessageInfo

func (m *ContainerInfo) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

type ZipInfo struct {
	// URL of the zip file to deploy from. Must be a URL to a resource in
	// Google Cloud Storage in the form
	// 'http(s)://storage.googleapis.com/\<bucket\>/\<object\>'.
	SourceUrl string `protobuf:"bytes,3,opt,name=source_url,json=sourceUrl,proto3" json:"source_url,omitempty"`
	// An estimate of the number of files in a zip for a zip deployment.
	// If set, must be greater than or equal to the actual number of files.
	// Used for optimizing performance; if not provided, deployment may be slow.
	FilesCount           int32    `protobuf:"varint,4,opt,name=files_count,json=filesCount,proto3" json:"files_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ZipInfo) Reset()         { *m = ZipInfo{} }
func (m *ZipInfo) String() string { return proto.CompactTextString(m) }
func (*ZipInfo) ProtoMessage()    {}
func (*ZipInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_deploy_ed99bebbe4262e1c, []int{3}
}
func (m *ZipInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ZipInfo.Unmarshal(m, b)
}
func (m *ZipInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ZipInfo.Marshal(b, m, deterministic)
}
func (dst *ZipInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ZipInfo.Merge(dst, src)
}
func (m *ZipInfo) XXX_Size() int {
	return xxx_messageInfo_ZipInfo.Size(m)
}
func (m *ZipInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ZipInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ZipInfo proto.InternalMessageInfo

func (m *ZipInfo) GetSourceUrl() string {
	if m != nil {
		return m.SourceUrl
	}
	return ""
}

func (m *ZipInfo) GetFilesCount() int32 {
	if m != nil {
		return m.FilesCount
	}
	return 0
}

func init() {
	proto.RegisterType((*Deployment)(nil), "google.appengine.v1.Deployment")
	proto.RegisterMapType((map[string]*FileInfo)(nil), "google.appengine.v1.Deployment.FilesEntry")
	proto.RegisterType((*FileInfo)(nil), "google.appengine.v1.FileInfo")
	proto.RegisterType((*ContainerInfo)(nil), "google.appengine.v1.ContainerInfo")
	proto.RegisterType((*ZipInfo)(nil), "google.appengine.v1.ZipInfo")
}

func init() {
	proto.RegisterFile("google/appengine/v1/deploy.proto", fileDescriptor_deploy_ed99bebbe4262e1c)
}

var fileDescriptor_deploy_ed99bebbe4262e1c = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xd1, 0xab, 0xd3, 0x30,
	0x14, 0xc6, 0xe9, 0x6a, 0xbd, 0xeb, 0x29, 0x82, 0x44, 0xc1, 0x7a, 0xbd, 0x17, 0x4b, 0x41, 0x28,
	0x3e, 0xa4, 0xec, 0xde, 0x17, 0x51, 0x1f, 0x2e, 0x9b, 0x0a, 0x7b, 0x1b, 0x55, 0x11, 0xf6, 0x52,
	0x62, 0xcd, 0x62, 0xb0, 0x4d, 0x42, 0x9b, 0x0e, 0xea, 0x7f, 0xe2, 0x7f, 0x2b, 0x49, 0xba, 0x8d,
	0x8d, 0xbe, 0xf5, 0x7c, 0xfd, 0x7d, 0x5f, 0x4e, 0x72, 0x0e, 0x24, 0x4c, 0x4a, 0x56, 0xd3, 0x9c,
	0x28, 0x45, 0x05, 0xe3, 0x82, 0xe6, 0xfb, 0x45, 0xfe, 0x8b, 0xaa, 0x5a, 0x0e, 0x58, 0xb5, 0x52,
	0x4b, 0xf4, 0xcc, 0x11, 0xf8, 0x48, 0xe0, 0xfd, 0xe2, 0xfa, 0xe6, 0x68, 0xe3, 0x39, 0x11, 0x42,
	0x6a, 0xa2, 0xb9, 0x14, 0x9d, 0xb3, 0xa4, 0xff, 0x66, 0x00, 0x9f, 0x6c, 0x46, 0x43, 0x85, 0x46,
	0x0f, 0x10, 0xec, 0x78, 0x4d, 0xbb, 0xd8, 0x4b, 0xfc, 0x2c, 0xba, 0x7b, 0x8b, 0x27, 0x12, 0xf1,
	0x89, 0xc7, 0x5f, 0x0c, 0xfc, 0x59, 0xe8, 0x76, 0x28, 0x9c, 0x11, 0x3d, 0x40, 0x58, 0x49, 0xa1,
	0x09, 0x17, 0xb4, 0x8d, 0x67, 0x89, 0x97, 0x45, 0x77, 0xe9, 0x64, 0xca, 0xea, 0x40, 0xad, 0xc5,
	0x4e, 0x16, 0x27, 0x13, 0xc2, 0xe0, 0xff, 0xe5, 0x2a, 0xf6, 0xad, 0xf7, 0x66, 0xd2, 0xbb, 0xe5,
	0xca, 0xba, 0x0c, 0x78, 0xfd, 0x03, 0xe0, 0xd4, 0x06, 0x7a, 0x0a, 0xfe, 0x1f, 0x3a, 0xc4, 0x5e,
	0xe2, 0x65, 0x61, 0x61, 0x3e, 0xd1, 0x3d, 0x04, 0x7b, 0x52, 0xf7, 0x74, 0xec, 0xe6, 0x76, 0x32,
	0xd1, 0x24, 0xd8, 0x48, 0xc7, 0xbe, 0x9f, 0xbd, 0xf3, 0x52, 0x02, 0xf3, 0x83, 0x8c, 0x6e, 0x01,
	0x3a, 0xd9, 0xb7, 0x15, 0x2d, 0xfb, 0xb6, 0x1e, 0xd3, 0x43, 0xa7, 0x7c, 0x6f, 0x6b, 0xf4, 0x12,
	0xe6, 0xdd, 0x6f, 0xb2, 0x28, 0xbb, 0xbe, 0xb1, 0xc7, 0x84, 0xc5, 0x95, 0xa9, 0xbf, 0xf6, 0x0d,
	0x7a, 0x05, 0x61, 0xc3, 0x1b, 0x5a, 0xea, 0x41, 0x51, 0x7b, 0xa9, 0xb0, 0x98, 0x1b, 0xe1, 0xdb,
	0xa0, 0x68, 0xfa, 0x06, 0x9e, 0x9c, 0xbd, 0x03, 0x7a, 0x0e, 0x01, 0x6f, 0x08, 0xa3, 0xe3, 0x11,
	0xae, 0x48, 0xd7, 0x70, 0x35, 0x5e, 0xf9, 0xa2, 0x11, 0xff, 0xb2, 0x91, 0xd7, 0x10, 0xd9, 0x39,
	0x94, 0x95, 0xec, 0x85, 0x8e, 0x1f, 0x25, 0x5e, 0x16, 0x14, 0x60, 0xa5, 0x95, 0x51, 0x96, 0x3b,
	0x78, 0x51, 0xc9, 0x66, 0xea, 0x0d, 0x96, 0x91, 0x1b, 0xec, 0xc6, 0x2c, 0xc6, 0xc6, 0xdb, 0x7e,
	0x1c, 0x19, 0x26, 0x6b, 0x22, 0x18, 0x96, 0x2d, 0xcb, 0x19, 0x15, 0x76, 0x6d, 0x72, 0xf7, 0x8b,
	0x28, 0xde, 0x9d, 0xad, 0xe3, 0x87, 0x63, 0xf1, 0xf3, 0xb1, 0x05, 0xef, 0xff, 0x07, 0x00, 0x00,
	0xff, 0xff, 0x6e, 0xeb, 0x52, 0x5a, 0xb6, 0x02, 0x00, 0x00,
}
