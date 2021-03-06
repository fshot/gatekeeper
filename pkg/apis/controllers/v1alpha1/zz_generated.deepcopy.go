// +build !ignore_autogenerated

/*
Copyright 2018 Replicated.

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
// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenPolicyAgent) DeepCopyInto(out *OpenPolicyAgent) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenPolicyAgent.
func (in *OpenPolicyAgent) DeepCopy() *OpenPolicyAgent {
	if in == nil {
		return nil
	}
	out := new(OpenPolicyAgent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenPolicyAgent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenPolicyAgentEnabledFailureModes) DeepCopyInto(out *OpenPolicyAgentEnabledFailureModes) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenPolicyAgentEnabledFailureModes.
func (in *OpenPolicyAgentEnabledFailureModes) DeepCopy() *OpenPolicyAgentEnabledFailureModes {
	if in == nil {
		return nil
	}
	out := new(OpenPolicyAgentEnabledFailureModes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenPolicyAgentList) DeepCopyInto(out *OpenPolicyAgentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpenPolicyAgent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenPolicyAgentList.
func (in *OpenPolicyAgentList) DeepCopy() *OpenPolicyAgentList {
	if in == nil {
		return nil
	}
	out := new(OpenPolicyAgentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenPolicyAgentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenPolicyAgentSpec) DeepCopyInto(out *OpenPolicyAgentSpec) {
	*out = *in
	if in.EnabledFailureModes != nil {
		in, out := &in.EnabledFailureModes, &out.EnabledFailureModes
		*out = new(OpenPolicyAgentEnabledFailureModes)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenPolicyAgentSpec.
func (in *OpenPolicyAgentSpec) DeepCopy() *OpenPolicyAgentSpec {
	if in == nil {
		return nil
	}
	out := new(OpenPolicyAgentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenPolicyAgentStatus) DeepCopyInto(out *OpenPolicyAgentStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenPolicyAgentStatus.
func (in *OpenPolicyAgentStatus) DeepCopy() *OpenPolicyAgentStatus {
	if in == nil {
		return nil
	}
	out := new(OpenPolicyAgentStatus)
	in.DeepCopyInto(out)
	return out
}
