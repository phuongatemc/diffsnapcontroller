package storage

import (
	"example.com/differentialsnapshot/pkg/apis/cbt/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var mocks = map[string]*v1alpha1.ChangedBlockList{
	"token-00": &v1alpha1.ChangedBlockList{
		metav1.TypeMeta{},
		metav1.ListMeta{},
		[]v1alpha1.ChangedBlock{
			v1alpha1.ChangedBlock{
				metav1.TypeMeta{},
				metav1.ObjectMeta{},
				v1alpha1.ChangedBlockStatus{
					Offset: uint64(1048576),
					Size:   uint64(260096),
				},
			},
			v1alpha1.ChangedBlock{
				metav1.TypeMeta{},
				metav1.ObjectMeta{},
				v1alpha1.ChangedBlockStatus{
					Offset: uint64(1048576),
					Size:   uint64(260096),
				},
			},
			v1alpha1.ChangedBlock{
				metav1.TypeMeta{},
				metav1.ObjectMeta{},
				v1alpha1.ChangedBlockStatus{
					Offset: uint64(1048576),
					Size:   uint64(260096),
				},
			},
		},
	},
	"token-01": &v1alpha1.ChangedBlockList{
		metav1.TypeMeta{},
		metav1.ListMeta{},
		[]v1alpha1.ChangedBlock{
			v1alpha1.ChangedBlock{
				metav1.TypeMeta{},
				metav1.ObjectMeta{},
				v1alpha1.ChangedBlockStatus{
					Offset: uint64(1048576),
					Size:   uint64(260096),
				},
			},
			v1alpha1.ChangedBlock{
				metav1.TypeMeta{},
				metav1.ObjectMeta{},
				v1alpha1.ChangedBlockStatus{
					Offset: uint64(1048576),
					Size:   uint64(260096),
				},
			},
			v1alpha1.ChangedBlock{
				metav1.TypeMeta{},
				metav1.ObjectMeta{},
				v1alpha1.ChangedBlockStatus{
					Offset: uint64(1048576),
					Size:   uint64(260096),
				},
			},
		},
	},
	"token-02": &v1alpha1.ChangedBlockList{
		metav1.TypeMeta{},
		metav1.ListMeta{},
		[]v1alpha1.ChangedBlock{
			v1alpha1.ChangedBlock{
				metav1.TypeMeta{},
				metav1.ObjectMeta{},
				v1alpha1.ChangedBlockStatus{
					Offset: uint64(1048576),
					Size:   uint64(260096),
				},
			},
			v1alpha1.ChangedBlock{
				metav1.TypeMeta{},
				metav1.ObjectMeta{},
				v1alpha1.ChangedBlockStatus{
					Offset: uint64(1048576),
					Size:   uint64(260096),
				},
			},
			v1alpha1.ChangedBlock{
				metav1.TypeMeta{},
				metav1.ObjectMeta{},
				v1alpha1.ChangedBlockStatus{
					Offset: uint64(1048576),
					Size:   uint64(260096),
				},
			},
		},
	},
}
