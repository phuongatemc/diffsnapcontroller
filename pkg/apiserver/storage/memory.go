package storage

import (
	"context"

	"example.com/differentialsnapshot/pkg/apis/cbt"
	"example.com/differentialsnapshot/pkg/apis/cbt/v1alpha1"
	metainternalversion "k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
)

type Memory struct{}

func NewMemory() Memory {
	return Memory{}
}

// New returns an empty object that can be used with Create and Update after request data has been put into it.
// This object must be a pointer type for use with Codec.DecodeInto([]byte, runtime.Object)
func (m *Memory) New() runtime.Object {
	return &cbt.ChangedBlock{}
}

// NamespaceScoped returns true if the storage is namespaced
func (m *Memory) NamespaceScoped() bool {
	return true
}

// Get finds a resource in the storage by name and returns it.
// Although it can return an arbitrary error value, IsNotFound(err) is true for the
// returned error value err when the specified resource is not found.
func (m *Memory) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	return &v1alpha1.ChangedBlock{}, nil
}

// NewList returns an empty object that can be used with the List call.
// This object must be a pointer type for use with Codec.DecodeInto([]byte, runtime.Object)
func (m *Memory) NewList() runtime.Object {
	return &cbt.ChangedBlockList{}
}

// List selects resources in the storage which match to the selector. 'options' can be nil.
func (m *Memory) List(ctx context.Context, options *metainternalversion.ListOptions) (runtime.Object, error) {
	var (
		next      string
		remaining int64
	)

	switch options.Continue {
	case "token-01":
		next = "token-02"
		remaining = 3
	case "token-02":
		next = ""
		remaining = 0
	default:
		next = "token-01"
		remaining = 6
	}

	list := &cbt.ChangedBlockList{}
	list.RemainingItemCount = &remaining
	if remaining > 0 {
		list.Continue = next
	}

	return list, nil
}

// 'label' selects on labels; 'field' selects on the object's fields. Not all fields
// are supported; an error should be returned if 'field' tries to select on a field that
// isn't supported. 'resourceVersion' allows for continuing/starting a watch at a
// particular version.
func (m *Memory) Watch(ctx context.Context, options *metainternalversion.ListOptions) (watch.Interface, error) {
	return watch.NewFake(), nil
}
