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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/zeiss/knative-extension-sample-source/pkg/apis/sources/v1alpha1"
	sourcesv1alpha1 "github.com/zeiss/knative-extension-sample-source/pkg/client/clientset/versioned/typed/sources/v1alpha1"
	gentype "k8s.io/client-go/gentype"
)

// fakeSampleSources implements SampleSourceInterface
type fakeSampleSources struct {
	*gentype.FakeClientWithList[*v1alpha1.SampleSource, *v1alpha1.SampleSourceList]
	Fake *FakeSamplesV1alpha1
}

func newFakeSampleSources(fake *FakeSamplesV1alpha1, namespace string) sourcesv1alpha1.SampleSourceInterface {
	return &fakeSampleSources{
		gentype.NewFakeClientWithList[*v1alpha1.SampleSource, *v1alpha1.SampleSourceList](
			fake.Fake,
			namespace,
			v1alpha1.SchemeGroupVersion.WithResource("samplesources"),
			v1alpha1.SchemeGroupVersion.WithKind("SampleSource"),
			func() *v1alpha1.SampleSource { return &v1alpha1.SampleSource{} },
			func() *v1alpha1.SampleSourceList { return &v1alpha1.SampleSourceList{} },
			func(dst, src *v1alpha1.SampleSourceList) { dst.ListMeta = src.ListMeta },
			func(list *v1alpha1.SampleSourceList) []*v1alpha1.SampleSource {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v1alpha1.SampleSourceList, items []*v1alpha1.SampleSource) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
