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

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	testapigroup "k8s.io/apimachinery/pkg/apis/testapigroup"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*Carp)(nil), (*testapigroup.Carp)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Carp_To_testapigroup_Carp(a.(*Carp), b.(*testapigroup.Carp), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*testapigroup.Carp)(nil), (*Carp)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_testapigroup_Carp_To_v1_Carp(a.(*testapigroup.Carp), b.(*Carp), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*CarpCondition)(nil), (*testapigroup.CarpCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_CarpCondition_To_testapigroup_CarpCondition(a.(*CarpCondition), b.(*testapigroup.CarpCondition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*testapigroup.CarpCondition)(nil), (*CarpCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_testapigroup_CarpCondition_To_v1_CarpCondition(a.(*testapigroup.CarpCondition), b.(*CarpCondition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*CarpList)(nil), (*testapigroup.CarpList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_CarpList_To_testapigroup_CarpList(a.(*CarpList), b.(*testapigroup.CarpList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*testapigroup.CarpList)(nil), (*CarpList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_testapigroup_CarpList_To_v1_CarpList(a.(*testapigroup.CarpList), b.(*CarpList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*CarpSpec)(nil), (*testapigroup.CarpSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_CarpSpec_To_testapigroup_CarpSpec(a.(*CarpSpec), b.(*testapigroup.CarpSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*testapigroup.CarpSpec)(nil), (*CarpSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_testapigroup_CarpSpec_To_v1_CarpSpec(a.(*testapigroup.CarpSpec), b.(*CarpSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*CarpStatus)(nil), (*testapigroup.CarpStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_CarpStatus_To_testapigroup_CarpStatus(a.(*CarpStatus), b.(*testapigroup.CarpStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*testapigroup.CarpStatus)(nil), (*CarpStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_testapigroup_CarpStatus_To_v1_CarpStatus(a.(*testapigroup.CarpStatus), b.(*CarpStatus), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1_Carp_To_testapigroup_Carp(in *Carp, out *testapigroup.Carp, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_CarpSpec_To_testapigroup_CarpSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_CarpStatus_To_testapigroup_CarpStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_Carp_To_testapigroup_Carp is an autogenerated conversion function.
func Convert_v1_Carp_To_testapigroup_Carp(in *Carp, out *testapigroup.Carp, s conversion.Scope) error {
	return autoConvert_v1_Carp_To_testapigroup_Carp(in, out, s)
}

func autoConvert_testapigroup_Carp_To_v1_Carp(in *testapigroup.Carp, out *Carp, s conversion.Scope) error {
	// Auto-generated external APIVersion/Kind
	// Disable with a `+k8s:conversion-gen:set-api-version-kind=false` comment
	// Customize with `+groupName`, `+version`, or `+kind` comments
	out.APIVersion = "testapigroup.apimachinery.k8s.io/v1"
	out.Kind = "Carp"

	out.ObjectMeta = in.ObjectMeta
	if err := Convert_testapigroup_CarpSpec_To_v1_CarpSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_testapigroup_CarpStatus_To_v1_CarpStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_testapigroup_Carp_To_v1_Carp is an autogenerated conversion function.
func Convert_testapigroup_Carp_To_v1_Carp(in *testapigroup.Carp, out *Carp, s conversion.Scope) error {
	return autoConvert_testapigroup_Carp_To_v1_Carp(in, out, s)
}

func autoConvert_v1_CarpCondition_To_testapigroup_CarpCondition(in *CarpCondition, out *testapigroup.CarpCondition, s conversion.Scope) error {
	out.Type = testapigroup.CarpConditionType(in.Type)
	out.Status = testapigroup.ConditionStatus(in.Status)
	out.LastProbeTime = in.LastProbeTime
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_v1_CarpCondition_To_testapigroup_CarpCondition is an autogenerated conversion function.
func Convert_v1_CarpCondition_To_testapigroup_CarpCondition(in *CarpCondition, out *testapigroup.CarpCondition, s conversion.Scope) error {
	return autoConvert_v1_CarpCondition_To_testapigroup_CarpCondition(in, out, s)
}

func autoConvert_testapigroup_CarpCondition_To_v1_CarpCondition(in *testapigroup.CarpCondition, out *CarpCondition, s conversion.Scope) error {
	out.Type = CarpConditionType(in.Type)
	out.Status = ConditionStatus(in.Status)
	out.LastProbeTime = in.LastProbeTime
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_testapigroup_CarpCondition_To_v1_CarpCondition is an autogenerated conversion function.
func Convert_testapigroup_CarpCondition_To_v1_CarpCondition(in *testapigroup.CarpCondition, out *CarpCondition, s conversion.Scope) error {
	return autoConvert_testapigroup_CarpCondition_To_v1_CarpCondition(in, out, s)
}

func autoConvert_v1_CarpList_To_testapigroup_CarpList(in *CarpList, out *testapigroup.CarpList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]testapigroup.Carp, len(*in))
		for i := range *in {
			if err := Convert_v1_Carp_To_testapigroup_Carp(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1_CarpList_To_testapigroup_CarpList is an autogenerated conversion function.
func Convert_v1_CarpList_To_testapigroup_CarpList(in *CarpList, out *testapigroup.CarpList, s conversion.Scope) error {
	return autoConvert_v1_CarpList_To_testapigroup_CarpList(in, out, s)
}

func autoConvert_testapigroup_CarpList_To_v1_CarpList(in *testapigroup.CarpList, out *CarpList, s conversion.Scope) error {
	// Auto-generated external APIVersion/Kind
	// Disable with a `+k8s:conversion-gen:set-api-version-kind=false` comment
	// Customize with `+groupName`, `+version`, or `+kind` comments
	out.APIVersion = "testapigroup.apimachinery.k8s.io/v1"
	out.Kind = "CarpList"

	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Carp, len(*in))
		for i := range *in {
			if err := Convert_testapigroup_Carp_To_v1_Carp(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_testapigroup_CarpList_To_v1_CarpList is an autogenerated conversion function.
func Convert_testapigroup_CarpList_To_v1_CarpList(in *testapigroup.CarpList, out *CarpList, s conversion.Scope) error {
	return autoConvert_testapigroup_CarpList_To_v1_CarpList(in, out, s)
}

func autoConvert_v1_CarpSpec_To_testapigroup_CarpSpec(in *CarpSpec, out *testapigroup.CarpSpec, s conversion.Scope) error {
	out.RestartPolicy = testapigroup.RestartPolicy(in.RestartPolicy)
	out.TerminationGracePeriodSeconds = (*int64)(unsafe.Pointer(in.TerminationGracePeriodSeconds))
	out.ActiveDeadlineSeconds = (*int64)(unsafe.Pointer(in.ActiveDeadlineSeconds))
	out.NodeSelector = *(*map[string]string)(unsafe.Pointer(&in.NodeSelector))
	out.ServiceAccountName = in.ServiceAccountName
	// INFO: in.DeprecatedServiceAccount opted out of conversion generation
	out.NodeName = in.NodeName
	// INFO: in.HostNetwork opted out of conversion generation
	// INFO: in.HostPID opted out of conversion generation
	// INFO: in.HostIPC opted out of conversion generation
	out.Hostname = in.Hostname
	out.Subdomain = in.Subdomain
	out.SchedulerName = in.SchedulerName
	return nil
}

// Convert_v1_CarpSpec_To_testapigroup_CarpSpec is an autogenerated conversion function.
func Convert_v1_CarpSpec_To_testapigroup_CarpSpec(in *CarpSpec, out *testapigroup.CarpSpec, s conversion.Scope) error {
	return autoConvert_v1_CarpSpec_To_testapigroup_CarpSpec(in, out, s)
}

func autoConvert_testapigroup_CarpSpec_To_v1_CarpSpec(in *testapigroup.CarpSpec, out *CarpSpec, s conversion.Scope) error {
	out.RestartPolicy = RestartPolicy(in.RestartPolicy)
	out.TerminationGracePeriodSeconds = (*int64)(unsafe.Pointer(in.TerminationGracePeriodSeconds))
	out.ActiveDeadlineSeconds = (*int64)(unsafe.Pointer(in.ActiveDeadlineSeconds))
	out.NodeSelector = *(*map[string]string)(unsafe.Pointer(&in.NodeSelector))
	out.ServiceAccountName = in.ServiceAccountName
	out.NodeName = in.NodeName
	out.Hostname = in.Hostname
	out.Subdomain = in.Subdomain
	out.SchedulerName = in.SchedulerName
	return nil
}

// Convert_testapigroup_CarpSpec_To_v1_CarpSpec is an autogenerated conversion function.
func Convert_testapigroup_CarpSpec_To_v1_CarpSpec(in *testapigroup.CarpSpec, out *CarpSpec, s conversion.Scope) error {
	return autoConvert_testapigroup_CarpSpec_To_v1_CarpSpec(in, out, s)
}

func autoConvert_v1_CarpStatus_To_testapigroup_CarpStatus(in *CarpStatus, out *testapigroup.CarpStatus, s conversion.Scope) error {
	out.Phase = testapigroup.CarpPhase(in.Phase)
	out.Conditions = *(*[]testapigroup.CarpCondition)(unsafe.Pointer(&in.Conditions))
	out.Message = in.Message
	out.Reason = in.Reason
	out.HostIP = in.HostIP
	out.CarpIP = in.CarpIP
	out.StartTime = (*metav1.Time)(unsafe.Pointer(in.StartTime))
	return nil
}

// Convert_v1_CarpStatus_To_testapigroup_CarpStatus is an autogenerated conversion function.
func Convert_v1_CarpStatus_To_testapigroup_CarpStatus(in *CarpStatus, out *testapigroup.CarpStatus, s conversion.Scope) error {
	return autoConvert_v1_CarpStatus_To_testapigroup_CarpStatus(in, out, s)
}

func autoConvert_testapigroup_CarpStatus_To_v1_CarpStatus(in *testapigroup.CarpStatus, out *CarpStatus, s conversion.Scope) error {
	out.Phase = CarpPhase(in.Phase)
	out.Conditions = *(*[]CarpCondition)(unsafe.Pointer(&in.Conditions))
	out.Message = in.Message
	out.Reason = in.Reason
	out.HostIP = in.HostIP
	out.CarpIP = in.CarpIP
	out.StartTime = (*metav1.Time)(unsafe.Pointer(in.StartTime))
	return nil
}

// Convert_testapigroup_CarpStatus_To_v1_CarpStatus is an autogenerated conversion function.
func Convert_testapigroup_CarpStatus_To_v1_CarpStatus(in *testapigroup.CarpStatus, out *CarpStatus, s conversion.Scope) error {
	return autoConvert_testapigroup_CarpStatus_To_v1_CarpStatus(in, out, s)
}
