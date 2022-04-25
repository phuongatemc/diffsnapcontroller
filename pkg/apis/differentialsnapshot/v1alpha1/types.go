/*
Copyright 2017 The Kubernetes Authors.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GetChangedBlocks is a specification for a GetChangedBlocks resource
type GetChangedBlocks struct {
	metav1.TypeMeta   `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GetChangedBlocksSpec   `json:"spec"`
	// +optional
	Status GetChangedBlocksStatus `json:"status,omitempty"`
}

// GetChangedBlocksSpec is the spec for a GetChangedBlocks resource
type GetChangedBlocksSpec struct {
	// If SnapshotBase is not specified, return all used blocks.
	SnapshotBase   string `json:"snapshotBase,omitempty"` // Snapshot handle, optional.
	SnapshotTarget string `json:"snapshotTarget"`         // Snapshot handle, required.
	VolumeId       string `json:"volumeId,omitempty"`     // optional
	StartOffset    string `json:"startOffset,omitempty"`  // Logical offset from beginning of disk/volume.
	// Use string instead of uint64 to give vendor
	// the flexibility of implementing it either
	// string "token" or a number.
	MaxEntries uint64            `json:"maxEntries"`           // Maximum number of entries in the response
	Secrets    map[string]string `json:"secrets,omitempty"`    // Secrets required by Vendor to access snapshots.  Optional.
	Parameters map[string]string `json:"parameters,omitempty"` // Vendor specific parameters passed in as opaque key-value pairs.  Optional.
}

// GetChangedBlocksStatus is the status for a GetChangedBlocks resource
type GetChangedBlocksStatus struct {
	State           string         `json:"state"`
	Error           string         `json:"error,omitempty"`
	ChangeBlockList []ChangedBlock `json:"changeBlockList"`      //array of ChangedBlock
	NextOffset      string         `json:"nextOffset,omitempty"` // StartOffset of the next “page”.
	VolumeSize      uint64         `json:"volumeSize"`           // size of volume in bytes
	Timeout         uint64         `json:"timeout"`              //second since epoch
}

type ChangedBlock struct {
	Offset  uint64 `json:"offset"`            // logical offset
	Size    uint64 `json:"size"`              // size of the block data
	Context []byte `json:"context,omitempty"` // additional vendor specific info.  Optional.
	ZeroOut bool   `json:"zeroOut"`           // If ZeroOut is true, this block in SnapshotTarget is zero out.
	// This is for optimization to avoid data mover to transfer zero blocks.
	// Not all vendors support this zeroout.
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GetChangedBlocksList is a list of GetChangedBlocks resources
type GetChangedBlocksList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []GetChangedBlocks `json:"items"`
}
