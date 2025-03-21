//go:build !ignore_autogenerated

// File generated by the kubebuilder framework:
// https://github.com/kubernetes-sigs/kubebuilder

/*
Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneAwareUpdate) DeepCopyInto(out *ZoneAwareUpdate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneAwareUpdate.
func (in *ZoneAwareUpdate) DeepCopy() *ZoneAwareUpdate {
	if in == nil {
		return nil
	}
	out := new(ZoneAwareUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ZoneAwareUpdate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneAwareUpdateList) DeepCopyInto(out *ZoneAwareUpdateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ZoneAwareUpdate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneAwareUpdateList.
func (in *ZoneAwareUpdateList) DeepCopy() *ZoneAwareUpdateList {
	if in == nil {
		return nil
	}
	out := new(ZoneAwareUpdateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ZoneAwareUpdateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneAwareUpdateSpec) DeepCopyInto(out *ZoneAwareUpdateSpec) {
	*out = *in
	if in.MaxUnavailable != nil {
		in, out := &in.MaxUnavailable, &out.MaxUnavailable
		*out = new(intstr.IntOrString)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneAwareUpdateSpec.
func (in *ZoneAwareUpdateSpec) DeepCopy() *ZoneAwareUpdateSpec {
	if in == nil {
		return nil
	}
	out := new(ZoneAwareUpdateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneAwareUpdateStatus) DeepCopyInto(out *ZoneAwareUpdateStatus) {
	*out = *in
	if in.OldReplicas != nil {
		in, out := &in.OldReplicas, &out.OldReplicas
		*out = make(map[string]int32, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneAwareUpdateStatus.
func (in *ZoneAwareUpdateStatus) DeepCopy() *ZoneAwareUpdateStatus {
	if in == nil {
		return nil
	}
	out := new(ZoneAwareUpdateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneDisruptionBudget) DeepCopyInto(out *ZoneDisruptionBudget) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneDisruptionBudget.
func (in *ZoneDisruptionBudget) DeepCopy() *ZoneDisruptionBudget {
	if in == nil {
		return nil
	}
	out := new(ZoneDisruptionBudget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ZoneDisruptionBudget) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneDisruptionBudgetList) DeepCopyInto(out *ZoneDisruptionBudgetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ZoneDisruptionBudget, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneDisruptionBudgetList.
func (in *ZoneDisruptionBudgetList) DeepCopy() *ZoneDisruptionBudgetList {
	if in == nil {
		return nil
	}
	out := new(ZoneDisruptionBudgetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ZoneDisruptionBudgetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneDisruptionBudgetSpec) DeepCopyInto(out *ZoneDisruptionBudgetSpec) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.MaxUnavailable != nil {
		in, out := &in.MaxUnavailable, &out.MaxUnavailable
		*out = new(intstr.IntOrString)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneDisruptionBudgetSpec.
func (in *ZoneDisruptionBudgetSpec) DeepCopy() *ZoneDisruptionBudgetSpec {
	if in == nil {
		return nil
	}
	out := new(ZoneDisruptionBudgetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneDisruptionBudgetStatus) DeepCopyInto(out *ZoneDisruptionBudgetStatus) {
	*out = *in
	if in.DisruptedPods != nil {
		in, out := &in.DisruptedPods, &out.DisruptedPods
		*out = make(map[string]metav1.Time, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.DisruptionsAllowed != nil {
		in, out := &in.DisruptionsAllowed, &out.DisruptionsAllowed
		*out = make(map[string]int32, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.CurrentHealthy != nil {
		in, out := &in.CurrentHealthy, &out.CurrentHealthy
		*out = make(map[string]int32, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.CurrentUnhealthy != nil {
		in, out := &in.CurrentUnhealthy, &out.CurrentUnhealthy
		*out = make(map[string]int32, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DesiredHealthy != nil {
		in, out := &in.DesiredHealthy, &out.DesiredHealthy
		*out = make(map[string]int32, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ExpectedPods != nil {
		in, out := &in.ExpectedPods, &out.ExpectedPods
		*out = make(map[string]int32, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneDisruptionBudgetStatus.
func (in *ZoneDisruptionBudgetStatus) DeepCopy() *ZoneDisruptionBudgetStatus {
	if in == nil {
		return nil
	}
	out := new(ZoneDisruptionBudgetStatus)
	in.DeepCopyInto(out)
	return out
}
