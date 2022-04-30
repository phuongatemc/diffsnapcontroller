/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	cbt "example.com/differentialsnapshot/pkg/apis/cbt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeChangedBlocks implements ChangedBlockInterface
type FakeChangedBlocks struct {
	Fake *FakeCbt
	ns   string
}

var changedblocksResource = schema.GroupVersionResource{Group: "cbt.example.com", Version: "", Resource: "changedblocks"}

var changedblocksKind = schema.GroupVersionKind{Group: "cbt.example.com", Version: "", Kind: "ChangedBlock"}

// Get takes name of the changedBlock, and returns the corresponding changedBlock object, and an error if there is any.
func (c *FakeChangedBlocks) Get(ctx context.Context, name string, options v1.GetOptions) (result *cbt.ChangedBlock, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(changedblocksResource, c.ns, name), &cbt.ChangedBlock{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cbt.ChangedBlock), err
}

// List takes label and field selectors, and returns the list of ChangedBlocks that match those selectors.
func (c *FakeChangedBlocks) List(ctx context.Context, opts v1.ListOptions) (result *cbt.ChangedBlockList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(changedblocksResource, changedblocksKind, c.ns, opts), &cbt.ChangedBlockList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &cbt.ChangedBlockList{ListMeta: obj.(*cbt.ChangedBlockList).ListMeta}
	for _, item := range obj.(*cbt.ChangedBlockList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested changedBlocks.
func (c *FakeChangedBlocks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(changedblocksResource, c.ns, opts))

}

// Create takes the representation of a changedBlock and creates it.  Returns the server's representation of the changedBlock, and an error, if there is any.
func (c *FakeChangedBlocks) Create(ctx context.Context, changedBlock *cbt.ChangedBlock, opts v1.CreateOptions) (result *cbt.ChangedBlock, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(changedblocksResource, c.ns, changedBlock), &cbt.ChangedBlock{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cbt.ChangedBlock), err
}

// Update takes the representation of a changedBlock and updates it. Returns the server's representation of the changedBlock, and an error, if there is any.
func (c *FakeChangedBlocks) Update(ctx context.Context, changedBlock *cbt.ChangedBlock, opts v1.UpdateOptions) (result *cbt.ChangedBlock, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(changedblocksResource, c.ns, changedBlock), &cbt.ChangedBlock{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cbt.ChangedBlock), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeChangedBlocks) UpdateStatus(ctx context.Context, changedBlock *cbt.ChangedBlock, opts v1.UpdateOptions) (*cbt.ChangedBlock, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(changedblocksResource, "status", c.ns, changedBlock), &cbt.ChangedBlock{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cbt.ChangedBlock), err
}

// Delete takes name of the changedBlock and deletes it. Returns an error if one occurs.
func (c *FakeChangedBlocks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(changedblocksResource, c.ns, name, opts), &cbt.ChangedBlock{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeChangedBlocks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(changedblocksResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &cbt.ChangedBlockList{})
	return err
}

// Patch applies the patch and returns the patched changedBlock.
func (c *FakeChangedBlocks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *cbt.ChangedBlock, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(changedblocksResource, c.ns, name, pt, data, subresources...), &cbt.ChangedBlock{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cbt.ChangedBlock), err
}
