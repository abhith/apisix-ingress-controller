// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by lister-gen. DO NOT EDIT.

package v2beta3

import (
	v2beta3 "github.com/apache/apisix-ingress-controller/pkg/kube/apisix/apis/config/v2beta3"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ApisixConsumerLister helps list ApisixConsumers.
// All objects returned here must be treated as read-only.
type ApisixConsumerLister interface {
	// List lists all ApisixConsumers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v2beta3.ApisixConsumer, err error)
	// ApisixConsumers returns an object that can list and get ApisixConsumers.
	ApisixConsumers(namespace string) ApisixConsumerNamespaceLister
	ApisixConsumerListerExpansion
}

// apisixConsumerLister implements the ApisixConsumerLister interface.
type apisixConsumerLister struct {
	indexer cache.Indexer
}

// NewApisixConsumerLister returns a new ApisixConsumerLister.
func NewApisixConsumerLister(indexer cache.Indexer) ApisixConsumerLister {
	return &apisixConsumerLister{indexer: indexer}
}

// List lists all ApisixConsumers in the indexer.
func (s *apisixConsumerLister) List(selector labels.Selector) (ret []*v2beta3.ApisixConsumer, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v2beta3.ApisixConsumer))
	})
	return ret, err
}

// ApisixConsumers returns an object that can list and get ApisixConsumers.
func (s *apisixConsumerLister) ApisixConsumers(namespace string) ApisixConsumerNamespaceLister {
	return apisixConsumerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ApisixConsumerNamespaceLister helps list and get ApisixConsumers.
// All objects returned here must be treated as read-only.
type ApisixConsumerNamespaceLister interface {
	// List lists all ApisixConsumers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v2beta3.ApisixConsumer, err error)
	// Get retrieves the ApisixConsumer from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v2beta3.ApisixConsumer, error)
	ApisixConsumerNamespaceListerExpansion
}

// apisixConsumerNamespaceLister implements the ApisixConsumerNamespaceLister
// interface.
type apisixConsumerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ApisixConsumers in the indexer for a given namespace.
func (s apisixConsumerNamespaceLister) List(selector labels.Selector) (ret []*v2beta3.ApisixConsumer, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v2beta3.ApisixConsumer))
	})
	return ret, err
}

// Get retrieves the ApisixConsumer from the indexer for a given namespace and name.
func (s apisixConsumerNamespaceLister) Get(name string) (*v2beta3.ApisixConsumer, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v2beta3.Resource("apisixconsumer"), name)
	}
	return obj.(*v2beta3.ApisixConsumer), nil
}
