// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// source: envoy/config/grpc_credential/v3/file_based_metadata.proto

package grpc_credentialv3

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

func (m *FileBasedMetadataConfig) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *FileBasedMetadataConfig) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *FileBasedMetadataConfig) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if len(m.HeaderPrefix) > 0 {
		i -= len(m.HeaderPrefix)
		copy(dAtA[i:], m.HeaderPrefix)
		i = encodeVarint(dAtA, i, uint64(len(m.HeaderPrefix)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.HeaderKey) > 0 {
		i -= len(m.HeaderKey)
		copy(dAtA[i:], m.HeaderKey)
		i = encodeVarint(dAtA, i, uint64(len(m.HeaderKey)))
		i--
		dAtA[i] = 0x12
	}
	if m.SecretData != nil {
		if vtmsg, ok := interface{}(m.SecretData).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.SecretData)
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

func (m *FileBasedMetadataConfig) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SecretData != nil {
		if size, ok := interface{}(m.SecretData).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.SecretData)
		}
		n += 1 + l + sov(uint64(l))
	}
	l = len(m.HeaderKey)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	l = len(m.HeaderPrefix)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}
