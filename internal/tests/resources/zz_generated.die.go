//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by diegen. DO NOT EDIT.

package resources

import (
	"dies.dev/apis/meta/v1"
	json "encoding/json"
	fmtx "fmt"
	"github.com/garethjevans/bundle-controller/api/v1alpha1"
	apis "github.com/vmware-labs/reconciler-runtime/apis"
	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	jsonpath "k8s.io/client-go/util/jsonpath"
	osx "os"
	reflectx "reflect"
	yaml "sigs.k8s.io/yaml"
)

var BundleBlank = (&BundleDie{}).DieFeed(v1alpha1.Bundle{})

type BundleDie struct {
	v1.FrozenObjectMeta
	mutable bool
	r       v1alpha1.Bundle
}

// DieImmutable returns a new die for the current die's state that is either mutable (`false`) or immutable (`true`).
func (d *BundleDie) DieImmutable(immutable bool) *BundleDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

// DieFeed returns a new die with the provided resource.
func (d *BundleDie) DieFeed(r v1alpha1.Bundle) *BundleDie {
	if d.mutable {
		d.FrozenObjectMeta = v1.FreezeObjectMeta(r.ObjectMeta)
		d.r = r
		return d
	}
	return &BundleDie{
		FrozenObjectMeta: v1.FreezeObjectMeta(r.ObjectMeta),
		mutable:          d.mutable,
		r:                r,
	}
}

// DieFeedPtr returns a new die with the provided resource pointer. If the resource is nil, the empty value is used instead.
func (d *BundleDie) DieFeedPtr(r *v1alpha1.Bundle) *BundleDie {
	if r == nil {
		r = &v1alpha1.Bundle{}
	}
	return d.DieFeed(*r)
}

// DieFeedJSON returns a new die with the provided JSON. Panics on error.
func (d *BundleDie) DieFeedJSON(j []byte) *BundleDie {
	r := v1alpha1.Bundle{}
	if err := json.Unmarshal(j, &r); err != nil {
		panic(err)
	}
	return d.DieFeed(r)
}

// DieFeedYAML returns a new die with the provided YAML. Panics on error.
func (d *BundleDie) DieFeedYAML(y []byte) *BundleDie {
	r := v1alpha1.Bundle{}
	if err := yaml.Unmarshal(y, &r); err != nil {
		panic(err)
	}
	return d.DieFeed(r)
}

// DieFeedYAMLFile returns a new die loading YAML from a file path. Panics on error.
func (d *BundleDie) DieFeedYAMLFile(name string) *BundleDie {
	y, err := osx.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return d.DieFeedYAML(y)
}

// DieFeedRawExtension returns the resource managed by the die as an raw extension. Panics on error.
func (d *BundleDie) DieFeedRawExtension(raw runtime.RawExtension) *BundleDie {
	j, err := json.Marshal(raw)
	if err != nil {
		panic(err)
	}
	return d.DieFeedJSON(j)
}

// DieRelease returns the resource managed by the die.
func (d *BundleDie) DieRelease() v1alpha1.Bundle {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

// DieReleasePtr returns a pointer to the resource managed by the die.
func (d *BundleDie) DieReleasePtr() *v1alpha1.Bundle {
	r := d.DieRelease()
	return &r
}

// DieReleaseUnstructured returns the resource managed by the die as an unstructured object. Panics on error.
func (d *BundleDie) DieReleaseUnstructured() *unstructured.Unstructured {
	r := d.DieReleasePtr()
	u, err := runtime.DefaultUnstructuredConverter.ToUnstructured(r)
	if err != nil {
		panic(err)
	}
	return &unstructured.Unstructured{
		Object: u,
	}
}

// DieReleaseJSON returns the resource managed by the die as JSON. Panics on error.
func (d *BundleDie) DieReleaseJSON() []byte {
	r := d.DieReleasePtr()
	j, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return j
}

// DieReleaseYAML returns the resource managed by the die as YAML. Panics on error.
func (d *BundleDie) DieReleaseYAML() []byte {
	r := d.DieReleasePtr()
	y, err := yaml.Marshal(r)
	if err != nil {
		panic(err)
	}
	return y
}

// DieReleaseRawExtension returns the resource managed by the die as an raw extension. Panics on error.
func (d *BundleDie) DieReleaseRawExtension() runtime.RawExtension {
	j := d.DieReleaseJSON()
	raw := runtime.RawExtension{}
	if err := json.Unmarshal(j, &raw); err != nil {
		panic(err)
	}
	return raw
}

// DieStamp returns a new die with the resource passed to the callback function. The resource is mutable.
func (d *BundleDie) DieStamp(fn func(r *v1alpha1.Bundle)) *BundleDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

// Experimental: DieStampAt uses a JSON path (http://goessner.net/articles/JsonPath/) expression to stamp portions of the resource. The callback is invoked with each JSON path match. Panics if the callback function does not accept a single argument of the same type as found on the resource at the target location.
//
// Future iterations will improve type coercion from the resource to the callback argument.
func (d *BundleDie) DieStampAt(jp string, fn interface{}) *BundleDie {
	return d.DieStamp(func(r *v1alpha1.Bundle) {
		cp := jsonpath.New("")
		if err := cp.Parse(fmtx.Sprintf("{%s}", jp)); err != nil {
			panic(err)
		}
		cr, err := cp.FindResults(r)
		if err != nil {
			// errors are expected if a path is not found
			return
		}
		for _, cv := range cr[0] {
			args := []reflectx.Value{cv}
			reflectx.ValueOf(fn).Call(args)
		}
	})
}

// DeepCopy returns a new die with equivalent state. Useful for snapshotting a mutable die.
func (d *BundleDie) DeepCopy() *BundleDie {
	r := *d.r.DeepCopy()
	return &BundleDie{
		FrozenObjectMeta: v1.FreezeObjectMeta(r.ObjectMeta),
		mutable:          d.mutable,
		r:                r,
	}
}

var _ runtime.Object = (*BundleDie)(nil)

func (d *BundleDie) DeepCopyObject() runtime.Object {
	return d.r.DeepCopy()
}

func (d *BundleDie) GetObjectKind() schema.ObjectKind {
	r := d.DieRelease()
	return r.GetObjectKind()
}

func (d *BundleDie) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.r)
}

func (d *BundleDie) UnmarshalJSON(b []byte) error {
	if d == BundleBlank {
		return fmtx.Errorf("cannot unmarshal into the blank die, create a copy first")
	}
	if !d.mutable {
		return fmtx.Errorf("cannot unmarshal into immutable dies, create a mutable version first")
	}
	r := &v1alpha1.Bundle{}
	err := json.Unmarshal(b, r)
	*d = *d.DieFeed(*r)
	return err
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (d *BundleDie) APIVersion(v string) *BundleDie {
	return d.DieStamp(func(r *v1alpha1.Bundle) {
		r.APIVersion = v
	})
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (d *BundleDie) Kind(v string) *BundleDie {
	return d.DieStamp(func(r *v1alpha1.Bundle) {
		r.Kind = v
	})
}

// MetadataDie stamps the resource's ObjectMeta field with a mutable die.
func (d *BundleDie) MetadataDie(fn func(d *v1.ObjectMetaDie)) *BundleDie {
	return d.DieStamp(func(r *v1alpha1.Bundle) {
		d := v1.ObjectMetaBlank.DieImmutable(false).DieFeed(r.ObjectMeta)
		fn(d)
		r.ObjectMeta = d.DieRelease()
	})
}

// SpecDie stamps the resource's spec field with a mutable die.
func (d *BundleDie) SpecDie(fn func(d *BundleSpecDie)) *BundleDie {
	return d.DieStamp(func(r *v1alpha1.Bundle) {
		d := BundleSpecBlank.DieImmutable(false).DieFeed(r.Spec)
		fn(d)
		r.Spec = d.DieRelease()
	})
}

// StatusDie stamps the resource's status field with a mutable die.
func (d *BundleDie) StatusDie(fn func(d *BundleStatusDie)) *BundleDie {
	return d.DieStamp(func(r *v1alpha1.Bundle) {
		d := BundleStatusBlank.DieImmutable(false).DieFeed(r.Status)
		fn(d)
		r.Status = d.DieRelease()
	})
}

func (d *BundleDie) Spec(v v1alpha1.BundleSpec) *BundleDie {
	return d.DieStamp(func(r *v1alpha1.Bundle) {
		r.Spec = v
	})
}

func (d *BundleDie) Status(v v1alpha1.BundleStatus) *BundleDie {
	return d.DieStamp(func(r *v1alpha1.Bundle) {
		r.Status = v
	})
}

var BundleSpecBlank = (&BundleSpecDie{}).DieFeed(v1alpha1.BundleSpec{})

type BundleSpecDie struct {
	mutable bool
	r       v1alpha1.BundleSpec
}

// DieImmutable returns a new die for the current die's state that is either mutable (`false`) or immutable (`true`).
func (d *BundleSpecDie) DieImmutable(immutable bool) *BundleSpecDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

// DieFeed returns a new die with the provided resource.
func (d *BundleSpecDie) DieFeed(r v1alpha1.BundleSpec) *BundleSpecDie {
	if d.mutable {
		d.r = r
		return d
	}
	return &BundleSpecDie{
		mutable: d.mutable,
		r:       r,
	}
}

// DieFeedPtr returns a new die with the provided resource pointer. If the resource is nil, the empty value is used instead.
func (d *BundleSpecDie) DieFeedPtr(r *v1alpha1.BundleSpec) *BundleSpecDie {
	if r == nil {
		r = &v1alpha1.BundleSpec{}
	}
	return d.DieFeed(*r)
}

// DieFeedJSON returns a new die with the provided JSON. Panics on error.
func (d *BundleSpecDie) DieFeedJSON(j []byte) *BundleSpecDie {
	r := v1alpha1.BundleSpec{}
	if err := json.Unmarshal(j, &r); err != nil {
		panic(err)
	}
	return d.DieFeed(r)
}

// DieFeedYAML returns a new die with the provided YAML. Panics on error.
func (d *BundleSpecDie) DieFeedYAML(y []byte) *BundleSpecDie {
	r := v1alpha1.BundleSpec{}
	if err := yaml.Unmarshal(y, &r); err != nil {
		panic(err)
	}
	return d.DieFeed(r)
}

// DieFeedYAMLFile returns a new die loading YAML from a file path. Panics on error.
func (d *BundleSpecDie) DieFeedYAMLFile(name string) *BundleSpecDie {
	y, err := osx.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return d.DieFeedYAML(y)
}

// DieFeedRawExtension returns the resource managed by the die as an raw extension. Panics on error.
func (d *BundleSpecDie) DieFeedRawExtension(raw runtime.RawExtension) *BundleSpecDie {
	j, err := json.Marshal(raw)
	if err != nil {
		panic(err)
	}
	return d.DieFeedJSON(j)
}

// DieRelease returns the resource managed by the die.
func (d *BundleSpecDie) DieRelease() v1alpha1.BundleSpec {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

// DieReleasePtr returns a pointer to the resource managed by the die.
func (d *BundleSpecDie) DieReleasePtr() *v1alpha1.BundleSpec {
	r := d.DieRelease()
	return &r
}

// DieReleaseJSON returns the resource managed by the die as JSON. Panics on error.
func (d *BundleSpecDie) DieReleaseJSON() []byte {
	r := d.DieReleasePtr()
	j, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return j
}

// DieReleaseYAML returns the resource managed by the die as YAML. Panics on error.
func (d *BundleSpecDie) DieReleaseYAML() []byte {
	r := d.DieReleasePtr()
	y, err := yaml.Marshal(r)
	if err != nil {
		panic(err)
	}
	return y
}

// DieReleaseRawExtension returns the resource managed by the die as an raw extension. Panics on error.
func (d *BundleSpecDie) DieReleaseRawExtension() runtime.RawExtension {
	j := d.DieReleaseJSON()
	raw := runtime.RawExtension{}
	if err := json.Unmarshal(j, &raw); err != nil {
		panic(err)
	}
	return raw
}

// DieStamp returns a new die with the resource passed to the callback function. The resource is mutable.
func (d *BundleSpecDie) DieStamp(fn func(r *v1alpha1.BundleSpec)) *BundleSpecDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

// Experimental: DieStampAt uses a JSON path (http://goessner.net/articles/JsonPath/) expression to stamp portions of the resource. The callback is invoked with each JSON path match. Panics if the callback function does not accept a single argument of the same type as found on the resource at the target location.
//
// Future iterations will improve type coercion from the resource to the callback argument.
func (d *BundleSpecDie) DieStampAt(jp string, fn interface{}) *BundleSpecDie {
	return d.DieStamp(func(r *v1alpha1.BundleSpec) {
		cp := jsonpath.New("")
		if err := cp.Parse(fmtx.Sprintf("{%s}", jp)); err != nil {
			panic(err)
		}
		cr, err := cp.FindResults(r)
		if err != nil {
			// errors are expected if a path is not found
			return
		}
		for _, cv := range cr[0] {
			args := []reflectx.Value{cv}
			reflectx.ValueOf(fn).Call(args)
		}
	})
}

// DeepCopy returns a new die with equivalent state. Useful for snapshotting a mutable die.
func (d *BundleSpecDie) DeepCopy() *BundleSpecDie {
	r := *d.r.DeepCopy()
	return &BundleSpecDie{
		mutable: d.mutable,
		r:       r,
	}
}

func (d *BundleSpecDie) Include(v ...v1alpha1.Source) *BundleSpecDie {
	return d.DieStamp(func(r *v1alpha1.BundleSpec) {
		r.Include = v
	})
}

var BundleStatusBlank = (&BundleStatusDie{}).DieFeed(v1alpha1.BundleStatus{})

type BundleStatusDie struct {
	mutable bool
	r       v1alpha1.BundleStatus
}

// DieImmutable returns a new die for the current die's state that is either mutable (`false`) or immutable (`true`).
func (d *BundleStatusDie) DieImmutable(immutable bool) *BundleStatusDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

// DieFeed returns a new die with the provided resource.
func (d *BundleStatusDie) DieFeed(r v1alpha1.BundleStatus) *BundleStatusDie {
	if d.mutable {
		d.r = r
		return d
	}
	return &BundleStatusDie{
		mutable: d.mutable,
		r:       r,
	}
}

// DieFeedPtr returns a new die with the provided resource pointer. If the resource is nil, the empty value is used instead.
func (d *BundleStatusDie) DieFeedPtr(r *v1alpha1.BundleStatus) *BundleStatusDie {
	if r == nil {
		r = &v1alpha1.BundleStatus{}
	}
	return d.DieFeed(*r)
}

// DieFeedJSON returns a new die with the provided JSON. Panics on error.
func (d *BundleStatusDie) DieFeedJSON(j []byte) *BundleStatusDie {
	r := v1alpha1.BundleStatus{}
	if err := json.Unmarshal(j, &r); err != nil {
		panic(err)
	}
	return d.DieFeed(r)
}

// DieFeedYAML returns a new die with the provided YAML. Panics on error.
func (d *BundleStatusDie) DieFeedYAML(y []byte) *BundleStatusDie {
	r := v1alpha1.BundleStatus{}
	if err := yaml.Unmarshal(y, &r); err != nil {
		panic(err)
	}
	return d.DieFeed(r)
}

// DieFeedYAMLFile returns a new die loading YAML from a file path. Panics on error.
func (d *BundleStatusDie) DieFeedYAMLFile(name string) *BundleStatusDie {
	y, err := osx.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return d.DieFeedYAML(y)
}

// DieFeedRawExtension returns the resource managed by the die as an raw extension. Panics on error.
func (d *BundleStatusDie) DieFeedRawExtension(raw runtime.RawExtension) *BundleStatusDie {
	j, err := json.Marshal(raw)
	if err != nil {
		panic(err)
	}
	return d.DieFeedJSON(j)
}

// DieRelease returns the resource managed by the die.
func (d *BundleStatusDie) DieRelease() v1alpha1.BundleStatus {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

// DieReleasePtr returns a pointer to the resource managed by the die.
func (d *BundleStatusDie) DieReleasePtr() *v1alpha1.BundleStatus {
	r := d.DieRelease()
	return &r
}

// DieReleaseJSON returns the resource managed by the die as JSON. Panics on error.
func (d *BundleStatusDie) DieReleaseJSON() []byte {
	r := d.DieReleasePtr()
	j, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return j
}

// DieReleaseYAML returns the resource managed by the die as YAML. Panics on error.
func (d *BundleStatusDie) DieReleaseYAML() []byte {
	r := d.DieReleasePtr()
	y, err := yaml.Marshal(r)
	if err != nil {
		panic(err)
	}
	return y
}

// DieReleaseRawExtension returns the resource managed by the die as an raw extension. Panics on error.
func (d *BundleStatusDie) DieReleaseRawExtension() runtime.RawExtension {
	j := d.DieReleaseJSON()
	raw := runtime.RawExtension{}
	if err := json.Unmarshal(j, &raw); err != nil {
		panic(err)
	}
	return raw
}

// DieStamp returns a new die with the resource passed to the callback function. The resource is mutable.
func (d *BundleStatusDie) DieStamp(fn func(r *v1alpha1.BundleStatus)) *BundleStatusDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

// Experimental: DieStampAt uses a JSON path (http://goessner.net/articles/JsonPath/) expression to stamp portions of the resource. The callback is invoked with each JSON path match. Panics if the callback function does not accept a single argument of the same type as found on the resource at the target location.
//
// Future iterations will improve type coercion from the resource to the callback argument.
func (d *BundleStatusDie) DieStampAt(jp string, fn interface{}) *BundleStatusDie {
	return d.DieStamp(func(r *v1alpha1.BundleStatus) {
		cp := jsonpath.New("")
		if err := cp.Parse(fmtx.Sprintf("{%s}", jp)); err != nil {
			panic(err)
		}
		cr, err := cp.FindResults(r)
		if err != nil {
			// errors are expected if a path is not found
			return
		}
		for _, cv := range cr[0] {
			args := []reflectx.Value{cv}
			reflectx.ValueOf(fn).Call(args)
		}
	})
}

// DeepCopy returns a new die with equivalent state. Useful for snapshotting a mutable die.
func (d *BundleStatusDie) DeepCopy() *BundleStatusDie {
	r := *d.r.DeepCopy()
	return &BundleStatusDie{
		mutable: d.mutable,
		r:       r,
	}
}

func (d *BundleStatusDie) Status(v apis.Status) *BundleStatusDie {
	return d.DieStamp(func(r *v1alpha1.BundleStatus) {
		r.Status = v
	})
}
