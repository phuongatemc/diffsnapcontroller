package validation

import (
	"example.com/differentialsnapshot/pkg/apis/cbt"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

func ValidateChangedBlock(c *cbt.ChangedBlock) field.ErrorList {
	return ValidateChangedBlockStatus(&c.Status, field.NewPath("status"))
}

func ValidateChangedBlockStatus(c *cbt.ChangedBlockStatus, f *field.Path) field.ErrorList {
	return field.ErrorList{}
}
