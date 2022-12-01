# Storage Type


1. Ephemeral Ephemeral volumes are designed for these use cases. Because volumes follow the Pod's lifetime and get created and deleted along with the Pod, Pods can be stopped and restarted without being limited to where some persistent volume is available. Ephemeral volumes are specified inline in the Pod spec, which simplifies application deployment and management.

2. NAS Storage / FileStorage (SMB, NFS, AFS)
3. SAN Storage / BlockStorage (iSCSI)


# Container Storage Interface



## Controller
Controller service interface. This interface is responsible of **controlling and managing the volumes, such as: creating, deleting, attaching/detaching, snapshotting**

If the volumes are part of a Cloud Provider (such as DigitalOcean, GKE, AWS), this interface must be implemented. However if you're planning not to use any kind of block storage or have other ways of providing storage space, you don't have to create this interface

## Node
**The Node plugin is a gRPC server that needs to run on the Node where the volume will be provisioned.** So suppose you have a Kubernetes cluster with  three nodes where your Pod's are scheduled, you would deploy this to all three nodes.


## Identity

This service needs to be implemented for each individual plugin.For example, if you have two separate plugins running, Node and Controller, both binaries  need to implement the Identity gRPC interface individually.

## driver-registrar  
Sidecar container that registers the CSI driver with kubelet, and adds  the drivers custom NodeId to a label on the Kubernetes Node API Object.

## external-provisioner 

Sidecar container that watches Kubernetes PersistentVolumeClaim objects and triggers CreateVolume/DeleteVolume against a CSI endpoint.

## external-attacher  
Sidecar container that watches Kubernetes VolumeAttachment objects and triggers ControllerPublish/Unpublish against a CSI endpoint
## Volume Life Cycle

<p align="center"><img src="images/volumelifecycle/fig1.png" /></p>

Fig.1 The lifecycle of a dynamically provisioned volume, from
creation to destruction.

<p align="center"><img src="images/volumelifecycle/fig2.png" /></p>

Fig.2 The lifecycle of a dynamically provisioned volume, from
creation to destruction, when the Node Plugin advertises the
STAGE_UNSTAGE_VOLUME capability.

<p align="center"><img src="images/volumelifecycle/fig3.png" /></p>

Fig.3 The lifecycle of a pre-provisioned volume that requires
controller to publish to a node (`ControllerPublishVolume`) prior to
publishing on the node (`NodePublishVolume`).

<p align="center"><img src="images/volumelifecycle/fig4.png" /></p>

Fig.4 Plugins MAY forego other lifecycle steps by contraindicating
them via the capabilities API. Interactions with the volumes of such
plugins is reduced to `NodePublishVolume` and `NodeUnpublishVolume`
calls.
# CSI Volume Snapshots

A snapshot represents the state of the storage volume in a cluster at a particular point in time. Volume snapshots can be used to provision a new volume. volume snapshots with supported Container Storage Interface (CSI) drivers to help protect against data loss

With CSI volume snapshots, a cluster administrator can:
- Deploy a third-party CSI driver that supports snapshots.  
- Create a new persistent volume claim (PVC) from an existing volume snapshot.
- Take a snapshot of an existing PVC.
- Restore a snapshot as a different PVC.
- Delete an existing volume snapshot.

With CSI volume snapshots, an app developer can:
- Use volume snapshots as building blocks for developing application- or cluster-level storage backup solutions.
- Rapidly rollback to a previous development version.
- Use storage more efficiently by not having to make a full copy each time.

Be aware of the following when using volume snapshots:
- Support is only available for CSI drivers. In-tree and FlexVolumes are not supported.
- CSI drivers may or may not have implemented the volume snapshot functionality. CSI drivers that have provided support for volume snapshots will likely use the csi-external-snapshotter sidecar. See documentation provided by the CSI driver for details.

# REF

[Kubernetes CSI NFS](https://github.com/kubernetes-retired/drivers/tree/master/pkg/nfs)

