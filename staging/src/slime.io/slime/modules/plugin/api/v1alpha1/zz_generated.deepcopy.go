// +build !ignore_autogenerated

/*
Copyright The Slime Authors.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvoyPlugin) DeepCopyInto(out *EnvoyPlugin) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyPlugin.
func (in *EnvoyPlugin) DeepCopy() *EnvoyPlugin {
	if in == nil {
		return nil
	}
	out := new(EnvoyPlugin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnvoyPlugin) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvoyPluginList) DeepCopyInto(out *EnvoyPluginList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EnvoyPlugin, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyPluginList.
func (in *EnvoyPluginList) DeepCopy() *EnvoyPluginList {
	if in == nil {
		return nil
	}
	out := new(EnvoyPluginList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnvoyPluginList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PluginManager) DeepCopyInto(out *PluginManager) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PluginManager.
func (in *PluginManager) DeepCopy() *PluginManager {
	if in == nil {
		return nil
	}
	out := new(PluginManager)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PluginManager) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PluginManagerList) DeepCopyInto(out *PluginManagerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EnvoyPlugin, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PluginManagerList.
func (in *PluginManagerList) DeepCopy() *PluginManagerList {
	if in == nil {
		return nil
	}
	out := new(PluginManagerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PluginManagerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Plugin_Inline) DeepCopyInto(out *Plugin_Inline) {
	*out = *in
	if in.Inline != nil {
		in, out := &in.Inline, &out.Inline
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Plugin_Inline.
func (in *Plugin_Inline) DeepCopy() *Plugin_Inline {
	if in == nil {
		return nil
	}
	out := new(Plugin_Inline)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Plugin_Rider) DeepCopyInto(out *Plugin_Rider) {
	*out = *in
	if in.Rider != nil {
		in, out := &in.Rider, &out.Rider
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Plugin_Rider.
func (in *Plugin_Rider) DeepCopy() *Plugin_Rider {
	if in == nil {
		return nil
	}
	out := new(Plugin_Rider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Plugin_Wasm) DeepCopyInto(out *Plugin_Wasm) {
	*out = *in
	if in.Wasm != nil {
		in, out := &in.Wasm, &out.Wasm
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Plugin_Wasm.
func (in *Plugin_Wasm) DeepCopy() *Plugin_Wasm {
	if in == nil {
		return nil
	}
	out := new(Plugin_Wasm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rider_ImagePullSecretContent) DeepCopyInto(out *Rider_ImagePullSecretContent) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rider_ImagePullSecretContent.
func (in *Rider_ImagePullSecretContent) DeepCopy() *Rider_ImagePullSecretContent {
	if in == nil {
		return nil
	}
	out := new(Rider_ImagePullSecretContent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rider_ImagePullSecretName) DeepCopyInto(out *Rider_ImagePullSecretName) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rider_ImagePullSecretName.
func (in *Rider_ImagePullSecretName) DeepCopy() *Rider_ImagePullSecretName {
	if in == nil {
		return nil
	}
	out := new(Rider_ImagePullSecretName)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Wasm_ImagePullSecretContent) DeepCopyInto(out *Wasm_ImagePullSecretContent) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Wasm_ImagePullSecretContent.
func (in *Wasm_ImagePullSecretContent) DeepCopy() *Wasm_ImagePullSecretContent {
	if in == nil {
		return nil
	}
	out := new(Wasm_ImagePullSecretContent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Wasm_ImagePullSecretName) DeepCopyInto(out *Wasm_ImagePullSecretName) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Wasm_ImagePullSecretName.
func (in *Wasm_ImagePullSecretName) DeepCopy() *Wasm_ImagePullSecretName {
	if in == nil {
		return nil
	}
	out := new(Wasm_ImagePullSecretName)
	in.DeepCopyInto(out)
	return out
}
