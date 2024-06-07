// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"net/http"

	v1 "github.com/JordanGunn/dask-kubernetes/v2024/dask_kubernetes/operator/go_client/pkg/apis/kubernetes.dask.org/v1"
	"github.com/JordanGunn/dask-kubernetes/v2024/dask_kubernetes/operator/go_client/pkg/client/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type KubernetesV1Interface interface {
	RESTClient() rest.Interface
	DaskAutoscalersGetter
	DaskClustersGetter
	DaskJobsGetter
	DaskWorkerGroupsGetter
}

// KubernetesV1Client is used to interact with features provided by the kubernetes.dask.org group.
type KubernetesV1Client struct {
	restClient rest.Interface
}

func (c *KubernetesV1Client) DaskAutoscalers(namespace string) DaskAutoscalerInterface {
	return newDaskAutoscalers(c, namespace)
}

func (c *KubernetesV1Client) DaskClusters(namespace string) DaskClusterInterface {
	return newDaskClusters(c, namespace)
}

func (c *KubernetesV1Client) DaskJobs(namespace string) DaskJobInterface {
	return newDaskJobs(c, namespace)
}

func (c *KubernetesV1Client) DaskWorkerGroups(namespace string) DaskWorkerGroupInterface {
	return newDaskWorkerGroups(c, namespace)
}

// NewForConfig creates a new KubernetesV1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*KubernetesV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new KubernetesV1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*KubernetesV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &KubernetesV1Client{client}, nil
}

// NewForConfigOrDie creates a new KubernetesV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *KubernetesV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new KubernetesV1Client for the given RESTClient.
func New(c rest.Interface) *KubernetesV1Client {
	return &KubernetesV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *KubernetesV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
