// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/JordanGunn/dask-kubernetes/v2024/dask_kubernetes/operator/go_client/pkg/apis/kubernetes.dask.org/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DaskClusterLister helps list DaskClusters.
// All objects returned here must be treated as read-only.
type DaskClusterLister interface {
	// List lists all DaskClusters in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.DaskCluster, err error)
	// DaskClusters returns an object that can list and get DaskClusters.
	DaskClusters(namespace string) DaskClusterNamespaceLister
	DaskClusterListerExpansion
}

// daskClusterLister implements the DaskClusterLister interface.
type daskClusterLister struct {
	indexer cache.Indexer
}

// NewDaskClusterLister returns a new DaskClusterLister.
func NewDaskClusterLister(indexer cache.Indexer) DaskClusterLister {
	return &daskClusterLister{indexer: indexer}
}

// List lists all DaskClusters in the indexer.
func (s *daskClusterLister) List(selector labels.Selector) (ret []*v1.DaskCluster, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.DaskCluster))
	})
	return ret, err
}

// DaskClusters returns an object that can list and get DaskClusters.
func (s *daskClusterLister) DaskClusters(namespace string) DaskClusterNamespaceLister {
	return daskClusterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DaskClusterNamespaceLister helps list and get DaskClusters.
// All objects returned here must be treated as read-only.
type DaskClusterNamespaceLister interface {
	// List lists all DaskClusters in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.DaskCluster, err error)
	// Get retrieves the DaskCluster from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.DaskCluster, error)
	DaskClusterNamespaceListerExpansion
}

// daskClusterNamespaceLister implements the DaskClusterNamespaceLister
// interface.
type daskClusterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DaskClusters in the indexer for a given namespace.
func (s daskClusterNamespaceLister) List(selector labels.Selector) (ret []*v1.DaskCluster, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.DaskCluster))
	})
	return ret, err
}

// Get retrieves the DaskCluster from the indexer for a given namespace and name.
func (s daskClusterNamespaceLister) Get(name string) (*v1.DaskCluster, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("daskcluster"), name)
	}
	return obj.(*v1.DaskCluster), nil
}
