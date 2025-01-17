/*
The GNU AFFERO GENERAL PUBLIC LICENSE

Copyright (c) 2020-2025 Traefik Labs

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published
by the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/traefik/hub-crds/pkg/apis/hub/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAPIPortals implements APIPortalInterface
type FakeAPIPortals struct {
	Fake *FakeHubV1alpha1
	ns   string
}

var apiportalsResource = v1alpha1.SchemeGroupVersion.WithResource("apiportals")

var apiportalsKind = v1alpha1.SchemeGroupVersion.WithKind("APIPortal")

// Get takes name of the aPIPortal, and returns the corresponding aPIPortal object, and an error if there is any.
func (c *FakeAPIPortals) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.APIPortal, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(apiportalsResource, c.ns, name), &v1alpha1.APIPortal{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APIPortal), err
}

// List takes label and field selectors, and returns the list of APIPortals that match those selectors.
func (c *FakeAPIPortals) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.APIPortalList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(apiportalsResource, apiportalsKind, c.ns, opts), &v1alpha1.APIPortalList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.APIPortalList{ListMeta: obj.(*v1alpha1.APIPortalList).ListMeta}
	for _, item := range obj.(*v1alpha1.APIPortalList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested aPIPortals.
func (c *FakeAPIPortals) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(apiportalsResource, c.ns, opts))

}

// Create takes the representation of a aPIPortal and creates it.  Returns the server's representation of the aPIPortal, and an error, if there is any.
func (c *FakeAPIPortals) Create(ctx context.Context, aPIPortal *v1alpha1.APIPortal, opts v1.CreateOptions) (result *v1alpha1.APIPortal, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(apiportalsResource, c.ns, aPIPortal), &v1alpha1.APIPortal{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APIPortal), err
}

// Update takes the representation of a aPIPortal and updates it. Returns the server's representation of the aPIPortal, and an error, if there is any.
func (c *FakeAPIPortals) Update(ctx context.Context, aPIPortal *v1alpha1.APIPortal, opts v1.UpdateOptions) (result *v1alpha1.APIPortal, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(apiportalsResource, c.ns, aPIPortal), &v1alpha1.APIPortal{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APIPortal), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAPIPortals) UpdateStatus(ctx context.Context, aPIPortal *v1alpha1.APIPortal, opts v1.UpdateOptions) (*v1alpha1.APIPortal, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(apiportalsResource, "status", c.ns, aPIPortal), &v1alpha1.APIPortal{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APIPortal), err
}

// Delete takes name of the aPIPortal and deletes it. Returns an error if one occurs.
func (c *FakeAPIPortals) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(apiportalsResource, c.ns, name, opts), &v1alpha1.APIPortal{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAPIPortals) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(apiportalsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.APIPortalList{})
	return err
}

// Patch applies the patch and returns the patched aPIPortal.
func (c *FakeAPIPortals) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.APIPortal, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(apiportalsResource, c.ns, name, pt, data, subresources...), &v1alpha1.APIPortal{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APIPortal), err
}
