// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/thrift_proxy/v3alpha/route.proto

package envoy_config_filter_network_thrift_proxy_v3alpha

import (
	fmt "fmt"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v3alpha/core"
	route "github.com/envoyproxy/go-control-plane/envoy/api/v3alpha/route"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type RouteConfiguration struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Routes               []*Route `protobuf:"bytes,2,rep,name=routes,proto3" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteConfiguration) Reset()         { *m = RouteConfiguration{} }
func (m *RouteConfiguration) String() string { return proto.CompactTextString(m) }
func (*RouteConfiguration) ProtoMessage()    {}
func (*RouteConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b3b4ebd82e248c, []int{0}
}

func (m *RouteConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteConfiguration.Unmarshal(m, b)
}
func (m *RouteConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteConfiguration.Marshal(b, m, deterministic)
}
func (m *RouteConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteConfiguration.Merge(m, src)
}
func (m *RouteConfiguration) XXX_Size() int {
	return xxx_messageInfo_RouteConfiguration.Size(m)
}
func (m *RouteConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_RouteConfiguration proto.InternalMessageInfo

func (m *RouteConfiguration) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RouteConfiguration) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

type Route struct {
	Match                *RouteMatch  `protobuf:"bytes,1,opt,name=match,proto3" json:"match,omitempty"`
	Route                *RouteAction `protobuf:"bytes,2,opt,name=route,proto3" json:"route,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Route) Reset()         { *m = Route{} }
func (m *Route) String() string { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()    {}
func (*Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b3b4ebd82e248c, []int{1}
}

func (m *Route) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Route.Unmarshal(m, b)
}
func (m *Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Route.Marshal(b, m, deterministic)
}
func (m *Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route.Merge(m, src)
}
func (m *Route) XXX_Size() int {
	return xxx_messageInfo_Route.Size(m)
}
func (m *Route) XXX_DiscardUnknown() {
	xxx_messageInfo_Route.DiscardUnknown(m)
}

var xxx_messageInfo_Route proto.InternalMessageInfo

func (m *Route) GetMatch() *RouteMatch {
	if m != nil {
		return m.Match
	}
	return nil
}

func (m *Route) GetRoute() *RouteAction {
	if m != nil {
		return m.Route
	}
	return nil
}

type RouteMatch struct {
	// Types that are valid to be assigned to MatchSpecifier:
	//	*RouteMatch_MethodName
	//	*RouteMatch_ServiceName
	MatchSpecifier       isRouteMatch_MatchSpecifier `protobuf_oneof:"match_specifier"`
	Invert               bool                        `protobuf:"varint,3,opt,name=invert,proto3" json:"invert,omitempty"`
	Headers              []*route.HeaderMatcher      `protobuf:"bytes,4,rep,name=headers,proto3" json:"headers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *RouteMatch) Reset()         { *m = RouteMatch{} }
func (m *RouteMatch) String() string { return proto.CompactTextString(m) }
func (*RouteMatch) ProtoMessage()    {}
func (*RouteMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b3b4ebd82e248c, []int{2}
}

func (m *RouteMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteMatch.Unmarshal(m, b)
}
func (m *RouteMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteMatch.Marshal(b, m, deterministic)
}
func (m *RouteMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteMatch.Merge(m, src)
}
func (m *RouteMatch) XXX_Size() int {
	return xxx_messageInfo_RouteMatch.Size(m)
}
func (m *RouteMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteMatch.DiscardUnknown(m)
}

var xxx_messageInfo_RouteMatch proto.InternalMessageInfo

type isRouteMatch_MatchSpecifier interface {
	isRouteMatch_MatchSpecifier()
}

type RouteMatch_MethodName struct {
	MethodName string `protobuf:"bytes,1,opt,name=method_name,json=methodName,proto3,oneof"`
}

type RouteMatch_ServiceName struct {
	ServiceName string `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3,oneof"`
}

func (*RouteMatch_MethodName) isRouteMatch_MatchSpecifier() {}

func (*RouteMatch_ServiceName) isRouteMatch_MatchSpecifier() {}

func (m *RouteMatch) GetMatchSpecifier() isRouteMatch_MatchSpecifier {
	if m != nil {
		return m.MatchSpecifier
	}
	return nil
}

func (m *RouteMatch) GetMethodName() string {
	if x, ok := m.GetMatchSpecifier().(*RouteMatch_MethodName); ok {
		return x.MethodName
	}
	return ""
}

func (m *RouteMatch) GetServiceName() string {
	if x, ok := m.GetMatchSpecifier().(*RouteMatch_ServiceName); ok {
		return x.ServiceName
	}
	return ""
}

func (m *RouteMatch) GetInvert() bool {
	if m != nil {
		return m.Invert
	}
	return false
}

func (m *RouteMatch) GetHeaders() []*route.HeaderMatcher {
	if m != nil {
		return m.Headers
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RouteMatch) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RouteMatch_MethodName)(nil),
		(*RouteMatch_ServiceName)(nil),
	}
}

type RouteAction struct {
	// Types that are valid to be assigned to ClusterSpecifier:
	//	*RouteAction_Cluster
	//	*RouteAction_WeightedClusters
	//	*RouteAction_ClusterHeader
	ClusterSpecifier     isRouteAction_ClusterSpecifier `protobuf_oneof:"cluster_specifier"`
	MetadataMatch        *core.Metadata                 `protobuf:"bytes,3,opt,name=metadata_match,json=metadataMatch,proto3" json:"metadata_match,omitempty"`
	RateLimits           []*route.RateLimit             `protobuf:"bytes,4,rep,name=rate_limits,json=rateLimits,proto3" json:"rate_limits,omitempty"`
	StripServiceName     bool                           `protobuf:"varint,5,opt,name=strip_service_name,json=stripServiceName,proto3" json:"strip_service_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *RouteAction) Reset()         { *m = RouteAction{} }
func (m *RouteAction) String() string { return proto.CompactTextString(m) }
func (*RouteAction) ProtoMessage()    {}
func (*RouteAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b3b4ebd82e248c, []int{3}
}

func (m *RouteAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteAction.Unmarshal(m, b)
}
func (m *RouteAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteAction.Marshal(b, m, deterministic)
}
func (m *RouteAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteAction.Merge(m, src)
}
func (m *RouteAction) XXX_Size() int {
	return xxx_messageInfo_RouteAction.Size(m)
}
func (m *RouteAction) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteAction.DiscardUnknown(m)
}

var xxx_messageInfo_RouteAction proto.InternalMessageInfo

type isRouteAction_ClusterSpecifier interface {
	isRouteAction_ClusterSpecifier()
}

type RouteAction_Cluster struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster,proto3,oneof"`
}

type RouteAction_WeightedClusters struct {
	WeightedClusters *WeightedCluster `protobuf:"bytes,2,opt,name=weighted_clusters,json=weightedClusters,proto3,oneof"`
}

type RouteAction_ClusterHeader struct {
	ClusterHeader string `protobuf:"bytes,6,opt,name=cluster_header,json=clusterHeader,proto3,oneof"`
}

func (*RouteAction_Cluster) isRouteAction_ClusterSpecifier() {}

func (*RouteAction_WeightedClusters) isRouteAction_ClusterSpecifier() {}

func (*RouteAction_ClusterHeader) isRouteAction_ClusterSpecifier() {}

func (m *RouteAction) GetClusterSpecifier() isRouteAction_ClusterSpecifier {
	if m != nil {
		return m.ClusterSpecifier
	}
	return nil
}

func (m *RouteAction) GetCluster() string {
	if x, ok := m.GetClusterSpecifier().(*RouteAction_Cluster); ok {
		return x.Cluster
	}
	return ""
}

func (m *RouteAction) GetWeightedClusters() *WeightedCluster {
	if x, ok := m.GetClusterSpecifier().(*RouteAction_WeightedClusters); ok {
		return x.WeightedClusters
	}
	return nil
}

func (m *RouteAction) GetClusterHeader() string {
	if x, ok := m.GetClusterSpecifier().(*RouteAction_ClusterHeader); ok {
		return x.ClusterHeader
	}
	return ""
}

func (m *RouteAction) GetMetadataMatch() *core.Metadata {
	if m != nil {
		return m.MetadataMatch
	}
	return nil
}

func (m *RouteAction) GetRateLimits() []*route.RateLimit {
	if m != nil {
		return m.RateLimits
	}
	return nil
}

func (m *RouteAction) GetStripServiceName() bool {
	if m != nil {
		return m.StripServiceName
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RouteAction) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RouteAction_Cluster)(nil),
		(*RouteAction_WeightedClusters)(nil),
		(*RouteAction_ClusterHeader)(nil),
	}
}

type WeightedCluster struct {
	Clusters             []*WeightedCluster_ClusterWeight `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *WeightedCluster) Reset()         { *m = WeightedCluster{} }
func (m *WeightedCluster) String() string { return proto.CompactTextString(m) }
func (*WeightedCluster) ProtoMessage()    {}
func (*WeightedCluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b3b4ebd82e248c, []int{4}
}

func (m *WeightedCluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WeightedCluster.Unmarshal(m, b)
}
func (m *WeightedCluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WeightedCluster.Marshal(b, m, deterministic)
}
func (m *WeightedCluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeightedCluster.Merge(m, src)
}
func (m *WeightedCluster) XXX_Size() int {
	return xxx_messageInfo_WeightedCluster.Size(m)
}
func (m *WeightedCluster) XXX_DiscardUnknown() {
	xxx_messageInfo_WeightedCluster.DiscardUnknown(m)
}

var xxx_messageInfo_WeightedCluster proto.InternalMessageInfo

func (m *WeightedCluster) GetClusters() []*WeightedCluster_ClusterWeight {
	if m != nil {
		return m.Clusters
	}
	return nil
}

type WeightedCluster_ClusterWeight struct {
	Name                 string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Weight               *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=weight,proto3" json:"weight,omitempty"`
	MetadataMatch        *core.Metadata        `protobuf:"bytes,3,opt,name=metadata_match,json=metadataMatch,proto3" json:"metadata_match,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *WeightedCluster_ClusterWeight) Reset()         { *m = WeightedCluster_ClusterWeight{} }
func (m *WeightedCluster_ClusterWeight) String() string { return proto.CompactTextString(m) }
func (*WeightedCluster_ClusterWeight) ProtoMessage()    {}
func (*WeightedCluster_ClusterWeight) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b3b4ebd82e248c, []int{4, 0}
}

func (m *WeightedCluster_ClusterWeight) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WeightedCluster_ClusterWeight.Unmarshal(m, b)
}
func (m *WeightedCluster_ClusterWeight) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WeightedCluster_ClusterWeight.Marshal(b, m, deterministic)
}
func (m *WeightedCluster_ClusterWeight) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeightedCluster_ClusterWeight.Merge(m, src)
}
func (m *WeightedCluster_ClusterWeight) XXX_Size() int {
	return xxx_messageInfo_WeightedCluster_ClusterWeight.Size(m)
}
func (m *WeightedCluster_ClusterWeight) XXX_DiscardUnknown() {
	xxx_messageInfo_WeightedCluster_ClusterWeight.DiscardUnknown(m)
}

var xxx_messageInfo_WeightedCluster_ClusterWeight proto.InternalMessageInfo

func (m *WeightedCluster_ClusterWeight) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WeightedCluster_ClusterWeight) GetWeight() *wrappers.UInt32Value {
	if m != nil {
		return m.Weight
	}
	return nil
}

func (m *WeightedCluster_ClusterWeight) GetMetadataMatch() *core.Metadata {
	if m != nil {
		return m.MetadataMatch
	}
	return nil
}

func init() {
	proto.RegisterType((*RouteConfiguration)(nil), "envoy.config.filter.network.thrift_proxy.v3alpha.RouteConfiguration")
	proto.RegisterType((*Route)(nil), "envoy.config.filter.network.thrift_proxy.v3alpha.Route")
	proto.RegisterType((*RouteMatch)(nil), "envoy.config.filter.network.thrift_proxy.v3alpha.RouteMatch")
	proto.RegisterType((*RouteAction)(nil), "envoy.config.filter.network.thrift_proxy.v3alpha.RouteAction")
	proto.RegisterType((*WeightedCluster)(nil), "envoy.config.filter.network.thrift_proxy.v3alpha.WeightedCluster")
	proto.RegisterType((*WeightedCluster_ClusterWeight)(nil), "envoy.config.filter.network.thrift_proxy.v3alpha.WeightedCluster.ClusterWeight")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/thrift_proxy/v3alpha/route.proto", fileDescriptor_82b3b4ebd82e248c)
}

var fileDescriptor_82b3b4ebd82e248c = []byte{
	// 666 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4f, 0x6f, 0x13, 0x3f,
	0x10, 0xad, 0x37, 0x4d, 0x9a, 0xdf, 0xe4, 0xd7, 0x7f, 0x3e, 0x94, 0x55, 0x40, 0x28, 0x4d, 0x25,
	0x14, 0x21, 0xe4, 0xad, 0xda, 0x03, 0x97, 0x16, 0xd1, 0xed, 0x81, 0x20, 0xd1, 0x3f, 0x32, 0x02,
	0x2e, 0xa0, 0x95, 0xbb, 0x71, 0x12, 0x8b, 0xcd, 0x7a, 0xe5, 0x75, 0x12, 0xfa, 0x15, 0x38, 0xf2,
	0x39, 0x38, 0x72, 0xe7, 0x82, 0xf8, 0x4c, 0xa8, 0x27, 0xb4, 0xb6, 0x37, 0x4d, 0x0b, 0x1c, 0x5a,
	0x38, 0xc5, 0x19, 0xcf, 0x7b, 0x6f, 0xe6, 0xcd, 0xac, 0x61, 0x8f, 0xa7, 0x13, 0x79, 0x1e, 0xc4,
	0x32, 0xed, 0x8b, 0x41, 0xd0, 0x17, 0x89, 0xe6, 0x2a, 0x48, 0xb9, 0x9e, 0x4a, 0xf5, 0x3e, 0xd0,
	0x43, 0x25, 0xfa, 0x3a, 0xca, 0x94, 0xfc, 0x70, 0x1e, 0x4c, 0x76, 0x59, 0x92, 0x0d, 0x59, 0xa0,
	0xe4, 0x58, 0x73, 0x92, 0x29, 0xa9, 0x25, 0xde, 0x36, 0x68, 0x62, 0xd1, 0xc4, 0xa2, 0x89, 0x43,
	0x93, 0x79, 0x34, 0x71, 0xe8, 0xe6, 0xa6, 0xd5, 0x63, 0x99, 0x98, 0x11, 0xc6, 0x52, 0xf1, 0xe0,
	0x8c, 0xe5, 0x8e, 0xb4, 0xb9, 0xf5, 0x6b, 0x8a, 0xd1, 0x9c, 0x57, 0x6e, 0xde, 0x1f, 0x48, 0x39,
	0x48, 0x78, 0x60, 0xfe, 0x9d, 0x8d, 0xfb, 0xc1, 0x54, 0xb1, 0x2c, 0xe3, 0x2a, 0x77, 0xf7, 0x77,
	0x26, 0x2c, 0x11, 0x3d, 0xa6, 0x79, 0x50, 0x1e, 0xec, 0x45, 0xfb, 0x1c, 0x30, 0x2d, 0x78, 0x0e,
	0x4d, 0xcd, 0x63, 0xc5, 0xb4, 0x90, 0x29, 0xc6, 0xb0, 0x98, 0xb2, 0x11, 0xf7, 0x51, 0x0b, 0x75,
	0xfe, 0xa3, 0xe6, 0x8c, 0x4f, 0xa0, 0x66, 0x14, 0x73, 0xdf, 0x6b, 0x55, 0x3a, 0x8d, 0x9d, 0xc7,
	0xe4, 0xa6, 0xdd, 0x12, 0xa3, 0x44, 0x1d, 0x4d, 0xfb, 0x1b, 0x82, 0xaa, 0x89, 0xe0, 0xb7, 0x50,
	0x1d, 0x31, 0x1d, 0x0f, 0x8d, 0x5e, 0x63, 0x67, 0xef, 0x96, 0xcc, 0x47, 0x05, 0x47, 0x58, 0xbf,
	0x08, 0xab, 0x1f, 0x91, 0xb7, 0x86, 0xa8, 0x25, 0xc5, 0xef, 0xa0, 0x6a, 0x14, 0x7d, 0xcf, 0xb0,
	0xef, 0xdf, 0x92, 0xfd, 0x20, 0x2e, 0xac, 0x99, 0xa7, 0x37, 0xac, 0xed, 0xef, 0x08, 0xe0, 0x52,
	0x1e, 0x6f, 0x42, 0x63, 0xc4, 0xf5, 0x50, 0xf6, 0xa2, 0x4b, 0x07, 0xbb, 0x0b, 0x14, 0x6c, 0xf0,
	0xb8, 0x70, 0x72, 0x0b, 0xfe, 0xcf, 0xb9, 0x9a, 0x88, 0x98, 0xdb, 0x1c, 0xcf, 0xe5, 0x34, 0x5c,
	0xd4, 0x24, 0x6d, 0x40, 0x4d, 0xa4, 0x13, 0xae, 0xb4, 0x5f, 0x69, 0xa1, 0x4e, 0x9d, 0xba, 0x7f,
	0xf8, 0x29, 0x2c, 0x0d, 0x39, 0xeb, 0x71, 0x95, 0xfb, 0x8b, 0x66, 0x0e, 0x0f, 0x5c, 0x3f, 0x2c,
	0x13, 0xb3, 0x82, 0xed, 0x6a, 0x74, 0x4d, 0x9e, 0x29, 0x8b, 0x2b, 0x5a, 0xc2, 0xc2, 0x0d, 0x58,
	0x35, 0xc6, 0x44, 0x79, 0xc6, 0x63, 0xd1, 0x17, 0x5c, 0xe1, 0xca, 0x8f, 0x10, 0xb5, 0x3f, 0x57,
	0xa0, 0x31, 0xd7, 0x29, 0xde, 0x82, 0xa5, 0x38, 0x19, 0xe7, 0x9a, 0x2b, 0xdb, 0x45, 0xb8, 0x74,
	0x11, 0x2e, 0x2a, 0xaf, 0x85, 0xba, 0x0b, 0xb4, 0xbc, 0xc1, 0x19, 0xac, 0x4f, 0xb9, 0x18, 0x0c,
	0x35, 0xef, 0x45, 0x2e, 0x96, 0x3b, 0xa3, 0x0f, 0x6e, 0x6e, 0xf4, 0x1b, 0x47, 0x75, 0x68, 0x99,
	0xba, 0x0b, 0x74, 0x6d, 0x7a, 0x35, 0x94, 0xe3, 0x6d, 0x58, 0x71, 0x42, 0x91, 0xed, 0xc8, 0xaf,
	0x5d, 0xaf, 0x6e, 0xd9, 0x25, 0x58, 0x03, 0xf0, 0x33, 0x58, 0x19, 0x71, 0xcd, 0x7a, 0x4c, 0xb3,
	0xc8, 0xee, 0x59, 0xc5, 0x14, 0xd8, 0xfa, 0x8d, 0x73, 0xc5, 0xd7, 0x47, 0x8e, 0x5c, 0x36, 0x5d,
	0x2e, 0x71, 0x76, 0xb6, 0x87, 0xd0, 0x50, 0x4c, 0xf3, 0x28, 0x11, 0x23, 0xa1, 0x4b, 0xff, 0xdb,
	0x7f, 0xf4, 0x9f, 0x32, 0xcd, 0x5f, 0x14, 0xa9, 0x14, 0x54, 0x79, 0xcc, 0xf1, 0x23, 0xc0, 0xb9,
	0x56, 0x22, 0x8b, 0xae, 0xec, 0x40, 0xd5, 0x0c, 0x79, 0xcd, 0xdc, 0xbc, 0xbc, 0x5c, 0x83, 0xd0,
	0x87, 0xf5, 0xb2, 0xdb, 0x6b, 0xe3, 0xfa, 0xea, 0xc1, 0xea, 0x35, 0xbf, 0xf0, 0x18, 0xea, 0xb3,
	0x21, 0x20, 0x53, 0xdd, 0xc9, 0x5f, 0x0f, 0x81, 0xb8, 0x5f, 0x1b, 0x36, 0xfb, 0xff, 0x09, 0x79,
	0x75, 0x44, 0x67, 0x52, 0xcd, 0x2f, 0x08, 0x96, 0xaf, 0x64, 0xe1, 0xbb, 0xf3, 0x0f, 0xc8, 0x6c,
	0x34, 0xee, 0x25, 0xd9, 0x87, 0x9a, 0x9d, 0xaa, 0x5b, 0x94, 0x7b, 0xc4, 0xbe, 0x5e, 0xa4, 0x7c,
	0xbd, 0xc8, 0xab, 0xe7, 0xa9, 0xde, 0xdd, 0x79, 0xcd, 0x92, 0x31, 0x37, 0xe0, 0x87, 0x5e, 0x07,
	0x51, 0x07, 0xfa, 0x67, 0xe3, 0x0c, 0x8f, 0xe1, 0x89, 0x90, 0x16, 0x64, 0x1d, 0xb8, 0xa9, 0x55,
	0xa1, 0xfd, 0xf0, 0x4f, 0x8b, 0xb2, 0x4f, 0xd1, 0x59, 0xcd, 0xd4, 0xbf, 0xfb, 0x33, 0x00, 0x00,
	0xff, 0xff, 0x18, 0x8c, 0x3d, 0x8b, 0x45, 0x06, 0x00, 0x00,
}