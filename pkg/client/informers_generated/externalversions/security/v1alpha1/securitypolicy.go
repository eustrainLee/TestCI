/*
Copyright 2021 The Everoute Authors.

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

package v1alpha1

import (
	"context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"

	securityv1alpha1 "github.com/everoute/everoute/pkg/apis/security/v1alpha1"
	clientset "github.com/everoute/everoute/pkg/client/clientset_generated/clientset"
	internalinterfaces "github.com/everoute/everoute/pkg/client/informers_generated/externalversions/internalinterfaces"
	v1alpha1 "github.com/everoute/everoute/pkg/client/listers_generated/security/v1alpha1"
)

// SecurityPolicyInformer provides access to a shared informer and lister for
// SecurityPolicies.
type SecurityPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.SecurityPolicyLister
}

type securityPolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSecurityPolicyInformer constructs a new informer for SecurityPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSecurityPolicyInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSecurityPolicyInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSecurityPolicyInformer constructs a new informer for SecurityPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSecurityPolicyInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SecurityV1alpha1().SecurityPolicies(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SecurityV1alpha1().SecurityPolicies(namespace).Watch(context.TODO(), options)
			},
		},
		&securityv1alpha1.SecurityPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *securityPolicyInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSecurityPolicyInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *securityPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&securityv1alpha1.SecurityPolicy{}, f.defaultInformer)
}

func (f *securityPolicyInformer) Lister() v1alpha1.SecurityPolicyLister {
	return v1alpha1.NewSecurityPolicyLister(f.Informer().GetIndexer())
}
