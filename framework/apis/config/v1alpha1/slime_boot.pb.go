// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: slime_boot.proto

package v1alpha1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	v1 "k8s.io/api/core/v1"
	_ "k8s.io/apimachinery/pkg/api/resource"
	_ "k8s.io/apimachinery/pkg/apis/meta/v1"
	_ "k8s.io/apimachinery/pkg/runtime"
	_ "k8s.io/apimachinery/pkg/runtime/schema"
	_ "k8s.io/apimachinery/pkg/util/intstr"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type SlimeBootStatus struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SlimeBootStatus) Reset()         { *m = SlimeBootStatus{} }
func (m *SlimeBootStatus) String() string { return proto.CompactTextString(m) }
func (*SlimeBootStatus) ProtoMessage()    {}
func (*SlimeBootStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bf8433223eb5dea, []int{0}
}
func (m *SlimeBootStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SlimeBootStatus.Unmarshal(m, b)
}
func (m *SlimeBootStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SlimeBootStatus.Marshal(b, m, deterministic)
}
func (m *SlimeBootStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SlimeBootStatus.Merge(m, src)
}
func (m *SlimeBootStatus) XXX_Size() int {
	return xxx_messageInfo_SlimeBootStatus.Size(m)
}
func (m *SlimeBootStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_SlimeBootStatus.DiscardUnknown(m)
}

var xxx_messageInfo_SlimeBootStatus proto.InternalMessageInfo

type SlimeBootSpec struct {
	Module                   []*Config                  `protobuf:"bytes,1,rep,name=module,proto3" json:"module,omitempty"`
	Component                *Component                 `protobuf:"bytes,2,opt,name=component,proto3" json:"component,omitempty"`
	Namespace                string                     `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	IstioNamespace           string                     `protobuf:"bytes,4,opt,name=istioNamespace,proto3" json:"istioNamespace,omitempty"`
	Image                    *Image                     `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
	ReplicaCount             int32                      `protobuf:"varint,6,opt,name=replicaCount,proto3" json:"replicaCount,omitempty"`
	Service                  *Service                   `protobuf:"bytes,7,opt,name=service,proto3" json:"service,omitempty"`
	Resources                *v1.ResourceRequirements   `protobuf:"bytes,8,opt,name=resources,proto3" json:"resources,omitempty"`
	ImagePullSecrets         []*v1.LocalObjectReference `protobuf:"bytes,9,rep,name=imagePullSecrets,proto3" json:"imagePullSecrets,omitempty"`
	NodeSelector             map[string]string          `protobuf:"bytes,10,rep,name=nodeSelector,proto3" json:"nodeSelector,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	PodSecurityContext       *v1.PodSecurityContext     `protobuf:"bytes,11,opt,name=podSecurityContext,proto3" json:"podSecurityContext,omitempty"`
	ContainerSecurityContext *v1.SecurityContext        `protobuf:"bytes,12,opt,name=containerSecurityContext,proto3" json:"containerSecurityContext,omitempty"`
	Tolerations              []*v1.Toleration           `protobuf:"bytes,13,rep,name=tolerations,proto3" json:"tolerations,omitempty"`
	Affinity                 *v1.Affinity               `protobuf:"bytes,14,opt,name=affinity,proto3" json:"affinity,omitempty"`
	Volumes                  []*v1.Volume               `protobuf:"bytes,15,rep,name=volumes,proto3" json:"volumes,omitempty"`
	VolumeMounts             []*v1.VolumeMount          `protobuf:"bytes,16,rep,name=volumeMounts,proto3" json:"volumeMounts,omitempty"`
	ServiceAccount           *ServiceAccount            `protobuf:"bytes,17,opt,name=serviceAccount,proto3" json:"serviceAccount,omitempty"`
	Args                     []string                   `protobuf:"bytes,18,rep,name=args,proto3" json:"args,omitempty"`
	Env                      []*v1.EnvVar               `protobuf:"bytes,19,rep,name=env,proto3" json:"env,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                   `json:"-"`
	XXX_unrecognized         []byte                     `json:"-"`
	XXX_sizecache            int32                      `json:"-"`
}

func (m *SlimeBootSpec) Reset()         { *m = SlimeBootSpec{} }
func (m *SlimeBootSpec) String() string { return proto.CompactTextString(m) }
func (*SlimeBootSpec) ProtoMessage()    {}
func (*SlimeBootSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bf8433223eb5dea, []int{1}
}
func (m *SlimeBootSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SlimeBootSpec.Unmarshal(m, b)
}
func (m *SlimeBootSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SlimeBootSpec.Marshal(b, m, deterministic)
}
func (m *SlimeBootSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SlimeBootSpec.Merge(m, src)
}
func (m *SlimeBootSpec) XXX_Size() int {
	return xxx_messageInfo_SlimeBootSpec.Size(m)
}
func (m *SlimeBootSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_SlimeBootSpec.DiscardUnknown(m)
}

var xxx_messageInfo_SlimeBootSpec proto.InternalMessageInfo

func (m *SlimeBootSpec) GetModule() []*Config {
	if m != nil {
		return m.Module
	}
	return nil
}

func (m *SlimeBootSpec) GetComponent() *Component {
	if m != nil {
		return m.Component
	}
	return nil
}

func (m *SlimeBootSpec) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *SlimeBootSpec) GetIstioNamespace() string {
	if m != nil {
		return m.IstioNamespace
	}
	return ""
}

func (m *SlimeBootSpec) GetImage() *Image {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *SlimeBootSpec) GetReplicaCount() int32 {
	if m != nil {
		return m.ReplicaCount
	}
	return 0
}

func (m *SlimeBootSpec) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *SlimeBootSpec) GetResources() *v1.ResourceRequirements {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *SlimeBootSpec) GetImagePullSecrets() []*v1.LocalObjectReference {
	if m != nil {
		return m.ImagePullSecrets
	}
	return nil
}

func (m *SlimeBootSpec) GetNodeSelector() map[string]string {
	if m != nil {
		return m.NodeSelector
	}
	return nil
}

func (m *SlimeBootSpec) GetPodSecurityContext() *v1.PodSecurityContext {
	if m != nil {
		return m.PodSecurityContext
	}
	return nil
}

func (m *SlimeBootSpec) GetContainerSecurityContext() *v1.SecurityContext {
	if m != nil {
		return m.ContainerSecurityContext
	}
	return nil
}

func (m *SlimeBootSpec) GetTolerations() []*v1.Toleration {
	if m != nil {
		return m.Tolerations
	}
	return nil
}

func (m *SlimeBootSpec) GetAffinity() *v1.Affinity {
	if m != nil {
		return m.Affinity
	}
	return nil
}

func (m *SlimeBootSpec) GetVolumes() []*v1.Volume {
	if m != nil {
		return m.Volumes
	}
	return nil
}

func (m *SlimeBootSpec) GetVolumeMounts() []*v1.VolumeMount {
	if m != nil {
		return m.VolumeMounts
	}
	return nil
}

func (m *SlimeBootSpec) GetServiceAccount() *ServiceAccount {
	if m != nil {
		return m.ServiceAccount
	}
	return nil
}

func (m *SlimeBootSpec) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *SlimeBootSpec) GetEnv() []*v1.EnvVar {
	if m != nil {
		return m.Env
	}
	return nil
}

type ServiceAccount struct {
	Create               bool     `protobuf:"varint,1,opt,name=create,proto3" json:"create,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceAccount) Reset()         { *m = ServiceAccount{} }
func (m *ServiceAccount) String() string { return proto.CompactTextString(m) }
func (*ServiceAccount) ProtoMessage()    {}
func (*ServiceAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bf8433223eb5dea, []int{2}
}
func (m *ServiceAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceAccount.Unmarshal(m, b)
}
func (m *ServiceAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceAccount.Marshal(b, m, deterministic)
}
func (m *ServiceAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceAccount.Merge(m, src)
}
func (m *ServiceAccount) XXX_Size() int {
	return xxx_messageInfo_ServiceAccount.Size(m)
}
func (m *ServiceAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceAccount.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceAccount proto.InternalMessageInfo

func (m *ServiceAccount) GetCreate() bool {
	if m != nil {
		return m.Create
	}
	return false
}

type Component struct {
	GlobalSidecar        *GlobalSidecar `protobuf:"bytes,1,opt,name=globalSidecar,proto3" json:"globalSidecar,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Component) Reset()         { *m = Component{} }
func (m *Component) String() string { return proto.CompactTextString(m) }
func (*Component) ProtoMessage()    {}
func (*Component) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bf8433223eb5dea, []int{3}
}
func (m *Component) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Component.Unmarshal(m, b)
}
func (m *Component) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Component.Marshal(b, m, deterministic)
}
func (m *Component) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Component.Merge(m, src)
}
func (m *Component) XXX_Size() int {
	return xxx_messageInfo_Component.Size(m)
}
func (m *Component) XXX_DiscardUnknown() {
	xxx_messageInfo_Component.DiscardUnknown(m)
}

var xxx_messageInfo_Component proto.InternalMessageInfo

func (m *Component) GetGlobalSidecar() *GlobalSidecar {
	if m != nil {
		return m.GlobalSidecar
	}
	return nil
}

type GlobalSidecar struct {
	Enable               bool                         `protobuf:"varint,1,opt,name=enable,proto3" json:"enable,omitempty"`
	Port                 int32                        `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	ProbePort            int32                        `protobuf:"varint,3,opt,name=probePort,proto3" json:"probePort,omitempty"`
	Replicas             int32                        `protobuf:"varint,4,opt,name=replicas,proto3" json:"replicas,omitempty"`
	Image                *Image                       `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
	SidecarInject        *GlobalSidecar_SidecarInject `protobuf:"bytes,6,opt,name=sidecarInject,proto3" json:"sidecarInject,omitempty"`
	Resources            *v1.ResourceRequirements     `protobuf:"bytes,7,opt,name=resources,proto3" json:"resources,omitempty"`
	LegacyFilterName     bool                         `protobuf:"varint,8,opt,name=legacyFilterName,proto3" json:"legacyFilterName,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *GlobalSidecar) Reset()         { *m = GlobalSidecar{} }
func (m *GlobalSidecar) String() string { return proto.CompactTextString(m) }
func (*GlobalSidecar) ProtoMessage()    {}
func (*GlobalSidecar) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bf8433223eb5dea, []int{4}
}
func (m *GlobalSidecar) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GlobalSidecar.Unmarshal(m, b)
}
func (m *GlobalSidecar) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GlobalSidecar.Marshal(b, m, deterministic)
}
func (m *GlobalSidecar) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GlobalSidecar.Merge(m, src)
}
func (m *GlobalSidecar) XXX_Size() int {
	return xxx_messageInfo_GlobalSidecar.Size(m)
}
func (m *GlobalSidecar) XXX_DiscardUnknown() {
	xxx_messageInfo_GlobalSidecar.DiscardUnknown(m)
}

var xxx_messageInfo_GlobalSidecar proto.InternalMessageInfo

func (m *GlobalSidecar) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

func (m *GlobalSidecar) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *GlobalSidecar) GetProbePort() int32 {
	if m != nil {
		return m.ProbePort
	}
	return 0
}

func (m *GlobalSidecar) GetReplicas() int32 {
	if m != nil {
		return m.Replicas
	}
	return 0
}

func (m *GlobalSidecar) GetImage() *Image {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *GlobalSidecar) GetSidecarInject() *GlobalSidecar_SidecarInject {
	if m != nil {
		return m.SidecarInject
	}
	return nil
}

func (m *GlobalSidecar) GetResources() *v1.ResourceRequirements {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *GlobalSidecar) GetLegacyFilterName() bool {
	if m != nil {
		return m.LegacyFilterName
	}
	return false
}

type GlobalSidecar_SidecarInject struct {
	Enable               bool              `protobuf:"varint,1,opt,name=enable,proto3" json:"enable,omitempty"`
	Mode                 string            `protobuf:"bytes,2,opt,name=mode,proto3" json:"mode,omitempty"`
	Labels               map[string]string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Annotations          map[string]string `protobuf:"bytes,4,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GlobalSidecar_SidecarInject) Reset()         { *m = GlobalSidecar_SidecarInject{} }
func (m *GlobalSidecar_SidecarInject) String() string { return proto.CompactTextString(m) }
func (*GlobalSidecar_SidecarInject) ProtoMessage()    {}
func (*GlobalSidecar_SidecarInject) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bf8433223eb5dea, []int{4, 0}
}
func (m *GlobalSidecar_SidecarInject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GlobalSidecar_SidecarInject.Unmarshal(m, b)
}
func (m *GlobalSidecar_SidecarInject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GlobalSidecar_SidecarInject.Marshal(b, m, deterministic)
}
func (m *GlobalSidecar_SidecarInject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GlobalSidecar_SidecarInject.Merge(m, src)
}
func (m *GlobalSidecar_SidecarInject) XXX_Size() int {
	return xxx_messageInfo_GlobalSidecar_SidecarInject.Size(m)
}
func (m *GlobalSidecar_SidecarInject) XXX_DiscardUnknown() {
	xxx_messageInfo_GlobalSidecar_SidecarInject.DiscardUnknown(m)
}

var xxx_messageInfo_GlobalSidecar_SidecarInject proto.InternalMessageInfo

func (m *GlobalSidecar_SidecarInject) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

func (m *GlobalSidecar_SidecarInject) GetMode() string {
	if m != nil {
		return m.Mode
	}
	return ""
}

func (m *GlobalSidecar_SidecarInject) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *GlobalSidecar_SidecarInject) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

type Image struct {
	PullPolicy           string   `protobuf:"bytes,1,opt,name=pullPolicy,proto3" json:"pullPolicy,omitempty"`
	Repository           string   `protobuf:"bytes,2,opt,name=repository,proto3" json:"repository,omitempty"`
	Tag                  string   `protobuf:"bytes,3,opt,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bf8433223eb5dea, []int{5}
}
func (m *Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Image.Unmarshal(m, b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Image.Marshal(b, m, deterministic)
}
func (m *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(m, src)
}
func (m *Image) XXX_Size() int {
	return xxx_messageInfo_Image.Size(m)
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetPullPolicy() string {
	if m != nil {
		return m.PullPolicy
	}
	return ""
}

func (m *Image) GetRepository() string {
	if m != nil {
		return m.Repository
	}
	return ""
}

func (m *Image) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

type Service struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Port                 int32    `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	AuxiliaryPort        int32    `protobuf:"varint,3,opt,name=auxiliaryPort,proto3" json:"auxiliaryPort,omitempty"`
	LogSourcePort        int32    `protobuf:"varint,4,opt,name=logSourcePort,proto3" json:"logSourcePort,omitempty"`
	McpOverXdsPort       int32    `protobuf:"varint,5,opt,name=mcpOverXdsPort,proto3" json:"mcpOverXdsPort,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Service) Reset()         { *m = Service{} }
func (m *Service) String() string { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()    {}
func (*Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bf8433223eb5dea, []int{6}
}
func (m *Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service.Unmarshal(m, b)
}
func (m *Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service.Marshal(b, m, deterministic)
}
func (m *Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service.Merge(m, src)
}
func (m *Service) XXX_Size() int {
	return xxx_messageInfo_Service.Size(m)
}
func (m *Service) XXX_DiscardUnknown() {
	xxx_messageInfo_Service.DiscardUnknown(m)
}

var xxx_messageInfo_Service proto.InternalMessageInfo

func (m *Service) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Service) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Service) GetAuxiliaryPort() int32 {
	if m != nil {
		return m.AuxiliaryPort
	}
	return 0
}

func (m *Service) GetLogSourcePort() int32 {
	if m != nil {
		return m.LogSourcePort
	}
	return 0
}

func (m *Service) GetMcpOverXdsPort() int32 {
	if m != nil {
		return m.McpOverXdsPort
	}
	return 0
}

func init() {
	proto.RegisterType((*SlimeBootStatus)(nil), "slime.config.v1alpha1.SlimeBootStatus")
	proto.RegisterType((*SlimeBootSpec)(nil), "slime.config.v1alpha1.SlimeBootSpec")
	proto.RegisterMapType((map[string]string)(nil), "slime.config.v1alpha1.SlimeBootSpec.NodeSelectorEntry")
	proto.RegisterType((*ServiceAccount)(nil), "slime.config.v1alpha1.ServiceAccount")
	proto.RegisterType((*Component)(nil), "slime.config.v1alpha1.Component")
	proto.RegisterType((*GlobalSidecar)(nil), "slime.config.v1alpha1.GlobalSidecar")
	proto.RegisterType((*GlobalSidecar_SidecarInject)(nil), "slime.config.v1alpha1.GlobalSidecar.SidecarInject")
	proto.RegisterMapType((map[string]string)(nil), "slime.config.v1alpha1.GlobalSidecar.SidecarInject.AnnotationsEntry")
	proto.RegisterMapType((map[string]string)(nil), "slime.config.v1alpha1.GlobalSidecar.SidecarInject.LabelsEntry")
	proto.RegisterType((*Image)(nil), "slime.config.v1alpha1.Image")
	proto.RegisterType((*Service)(nil), "slime.config.v1alpha1.Service")
}

func init() { proto.RegisterFile("slime_boot.proto", fileDescriptor_6bf8433223eb5dea) }

var fileDescriptor_6bf8433223eb5dea = []byte{
	// 1033 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xdb, 0x6e, 0x1b, 0x37,
	0x10, 0x85, 0xa2, 0xc8, 0xb6, 0x46, 0x96, 0x23, 0xb3, 0x17, 0x2c, 0x84, 0xd4, 0x15, 0xd4, 0x34,
	0x10, 0x8a, 0x76, 0x17, 0x56, 0xd2, 0xc0, 0xed, 0x83, 0x5b, 0xc7, 0x48, 0x8a, 0x14, 0xb9, 0x18,
	0x54, 0xe0, 0xa6, 0x79, 0x09, 0x28, 0x6a, 0x2c, 0xb3, 0xe6, 0x92, 0x5b, 0x92, 0xab, 0x46, 0x1f,
	0xd4, 0x9f, 0xea, 0x4f, 0xf4, 0xa9, 0xef, 0xc5, 0x72, 0x57, 0x77, 0x29, 0xad, 0xfd, 0x36, 0x9c,
	0x3d, 0xe7, 0x70, 0x96, 0x9c, 0x39, 0x20, 0x34, 0xac, 0x14, 0x31, 0xbe, 0xeb, 0x6b, 0xed, 0xc2,
	0xc4, 0x68, 0xa7, 0xc9, 0x27, 0x3e, 0x13, 0x72, 0xad, 0x2e, 0xc4, 0x30, 0x1c, 0x1d, 0x32, 0x99,
	0x5c, 0xb2, 0xc3, 0xe6, 0x6e, 0x91, 0xf0, 0xa0, 0x66, 0xfb, 0xea, 0xc8, 0x86, 0x42, 0x47, 0x2c,
	0x11, 0x11, 0xd7, 0x06, 0xa3, 0xd1, 0x61, 0x34, 0x44, 0x85, 0x86, 0x39, 0x1c, 0x14, 0x98, 0x87,
	0x33, 0x4c, 0xcc, 0xf8, 0xa5, 0x50, 0x68, 0xc6, 0x51, 0x72, 0x35, 0xf4, 0x24, 0x83, 0x56, 0xa7,
	0x86, 0xe3, 0xb5, 0x58, 0x36, 0x8a, 0xd1, 0xb1, 0x75, 0x7b, 0x45, 0x9b, 0x58, 0x26, 0x55, 0x4e,
	0xc4, 0xab, 0xdb, 0x3c, 0xfa, 0x2f, 0x82, 0xe5, 0x97, 0x18, 0xb3, 0x15, 0xde, 0x83, 0x4d, 0xbc,
	0xd4, 0x09, 0x19, 0x09, 0xe5, 0xac, 0x33, 0xcb, 0xa4, 0xf6, 0x3e, 0xdc, 0xe9, 0x65, 0x87, 0xfa,
	0x58, 0x6b, 0xd7, 0x73, 0xcc, 0xa5, 0xb6, 0xfd, 0x57, 0x15, 0xea, 0xb3, 0x5c, 0x82, 0x9c, 0x7c,
	0x0b, 0x5b, 0xb1, 0x1e, 0xa4, 0x12, 0x83, 0x52, 0xab, 0xdc, 0xa9, 0x75, 0x3f, 0x0b, 0xd7, 0x5e,
	0x44, 0x78, 0xea, 0xd7, 0xb4, 0x00, 0x93, 0x63, 0xa8, 0x72, 0x1d, 0x27, 0x5a, 0xa1, 0x72, 0xc1,
	0xad, 0x56, 0xa9, 0x53, 0xeb, 0xb6, 0x36, 0x32, 0x0b, 0x1c, 0x9d, 0x51, 0xc8, 0x5d, 0xa8, 0x2a,
	0x16, 0xa3, 0x4d, 0x18, 0xc7, 0xa0, 0xdc, 0x2a, 0x75, 0xaa, 0x74, 0x96, 0x20, 0xf7, 0x61, 0x4f,
	0x58, 0x27, 0xf4, 0xcb, 0x29, 0xe4, 0xb6, 0x87, 0x2c, 0x65, 0x49, 0x17, 0x2a, 0x22, 0x66, 0x43,
	0x0c, 0x2a, 0xbe, 0x82, 0xbb, 0x1b, 0x2a, 0x78, 0x96, 0x61, 0x68, 0x0e, 0x25, 0x6d, 0xd8, 0x35,
	0x98, 0x48, 0xc1, 0xd9, 0xa9, 0x4e, 0x95, 0x0b, 0xb6, 0x5a, 0xa5, 0x4e, 0x85, 0x2e, 0xe4, 0xc8,
	0x11, 0x6c, 0x5b, 0x34, 0x23, 0xc1, 0x31, 0xd8, 0xf6, 0xca, 0x07, 0x1b, 0x94, 0x7b, 0x39, 0x8a,
	0x4e, 0xe0, 0xe4, 0x29, 0x54, 0x27, 0x3d, 0x66, 0x83, 0x1d, 0xcf, 0xed, 0x84, 0xf9, 0xe5, 0x85,
	0x2c, 0x11, 0x61, 0xd6, 0xb5, 0xe1, 0xe8, 0x30, 0xa4, 0x05, 0x88, 0xe2, 0xef, 0xa9, 0x30, 0x18,
	0xa3, 0x72, 0x96, 0xce, 0xa8, 0xe4, 0x35, 0x34, 0x7c, 0xb9, 0x67, 0xa9, 0x94, 0x3d, 0xe4, 0x06,
	0x9d, 0x0d, 0xaa, 0xfe, 0x82, 0xd6, 0xca, 0x3d, 0xd7, 0x9c, 0xc9, 0x57, 0xfd, 0xdf, 0x90, 0x3b,
	0x8a, 0x17, 0x68, 0x50, 0x71, 0xa4, 0x2b, 0x0a, 0xe4, 0x2d, 0xec, 0x2a, 0x3d, 0xc0, 0x1e, 0x4a,
	0xe4, 0x4e, 0x9b, 0x00, 0xbc, 0xe2, 0xa3, 0x4d, 0x3f, 0x37, 0xdf, 0x28, 0xe1, 0xcb, 0x39, 0xe2,
	0x13, 0xe5, 0xcc, 0x98, 0x2e, 0x68, 0x91, 0x73, 0x20, 0x89, 0x1e, 0xf4, 0x90, 0xa7, 0x46, 0xb8,
	0xf1, 0xa9, 0x56, 0x0e, 0xdf, 0xbb, 0xa0, 0xe6, 0x8f, 0xe0, 0xfe, 0xba, 0x9a, 0xcf, 0x56, 0xd0,
	0x74, 0x8d, 0x02, 0x79, 0x07, 0x01, 0xd7, 0xca, 0xb1, 0xac, 0xe7, 0x97, 0xd5, 0x77, 0xbd, 0xfa,
	0x17, 0xeb, 0xd4, 0x97, 0xa5, 0x37, 0x8a, 0x90, 0x1f, 0xa1, 0xe6, 0xb4, 0xcc, 0x26, 0x47, 0x68,
	0x65, 0x83, 0xba, 0x3f, 0x93, 0x83, 0x75, 0x9a, 0xaf, 0xa7, 0x30, 0x3a, 0x4f, 0x21, 0x47, 0xb0,
	0xc3, 0x2e, 0x2e, 0x84, 0x12, 0x6e, 0x1c, 0xec, 0x15, 0x9d, 0xb8, 0x86, 0x7e, 0x52, 0x60, 0xe8,
	0x14, 0x4d, 0x1e, 0xc2, 0xf6, 0x48, 0xcb, 0x34, 0x46, 0x1b, 0xdc, 0xf1, 0xfb, 0x36, 0xd7, 0x11,
	0xcf, 0x3d, 0x84, 0x4e, 0xa0, 0xe4, 0x14, 0x76, 0xf3, 0xf0, 0x45, 0xd6, 0xad, 0x36, 0x68, 0x78,
	0xea, 0xe7, 0x9b, 0xa9, 0x1e, 0x47, 0x17, 0x48, 0xe4, 0x05, 0xec, 0x15, 0x4d, 0x7b, 0xc2, 0xb9,
	0x9f, 0x84, 0x7d, 0x5f, 0xfa, 0x97, 0x1f, 0x6e, 0xf5, 0x02, 0x4c, 0x97, 0xc8, 0x84, 0xc0, 0x6d,
	0x66, 0x86, 0x36, 0x20, 0xad, 0x72, 0xa7, 0x4a, 0x7d, 0x4c, 0xbe, 0x86, 0x32, 0xaa, 0x51, 0xf0,
	0xd1, 0xe6, 0x3f, 0x7b, 0xa2, 0x46, 0xe7, 0xcc, 0xd0, 0x0c, 0xd6, 0xfc, 0x01, 0xf6, 0x57, 0x7a,
	0x8c, 0x34, 0xa0, 0x7c, 0x85, 0xe3, 0xa0, 0xe4, 0xc7, 0x3f, 0x0b, 0xc9, 0xc7, 0x50, 0x19, 0x31,
	0x99, 0xa2, 0x77, 0x9d, 0x2a, 0xcd, 0x17, 0xdf, 0xdf, 0x3a, 0x2a, 0xb5, 0x3b, 0xb0, 0xb7, 0x58,
	0x24, 0xf9, 0x14, 0xb6, 0xb8, 0x41, 0xe6, 0xd0, 0x0b, 0xec, 0xd0, 0x62, 0xd5, 0xfe, 0x05, 0xaa,
	0x53, 0x57, 0x22, 0x3f, 0x43, 0x7d, 0x28, 0x75, 0x9f, 0xc9, 0x9e, 0x18, 0x20, 0x67, 0xc6, 0x63,
	0x6b, 0xdd, 0x7b, 0x1b, 0xce, 0xe1, 0xa7, 0x79, 0x2c, 0x5d, 0xa4, 0xb6, 0xff, 0xae, 0x40, 0x7d,
	0x01, 0x90, 0x95, 0x80, 0x8a, 0xf5, 0xe5, 0xb4, 0x84, 0x7c, 0x95, 0x9d, 0x57, 0xa2, 0x4d, 0xee,
	0x9d, 0x15, 0xea, 0xe3, 0xcc, 0x14, 0x13, 0xa3, 0xfb, 0x78, 0x96, 0x7d, 0x28, 0xfb, 0x0f, 0xb3,
	0x04, 0x69, 0xc2, 0x4e, 0x61, 0x52, 0xd6, 0xdb, 0x61, 0x85, 0x4e, 0xd7, 0x37, 0x32, 0xc2, 0x37,
	0x50, 0xb7, 0x79, 0x91, 0xcf, 0x54, 0x66, 0x1c, 0xde, 0x09, 0x6b, 0xdd, 0xee, 0xff, 0xf9, 0xef,
	0xb0, 0x37, 0xcf, 0xa4, 0x8b, 0x42, 0x8b, 0x26, 0xb8, 0x7d, 0x73, 0x13, 0xfc, 0x0a, 0x1a, 0x12,
	0x87, 0x8c, 0x8f, 0x9f, 0x0a, 0xe9, 0xd0, 0x64, 0xbe, 0xef, 0x3d, 0x75, 0x87, 0xae, 0xe4, 0x9b,
	0xff, 0xdc, 0x82, 0xfa, 0x42, 0x51, 0x1f, 0x3a, 0xf9, 0x58, 0x0f, 0x26, 0xfd, 0xe3, 0x63, 0x72,
	0x0e, 0x5b, 0x92, 0xf5, 0x51, 0xda, 0xa0, 0xec, 0x9b, 0xf5, 0xf8, 0xfa, 0x87, 0x10, 0x3e, 0xf7,
	0x02, 0xb9, 0x35, 0x16, 0x6a, 0x04, 0xa1, 0xc6, 0x94, 0xd2, 0xae, 0xf0, 0x96, 0xdb, 0x5e, 0xfc,
	0xf4, 0x06, 0xe2, 0x27, 0x33, 0x95, 0x7c, 0x87, 0x79, 0xdd, 0xe6, 0x77, 0x50, 0x9b, 0xdb, 0xfd,
	0x3a, 0x43, 0xd3, 0x3c, 0x86, 0xc6, 0xb2, 0xf6, 0xb5, 0x86, 0xee, 0x57, 0xa8, 0xf8, 0xae, 0x22,
	0x07, 0x00, 0x49, 0x2a, 0xe5, 0x99, 0x96, 0x82, 0x4f, 0xb8, 0x73, 0x99, 0xec, 0xbb, 0xc1, 0x44,
	0x5b, 0xe1, 0xb4, 0x19, 0x17, 0x3a, 0x73, 0x99, 0x6c, 0x53, 0xc7, 0x86, 0xc5, 0x5b, 0x20, 0x0b,
	0xdb, 0x7f, 0x96, 0x60, 0xbb, 0x18, 0xe8, 0xec, 0xd2, 0xdc, 0x38, 0xc1, 0x42, 0xd7, 0xc7, 0x6b,
	0x47, 0xe8, 0x1e, 0xd4, 0x59, 0xfa, 0x5e, 0x48, 0xc1, 0xcc, 0x78, 0x6e, 0x8c, 0x16, 0x93, 0x19,
	0x4a, 0xea, 0x61, 0xcf, 0xb7, 0x99, 0x47, 0xe5, 0xf3, 0xb4, 0x98, 0xcc, 0x5e, 0x21, 0x31, 0x4f,
	0x5e, 0x8d, 0xd0, 0xbc, 0x19, 0x58, 0x0f, 0xab, 0x78, 0xd8, 0x52, 0xf6, 0x71, 0xf4, 0xf6, 0x9b,
	0xfc, 0x42, 0x85, 0x8e, 0x7c, 0x10, 0x5d, 0x18, 0x16, 0xe3, 0x1f, 0xda, 0x5c, 0xe5, 0xcf, 0xc7,
	0xfc, 0xa2, 0xa3, 0xc9, 0x45, 0xf7, 0xb7, 0xfc, 0xfb, 0xec, 0xc1, 0xbf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xe7, 0x4f, 0x67, 0xf6, 0x06, 0x0b, 0x00, 0x00,
}
