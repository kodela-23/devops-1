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

// this file contains factories with no other dependencies

package util

import (
	"fmt"
	"reflect"
	"sort"
	"sync"
	"time"

	"k8s.io/api/core/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/dynamic"
	restclient "k8s.io/client-go/rest"
	"k8s.io/kubernetes/pkg/apis/batch"
	api "k8s.io/kubernetes/pkg/apis/core"
	apiv1 "k8s.io/kubernetes/pkg/apis/core/v1"
	"k8s.io/kubernetes/pkg/apis/extensions"
	"k8s.io/kubernetes/pkg/controller"
	"k8s.io/kubernetes/pkg/kubectl"
	"k8s.io/kubernetes/pkg/kubectl/cmd/util/openapi"
	openapivalidation "k8s.io/kubernetes/pkg/kubectl/cmd/util/openapi/validation"
	"k8s.io/kubernetes/pkg/kubectl/genericclioptions/resource"
	"k8s.io/kubernetes/pkg/kubectl/polymorphichelpers"
	"k8s.io/kubernetes/pkg/kubectl/validation"
	"k8s.io/kubernetes/pkg/printers"
	printersinternal "k8s.io/kubernetes/pkg/printers/internalversion"
)

type ring1Factory struct {
	clientAccessFactory ClientAccessFactory

	// openAPIGetter loads and caches openapi specs
	openAPIGetter openAPIGetter
}

type openAPIGetter struct {
	once   sync.Once
	getter openapi.Getter
}

func NewObjectMappingFactory(clientAccessFactory ClientAccessFactory) ObjectMappingFactory {
	f := &ring1Factory{
		clientAccessFactory: clientAccessFactory,
	}
	return f
}

func (f *ring1Factory) ClientForMapping(mapping *meta.RESTMapping) (resource.RESTClient, error) {
	cfg, err := f.clientAccessFactory.ToRESTConfig()
	if err != nil {
		return nil, err
	}
	if err := setKubernetesDefaults(cfg); err != nil {
		return nil, err
	}
	gvk := mapping.GroupVersionKind
	switch gvk.Group {
	case api.GroupName:
		cfg.APIPath = "/api"
	default:
		cfg.APIPath = "/apis"
	}
	gv := gvk.GroupVersion()
	cfg.GroupVersion = &gv
	return restclient.RESTClientFor(cfg)
}

func (f *ring1Factory) UnstructuredClientForMapping(mapping *meta.RESTMapping) (resource.RESTClient, error) {
	cfg, err := f.clientAccessFactory.ToRESTConfig()
	if err != nil {
		return nil, err
	}
	if err := restclient.SetKubernetesDefaults(cfg); err != nil {
		return nil, err
	}
	cfg.APIPath = "/apis"
	if mapping.GroupVersionKind.Group == api.GroupName {
		cfg.APIPath = "/api"
	}
	gv := mapping.GroupVersionKind.GroupVersion()
	cfg.ContentConfig = resource.UnstructuredPlusDefaultContentConfig()
	cfg.GroupVersion = &gv
	return restclient.RESTClientFor(cfg)
}

func (f *ring1Factory) Describer(mapping *meta.RESTMapping) (printers.Describer, error) {
	clientConfig, err := f.clientAccessFactory.ToRESTConfig()
	if err != nil {
		return nil, err
	}
	// try to get a describer
	if describer, ok := printersinternal.DescriberFor(mapping.GroupVersionKind.GroupKind(), clientConfig); ok {
		return describer, nil
	}
	// if this is a kind we don't have a describer for yet, go generic if possible
	if genericDescriber, genericErr := genericDescriber(f.clientAccessFactory, mapping); genericErr == nil {
		return genericDescriber, nil
	}
	// otherwise return an unregistered error
	return nil, fmt.Errorf("no description has been implemented for %s", mapping.GroupVersionKind.String())
}

// helper function to make a generic describer, or return an error
func genericDescriber(clientAccessFactory ClientAccessFactory, mapping *meta.RESTMapping) (printers.Describer, error) {
	clientConfig, err := clientAccessFactory.ToRESTConfig()
	if err != nil {
		return nil, err
	}

	// used to fetch the resource
	dynamicClient, err := dynamic.NewForConfig(clientConfig)
	if err != nil {
		return nil, err
	}

	// used to get events for the resource
	clientSet, err := clientAccessFactory.ClientSet()
	if err != nil {
		return nil, err
	}
	eventsClient := clientSet.Core()

	return printersinternal.GenericDescriberFor(mapping, dynamicClient, eventsClient), nil
}

func (f *ring1Factory) HistoryViewer(mapping *meta.RESTMapping) (kubectl.HistoryViewer, error) {
	external, err := f.clientAccessFactory.KubernetesClientSet()
	if err != nil {
		return nil, err
	}
	return kubectl.HistoryViewerFor(mapping.GroupVersionKind.GroupKind(), external)
}

func (f *ring1Factory) Rollbacker(mapping *meta.RESTMapping) (kubectl.Rollbacker, error) {
	external, err := f.clientAccessFactory.KubernetesClientSet()
	if err != nil {
		return nil, err
	}
	return kubectl.RollbackerFor(mapping.GroupVersionKind.GroupKind(), external)
}

func (f *ring1Factory) StatusViewer(mapping *meta.RESTMapping) (kubectl.StatusViewer, error) {
	clientset, err := f.clientAccessFactory.KubernetesClientSet()
	if err != nil {
		return nil, err
	}
	return kubectl.StatusViewerFor(mapping.GroupVersionKind.GroupKind(), clientset)
}

func (f *ring1Factory) ApproximatePodTemplateForObject(object runtime.Object) (*api.PodTemplateSpec, error) {
	switch t := object.(type) {
	case *api.Pod:
		return &api.PodTemplateSpec{
			ObjectMeta: t.ObjectMeta,
			Spec:       t.Spec,
		}, nil
	case *api.ReplicationController:
		return t.Spec.Template, nil
	case *extensions.ReplicaSet:
		return &t.Spec.Template, nil
	case *extensions.DaemonSet:
		return &t.Spec.Template, nil
	case *extensions.Deployment:
		return &t.Spec.Template, nil
	case *batch.Job:
		return &t.Spec.Template, nil
	}

	return nil, fmt.Errorf("unable to extract pod template from type %v", reflect.TypeOf(object))
}

func (f *ring1Factory) AttachablePodForObject(object runtime.Object, timeout time.Duration) (*api.Pod, error) {
	clientset, err := f.clientAccessFactory.ClientSet()
	if err != nil {
		return nil, err
	}

	switch t := object.(type) {
	case *api.Pod:
		return t, nil
	case *corev1.Pod:
		internalPod := &api.Pod{}
		err := apiv1.Convert_v1_Pod_To_core_Pod(t, internalPod, nil)
		return internalPod, err

	}

	namespace, selector, err := polymorphichelpers.SelectorsForObject(object)
	if err != nil {
		return nil, fmt.Errorf("cannot attach to %T: %v", object, err)
	}
	sortBy := func(pods []*v1.Pod) sort.Interface { return sort.Reverse(controller.ActivePods(pods)) }
	pod, _, err := polymorphichelpers.GetFirstPod(clientset.Core(), namespace, selector.String(), timeout, sortBy)
	return pod, err
}

func (f *ring1Factory) Validator(validate bool) (validation.Schema, error) {
	if !validate {
		return validation.NullSchema{}, nil
	}

	resources, err := f.OpenAPISchema()
	if err != nil {
		return nil, err
	}

	return validation.ConjunctiveSchema{
		openapivalidation.NewSchemaValidation(resources),
		validation.NoDoubleKeySchema{},
	}, nil
}

// OpenAPISchema returns metadata and structural information about Kubernetes object definitions.
func (f *ring1Factory) OpenAPISchema() (openapi.Resources, error) {
	discovery, err := f.clientAccessFactory.ToDiscoveryClient()
	if err != nil {
		return nil, err
	}

	// Lazily initialize the OpenAPIGetter once
	f.openAPIGetter.once.Do(func() {
		// Create the caching OpenAPIGetter
		f.openAPIGetter.getter = openapi.NewOpenAPIGetter(discovery)
	})

	// Delegate to the OpenAPIGetter
	return f.openAPIGetter.getter.Get()
}
