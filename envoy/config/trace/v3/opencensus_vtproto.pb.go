// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// source: envoy/config/trace/v3/opencensus.proto

package tracev3

import (
	proto "google.golang.org/protobuf/proto"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *OpenCensusConfig) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *OpenCensusConfig) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *OpenCensusConfig) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if m.OcagentGrpcService != nil {
		if vtmsg, ok := interface{}(m.OcagentGrpcService).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.OcagentGrpcService)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = encodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0x72
	}
	if m.StackdriverGrpcService != nil {
		if vtmsg, ok := interface{}(m.StackdriverGrpcService).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.StackdriverGrpcService)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = encodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0x6a
	}
	if len(m.OcagentAddress) > 0 {
		i -= len(m.OcagentAddress)
		copy(dAtA[i:], m.OcagentAddress)
		i = encodeVarint(dAtA, i, uint64(len(m.OcagentAddress)))
		i--
		dAtA[i] = 0x62
	}
	if m.OcagentExporterEnabled {
		i--
		if m.OcagentExporterEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x58
	}
	if len(m.StackdriverAddress) > 0 {
		i -= len(m.StackdriverAddress)
		copy(dAtA[i:], m.StackdriverAddress)
		i = encodeVarint(dAtA, i, uint64(len(m.StackdriverAddress)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.OutgoingTraceContext) > 0 {
		var pksize2 int
		for _, num := range m.OutgoingTraceContext {
			pksize2 += sov(uint64(num))
		}
		i -= pksize2
		j1 := i
		for _, num1 := range m.OutgoingTraceContext {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA[j1] = uint8(num)
			j1++
		}
		i = encodeVarint(dAtA, i, uint64(pksize2))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.IncomingTraceContext) > 0 {
		var pksize4 int
		for _, num := range m.IncomingTraceContext {
			pksize4 += sov(uint64(num))
		}
		i -= pksize4
		j3 := i
		for _, num1 := range m.IncomingTraceContext {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA[j3] = uint8(num)
			j3++
		}
		i = encodeVarint(dAtA, i, uint64(pksize4))
		i--
		dAtA[i] = 0x42
	}
	if len(m.ZipkinUrl) > 0 {
		i -= len(m.ZipkinUrl)
		copy(dAtA[i:], m.ZipkinUrl)
		i = encodeVarint(dAtA, i, uint64(len(m.ZipkinUrl)))
		i--
		dAtA[i] = 0x32
	}
	if m.ZipkinExporterEnabled {
		i--
		if m.ZipkinExporterEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if len(m.StackdriverProjectId) > 0 {
		i -= len(m.StackdriverProjectId)
		copy(dAtA[i:], m.StackdriverProjectId)
		i = encodeVarint(dAtA, i, uint64(len(m.StackdriverProjectId)))
		i--
		dAtA[i] = 0x22
	}
	if m.StackdriverExporterEnabled {
		i--
		if m.StackdriverExporterEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.StdoutExporterEnabled {
		i--
		if m.StdoutExporterEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.TraceConfig != nil {
		if vtmsg, ok := interface{}(m.TraceConfig).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.TraceConfig)
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

func (m *OpenCensusConfig) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TraceConfig != nil {
		if size, ok := interface{}(m.TraceConfig).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.TraceConfig)
		}
		n += 1 + l + sov(uint64(l))
	}
	if m.StdoutExporterEnabled {
		n += 2
	}
	if m.StackdriverExporterEnabled {
		n += 2
	}
	l = len(m.StackdriverProjectId)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if m.ZipkinExporterEnabled {
		n += 2
	}
	l = len(m.ZipkinUrl)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if len(m.IncomingTraceContext) > 0 {
		l = 0
		for _, e := range m.IncomingTraceContext {
			l += sov(uint64(e))
		}
		n += 1 + sov(uint64(l)) + l
	}
	if len(m.OutgoingTraceContext) > 0 {
		l = 0
		for _, e := range m.OutgoingTraceContext {
			l += sov(uint64(e))
		}
		n += 1 + sov(uint64(l)) + l
	}
	l = len(m.StackdriverAddress)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if m.OcagentExporterEnabled {
		n += 2
	}
	l = len(m.OcagentAddress)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if m.StackdriverGrpcService != nil {
		if size, ok := interface{}(m.StackdriverGrpcService).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.StackdriverGrpcService)
		}
		n += 1 + l + sov(uint64(l))
	}
	if m.OcagentGrpcService != nil {
		if size, ok := interface{}(m.OcagentGrpcService).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.OcagentGrpcService)
		}
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}
