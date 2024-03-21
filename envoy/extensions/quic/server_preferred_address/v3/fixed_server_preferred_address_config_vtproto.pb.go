//go:build vtprotobuf
// +build vtprotobuf
// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// source: envoy/extensions/quic/server_preferred_address/v3/fixed_server_preferred_address_config.proto

package server_preferred_addressv3

import (
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	bits "math/bits"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *FixedServerPreferredAddressConfig) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FixedServerPreferredAddressConfig) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *FixedServerPreferredAddressConfig) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if msg, ok := m.Ipv6Type.(*FixedServerPreferredAddressConfig_Ipv6Address); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if msg, ok := m.Ipv4Type.(*FixedServerPreferredAddressConfig_Ipv4Address); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	return len(dAtA) - i, nil
}

func (m *FixedServerPreferredAddressConfig_Ipv4Address) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *FixedServerPreferredAddressConfig_Ipv4Address) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.Ipv4Address)
	copy(dAtA[i:], m.Ipv4Address)
	i = encodeVarint(dAtA, i, uint64(len(m.Ipv4Address)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}
func (m *FixedServerPreferredAddressConfig_Ipv6Address) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *FixedServerPreferredAddressConfig_Ipv6Address) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.Ipv6Address)
	copy(dAtA[i:], m.Ipv6Address)
	i = encodeVarint(dAtA, i, uint64(len(m.Ipv6Address)))
	i--
	dAtA[i] = 0x12
	return len(dAtA) - i, nil
}
func encodeVarint(dAtA []byte, offset int, v uint64) int {
	offset -= sov(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FixedServerPreferredAddressConfig) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if vtmsg, ok := m.Ipv4Type.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	if vtmsg, ok := m.Ipv6Type.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	n += len(m.unknownFields)
	return n
}

func (m *FixedServerPreferredAddressConfig_Ipv4Address) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Ipv4Address)
	n += 1 + l + sov(uint64(l))
	return n
}
func (m *FixedServerPreferredAddressConfig_Ipv6Address) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Ipv6Address)
	n += 1 + l + sov(uint64(l))
	return n
}

func sov(x uint64) (n int) {
	return (bits.Len64(x|1) + 6) / 7
}
func soz(x uint64) (n int) {
	return sov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}