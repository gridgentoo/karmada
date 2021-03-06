// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/karmada-io/karmada/pkg/apis/networking/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MultiClusterIngressLister helps list MultiClusterIngresses.
// All objects returned here must be treated as read-only.
type MultiClusterIngressLister interface {
	// List lists all MultiClusterIngresses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.MultiClusterIngress, err error)
	// MultiClusterIngresses returns an object that can list and get MultiClusterIngresses.
	MultiClusterIngresses(namespace string) MultiClusterIngressNamespaceLister
	MultiClusterIngressListerExpansion
}

// multiClusterIngressLister implements the MultiClusterIngressLister interface.
type multiClusterIngressLister struct {
	indexer cache.Indexer
}

// NewMultiClusterIngressLister returns a new MultiClusterIngressLister.
func NewMultiClusterIngressLister(indexer cache.Indexer) MultiClusterIngressLister {
	return &multiClusterIngressLister{indexer: indexer}
}

// List lists all MultiClusterIngresses in the indexer.
func (s *multiClusterIngressLister) List(selector labels.Selector) (ret []*v1alpha1.MultiClusterIngress, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MultiClusterIngress))
	})
	return ret, err
}

// MultiClusterIngresses returns an object that can list and get MultiClusterIngresses.
func (s *multiClusterIngressLister) MultiClusterIngresses(namespace string) MultiClusterIngressNamespaceLister {
	return multiClusterIngressNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MultiClusterIngressNamespaceLister helps list and get MultiClusterIngresses.
// All objects returned here must be treated as read-only.
type MultiClusterIngressNamespaceLister interface {
	// List lists all MultiClusterIngresses in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.MultiClusterIngress, err error)
	// Get retrieves the MultiClusterIngress from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.MultiClusterIngress, error)
	MultiClusterIngressNamespaceListerExpansion
}

// multiClusterIngressNamespaceLister implements the MultiClusterIngressNamespaceLister
// interface.
type multiClusterIngressNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MultiClusterIngresses in the indexer for a given namespace.
func (s multiClusterIngressNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MultiClusterIngress, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MultiClusterIngress))
	})
	return ret, err
}

// Get retrieves the MultiClusterIngress from the indexer for a given namespace and name.
func (s multiClusterIngressNamespaceLister) Get(name string) (*v1alpha1.MultiClusterIngress, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("multiclusteringress"), name)
	}
	return obj.(*v1alpha1.MultiClusterIngress), nil
}
