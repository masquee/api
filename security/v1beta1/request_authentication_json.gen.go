// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: security/v1beta1/request_authentication.proto

package v1beta1

import (
	bytes "bytes"
	fmt "fmt"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "istio.io/api/type/v1beta1"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for RequestAuthentication
func (this *RequestAuthentication) MarshalJSON() ([]byte, error) {
	str, err := RequestAuthenticationMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for RequestAuthentication
func (this *RequestAuthentication) UnmarshalJSON(b []byte) error {
	return RequestAuthenticationUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	RequestAuthenticationMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	RequestAuthenticationUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{AllowUnknownFields: true}
)
