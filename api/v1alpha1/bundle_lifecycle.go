package v1alpha1

import (
	"fmt"
	"github.com/vmware-labs/reconciler-runtime/apis"
)

const (
	BundleConditionReady                       = apis.ConditionReady
	BundleResourceMapping                      = "BundleResourceMapping"
	BundleResourceMappingNoSuchComponentReason = "NoSuchResource"
)

var containerCondSet = apis.NewLivingConditionSet(
	BundleResourceMapping,
)

func (b *BundleStatus) MarkResourceMissing(resource string, component string, namespace string) {
	template := "resource `%s` missing. " +
		"bundle is trying to find resource `%s` in namespace `%s`"

	message := fmt.Sprintf(template, resource, component, namespace)

	containerCondSet.Manage(b).MarkFalse(BundleResourceMapping, BundleResourceMappingNoSuchComponentReason, message)
}

func (b *BundleStatus) MarkFailed(err error) {
	containerCondSet.Manage(b).MarkFalse(BundleConditionReady, "Failed", err.Error())
}
