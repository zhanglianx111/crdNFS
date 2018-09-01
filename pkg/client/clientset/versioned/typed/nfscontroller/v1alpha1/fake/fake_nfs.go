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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/zhanglianx111/nfscontroller/pkg/apis/nfscontroller/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNfses implements NfsInterface
type FakeNfses struct {
	Fake *FakeNfscontrollerV1alpha1
	ns   string
}

var nfsesResource = schema.GroupVersionResource{Group: "nfscontroller.k8s.io", Version: "v1alpha1", Resource: "nfses"}

var nfsesKind = schema.GroupVersionKind{Group: "nfscontroller.k8s.io", Version: "v1alpha1", Kind: "Nfs"}

// Get takes name of the nfs, and returns the corresponding nfs object, and an error if there is any.
func (c *FakeNfses) Get(name string, options v1.GetOptions) (result *v1alpha1.Nfs, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(nfsesResource, c.ns, name), &v1alpha1.Nfs{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Nfs), err
}

// List takes label and field selectors, and returns the list of Nfses that match those selectors.
func (c *FakeNfses) List(opts v1.ListOptions) (result *v1alpha1.NfsList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(nfsesResource, nfsesKind, c.ns, opts), &v1alpha1.NfsList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NfsList{}
	for _, item := range obj.(*v1alpha1.NfsList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested nfses.
func (c *FakeNfses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(nfsesResource, c.ns, opts))

}

// Create takes the representation of a nfs and creates it.  Returns the server's representation of the nfs, and an error, if there is any.
func (c *FakeNfses) Create(nfs *v1alpha1.Nfs) (result *v1alpha1.Nfs, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(nfsesResource, c.ns, nfs), &v1alpha1.Nfs{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Nfs), err
}

// Update takes the representation of a nfs and updates it. Returns the server's representation of the nfs, and an error, if there is any.
func (c *FakeNfses) Update(nfs *v1alpha1.Nfs) (result *v1alpha1.Nfs, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(nfsesResource, c.ns, nfs), &v1alpha1.Nfs{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Nfs), err
}

// Delete takes name of the nfs and deletes it. Returns an error if one occurs.
func (c *FakeNfses) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(nfsesResource, c.ns, name), &v1alpha1.Nfs{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNfses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(nfsesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.NfsList{})
	return err
}

// Patch applies the patch and returns the patched nfs.
func (c *FakeNfses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Nfs, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(nfsesResource, c.ns, name, data, subresources...), &v1alpha1.Nfs{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Nfs), err
}