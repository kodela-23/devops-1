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

	v1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	batch "k8s.io/kubernetes/pkg/apis/batch"
	core "k8s.io/kubernetes/pkg/apis/core"
	apiscorev1 "k8s.io/kubernetes/pkg/apis/core/v1"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*v1.CronJob)(nil), (*batch.CronJob)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_CronJob_To_batch_CronJob(a.(*v1.CronJob), b.(*batch.CronJob), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*batch.CronJob)(nil), (*v1.CronJob)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_batch_CronJob_To_v1_CronJob(a.(*batch.CronJob), b.(*v1.CronJob), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.CronJobList)(nil), (*batch.CronJobList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_CronJobList_To_batch_CronJobList(a.(*v1.CronJobList), b.(*batch.CronJobList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*batch.CronJobList)(nil), (*v1.CronJobList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_batch_CronJobList_To_v1_CronJobList(a.(*batch.CronJobList), b.(*v1.CronJobList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.CronJobSpec)(nil), (*batch.CronJobSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_CronJobSpec_To_batch_CronJobSpec(a.(*v1.CronJobSpec), b.(*batch.CronJobSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*batch.CronJobSpec)(nil), (*v1.CronJobSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_batch_CronJobSpec_To_v1_CronJobSpec(a.(*batch.CronJobSpec), b.(*v1.CronJobSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.CronJobStatus)(nil), (*batch.CronJobStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_CronJobStatus_To_batch_CronJobStatus(a.(*v1.CronJobStatus), b.(*batch.CronJobStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*batch.CronJobStatus)(nil), (*v1.CronJobStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_batch_CronJobStatus_To_v1_CronJobStatus(a.(*batch.CronJobStatus), b.(*v1.CronJobStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.Job)(nil), (*batch.Job)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Job_To_batch_Job(a.(*v1.Job), b.(*batch.Job), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*batch.Job)(nil), (*v1.Job)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_batch_Job_To_v1_Job(a.(*batch.Job), b.(*v1.Job), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.JobCondition)(nil), (*batch.JobCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_JobCondition_To_batch_JobCondition(a.(*v1.JobCondition), b.(*batch.JobCondition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*batch.JobCondition)(nil), (*v1.JobCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_batch_JobCondition_To_v1_JobCondition(a.(*batch.JobCondition), b.(*v1.JobCondition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.JobList)(nil), (*batch.JobList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_JobList_To_batch_JobList(a.(*v1.JobList), b.(*batch.JobList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*batch.JobList)(nil), (*v1.JobList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_batch_JobList_To_v1_JobList(a.(*batch.JobList), b.(*v1.JobList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.JobStatus)(nil), (*batch.JobStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_JobStatus_To_batch_JobStatus(a.(*v1.JobStatus), b.(*batch.JobStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*batch.JobStatus)(nil), (*v1.JobStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_batch_JobStatus_To_v1_JobStatus(a.(*batch.JobStatus), b.(*v1.JobStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.JobTemplateSpec)(nil), (*batch.JobTemplateSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_JobTemplateSpec_To_batch_JobTemplateSpec(a.(*v1.JobTemplateSpec), b.(*batch.JobTemplateSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*batch.JobTemplateSpec)(nil), (*v1.JobTemplateSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_batch_JobTemplateSpec_To_v1_JobTemplateSpec(a.(*batch.JobTemplateSpec), b.(*v1.JobTemplateSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*batch.JobSpec)(nil), (*v1.JobSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_batch_JobSpec_To_v1_JobSpec(a.(*batch.JobSpec), b.(*v1.JobSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1.JobSpec)(nil), (*batch.JobSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_JobSpec_To_batch_JobSpec(a.(*v1.JobSpec), b.(*batch.JobSpec), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1_CronJob_To_batch_CronJob(in *v1.CronJob, out *batch.CronJob, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_CronJobSpec_To_batch_CronJobSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_CronJobStatus_To_batch_CronJobStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_CronJob_To_batch_CronJob is an autogenerated conversion function.
func Convert_v1_CronJob_To_batch_CronJob(in *v1.CronJob, out *batch.CronJob, s conversion.Scope) error {
	return autoConvert_v1_CronJob_To_batch_CronJob(in, out, s)
}

func autoConvert_batch_CronJob_To_v1_CronJob(in *batch.CronJob, out *v1.CronJob, s conversion.Scope) error {
	// Auto-generated external APIVersion/Kind
	// Disable with a `+k8s:conversion-gen:set-api-version-kind=false` comment
	// Customize with `+groupName`, `+version`, or `+kind` comments
	out.APIVersion = "batch/v1"
	out.Kind = "CronJob"

	out.ObjectMeta = in.ObjectMeta
	if err := Convert_batch_CronJobSpec_To_v1_CronJobSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_batch_CronJobStatus_To_v1_CronJobStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_batch_CronJob_To_v1_CronJob is an autogenerated conversion function.
func Convert_batch_CronJob_To_v1_CronJob(in *batch.CronJob, out *v1.CronJob, s conversion.Scope) error {
	return autoConvert_batch_CronJob_To_v1_CronJob(in, out, s)
}

func autoConvert_v1_CronJobList_To_batch_CronJobList(in *v1.CronJobList, out *batch.CronJobList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]batch.CronJob, len(*in))
		for i := range *in {
			if err := Convert_v1_CronJob_To_batch_CronJob(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1_CronJobList_To_batch_CronJobList is an autogenerated conversion function.
func Convert_v1_CronJobList_To_batch_CronJobList(in *v1.CronJobList, out *batch.CronJobList, s conversion.Scope) error {
	return autoConvert_v1_CronJobList_To_batch_CronJobList(in, out, s)
}

func autoConvert_batch_CronJobList_To_v1_CronJobList(in *batch.CronJobList, out *v1.CronJobList, s conversion.Scope) error {
	// Auto-generated external APIVersion/Kind
	// Disable with a `+k8s:conversion-gen:set-api-version-kind=false` comment
	// Customize with `+groupName`, `+version`, or `+kind` comments
	out.APIVersion = "batch/v1"
	out.Kind = "CronJobList"

	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1.CronJob, len(*in))
		for i := range *in {
			if err := Convert_batch_CronJob_To_v1_CronJob(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_batch_CronJobList_To_v1_CronJobList is an autogenerated conversion function.
func Convert_batch_CronJobList_To_v1_CronJobList(in *batch.CronJobList, out *v1.CronJobList, s conversion.Scope) error {
	return autoConvert_batch_CronJobList_To_v1_CronJobList(in, out, s)
}

func autoConvert_v1_CronJobSpec_To_batch_CronJobSpec(in *v1.CronJobSpec, out *batch.CronJobSpec, s conversion.Scope) error {
	out.Schedule = in.Schedule
	out.StartingDeadlineSeconds = (*int64)(unsafe.Pointer(in.StartingDeadlineSeconds))
	out.ConcurrencyPolicy = batch.ConcurrencyPolicy(in.ConcurrencyPolicy)
	out.Suspend = (*bool)(unsafe.Pointer(in.Suspend))
	if err := Convert_v1_JobTemplateSpec_To_batch_JobTemplateSpec(&in.JobTemplate, &out.JobTemplate, s); err != nil {
		return err
	}
	out.SuccessfulJobsHistoryLimit = (*int32)(unsafe.Pointer(in.SuccessfulJobsHistoryLimit))
	out.FailedJobsHistoryLimit = (*int32)(unsafe.Pointer(in.FailedJobsHistoryLimit))
	return nil
}

// Convert_v1_CronJobSpec_To_batch_CronJobSpec is an autogenerated conversion function.
func Convert_v1_CronJobSpec_To_batch_CronJobSpec(in *v1.CronJobSpec, out *batch.CronJobSpec, s conversion.Scope) error {
	return autoConvert_v1_CronJobSpec_To_batch_CronJobSpec(in, out, s)
}

func autoConvert_batch_CronJobSpec_To_v1_CronJobSpec(in *batch.CronJobSpec, out *v1.CronJobSpec, s conversion.Scope) error {
	out.Schedule = in.Schedule
	out.StartingDeadlineSeconds = (*int64)(unsafe.Pointer(in.StartingDeadlineSeconds))
	out.ConcurrencyPolicy = v1.ConcurrencyPolicy(in.ConcurrencyPolicy)
	out.Suspend = (*bool)(unsafe.Pointer(in.Suspend))
	if err := Convert_batch_JobTemplateSpec_To_v1_JobTemplateSpec(&in.JobTemplate, &out.JobTemplate, s); err != nil {
		return err
	}
	out.SuccessfulJobsHistoryLimit = (*int32)(unsafe.Pointer(in.SuccessfulJobsHistoryLimit))
	out.FailedJobsHistoryLimit = (*int32)(unsafe.Pointer(in.FailedJobsHistoryLimit))
	return nil
}

// Convert_batch_CronJobSpec_To_v1_CronJobSpec is an autogenerated conversion function.
func Convert_batch_CronJobSpec_To_v1_CronJobSpec(in *batch.CronJobSpec, out *v1.CronJobSpec, s conversion.Scope) error {
	return autoConvert_batch_CronJobSpec_To_v1_CronJobSpec(in, out, s)
}

func autoConvert_v1_CronJobStatus_To_batch_CronJobStatus(in *v1.CronJobStatus, out *batch.CronJobStatus, s conversion.Scope) error {
	out.Active = *(*[]core.ObjectReference)(unsafe.Pointer(&in.Active))
	out.LastScheduleTime = (*metav1.Time)(unsafe.Pointer(in.LastScheduleTime))
	out.LastSuccessfulTime = (*metav1.Time)(unsafe.Pointer(in.LastSuccessfulTime))
	return nil
}

// Convert_v1_CronJobStatus_To_batch_CronJobStatus is an autogenerated conversion function.
func Convert_v1_CronJobStatus_To_batch_CronJobStatus(in *v1.CronJobStatus, out *batch.CronJobStatus, s conversion.Scope) error {
	return autoConvert_v1_CronJobStatus_To_batch_CronJobStatus(in, out, s)
}

func autoConvert_batch_CronJobStatus_To_v1_CronJobStatus(in *batch.CronJobStatus, out *v1.CronJobStatus, s conversion.Scope) error {
	out.Active = *(*[]corev1.ObjectReference)(unsafe.Pointer(&in.Active))
	out.LastScheduleTime = (*metav1.Time)(unsafe.Pointer(in.LastScheduleTime))
	out.LastSuccessfulTime = (*metav1.Time)(unsafe.Pointer(in.LastSuccessfulTime))
	return nil
}

// Convert_batch_CronJobStatus_To_v1_CronJobStatus is an autogenerated conversion function.
func Convert_batch_CronJobStatus_To_v1_CronJobStatus(in *batch.CronJobStatus, out *v1.CronJobStatus, s conversion.Scope) error {
	return autoConvert_batch_CronJobStatus_To_v1_CronJobStatus(in, out, s)
}

func autoConvert_v1_Job_To_batch_Job(in *v1.Job, out *batch.Job, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_JobSpec_To_batch_JobSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_JobStatus_To_batch_JobStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_Job_To_batch_Job is an autogenerated conversion function.
func Convert_v1_Job_To_batch_Job(in *v1.Job, out *batch.Job, s conversion.Scope) error {
	return autoConvert_v1_Job_To_batch_Job(in, out, s)
}

func autoConvert_batch_Job_To_v1_Job(in *batch.Job, out *v1.Job, s conversion.Scope) error {
	// Auto-generated external APIVersion/Kind
	// Disable with a `+k8s:conversion-gen:set-api-version-kind=false` comment
	// Customize with `+groupName`, `+version`, or `+kind` comments
	out.APIVersion = "batch/v1"
	out.Kind = "Job"

	out.ObjectMeta = in.ObjectMeta
	if err := Convert_batch_JobSpec_To_v1_JobSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_batch_JobStatus_To_v1_JobStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_batch_Job_To_v1_Job is an autogenerated conversion function.
func Convert_batch_Job_To_v1_Job(in *batch.Job, out *v1.Job, s conversion.Scope) error {
	return autoConvert_batch_Job_To_v1_Job(in, out, s)
}

func autoConvert_v1_JobCondition_To_batch_JobCondition(in *v1.JobCondition, out *batch.JobCondition, s conversion.Scope) error {
	out.Type = batch.JobConditionType(in.Type)
	out.Status = core.ConditionStatus(in.Status)
	out.LastProbeTime = in.LastProbeTime
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_v1_JobCondition_To_batch_JobCondition is an autogenerated conversion function.
func Convert_v1_JobCondition_To_batch_JobCondition(in *v1.JobCondition, out *batch.JobCondition, s conversion.Scope) error {
	return autoConvert_v1_JobCondition_To_batch_JobCondition(in, out, s)
}

func autoConvert_batch_JobCondition_To_v1_JobCondition(in *batch.JobCondition, out *v1.JobCondition, s conversion.Scope) error {
	out.Type = v1.JobConditionType(in.Type)
	out.Status = corev1.ConditionStatus(in.Status)
	out.LastProbeTime = in.LastProbeTime
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_batch_JobCondition_To_v1_JobCondition is an autogenerated conversion function.
func Convert_batch_JobCondition_To_v1_JobCondition(in *batch.JobCondition, out *v1.JobCondition, s conversion.Scope) error {
	return autoConvert_batch_JobCondition_To_v1_JobCondition(in, out, s)
}

func autoConvert_v1_JobList_To_batch_JobList(in *v1.JobList, out *batch.JobList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]batch.Job, len(*in))
		for i := range *in {
			if err := Convert_v1_Job_To_batch_Job(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1_JobList_To_batch_JobList is an autogenerated conversion function.
func Convert_v1_JobList_To_batch_JobList(in *v1.JobList, out *batch.JobList, s conversion.Scope) error {
	return autoConvert_v1_JobList_To_batch_JobList(in, out, s)
}

func autoConvert_batch_JobList_To_v1_JobList(in *batch.JobList, out *v1.JobList, s conversion.Scope) error {
	// Auto-generated external APIVersion/Kind
	// Disable with a `+k8s:conversion-gen:set-api-version-kind=false` comment
	// Customize with `+groupName`, `+version`, or `+kind` comments
	out.APIVersion = "batch/v1"
	out.Kind = "JobList"

	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1.Job, len(*in))
		for i := range *in {
			if err := Convert_batch_Job_To_v1_Job(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_batch_JobList_To_v1_JobList is an autogenerated conversion function.
func Convert_batch_JobList_To_v1_JobList(in *batch.JobList, out *v1.JobList, s conversion.Scope) error {
	return autoConvert_batch_JobList_To_v1_JobList(in, out, s)
}

func autoConvert_v1_JobSpec_To_batch_JobSpec(in *v1.JobSpec, out *batch.JobSpec, s conversion.Scope) error {
	out.Parallelism = (*int32)(unsafe.Pointer(in.Parallelism))
	out.Completions = (*int32)(unsafe.Pointer(in.Completions))
	out.ActiveDeadlineSeconds = (*int64)(unsafe.Pointer(in.ActiveDeadlineSeconds))
	out.BackoffLimit = (*int32)(unsafe.Pointer(in.BackoffLimit))
	out.Selector = (*metav1.LabelSelector)(unsafe.Pointer(in.Selector))
	out.ManualSelector = (*bool)(unsafe.Pointer(in.ManualSelector))
	if err := apiscorev1.Convert_v1_PodTemplateSpec_To_core_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
		return err
	}
	out.TTLSecondsAfterFinished = (*int32)(unsafe.Pointer(in.TTLSecondsAfterFinished))
	out.CompletionMode = (*batch.CompletionMode)(unsafe.Pointer(in.CompletionMode))
	out.Suspend = (*bool)(unsafe.Pointer(in.Suspend))
	return nil
}

func autoConvert_batch_JobSpec_To_v1_JobSpec(in *batch.JobSpec, out *v1.JobSpec, s conversion.Scope) error {
	out.Parallelism = (*int32)(unsafe.Pointer(in.Parallelism))
	out.Completions = (*int32)(unsafe.Pointer(in.Completions))
	out.ActiveDeadlineSeconds = (*int64)(unsafe.Pointer(in.ActiveDeadlineSeconds))
	out.BackoffLimit = (*int32)(unsafe.Pointer(in.BackoffLimit))
	out.Selector = (*metav1.LabelSelector)(unsafe.Pointer(in.Selector))
	out.ManualSelector = (*bool)(unsafe.Pointer(in.ManualSelector))
	if err := apiscorev1.Convert_core_PodTemplateSpec_To_v1_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
		return err
	}
	out.TTLSecondsAfterFinished = (*int32)(unsafe.Pointer(in.TTLSecondsAfterFinished))
	out.CompletionMode = (*v1.CompletionMode)(unsafe.Pointer(in.CompletionMode))
	out.Suspend = (*bool)(unsafe.Pointer(in.Suspend))
	return nil
}

func autoConvert_v1_JobStatus_To_batch_JobStatus(in *v1.JobStatus, out *batch.JobStatus, s conversion.Scope) error {
	out.Conditions = *(*[]batch.JobCondition)(unsafe.Pointer(&in.Conditions))
	out.StartTime = (*metav1.Time)(unsafe.Pointer(in.StartTime))
	out.CompletionTime = (*metav1.Time)(unsafe.Pointer(in.CompletionTime))
	out.Active = in.Active
	out.Succeeded = in.Succeeded
	out.Failed = in.Failed
	out.CompletedIndexes = in.CompletedIndexes
	return nil
}

// Convert_v1_JobStatus_To_batch_JobStatus is an autogenerated conversion function.
func Convert_v1_JobStatus_To_batch_JobStatus(in *v1.JobStatus, out *batch.JobStatus, s conversion.Scope) error {
	return autoConvert_v1_JobStatus_To_batch_JobStatus(in, out, s)
}

func autoConvert_batch_JobStatus_To_v1_JobStatus(in *batch.JobStatus, out *v1.JobStatus, s conversion.Scope) error {
	out.Conditions = *(*[]v1.JobCondition)(unsafe.Pointer(&in.Conditions))
	out.StartTime = (*metav1.Time)(unsafe.Pointer(in.StartTime))
	out.CompletionTime = (*metav1.Time)(unsafe.Pointer(in.CompletionTime))
	out.Active = in.Active
	out.Succeeded = in.Succeeded
	out.Failed = in.Failed
	out.CompletedIndexes = in.CompletedIndexes
	return nil
}

// Convert_batch_JobStatus_To_v1_JobStatus is an autogenerated conversion function.
func Convert_batch_JobStatus_To_v1_JobStatus(in *batch.JobStatus, out *v1.JobStatus, s conversion.Scope) error {
	return autoConvert_batch_JobStatus_To_v1_JobStatus(in, out, s)
}

func autoConvert_v1_JobTemplateSpec_To_batch_JobTemplateSpec(in *v1.JobTemplateSpec, out *batch.JobTemplateSpec, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_JobSpec_To_batch_JobSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_JobTemplateSpec_To_batch_JobTemplateSpec is an autogenerated conversion function.
func Convert_v1_JobTemplateSpec_To_batch_JobTemplateSpec(in *v1.JobTemplateSpec, out *batch.JobTemplateSpec, s conversion.Scope) error {
	return autoConvert_v1_JobTemplateSpec_To_batch_JobTemplateSpec(in, out, s)
}

func autoConvert_batch_JobTemplateSpec_To_v1_JobTemplateSpec(in *batch.JobTemplateSpec, out *v1.JobTemplateSpec, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_batch_JobSpec_To_v1_JobSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_batch_JobTemplateSpec_To_v1_JobTemplateSpec is an autogenerated conversion function.
func Convert_batch_JobTemplateSpec_To_v1_JobTemplateSpec(in *batch.JobTemplateSpec, out *v1.JobTemplateSpec, s conversion.Scope) error {
	return autoConvert_batch_JobTemplateSpec_To_v1_JobTemplateSpec(in, out, s)
}
