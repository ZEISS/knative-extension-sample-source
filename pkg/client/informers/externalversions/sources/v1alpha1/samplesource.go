/*
Copyright (c) 2025 ZEISS Group, Sebastian Doell (@katallaxie)

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
	context "context"
	time "time"

	apissourcesv1alpha1 "github.com/zeiss/knative-extension-sample-source/pkg/apis/sources/v1alpha1"
	versioned "github.com/zeiss/knative-extension-sample-source/pkg/client/clientset/versioned"
	internalinterfaces "github.com/zeiss/knative-extension-sample-source/pkg/client/informers/externalversions/internalinterfaces"
	sourcesv1alpha1 "github.com/zeiss/knative-extension-sample-source/pkg/client/listers/sources/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// SampleSourceInformer provides access to a shared informer and lister for
// SampleSources.
type SampleSourceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() sourcesv1alpha1.SampleSourceLister
}

type sampleSourceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSampleSourceInformer constructs a new informer for SampleSource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSampleSourceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSampleSourceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSampleSourceInformer constructs a new informer for SampleSource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSampleSourceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SamplesV1alpha1().SampleSources(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SamplesV1alpha1().SampleSources(namespace).Watch(context.TODO(), options)
			},
		},
		&apissourcesv1alpha1.SampleSource{},
		resyncPeriod,
		indexers,
	)
}

func (f *sampleSourceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSampleSourceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *sampleSourceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apissourcesv1alpha1.SampleSource{}, f.defaultInformer)
}

func (f *sampleSourceInformer) Lister() sourcesv1alpha1.SampleSourceLister {
	return sourcesv1alpha1.NewSampleSourceLister(f.Informer().GetIndexer())
}
