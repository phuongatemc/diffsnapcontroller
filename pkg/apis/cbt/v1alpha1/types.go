package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ChangedBlock contains information about the blocks that are different
// between two snapshots.
type ChangedBlock struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Status ChangedBlockStatus `json:"status,omitempty"`
}

// ChangedBlockStatus is a changed block status.
type ChangedBlockStatus struct {
	Offset  uint64 `json:"offset"`            // logical offset
	Size    uint64 `json:"size"`              // size of the block data
	Context []byte `json:"context,omitempty"` // additional vendor specific info.  Optional.
	ZeroOut bool   `json:"zeroOut"`           // If ZeroOut is true, this block in SnapshotTarget is zero out.
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ChangedBlockList is a list of ChangedBlock objects.
type ChangedBlockList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ChangedBlock `json:"items"`
}
