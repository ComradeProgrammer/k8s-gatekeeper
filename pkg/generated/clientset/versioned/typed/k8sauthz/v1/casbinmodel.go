// Copyright 2022 The Casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/casbin/k8s-gatekeeper/pkg/apis/k8sauthz/v1"
	scheme "github.com/casbin/k8s-gatekeeper/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CasbinModelsGetter has a method to return a CasbinModelInterface.
// A group's client should implement this interface.
type CasbinModelsGetter interface {
	CasbinModels(namespace string) CasbinModelInterface
}

// CasbinModelInterface has methods to work with CasbinModel resources.
type CasbinModelInterface interface {
	Create(ctx context.Context, casbinModel *v1.CasbinModel, opts metav1.CreateOptions) (*v1.CasbinModel, error)
	Update(ctx context.Context, casbinModel *v1.CasbinModel, opts metav1.UpdateOptions) (*v1.CasbinModel, error)
	UpdateStatus(ctx context.Context, casbinModel *v1.CasbinModel, opts metav1.UpdateOptions) (*v1.CasbinModel, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.CasbinModel, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.CasbinModelList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.CasbinModel, err error)
	CasbinModelExpansion
}

// casbinModels implements CasbinModelInterface
type casbinModels struct {
	client rest.Interface
	ns     string
}

// newCasbinModels returns a CasbinModels
func newCasbinModels(c *AuthV1Client, namespace string) *casbinModels {
	return &casbinModels{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the casbinModel, and returns the corresponding casbinModel object, and an error if there is any.
func (c *casbinModels) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.CasbinModel, err error) {
	result = &v1.CasbinModel{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("casbinmodels").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CasbinModels that match those selectors.
func (c *casbinModels) List(ctx context.Context, opts metav1.ListOptions) (result *v1.CasbinModelList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.CasbinModelList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("casbinmodels").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested casbinModels.
func (c *casbinModels) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("casbinmodels").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a casbinModel and creates it.  Returns the server's representation of the casbinModel, and an error, if there is any.
func (c *casbinModels) Create(ctx context.Context, casbinModel *v1.CasbinModel, opts metav1.CreateOptions) (result *v1.CasbinModel, err error) {
	result = &v1.CasbinModel{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("casbinmodels").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(casbinModel).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a casbinModel and updates it. Returns the server's representation of the casbinModel, and an error, if there is any.
func (c *casbinModels) Update(ctx context.Context, casbinModel *v1.CasbinModel, opts metav1.UpdateOptions) (result *v1.CasbinModel, err error) {
	result = &v1.CasbinModel{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("casbinmodels").
		Name(casbinModel.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(casbinModel).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *casbinModels) UpdateStatus(ctx context.Context, casbinModel *v1.CasbinModel, opts metav1.UpdateOptions) (result *v1.CasbinModel, err error) {
	result = &v1.CasbinModel{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("casbinmodels").
		Name(casbinModel.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(casbinModel).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the casbinModel and deletes it. Returns an error if one occurs.
func (c *casbinModels) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("casbinmodels").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *casbinModels) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("casbinmodels").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched casbinModel.
func (c *casbinModels) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.CasbinModel, err error) {
	result = &v1.CasbinModel{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("casbinmodels").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
