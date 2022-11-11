# CSI Hostpath

```go
type hostPath struct {
	config Config

	// gRPC calls involving any of the fields below must be serialized
	// by locking this mutex before starting. Internal helper
	// functions assume that the mutex has been locked.
	mutex sync.Mutex
	state state.State
}

type Config struct {
	DriverName                 string
	Endpoint                   string
	ProxyEndpoint              string
	NodeID                     string
	VendorVersion              string
	StateDir                   string
	MaxVolumesPerNode          int64
	MaxVolumeSize              int64
	AttachLimit                int64
	Capacity                   Capacity
	Ephemeral                  bool
	ShowVersion                bool
	EnableAttach               bool
	EnableTopology             bool
	EnableVolumeExpansion      bool
	DisableControllerExpansion bool
	DisableNodeExpansion       bool
	MaxVolumeExpansionSizeNode int64
	CheckVolumeLifecycle       bool
}

```

# CSI NFS

```go
// DriverOptions defines driver parameters specified in driver deployment
type DriverOptions struct {
	NodeID           string
	DriverName       string
	Endpoint         string
	MountPermissions uint64
	WorkingMountDir  string
}

type Driver struct {
	name             string
	nodeID           string
	version          string
	endpoint         string
	mountPermissions uint64
	workingMountDir  string

	//ids *identityServer
	ns          *NodeServer
	cscap       []*csi.ControllerServiceCapability
	nscap       []*csi.NodeServiceCapability
	volumeLocks *VolumeLocks
}

const (
	DefaultDriverName = "nfs.csi.k8s.io"
	// Address of the NFS server
	paramServer = "server"
	// Base directory of the NFS server to create volumes under.
	// The base directory must be a direct child of the root directory.
	// The root directory is omitted from the string, for example:
	//     "base" instead of "/base"
	paramShare            = "share"
	paramSubDir           = "subdir"
	mountOptionsField     = "mountoptions"
	mountPermissionsField = "mountpermissions"
	pvcNameKey            = "csi.storage.k8s.io/pvc/name"
	pvcNamespaceKey       = "csi.storage.k8s.io/pvc/namespace"
	pvNameKey             = "csi.storage.k8s.io/pv/name"
	pvcNameMetadata       = "${pvc.metadata.name}"
	pvcNamespaceMetadata  = "${pvc.metadata.namespace}"
	pvNameMetadata        = "${pv.metadata.name}"
)

```