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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	sourcesv1alpha1 "github.com/zeiss/knative-extension-sample-source/pkg/apis/sources/v1alpha1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// SampleSourceLister helps list SampleSources.
// All objects returned here must be treated as read-only.
type SampleSourceLister interface {
	// List lists all SampleSources in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*sourcesv1alpha1.SampleSource, err error)
	// SampleSources returns an object that can list and get SampleSources.
	SampleSources(namespace string) SampleSourceNamespaceLister
	SampleSourceListerExpansion
}

// sampleSourceLister implements the SampleSourceLister interface.
type sampleSourceLister struct {
	listers.ResourceIndexer[*sourcesv1alpha1.SampleSource]
}

// NewSampleSourceLister returns a new SampleSourceLister.
func NewSampleSourceLister(indexer cache.Indexer) SampleSourceLister {
	return &sampleSourceLister{listers.New[*sourcesv1alpha1.SampleSource](indexer, sourcesv1alpha1.Resource("samplesource"))}
}

// SampleSources returns an object that can list and get SampleSources.
func (s *sampleSourceLister) SampleSources(namespace string) SampleSourceNamespaceLister {
	return sampleSourceNamespaceLister{listers.NewNamespaced[*sourcesv1alpha1.SampleSource](s.ResourceIndexer, namespace)}
}

// SampleSourceNamespaceLister helps list and get SampleSources.
// All objects returned here must be treated as read-only.
type SampleSourceNamespaceLister interface {
	// List lists all SampleSources in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*sourcesv1alpha1.SampleSource, err error)
	// Get retrieves the SampleSource from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*sourcesv1alpha1.SampleSource, error)
	SampleSourceNamespaceListerExpansion
}

// sampleSourceNamespaceLister implements the SampleSourceNamespaceLister
// interface.
type sampleSourceNamespaceLister struct {
	listers.ResourceIndexer[*sourcesv1alpha1.SampleSource]
}
