#!/bin/bash

## Steps to generate structural CRDs
## 1: Manually run 'mvn -Pgo-env -N' at golang/controller to set the GO ENV variables in the current workspace.
## 2: Run this script from this path -> vproxy-kubernetes/golang/controller/src
## 3: CRD YAML files will be available at vproxy-kubernetes/golang/controller/src/config/crd/bases

## find or download controller-gen
CONTROLLER_GEN=$(which controller-gen)

if [ "$CONTROLLER_GEN" = "" ]
then
  go install sigs.k8s.io/controller-tools/cmd/controller-gen@v0.4.1;
  CONTROLLER_GEN=$(which controller-gen)
fi

if [ "$CONTROLLER_GEN" = "" ]
then
  echo "ERROR: failed to get controller-gen";
  exit 1;
fi

SCRIPT_ROOT=$(pwd)


$CONTROLLER_GEN crd:crdVersions=v1,trivialVersions=true,preserveUnknownFields=false \
paths=../pkg/apis/differentialsnapshot/v1alpha1 \
output:crd:artifacts:config=../artifacts/examples


