// +build !ignore_autogenerated

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

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1alpha1

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	wardle "k8s.io/sample-apiserver/pkg/apis/wardle"
	unsafe "unsafe"
)

func init() {
	SchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_Flunder_To_wardle_Flunder,
		Convert_wardle_Flunder_To_v1alpha1_Flunder,
		Convert_v1alpha1_FlunderList_To_wardle_FlunderList,
		Convert_wardle_FlunderList_To_v1alpha1_FlunderList,
		Convert_v1alpha1_FlunderSpec_To_wardle_FlunderSpec,
		Convert_wardle_FlunderSpec_To_v1alpha1_FlunderSpec,
		Convert_v1alpha1_FlunderStatus_To_wardle_FlunderStatus,
		Convert_wardle_FlunderStatus_To_v1alpha1_FlunderStatus,
	)
}

func autoConvert_v1alpha1_Flunder_To_wardle_Flunder(in *Flunder, out *wardle.Flunder, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_FlunderSpec_To_wardle_FlunderSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_FlunderStatus_To_wardle_FlunderStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_Flunder_To_wardle_Flunder is an autogenerated conversion function.
func Convert_v1alpha1_Flunder_To_wardle_Flunder(in *Flunder, out *wardle.Flunder, s conversion.Scope) error {
	return autoConvert_v1alpha1_Flunder_To_wardle_Flunder(in, out, s)
}

func autoConvert_wardle_Flunder_To_v1alpha1_Flunder(in *wardle.Flunder, out *Flunder, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_wardle_FlunderSpec_To_v1alpha1_FlunderSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_wardle_FlunderStatus_To_v1alpha1_FlunderStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_wardle_Flunder_To_v1alpha1_Flunder is an autogenerated conversion function.
func Convert_wardle_Flunder_To_v1alpha1_Flunder(in *wardle.Flunder, out *Flunder, s conversion.Scope) error {
	return autoConvert_wardle_Flunder_To_v1alpha1_Flunder(in, out, s)
}

func autoConvert_v1alpha1_FlunderList_To_wardle_FlunderList(in *FlunderList, out *wardle.FlunderList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]wardle.Flunder)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_FlunderList_To_wardle_FlunderList is an autogenerated conversion function.
func Convert_v1alpha1_FlunderList_To_wardle_FlunderList(in *FlunderList, out *wardle.FlunderList, s conversion.Scope) error {
	return autoConvert_v1alpha1_FlunderList_To_wardle_FlunderList(in, out, s)
}

func autoConvert_wardle_FlunderList_To_v1alpha1_FlunderList(in *wardle.FlunderList, out *FlunderList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items == nil {
		out.Items = make([]Flunder, 0)
	} else {
		out.Items = *(*[]Flunder)(unsafe.Pointer(&in.Items))
	}
	return nil
}

// Convert_wardle_FlunderList_To_v1alpha1_FlunderList is an autogenerated conversion function.
func Convert_wardle_FlunderList_To_v1alpha1_FlunderList(in *wardle.FlunderList, out *FlunderList, s conversion.Scope) error {
	return autoConvert_wardle_FlunderList_To_v1alpha1_FlunderList(in, out, s)
}

func autoConvert_v1alpha1_FlunderSpec_To_wardle_FlunderSpec(in *FlunderSpec, out *wardle.FlunderSpec, s conversion.Scope) error {
	return nil
}

// Convert_v1alpha1_FlunderSpec_To_wardle_FlunderSpec is an autogenerated conversion function.
func Convert_v1alpha1_FlunderSpec_To_wardle_FlunderSpec(in *FlunderSpec, out *wardle.FlunderSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_FlunderSpec_To_wardle_FlunderSpec(in, out, s)
}

func autoConvert_wardle_FlunderSpec_To_v1alpha1_FlunderSpec(in *wardle.FlunderSpec, out *FlunderSpec, s conversion.Scope) error {
	return nil
}

// Convert_wardle_FlunderSpec_To_v1alpha1_FlunderSpec is an autogenerated conversion function.
func Convert_wardle_FlunderSpec_To_v1alpha1_FlunderSpec(in *wardle.FlunderSpec, out *FlunderSpec, s conversion.Scope) error {
	return autoConvert_wardle_FlunderSpec_To_v1alpha1_FlunderSpec(in, out, s)
}

func autoConvert_v1alpha1_FlunderStatus_To_wardle_FlunderStatus(in *FlunderStatus, out *wardle.FlunderStatus, s conversion.Scope) error {
	return nil
}

// Convert_v1alpha1_FlunderStatus_To_wardle_FlunderStatus is an autogenerated conversion function.
func Convert_v1alpha1_FlunderStatus_To_wardle_FlunderStatus(in *FlunderStatus, out *wardle.FlunderStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_FlunderStatus_To_wardle_FlunderStatus(in, out, s)
}

func autoConvert_wardle_FlunderStatus_To_v1alpha1_FlunderStatus(in *wardle.FlunderStatus, out *FlunderStatus, s conversion.Scope) error {
	return nil
}

// Convert_wardle_FlunderStatus_To_v1alpha1_FlunderStatus is an autogenerated conversion function.
func Convert_wardle_FlunderStatus_To_v1alpha1_FlunderStatus(in *wardle.FlunderStatus, out *FlunderStatus, s conversion.Scope) error {
	return autoConvert_wardle_FlunderStatus_To_v1alpha1_FlunderStatus(in, out, s)
}
