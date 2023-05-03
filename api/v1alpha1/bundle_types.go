/*
Copyright 2023 VMware Inc.

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

package v1alpha1

import (
	"github.com/vmware-labs/reconciler-runtime/apis"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Source struct {
	Ref      SourceRef `json:"source"`
	FromPath string    `json:"fromPath"`
	ToPath   string    `json:"toPath"`
}

func (in *Source) Key(namespace string) client.ObjectKey {
	return client.ObjectKey{Name: in.Ref.Name, Namespace: namespace}
}

// SourceRef the GVK and name of the source object
type SourceRef struct {
	ApiVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Name       string `json:"name"`
}

// BundleSpec defines the structure of the bundle
type BundleSpec struct {
	Include []Source `json:"include"`
}

// BundleStatus defines the observed state of Bundle
type BundleStatus struct {
	apis.Status `json:",inline"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Bundle is the Schema for the bundle API
type Bundle struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BundleSpec   `json:"spec,omitempty"`
	Status BundleStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// BundleList contains a list of Bundle
type BundleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Bundle `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Bundle{}, &BundleList{})
}
