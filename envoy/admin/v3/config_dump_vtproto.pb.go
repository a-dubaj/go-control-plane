// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// source: envoy/admin/v3/config_dump.proto

package adminv3

import (
	anypb "github.com/planetscale/vtprotobuf/types/known/anypb"
	timestamppb "github.com/planetscale/vtprotobuf/types/known/timestamppb"
	proto "google.golang.org/protobuf/proto"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *ConfigDump) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *ConfigDump) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *ConfigDump) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if len(m.Configs) > 0 {
		for iNdEx := len(m.Configs) - 1; iNdEx >= 0; iNdEx-- {
			size, err := (*anypb.Any)(m.Configs[iNdEx]).MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *BootstrapConfigDump) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *BootstrapConfigDump) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *BootstrapConfigDump) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if m.LastUpdated != nil {
		size, err := (*timestamppb.Timestamp)(m.LastUpdated).MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x12
	}
	if m.Bootstrap != nil {
		if vtmsg, ok := interface{}(m.Bootstrap).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.Bootstrap)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = encodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SecretsConfigDump_DynamicSecret) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *SecretsConfigDump_DynamicSecret) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *SecretsConfigDump_DynamicSecret) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if m.ClientStatus != 0 {
		i = encodeVarint(dAtA, i, uint64(m.ClientStatus))
		i--
		dAtA[i] = 0x30
	}
	if m.ErrorState != nil {
		size, err := m.ErrorState.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x2a
	}
	if m.Secret != nil {
		size, err := (*anypb.Any)(m.Secret).MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x22
	}
	if m.LastUpdated != nil {
		size, err := (*timestamppb.Timestamp)(m.LastUpdated).MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.VersionInfo) > 0 {
		i -= len(m.VersionInfo)
		copy(dAtA[i:], m.VersionInfo)
		i = encodeVarint(dAtA, i, uint64(len(m.VersionInfo)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarint(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SecretsConfigDump_StaticSecret) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *SecretsConfigDump_StaticSecret) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *SecretsConfigDump_StaticSecret) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if m.Secret != nil {
		size, err := (*anypb.Any)(m.Secret).MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1a
	}
	if m.LastUpdated != nil {
		size, err := (*timestamppb.Timestamp)(m.LastUpdated).MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarint(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SecretsConfigDump) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *SecretsConfigDump) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *SecretsConfigDump) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if len(m.DynamicWarmingSecrets) > 0 {
		for iNdEx := len(m.DynamicWarmingSecrets) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.DynamicWarmingSecrets[iNdEx].MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.DynamicActiveSecrets) > 0 {
		for iNdEx := len(m.DynamicActiveSecrets) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.DynamicActiveSecrets[iNdEx].MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.StaticSecrets) > 0 {
		for iNdEx := len(m.StaticSecrets) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.StaticSecrets[iNdEx].MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ConfigDump) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Configs) > 0 {
		for _, e := range m.Configs {
			l = (*anypb.Any)(e).SizeVT()
			n += 1 + l + sov(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *BootstrapConfigDump) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Bootstrap != nil {
		if size, ok := interface{}(m.Bootstrap).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.Bootstrap)
		}
		n += 1 + l + sov(uint64(l))
	}
	if m.LastUpdated != nil {
		l = (*timestamppb.Timestamp)(m.LastUpdated).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *SecretsConfigDump_DynamicSecret) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	l = len(m.VersionInfo)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if m.LastUpdated != nil {
		l = (*timestamppb.Timestamp)(m.LastUpdated).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.Secret != nil {
		l = (*anypb.Any)(m.Secret).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.ErrorState != nil {
		l = m.ErrorState.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.ClientStatus != 0 {
		n += 1 + sov(uint64(m.ClientStatus))
	}
	n += len(m.unknownFields)
	return n
}

func (m *SecretsConfigDump_StaticSecret) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if m.LastUpdated != nil {
		l = (*timestamppb.Timestamp)(m.LastUpdated).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.Secret != nil {
		l = (*anypb.Any)(m.Secret).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *SecretsConfigDump) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.StaticSecrets) > 0 {
		for _, e := range m.StaticSecrets {
			l = e.SizeVT()
			n += 1 + l + sov(uint64(l))
		}
	}
	if len(m.DynamicActiveSecrets) > 0 {
		for _, e := range m.DynamicActiveSecrets {
			l = e.SizeVT()
			n += 1 + l + sov(uint64(l))
		}
	}
	if len(m.DynamicWarmingSecrets) > 0 {
		for _, e := range m.DynamicWarmingSecrets {
			l = e.SizeVT()
			n += 1 + l + sov(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}
