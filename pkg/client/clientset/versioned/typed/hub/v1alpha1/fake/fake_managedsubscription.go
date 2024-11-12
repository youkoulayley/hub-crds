/*
The GNU AFFERO GENERAL PUBLIC LICENSE

Copyright (c) 2020-2024 Traefik Labs

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

// FakeManagedSubscriptions implements ManagedSubscriptionInterface
type FakeManagedSubscriptions struct {
	Fake *FakeHubV1alpha1
	ns   string
}

var managedsubscriptionsResource = v1alpha1.SchemeGroupVersion.WithResource("managedsubscriptions")

var managedsubscriptionsKind = v1alpha1.SchemeGroupVersion.WithKind("ManagedSubscription")

// Get takes name of the managedSubscription, and returns the corresponding managedSubscription object, and an error if there is any.
func (c *FakeManagedSubscriptions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ManagedSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(managedsubscriptionsResource, c.ns, name), &v1alpha1.ManagedSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagedSubscription), err
}

// List takes label and field selectors, and returns the list of ManagedSubscriptions that match those selectors.
func (c *FakeManagedSubscriptions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ManagedSubscriptionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(managedsubscriptionsResource, managedsubscriptionsKind, c.ns, opts), &v1alpha1.ManagedSubscriptionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ManagedSubscriptionList{ListMeta: obj.(*v1alpha1.ManagedSubscriptionList).ListMeta}
	for _, item := range obj.(*v1alpha1.ManagedSubscriptionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested managedSubscriptions.
func (c *FakeManagedSubscriptions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(managedsubscriptionsResource, c.ns, opts))

}

// Create takes the representation of a managedSubscription and creates it.  Returns the server's representation of the managedSubscription, and an error, if there is any.
func (c *FakeManagedSubscriptions) Create(ctx context.Context, managedSubscription *v1alpha1.ManagedSubscription, opts v1.CreateOptions) (result *v1alpha1.ManagedSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(managedsubscriptionsResource, c.ns, managedSubscription), &v1alpha1.ManagedSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagedSubscription), err
}

// Update takes the representation of a managedSubscription and updates it. Returns the server's representation of the managedSubscription, and an error, if there is any.
func (c *FakeManagedSubscriptions) Update(ctx context.Context, managedSubscription *v1alpha1.ManagedSubscription, opts v1.UpdateOptions) (result *v1alpha1.ManagedSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(managedsubscriptionsResource, c.ns, managedSubscription), &v1alpha1.ManagedSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagedSubscription), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeManagedSubscriptions) UpdateStatus(ctx context.Context, managedSubscription *v1alpha1.ManagedSubscription, opts v1.UpdateOptions) (*v1alpha1.ManagedSubscription, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(managedsubscriptionsResource, "status", c.ns, managedSubscription), &v1alpha1.ManagedSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagedSubscription), err
}

// Delete takes name of the managedSubscription and deletes it. Returns an error if one occurs.
func (c *FakeManagedSubscriptions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(managedsubscriptionsResource, c.ns, name, opts), &v1alpha1.ManagedSubscription{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeManagedSubscriptions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(managedsubscriptionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ManagedSubscriptionList{})
	return err
}

// Patch applies the patch and returns the patched managedSubscription.
func (c *FakeManagedSubscriptions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ManagedSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(managedsubscriptionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ManagedSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagedSubscription), err
}