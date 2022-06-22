/*
Copyright 2022 The Kubernetes Authors.

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

// VolumeSnapshotDelta is a specification for a VolumeSnapshotDelta resource
type VolumeSnapshotDelta struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec VolumeSnapshotDeltaSpec `json:"spec"`
	// +optional
	Status VolumeSnapshotDeltaStatus `json:"status,omitempty"`
}

// VolumeSnapshotDeltaSpec is the spec for a VolumeSnapshotDelta resource
type VolumeSnapshotDeltaSpec struct {
	// If BaseVolumeSnapshotName is not specified, return all used blocks.
	BaseVolumeSnapshotName   string `json:"baseVolumeSnapshotName,omitempty"` // Base VolumeSnapshot, optional.
	TargetVolumeSnapshotName string `json:"targetVolumeSnapshotName"`         // Target VolumeSnapshot, required.
	Mode                     string `json:"mode,omitempty"`                   // default 'Block'
	StartOffset              string `json:"startOffset,omitempty"`            // Logical offset from beginning of disk/volume.
	// Use string instead of uint64 to give vendor
	// the flexibility of implementing it either
	// string "token" or a number.
	MaxEntries uint64            `json:"maxEntries"`           // Maximum number of entries in the response
	Parameters map[string]string `json:"parameters,omitempty"` // Vendor specific parameters passed in as opaque key-value pairs.  Optional.
}

// VolumeSnapshotDeltaStatus is the status for a VolumeSnapshotDelta resource
type VolumeSnapshotDeltaStatus struct {
	State      string `json:"state"`
	Error      string `json:"error,omitempty"`
	StreamURL  string `json:"streamURL,omitempty"`
	VolumeSize uint64 `json:"volumeSize"` // size of volume in bytes
	Timeout    uint64 `json:"timeout"`    //second since epoch
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VolumeSnapshotDeltaList is a list of VolumeSnapshotDelta resources
type VolumeSnapshotDeltaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []VolumeSnapshotDelta `json:"items"`
}
