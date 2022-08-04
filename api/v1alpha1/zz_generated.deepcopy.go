//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 The RamenDR authors.

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
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DRCluster) DeepCopyInto(out *DRCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DRCluster.
func (in *DRCluster) DeepCopy() *DRCluster {
	if in == nil {
		return nil
	}
	out := new(DRCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DRCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DRClusterList) DeepCopyInto(out *DRClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DRCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DRClusterList.
func (in *DRClusterList) DeepCopy() *DRClusterList {
	if in == nil {
		return nil
	}
	out := new(DRClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DRClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DRClusterSpec) DeepCopyInto(out *DRClusterSpec) {
	*out = *in
	if in.CIDRs != nil {
		in, out := &in.CIDRs, &out.CIDRs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DRClusterSpec.
func (in *DRClusterSpec) DeepCopy() *DRClusterSpec {
	if in == nil {
		return nil
	}
	out := new(DRClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DRClusterStatus) DeepCopyInto(out *DRClusterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DRClusterStatus.
func (in *DRClusterStatus) DeepCopy() *DRClusterStatus {
	if in == nil {
		return nil
	}
	out := new(DRClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DRPlacementControl) DeepCopyInto(out *DRPlacementControl) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DRPlacementControl.
func (in *DRPlacementControl) DeepCopy() *DRPlacementControl {
	if in == nil {
		return nil
	}
	out := new(DRPlacementControl)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DRPlacementControl) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DRPlacementControlList) DeepCopyInto(out *DRPlacementControlList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DRPlacementControl, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DRPlacementControlList.
func (in *DRPlacementControlList) DeepCopy() *DRPlacementControlList {
	if in == nil {
		return nil
	}
	out := new(DRPlacementControlList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DRPlacementControlList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DRPlacementControlSpec) DeepCopyInto(out *DRPlacementControlSpec) {
	*out = *in
	out.PlacementRef = in.PlacementRef
	out.DRPolicyRef = in.DRPolicyRef
	in.PVCSelector.DeepCopyInto(&out.PVCSelector)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DRPlacementControlSpec.
func (in *DRPlacementControlSpec) DeepCopy() *DRPlacementControlSpec {
	if in == nil {
		return nil
	}
	out := new(DRPlacementControlSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DRPlacementControlStatus) DeepCopyInto(out *DRPlacementControlStatus) {
	*out = *in
	if in.ActionStartTime != nil {
		in, out := &in.ActionStartTime, &out.ActionStartTime
		*out = (*in).DeepCopy()
	}
	if in.ActionDuration != nil {
		in, out := &in.ActionDuration, &out.ActionDuration
		*out = new(v1.Duration)
		**out = **in
	}
	out.PreferredDecision = in.PreferredDecision
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.ResourceConditions.DeepCopyInto(&out.ResourceConditions)
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DRPlacementControlStatus.
func (in *DRPlacementControlStatus) DeepCopy() *DRPlacementControlStatus {
	if in == nil {
		return nil
	}
	out := new(DRPlacementControlStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DRPolicy) DeepCopyInto(out *DRPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DRPolicy.
func (in *DRPolicy) DeepCopy() *DRPolicy {
	if in == nil {
		return nil
	}
	out := new(DRPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DRPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DRPolicyList) DeepCopyInto(out *DRPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DRPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DRPolicyList.
func (in *DRPolicyList) DeepCopy() *DRPolicyList {
	if in == nil {
		return nil
	}
	out := new(DRPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DRPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DRPolicySpec) DeepCopyInto(out *DRPolicySpec) {
	*out = *in
	in.ReplicationClassSelector.DeepCopyInto(&out.ReplicationClassSelector)
	in.VolumeSnapshotClassSelector.DeepCopyInto(&out.VolumeSnapshotClassSelector)
	if in.DRClusters != nil {
		in, out := &in.DRClusters, &out.DRClusters
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DRPolicySpec.
func (in *DRPolicySpec) DeepCopy() *DRPolicySpec {
	if in == nil {
		return nil
	}
	out := new(DRPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DRPolicyStatus) DeepCopyInto(out *DRPolicyStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DRPolicyStatus.
func (in *DRPolicyStatus) DeepCopy() *DRPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(DRPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeObjectProtection) DeepCopyInto(out *KubeObjectProtection) {
	*out = *in
	if in.Disabled != nil {
		in, out := &in.Disabled, &out.Disabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeObjectProtection.
func (in *KubeObjectProtection) DeepCopy() *KubeObjectProtection {
	if in == nil {
		return nil
	}
	out := new(KubeObjectProtection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeObjectProtectionSpec) DeepCopyInto(out *KubeObjectProtectionSpec) {
	*out = *in
	if in.CaptureInterval != nil {
		in, out := &in.CaptureInterval, &out.CaptureInterval
		*out = new(v1.Duration)
		**out = **in
	}
	if in.CaptureOrder != nil {
		in, out := &in.CaptureOrder, &out.CaptureOrder
		*out = make([]KubeObjectsCaptureGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RecoverOrder != nil {
		in, out := &in.RecoverOrder, &out.RecoverOrder
		*out = make([]KubeObjectsRecoverGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeObjectProtectionSpec.
func (in *KubeObjectProtectionSpec) DeepCopy() *KubeObjectProtectionSpec {
	if in == nil {
		return nil
	}
	out := new(KubeObjectProtectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeObjectProtectionStatus) DeepCopyInto(out *KubeObjectProtectionStatus) {
	*out = *in
	if in.CaptureToRecoverFrom != nil {
		in, out := &in.CaptureToRecoverFrom, &out.CaptureToRecoverFrom
		*out = new(KubeObjectsCaptureIdentifier)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeObjectProtectionStatus.
func (in *KubeObjectProtectionStatus) DeepCopy() *KubeObjectProtectionStatus {
	if in == nil {
		return nil
	}
	out := new(KubeObjectProtectionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeObjectsCaptureGroup) DeepCopyInto(out *KubeObjectsCaptureGroup) {
	*out = *in
	in.KubeObjectsSpec.DeepCopyInto(&out.KubeObjectsSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeObjectsCaptureGroup.
func (in *KubeObjectsCaptureGroup) DeepCopy() *KubeObjectsCaptureGroup {
	if in == nil {
		return nil
	}
	out := new(KubeObjectsCaptureGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeObjectsCaptureIdentifier) DeepCopyInto(out *KubeObjectsCaptureIdentifier) {
	*out = *in
	in.StartTime.DeepCopyInto(&out.StartTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeObjectsCaptureIdentifier.
func (in *KubeObjectsCaptureIdentifier) DeepCopy() *KubeObjectsCaptureIdentifier {
	if in == nil {
		return nil
	}
	out := new(KubeObjectsCaptureIdentifier)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeObjectsRecoverGroup) DeepCopyInto(out *KubeObjectsRecoverGroup) {
	*out = *in
	in.KubeObjectsSpec.DeepCopyInto(&out.KubeObjectsSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeObjectsRecoverGroup.
func (in *KubeObjectsRecoverGroup) DeepCopy() *KubeObjectsRecoverGroup {
	if in == nil {
		return nil
	}
	out := new(KubeObjectsRecoverGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeObjectsSpec) DeepCopyInto(out *KubeObjectsSpec) {
	*out = *in
	if in.IncludedResources != nil {
		in, out := &in.IncludedResources, &out.IncludedResources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExcludedResources != nil {
		in, out := &in.ExcludedResources, &out.ExcludedResources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.IncludeClusterResources != nil {
		in, out := &in.IncludeClusterResources, &out.IncludeClusterResources
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeObjectsSpec.
func (in *KubeObjectsSpec) DeepCopy() *KubeObjectsSpec {
	if in == nil {
		return nil
	}
	out := new(KubeObjectsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtectedPVC) DeepCopyInto(out *ProtectedPVC) {
	*out = *in
	if in.StorageClassName != nil {
		in, out := &in.StorageClassName, &out.StorageClassName
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AccessModes != nil {
		in, out := &in.AccessModes, &out.AccessModes
		*out = make([]corev1.PersistentVolumeAccessMode, len(*in))
		copy(*out, *in)
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtectedPVC.
func (in *ProtectedPVC) DeepCopy() *ProtectedPVC {
	if in == nil {
		return nil
	}
	out := new(ProtectedPVC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RamenConfig) DeepCopyInto(out *RamenConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ControllerManagerConfigurationSpec.DeepCopyInto(&out.ControllerManagerConfigurationSpec)
	if in.S3StoreProfiles != nil {
		in, out := &in.S3StoreProfiles, &out.S3StoreProfiles
		*out = make([]S3StoreProfile, len(*in))
		copy(*out, *in)
	}
	out.DrClusterOperator = in.DrClusterOperator
	out.VolSync = in.VolSync
	if in.KubeObjectProtection != nil {
		in, out := &in.KubeObjectProtection, &out.KubeObjectProtection
		*out = new(KubeObjectProtection)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RamenConfig.
func (in *RamenConfig) DeepCopy() *RamenConfig {
	if in == nil {
		return nil
	}
	out := new(RamenConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RamenConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3BucketView) DeepCopyInto(out *S3BucketView) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(S3BucketViewStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3BucketView.
func (in *S3BucketView) DeepCopy() *S3BucketView {
	if in == nil {
		return nil
	}
	out := new(S3BucketView)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *S3BucketView) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3BucketViewList) DeepCopyInto(out *S3BucketViewList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]S3BucketView, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3BucketViewList.
func (in *S3BucketViewList) DeepCopy() *S3BucketViewList {
	if in == nil {
		return nil
	}
	out := new(S3BucketViewList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *S3BucketViewList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3BucketViewSpec) DeepCopyInto(out *S3BucketViewSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3BucketViewSpec.
func (in *S3BucketViewSpec) DeepCopy() *S3BucketViewSpec {
	if in == nil {
		return nil
	}
	out := new(S3BucketViewSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3BucketViewStatus) DeepCopyInto(out *S3BucketViewStatus) {
	*out = *in
	in.SampleTime.DeepCopyInto(&out.SampleTime)
	if in.VolumeReplicationGroups != nil {
		in, out := &in.VolumeReplicationGroups, &out.VolumeReplicationGroups
		*out = make([]VolumeReplicationGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3BucketViewStatus.
func (in *S3BucketViewStatus) DeepCopy() *S3BucketViewStatus {
	if in == nil {
		return nil
	}
	out := new(S3BucketViewStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3StoreProfile) DeepCopyInto(out *S3StoreProfile) {
	*out = *in
	out.S3SecretRef = in.S3SecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3StoreProfile.
func (in *S3StoreProfile) DeepCopy() *S3StoreProfile {
	if in == nil {
		return nil
	}
	out := new(S3StoreProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VRGAsyncSpec) DeepCopyInto(out *VRGAsyncSpec) {
	*out = *in
	in.ReplicationClassSelector.DeepCopyInto(&out.ReplicationClassSelector)
	in.VolumeSnapshotClassSelector.DeepCopyInto(&out.VolumeSnapshotClassSelector)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VRGAsyncSpec.
func (in *VRGAsyncSpec) DeepCopy() *VRGAsyncSpec {
	if in == nil {
		return nil
	}
	out := new(VRGAsyncSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VRGConditions) DeepCopyInto(out *VRGConditions) {
	*out = *in
	in.ResourceMeta.DeepCopyInto(&out.ResourceMeta)
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VRGConditions.
func (in *VRGConditions) DeepCopy() *VRGConditions {
	if in == nil {
		return nil
	}
	out := new(VRGConditions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VRGResourceMeta) DeepCopyInto(out *VRGResourceMeta) {
	*out = *in
	if in.ProtectedPVCs != nil {
		in, out := &in.ProtectedPVCs, &out.ProtectedPVCs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VRGResourceMeta.
func (in *VRGResourceMeta) DeepCopy() *VRGResourceMeta {
	if in == nil {
		return nil
	}
	out := new(VRGResourceMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VRGSyncSpec) DeepCopyInto(out *VRGSyncSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VRGSyncSpec.
func (in *VRGSyncSpec) DeepCopy() *VRGSyncSpec {
	if in == nil {
		return nil
	}
	out := new(VRGSyncSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolSyncReplicationDestinationSpec) DeepCopyInto(out *VolSyncReplicationDestinationSpec) {
	*out = *in
	in.ProtectedPVC.DeepCopyInto(&out.ProtectedPVC)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolSyncReplicationDestinationSpec.
func (in *VolSyncReplicationDestinationSpec) DeepCopy() *VolSyncReplicationDestinationSpec {
	if in == nil {
		return nil
	}
	out := new(VolSyncReplicationDestinationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolSyncReplicationSourceSpec) DeepCopyInto(out *VolSyncReplicationSourceSpec) {
	*out = *in
	in.ProtectedPVC.DeepCopyInto(&out.ProtectedPVC)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolSyncReplicationSourceSpec.
func (in *VolSyncReplicationSourceSpec) DeepCopy() *VolSyncReplicationSourceSpec {
	if in == nil {
		return nil
	}
	out := new(VolSyncReplicationSourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolSyncSpec) DeepCopyInto(out *VolSyncSpec) {
	*out = *in
	if in.RDSpec != nil {
		in, out := &in.RDSpec, &out.RDSpec
		*out = make([]VolSyncReplicationDestinationSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolSyncSpec.
func (in *VolSyncSpec) DeepCopy() *VolSyncSpec {
	if in == nil {
		return nil
	}
	out := new(VolSyncSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeReplicationGroup) DeepCopyInto(out *VolumeReplicationGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeReplicationGroup.
func (in *VolumeReplicationGroup) DeepCopy() *VolumeReplicationGroup {
	if in == nil {
		return nil
	}
	out := new(VolumeReplicationGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VolumeReplicationGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeReplicationGroupList) DeepCopyInto(out *VolumeReplicationGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VolumeReplicationGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeReplicationGroupList.
func (in *VolumeReplicationGroupList) DeepCopy() *VolumeReplicationGroupList {
	if in == nil {
		return nil
	}
	out := new(VolumeReplicationGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VolumeReplicationGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeReplicationGroupSpec) DeepCopyInto(out *VolumeReplicationGroupSpec) {
	*out = *in
	in.PVCSelector.DeepCopyInto(&out.PVCSelector)
	if in.S3Profiles != nil {
		in, out := &in.S3Profiles, &out.S3Profiles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Async.DeepCopyInto(&out.Async)
	out.Sync = in.Sync
	in.VolSync.DeepCopyInto(&out.VolSync)
	if in.KubeObjectProtection != nil {
		in, out := &in.KubeObjectProtection, &out.KubeObjectProtection
		*out = new(KubeObjectProtectionSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeReplicationGroupSpec.
func (in *VolumeReplicationGroupSpec) DeepCopy() *VolumeReplicationGroupSpec {
	if in == nil {
		return nil
	}
	out := new(VolumeReplicationGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeReplicationGroupStatus) DeepCopyInto(out *VolumeReplicationGroupStatus) {
	*out = *in
	if in.ProtectedPVCs != nil {
		in, out := &in.ProtectedPVCs, &out.ProtectedPVCs
		*out = make([]ProtectedPVC, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	in.KubeObjectProtection.DeepCopyInto(&out.KubeObjectProtection)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeReplicationGroupStatus.
func (in *VolumeReplicationGroupStatus) DeepCopy() *VolumeReplicationGroupStatus {
	if in == nil {
		return nil
	}
	out := new(VolumeReplicationGroupStatus)
	in.DeepCopyInto(out)
	return out
}
