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

// CasbinPolicyLister helps list CasbinPolicies.
// All objects returned here must be treated as read-only.
type CasbinPolicyLister interface {
	// List lists all CasbinPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.CasbinPolicy, err error)
	// CasbinPolicies returns an object that can list and get CasbinPolicies.
	CasbinPolicies(namespace string) CasbinPolicyNamespaceLister
	CasbinPolicyListerExpansion
}

// casbinPolicyLister implements the CasbinPolicyLister interface.
type casbinPolicyLister struct {
	indexer cache.Indexer
}

// NewCasbinPolicyLister returns a new CasbinPolicyLister.
func NewCasbinPolicyLister(indexer cache.Indexer) CasbinPolicyLister {
	return &casbinPolicyLister{indexer: indexer}
}

// List lists all CasbinPolicies in the indexer.
func (s *casbinPolicyLister) List(selector labels.Selector) (ret []*v1.CasbinPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.CasbinPolicy))
	})
	return ret, err
}

// CasbinPolicies returns an object that can list and get CasbinPolicies.
func (s *casbinPolicyLister) CasbinPolicies(namespace string) CasbinPolicyNamespaceLister {
	return casbinPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CasbinPolicyNamespaceLister helps list and get CasbinPolicies.
// All objects returned here must be treated as read-only.
type CasbinPolicyNamespaceLister interface {
	// List lists all CasbinPolicies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.CasbinPolicy, err error)
	// Get retrieves the CasbinPolicy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.CasbinPolicy, error)
	CasbinPolicyNamespaceListerExpansion
}

// casbinPolicyNamespaceLister implements the CasbinPolicyNamespaceLister
// interface.
type casbinPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CasbinPolicies in the indexer for a given namespace.
func (s casbinPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1.CasbinPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.CasbinPolicy))
	})
	return ret, err
}

// Get retrieves the CasbinPolicy from the indexer for a given namespace and name.
func (s casbinPolicyNamespaceLister) Get(name string) (*v1.CasbinPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("casbinpolicy"), name)
	}
	return obj.(*v1.CasbinPolicy), nil
}