package cbt

import (
	"fmt"
	"io"

	"example.com/differentialsnapshot/pkg/apis/cbt"
	informers "example.com/differentialsnapshot/pkg/generated/cbt/informers/externalversions"
	listers "example.com/differentialsnapshot/pkg/generated/cbt/listers/cbt/v1alpha1"
	"k8s.io/apiserver/pkg/admission"
)

// Register registers a plugin
func Register(plugins *admission.Plugins) {
	plugins.Register("ChangedBlocks", func(config io.Reader) (admission.Interface, error) {
		return New()
	})
}

// New creates a new changed block admission plugin.
func New() (*changedBlockPlugin, error) {
	return &changedBlockPlugin{
		Handler: admission.NewHandler(admission.Create, admission.Update, admission.Connect),
	}, nil
}

type changedBlockPlugin struct {
	*admission.Handler
	lister listers.ChangedBlockLister
}

func (c *changedBlockPlugin) Validate(a admission.Attributes, _ admission.ObjectInterfaces) error {
	if a.GetKind().GroupKind() != cbt.Kind("ChangedBlock") {
		return nil
	}

	if !c.WaitForReady() {
		return admission.NewForbidden(a, fmt.Errorf("not yet ready to handle request"))
	}

	return nil
}

func (c *changedBlockPlugin) SetStorageInformerFactory(f informers.SharedInformerFactory) {
	c.lister = f.Cbt().V1alpha1().ChangedBlocks().Lister()
	c.SetReadyFunc(f.Cbt().V1alpha1().ChangedBlocks().Informer().HasSynced)
}

func (c *changedBlockPlugin) ValidateInitialization() error {
	if c.lister == nil {
		return fmt.Errorf("missing policy lister")
	}
	return nil
}
