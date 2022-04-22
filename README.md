# Differential Snapshot Controller

<<<<<<< HEAD
This is a prototype that implements the controller for GetChangedBlocks API

# Prerequisite:
- AWS EKS: See instruction in deploy/eks.md and sample in deploy/k8s
- AWS CSI Driver: See instruction in deploy/README.md

# Build:

```sh
go build -o diffsnapcontroller .
```

# Deploy:

Currently we have not yet have the docker container image for this controller yet so follow the steps below to deploy an Ubuntu pod and copy the binary there to run:

```sh
kubectl create ns testns
kubectl apply -f artifacts/examples/controller-rbac.yaml
kubectl apply -f artifacts/examples/differentialsnapshot.example.com_getchangedblocks.yaml
kubectl apply -f artifacts/examples/testpod.yaml
kubectl cp diffsnapcontroller -n testns phtest:/root
```

Also copy additional files for AWS such as credential, config etc. as needed to the pod.

# Run:

```sh
kubectl exec -it -n testns phtest /root/diffsnapcontroller
```

# Create GetChangedBlock CR:

See the example GetChangedBlock CR in the file artifacts/examples/getchangedblocks.yaml
You can edit it and run command below to create it.
```sh
kubectl apply -f artifacts/examples/getchangedblocks.yaml
```
