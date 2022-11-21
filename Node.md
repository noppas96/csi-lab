# Node

- NodeStageVolume

```go
type NodeStageVolumeRequest struct {
    VolumeId string `protobuf:"bytes,1,opt,name=volume_id,json=volumeId,proto3" json:"volume_id,omitempty"`
    PublishContext map[string]string `protobuf:"bytes,2,rep,name=publish_context,json=publishContext,proto3" json:"publish_context,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
    StagingTargetPath string `protobuf:"bytes,3,opt,name=staging_target_path,json=stagingTargetPath,proto3" json:"staging_target_path,omitempty"`
    VolumeCapability *VolumeCapability `protobuf:"bytes,4,opt,name=volume_capability,json=volumeCapability,proto3" json:"volume_capability,omitempty"`
    Secrets map[string]string `protobuf:"bytes,5,rep,name=secrets,proto3" json:"secrets,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
    VolumeContext        map[string]string `protobuf:"bytes,6,rep,name=volume_context,json=volumeContext,proto3" json:"volume_context,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
    XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`

func (*NodeStageVolumeRequest) Descriptor() ([]byte, []int)
func (m *NodeStageVolumeRequest) GetPublishContext() map[string]string
func (m *NodeStageVolumeRequest) GetSecrets() map[string]string
func (m *NodeStageVolumeRequest) GetStagingTargetPath() string
func (m *NodeStageVolumeRequest) GetVolumeCapability() *VolumeCapability
func (m *NodeStageVolumeRequest) GetVolumeContext() map[string]string
func (m *NodeStageVolumeRequest) GetVolumeId() string
func (*NodeStageVolumeRequest) ProtoMessage()
func (m *NodeStageVolumeRequest) Reset()
func (m *NodeStageVolumeRequest) String() string
func (m *NodeStageVolumeRequest) XXX_DiscardUnknown()
func (m *NodeStageVolumeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
func (m *NodeStageVolumeRequest) XXX_Merge(src proto.Message)
func (m *NodeStageVolumeRequest) XXX_Size() int
func (m *NodeStageVolumeRequest) XXX_Unmarshal(b []byte) error    
```

- NodePublishVolume

```go
type NodePublishVolumeRequest struct {
    VolumeId string `protobuf:"bytes,1,opt,name=volume_id,json=volumeId,proto3" json:"volume_id,omitempty"`
    PublishContext map[string]string `protobuf:"bytes,2,rep,name=publish_context,json=publishContext,proto3" json:"publish_context,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
    StagingTargetPath string `protobuf:"bytes,3,opt,name=staging_target_path,json=stagingTargetPath,proto3" json:"staging_target_path,omitempty"`
    TargetPath string `protobuf:"bytes,4,opt,name=target_path,json=targetPath,proto3" json:"target_path,omitempty"`
    VolumeCapability *VolumeCapability `protobuf:"bytes,5,opt,name=volume_capability,json=volumeCapability,proto3" json:"volume_capability,omitempty"`
    Readonly bool `protobuf:"varint,6,opt,name=readonly,proto3" json:"readonly,omitempty"`
    Secrets map[string]string `protobuf:"bytes,7,rep,name=secrets,proto3" json:"secrets,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
    VolumeContext        map[string]string `protobuf:"bytes,8,rep,name=volume_context,json=volumeContext,proto3" json:"volume_context,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
    XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`

func (*NodePublishVolumeRequest) Descriptor() ([]byte, []int)
func (m *NodePublishVolumeRequest) GetPublishContext() map[string]string
func (m *NodePublishVolumeRequest) GetReadonly() bool
func (m *NodePublishVolumeRequest) GetSecrets() map[string]string
func (m *NodePublishVolumeRequest) GetStagingTargetPath() string
func (m *NodePublishVolumeRequest) GetTargetPath() string
func (m *NodePublishVolumeRequest) GetVolumeCapability() *VolumeCapability
func (m *NodePublishVolumeRequest) GetVolumeContext() map[string]string
func (m *NodePublishVolumeRequest) GetVolumeId() string
func (*NodePublishVolumeRequest) ProtoMessage()
func (m *NodePublishVolumeRequest) Reset()
func (m *NodePublishVolumeRequest) String() string
func (m *NodePublishVolumeRequest) XXX_DiscardUnknown()
func (m *NodePublishVolumeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
func (m *NodePublishVolumeRequest) XXX_Merge(src proto.Message)
func (m *NodePublishVolumeRequest) XXX_Size() int
func (m *NodePublishVolumeRequest) XXX_Unmarshal(b []byte) error
```