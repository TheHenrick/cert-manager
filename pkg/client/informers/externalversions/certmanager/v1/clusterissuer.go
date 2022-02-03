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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	certmanagerv1 "github.com/cert-manager/cert-manager/pkg/apis/certmanager/v1"
	versioned "github.com/cert-manager/cert-manager/pkg/client/clientset/versioned"
	internalinterfaces "github.com/cert-manager/cert-manager/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/cert-manager/cert-manager/pkg/client/listers/certmanager/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterIssuerInformer provides access to a shared informer and lister for
// ClusterIssuers.
type ClusterIssuerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ClusterIssuerLister
}

type clusterIssuerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewClusterIssuerInformer constructs a new informer for ClusterIssuer type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterIssuerInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterIssuerInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredClusterIssuerInformer constructs a new informer for ClusterIssuer type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterIssuerInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CertmanagerV1().ClusterIssuers().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CertmanagerV1().ClusterIssuers().Watch(context.TODO(), options)
			},
		},
		&certmanagerv1.ClusterIssuer{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterIssuerInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterIssuerInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterIssuerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&certmanagerv1.ClusterIssuer{}, f.defaultInformer)
}

func (f *clusterIssuerInformer) Lister() v1.ClusterIssuerLister {
	return v1.NewClusterIssuerLister(f.Informer().GetIndexer())
}
