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

package v1

import (
	unsafe "unsafe"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/node/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/kubernetes/pkg/apis/core"
	node "k8s.io/kubernetes/pkg/apis/node"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*v1.Overhead)(nil), (*node.Overhead)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Overhead_To_node_Overhead(a.(*v1.Overhead), b.(*node.Overhead), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*node.Overhead)(nil), (*v1.Overhead)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_node_Overhead_To_v1_Overhead(a.(*node.Overhead), b.(*v1.Overhead), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.RuntimeClass)(nil), (*node.RuntimeClass)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RuntimeClass_To_node_RuntimeClass(a.(*v1.RuntimeClass), b.(*node.RuntimeClass), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*node.RuntimeClass)(nil), (*v1.RuntimeClass)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_node_RuntimeClass_To_v1_RuntimeClass(a.(*node.RuntimeClass), b.(*v1.RuntimeClass), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.RuntimeClassList)(nil), (*node.RuntimeClassList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RuntimeClassList_To_node_RuntimeClassList(a.(*v1.RuntimeClassList), b.(*node.RuntimeClassList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*node.RuntimeClassList)(nil), (*v1.RuntimeClassList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_node_RuntimeClassList_To_v1_RuntimeClassList(a.(*node.RuntimeClassList), b.(*v1.RuntimeClassList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.Scheduling)(nil), (*node.Scheduling)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Scheduling_To_node_Scheduling(a.(*v1.Scheduling), b.(*node.Scheduling), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*node.Scheduling)(nil), (*v1.Scheduling)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_node_Scheduling_To_v1_Scheduling(a.(*node.Scheduling), b.(*v1.Scheduling), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1_Overhead_To_node_Overhead(in *v1.Overhead, out *node.Overhead, s conversion.Scope) error {
	out.PodFixed = *(*core.ResourceList)(unsafe.Pointer(&in.PodFixed))
	return nil
}

// Convert_v1_Overhead_To_node_Overhead is an autogenerated conversion function.
func Convert_v1_Overhead_To_node_Overhead(in *v1.Overhead, out *node.Overhead, s conversion.Scope) error {
	return autoConvert_v1_Overhead_To_node_Overhead(in, out, s)
}

func autoConvert_node_Overhead_To_v1_Overhead(in *node.Overhead, out *v1.Overhead, s conversion.Scope) error {
	out.PodFixed = *(*corev1.ResourceList)(unsafe.Pointer(&in.PodFixed))
	return nil
}

// Convert_node_Overhead_To_v1_Overhead is an autogenerated conversion function.
func Convert_node_Overhead_To_v1_Overhead(in *node.Overhead, out *v1.Overhead, s conversion.Scope) error {
	return autoConvert_node_Overhead_To_v1_Overhead(in, out, s)
}

func autoConvert_v1_RuntimeClass_To_node_RuntimeClass(in *v1.RuntimeClass, out *node.RuntimeClass, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Handler = in.Handler
	out.Overhead = (*node.Overhead)(unsafe.Pointer(in.Overhead))
	out.Scheduling = (*node.Scheduling)(unsafe.Pointer(in.Scheduling))
	return nil
}

// Convert_v1_RuntimeClass_To_node_RuntimeClass is an autogenerated conversion function.
func Convert_v1_RuntimeClass_To_node_RuntimeClass(in *v1.RuntimeClass, out *node.RuntimeClass, s conversion.Scope) error {
	return autoConvert_v1_RuntimeClass_To_node_RuntimeClass(in, out, s)
}

func autoConvert_node_RuntimeClass_To_v1_RuntimeClass(in *node.RuntimeClass, out *v1.RuntimeClass, s conversion.Scope) error {
	// Auto-generated external APIVersion/Kind
	// Disable with a `+k8s:conversion-gen:set-api-version-kind=false` comment
	// Customize with `+groupName`, `+version`, or `+kind` comments
	out.APIVersion = "node.k8s.io/v1"
	out.Kind = "RuntimeClass"

	out.ObjectMeta = in.ObjectMeta
	out.Handler = in.Handler
	out.Overhead = (*v1.Overhead)(unsafe.Pointer(in.Overhead))
	out.Scheduling = (*v1.Scheduling)(unsafe.Pointer(in.Scheduling))
	return nil
}

// Convert_node_RuntimeClass_To_v1_RuntimeClass is an autogenerated conversion function.
func Convert_node_RuntimeClass_To_v1_RuntimeClass(in *node.RuntimeClass, out *v1.RuntimeClass, s conversion.Scope) error {
	return autoConvert_node_RuntimeClass_To_v1_RuntimeClass(in, out, s)
}

func autoConvert_v1_RuntimeClassList_To_node_RuntimeClassList(in *v1.RuntimeClassList, out *node.RuntimeClassList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]node.RuntimeClass, len(*in))
		for i := range *in {
			if err := Convert_v1_RuntimeClass_To_node_RuntimeClass(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1_RuntimeClassList_To_node_RuntimeClassList is an autogenerated conversion function.
func Convert_v1_RuntimeClassList_To_node_RuntimeClassList(in *v1.RuntimeClassList, out *node.RuntimeClassList, s conversion.Scope) error {
	return autoConvert_v1_RuntimeClassList_To_node_RuntimeClassList(in, out, s)
}

func autoConvert_node_RuntimeClassList_To_v1_RuntimeClassList(in *node.RuntimeClassList, out *v1.RuntimeClassList, s conversion.Scope) error {
	// Auto-generated external APIVersion/Kind
	// Disable with a `+k8s:conversion-gen:set-api-version-kind=false` comment
	// Customize with `+groupName`, `+version`, or `+kind` comments
	out.APIVersion = "node.k8s.io/v1"
	out.Kind = "RuntimeClassList"

	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1.RuntimeClass, len(*in))
		for i := range *in {
			if err := Convert_node_RuntimeClass_To_v1_RuntimeClass(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_node_RuntimeClassList_To_v1_RuntimeClassList is an autogenerated conversion function.
func Convert_node_RuntimeClassList_To_v1_RuntimeClassList(in *node.RuntimeClassList, out *v1.RuntimeClassList, s conversion.Scope) error {
	return autoConvert_node_RuntimeClassList_To_v1_RuntimeClassList(in, out, s)
}

func autoConvert_v1_Scheduling_To_node_Scheduling(in *v1.Scheduling, out *node.Scheduling, s conversion.Scope) error {
	out.NodeSelector = *(*map[string]string)(unsafe.Pointer(&in.NodeSelector))
	out.Tolerations = *(*[]core.Toleration)(unsafe.Pointer(&in.Tolerations))
	return nil
}

// Convert_v1_Scheduling_To_node_Scheduling is an autogenerated conversion function.
func Convert_v1_Scheduling_To_node_Scheduling(in *v1.Scheduling, out *node.Scheduling, s conversion.Scope) error {
	return autoConvert_v1_Scheduling_To_node_Scheduling(in, out, s)
}

func autoConvert_node_Scheduling_To_v1_Scheduling(in *node.Scheduling, out *v1.Scheduling, s conversion.Scope) error {
	out.NodeSelector = *(*map[string]string)(unsafe.Pointer(&in.NodeSelector))
	out.Tolerations = *(*[]corev1.Toleration)(unsafe.Pointer(&in.Tolerations))
	return nil
}

// Convert_node_Scheduling_To_v1_Scheduling is an autogenerated conversion function.
func Convert_node_Scheduling_To_v1_Scheduling(in *node.Scheduling, out *v1.Scheduling, s conversion.Scope) error {
	return autoConvert_node_Scheduling_To_v1_Scheduling(in, out, s)
}
