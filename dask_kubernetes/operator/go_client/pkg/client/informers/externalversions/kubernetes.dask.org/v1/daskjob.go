// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	kubernetesdaskorgv1 "github.com/geoanalytics-ca/dask-kubernetes/v2024/dask_kubernetes/operator/go_client/pkg/apis/kubernetes.dask.org/v1"
	versioned "github.com/geoanalytics-ca/dask-kubernetes/v2024/dask_kubernetes/operator/go_client/pkg/client/clientset/versioned"
	internalinterfaces "github.com/geoanalytics-ca/dask-kubernetes/v2024/dask_kubernetes/operator/go_client/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/geoanalytics-ca/dask-kubernetes/v2024/dask_kubernetes/operator/go_client/pkg/client/listers/kubernetes.dask.org/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// DaskJobInformer provides access to a shared informer and lister for
// DaskJobs.
type DaskJobInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.DaskJobLister
}

type daskJobInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewDaskJobInformer constructs a new informer for DaskJob type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDaskJobInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDaskJobInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredDaskJobInformer constructs a new informer for DaskJob type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDaskJobInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubernetesV1().DaskJobs(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubernetesV1().DaskJobs(namespace).Watch(context.TODO(), options)
			},
		},
		&kubernetesdaskorgv1.DaskJob{},
		resyncPeriod,
		indexers,
	)
}

func (f *daskJobInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDaskJobInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *daskJobInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&kubernetesdaskorgv1.DaskJob{}, f.defaultInformer)
}

func (f *daskJobInformer) Lister() v1.DaskJobLister {
	return v1.NewDaskJobLister(f.Informer().GetIndexer())
}
