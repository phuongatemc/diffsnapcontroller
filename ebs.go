package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ebs"
	"k8s.io/klog/v2"

	differentialsnapshotv1alpha1 "example.com/differentialsnapshot/pkg/apis/differentialsnapshot/v1alpha1"
)

func processGetChangedBlock(cEBS *ebs.EBS, getChangedBlocks *differentialsnapshotv1alpha1.GetChangedBlocks) (err error) {
	klog.Info("Processing GetChangedBlocks...")
	defer func() {
		if err != nil {
			klog.Info("Unable to GetChangedBlocks: %s.", err)
			getChangedBlocks.Status.State = "Failure"
			getChangedBlocks.Status.Error = fmt.Sprintf("Unable to get ChangedBlocks: %s.", err)
		}
	}()
	maxResult := int64(getChangedBlocks.Spec.MaxEntries)

	ebsListChangedBlocksInput := &ebs.ListChangedBlocksInput{
		FirstSnapshotId:  &getChangedBlocks.Spec.SnapshotBase,
		SecondSnapshotId: &getChangedBlocks.Spec.SnapshotTarget,
		MaxResults:       &maxResult,
		NextToken:        &getChangedBlocks.Spec.StartOffset,
	}
	klog.Info("Calling EBS ListChangedBlocks with input: ", printObject(ebsListChangedBlocksInput))
	listChangedBlocksOutput, err := cEBS.ListChangedBlocks(ebsListChangedBlocksInput)
	if err != nil {
		err = fmt.Errorf("EBS ListChangedBlocks failed %s.", err)
		return
	}
	klog.Info("EBS ListChangedBlocks return: ", printObject(listChangedBlocksOutput))
	blockSize := *listChangedBlocksOutput.BlockSize
	for _, ebsChangedBlock := range listChangedBlocksOutput.ChangedBlocks {
		var blockIndex int64
		if ebsChangedBlock.BlockIndex == nil {
			err = fmt.Errorf("Invalid BlockIndex in EBS ChangedBlock %s.", printObject(ebsChangedBlock))
		} else {
			blockIndex = *ebsChangedBlock.BlockIndex
		}
		changedBlock := differentialsnapshotv1alpha1.ChangedBlock{
			Offset: uint64(blockIndex * blockSize),
			Size:   uint64(blockSize),
		}
		getChangedBlocks.Status.ChangeBlockList = append(getChangedBlocks.Status.ChangeBlockList, changedBlock)
	}
	if listChangedBlocksOutput.ExpiryTime != nil {
		expiry := *listChangedBlocksOutput.ExpiryTime
		getChangedBlocks.Status.Timeout = uint64(expiry.Unix())
	}
	if listChangedBlocksOutput.NextToken != nil {
		getChangedBlocks.Status.NextOffset = *listChangedBlocksOutput.NextToken
	}
	getChangedBlocks.Status.State = "Success"
	klog.Info("processGetChangedBlocks completes successfully: ", printObject(getChangedBlocks))
	return
}

// takes AWS session options to create a new session
func getSession(options session.Options) (*session.Session, error) {
	sess, err := session.NewSessionWithOptions(options)
	if err != nil {
		klog.Info("NewSessionWithOptions failed: ", err)
		return nil, err
	}

	if _, err := sess.Config.Credentials.Get(); err != nil {
		klog.Info("Get Credentials failed: ", err)
		return nil, err
	}
	klog.Info("Successfully get aws Session.")
	return sess, nil
}

func newEBS() (EBS *ebs.EBS, err error) {
	region := "us-west-1"

	awsConfig := aws.NewConfig().WithRegion(region)
	sessionOptions := session.Options{Config: *awsConfig}
	sess, err := getSession(sessionOptions)
	if err != nil {
		return nil, err
	}
	klog.Info("Create EBS client...")
	EBS = ebs.New(sess)
	klog.Info("EBS client returns: ", printObject(EBS))
	return
}

func printObject(obj interface{}) string {
	buf, err := json.MarshalIndent(obj, "", "   ")
	if err == nil {
		return fmt.Sprintf("%s", buf)
	}
	return fmt.Sprintf("%v", obj)
}
