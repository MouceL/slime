// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: limiter_module.proto

package config

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/durationpb"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// DeepCopyInto supports using Limiter within kubernetes types, where deepcopy-gen is used.
func (in *Limiter) DeepCopyInto(out *Limiter) {
	p := proto.Clone(in).(*Limiter)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Limiter. Required by controller-gen.
func (in *Limiter) DeepCopy() *Limiter {
	if in == nil {
		return nil
	}
	out := new(Limiter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Limiter. Required by controller-gen.
func (in *Limiter) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using RlsConfigMap within kubernetes types, where deepcopy-gen is used.
func (in *RlsConfigMap) DeepCopyInto(out *RlsConfigMap) {
	p := proto.Clone(in).(*RlsConfigMap)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RlsConfigMap. Required by controller-gen.
func (in *RlsConfigMap) DeepCopy() *RlsConfigMap {
	if in == nil {
		return nil
	}
	out := new(RlsConfigMap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new RlsConfigMap. Required by controller-gen.
func (in *RlsConfigMap) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using RateLimitService within kubernetes types, where deepcopy-gen is used.
func (in *RateLimitService) DeepCopyInto(out *RateLimitService) {
	p := proto.Clone(in).(*RateLimitService)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitService. Required by controller-gen.
func (in *RateLimitService) DeepCopy() *RateLimitService {
	if in == nil {
		return nil
	}
	out := new(RateLimitService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitService. Required by controller-gen.
func (in *RateLimitService) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}
