# Differential Snapshot Controller

This repository contains the source of the differential snapshot controller
prototype.

Prerequisites:

* Go
* Docker

To compile the controller Go binary locally:

```sh
make compile
```

This will produce a local binary named `diffsnap-controller`. To run it:

```sh
./diffsnap-controller
```

To build and push the controller's image:

```sh
make build

make push
```

The image name and tag can be overridden using the `IMAGE_NAME` and `IMAGE_TAG`
variables.

To deploy to Kubernetes:

```sh
kubectl apply -f artifacts/examples/differentialsnapshot.example.com_getchangedblocks.yaml

kubectl apply -f deploy/controller/ns.yaml

kubectl apply -f deploy/controller/rbac.yaml

kubectl apply -f deploy/controller/deploy.yaml
```

The controller will be deployed to the new `diffsnap` namespace.

An example of the GetChangedBlock custom resource can be found in
`artifacts/examples/getchangedblocks.yaml`.

## On AWS EKS

Instructions on how to set up an EKS cluster with the CSI snapshot controller
and EBS CSI driver can be found in the `deploy/eks/README.md` file.
