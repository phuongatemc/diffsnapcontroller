package initializer

import (
	informers "example.com/differentialsnapshot/pkg/generated/cbt/informers/externalversions"
	"k8s.io/apiserver/pkg/admission"
)

// WantsStorageInformerFactory defines a function which sets the storage
// InformerFactory for admission plugins that need it.
type WantsStorageInformerFactory interface {
	SetStorageInformerFactory(informers.SharedInformerFactory)
	admission.InitializationValidator
}

type storageInformerPluginInitializer struct {
	informers informers.SharedInformerFactory
}

var _ admission.PluginInitializer = storageInformerPluginInitializer{}

// New creates an instance of the custom admission plugins initializer.
func New(informers informers.SharedInformerFactory) storageInformerPluginInitializer {
	return storageInformerPluginInitializer{
		informers: informers,
	}
}

// Initialize checks the initialization interfaces implemented by a plugin and
// provide the appropriate initialization data.
func (s storageInformerPluginInitializer) Initialize(plugin admission.Interface) {
	if wants, ok := plugin.(WantsStorageInformerFactory); ok {
		wants.SetStorageInformerFactory(s.informers)
	}
}
