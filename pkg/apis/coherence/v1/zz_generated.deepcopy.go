// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoherenceCluster) DeepCopyInto(out *CoherenceCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoherenceCluster.
func (in *CoherenceCluster) DeepCopy() *CoherenceCluster {
	if in == nil {
		return nil
	}
	out := new(CoherenceCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CoherenceCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoherenceClusterList) DeepCopyInto(out *CoherenceClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CoherenceCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoherenceClusterList.
func (in *CoherenceClusterList) DeepCopy() *CoherenceClusterList {
	if in == nil {
		return nil
	}
	out := new(CoherenceClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CoherenceClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoherenceClusterSpec) DeepCopyInto(out *CoherenceClusterSpec) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.CoherenceRoleSpec.DeepCopyInto(&out.CoherenceRoleSpec)
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]CoherenceRoleSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoherenceClusterSpec.
func (in *CoherenceClusterSpec) DeepCopy() *CoherenceClusterSpec {
	if in == nil {
		return nil
	}
	out := new(CoherenceClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoherenceClusterStatus) DeepCopyInto(out *CoherenceClusterStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoherenceClusterStatus.
func (in *CoherenceClusterStatus) DeepCopy() *CoherenceClusterStatus {
	if in == nil {
		return nil
	}
	out := new(CoherenceClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoherenceInternal) DeepCopyInto(out *CoherenceInternal) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoherenceInternal.
func (in *CoherenceInternal) DeepCopy() *CoherenceInternal {
	if in == nil {
		return nil
	}
	out := new(CoherenceInternal)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CoherenceInternal) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoherenceInternalList) DeepCopyInto(out *CoherenceInternalList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CoherenceInternal, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoherenceInternalList.
func (in *CoherenceInternalList) DeepCopy() *CoherenceInternalList {
	if in == nil {
		return nil
	}
	out := new(CoherenceInternalList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CoherenceInternalList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoherenceInternalSpec) DeepCopyInto(out *CoherenceInternalSpec) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Coherence != nil {
		in, out := &in.Coherence, &out.Coherence
		*out = new(ImageSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.CoherenceUtils != nil {
		in, out := &in.CoherenceUtils, &out.CoherenceUtils
		*out = new(ImageSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Store != nil {
		in, out := &in.Store, &out.Store
		*out = new(CoherenceInternalStoreSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Fluentd != nil {
		in, out := &in.Fluentd, &out.Fluentd
		*out = new(FluentdImageSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.UserArtifacts != nil {
		in, out := &in.UserArtifacts, &out.UserArtifacts
		*out = new(UserArtifactsImageSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoherenceInternalSpec.
func (in *CoherenceInternalSpec) DeepCopy() *CoherenceInternalSpec {
	if in == nil {
		return nil
	}
	out := new(CoherenceInternalSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoherenceInternalStatus) DeepCopyInto(out *CoherenceInternalStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoherenceInternalStatus.
func (in *CoherenceInternalStatus) DeepCopy() *CoherenceInternalStatus {
	if in == nil {
		return nil
	}
	out := new(CoherenceInternalStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoherenceInternalStoreSpec) DeepCopyInto(out *CoherenceInternalStoreSpec) {
	*out = *in
	if in.StorageEnabled != nil {
		in, out := &in.StorageEnabled, &out.StorageEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ReadinessProbe != nil {
		in, out := &in.ReadinessProbe, &out.ReadinessProbe
		*out = new(ReadinessProbeSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.CacheConfig != nil {
		in, out := &in.CacheConfig, &out.CacheConfig
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoherenceInternalStoreSpec.
func (in *CoherenceInternalStoreSpec) DeepCopy() *CoherenceInternalStoreSpec {
	if in == nil {
		return nil
	}
	out := new(CoherenceInternalStoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoherenceRole) DeepCopyInto(out *CoherenceRole) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoherenceRole.
func (in *CoherenceRole) DeepCopy() *CoherenceRole {
	if in == nil {
		return nil
	}
	out := new(CoherenceRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CoherenceRole) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoherenceRoleList) DeepCopyInto(out *CoherenceRoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CoherenceRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoherenceRoleList.
func (in *CoherenceRoleList) DeepCopy() *CoherenceRoleList {
	if in == nil {
		return nil
	}
	out := new(CoherenceRoleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CoherenceRoleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoherenceRoleSpec) DeepCopyInto(out *CoherenceRoleSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = new(Images)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageEnabled != nil {
		in, out := &in.StorageEnabled, &out.StorageEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ScalingPolicy != nil {
		in, out := &in.ScalingPolicy, &out.ScalingPolicy
		*out = new(ScalingPolicy)
		**out = **in
	}
	if in.ReadinessProbe != nil {
		in, out := &in.ReadinessProbe, &out.ReadinessProbe
		*out = new(ReadinessProbeSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.CacheConfig != nil {
		in, out := &in.CacheConfig, &out.CacheConfig
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoherenceRoleSpec.
func (in *CoherenceRoleSpec) DeepCopy() *CoherenceRoleSpec {
	if in == nil {
		return nil
	}
	out := new(CoherenceRoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoherenceRoleStatus) DeepCopyInto(out *CoherenceRoleStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoherenceRoleStatus.
func (in *CoherenceRoleStatus) DeepCopy() *CoherenceRoleStatus {
	if in == nil {
		return nil
	}
	out := new(CoherenceRoleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FluentdApplicationSpec) DeepCopyInto(out *FluentdApplicationSpec) {
	*out = *in
	if in.ConfigFile != nil {
		in, out := &in.ConfigFile, &out.ConfigFile
		*out = new(string)
		**out = **in
	}
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FluentdApplicationSpec.
func (in *FluentdApplicationSpec) DeepCopy() *FluentdApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(FluentdApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FluentdImageSpec) DeepCopyInto(out *FluentdImageSpec) {
	*out = *in
	in.ImageSpec.DeepCopyInto(&out.ImageSpec)
	if in.Application != nil {
		in, out := &in.Application, &out.Application
		*out = new(FluentdApplicationSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FluentdImageSpec.
func (in *FluentdImageSpec) DeepCopy() *FluentdImageSpec {
	if in == nil {
		return nil
	}
	out := new(FluentdImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSpec) DeepCopyInto(out *ImageSpec) {
	*out = *in
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	if in.ImagePullPolicy != nil {
		in, out := &in.ImagePullPolicy, &out.ImagePullPolicy
		*out = new(corev1.PullPolicy)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSpec.
func (in *ImageSpec) DeepCopy() *ImageSpec {
	if in == nil {
		return nil
	}
	out := new(ImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Images) DeepCopyInto(out *Images) {
	*out = *in
	if in.Coherence != nil {
		in, out := &in.Coherence, &out.Coherence
		*out = new(ImageSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.CoherenceUtils != nil {
		in, out := &in.CoherenceUtils, &out.CoherenceUtils
		*out = new(ImageSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.UserArtifacts != nil {
		in, out := &in.UserArtifacts, &out.UserArtifacts
		*out = new(UserArtifactsImageSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Fluentd != nil {
		in, out := &in.Fluentd, &out.Fluentd
		*out = new(FluentdImageSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Images.
func (in *Images) DeepCopy() *Images {
	if in == nil {
		return nil
	}
	out := new(Images)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReadinessProbeSpec) DeepCopyInto(out *ReadinessProbeSpec) {
	*out = *in
	if in.InitialDelaySeconds != nil {
		in, out := &in.InitialDelaySeconds, &out.InitialDelaySeconds
		*out = new(int32)
		**out = **in
	}
	if in.TimeoutSeconds != nil {
		in, out := &in.TimeoutSeconds, &out.TimeoutSeconds
		*out = new(int32)
		**out = **in
	}
	if in.PeriodSeconds != nil {
		in, out := &in.PeriodSeconds, &out.PeriodSeconds
		*out = new(int32)
		**out = **in
	}
	if in.SuccessThreshold != nil {
		in, out := &in.SuccessThreshold, &out.SuccessThreshold
		*out = new(int32)
		**out = **in
	}
	if in.FailureThreshold != nil {
		in, out := &in.FailureThreshold, &out.FailureThreshold
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReadinessProbeSpec.
func (in *ReadinessProbeSpec) DeepCopy() *ReadinessProbeSpec {
	if in == nil {
		return nil
	}
	out := new(ReadinessProbeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserArtifactsImageSpec) DeepCopyInto(out *UserArtifactsImageSpec) {
	*out = *in
	in.ImageSpec.DeepCopyInto(&out.ImageSpec)
	if in.LibDir != nil {
		in, out := &in.LibDir, &out.LibDir
		*out = new(string)
		**out = **in
	}
	if in.ConfigDir != nil {
		in, out := &in.ConfigDir, &out.ConfigDir
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserArtifactsImageSpec.
func (in *UserArtifactsImageSpec) DeepCopy() *UserArtifactsImageSpec {
	if in == nil {
		return nil
	}
	out := new(UserArtifactsImageSpec)
	in.DeepCopyInto(out)
	return out
}
