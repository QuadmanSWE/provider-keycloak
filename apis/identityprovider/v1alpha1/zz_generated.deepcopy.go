//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityProviderMapper) DeepCopyInto(out *IdentityProviderMapper) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityProviderMapper.
func (in *IdentityProviderMapper) DeepCopy() *IdentityProviderMapper {
	if in == nil {
		return nil
	}
	out := new(IdentityProviderMapper)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IdentityProviderMapper) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityProviderMapperInitParameters) DeepCopyInto(out *IdentityProviderMapperInitParameters) {
	*out = *in
	if in.ExtraConfig != nil {
		in, out := &in.ExtraConfig, &out.ExtraConfig
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.IdentityProviderAlias != nil {
		in, out := &in.IdentityProviderAlias, &out.IdentityProviderAlias
		*out = new(string)
		**out = **in
	}
	if in.IdentityProviderMapper != nil {
		in, out := &in.IdentityProviderMapper, &out.IdentityProviderMapper
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Realm != nil {
		in, out := &in.Realm, &out.Realm
		*out = new(string)
		**out = **in
	}
	if in.RealmRef != nil {
		in, out := &in.RealmRef, &out.RealmRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.RealmSelector != nil {
		in, out := &in.RealmSelector, &out.RealmSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityProviderMapperInitParameters.
func (in *IdentityProviderMapperInitParameters) DeepCopy() *IdentityProviderMapperInitParameters {
	if in == nil {
		return nil
	}
	out := new(IdentityProviderMapperInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityProviderMapperList) DeepCopyInto(out *IdentityProviderMapperList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IdentityProviderMapper, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityProviderMapperList.
func (in *IdentityProviderMapperList) DeepCopy() *IdentityProviderMapperList {
	if in == nil {
		return nil
	}
	out := new(IdentityProviderMapperList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IdentityProviderMapperList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityProviderMapperObservation) DeepCopyInto(out *IdentityProviderMapperObservation) {
	*out = *in
	if in.ExtraConfig != nil {
		in, out := &in.ExtraConfig, &out.ExtraConfig
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IdentityProviderAlias != nil {
		in, out := &in.IdentityProviderAlias, &out.IdentityProviderAlias
		*out = new(string)
		**out = **in
	}
	if in.IdentityProviderMapper != nil {
		in, out := &in.IdentityProviderMapper, &out.IdentityProviderMapper
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Realm != nil {
		in, out := &in.Realm, &out.Realm
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityProviderMapperObservation.
func (in *IdentityProviderMapperObservation) DeepCopy() *IdentityProviderMapperObservation {
	if in == nil {
		return nil
	}
	out := new(IdentityProviderMapperObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityProviderMapperParameters) DeepCopyInto(out *IdentityProviderMapperParameters) {
	*out = *in
	if in.ExtraConfig != nil {
		in, out := &in.ExtraConfig, &out.ExtraConfig
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.IdentityProviderAlias != nil {
		in, out := &in.IdentityProviderAlias, &out.IdentityProviderAlias
		*out = new(string)
		**out = **in
	}
	if in.IdentityProviderMapper != nil {
		in, out := &in.IdentityProviderMapper, &out.IdentityProviderMapper
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Realm != nil {
		in, out := &in.Realm, &out.Realm
		*out = new(string)
		**out = **in
	}
	if in.RealmRef != nil {
		in, out := &in.RealmRef, &out.RealmRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.RealmSelector != nil {
		in, out := &in.RealmSelector, &out.RealmSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityProviderMapperParameters.
func (in *IdentityProviderMapperParameters) DeepCopy() *IdentityProviderMapperParameters {
	if in == nil {
		return nil
	}
	out := new(IdentityProviderMapperParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityProviderMapperSpec) DeepCopyInto(out *IdentityProviderMapperSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityProviderMapperSpec.
func (in *IdentityProviderMapperSpec) DeepCopy() *IdentityProviderMapperSpec {
	if in == nil {
		return nil
	}
	out := new(IdentityProviderMapperSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityProviderMapperStatus) DeepCopyInto(out *IdentityProviderMapperStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityProviderMapperStatus.
func (in *IdentityProviderMapperStatus) DeepCopy() *IdentityProviderMapperStatus {
	if in == nil {
		return nil
	}
	out := new(IdentityProviderMapperStatus)
	in.DeepCopyInto(out)
	return out
}
