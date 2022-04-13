# AWS EBS CSI Driver

This document provides instructions on how to deploy the EBS CSI driver and
the CSI snapshot controller on an EKS cluster.

Information on how to set up an EKS cluster with the appropriate IAM permissions
can be found [here](eks.md).

## Deploying CSI Snapshot Controller

Per issue
[#1065](https://github.com/kubernetes-sigs/aws-ebs-csi-driver/issues/1065),
the CSI snapshot controller **must** be deployed before the EBS CSI driver:

```sh
kubectl apply -f https://raw.githubusercontent.com/kubernetes-csi/external-snapshotter/master/client/config/crd/snapshot.storage.k8s.io_volumesnapshotclasses.yaml

kubectl apply -f https://raw.githubusercontent.com/kubernetes-csi/external-snapshotter/master/client/config/crd/snapshot.storage.k8s.io_volumesnapshotcontents.yaml

kubectl apply -f https://raw.githubusercontent.com/kubernetes-csi/external-snapshotter/master/client/config/crd/snapshot.storage.k8s.io_volumesnapshots.yaml

kubectl apply -f https://raw.githubusercontent.com/kubernetes-csi/external-snapshotter/master/deploy/kubernetes/snapshot-controller/rbac-snapshot-controller.yaml

kubectl apply -f https://raw.githubusercontent.com/kubernetes-csi/external-snapshotter/master/deploy/kubernetes/snapshot-controller/setup-snapshot-controller.yaml
```

Confirm that the snapshot controller pods are ready:

```sh
$ kubectl -n kube-system get po -lapp=snapshot-controller
NAME                                   READY   STATUS    RESTARTS   AGE
snapshot-controller-75fd799dc8-5wzxv   1/1     Running   0          45s
snapshot-controller-75fd799dc8-g7v79   1/1     Running   0          45s
```

## Deploying EBS CSI Driver

Use [Helm](https://helm.sh/docs/intro/install/) v3.8.0 to deploy the EBS CSI
driver:

```sh
helm repo add aws-ebs-csi-driver https://kubernetes-sigs.github.io/aws-ebs-csi-driver

helm repo update

helm upgrade --install aws-ebs-csi-driver \
    --namespace kube-system \
    aws-ebs-csi-driver/aws-ebs-csi-driver
```

Confirm that the driver controller pods are ready:

```sh
$ kubectl get pod -n kube-system -l "app.kubernetes.io/name=aws-ebs-csi-driver,app.kubernetes.io/instance=aws-ebs-csi-driver"
NAME                                  READY   STATUS    RESTARTS   AGE
ebs-csi-controller-5d8f969694-5p77x   6/6     Running   0          20h
ebs-csi-controller-5d8f969694-7kx96   6/6     Running   0          20h
ebs-csi-node-gkqwq                    3/3     Running   0          20h
ebs-csi-node-rbxsr                    3/3     Running   0          20h
```

Follow the steps in the [eks.md](eks.md) documentation to attach the appropriate
IAM permissions to the driver controller's service account.

## Deploy Client Application

Deploy the namespace, pod, PVC and storage class resources used for testing:

```sh
kubectl apply -f k8s/ns.yaml

kubectl apply -f k8s/class.yaml

kubectl apply -f k8s/app.yaml
```

Confirm that the pod has a block PVC attached:

```sh
$ kubectl -n diffsnap get po,pvc,pv
NAME              READY   STATUS    RESTARTS   AGE
pod/diffsnap-io   2/2     Running   0          32s

NAME                                   STATUS   VOLUME                                     CAPACITY   ACCESS MODES   STORAGECLASS   AGE
persistentvolumeclaim/diffsnap-block   Bound    pvc-5615fa3e-0847-4544-b227-88d6a336a3dd   4Gi        RWO            ebs-csi        32s

NAME                                                        CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS   CLAIM                     STORAGECLASS   REASON   AGE
persistentvolume/pvc-5615fa3e-0847-4544-b227-88d6a336a3dd   4Gi        RWO            Delete           Bound    diffsnap/diffsnap-block   ebs-csi                 29s
```

### Creating The Snapshots

Use the `kubectl exec` command to invoke the `dd` tool to write 512KB of
data to the 1st block on the EBS volume:

```sh
kubectl -n diffsnap exec diffsnap-client -c writer -- /bin/sh -c "dd if=/dev/urandom of=/dev/xvda bs=512k count=1"
```

The `dd` command uses the following options to alter the write behaviour:

* `bs` the total amount of data to write
* `count` the total number of blocks to write

Take the 1st snapshot:

```sh
kubectl create -f k8s/snapshot.yaml
```

The volume snapshots' `READYTOUSE` field will be set to `true` once the actual
snapshot is ready:

```sh
$ kubectl -n diffsnap get vs,vsc
NAME                                                    READYTOUSE   SOURCEPVC        SOURCESNAPSHOTCONTENT   RESTORESIZE   SNAPSHOTCLASS   SNAPSHOTCONTENT                                    CREATIONTIME   AGE
volumesnapshot.snapshot.storage.k8s.io/diffsnap-tjkx6   true         diffsnap-block                           4Gi           ebs-csi         snapcontent-f591a7de-bbf3-4266-a042-9f16ce94d7ac   40s            40s

NAME                                                                                             READYTOUSE   RESTORESIZE   DELETIONPOLICY   DRIVER            VOLUMESNAPSHOTCLASS   VOLUMESNAPSHOT   VOLUMESNAPSHOTNAMESPACE   AGE
volumesnapshotcontent.snapshot.storage.k8s.io/snapcontent-f591a7de-bbf3-4266-a042-9f16ce94d7ac   true         4294967296    Delete           ebs.csi.aws.com   ebs-csi               diffsnap-tjkx6   diffsnap                  40s
```

Use the same `kubectl exec` command to write a different set of data to the same
block on the volume:

```sh
kubectl -n diffsnap exec diffsnap-client -c writer -- /bin/sh -c "dd if=/dev/urandom of=/dev/xvda bs=512k count=1"
```

Take the 2nd snapshot:

```sh
kubectl create -f k8s/snapshot.yaml
```

Confirm that the volume snapshot and its content are created:

```sh
$ kubectl -n diffsnap get vs,vsc
NAME                                                    READYTOUSE   SOURCEPVC        SOURCESNAPSHOTCONTENT   RESTORESIZE   SNAPSHOTCLASS   SNAPSHOTCONTENT                                    CREATIONTIME   AGE
volumesnapshot.snapshot.storage.k8s.io/diffsnap-tjkx6   true         diffsnap-block                           4Gi           ebs-csi         snapcontent-f591a7de-bbf3-4266-a042-9f16ce94d7ac   129m           129m
volumesnapshot.snapshot.storage.k8s.io/diffsnap-x2lxg   true         diffsnap-block                           4Gi           ebs-csi         snapcontent-6865ccde-6854-483a-9dc1-f32266ffe904   3m3s           3m4s

NAME                                                                                             READYTOUSE   RESTORESIZE   DELETIONPOLICY   DRIVER            VOLUMESNAPSHOTCLASS   VOLUMESNAPSHOT   VOLUMESNAPSHOTNAMESPACE   AGE
volumesnapshotcontent.snapshot.storage.k8s.io/snapcontent-6865ccde-6854-483a-9dc1-f32266ffe904   true         4294967296    Delete           ebs.csi.aws.com   ebs-csi               diffsnap-x2lxg   diffsnap                  3m4s
volumesnapshotcontent.snapshot.storage.k8s.io/snapcontent-f591a7de-bbf3-4266-a042-9f16ce94d7ac   true         4294967296    Delete           ebs.csi.aws.com   ebs-csi               diffsnap-tjkx6   diffsnap                  129m
```

Retrieve the snapshot IDs from the `VolumeSnapshotContent` resources:

```sh
kubectl get vsc -ojsonpath='{.items[*].status.snapshotHandle}'
```

Use the `aws` CLI to list the changed blocks:

```sh
$ aws ebs list-changed-blocks --first-snapshot-id <aws_first_snapshot_id> --second-snapshot-id <aws_second_snapshot_id>
- BlockSize: 524288
  ChangedBlocks:
  - BlockIndex: 0
    FirstBlockToken: dG9rZW4tMDAK
    SecondBlockToken: dG9rZW4tMDEK
  ExpiryTime: '2022-04-19T11:28:37.982000-07:00'
  VolumeSize: 4
```

The `ebs:ListChangedBlocks` IAM permission is required in order for the
`list-changed-blocks` subcommand to work.

## Clean Up

Use the following commands to delete the unwanted resources:

```sh
kubectl delete -f k8s/class.yaml

kubectl --diffsnap delete volumesnapshot -lapp=diffsnap

kubectl --diffsnap delete po -lapp=diffsnap

kubectl delete ns -lapp=diffsnap
```
