# CSI Research

## Identity service interface
```
service Identity {
  rpc GetPluginInfo(GetPluginInfoRequest)
	returns (GetPluginInfoResponse) {}

  rpc GetPluginCapabilities(GetPluginCapabilitiesRequest)
	returns (GetPluginCapabilitiesResponse) {}

  rpc Probe (ProbeRequest)
	returns (ProbeResponse) {}
}
```

The Identity service consists of very basic methods and is mainly for identifying the service, making sure it’s healthy, and returning basic information about the plugin itself (whether it’s a Node plugin or Controller plugin). Here is a basic Go implementation of this service

**GetPluginInfo:** This method needs to return the version and name of the plugin.

**GetPluginCapabilities:** This method returns the capabilities of the plugin. Currently it reports whether the plugin has the ability of serving the Controller interface. The CO calls the Controller interface methods depending on whether this method returns the capability or not.

**Probe:** This is called by the CO just to check whether the plugin is running or not. This method doesn’t need to return anything. Currently the spec doesn’t dictate what you should return either. Hence, return an empty response.
## Node service interface
```
service Node {
  rpc NodeStageVolume (NodeStageVolumeRequest)
	returns (NodeStageVolumeResponse) {}

  rpc NodeUnstageVolume (NodeUnstageVolumeRequest)
	returns (NodeUnstageVolumeResponse) {}

  rpc NodePublishVolume (NodePublishVolumeRequest)
	returns (NodePublishVolumeResponse) {}

  rpc NodeUnpublishVolume (NodeUnpublishVolumeRequest)
	returns (NodeUnpublishVolumeResponse) {}

  rpc NodeGetId (NodeGetIdRequest)
	returns (NodeGetIdResponse) {}

  rpc NodeGetCapabilities (NodeGetCapabilitiesRequest)
	returns (NodeGetCapabilitiesResponse) {}
}
```

**NodeStageVolume:** This method is called by the CO to temporarily mount the volume to a staging path. Usually this staging path is a global directory on the node. In Kubernetes, after it’s mounted to the global directory, you mount it into the pod directory (via NodePublishVolume). The reason that mounting is a two step operation is because Kubernetes allows you to use a single volume by multiple pods. This is allowed when the storage system supports it (say NFS) or if all pods run on the same node. One thing to note is that you also need to format the volume if it’s not formatted already. Keep that in mind.

**NodeUnstageVolume:** This method is called by the CO to unmount the volume from the staging path. It’s the reverse of NodeStageVolume

**NodePublishVolume:** This method is called to mount the volume from staging to target path. Usually what you do here is a bind mount. A bind mount allows you to mount a path to a different path (instead of mounting a device to a path). In Kubernetes, this allows us for example to use the mounted volume from the staging path (i.e global directory) to the target path (pod directory). Here, formatting is not needed because we already did it in NodeStageVolume.

**NodeUnpublishVolume:** This is the reverse of NodePublishVolume. It unmounts the volume from the target path

**NodeGetId:** This method should return a unique ID of the node on which this plugin is running. For example, for DigitalOcean we return the Droplet ID.

**NodeGetCapabilities:** Just like ControllerGetCapabilities, this returns the capabilities of the Node plugin. For example if you don’t advertise RPC_STAGE_UNSTAGE_VOLUME capability, the CO will not call NodeStageVolume and NodeUnstageVolume as you don’t provide it.

[Ref. CSI Repo](https://github.com/container-storage-interface/spec)