package registry

import (
	"example.com/differentialsnapshot/pkg/apis/cbt"
	reststorage "example.com/differentialsnapshot/pkg/apiserver/storage"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/apiserver/pkg/registry/rest"
)

// REST implements a RESTStorage for API services.
type REST struct {
	watch chan watch.Event
	reststorage.Memory
	rest.TableConvertor
}

// NewRest returns a RESTStorage object that will work against API services.
func NewREST() *REST {
	return &REST{
		make(chan watch.Event),
		reststorage.NewMemory(),
		rest.NewDefaultTableConvertor(cbt.Resource("changedblocks")),
	}
}
