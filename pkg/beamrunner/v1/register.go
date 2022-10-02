package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	beamrunner "github.com/slai-labs/beamcrds/pkg/beamrunner"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	SchemeGroupVersion = schema.GroupVersion{
		Group:   beamrunner.GroupName,
		Version: beamrunner.V1,
	}
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)

func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&BeamRunner{},
		&BeamRunnerList{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
