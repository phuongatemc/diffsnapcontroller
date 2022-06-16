//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	cbt "example.com/differentialsnapshot/pkg/apis/cbt"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*ChangedBlock)(nil), (*cbt.ChangedBlock)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ChangedBlock_To_cbt_ChangedBlock(a.(*ChangedBlock), b.(*cbt.ChangedBlock), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cbt.ChangedBlock)(nil), (*ChangedBlock)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_cbt_ChangedBlock_To_v1alpha1_ChangedBlock(a.(*cbt.ChangedBlock), b.(*ChangedBlock), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ChangedBlockList)(nil), (*cbt.ChangedBlockList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ChangedBlockList_To_cbt_ChangedBlockList(a.(*ChangedBlockList), b.(*cbt.ChangedBlockList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cbt.ChangedBlockList)(nil), (*ChangedBlockList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_cbt_ChangedBlockList_To_v1alpha1_ChangedBlockList(a.(*cbt.ChangedBlockList), b.(*ChangedBlockList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ChangedBlockStatus)(nil), (*cbt.ChangedBlockStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ChangedBlockStatus_To_cbt_ChangedBlockStatus(a.(*ChangedBlockStatus), b.(*cbt.ChangedBlockStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cbt.ChangedBlockStatus)(nil), (*ChangedBlockStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_cbt_ChangedBlockStatus_To_v1alpha1_ChangedBlockStatus(a.(*cbt.ChangedBlockStatus), b.(*ChangedBlockStatus), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_ChangedBlock_To_cbt_ChangedBlock(in *ChangedBlock, out *cbt.ChangedBlock, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_ChangedBlockStatus_To_cbt_ChangedBlockStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_ChangedBlock_To_cbt_ChangedBlock is an autogenerated conversion function.
func Convert_v1alpha1_ChangedBlock_To_cbt_ChangedBlock(in *ChangedBlock, out *cbt.ChangedBlock, s conversion.Scope) error {
	return autoConvert_v1alpha1_ChangedBlock_To_cbt_ChangedBlock(in, out, s)
}

func autoConvert_cbt_ChangedBlock_To_v1alpha1_ChangedBlock(in *cbt.ChangedBlock, out *ChangedBlock, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_cbt_ChangedBlockStatus_To_v1alpha1_ChangedBlockStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_cbt_ChangedBlock_To_v1alpha1_ChangedBlock is an autogenerated conversion function.
func Convert_cbt_ChangedBlock_To_v1alpha1_ChangedBlock(in *cbt.ChangedBlock, out *ChangedBlock, s conversion.Scope) error {
	return autoConvert_cbt_ChangedBlock_To_v1alpha1_ChangedBlock(in, out, s)
}

func autoConvert_v1alpha1_ChangedBlockList_To_cbt_ChangedBlockList(in *ChangedBlockList, out *cbt.ChangedBlockList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]cbt.ChangedBlock)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_ChangedBlockList_To_cbt_ChangedBlockList is an autogenerated conversion function.
func Convert_v1alpha1_ChangedBlockList_To_cbt_ChangedBlockList(in *ChangedBlockList, out *cbt.ChangedBlockList, s conversion.Scope) error {
	return autoConvert_v1alpha1_ChangedBlockList_To_cbt_ChangedBlockList(in, out, s)
}

func autoConvert_cbt_ChangedBlockList_To_v1alpha1_ChangedBlockList(in *cbt.ChangedBlockList, out *ChangedBlockList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]ChangedBlock)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_cbt_ChangedBlockList_To_v1alpha1_ChangedBlockList is an autogenerated conversion function.
func Convert_cbt_ChangedBlockList_To_v1alpha1_ChangedBlockList(in *cbt.ChangedBlockList, out *ChangedBlockList, s conversion.Scope) error {
	return autoConvert_cbt_ChangedBlockList_To_v1alpha1_ChangedBlockList(in, out, s)
}

func autoConvert_v1alpha1_ChangedBlockStatus_To_cbt_ChangedBlockStatus(in *ChangedBlockStatus, out *cbt.ChangedBlockStatus, s conversion.Scope) error {
	out.Offset = in.Offset
	out.Size = in.Size
	out.Context = *(*[]byte)(unsafe.Pointer(&in.Context))
	out.ZeroOut = in.ZeroOut
	return nil
}

// Convert_v1alpha1_ChangedBlockStatus_To_cbt_ChangedBlockStatus is an autogenerated conversion function.
func Convert_v1alpha1_ChangedBlockStatus_To_cbt_ChangedBlockStatus(in *ChangedBlockStatus, out *cbt.ChangedBlockStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_ChangedBlockStatus_To_cbt_ChangedBlockStatus(in, out, s)
}

func autoConvert_cbt_ChangedBlockStatus_To_v1alpha1_ChangedBlockStatus(in *cbt.ChangedBlockStatus, out *ChangedBlockStatus, s conversion.Scope) error {
	out.Offset = in.Offset
	out.Size = in.Size
	out.Context = *(*[]byte)(unsafe.Pointer(&in.Context))
	out.ZeroOut = in.ZeroOut
	return nil
}

// Convert_cbt_ChangedBlockStatus_To_v1alpha1_ChangedBlockStatus is an autogenerated conversion function.
func Convert_cbt_ChangedBlockStatus_To_v1alpha1_ChangedBlockStatus(in *cbt.ChangedBlockStatus, out *ChangedBlockStatus, s conversion.Scope) error {
	return autoConvert_cbt_ChangedBlockStatus_To_v1alpha1_ChangedBlockStatus(in, out, s)
}