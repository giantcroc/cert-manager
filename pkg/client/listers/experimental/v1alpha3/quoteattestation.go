/*
Copyright The cert-manager Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha3

import (
	v1alpha3 "github.com/cert-manager/cert-manager/pkg/apis/experimental/v1alpha3"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// QuoteAttestationLister helps list QuoteAttestations.
// All objects returned here must be treated as read-only.
type QuoteAttestationLister interface {
	// List lists all QuoteAttestations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha3.QuoteAttestation, err error)
	// QuoteAttestations returns an object that can list and get QuoteAttestations.
	QuoteAttestations(namespace string) QuoteAttestationNamespaceLister
	QuoteAttestationListerExpansion
}

// quoteAttestationLister implements the QuoteAttestationLister interface.
type quoteAttestationLister struct {
	indexer cache.Indexer
}

// NewQuoteAttestationLister returns a new QuoteAttestationLister.
func NewQuoteAttestationLister(indexer cache.Indexer) QuoteAttestationLister {
	return &quoteAttestationLister{indexer: indexer}
}

// List lists all QuoteAttestations in the indexer.
func (s *quoteAttestationLister) List(selector labels.Selector) (ret []*v1alpha3.QuoteAttestation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha3.QuoteAttestation))
	})
	return ret, err
}

// QuoteAttestations returns an object that can list and get QuoteAttestations.
func (s *quoteAttestationLister) QuoteAttestations(namespace string) QuoteAttestationNamespaceLister {
	return quoteAttestationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// QuoteAttestationNamespaceLister helps list and get QuoteAttestations.
// All objects returned here must be treated as read-only.
type QuoteAttestationNamespaceLister interface {
	// List lists all QuoteAttestations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha3.QuoteAttestation, err error)
	// Get retrieves the QuoteAttestation from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha3.QuoteAttestation, error)
	QuoteAttestationNamespaceListerExpansion
}

// quoteAttestationNamespaceLister implements the QuoteAttestationNamespaceLister
// interface.
type quoteAttestationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all QuoteAttestations in the indexer for a given namespace.
func (s quoteAttestationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha3.QuoteAttestation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha3.QuoteAttestation))
	})
	return ret, err
}

// Get retrieves the QuoteAttestation from the indexer for a given namespace and name.
func (s quoteAttestationNamespaceLister) Get(name string) (*v1alpha3.QuoteAttestation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha3.Resource("quoteattestation"), name)
	}
	return obj.(*v1alpha3.QuoteAttestation), nil
}
