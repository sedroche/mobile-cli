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

package internalversion

import (
	mobile "github.com/aerogear/mobile-cli/pkg/apis/mobile"
	scheme "github.com/aerogear/mobile-cli/pkg/client/mobile/clientset/internalversion/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MobileClientsGetter has a method to return a MobileClientInterface.
// A group's client should implement this interface.
type MobileClientsGetter interface {
	MobileClients(namespace string) MobileClientInterface
}

// MobileClientInterface has methods to work with MobileClient resources.
type MobileClientInterface interface {
	Create(*mobile.MobileClient) (*mobile.MobileClient, error)
	Update(*mobile.MobileClient) (*mobile.MobileClient, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*mobile.MobileClient, error)
	List(opts v1.ListOptions) (*mobile.MobileClientList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *mobile.MobileClient, err error)
	MobileClientExpansion
}

// mobileClients implements MobileClientInterface
type mobileClients struct {
	client rest.Interface
	ns     string
}

// newMobileClients returns a MobileClients
func newMobileClients(c *MobileClient, namespace string) *mobileClients {
	return &mobileClients{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the mobileClient, and returns the corresponding mobileClient object, and an error if there is any.
func (c *mobileClients) Get(name string, options v1.GetOptions) (result *mobile.MobileClient, err error) {
	result = &mobile.MobileClient{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mobileclients").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MobileClients that match those selectors.
func (c *mobileClients) List(opts v1.ListOptions) (result *mobile.MobileClientList, err error) {
	result = &mobile.MobileClientList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mobileclients").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested mobileClients.
func (c *mobileClients) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("mobileclients").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a mobileClient and creates it.  Returns the server's representation of the mobileClient, and an error, if there is any.
func (c *mobileClients) Create(mobileClient *mobile.MobileClient) (result *mobile.MobileClient, err error) {
	result = &mobile.MobileClient{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("mobileclients").
		Body(mobileClient).
		Do().
		Into(result)
	return
}

// Update takes the representation of a mobileClient and updates it. Returns the server's representation of the mobileClient, and an error, if there is any.
func (c *mobileClients) Update(mobileClient *mobile.MobileClient) (result *mobile.MobileClient, err error) {
	result = &mobile.MobileClient{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mobileclients").
		Name(mobileClient.Name).
		Body(mobileClient).
		Do().
		Into(result)
	return
}

// Delete takes name of the mobileClient and deletes it. Returns an error if one occurs.
func (c *mobileClients) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mobileclients").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *mobileClients) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mobileclients").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched mobileClient.
func (c *mobileClients) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *mobile.MobileClient, err error) {
	result = &mobile.MobileClient{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("mobileclients").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
