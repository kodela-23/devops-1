/*
Copyright 2018 The Kubernetes Authors.

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

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CSIDriver captures information about a Container Storage Interface (CSI)
// volume driver deployed on the cluster.
//
// CSIDriver objects are non-namespaced.
type CSIDriver struct {
	metav1.TypeMeta `json:",inline"`

	// Standard object metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata"`

	// Specification describing the CSI volume driver and any custom
	// configuration for it.
	Spec CSIDriverSpec `json:"spec"`

	// Status of the CSI volume driver.
	Status CSIDriverStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CSIDriverList is a collection of CSIDriver objects.
type CSIDriverList struct {
	metav1.TypeMeta `json:",inline"`

	// Standard list metadata
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	metav1.ListMeta `json:"metadata,omitempty"`

	// Items is the list of CSIDrivers
	Items []CSIDriver `json:"items"`
}

// CSIDriverSpec is the specification of a CSIDriver.
type CSIDriverSpec struct {
	// Driver indicates the name of the CSI driver that this object refers to.
	// This MUST be the same name returned by the CSI GetPluginName() call for
	// that driver.
	Driver string `json:"driver"`

	// Indicates this CSI volume driver requires an attach operation (because it
	// implements the CSI ControllerPublishVolume() method), and that Kubernetes
	// should call attach and wait for any attach operation to complete before
	// proceeding to mounting.
	// If value is not specified, default is true -- meaning attach will be
	// called.
	// +optional
	AttachRequired *bool `json:"attachRequired"`

	// Indicates this CSI volume driver requires additional pod information
	// (like podName, podUID, etc.) during mount operations.
	// If this is set to true, Kubelet will pass pod information as
	// VolumeAttributes in the CSI NodePublishVolume() calls.
	// If value is not specified, default is false -- meaning pod information
	// will not be passed on mount.
	// +optional
	PodInfoRequiredOnMount *bool `json:"podInfoRequiredOnMount"`
}

// CSIDriverStatus is the status of a CSIDriver.
type CSIDriverStatus struct {
	// Indicates the volume driver has been successfully deployed on the cluster
	// and is ready to use.
	Ready bool `json:"ready"`
}
