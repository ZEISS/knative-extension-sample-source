package v1alpha1

import (
	"testing"

	"k8s.io/apimachinery/pkg/runtime"
)

func TestRegisterHelpers(t *testing.T) {
	if got, want := Kind("SampleSource"), "SampleSource.sources.eventing.zeiss.com"; got.String() != want {
		t.Errorf("Kind(SampleSource) = %v, want %v", got.String(), want)
	}

	if got, want := Resource("SampleSource"), "SampleSource.sources.eventing.zeiss.com"; got.String() != want {
		t.Errorf("Resource(SampleSource) = %v, want %v", got.String(), want)
	}

	if got, want := SchemeGroupVersion.String(), "sources.eventing.zeiss.com/v1alpha1"; got != want {
		t.Errorf("SchemeGroupVersion() = %v, want %v", got, want)
	}

	scheme := runtime.NewScheme()
	if err := addKnownTypes(scheme); err != nil {
		t.Errorf("addKnownTypes() = %v", err)
	}
}
