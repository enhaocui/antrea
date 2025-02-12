/*
Copyright 2021 Antrea Authors.

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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "antrea.io/antrea/multicluster/apis/multicluster/v1alpha1"
	scheme "antrea.io/antrea/multicluster/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ResourceExportsGetter has a method to return a ResourceExportInterface.
// A group's client should implement this interface.
type ResourceExportsGetter interface {
	ResourceExports(namespace string) ResourceExportInterface
}

// ResourceExportInterface has methods to work with ResourceExport resources.
type ResourceExportInterface interface {
	Create(ctx context.Context, resourceExport *v1alpha1.ResourceExport, opts v1.CreateOptions) (*v1alpha1.ResourceExport, error)
	Update(ctx context.Context, resourceExport *v1alpha1.ResourceExport, opts v1.UpdateOptions) (*v1alpha1.ResourceExport, error)
	UpdateStatus(ctx context.Context, resourceExport *v1alpha1.ResourceExport, opts v1.UpdateOptions) (*v1alpha1.ResourceExport, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ResourceExport, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ResourceExportList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ResourceExport, err error)
	ResourceExportExpansion
}

// resourceExports implements ResourceExportInterface
type resourceExports struct {
	client rest.Interface
	ns     string
}

// newResourceExports returns a ResourceExports
func newResourceExports(c *MulticlusterV1alpha1Client, namespace string) *resourceExports {
	return &resourceExports{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the resourceExport, and returns the corresponding resourceExport object, and an error if there is any.
func (c *resourceExports) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ResourceExport, err error) {
	result = &v1alpha1.ResourceExport{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("resourceexports").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ResourceExports that match those selectors.
func (c *resourceExports) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ResourceExportList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ResourceExportList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("resourceexports").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested resourceExports.
func (c *resourceExports) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("resourceexports").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a resourceExport and creates it.  Returns the server's representation of the resourceExport, and an error, if there is any.
func (c *resourceExports) Create(ctx context.Context, resourceExport *v1alpha1.ResourceExport, opts v1.CreateOptions) (result *v1alpha1.ResourceExport, err error) {
	result = &v1alpha1.ResourceExport{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("resourceexports").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(resourceExport).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a resourceExport and updates it. Returns the server's representation of the resourceExport, and an error, if there is any.
func (c *resourceExports) Update(ctx context.Context, resourceExport *v1alpha1.ResourceExport, opts v1.UpdateOptions) (result *v1alpha1.ResourceExport, err error) {
	result = &v1alpha1.ResourceExport{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("resourceexports").
		Name(resourceExport.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(resourceExport).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *resourceExports) UpdateStatus(ctx context.Context, resourceExport *v1alpha1.ResourceExport, opts v1.UpdateOptions) (result *v1alpha1.ResourceExport, err error) {
	result = &v1alpha1.ResourceExport{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("resourceexports").
		Name(resourceExport.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(resourceExport).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the resourceExport and deletes it. Returns an error if one occurs.
func (c *resourceExports) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("resourceexports").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *resourceExports) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("resourceexports").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched resourceExport.
func (c *resourceExports) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ResourceExport, err error) {
	result = &v1alpha1.ResourceExport{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("resourceexports").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
