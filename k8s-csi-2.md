# CSI NFS

CSI driver is typically deployed in Kubernetes as two components: a controller component and a per-node component.

# Install the CSI driver for NFS

1. which, deploy successful, should get csi-controller, csi -node pod, and csidrivers.kubernetes.resource, when we use kubectl get csidrivers result should be like this nfs.csi.k8s.io

2. Create a Kubernetes Storage Class that uses the nfs.csi.k8s.io CSI driver. Assuming you have configured an NFS share /srv/nfs and the address of your NFS server is 10.0.0.42

```yaml
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: nfs-csi
provisioner: nfs.csi.k8s.io
parameters:
  server: 10.0.0.42
  share: /srv/nfs
reclaimPolicy: Delete
volumeBindingMode: Immediate
mountOptions:
  - hard
  - nfsvers=4.1
```

3. create a new PersistentVolumeClaim using the nfs-csi storage class. This is as simple as specifying storageClassName: nfs-csi in the PVC definition

```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: my-pvc
spec:
  storageClassName: nfs-csi
  accessModes: [ReadWriteOnce]
  resources:
    requests:
      storage: 5Gi
```
