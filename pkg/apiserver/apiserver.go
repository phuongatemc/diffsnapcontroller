package apiserver

import (
	"example.com/differentialsnapshot/pkg/apis/cbt/install"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	genericapiserver "k8s.io/apiserver/pkg/server"
)

var (
	Scheme = runtime.NewScheme()
	Codecs = serializer.NewCodecFactory(Scheme)
)

func init() {
	install.Install(Scheme)

	// we need to add the options to empty v1
	// to satisfy the generic API server from wanting this
	metav1.AddToGroupVersion(Scheme,
		schema.GroupVersion{Version: "v1"})

	Scheme.AddUnversionedTypes(
		schema.GroupVersion{Group: "", Version: "v1"},
		&metav1.Status{},
		&metav1.APIVersions{},
		&metav1.APIGroupList{},
		&metav1.APIGroup{},
		&metav1.APIResourceList{},
	)
}

// Server contains the state for a Kubernetes custom api server.
type Server struct {
	GenericAPIServer *genericapiserver.GenericAPIServer
}
