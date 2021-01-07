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

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	scheme "k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
)

// CSIDriversGetter has a method to return a CSIDriverInterface.
// A group's client should implement this interface.
type CSIDriversGetter interface {
	CSIDrivers() CSIDriverInterface
}

// CSIDriverInterface has methods to work with CSIDriver resources.
type CSIDriverInterface interface {
	Create(ctx context.Context, cSIDriver *v1.CSIDriver, opts metav1.CreateOptions) (*v1.CSIDriver, error)
	Update(ctx context.Context, cSIDriver *v1.CSIDriver, opts metav1.UpdateOptions) (*v1.CSIDriver, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.CSIDriver, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.CSIDriverList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.CSIDriver, err error)
	CSIDriverExpansion
}

// cSIDrivers implements CSIDriverInterface
type cSIDrivers struct {
	client rest.Interface
}

// newCSIDrivers returns a CSIDrivers
func newCSIDrivers(c *StorageV1Client) *cSIDrivers {
	return &cSIDrivers{
		client: c.RESTClient(),
	}
}

// Get takes name of the cSIDriver, and returns the corresponding cSIDriver object, and an error if there is any.
func (c *cSIDrivers) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.CSIDriver, err error) {
	result = &v1.CSIDriver{}
	err = c.client.Get().
		Resource("csidrivers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CSIDrivers that match those selectors.
func (c *cSIDrivers) List(ctx context.Context, opts metav1.ListOptions) (result *v1.CSIDriverList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.CSIDriverList{}
	err = c.client.Get().
		Prefix("/apis", "storage.k8s.io", "v1").
		GroupVersion(v1.SchemeGroupVersion).
		Resource("csidrivers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cSIDrivers.
func (c *cSIDrivers) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("csidrivers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a cSIDriver and creates it.  Returns the server's representation of the cSIDriver, and an error, if there is any.
func (c *cSIDrivers) Create(ctx context.Context, cSIDriver *v1.CSIDriver, opts metav1.CreateOptions) (result *v1.CSIDriver, err error) {
	result = &v1.CSIDriver{}
	err = c.client.Post().
		Resource("csidrivers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cSIDriver).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a cSIDriver and updates it. Returns the server's representation of the cSIDriver, and an error, if there is any.
func (c *cSIDrivers) Update(ctx context.Context, cSIDriver *v1.CSIDriver, opts metav1.UpdateOptions) (result *v1.CSIDriver, err error) {
	result = &v1.CSIDriver{}
	err = c.client.Put().
		Resource("csidrivers").
		Name(cSIDriver.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cSIDriver).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the cSIDriver and deletes it. Returns an error if one occurs.
func (c *cSIDrivers) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("csidrivers").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cSIDrivers) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("csidrivers").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched cSIDriver.
func (c *cSIDrivers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.CSIDriver, err error) {
	result = &v1.CSIDriver{}
	err = c.client.Patch(pt).
		Resource("csidrivers").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
