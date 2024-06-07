// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/JordanGunn/dask-kubernetes/v2024/dask_kubernetes/operator/go_client/pkg/apis/kubernetes.dask.org/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DaskAutoscalerLister helps list DaskAutoscalers.
// All objects returned here must be treated as read-only.
type DaskAutoscalerLister interface {
	// List lists all DaskAutoscalers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.DaskAutoscaler, err error)
	// DaskAutoscalers returns an object that can list and get DaskAutoscalers.
	DaskAutoscalers(namespace string) DaskAutoscalerNamespaceLister
	DaskAutoscalerListerExpansion
}

// daskAutoscalerLister implements the DaskAutoscalerLister interface.
type daskAutoscalerLister struct {
	indexer cache.Indexer
}

// NewDaskAutoscalerLister returns a new DaskAutoscalerLister.
func NewDaskAutoscalerLister(indexer cache.Indexer) DaskAutoscalerLister {
	return &daskAutoscalerLister{indexer: indexer}
}

// List lists all DaskAutoscalers in the indexer.
func (s *daskAutoscalerLister) List(selector labels.Selector) (ret []*v1.DaskAutoscaler, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.DaskAutoscaler))
	})
	return ret, err
}

// DaskAutoscalers returns an object that can list and get DaskAutoscalers.
func (s *daskAutoscalerLister) DaskAutoscalers(namespace string) DaskAutoscalerNamespaceLister {
	return daskAutoscalerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DaskAutoscalerNamespaceLister helps list and get DaskAutoscalers.
// All objects returned here must be treated as read-only.
type DaskAutoscalerNamespaceLister interface {
	// List lists all DaskAutoscalers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.DaskAutoscaler, err error)
	// Get retrieves the DaskAutoscaler from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.DaskAutoscaler, error)
	DaskAutoscalerNamespaceListerExpansion
}

// daskAutoscalerNamespaceLister implements the DaskAutoscalerNamespaceLister
// interface.
type daskAutoscalerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DaskAutoscalers in the indexer for a given namespace.
func (s daskAutoscalerNamespaceLister) List(selector labels.Selector) (ret []*v1.DaskAutoscaler, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.DaskAutoscaler))
	})
	return ret, err
}

// Get retrieves the DaskAutoscaler from the indexer for a given namespace and name.
func (s daskAutoscalerNamespaceLister) Get(name string) (*v1.DaskAutoscaler, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("daskautoscaler"), name)
	}
	return obj.(*v1.DaskAutoscaler), nil
}
