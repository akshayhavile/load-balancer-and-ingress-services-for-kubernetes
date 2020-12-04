// +build !ignore_autogenerated

/*


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
func (in *AKOConfig) DeepCopyInto(out *AKOConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AKOConfig.
func (in *AKOConfig) DeepCopy() *AKOConfig {
	if in == nil {
		return nil
	}
	out := new(AKOConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AKOConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AKOConfigList) DeepCopyInto(out *AKOConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AKOConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AKOConfigList.
func (in *AKOConfigList) DeepCopy() *AKOConfigList {
	if in == nil {
		return nil
	}
	out := new(AKOConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AKOConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AKOConfigSpec) DeepCopyInto(out *AKOConfigSpec) {
	*out = *in
	out.AKOSettings = in.AKOSettings
	in.NetworkSettings.DeepCopyInto(&out.NetworkSettings)
	out.L7Settings = in.L7Settings
	out.L4Settings = in.L4Settings
	out.ControllerSettings = in.ControllerSettings
	out.NodePortSelector = in.NodePortSelector
	out.Resources = in.Resources
	out.Rbac = in.Rbac
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AKOConfigSpec.
func (in *AKOConfigSpec) DeepCopy() *AKOConfigSpec {
	if in == nil {
		return nil
	}
	out := new(AKOConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AKOConfigStatus) DeepCopyInto(out *AKOConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AKOConfigStatus.
func (in *AKOConfigStatus) DeepCopy() *AKOConfigStatus {
	if in == nil {
		return nil
	}
	out := new(AKOConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AKOSettings) DeepCopyInto(out *AKOSettings) {
	*out = *in
	out.NSSelector = in.NSSelector
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AKOSettings.
func (in *AKOSettings) DeepCopy() *AKOSettings {
	if in == nil {
		return nil
	}
	out := new(AKOSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerSettings) DeepCopyInto(out *ControllerSettings) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerSettings.
func (in *ControllerSettings) DeepCopy() *ControllerSettings {
	if in == nil {
		return nil
	}
	out := new(ControllerSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *L4Settings) DeepCopyInto(out *L4Settings) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L4Settings.
func (in *L4Settings) DeepCopy() *L4Settings {
	if in == nil {
		return nil
	}
	out := new(L4Settings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *L7Settings) DeepCopyInto(out *L7Settings) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L7Settings.
func (in *L7Settings) DeepCopy() *L7Settings {
	if in == nil {
		return nil
	}
	out := new(L7Settings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceSelector) DeepCopyInto(out *NamespaceSelector) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceSelector.
func (in *NamespaceSelector) DeepCopy() *NamespaceSelector {
	if in == nil {
		return nil
	}
	out := new(NamespaceSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkSettings) DeepCopyInto(out *NetworkSettings) {
	*out = *in
	if in.NodeNetworkList != nil {
		in, out := &in.NodeNetworkList, &out.NodeNetworkList
		*out = make([]NodeNetwork, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkSettings.
func (in *NetworkSettings) DeepCopy() *NetworkSettings {
	if in == nil {
		return nil
	}
	out := new(NetworkSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeNetwork) DeepCopyInto(out *NodeNetwork) {
	*out = *in
	if in.Cidrs != nil {
		in, out := &in.Cidrs, &out.Cidrs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeNetwork.
func (in *NodeNetwork) DeepCopy() *NodeNetwork {
	if in == nil {
		return nil
	}
	out := new(NodeNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePortSelector) DeepCopyInto(out *NodePortSelector) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePortSelector.
func (in *NodePortSelector) DeepCopy() *NodePortSelector {
	if in == nil {
		return nil
	}
	out := new(NodePortSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rbac) DeepCopyInto(out *Rbac) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rbac.
func (in *Rbac) DeepCopy() *Rbac {
	if in == nil {
		return nil
	}
	out := new(Rbac)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceLimits) DeepCopyInto(out *ResourceLimits) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceLimits.
func (in *ResourceLimits) DeepCopy() *ResourceLimits {
	if in == nil {
		return nil
	}
	out := new(ResourceLimits)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRequests) DeepCopyInto(out *ResourceRequests) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRequests.
func (in *ResourceRequests) DeepCopy() *ResourceRequests {
	if in == nil {
		return nil
	}
	out := new(ResourceRequests)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resources) DeepCopyInto(out *Resources) {
	*out = *in
	out.Limits = in.Limits
	out.Requests = in.Requests
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resources.
func (in *Resources) DeepCopy() *Resources {
	if in == nil {
		return nil
	}
	out := new(Resources)
	in.DeepCopyInto(out)
	return out
}
