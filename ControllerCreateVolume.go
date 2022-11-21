type VolumeCapability_AccessMode_Mode int32

const (
	VolumeCapability_AccessMode_UNKNOWN VolumeCapability_AccessMode_Mode = 0
	// Can only be published once as read/write on a single node, at
	// any given time.
	VolumeCapability_AccessMode_SINGLE_NODE_WRITER VolumeCapability_AccessMode_Mode = 1
	// Can only be published once as readonly on a single node, at
	// any given time.
	VolumeCapability_AccessMode_SINGLE_NODE_READER_ONLY VolumeCapability_AccessMode_Mode = 2
	// Can be published as readonly at multiple nodes simultaneously.
	VolumeCapability_AccessMode_MULTI_NODE_READER_ONLY VolumeCapability_AccessMode_Mode = 3
	// Can be published at multiple nodes simultaneously. Only one of
	// the node can be used as read/write. The rest will be readonly.
	VolumeCapability_AccessMode_MULTI_NODE_SINGLE_WRITER VolumeCapability_AccessMode_Mode = 4
	// Can be published as read/write at multiple nodes
	// simultaneously.
	VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER VolumeCapability_AccessMode_Mode = 5
	// Can only be published once as read/write at a single workload
	// on a single node, at any given time. SHOULD be used instead of
	// SINGLE_NODE_WRITER for COs using the experimental
	// SINGLE_NODE_MULTI_WRITER capability.
	VolumeCapability_AccessMode_SINGLE_NODE_SINGLE_WRITER VolumeCapability_AccessMode_Mode = 6
	// Can be published as read/write at multiple workloads on a
	// single node simultaneously. SHOULD be used instead of
	// SINGLE_NODE_WRITER for COs using the experimental
	// SINGLE_NODE_MULTI_WRITER capability.
	VolumeCapability_AccessMode_SINGLE_NODE_MULTI_WRITER VolumeCapability_AccessMode_Mode = 7
)

var VolumeCapability_AccessMode_Mode_name = map[int32]string{
	0: "UNKNOWN",
	1: "SINGLE_NODE_WRITER",
	2: "SINGLE_NODE_READER_ONLY",
	3: "MULTI_NODE_READER_ONLY",
	4: "MULTI_NODE_SINGLE_WRITER",
	5: "MULTI_NODE_MULTI_WRITER",
	6: "SINGLE_NODE_SINGLE_WRITER",
	7: "SINGLE_NODE_MULTI_WRITER",
}

var VolumeCapability_AccessMode_Mode_value = map[string]int32{
	"UNKNOWN":                   0,
	"SINGLE_NODE_WRITER":        1,
	"SINGLE_NODE_READER_ONLY":   2,
	"MULTI_NODE_READER_ONLY":    3,
	"MULTI_NODE_SINGLE_WRITER":  4,
	"MULTI_NODE_MULTI_WRITER":   5,
	"SINGLE_NODE_SINGLE_WRITER": 6,
	"SINGLE_NODE_MULTI_WRITER":  7,
}

type Topology struct {
	Segments             map[string]string `protobuf:"bytes,1,rep,name=segments,proto3" 		json:"segments,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

type VolumeContentSource struct {
	Type                 isVolumeContentSource_Type `protobuf_oneof:"type"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

type VolumeCapability struct {
	AccessType           isVolumeCapability_AccessType `protobuf_oneof:"access_type"`
	AccessMode           *VolumeCapability_AccessMode  `protobuf:"bytes,3,opt,name=access_mode,json=accessMode,proto3"		json:"access_mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

type CapacityRange struct {
	RequiredBytes        int64    `protobuf:"varint,1,opt,name=required_bytes,json=requiredBytes,proto3"		json:"required_bytes,omitempty"`
	LimitBytes           int64    `protobuf:"varint,2,opt,name=limit_bytes,json=limitBytes,proto3"			json:"limit_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

type TopologyRequirement struct {
	Requisite            []*Topology `protobuf:"bytes,1,rep,name=requisite,proto3" 		json:"requisite,omitempty"`
	Preferred            []*Topology `protobuf:"bytes,2,rep,name=preferred,proto3"		json:"preferred,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

type CreateVolumeRequest struct {
	Name                      string               `protobuf:"bytes,1,opt,name=name,proto3" 								json:"name,omitempty"`
	CapacityRange             *CapacityRange       `protobuf:"bytes,2,opt,name=capacity_range,json=capacityRange,proto3" 					json:"capacity_range,omitempty"`
	VolumeCapabilities        []*VolumeCapability  `protobuf:"bytes,3,rep,name=volume_capabilities,json=volumeCapabilities,proto3" 			   	json:"volume_capabilities,omitempty"`
	Parameters                map[string]string    `protobuf:"bytes,4,rep,name=parameters,proto3" 								json:"parameters,omitempty" 			protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Secrets                   map[string]string    `protobuf:"bytes,5,rep,name=secrets,proto3" 								json:"secrets,omitempty" 			protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	VolumeContentSource       *VolumeContentSource `protobuf:"bytes,6,opt,name=volume_content_source,json=volumeContentSource,proto3" 		   	json:"volume_content_source,omitempty"`
	AccessibilityRequirements *TopologyRequirement `protobuf:"bytes,7,opt,name=accessibility_requirements,json=accessibilityRequirements,proto3" 		json:"accessibility_requirements,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}             `json:"-"`
	XXX_unrecognized          []byte               `json:"-"`
	XXX_sizecache             int32                `json:"-"`
}
