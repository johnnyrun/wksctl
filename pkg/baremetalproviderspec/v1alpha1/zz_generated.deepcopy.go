// +build !ignore_autogenerated

// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIServer) DeepCopyInto(out *APIServer) {
	*out = *in
	if in.AdditionalSANs != nil {
		in, out := &in.AdditionalSANs, &out.AdditionalSANs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIServer.
func (in *APIServer) DeepCopy() *APIServer {
	if in == nil {
		return nil
	}
	out := new(APIServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Addon) DeepCopyInto(out *Addon) {
	*out = *in
	if in.Params != nil {
		in, out := &in.Params, &out.Params
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Deps != nil {
		in, out := &in.Deps, &out.Deps
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Addon.
func (in *Addon) DeepCopy() *Addon {
	if in == nil {
		return nil
	}
	out := new(Addon)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthenticationWebhook) DeepCopyInto(out *AuthenticationWebhook) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthenticationWebhook.
func (in *AuthenticationWebhook) DeepCopy() *AuthenticationWebhook {
	if in == nil {
		return nil
	}
	out := new(AuthenticationWebhook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthorizationWebhook) DeepCopyInto(out *AuthorizationWebhook) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthorizationWebhook.
func (in *AuthorizationWebhook) DeepCopy() *AuthorizationWebhook {
	if in == nil {
		return nil
	}
	out := new(AuthorizationWebhook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BareMetalClusterProviderSpec) DeepCopyInto(out *BareMetalClusterProviderSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Authentication != nil {
		in, out := &in.Authentication, &out.Authentication
		*out = new(AuthenticationWebhook)
		**out = **in
	}
	if in.Authorization != nil {
		in, out := &in.Authorization, &out.Authorization
		*out = new(AuthorizationWebhook)
		**out = **in
	}
	in.OS.DeepCopyInto(&out.OS)
	out.CRI = in.CRI
	in.APIServer.DeepCopyInto(&out.APIServer)
	if in.Addons != nil {
		in, out := &in.Addons, &out.Addons
		*out = make([]Addon, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BareMetalClusterProviderSpec.
func (in *BareMetalClusterProviderSpec) DeepCopy() *BareMetalClusterProviderSpec {
	if in == nil {
		return nil
	}
	out := new(BareMetalClusterProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BareMetalClusterProviderSpec) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BareMetalMachineProviderSpec) DeepCopyInto(out *BareMetalMachineProviderSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.Private = in.Private
	out.Public = in.Public
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BareMetalMachineProviderSpec.
func (in *BareMetalMachineProviderSpec) DeepCopy() *BareMetalMachineProviderSpec {
	if in == nil {
		return nil
	}
	out := new(BareMetalMachineProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BareMetalMachineProviderSpec) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerRuntime) DeepCopyInto(out *ContainerRuntime) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerRuntime.
func (in *ContainerRuntime) DeepCopy() *ContainerRuntime {
	if in == nil {
		return nil
	}
	out := new(ContainerRuntime)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndPoint) DeepCopyInto(out *EndPoint) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndPoint.
func (in *EndPoint) DeepCopy() *EndPoint {
	if in == nil {
		return nil
	}
	out := new(EndPoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSpec) DeepCopyInto(out *FileSpec) {
	*out = *in
	out.Source = in.Source
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSpec.
func (in *FileSpec) DeepCopy() *FileSpec {
	if in == nil {
		return nil
	}
	out := new(FileSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OSConfig) DeepCopyInto(out *OSConfig) {
	*out = *in
	if in.Files != nil {
		in, out := &in.Files, &out.Files
		*out = make([]FileSpec, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OSConfig.
func (in *OSConfig) DeepCopy() *OSConfig {
	if in == nil {
		return nil
	}
	out := new(OSConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceSpec) DeepCopyInto(out *SourceSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceSpec.
func (in *SourceSpec) DeepCopy() *SourceSpec {
	if in == nil {
		return nil
	}
	out := new(SourceSpec)
	in.DeepCopyInto(out)
	return out
}
