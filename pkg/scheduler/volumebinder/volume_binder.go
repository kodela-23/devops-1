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

package volumebinder

import (
	"time"

	coreinformers "k8s.io/client-go/informers/core/v1"
	storageinformers "k8s.io/client-go/informers/storage/v1"
	clientset "k8s.io/client-go/kubernetes"
	volumescheduling "k8s.io/kubernetes/pkg/controller/volume/scheduling"
)

// VolumeBinder sets up the volume binding library
type VolumeBinder struct {
	Binder volumescheduling.SchedulerVolumeBinder
}

// NewVolumeBinder sets up the volume binding library and binding queue
func NewVolumeBinder(
	client clientset.Interface,
	nodeInformer coreinformers.NodeInformer,
	pvcInformer coreinformers.PersistentVolumeClaimInformer,
	pvInformer coreinformers.PersistentVolumeInformer,
	storageClassInformer storageinformers.StorageClassInformer,
	bindTimeout time.Duration) *VolumeBinder {

	return &VolumeBinder{
		Binder: volumescheduling.NewVolumeBinder(client, nodeInformer, pvcInformer, pvInformer, storageClassInformer, bindTimeout),
	}
}

// NewFakeVolumeBinder sets up a fake volume binder and binding queue
func NewFakeVolumeBinder(config *volumescheduling.FakeVolumeBinderConfig) *VolumeBinder {
	return &VolumeBinder{
		Binder: volumescheduling.NewFakeVolumeBinder(config),
	}
}
