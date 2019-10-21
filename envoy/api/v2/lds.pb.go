// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/lds.proto

package envoy_api_v2

import (
	context "context"
	fmt "fmt"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	listener "github.com/envoyproxy/go-control-plane/envoy/api/v2/listener"
	v2 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v2"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Listener_DrainType int32

const (
	Listener_DEFAULT     Listener_DrainType = 0
	Listener_MODIFY_ONLY Listener_DrainType = 1
)

var Listener_DrainType_name = map[int32]string{
	0: "DEFAULT",
	1: "MODIFY_ONLY",
}

var Listener_DrainType_value = map[string]int32{
	"DEFAULT":     0,
	"MODIFY_ONLY": 1,
}

func (x Listener_DrainType) String() string {
	return proto.EnumName(Listener_DrainType_name, int32(x))
}

func (Listener_DrainType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_34e2cd84a105bcd1, []int{0, 0}
}

type Listener struct {
	Name                             string                            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address                          *core.Address                     `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	FilterChains                     []*listener.FilterChain           `protobuf:"bytes,3,rep,name=filter_chains,json=filterChains,proto3" json:"filter_chains,omitempty"`
	UseOriginalDst                   *wrappers.BoolValue               `protobuf:"bytes,4,opt,name=use_original_dst,json=useOriginalDst,proto3" json:"use_original_dst,omitempty"` // Deprecated: Do not use.
	PerConnectionBufferLimitBytes    *wrappers.UInt32Value             `protobuf:"bytes,5,opt,name=per_connection_buffer_limit_bytes,json=perConnectionBufferLimitBytes,proto3" json:"per_connection_buffer_limit_bytes,omitempty"`
	Metadata                         *core.Metadata                    `protobuf:"bytes,6,opt,name=metadata,proto3" json:"metadata,omitempty"`
	DeprecatedV1                     *Listener_DeprecatedV1            `protobuf:"bytes,7,opt,name=deprecated_v1,json=deprecatedV1,proto3" json:"deprecated_v1,omitempty"`
	DrainType                        Listener_DrainType                `protobuf:"varint,8,opt,name=drain_type,json=drainType,proto3,enum=envoy.api.v2.Listener_DrainType" json:"drain_type,omitempty"`
	ListenerFilters                  []*listener.ListenerFilter        `protobuf:"bytes,9,rep,name=listener_filters,json=listenerFilters,proto3" json:"listener_filters,omitempty"`
	ListenerFiltersTimeout           *duration.Duration                `protobuf:"bytes,15,opt,name=listener_filters_timeout,json=listenerFiltersTimeout,proto3" json:"listener_filters_timeout,omitempty"`
	ContinueOnListenerFiltersTimeout bool                              `protobuf:"varint,17,opt,name=continue_on_listener_filters_timeout,json=continueOnListenerFiltersTimeout,proto3" json:"continue_on_listener_filters_timeout,omitempty"`
	Transparent                      *wrappers.BoolValue               `protobuf:"bytes,10,opt,name=transparent,proto3" json:"transparent,omitempty"`
	Freebind                         *wrappers.BoolValue               `protobuf:"bytes,11,opt,name=freebind,proto3" json:"freebind,omitempty"`
	SocketOptions                    []*core.SocketOption              `protobuf:"bytes,13,rep,name=socket_options,json=socketOptions,proto3" json:"socket_options,omitempty"`
	TcpFastOpenQueueLength           *wrappers.UInt32Value             `protobuf:"bytes,12,opt,name=tcp_fast_open_queue_length,json=tcpFastOpenQueueLength,proto3" json:"tcp_fast_open_queue_length,omitempty"`
	TrafficDirection                 core.TrafficDirection             `protobuf:"varint,16,opt,name=traffic_direction,json=trafficDirection,proto3,enum=envoy.api.v2.core.TrafficDirection" json:"traffic_direction,omitempty"`
	UdpListenerConfig                *listener.UdpListenerConfig       `protobuf:"bytes,18,opt,name=udp_listener_config,json=udpListenerConfig,proto3" json:"udp_listener_config,omitempty"`
	ApiListener                      *v2.ApiListener                   `protobuf:"bytes,19,opt,name=api_listener,json=apiListener,proto3" json:"api_listener,omitempty"`
	ConnectionBalanceConfig          *Listener_ConnectionBalanceConfig `protobuf:"bytes,20,opt,name=connection_balance_config,json=connectionBalanceConfig,proto3" json:"connection_balance_config,omitempty"`
	XXX_NoUnkeyedLiteral             struct{}                          `json:"-"`
	XXX_unrecognized                 []byte                            `json:"-"`
	XXX_sizecache                    int32                             `json:"-"`
}

func (m *Listener) Reset()         { *m = Listener{} }
func (m *Listener) String() string { return proto.CompactTextString(m) }
func (*Listener) ProtoMessage()    {}
func (*Listener) Descriptor() ([]byte, []int) {
	return fileDescriptor_34e2cd84a105bcd1, []int{0}
}

func (m *Listener) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Listener.Unmarshal(m, b)
}
func (m *Listener) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Listener.Marshal(b, m, deterministic)
}
func (m *Listener) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Listener.Merge(m, src)
}
func (m *Listener) XXX_Size() int {
	return xxx_messageInfo_Listener.Size(m)
}
func (m *Listener) XXX_DiscardUnknown() {
	xxx_messageInfo_Listener.DiscardUnknown(m)
}

var xxx_messageInfo_Listener proto.InternalMessageInfo

func (m *Listener) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Listener) GetAddress() *core.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Listener) GetFilterChains() []*listener.FilterChain {
	if m != nil {
		return m.FilterChains
	}
	return nil
}

// Deprecated: Do not use.
func (m *Listener) GetUseOriginalDst() *wrappers.BoolValue {
	if m != nil {
		return m.UseOriginalDst
	}
	return nil
}

func (m *Listener) GetPerConnectionBufferLimitBytes() *wrappers.UInt32Value {
	if m != nil {
		return m.PerConnectionBufferLimitBytes
	}
	return nil
}

func (m *Listener) GetMetadata() *core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Listener) GetDeprecatedV1() *Listener_DeprecatedV1 {
	if m != nil {
		return m.DeprecatedV1
	}
	return nil
}

func (m *Listener) GetDrainType() Listener_DrainType {
	if m != nil {
		return m.DrainType
	}
	return Listener_DEFAULT
}

func (m *Listener) GetListenerFilters() []*listener.ListenerFilter {
	if m != nil {
		return m.ListenerFilters
	}
	return nil
}

func (m *Listener) GetListenerFiltersTimeout() *duration.Duration {
	if m != nil {
		return m.ListenerFiltersTimeout
	}
	return nil
}

func (m *Listener) GetContinueOnListenerFiltersTimeout() bool {
	if m != nil {
		return m.ContinueOnListenerFiltersTimeout
	}
	return false
}

func (m *Listener) GetTransparent() *wrappers.BoolValue {
	if m != nil {
		return m.Transparent
	}
	return nil
}

func (m *Listener) GetFreebind() *wrappers.BoolValue {
	if m != nil {
		return m.Freebind
	}
	return nil
}

func (m *Listener) GetSocketOptions() []*core.SocketOption {
	if m != nil {
		return m.SocketOptions
	}
	return nil
}

func (m *Listener) GetTcpFastOpenQueueLength() *wrappers.UInt32Value {
	if m != nil {
		return m.TcpFastOpenQueueLength
	}
	return nil
}

func (m *Listener) GetTrafficDirection() core.TrafficDirection {
	if m != nil {
		return m.TrafficDirection
	}
	return core.TrafficDirection_UNSPECIFIED
}

func (m *Listener) GetUdpListenerConfig() *listener.UdpListenerConfig {
	if m != nil {
		return m.UdpListenerConfig
	}
	return nil
}

func (m *Listener) GetApiListener() *v2.ApiListener {
	if m != nil {
		return m.ApiListener
	}
	return nil
}

func (m *Listener) GetConnectionBalanceConfig() *Listener_ConnectionBalanceConfig {
	if m != nil {
		return m.ConnectionBalanceConfig
	}
	return nil
}

type Listener_DeprecatedV1 struct {
	BindToPort           *wrappers.BoolValue `protobuf:"bytes,1,opt,name=bind_to_port,json=bindToPort,proto3" json:"bind_to_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Listener_DeprecatedV1) Reset()         { *m = Listener_DeprecatedV1{} }
func (m *Listener_DeprecatedV1) String() string { return proto.CompactTextString(m) }
func (*Listener_DeprecatedV1) ProtoMessage()    {}
func (*Listener_DeprecatedV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_34e2cd84a105bcd1, []int{0, 0}
}

func (m *Listener_DeprecatedV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Listener_DeprecatedV1.Unmarshal(m, b)
}
func (m *Listener_DeprecatedV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Listener_DeprecatedV1.Marshal(b, m, deterministic)
}
func (m *Listener_DeprecatedV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Listener_DeprecatedV1.Merge(m, src)
}
func (m *Listener_DeprecatedV1) XXX_Size() int {
	return xxx_messageInfo_Listener_DeprecatedV1.Size(m)
}
func (m *Listener_DeprecatedV1) XXX_DiscardUnknown() {
	xxx_messageInfo_Listener_DeprecatedV1.DiscardUnknown(m)
}

var xxx_messageInfo_Listener_DeprecatedV1 proto.InternalMessageInfo

func (m *Listener_DeprecatedV1) GetBindToPort() *wrappers.BoolValue {
	if m != nil {
		return m.BindToPort
	}
	return nil
}

type Listener_ConnectionBalanceConfig struct {
	// Types that are valid to be assigned to BalanceType:
	//	*Listener_ConnectionBalanceConfig_ExactBalance_
	BalanceType          isListener_ConnectionBalanceConfig_BalanceType `protobuf_oneof:"balance_type"`
	XXX_NoUnkeyedLiteral struct{}                                       `json:"-"`
	XXX_unrecognized     []byte                                         `json:"-"`
	XXX_sizecache        int32                                          `json:"-"`
}

func (m *Listener_ConnectionBalanceConfig) Reset()         { *m = Listener_ConnectionBalanceConfig{} }
func (m *Listener_ConnectionBalanceConfig) String() string { return proto.CompactTextString(m) }
func (*Listener_ConnectionBalanceConfig) ProtoMessage()    {}
func (*Listener_ConnectionBalanceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_34e2cd84a105bcd1, []int{0, 1}
}

func (m *Listener_ConnectionBalanceConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Listener_ConnectionBalanceConfig.Unmarshal(m, b)
}
func (m *Listener_ConnectionBalanceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Listener_ConnectionBalanceConfig.Marshal(b, m, deterministic)
}
func (m *Listener_ConnectionBalanceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Listener_ConnectionBalanceConfig.Merge(m, src)
}
func (m *Listener_ConnectionBalanceConfig) XXX_Size() int {
	return xxx_messageInfo_Listener_ConnectionBalanceConfig.Size(m)
}
func (m *Listener_ConnectionBalanceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_Listener_ConnectionBalanceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_Listener_ConnectionBalanceConfig proto.InternalMessageInfo

type isListener_ConnectionBalanceConfig_BalanceType interface {
	isListener_ConnectionBalanceConfig_BalanceType()
}

type Listener_ConnectionBalanceConfig_ExactBalance_ struct {
	ExactBalance *Listener_ConnectionBalanceConfig_ExactBalance `protobuf:"bytes,1,opt,name=exact_balance,json=exactBalance,proto3,oneof"`
}

func (*Listener_ConnectionBalanceConfig_ExactBalance_) isListener_ConnectionBalanceConfig_BalanceType() {
}

func (m *Listener_ConnectionBalanceConfig) GetBalanceType() isListener_ConnectionBalanceConfig_BalanceType {
	if m != nil {
		return m.BalanceType
	}
	return nil
}

func (m *Listener_ConnectionBalanceConfig) GetExactBalance() *Listener_ConnectionBalanceConfig_ExactBalance {
	if x, ok := m.GetBalanceType().(*Listener_ConnectionBalanceConfig_ExactBalance_); ok {
		return x.ExactBalance
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Listener_ConnectionBalanceConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Listener_ConnectionBalanceConfig_ExactBalance_)(nil),
	}
}

type Listener_ConnectionBalanceConfig_ExactBalance struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Listener_ConnectionBalanceConfig_ExactBalance) Reset() {
	*m = Listener_ConnectionBalanceConfig_ExactBalance{}
}
func (m *Listener_ConnectionBalanceConfig_ExactBalance) String() string {
	return proto.CompactTextString(m)
}
func (*Listener_ConnectionBalanceConfig_ExactBalance) ProtoMessage() {}
func (*Listener_ConnectionBalanceConfig_ExactBalance) Descriptor() ([]byte, []int) {
	return fileDescriptor_34e2cd84a105bcd1, []int{0, 1, 0}
}

func (m *Listener_ConnectionBalanceConfig_ExactBalance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Listener_ConnectionBalanceConfig_ExactBalance.Unmarshal(m, b)
}
func (m *Listener_ConnectionBalanceConfig_ExactBalance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Listener_ConnectionBalanceConfig_ExactBalance.Marshal(b, m, deterministic)
}
func (m *Listener_ConnectionBalanceConfig_ExactBalance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Listener_ConnectionBalanceConfig_ExactBalance.Merge(m, src)
}
func (m *Listener_ConnectionBalanceConfig_ExactBalance) XXX_Size() int {
	return xxx_messageInfo_Listener_ConnectionBalanceConfig_ExactBalance.Size(m)
}
func (m *Listener_ConnectionBalanceConfig_ExactBalance) XXX_DiscardUnknown() {
	xxx_messageInfo_Listener_ConnectionBalanceConfig_ExactBalance.DiscardUnknown(m)
}

var xxx_messageInfo_Listener_ConnectionBalanceConfig_ExactBalance proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("envoy.api.v2.Listener_DrainType", Listener_DrainType_name, Listener_DrainType_value)
	proto.RegisterType((*Listener)(nil), "envoy.api.v2.Listener")
	proto.RegisterType((*Listener_DeprecatedV1)(nil), "envoy.api.v2.Listener.DeprecatedV1")
	proto.RegisterType((*Listener_ConnectionBalanceConfig)(nil), "envoy.api.v2.Listener.ConnectionBalanceConfig")
	proto.RegisterType((*Listener_ConnectionBalanceConfig_ExactBalance)(nil), "envoy.api.v2.Listener.ConnectionBalanceConfig.ExactBalance")
}

func init() { proto.RegisterFile("envoy/api/v2/lds.proto", fileDescriptor_34e2cd84a105bcd1) }

var fileDescriptor_34e2cd84a105bcd1 = []byte{
	// 1037 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0xee, 0xa6, 0x3f, 0x71, 0xc6, 0x3f, 0x71, 0x26, 0x28, 0xd9, 0x9a, 0xd0, 0x18, 0x37, 0x48,
	0x86, 0x4a, 0x6b, 0xea, 0x4a, 0x20, 0x95, 0x0a, 0x14, 0xc7, 0x35, 0x29, 0x72, 0xeb, 0xb0, 0x49,
	0xaa, 0xf4, 0x6a, 0x35, 0xde, 0x3d, 0x9b, 0x0c, 0x6c, 0x66, 0xb6, 0x33, 0xb3, 0xa6, 0xbe, 0xe5,
	0x0a, 0x71, 0xcb, 0x1b, 0xf0, 0x08, 0x3c, 0x08, 0x37, 0xbc, 0x02, 0x0f, 0x81, 0xb8, 0x42, 0x3b,
	0xbb, 0xe3, 0xd8, 0x8e, 0xdd, 0x82, 0xc4, 0xdd, 0x9c, 0x39, 0xdf, 0xf7, 0xf9, 0xcc, 0x39, 0xdf,
	0x9e, 0x04, 0x6d, 0x01, 0x1b, 0xf1, 0x71, 0x8b, 0xc4, 0xb4, 0x35, 0x6a, 0xb7, 0xa2, 0x40, 0x3a,
	0xb1, 0xe0, 0x8a, 0xe3, 0x92, 0xbe, 0x77, 0x48, 0x4c, 0x9d, 0x51, 0xbb, 0xb6, 0x3b, 0x83, 0xf2,
	0xb9, 0x80, 0x16, 0x09, 0x02, 0x01, 0x32, 0x87, 0xd7, 0x76, 0xae, 0x03, 0x86, 0x44, 0xc2, 0xc2,
	0x6c, 0x40, 0xa5, 0xcf, 0x47, 0x20, 0xc6, 0x79, 0x76, 0x6f, 0xb6, 0x04, 0x2a, 0x15, 0x30, 0x10,
	0x93, 0x43, 0x8e, 0x6a, 0x2d, 0x46, 0x25, 0x41, 0xec, 0x99, 0xc0, 0xf3, 0x39, 0x0b, 0xe9, 0x79,
	0x4e, 0x78, 0x90, 0x11, 0xb2, 0xbb, 0x2b, 0xc2, 0xa8, 0x9d, 0x6a, 0x78, 0x73, 0xea, 0x3b, 0xe7,
	0x9c, 0x9f, 0x47, 0xa0, 0xe5, 0x09, 0x63, 0x5c, 0x11, 0x45, 0x39, 0x33, 0xaf, 0xbb, 0x97, 0x67,
	0x75, 0x34, 0x4c, 0xc2, 0x56, 0x90, 0x08, 0x0d, 0x58, 0x96, 0xff, 0x41, 0x90, 0x38, 0x06, 0x61,
	0xf8, 0xdb, 0x23, 0x12, 0xd1, 0x80, 0x28, 0x68, 0x99, 0x43, 0x96, 0x68, 0xfc, 0x56, 0x46, 0x85,
	0x7e, 0x5e, 0x09, 0xc6, 0xe8, 0x16, 0x23, 0x97, 0x60, 0x5b, 0x75, 0xab, 0xb9, 0xe6, 0xea, 0x33,
	0xfe, 0x12, 0xad, 0xe6, 0x8d, 0xb6, 0x57, 0xea, 0x56, 0xb3, 0xd8, 0xae, 0x39, 0xd3, 0x83, 0x71,
	0xd2, 0x4e, 0x3b, 0xfb, 0x19, 0xa2, 0x53, 0xf8, 0xbb, 0x73, 0xfb, 0x67, 0x6b, 0xa5, 0x6a, 0xb9,
	0x86, 0x84, 0xbf, 0x46, 0xe5, 0x90, 0x46, 0x2a, 0xed, 0xcd, 0x05, 0xa1, 0x4c, 0xda, 0x37, 0xeb,
	0x37, 0x9b, 0xc5, 0x76, 0x63, 0x56, 0x65, 0xd2, 0x8c, 0x9e, 0xc6, 0x1e, 0xa4, 0x50, 0xb7, 0x14,
	0x5e, 0x05, 0x12, 0x1f, 0xa2, 0x6a, 0x22, 0xc1, 0xe3, 0x82, 0x9e, 0x53, 0x46, 0x22, 0x2f, 0x90,
	0xca, 0xbe, 0x95, 0x57, 0x94, 0xbd, 0xde, 0x31, 0xaf, 0x77, 0x3a, 0x9c, 0x47, 0x2f, 0x49, 0x94,
	0x40, 0x67, 0xc5, 0xb6, 0xdc, 0x4a, 0x22, 0x61, 0x90, 0xd3, 0xba, 0x52, 0xe1, 0x10, 0x7d, 0x18,
	0x67, 0xb3, 0x62, 0xe0, 0xa7, 0x4d, 0xf4, 0x86, 0x49, 0x18, 0x82, 0xf0, 0x22, 0x7a, 0x49, 0x95,
	0x37, 0x1c, 0x2b, 0x90, 0xf6, 0x6d, 0x2d, 0xbd, 0x73, 0x4d, 0xfa, 0xf4, 0x19, 0x53, 0x8f, 0xda,
	0x5a, 0xdc, 0xfd, 0x20, 0x06, 0x71, 0x30, 0x51, 0xe9, 0x68, 0x91, 0x7e, 0xaa, 0xd1, 0x49, 0x25,
	0xf0, 0xe7, 0xa8, 0x70, 0x09, 0x8a, 0x04, 0x44, 0x11, 0xfb, 0x8e, 0x96, 0x7b, 0x7f, 0x41, 0xef,
	0x9e, 0xe7, 0x10, 0x77, 0x02, 0xc6, 0x87, 0xa8, 0x1c, 0x40, 0x2c, 0xc0, 0x27, 0x0a, 0x02, 0x6f,
	0xf4, 0xd0, 0x5e, 0xd5, 0xec, 0xfb, 0xb3, 0x6c, 0x33, 0x36, 0xa7, 0x3b, 0xc1, 0xbe, 0x7c, 0xe8,
	0x96, 0x82, 0xa9, 0x08, 0x7f, 0x85, 0x50, 0x20, 0x08, 0x65, 0x9e, 0x1a, 0xc7, 0x60, 0x17, 0xea,
	0x56, 0xb3, 0xd2, 0xae, 0x2f, 0x93, 0x49, 0x81, 0x27, 0xe3, 0x18, 0xdc, 0xb5, 0xc0, 0x1c, 0xf1,
	0x11, 0xaa, 0x4e, 0xcc, 0x9d, 0x8d, 0x43, 0xda, 0x6b, 0x7a, 0x82, 0x1f, 0x2d, 0x99, 0xa0, 0xd1,
	0xcb, 0x26, 0xe9, 0xae, 0x47, 0x33, 0xb1, 0xc4, 0xc7, 0xc8, 0x9e, 0x57, 0xf4, 0x14, 0xbd, 0x04,
	0x9e, 0x28, 0x7b, 0x5d, 0xbf, 0xf3, 0xee, 0xb5, 0xa6, 0x77, 0x73, 0xb7, 0xbb, 0x5b, 0x73, 0x6a,
	0x27, 0x19, 0x11, 0xbf, 0x40, 0x7b, 0x3e, 0x67, 0x8a, 0xb2, 0x04, 0x3c, 0xce, 0xbc, 0xa5, 0x3f,
	0xb0, 0x51, 0xb7, 0x9a, 0x05, 0xb7, 0x6e, 0xb0, 0x03, 0xd6, 0x5f, 0xac, 0xf7, 0x04, 0x15, 0x95,
	0x20, 0x4c, 0xc6, 0x44, 0x00, 0x53, 0x36, 0x7a, 0x97, 0xcf, 0xdc, 0x69, 0x38, 0xfe, 0x0c, 0x15,
	0x42, 0x01, 0x30, 0xa4, 0x2c, 0xb0, 0x8b, 0xef, 0xa4, 0x4e, 0xb0, 0xb8, 0x87, 0x2a, 0x92, 0xfb,
	0xdf, 0x83, 0xf2, 0x78, 0xac, 0xbf, 0x7e, 0xbb, 0xac, 0x5b, 0xbd, 0xbb, 0xc0, 0x36, 0xc7, 0x1a,
	0x38, 0xd0, 0x38, 0xb7, 0x2c, 0xa7, 0x22, 0x89, 0xcf, 0x50, 0x4d, 0xf9, 0xb1, 0x17, 0x12, 0x99,
	0x2a, 0x01, 0xf3, 0x5e, 0x27, 0x90, 0x80, 0x17, 0x01, 0x3b, 0x57, 0x17, 0x76, 0xe9, 0x5f, 0x38,
	0x7b, 0x4b, 0xf9, 0x71, 0x8f, 0x48, 0x35, 0x88, 0x81, 0x7d, 0x9b, 0x92, 0xfb, 0x9a, 0x8b, 0x8f,
	0xd0, 0x86, 0x12, 0x24, 0x0c, 0xa9, 0xef, 0x05, 0x54, 0x64, 0xbe, 0xb7, 0xab, 0xda, 0x56, 0xf7,
	0x17, 0x14, 0x79, 0x92, 0x61, 0xbb, 0x06, 0xea, 0x56, 0xd5, 0xdc, 0x0d, 0x3e, 0x43, 0x9b, 0x0b,
	0x36, 0xa8, 0x8d, 0x75, 0x91, 0xcd, 0x25, 0x1e, 0x3b, 0x0d, 0x62, 0x33, 0xb8, 0x03, 0x8d, 0x77,
	0x37, 0x92, 0xf9, 0x2b, 0x7c, 0x88, 0x4a, 0xd3, 0x7b, 0xd6, 0xde, 0xd4, 0x92, 0xc6, 0xb6, 0xf9,
	0xa6, 0x9e, 0x48, 0x8e, 0xda, 0xce, 0x7e, 0x4c, 0x8d, 0x84, 0x5b, 0x24, 0x57, 0x01, 0xfe, 0x0e,
	0xdd, 0x9d, 0x5e, 0x16, 0x24, 0x22, 0xcc, 0x07, 0x53, 0xe9, 0x7b, 0x5a, 0xd6, 0x59, 0xf2, 0x51,
	0x4d, 0xad, 0x87, 0x8c, 0x96, 0xd7, 0xbb, 0xed, 0x2f, 0x4e, 0xd4, 0xfa, 0xa8, 0x34, 0xfd, 0x3d,
	0xe3, 0x27, 0xa8, 0x94, 0x7a, 0xc3, 0x53, 0xdc, 0x8b, 0xb9, 0x50, 0x7a, 0x37, 0xbf, 0xdd, 0x4f,
	0x28, 0xc5, 0x9f, 0xf0, 0x23, 0x2e, 0x54, 0xed, 0x57, 0x0b, 0x6d, 0x2f, 0x29, 0x01, 0x0f, 0x51,
	0x19, 0xde, 0x10, 0x5f, 0x99, 0x07, 0xe5, 0xd2, 0x5f, 0xfc, 0xb7, 0x97, 0x38, 0x4f, 0x53, 0x8d,
	0xfc, 0xea, 0xf0, 0x86, 0x5b, 0x82, 0xa9, 0xb8, 0x56, 0x41, 0xa5, 0xe9, 0x7c, 0x67, 0x13, 0x95,
	0x4c, 0xfb, 0xd2, 0x8d, 0x84, 0x6f, 0xfe, 0xd5, 0xb1, 0x1a, 0x1f, 0xa3, 0xb5, 0xc9, 0xee, 0xc1,
	0x45, 0xb4, 0xda, 0x7d, 0xda, 0xdb, 0x3f, 0xed, 0x9f, 0x54, 0x6f, 0xe0, 0x75, 0x54, 0x7c, 0x3e,
	0xe8, 0x3e, 0xeb, 0xbd, 0xf2, 0x06, 0x2f, 0xfa, 0xaf, 0xaa, 0xd6, 0x37, 0xb7, 0x0a, 0x95, 0xea,
	0x7a, 0xfb, 0xf7, 0x15, 0x64, 0x9b, 0xba, 0xba, 0xe6, 0x6f, 0xf9, 0x31, 0x88, 0x11, 0xf5, 0x01,
	0x13, 0x54, 0xe9, 0x42, 0xa4, 0x88, 0x01, 0x48, 0x3c, 0xe7, 0x4c, 0x9d, 0x9d, 0xd0, 0x5c, 0x78,
	0x9d, 0x80, 0x54, 0xb5, 0xbd, 0xb7, 0x83, 0x64, 0xcc, 0x99, 0x84, 0xc6, 0x8d, 0xa6, 0xf5, 0xa9,
	0x85, 0xcf, 0xd0, 0xfa, 0xb1, 0x12, 0x40, 0x2e, 0xaf, 0x7e, 0xe3, 0xde, 0x1c, 0x7d, 0x5e, 0x7e,
	0x77, 0x69, 0x7e, 0x46, 0x39, 0x41, 0x95, 0x1e, 0x28, 0xff, 0xe2, 0x7f, 0x14, 0x6e, 0xfc, 0xf8,
	0xc7, 0x9f, 0xbf, 0xac, 0xec, 0x34, 0xb6, 0x67, 0xfe, 0xf3, 0x79, 0x6c, 0x8c, 0x2f, 0x1f, 0x5b,
	0x9f, 0x74, 0x1e, 0xa0, 0x1a, 0xe5, 0x99, 0x50, 0x2c, 0xf8, 0x9b, 0xf1, 0x8c, 0x66, 0xa7, 0xd0,
	0x0f, 0xe4, 0x51, 0xea, 0xb3, 0x23, 0xeb, 0x27, 0xcb, 0x1a, 0xde, 0xd1, 0x9e, 0x7b, 0xf4, 0x4f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x33, 0x61, 0x78, 0x40, 0xb9, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ListenerDiscoveryServiceClient is the client API for ListenerDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ListenerDiscoveryServiceClient interface {
	DeltaListeners(ctx context.Context, opts ...grpc.CallOption) (ListenerDiscoveryService_DeltaListenersClient, error)
	StreamListeners(ctx context.Context, opts ...grpc.CallOption) (ListenerDiscoveryService_StreamListenersClient, error)
	FetchListeners(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error)
}

type listenerDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewListenerDiscoveryServiceClient(cc *grpc.ClientConn) ListenerDiscoveryServiceClient {
	return &listenerDiscoveryServiceClient{cc}
}

func (c *listenerDiscoveryServiceClient) DeltaListeners(ctx context.Context, opts ...grpc.CallOption) (ListenerDiscoveryService_DeltaListenersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ListenerDiscoveryService_serviceDesc.Streams[0], "/envoy.api.v2.ListenerDiscoveryService/DeltaListeners", opts...)
	if err != nil {
		return nil, err
	}
	x := &listenerDiscoveryServiceDeltaListenersClient{stream}
	return x, nil
}

type ListenerDiscoveryService_DeltaListenersClient interface {
	Send(*DeltaDiscoveryRequest) error
	Recv() (*DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type listenerDiscoveryServiceDeltaListenersClient struct {
	grpc.ClientStream
}

func (x *listenerDiscoveryServiceDeltaListenersClient) Send(m *DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *listenerDiscoveryServiceDeltaListenersClient) Recv() (*DeltaDiscoveryResponse, error) {
	m := new(DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *listenerDiscoveryServiceClient) StreamListeners(ctx context.Context, opts ...grpc.CallOption) (ListenerDiscoveryService_StreamListenersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ListenerDiscoveryService_serviceDesc.Streams[1], "/envoy.api.v2.ListenerDiscoveryService/StreamListeners", opts...)
	if err != nil {
		return nil, err
	}
	x := &listenerDiscoveryServiceStreamListenersClient{stream}
	return x, nil
}

type ListenerDiscoveryService_StreamListenersClient interface {
	Send(*DiscoveryRequest) error
	Recv() (*DiscoveryResponse, error)
	grpc.ClientStream
}

type listenerDiscoveryServiceStreamListenersClient struct {
	grpc.ClientStream
}

func (x *listenerDiscoveryServiceStreamListenersClient) Send(m *DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *listenerDiscoveryServiceStreamListenersClient) Recv() (*DiscoveryResponse, error) {
	m := new(DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *listenerDiscoveryServiceClient) FetchListeners(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error) {
	out := new(DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.api.v2.ListenerDiscoveryService/FetchListeners", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListenerDiscoveryServiceServer is the server API for ListenerDiscoveryService service.
type ListenerDiscoveryServiceServer interface {
	DeltaListeners(ListenerDiscoveryService_DeltaListenersServer) error
	StreamListeners(ListenerDiscoveryService_StreamListenersServer) error
	FetchListeners(context.Context, *DiscoveryRequest) (*DiscoveryResponse, error)
}

func RegisterListenerDiscoveryServiceServer(s *grpc.Server, srv ListenerDiscoveryServiceServer) {
	s.RegisterService(&_ListenerDiscoveryService_serviceDesc, srv)
}

func _ListenerDiscoveryService_DeltaListeners_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ListenerDiscoveryServiceServer).DeltaListeners(&listenerDiscoveryServiceDeltaListenersServer{stream})
}

type ListenerDiscoveryService_DeltaListenersServer interface {
	Send(*DeltaDiscoveryResponse) error
	Recv() (*DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type listenerDiscoveryServiceDeltaListenersServer struct {
	grpc.ServerStream
}

func (x *listenerDiscoveryServiceDeltaListenersServer) Send(m *DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *listenerDiscoveryServiceDeltaListenersServer) Recv() (*DeltaDiscoveryRequest, error) {
	m := new(DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ListenerDiscoveryService_StreamListeners_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ListenerDiscoveryServiceServer).StreamListeners(&listenerDiscoveryServiceStreamListenersServer{stream})
}

type ListenerDiscoveryService_StreamListenersServer interface {
	Send(*DiscoveryResponse) error
	Recv() (*DiscoveryRequest, error)
	grpc.ServerStream
}

type listenerDiscoveryServiceStreamListenersServer struct {
	grpc.ServerStream
}

func (x *listenerDiscoveryServiceStreamListenersServer) Send(m *DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *listenerDiscoveryServiceStreamListenersServer) Recv() (*DiscoveryRequest, error) {
	m := new(DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ListenerDiscoveryService_FetchListeners_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListenerDiscoveryServiceServer).FetchListeners(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.api.v2.ListenerDiscoveryService/FetchListeners",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListenerDiscoveryServiceServer).FetchListeners(ctx, req.(*DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ListenerDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.api.v2.ListenerDiscoveryService",
	HandlerType: (*ListenerDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchListeners",
			Handler:    _ListenerDiscoveryService_FetchListeners_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DeltaListeners",
			Handler:       _ListenerDiscoveryService_DeltaListeners_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamListeners",
			Handler:       _ListenerDiscoveryService_StreamListeners_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/api/v2/lds.proto",
}
