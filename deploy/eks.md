## AWS EKS

This document contains information on how to set up an AWS EKS cluster with the
appropriate IAM permissions required by the EBS CSI driver.

Provision a 2-node EKS cluster with the IAM OIDC provider enabled, using
[`eksctl`](https://eksctl.io/introduction/) 0.92.0:

```sh
eksctl create cluster --name=<cluster_name> --nodes=2 --with-oidc
```

After the EBS CSI driver is deployed, its controller's service account must be
attached to an AWS IAM role with the appropriate permissions to manage EBS
volumes.

Create the AWS IAM policy defined in the `iam-policy.json` file:

```sh
aws iam create-policy \
  policy-name <iam_policy_name> \
  --policy-document file://iam-policy.json
```

`eksctl` provides a helper command to create a new AWS IAM role and attach it to
to the CSI controller's `ebs-csi-controller-sa` service account:

```sh
eksctl create iamserviceaccount \
    --name ebs-csi-controller-sa \
    --namespace kube-system \
    --cluster <cluster_name> \
    --role-name <new_role_name> \
    --attach-policy-arn arn:aws:iam::<aws_account_id>:policy/<iam_policy_name> \
    --approve \
    --override-existing-serviceaccounts
```

Confirm that the controller's service account has the correct
`eks.amazonaws.com/role-arn` annotation:

```sh
$ kubectl -n kube-system get sa ebs-csi-controller-sa -ojsonpath='{.metadata.annotations.eks\.amazonaws\.com/role-arn}'
arn:aws:iam::<aws_account_id>:role/<new_role_name>
```

Restart the CSI controller:

```sh
kubectl -n kube-system rollout restart deploy/ebs-csi-controller
```

For more information on IAM roles for K8s service account, see AWS documentation
[here](https://docs.aws.amazon.com/eks/latest/userguide/iam-roles-for-service-accounts.html).
