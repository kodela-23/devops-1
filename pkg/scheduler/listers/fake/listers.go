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

package fake

import (
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
	storagev1beta1 "k8s.io/api/storage/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	appslisters "k8s.io/client-go/listers/apps/v1"
	corelisters "k8s.io/client-go/listers/core/v1"
	schedulerlisters "k8s.io/kubernetes/pkg/scheduler/listers"
	schedulernodeinfo "k8s.io/kubernetes/pkg/scheduler/nodeinfo"
)

var _ schedulerlisters.PodLister = &PodLister{}

// PodLister implements PodLister on an []v1.Pods for test purposes.
type PodLister []*v1.Pod

// List returns []*v1.Pod matching a query.
func (f PodLister) List(s labels.Selector) (selected []*v1.Pod, err error) {
	for _, pod := range f {
		if s.Matches(labels.Set(pod.Labels)) {
			selected = append(selected, pod)
		}
	}
	return selected, nil
}

// FilteredList returns pods matching a pod filter and a label selector.
func (f PodLister) FilteredList(podFilter schedulerlisters.PodFilter, s labels.Selector) (selected []*v1.Pod, err error) {
	for _, pod := range f {
		if podFilter(pod) && s.Matches(labels.Set(pod.Labels)) {
			selected = append(selected, pod)
		}
	}
	return selected, nil
}

var _ corelisters.ServiceLister = &ServiceLister{}

// ServiceLister implements ServiceLister on []v1.Service for test purposes.
type ServiceLister []*v1.Service

// Services returns nil.
func (f ServiceLister) Services(namespace string) corelisters.ServiceNamespaceLister {
	return nil
}

// List returns v1.ServiceList, the list of all services.
func (f ServiceLister) List(labels.Selector) ([]*v1.Service, error) {
	return f, nil
}

// GetPodServices gets the services that have the selector that match the labels on the given pod.
func (f ServiceLister) GetPodServices(pod *v1.Pod) (services []*v1.Service, err error) {
	var selector labels.Selector

	for i := range f {
		service := f[i]
		// consider only services that are in the same namespace as the pod
		if service.Namespace != pod.Namespace {
			continue
		}
		selector = labels.Set(service.Spec.Selector).AsSelectorPreValidated()
		if selector.Matches(labels.Set(pod.Labels)) {
			services = append(services, service)
		}
	}
	return
}

var _ corelisters.ReplicationControllerLister = &ControllerLister{}

// ControllerLister implements ControllerLister on []v1.ReplicationController for test purposes.
type ControllerLister []*v1.ReplicationController

// List returns []v1.ReplicationController, the list of all ReplicationControllers.
func (f ControllerLister) List(labels.Selector) ([]*v1.ReplicationController, error) {
	return f, nil
}

// GetPodControllers gets the ReplicationControllers that have the selector that match the labels on the given pod
func (f ControllerLister) GetPodControllers(pod *v1.Pod) (controllers []*v1.ReplicationController, err error) {
	var selector labels.Selector

	for i := range f {
		controller := f[i]
		if controller.Namespace != pod.Namespace {
			continue
		}
		selector = labels.Set(controller.Spec.Selector).AsSelectorPreValidated()
		if selector.Matches(labels.Set(pod.Labels)) {
			controllers = append(controllers, controller)
		}
	}
	if len(controllers) == 0 {
		err = fmt.Errorf("Could not find Replication Controller for pod %s in namespace %s with labels: %v", pod.Name, pod.Namespace, pod.Labels)
	}

	return
}

// ReplicationControllers returns nil
func (f ControllerLister) ReplicationControllers(namespace string) corelisters.ReplicationControllerNamespaceLister {
	return nil
}

var _ appslisters.ReplicaSetLister = &ReplicaSetLister{}

// ReplicaSetLister implements ControllerLister on []extensions.ReplicaSet for test purposes.
type ReplicaSetLister []*appsv1.ReplicaSet

// List returns replica sets.
func (f ReplicaSetLister) List(labels.Selector) ([]*appsv1.ReplicaSet, error) {
	return f, nil
}

// GetPodReplicaSets gets the ReplicaSets that have the selector that match the labels on the given pod
func (f ReplicaSetLister) GetPodReplicaSets(pod *v1.Pod) (rss []*appsv1.ReplicaSet, err error) {
	var selector labels.Selector

	for _, rs := range f {
		if rs.Namespace != pod.Namespace {
			continue
		}
		selector, err = metav1.LabelSelectorAsSelector(rs.Spec.Selector)
		if err != nil {
			return
		}

		if selector.Matches(labels.Set(pod.Labels)) {
			rss = append(rss, rs)
		}
	}
	if len(rss) == 0 {
		err = fmt.Errorf("Could not find ReplicaSet for pod %s in namespace %s with labels: %v", pod.Name, pod.Namespace, pod.Labels)
	}

	return
}

// ReplicaSets returns nil
func (f ReplicaSetLister) ReplicaSets(namespace string) appslisters.ReplicaSetNamespaceLister {
	return nil
}

var _ appslisters.StatefulSetLister = &StatefulSetLister{}

// StatefulSetLister implements ControllerLister on []appsv1.StatefulSet for testing purposes.
type StatefulSetLister []*appsv1.StatefulSet

// List returns stateful sets.
func (f StatefulSetLister) List(labels.Selector) ([]*appsv1.StatefulSet, error) {
	return f, nil
}

// GetPodStatefulSets gets the StatefulSets that have the selector that match the labels on the given pod.
func (f StatefulSetLister) GetPodStatefulSets(pod *v1.Pod) (sss []*appsv1.StatefulSet, err error) {
	var selector labels.Selector

	for _, ss := range f {
		if ss.Namespace != pod.Namespace {
			continue
		}
		selector, err = metav1.LabelSelectorAsSelector(ss.Spec.Selector)
		if err != nil {
			return
		}
		if selector.Matches(labels.Set(pod.Labels)) {
			sss = append(sss, ss)
		}
	}
	if len(sss) == 0 {
		err = fmt.Errorf("Could not find StatefulSet for pod %s in namespace %s with labels: %v", pod.Name, pod.Namespace, pod.Labels)
	}
	return
}

// StatefulSets returns nil
func (f StatefulSetLister) StatefulSets(namespace string) appslisters.StatefulSetNamespaceLister {
	return nil
}

// PersistentVolumeClaimLister implements PersistentVolumeClaimLister on []*v1.PersistentVolumeClaim for test purposes.
type PersistentVolumeClaimLister []*v1.PersistentVolumeClaim

var _ corelisters.PersistentVolumeClaimLister = PersistentVolumeClaimLister{}

// List returns not implemented error.
func (f PersistentVolumeClaimLister) List(selector labels.Selector) (ret []*v1.PersistentVolumeClaim, err error) {
	return nil, fmt.Errorf("not implemented")
}

// PersistentVolumeClaims returns a fake PersistentVolumeClaimLister object.
func (f PersistentVolumeClaimLister) PersistentVolumeClaims(namespace string) corelisters.PersistentVolumeClaimNamespaceLister {
	return &persistentVolumeClaimNamespaceLister{
		pvcs:      f,
		namespace: namespace,
	}
}

// persistentVolumeClaimNamespaceLister is implementation of PersistentVolumeClaimNamespaceLister returned by List() above.
type persistentVolumeClaimNamespaceLister struct {
	pvcs      []*v1.PersistentVolumeClaim
	namespace string
}

func (f *persistentVolumeClaimNamespaceLister) Get(name string) (*v1.PersistentVolumeClaim, error) {
	for _, pvc := range f.pvcs {
		if pvc.Name == name && pvc.Namespace == f.namespace {
			return pvc, nil
		}
	}
	return nil, fmt.Errorf("persistentvolumeclaim %q not found", name)
}

func (f persistentVolumeClaimNamespaceLister) List(selector labels.Selector) (ret []*v1.PersistentVolumeClaim, err error) {
	return nil, fmt.Errorf("not implemented")
}

// PersistentVolumeClaimInfo declares a []v1.PersistentVolumeClaim type for testing.
type PersistentVolumeClaimInfo []v1.PersistentVolumeClaim

// GetPersistentVolumeClaimInfo gets PVC matching the namespace and PVC ID.
func (pvcs PersistentVolumeClaimInfo) GetPersistentVolumeClaimInfo(namespace string, pvcID string) (*v1.PersistentVolumeClaim, error) {
	for _, pvc := range pvcs {
		if pvc.Name == pvcID && pvc.Namespace == namespace {
			return &pvc, nil
		}
	}
	return nil, fmt.Errorf("Unable to find persistent volume claim: %s/%s", namespace, pvcID)
}

// NodeInfo declares a schedulernodeinfo.NodeInfo type for testing.
type NodeInfoLister []*schedulernodeinfo.NodeInfo

// Get returns a fake node object in the fake nodes.
func (nodes NodeInfoLister) Get(nodeName string) (*schedulernodeinfo.NodeInfo, error) {
	for _, node := range nodes {
		if node != nil && node.Node().Name == nodeName {
			return node, nil
		}
	}
	return nil, fmt.Errorf("Unable to find node: %s", nodeName)
}

func (nodes NodeInfoLister) List() ([]*schedulernodeinfo.NodeInfo, error) {
	return nodes, nil
}

func NewNodeInfoLister(nodes []*v1.Node) schedulerlisters.NodeInfoLister {
	nodeInfoList := make([]*schedulernodeinfo.NodeInfo, len(nodes))
	for _, node := range nodes {
		nodeInfo := schedulernodeinfo.NewNodeInfo()
		nodeInfo.SetNode(node)
		nodeInfoList = append(nodeInfoList, nodeInfo)
	}

	return NodeInfoLister(nodeInfoList)
}

// CSINodeInfo declares a storagev1beta1.CSINode type for testing.
type CSINodeInfo storagev1beta1.CSINode

// GetCSINodeInfo returns a fake CSINode object.
func (n CSINodeInfo) GetCSINodeInfo(name string) (*storagev1beta1.CSINode, error) {
	csiNode := storagev1beta1.CSINode(n)
	return &csiNode, nil
}

// PersistentVolumeInfo declares a []v1.PersistentVolume type for testing.
type PersistentVolumeInfo []v1.PersistentVolume

// GetPersistentVolumeInfo returns a fake PV object in the fake PVs by PV ID.
func (pvs PersistentVolumeInfo) GetPersistentVolumeInfo(pvID string) (*v1.PersistentVolume, error) {
	for _, pv := range pvs {
		if pv.Name == pvID {
			return &pv, nil
		}
	}
	return nil, fmt.Errorf("Unable to find persistent volume: %s", pvID)
}

// StorageClassInfo declares a []storagev1.StorageClass type for testing.
type StorageClassInfo []storagev1.StorageClass

// GetStorageClassInfo returns a fake storage class object in the fake storage classes by name.
func (classes StorageClassInfo) GetStorageClassInfo(name string) (*storagev1.StorageClass, error) {
	for _, sc := range classes {
		if sc.Name == name {
			return &sc, nil
		}
	}
	return nil, fmt.Errorf("Unable to find storage class: %s", name)
}
