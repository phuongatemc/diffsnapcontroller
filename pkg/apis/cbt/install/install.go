package install

import (
	"example.com/differentialsnapshot/pkg/apis/cbt"
	"example.com/differentialsnapshot/pkg/apis/cbt/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

// Install registers the API group anda dds types to a scheme.
func Install(scheme *runtime.Scheme) {
	utilruntime.Must(cbt.AddToScheme(scheme))
	utilruntime.Must(v1alpha1.AddToScheme(scheme))
	utilruntime.Must(scheme.SetVersionPriority(v1alpha1.SchemeGroupVersion))
}
