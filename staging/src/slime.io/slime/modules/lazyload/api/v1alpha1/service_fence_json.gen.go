// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: service_fence.proto

package v1alpha1

import (
	bytes "bytes"
	fmt "fmt"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for Timestamp
func (this *Timestamp) MarshalJSON() ([]byte, error) {
	str, err := ServiceFenceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Timestamp
func (this *Timestamp) UnmarshalJSON(b []byte) error {
	return ServiceFenceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ServiceFenceSpec
func (this *ServiceFenceSpec) MarshalJSON() ([]byte, error) {
	str, err := ServiceFenceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ServiceFenceSpec
func (this *ServiceFenceSpec) UnmarshalJSON(b []byte) error {
	return ServiceFenceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Selector
func (this *Selector) MarshalJSON() ([]byte, error) {
	str, err := ServiceFenceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Selector
func (this *Selector) UnmarshalJSON(b []byte) error {
	return ServiceFenceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for WorkloadSelector
func (this *WorkloadSelector) MarshalJSON() ([]byte, error) {
	str, err := ServiceFenceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for WorkloadSelector
func (this *WorkloadSelector) UnmarshalJSON(b []byte) error {
	return ServiceFenceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for RecyclingStrategy
func (this *RecyclingStrategy) MarshalJSON() ([]byte, error) {
	str, err := ServiceFenceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for RecyclingStrategy
func (this *RecyclingStrategy) UnmarshalJSON(b []byte) error {
	return ServiceFenceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for RecyclingStrategy_Stable
func (this *RecyclingStrategy_Stable) MarshalJSON() ([]byte, error) {
	str, err := ServiceFenceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for RecyclingStrategy_Stable
func (this *RecyclingStrategy_Stable) UnmarshalJSON(b []byte) error {
	return ServiceFenceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for RecyclingStrategy_Deadline
func (this *RecyclingStrategy_Deadline) MarshalJSON() ([]byte, error) {
	str, err := ServiceFenceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for RecyclingStrategy_Deadline
func (this *RecyclingStrategy_Deadline) UnmarshalJSON(b []byte) error {
	return ServiceFenceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for RecyclingStrategy_Auto
func (this *RecyclingStrategy_Auto) MarshalJSON() ([]byte, error) {
	str, err := ServiceFenceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for RecyclingStrategy_Auto
func (this *RecyclingStrategy_Auto) UnmarshalJSON(b []byte) error {
	return ServiceFenceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Destinations
func (this *Destinations) MarshalJSON() ([]byte, error) {
	str, err := ServiceFenceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Destinations
func (this *Destinations) UnmarshalJSON(b []byte) error {
	return ServiceFenceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ServiceFenceStatus
func (this *ServiceFenceStatus) MarshalJSON() ([]byte, error) {
	str, err := ServiceFenceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ServiceFenceStatus
func (this *ServiceFenceStatus) UnmarshalJSON(b []byte) error {
	return ServiceFenceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	ServiceFenceMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	ServiceFenceUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{AllowUnknownFields: true}
)
