/*
Copyright The Kubernetes Authors.

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

	cbtv1alpha1 "example.com/differentialsnapshot/pkg/apis/cbt/v1alpha1"
	versioned "example.com/differentialsnapshot/pkg/generated/cbt/clientset/versioned"
	internalinterfaces "example.com/differentialsnapshot/pkg/generated/cbt/informers/externalversions/internalinterfaces"
	v1alpha1 "example.com/differentialsnapshot/pkg/generated/cbt/listers/cbt/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ChangedBlockInformer provides access to a shared informer and lister for
// ChangedBlocks.
type ChangedBlockInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ChangedBlockLister
}

type changedBlockInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewChangedBlockInformer constructs a new informer for ChangedBlock type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewChangedBlockInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredChangedBlockInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredChangedBlockInformer constructs a new informer for ChangedBlock type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredChangedBlockInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CbtV1alpha1().ChangedBlocks(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CbtV1alpha1().ChangedBlocks(namespace).Watch(context.TODO(), options)
			},
		},
		&cbtv1alpha1.ChangedBlock{},
		resyncPeriod,
		indexers,
	)
}

func (f *changedBlockInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredChangedBlockInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *changedBlockInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&cbtv1alpha1.ChangedBlock{}, f.defaultInformer)
}

func (f *changedBlockInformer) Lister() v1alpha1.ChangedBlockLister {
	return v1alpha1.NewChangedBlockLister(f.Informer().GetIndexer())
}