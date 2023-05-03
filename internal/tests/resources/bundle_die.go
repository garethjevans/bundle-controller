package resources

import (
	v1 "dies.dev/apis/meta/v1"
	"github.com/garethjevans/bundle-controller/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +die:object=true
type _ = v1alpha1.Bundle

// +die
type _ = v1alpha1.BundleSpec

// +die
type _ = v1alpha1.BundleStatus

func (d *BundleStatusDie) ConditionsDie(conditions ...*v1.ConditionDie) *BundleStatusDie {
	return d.DieStamp(func(r *v1alpha1.BundleStatus) {
		r.Conditions = make([]metav1.Condition, len(conditions))
		for i := range conditions {
			r.Conditions[i] = conditions[i].DieRelease()
		}
	})
}

var (
	BundleReadyBlank           = v1.ConditionBlank.Type(v1alpha1.BundleConditionReady)
	BundleResourceMappingBlank = v1.ConditionBlank.Type(v1alpha1.BundleResourceMapping)
)
