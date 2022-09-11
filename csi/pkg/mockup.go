package csilab

import (
	proto "github.com/golang/protobuf/proto"
)

type GetPluginInfoRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPluginInfoRequest) Reset()         { *m = GetPluginInfoRequest{} }
func (m *GetPluginInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetPluginInfoRequest) ProtoMessage()    {}
func (*GetPluginInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9cdb00adce470e01, []int{0}
}

func (m *GetPluginInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPluginInfoRequest.Unmarshal(m, b)
}
func (m *GetPluginInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPluginInfoRequest.Marshal(b, m, deterministic)
}
func (m *GetPluginInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPluginInfoRequest.Merge(m, src)
}
func (m *GetPluginInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetPluginInfoRequest.Size(m)
}
func (m *GetPluginInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPluginInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPluginInfoRequest proto.InternalMessageInfo

type GetPluginInfoResponse struct {
	// The name MUST follow domain name notation format
	// (https://tools.ietf.org/html/rfc1035#section-2.3.1). It SHOULD
	// include the plugin's host company name and the plugin name,
	// to minimize the possibility of collisions. It MUST be 63
	// characters or less, beginning and ending with an alphanumeric
	// character ([a-z0-9A-Z]) with dashes (-), dots (.), and
	// alphanumerics between. This field is REQUIRED.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// This field is REQUIRED. Value of this field is opaque to the CO.
	VendorVersion string `protobuf:"bytes,2,opt,name=vendor_version,json=vendorVersion,proto3" json:"vendor_version,omitempty"`
	// This field is OPTIONAL. Values are opaque to the CO.
	Manifest             map[string]string `protobuf:"bytes,3,rep,name=manifest,proto3" json:"manifest,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetPluginInfoResponse) Reset()         { *m = GetPluginInfoResponse{} }
func (m *GetPluginInfoResponse) String() string { return proto.CompactTextString(m) }
func (*GetPluginInfoResponse) ProtoMessage()    {}
func (*GetPluginInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9cdb00adce470e01, []int{1}
}

func (m *GetPluginInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPluginInfoResponse.Unmarshal(m, b)
}
func (m *GetPluginInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPluginInfoResponse.Marshal(b, m, deterministic)
}
func (m *GetPluginInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPluginInfoResponse.Merge(m, src)
}
func (m *GetPluginInfoResponse) XXX_Size() int {
	return xxx_messageInfo_GetPluginInfoResponse.Size(m)
}
func (m *GetPluginInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPluginInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPluginInfoResponse proto.InternalMessageInfo

func (m *GetPluginInfoResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetPluginInfoResponse) GetVendorVersion() string {
	if m != nil {
		return m.VendorVersion
	}
	return ""
}

func (m *GetPluginInfoResponse) GetManifest() map[string]string {
	if m != nil {
		return m.Manifest
	}
	return nil
}
