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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/casbin/k8s-gatekeeper/pkg/apis/k8sauthz/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CasbinModelLister helps list CasbinModels.
// All objects returned here must be treated as read-only.
type CasbinModelLister interface {
	// List lists all CasbinModels in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.CasbinModel, err error)
	// CasbinModels returns an object that can list and get CasbinModels.
	CasbinModels(namespace string) CasbinModelNamespaceLister
	CasbinModelListerExpansion
}

// casbinModelLister implements the CasbinModelLister interface.
type casbinModelLister struct {
	indexer cache.Indexer
}

// NewCasbinModelLister returns a new CasbinModelLister.
func NewCasbinModelLister(indexer cache.Indexer) CasbinModelLister {
	return &casbinModelLister{indexer: indexer}
}

// List lists all CasbinModels in the indexer.
func (s *casbinModelLister) List(selector labels.Selector) (ret []*v1.CasbinModel, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.CasbinModel))
	})
	return ret, err
}

// CasbinModels returns an object that can list and get CasbinModels.
func (s *casbinModelLister) CasbinModels(namespace string) CasbinModelNamespaceLister {
	return casbinModelNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CasbinModelNamespaceLister helps list and get CasbinModels.
// All objects returned here must be treated as read-only.
type CasbinModelNamespaceLister interface {
	// List lists all CasbinModels in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.CasbinModel, err error)
	// Get retrieves the CasbinModel from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.CasbinModel, error)
	CasbinModelNamespaceListerExpansion
}

// casbinModelNamespaceLister implements the CasbinModelNamespaceLister
// interface.
type casbinModelNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CasbinModels in the indexer for a given namespace.
func (s casbinModelNamespaceLister) List(selector labels.Selector) (ret []*v1.CasbinModel, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.CasbinModel))
	})
	return ret, err
}

// Get retrieves the CasbinModel from the indexer for a given namespace and name.
func (s casbinModelNamespaceLister) Get(name string) (*v1.CasbinModel, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("casbinmodel"), name)
	}
	return obj.(*v1.CasbinModel), nil
}
