// +build !ignore_autogenerated

/*
Copyright 2017 The Gardener Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Addon) DeepCopyInto(out *Addon) {
	*out = *in
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
func (in *Addons) DeepCopyInto(out *Addons) {
	*out = *in
	in.Kube2IAM.DeepCopyInto(&out.Kube2IAM)
	out.Heapster = in.Heapster
	out.KubernetesDashboard = in.KubernetesDashboard
	out.ClusterAutoscaler = in.ClusterAutoscaler
	out.NginxIngress = in.NginxIngress
	out.Monocular = in.Monocular
	out.KubeLego = in.KubeLego
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Addons.
func (in *Addons) DeepCopy() *Addons {
	if in == nil {
		return nil
	}
	out := new(Addons)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Backup) DeepCopyInto(out *Backup) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Backup.
func (in *Backup) DeepCopy() *Backup {
	if in == nil {
		return nil
	}
	out := new(Backup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAutoscaler) DeepCopyInto(out *ClusterAutoscaler) {
	*out = *in
	out.Addon = in.Addon
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAutoscaler.
func (in *ClusterAutoscaler) DeepCopy() *ClusterAutoscaler {
	if in == nil {
		return nil
	}
	out := new(ClusterAutoscaler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNS) DeepCopyInto(out *DNS) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNS.
func (in *DNS) DeepCopy() *DNS {
	if in == nil {
		return nil
	}
	out := new(DNS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GardenOperator) DeepCopyInto(out *GardenOperator) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GardenOperator.
func (in *GardenOperator) DeepCopy() *GardenOperator {
	if in == nil {
		return nil
	}
	out := new(GardenOperator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Heapster) DeepCopyInto(out *Heapster) {
	*out = *in
	out.Addon = in.Addon
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Heapster.
func (in *Heapster) DeepCopy() *Heapster {
	if in == nil {
		return nil
	}
	out := new(Heapster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmTiller) DeepCopyInto(out *HelmTiller) {
	*out = *in
	out.Addon = in.Addon
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmTiller.
func (in *HelmTiller) DeepCopy() *HelmTiller {
	if in == nil {
		return nil
	}
	out := new(HelmTiller)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Infrastructure) DeepCopyInto(out *Infrastructure) {
	*out = *in
	if in.VNet != nil {
		in, out := &in.VNet, &out.VNet
		if *in == nil {
			*out = nil
		} else {
			*out = new(VNet)
			**out = **in
		}
	}
	if in.VPC != nil {
		in, out := &in.VPC, &out.VPC
		if *in == nil {
			*out = nil
		} else {
			*out = new(VPC)
			**out = **in
		}
	}
	if in.AdditionalDNS != nil {
		in, out := &in.AdditionalDNS, &out.AdditionalDNS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Infrastructure.
func (in *Infrastructure) DeepCopy() *Infrastructure {
	if in == nil {
		return nil
	}
	out := new(Infrastructure)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Kube2IAM) DeepCopyInto(out *Kube2IAM) {
	*out = *in
	out.Addon = in.Addon
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]Kube2IAMRole, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Kube2IAM.
func (in *Kube2IAM) DeepCopy() *Kube2IAM {
	if in == nil {
		return nil
	}
	out := new(Kube2IAM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Kube2IAMRole) DeepCopyInto(out *Kube2IAMRole) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Kube2IAMRole.
func (in *Kube2IAMRole) DeepCopy() *Kube2IAMRole {
	if in == nil {
		return nil
	}
	out := new(Kube2IAMRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeLego) DeepCopyInto(out *KubeLego) {
	*out = *in
	out.Addon = in.Addon
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeLego.
func (in *KubeLego) DeepCopy() *KubeLego {
	if in == nil {
		return nil
	}
	out := new(KubeLego)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesDashboard) DeepCopyInto(out *KubernetesDashboard) {
	*out = *in
	out.Addon = in.Addon
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesDashboard.
func (in *KubernetesDashboard) DeepCopy() *KubernetesDashboard {
	if in == nil {
		return nil
	}
	out := new(KubernetesDashboard)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LastOperation) DeepCopyInto(out *LastOperation) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LastOperation.
func (in *LastOperation) DeepCopy() *LastOperation {
	if in == nil {
		return nil
	}
	out := new(LastOperation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Monocular) DeepCopyInto(out *Monocular) {
	*out = *in
	out.Addon = in.Addon
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Monocular.
func (in *Monocular) DeepCopy() *Monocular {
	if in == nil {
		return nil
	}
	out := new(Monocular)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Networks) DeepCopyInto(out *Networks) {
	*out = *in
	if in.Workers != nil {
		in, out := &in.Workers, &out.Workers
		*out = make([]CIDR, len(*in))
		copy(*out, *in)
	}
	if in.Public != nil {
		in, out := &in.Public, &out.Public
		*out = make([]CIDR, len(*in))
		copy(*out, *in)
	}
	if in.Internal != nil {
		in, out := &in.Internal, &out.Internal
		*out = make([]CIDR, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Networks.
func (in *Networks) DeepCopy() *Networks {
	if in == nil {
		return nil
	}
	out := new(Networks)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NginxIngress) DeepCopyInto(out *NginxIngress) {
	*out = *in
	out.Addon = in.Addon
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NginxIngress.
func (in *NginxIngress) DeepCopy() *NginxIngress {
	if in == nil {
		return nil
	}
	out := new(NginxIngress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Shoot) DeepCopyInto(out *Shoot) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.TypeMeta = in.TypeMeta
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Shoot.
func (in *Shoot) DeepCopy() *Shoot {
	if in == nil {
		return nil
	}
	out := new(Shoot)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Shoot) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShootCondition) DeepCopyInto(out *ShootCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShootCondition.
func (in *ShootCondition) DeepCopy() *ShootCondition {
	if in == nil {
		return nil
	}
	out := new(ShootCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShootList) DeepCopyInto(out *ShootList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Shoot, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShootList.
func (in *ShootList) DeepCopy() *ShootList {
	if in == nil {
		return nil
	}
	out := new(ShootList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ShootList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShootSpec) DeepCopyInto(out *ShootSpec) {
	*out = *in
	if in.Addons != nil {
		in, out := &in.Addons, &out.Addons
		if *in == nil {
			*out = nil
		} else {
			*out = new(Addons)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Backup != nil {
		in, out := &in.Backup, &out.Backup
		if *in == nil {
			*out = nil
		} else {
			*out = new(Backup)
			**out = **in
		}
	}
	if in.DNS != nil {
		in, out := &in.DNS, &out.DNS
		if *in == nil {
			*out = nil
		} else {
			*out = new(DNS)
			**out = **in
		}
	}
	if in.Infrastructure != nil {
		in, out := &in.Infrastructure, &out.Infrastructure
		if *in == nil {
			*out = nil
		} else {
			*out = new(Infrastructure)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Networks != nil {
		in, out := &in.Networks, &out.Networks
		if *in == nil {
			*out = nil
		} else {
			*out = new(Networks)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Workers != nil {
		in, out := &in.Workers, &out.Workers
		*out = make([]Worker, len(*in))
		copy(*out, *in)
	}
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
		*out = make([]Zone, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShootSpec.
func (in *ShootSpec) DeepCopy() *ShootSpec {
	if in == nil {
		return nil
	}
	out := new(ShootSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShootStatus) DeepCopyInto(out *ShootStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ShootCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.GardenOperator = in.GardenOperator
	if in.LastOperation != nil {
		in, out := &in.LastOperation, &out.LastOperation
		if *in == nil {
			*out = nil
		} else {
			*out = new(LastOperation)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.OperationStartTime != nil {
		in, out := &in.OperationStartTime, &out.OperationStartTime
		if *in == nil {
			*out = nil
		} else {
			*out = new(meta_v1.Time)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShootStatus.
func (in *ShootStatus) DeepCopy() *ShootStatus {
	if in == nil {
		return nil
	}
	out := new(ShootStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VNet) DeepCopyInto(out *VNet) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VNet.
func (in *VNet) DeepCopy() *VNet {
	if in == nil {
		return nil
	}
	out := new(VNet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPC) DeepCopyInto(out *VPC) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPC.
func (in *VPC) DeepCopy() *VPC {
	if in == nil {
		return nil
	}
	out := new(VPC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Worker) DeepCopyInto(out *Worker) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Worker.
func (in *Worker) DeepCopy() *Worker {
	if in == nil {
		return nil
	}
	out := new(Worker)
	in.DeepCopyInto(out)
	return out
}
