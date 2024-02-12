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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/traefik/hub-crds/pkg/apis/hub/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// APILister helps list APIs.
// All objects returned here must be treated as read-only.
type APILister interface {
	// List lists all APIs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.API, err error)
	// APIs returns an object that can list and get APIs.
	APIs(namespace string) APINamespaceLister
	APIListerExpansion
}

// aPILister implements the APILister interface.
type aPILister struct {
	indexer cache.Indexer
}

// NewAPILister returns a new APILister.
func NewAPILister(indexer cache.Indexer) APILister {
	return &aPILister{indexer: indexer}
}

// List lists all APIs in the indexer.
func (s *aPILister) List(selector labels.Selector) (ret []*v1alpha1.API, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.API))
	})
	return ret, err
}

// APIs returns an object that can list and get APIs.
func (s *aPILister) APIs(namespace string) APINamespaceLister {
	return aPINamespaceLister{indexer: s.indexer, namespace: namespace}
}

// APINamespaceLister helps list and get APIs.
// All objects returned here must be treated as read-only.
type APINamespaceLister interface {
	// List lists all APIs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.API, err error)
	// Get retrieves the API from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.API, error)
	APINamespaceListerExpansion
}

// aPINamespaceLister implements the APINamespaceLister
// interface.
type aPINamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all APIs in the indexer for a given namespace.
func (s aPINamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.API, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.API))
	})
	return ret, err
}

// Get retrieves the API from the indexer for a given namespace and name.
func (s aPINamespaceLister) Get(name string) (*v1alpha1.API, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("api"), name)
	}
	return obj.(*v1alpha1.API), nil
}