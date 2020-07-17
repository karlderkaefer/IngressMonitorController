// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressMonitor) DeepCopyInto(out *IngressMonitor) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressMonitor.
func (in *IngressMonitor) DeepCopy() *IngressMonitor {
	if in == nil {
		return nil
	}
	out := new(IngressMonitor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IngressMonitor) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressMonitorList) DeepCopyInto(out *IngressMonitorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IngressMonitor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressMonitorList.
func (in *IngressMonitorList) DeepCopy() *IngressMonitorList {
	if in == nil {
		return nil
	}
	out := new(IngressMonitorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IngressMonitorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressMonitorSpec) DeepCopyInto(out *IngressMonitorSpec) {
	*out = *in
	if in.URLFrom != nil {
		in, out := &in.URLFrom, &out.URLFrom
		*out = new(URLSource)
		(*in).DeepCopyInto(*out)
	}
	if in.UptimeRobotConfig != nil {
		in, out := &in.UptimeRobotConfig, &out.UptimeRobotConfig
		*out = new(UptimeRobotConfig)
		**out = **in
	}
	if in.UptimeConfig != nil {
		in, out := &in.UptimeConfig, &out.UptimeConfig
		*out = new(UptimeConfig)
		**out = **in
	}
	if in.UpdownConfig != nil {
		in, out := &in.UpdownConfig, &out.UpdownConfig
		*out = new(UpdownConfig)
		**out = **in
	}
	if in.StatusCakeConfig != nil {
		in, out := &in.StatusCakeConfig, &out.StatusCakeConfig
		*out = new(StatusCakeConfig)
		**out = **in
	}
	if in.PingdomConfig != nil {
		in, out := &in.PingdomConfig, &out.PingdomConfig
		*out = new(PingdomConfig)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressMonitorSpec.
func (in *IngressMonitorSpec) DeepCopy() *IngressMonitorSpec {
	if in == nil {
		return nil
	}
	out := new(IngressMonitorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressMonitorStatus) DeepCopyInto(out *IngressMonitorStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressMonitorStatus.
func (in *IngressMonitorStatus) DeepCopy() *IngressMonitorStatus {
	if in == nil {
		return nil
	}
	out := new(IngressMonitorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressURLSource) DeepCopyInto(out *IngressURLSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressURLSource.
func (in *IngressURLSource) DeepCopy() *IngressURLSource {
	if in == nil {
		return nil
	}
	out := new(IngressURLSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PingdomConfig) DeepCopyInto(out *PingdomConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PingdomConfig.
func (in *PingdomConfig) DeepCopy() *PingdomConfig {
	if in == nil {
		return nil
	}
	out := new(PingdomConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteURLSource) DeepCopyInto(out *RouteURLSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteURLSource.
func (in *RouteURLSource) DeepCopy() *RouteURLSource {
	if in == nil {
		return nil
	}
	out := new(RouteURLSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusCakeConfig) DeepCopyInto(out *StatusCakeConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusCakeConfig.
func (in *StatusCakeConfig) DeepCopy() *StatusCakeConfig {
	if in == nil {
		return nil
	}
	out := new(StatusCakeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *URLSource) DeepCopyInto(out *URLSource) {
	*out = *in
	if in.IngressRef != nil {
		in, out := &in.IngressRef, &out.IngressRef
		*out = new(IngressURLSource)
		**out = **in
	}
	if in.RouteRef != nil {
		in, out := &in.RouteRef, &out.RouteRef
		*out = new(RouteURLSource)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new URLSource.
func (in *URLSource) DeepCopy() *URLSource {
	if in == nil {
		return nil
	}
	out := new(URLSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpdownConfig) DeepCopyInto(out *UpdownConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdownConfig.
func (in *UpdownConfig) DeepCopy() *UpdownConfig {
	if in == nil {
		return nil
	}
	out := new(UpdownConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UptimeConfig) DeepCopyInto(out *UptimeConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UptimeConfig.
func (in *UptimeConfig) DeepCopy() *UptimeConfig {
	if in == nil {
		return nil
	}
	out := new(UptimeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UptimeRobotConfig) DeepCopyInto(out *UptimeRobotConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UptimeRobotConfig.
func (in *UptimeRobotConfig) DeepCopy() *UptimeRobotConfig {
	if in == nil {
		return nil
	}
	out := new(UptimeRobotConfig)
	in.DeepCopyInto(out)
	return out
}
