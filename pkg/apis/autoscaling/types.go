/*
Copyright 2016 The Kubernetes Authors.

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

package autoscaling

import (
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/resource"
	"k8s.io/kubernetes/pkg/api/unversioned"
)

// Scale represents a scaling request for a resource.
type Scale struct {
	unversioned.TypeMeta `json:",inline"`
	// Standard object metadata; More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata.
	// +optional
	api.ObjectMeta `json:"metadata,omitempty"`

	// defines the behavior of the scale. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status.
	// +optional
	Spec ScaleSpec `json:"spec,omitempty"`

	// current status of the scale. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status. Read-only.
	// +optional
	Status ScaleStatus `json:"status,omitempty"`
}

// ScaleSpec describes the attributes of a scale subresource.
type ScaleSpec struct {
	// desired number of instances for the scaled object.
	// +optional
	Replicas int32 `json:"replicas,omitempty"`
}

// ScaleStatus represents the current status of a scale subresource.
type ScaleStatus struct {
	// actual number of observed instances of the scaled object.
	Replicas int32 `json:"replicas"`

	// label query over pods that should match the replicas count. This is same
	// as the label selector but in the string format to avoid introspection
	// by clients. The string will be in the same format as the query-param syntax.
	// More info: http://kubernetes.io/docs/user-guide/labels#label-selectors
	// +optional
	Selector string `json:"selector,omitempty"`
}

// CrossVersionObjectReference contains enough information to let you identify the referred resource.
type CrossVersionObjectReference struct {
	// Kind of the referent; More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#types-kinds"
	Kind string `json:"kind" protobuf:"bytes,1,opt,name=kind"`
	// Name of the referent; More info: http://kubernetes.io/docs/user-guide/identifiers#names
	Name string `json:"name" protobuf:"bytes,2,opt,name=name"`
	// API version of the referent
	// +optional
	APIVersion string `json:"apiVersion,omitempty" protobuf:"bytes,3,opt,name=apiVersion"`
}

// specification of a horizontal pod autoscaler
type HorizontalPodAutoscalerSpec struct {
	// the target scalable object to autoscale
	ScaleTargetRef CrossVersionObjectReference `json:"scaleTargetRef"`

	// the minimum number of replicas to which the autoscaler may scale
	// +optional
	MinReplicas *int32 `json:"minReplicas,omitempty"`
	// the maximum number of replicas to which the autoscaler may scale
	MaxReplicas int32 `json:"maxReplicas"`

	// the metrics to use to calculate the desired replica count (the
	// maximum replica count across all metrics will be used).  It is
	// expected that any metrics used will decrease as the replica count
	// increases, and will eventually increase if we decrease the replica
	// count.
	Metrics []MetricSpec `json:"metrics,omitempty"`
}

// a type of metric source
type MetricSourceType string

var (
	// a metric describing a kubernetes object
	ObjectSourceType MetricSourceType = "object"
	// a metric describing pods in the scale target
	PodsSourceType MetricSourceType = "pods"
	// a resource metric known to Kubernetes
	ResourceSourceType MetricSourceType = "resource"
)

// a specification for how to scale based on a single metric
// (only `type` and one other matching field should be set at once)
type MetricSpec struct {
	// the type of metric source (should match one of the fields below)
	Type MetricSourceType `json:"type"`

	// metric describing a single Kubernetes object
	Object *ObjectMetricSource `json:"object,omitempty"`
	// metric describing pods in the scale target
	Pods *PodsMetricSource `json:"pods,omitempty"`
	// resource metric describing pods in the scale target
	// (guaranteed to be available and have the same names across clusters)
	Resource *ResourceMetricSource `json:"resource,omitempty"`
}

// a metric describing a Kubernetes object
type ObjectMetricSource struct {
	// the described Kubernetes object
	Target CrossVersionObjectReference `json:"target"`

	// the name of the metric in question
	MetricName string `json:"metricName"`
	// the target value of the metric (as a quantity)
	TargetValue resource.Quantity `json:"targetValue"`
}

// metric describing pods in the scale target
type PodsMetricSource struct {
	// the name of the metric in question
	MetricName string `json:"metricName"`
	// the target value of the metric (as a quantity)
	TargetValue resource.Quantity `json:"targetValue"`
}

// resource metric describing pods in the scale target
// (guaranteed to be available and have the same names across clusters)
type ResourceMetricSource struct {
	// the name of the resource in question
	Name api.ResourceName `json:"name"`
	// the target value of the resource metric as a percentage of the
	// request on the pods
	// +optional
	TargetPercentageOfRequest *int32 `json:"targetPercentageOfRequest,omitempty"`
	// the target value of the resource metric as a raw value
	// +optional
	TargetRawValue *resource.Quantity `json:"targetRawValue,omitempty"`
}

// the status of a horizontal pod autoscaler
type HorizontalPodAutoscalerStatus struct {
	// most recent generation observed by this autoscaler.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
	// last time the autoscaler scaled the number of pods;
	// used by the autoscaler to control how often the number of pods is changed.
	LastScaleTime *unversioned.Time `json:"lastScaleTime,omitempty"`

	// the last observed number of replicas from the target object.
	CurrentReplicas int32 `json:"currentReplicas"`
	// the desired number of replicas as last computed by the autoscaler
	DesiredReplicas int32 `json:"desiredReplicas"`

	// the last read state of the metrics used by this autoscaler
	CurrentMetrics []MetricStatus `json:"currentMetrics"`
}

// the status of a single metric
type MetricStatus struct {
	// the type of metric source
	Type MetricSourceType `json:"type"`

	// metric describing a single Kubernetes object
	Object *ObjectMetricStatus `json:"object,omitemtpy"`
	// metric describing pods in the scale target
	Pods *PodsMetricStatus `json:"pods,omitemtpy"`
	// resource metric describing pods in the scale target
	Resource *ResourceMetricStatus `json:"resource,omitempty"`
}

// a metric describing a Kubernetes object
type ObjectMetricStatus struct {
	// the described Kubernetes object
	Target CrossVersionObjectReference `json:"target"`

	// the name of the metric in question
	MetricName string `json:"metricName"`
	// the current value of the metric (as a quantity)
	CurrentValue resource.Quantity `json:"targetValue"`
}

// metric describing pods in the scale target
type PodsMetricStatus struct {
	// the name of the metric in question
	MetricName string `json:"metricName"`
	// the current value of the metric (as a quantity)
	CurrentValue resource.Quantity `json:"targetValue"`
}

// resource metric describing pods in the scale target
type ResourceMetricStatus struct {
	// the name of the resource in question
	Name api.ResourceName `json:"name"`
	// the target value of the resource metric as a percentage of the
	// request on the pods (only populated if request is available)
	// +optional
	CurrentPercentageOfRequest *int32 `json:"targetPercentageOfRequest,omitempty"`
	// the target value of the resource metric as a raw value
	CurrentRawValue *resource.Quantity `json:"targetRawValue"`
}

// +genclient=true

// configuration of a horizontal pod autoscaler.
type HorizontalPodAutoscaler struct {
	unversioned.TypeMeta `json:",inline"`
	// +optional
	api.ObjectMeta `json:"metadata,omitempty"`

	// behaviour of autoscaler. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status.
	// +optional
	Spec HorizontalPodAutoscalerSpec `json:"spec,omitempty"`

	// current information about the autoscaler.
	// +optional
	Status HorizontalPodAutoscalerStatus `json:"status,omitempty"`
}

// list of horizontal pod autoscaler objects.
type HorizontalPodAutoscalerList struct {
	unversioned.TypeMeta `json:",inline"`
	// +optional
	unversioned.ListMeta `json:"metadata,omitempty"`

	// list of horizontal pod autoscaler objects.
	Items []HorizontalPodAutoscaler `json:"items"`
}
