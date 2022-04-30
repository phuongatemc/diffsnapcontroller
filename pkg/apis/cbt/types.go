package cbt

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ChangedBlock contains information about the blocks that are different
// between two snapshots.
type ChangedBlock struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Status ChangedBlockStatus
}

// ChangedBlockStatus is a changed block status.
type ChangedBlockStatus struct {
	Offset  uint64 // logical offset
	Size    uint64 // size of the block data
	Context []byte // additional vendor specific info.  Optional.
	ZeroOut bool   // If ZeroOut is true, this block in SnapshotTarget is zero out.
}

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ChangedBlockList is a list of ChangedBlock objects.
type ChangedBlockList struct {
	metav1.TypeMeta
	metav1.ListMeta

	Items []ChangedBlock
}
