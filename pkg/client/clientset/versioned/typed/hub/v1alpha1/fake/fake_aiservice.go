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

// FakeAIServices implements AIServiceInterface
type FakeAIServices struct {
	Fake *FakeHubV1alpha1
	ns   string
}

var aiservicesResource = v1alpha1.SchemeGroupVersion.WithResource("aiservices")

var aiservicesKind = v1alpha1.SchemeGroupVersion.WithKind("AIService")

// Get takes name of the aIService, and returns the corresponding aIService object, and an error if there is any.
func (c *FakeAIServices) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AIService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(aiservicesResource, c.ns, name), &v1alpha1.AIService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AIService), err
}

// List takes label and field selectors, and returns the list of AIServices that match those selectors.
func (c *FakeAIServices) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AIServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(aiservicesResource, aiservicesKind, c.ns, opts), &v1alpha1.AIServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AIServiceList{ListMeta: obj.(*v1alpha1.AIServiceList).ListMeta}
	for _, item := range obj.(*v1alpha1.AIServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested aIServices.
func (c *FakeAIServices) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(aiservicesResource, c.ns, opts))

}

// Create takes the representation of a aIService and creates it.  Returns the server's representation of the aIService, and an error, if there is any.
func (c *FakeAIServices) Create(ctx context.Context, aIService *v1alpha1.AIService, opts v1.CreateOptions) (result *v1alpha1.AIService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(aiservicesResource, c.ns, aIService), &v1alpha1.AIService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AIService), err
}

// Update takes the representation of a aIService and updates it. Returns the server's representation of the aIService, and an error, if there is any.
func (c *FakeAIServices) Update(ctx context.Context, aIService *v1alpha1.AIService, opts v1.UpdateOptions) (result *v1alpha1.AIService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(aiservicesResource, c.ns, aIService), &v1alpha1.AIService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AIService), err
}

// Delete takes name of the aIService and deletes it. Returns an error if one occurs.
func (c *FakeAIServices) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(aiservicesResource, c.ns, name, opts), &v1alpha1.AIService{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAIServices) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(aiservicesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AIServiceList{})
	return err
}

// Patch applies the patch and returns the patched aIService.
func (c *FakeAIServices) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AIService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(aiservicesResource, c.ns, name, pt, data, subresources...), &v1alpha1.AIService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AIService), err
}
