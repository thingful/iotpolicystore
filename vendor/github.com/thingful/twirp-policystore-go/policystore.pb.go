// Code generated by protoc-gen-go. DO NOT EDIT.
// source: policystore.proto

package policystore

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// An enumeration which allows us to specify what type of sharing is to be
// defined for the specified sensor type. The default value is `SHARE` which
// implies sharing the data at full resolution. If this type is specified, it
// is an error if either of `buckets` or `interval` is also supplied.
type Operation_Action int32

const (
	Operation_SHARE      Operation_Action = 0
	Operation_BIN        Operation_Action = 1
	Operation_MOVING_AVG Operation_Action = 2
)

var Operation_Action_name = map[int32]string{
	0: "SHARE",
	1: "BIN",
	2: "MOVING_AVG",
}
var Operation_Action_value = map[string]int32{
	"SHARE":      0,
	"BIN":        1,
	"MOVING_AVG": 2,
}

func (x Operation_Action) String() string {
	return proto.EnumName(Operation_Action_name, int32(x))
}
func (Operation_Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_policystore_4484fb306aaefb07, []int{0, 0}
}

// Operation is a message used to describe an operation that may be applied to
// a specific data type published by a SmartCitizen device. The message contains
// two required fields: the sensor_id (this is the type of data we are entitling
// over), and a specified operation to be performed on that sensor type. This
// can be one of three actions: to share the sensor without modification, to
// apply a binning algorithm to the data so we output a bucketed value, or a
// moving average calculated dynamically for incoming values.
//
// If an operation specifies an Action type of `BIN`, then the optional
// `buckets` parameter is required, similarly if an action type of `MOVING_AVG`
// is specified, then `interval` is a required field.
type Operation struct {
	// The unique id of the sensor type for which this specific entitlement is
	// defined. This is a required field.
	SensorId uint32 `protobuf:"varint,1,opt,name=sensor_id,json=sensorId,proto3" json:"sensor_id,omitempty"`
	// The specific action this operation defines for the sensor type. This is a
	// required field.
	Action Operation_Action `protobuf:"varint,2,opt,name=action,proto3,enum=decode.iot.policystore.Operation_Action" json:"action,omitempty"`
	// The bins attribute is used to specify the the bins into which incoming
	// values should be classified. Each element in the list is the upper
	// inclusive bound of a bin. The values submitted must be sorted in strictly
	// increasing order. There is no need to add a highest bin with +Inf bound, it
	// will be added implicitly. This field is optional unless an Action of `BIN`
	// has been requested, in which case it is required. It is an error to send
	// values for this attribute unless the value of Action is `BIN`.
	Bins []float64 `protobuf:"fixed64,3,rep,packed,name=bins,proto3" json:"bins,omitempty"`
	// This attribute is used to control the entitlement in the case for which we
	// have specified an action type representing a moving average. It represents
	// the interval in seconds over which the moving average should be calculated,
	// e.g. for a 15 minute moving average the value supplied here would be 900.
	// This field is optional unless an Action of `MOVING_AVG` has been specified,
	// in which case it is required. It is an error to send a value for this
	// attribute unless the value of Action is `MOVING_AVG`.
	Interval             uint32   `protobuf:"varint,4,opt,name=interval,proto3" json:"interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Operation) Reset()         { *m = Operation{} }
func (m *Operation) String() string { return proto.CompactTextString(m) }
func (*Operation) ProtoMessage()    {}
func (*Operation) Descriptor() ([]byte, []int) {
	return fileDescriptor_policystore_4484fb306aaefb07, []int{0}
}
func (m *Operation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Operation.Unmarshal(m, b)
}
func (m *Operation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Operation.Marshal(b, m, deterministic)
}
func (dst *Operation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Operation.Merge(dst, src)
}
func (m *Operation) XXX_Size() int {
	return xxx_messageInfo_Operation.Size(m)
}
func (m *Operation) XXX_DiscardUnknown() {
	xxx_messageInfo_Operation.DiscardUnknown(m)
}

var xxx_messageInfo_Operation proto.InternalMessageInfo

func (m *Operation) GetSensorId() uint32 {
	if m != nil {
		return m.SensorId
	}
	return 0
}

func (m *Operation) GetAction() Operation_Action {
	if m != nil {
		return m.Action
	}
	return Operation_SHARE
}

func (m *Operation) GetBins() []float64 {
	if m != nil {
		return m.Bins
	}
	return nil
}

func (m *Operation) GetInterval() uint32 {
	if m != nil {
		return m.Interval
	}
	return 0
}

// CreateEntitlementPolicyRequest is a message sent to the policy registration
// service to create a new entitlement policy. An entitlement policy is a
// collection of one or more "Operations". A single Operation specifies an
// functional transformation to be performed on a single data channel being
// published by a SmartCitizen device. The policy as a whole is comprised of
// one or more Entitlements.
type CreateEntitlementPolicyRequest struct {
	// This attribute contains the public part of a key pair created by the
	// caller. The caller must keep the private key secret as this is will be
	// required for them to be able to decrypt data.
	PublicKey string `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// This attribute is used to attach a human friendly label to the policy
	// suitable for presenting to the end user in the DECODE wallet. This is a
	// required field.
	Label string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	// The list of operations we wish to create for the policy. This field is
	// required, and it is required that the client supplies at least one
	// Operation.
	Operations           []*Operation `protobuf:"bytes,3,rep,name=operations,proto3" json:"operations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateEntitlementPolicyRequest) Reset()         { *m = CreateEntitlementPolicyRequest{} }
func (m *CreateEntitlementPolicyRequest) String() string { return proto.CompactTextString(m) }
func (*CreateEntitlementPolicyRequest) ProtoMessage()    {}
func (*CreateEntitlementPolicyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_policystore_4484fb306aaefb07, []int{1}
}
func (m *CreateEntitlementPolicyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEntitlementPolicyRequest.Unmarshal(m, b)
}
func (m *CreateEntitlementPolicyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEntitlementPolicyRequest.Marshal(b, m, deterministic)
}
func (dst *CreateEntitlementPolicyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEntitlementPolicyRequest.Merge(dst, src)
}
func (m *CreateEntitlementPolicyRequest) XXX_Size() int {
	return xxx_messageInfo_CreateEntitlementPolicyRequest.Size(m)
}
func (m *CreateEntitlementPolicyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEntitlementPolicyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEntitlementPolicyRequest proto.InternalMessageInfo

func (m *CreateEntitlementPolicyRequest) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

func (m *CreateEntitlementPolicyRequest) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *CreateEntitlementPolicyRequest) GetOperations() []*Operation {
	if m != nil {
		return m.Operations
	}
	return nil
}

// CreateEntitlementPolicyResponse is a message returned by the service after a
// policy has been created. The message simply contains an identifier for the
// policy, as well as a token that the caller must protect.
type CreateEntitlementPolicyResponse struct {
	// This attribute contains a unique identifier for the policy that can be used
	// for later requests to either apply a policy to a specific device, or to
	// delete the policy and so prevent new instances being applied to devices.
	PolicyId string `protobuf:"bytes,1,opt,name=policy_id,json=policyId,proto3" json:"policy_id,omitempty"`
	// This attribute contains a secret generated by the service that is
	// associated with the policy. This token is required to be presented by a
	// caller when deleting a policy, so must be treated as confidential by the
	// caller.
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateEntitlementPolicyResponse) Reset()         { *m = CreateEntitlementPolicyResponse{} }
func (m *CreateEntitlementPolicyResponse) String() string { return proto.CompactTextString(m) }
func (*CreateEntitlementPolicyResponse) ProtoMessage()    {}
func (*CreateEntitlementPolicyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_policystore_4484fb306aaefb07, []int{2}
}
func (m *CreateEntitlementPolicyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEntitlementPolicyResponse.Unmarshal(m, b)
}
func (m *CreateEntitlementPolicyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEntitlementPolicyResponse.Marshal(b, m, deterministic)
}
func (dst *CreateEntitlementPolicyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEntitlementPolicyResponse.Merge(dst, src)
}
func (m *CreateEntitlementPolicyResponse) XXX_Size() int {
	return xxx_messageInfo_CreateEntitlementPolicyResponse.Size(m)
}
func (m *CreateEntitlementPolicyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEntitlementPolicyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEntitlementPolicyResponse proto.InternalMessageInfo

func (m *CreateEntitlementPolicyResponse) GetPolicyId() string {
	if m != nil {
		return m.PolicyId
	}
	return ""
}

func (m *CreateEntitlementPolicyResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// DeleteEntitlementPolicyRequest is a message that can be sent to the
// registration service in order to delete an existing policy.
//
// Deleting a policy does not affect any already existing streams configured for
// the policy, it just stops any new instances of this policy being applied to
// other devices.
type DeleteEntitlementPolicyRequest struct {
	// This attribute contains the unique policy identifier returned when creating
	// the policy. This is a requiredi field.
	PolicyId string `protobuf:"bytes,1,opt,name=policy_id,json=policyId,proto3" json:"policy_id,omitempty"`
	// This attribute contains the token returned to the creator when they
	// created the policy, and must match the value stored within the
	// PolicyStore. This is a required field.
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteEntitlementPolicyRequest) Reset()         { *m = DeleteEntitlementPolicyRequest{} }
func (m *DeleteEntitlementPolicyRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteEntitlementPolicyRequest) ProtoMessage()    {}
func (*DeleteEntitlementPolicyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_policystore_4484fb306aaefb07, []int{3}
}
func (m *DeleteEntitlementPolicyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteEntitlementPolicyRequest.Unmarshal(m, b)
}
func (m *DeleteEntitlementPolicyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteEntitlementPolicyRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteEntitlementPolicyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteEntitlementPolicyRequest.Merge(dst, src)
}
func (m *DeleteEntitlementPolicyRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteEntitlementPolicyRequest.Size(m)
}
func (m *DeleteEntitlementPolicyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteEntitlementPolicyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteEntitlementPolicyRequest proto.InternalMessageInfo

func (m *DeleteEntitlementPolicyRequest) GetPolicyId() string {
	if m != nil {
		return m.PolicyId
	}
	return ""
}

func (m *DeleteEntitlementPolicyRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// DeleteEntitlementPolicyResponse is a placeholder response returned from a
// delete request. Currently empty, but reserved for any fields identified for
// future iterations.
type DeleteEntitlementPolicyResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteEntitlementPolicyResponse) Reset()         { *m = DeleteEntitlementPolicyResponse{} }
func (m *DeleteEntitlementPolicyResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteEntitlementPolicyResponse) ProtoMessage()    {}
func (*DeleteEntitlementPolicyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_policystore_4484fb306aaefb07, []int{4}
}
func (m *DeleteEntitlementPolicyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteEntitlementPolicyResponse.Unmarshal(m, b)
}
func (m *DeleteEntitlementPolicyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteEntitlementPolicyResponse.Marshal(b, m, deterministic)
}
func (dst *DeleteEntitlementPolicyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteEntitlementPolicyResponse.Merge(dst, src)
}
func (m *DeleteEntitlementPolicyResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteEntitlementPolicyResponse.Size(m)
}
func (m *DeleteEntitlementPolicyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteEntitlementPolicyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteEntitlementPolicyResponse proto.InternalMessageInfo

// ListEntitlementPoliciesRequest is the message sent to the service in order
// to receive a list of currently defined entitlement policies. Currently this
// message is empty as we simply return a list of all known policies, but this
// message may be extended should a need be identified to paginate through
// policies, or apply any search or filtering techniques.
type ListEntitlementPoliciesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListEntitlementPoliciesRequest) Reset()         { *m = ListEntitlementPoliciesRequest{} }
func (m *ListEntitlementPoliciesRequest) String() string { return proto.CompactTextString(m) }
func (*ListEntitlementPoliciesRequest) ProtoMessage()    {}
func (*ListEntitlementPoliciesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_policystore_4484fb306aaefb07, []int{5}
}
func (m *ListEntitlementPoliciesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListEntitlementPoliciesRequest.Unmarshal(m, b)
}
func (m *ListEntitlementPoliciesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListEntitlementPoliciesRequest.Marshal(b, m, deterministic)
}
func (dst *ListEntitlementPoliciesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListEntitlementPoliciesRequest.Merge(dst, src)
}
func (m *ListEntitlementPoliciesRequest) XXX_Size() int {
	return xxx_messageInfo_ListEntitlementPoliciesRequest.Size(m)
}
func (m *ListEntitlementPoliciesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListEntitlementPoliciesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListEntitlementPoliciesRequest proto.InternalMessageInfo

// ListEntitlementPoliciesResponse is the response to the method call to list
// policies. It simply returns a list of all currently registered and
// non-deleted policies. This is intended to be able to be fed to the DECODE
// wallet in order to allow participant to choose which entitlements to apply to
// their devices.
type ListEntitlementPoliciesResponse struct {
	// This attribute contains the list of all policies currently available on
	// the device registration service.
	Policies             []*ListEntitlementPoliciesResponse_Policy `protobuf:"bytes,1,rep,name=policies,proto3" json:"policies,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *ListEntitlementPoliciesResponse) Reset()         { *m = ListEntitlementPoliciesResponse{} }
func (m *ListEntitlementPoliciesResponse) String() string { return proto.CompactTextString(m) }
func (*ListEntitlementPoliciesResponse) ProtoMessage()    {}
func (*ListEntitlementPoliciesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_policystore_4484fb306aaefb07, []int{6}
}
func (m *ListEntitlementPoliciesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListEntitlementPoliciesResponse.Unmarshal(m, b)
}
func (m *ListEntitlementPoliciesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListEntitlementPoliciesResponse.Marshal(b, m, deterministic)
}
func (dst *ListEntitlementPoliciesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListEntitlementPoliciesResponse.Merge(dst, src)
}
func (m *ListEntitlementPoliciesResponse) XXX_Size() int {
	return xxx_messageInfo_ListEntitlementPoliciesResponse.Size(m)
}
func (m *ListEntitlementPoliciesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListEntitlementPoliciesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListEntitlementPoliciesResponse proto.InternalMessageInfo

func (m *ListEntitlementPoliciesResponse) GetPolicies() []*ListEntitlementPoliciesResponse_Policy {
	if m != nil {
		return m.Policies
	}
	return nil
}

// Policy is a nested type used to be able to cleanly return a list of
// Policies within a single response. Each Policy instance contains the id of
// the policy, the list of entitlements defined by the policy, as well as the
// policy's public key.
type ListEntitlementPoliciesResponse_Policy struct {
	// This attribute contains the unique identifier of the policy.
	PolicyId string `protobuf:"bytes,1,opt,name=policy_id,json=policyId,proto3" json:"policy_id,omitempty"`
	// This attribute contains a human friendly label describing the policy
	// suitable for rendering in the DECODE wallet
	Label string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	// This field contains a list of the operations that define the policy.
	Operations []*Operation `protobuf:"bytes,3,rep,name=operations,proto3" json:"operations,omitempty"`
	// This attribute contains the public key of the policy. This public key
	// attribute is the label applied to the bucket within the datastore which
	// will be how data can be downloaded for the entitlement policy.
	PublicKey            string   `protobuf:"bytes,4,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListEntitlementPoliciesResponse_Policy) Reset() {
	*m = ListEntitlementPoliciesResponse_Policy{}
}
func (m *ListEntitlementPoliciesResponse_Policy) String() string { return proto.CompactTextString(m) }
func (*ListEntitlementPoliciesResponse_Policy) ProtoMessage()    {}
func (*ListEntitlementPoliciesResponse_Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_policystore_4484fb306aaefb07, []int{6, 0}
}
func (m *ListEntitlementPoliciesResponse_Policy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListEntitlementPoliciesResponse_Policy.Unmarshal(m, b)
}
func (m *ListEntitlementPoliciesResponse_Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListEntitlementPoliciesResponse_Policy.Marshal(b, m, deterministic)
}
func (dst *ListEntitlementPoliciesResponse_Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListEntitlementPoliciesResponse_Policy.Merge(dst, src)
}
func (m *ListEntitlementPoliciesResponse_Policy) XXX_Size() int {
	return xxx_messageInfo_ListEntitlementPoliciesResponse_Policy.Size(m)
}
func (m *ListEntitlementPoliciesResponse_Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_ListEntitlementPoliciesResponse_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_ListEntitlementPoliciesResponse_Policy proto.InternalMessageInfo

func (m *ListEntitlementPoliciesResponse_Policy) GetPolicyId() string {
	if m != nil {
		return m.PolicyId
	}
	return ""
}

func (m *ListEntitlementPoliciesResponse_Policy) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *ListEntitlementPoliciesResponse_Policy) GetOperations() []*Operation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *ListEntitlementPoliciesResponse_Policy) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

func init() {
	proto.RegisterType((*Operation)(nil), "decode.iot.policystore.Operation")
	proto.RegisterType((*CreateEntitlementPolicyRequest)(nil), "decode.iot.policystore.CreateEntitlementPolicyRequest")
	proto.RegisterType((*CreateEntitlementPolicyResponse)(nil), "decode.iot.policystore.CreateEntitlementPolicyResponse")
	proto.RegisterType((*DeleteEntitlementPolicyRequest)(nil), "decode.iot.policystore.DeleteEntitlementPolicyRequest")
	proto.RegisterType((*DeleteEntitlementPolicyResponse)(nil), "decode.iot.policystore.DeleteEntitlementPolicyResponse")
	proto.RegisterType((*ListEntitlementPoliciesRequest)(nil), "decode.iot.policystore.ListEntitlementPoliciesRequest")
	proto.RegisterType((*ListEntitlementPoliciesResponse)(nil), "decode.iot.policystore.ListEntitlementPoliciesResponse")
	proto.RegisterType((*ListEntitlementPoliciesResponse_Policy)(nil), "decode.iot.policystore.ListEntitlementPoliciesResponse.Policy")
	proto.RegisterEnum("decode.iot.policystore.Operation_Action", Operation_Action_name, Operation_Action_value)
}

func init() { proto.RegisterFile("policystore.proto", fileDescriptor_policystore_4484fb306aaefb07) }

var fileDescriptor_policystore_4484fb306aaefb07 = []byte{
	// 469 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xd9, 0x24, 0x0d, 0xf1, 0x44, 0xad, 0xc2, 0x0a, 0x81, 0x15, 0x44, 0x92, 0xfa, 0x94,
	0x03, 0xf2, 0x21, 0x48, 0x70, 0x43, 0xa4, 0x50, 0x95, 0x08, 0x68, 0xd1, 0x06, 0xf5, 0xd0, 0x4b,
	0xe4, 0x3f, 0x73, 0x58, 0xd5, 0xec, 0x1a, 0xef, 0x16, 0x29, 0xaf, 0xc0, 0x13, 0x70, 0xe9, 0x5b,
	0xf0, 0x10, 0x3c, 0x16, 0xf2, 0xae, 0x5d, 0x45, 0x11, 0x6b, 0x0b, 0x24, 0x6e, 0xd9, 0xc9, 0xec,
	0xb7, 0xbf, 0x6f, 0xe6, 0x93, 0xe1, 0x41, 0x2e, 0x33, 0x9e, 0x6c, 0x95, 0x96, 0x05, 0x86, 0x79,
	0x21, 0xb5, 0xa4, 0x8f, 0x52, 0x4c, 0x64, 0x8a, 0x21, 0x97, 0x3a, 0xdc, 0xf9, 0x37, 0xf8, 0x45,
	0xc0, 0xbb, 0xc8, 0xb1, 0x88, 0x34, 0x97, 0x82, 0x3e, 0x01, 0x4f, 0xa1, 0x50, 0xb2, 0xd8, 0xf0,
	0xd4, 0x27, 0x33, 0x32, 0x3f, 0x64, 0x03, 0x5b, 0x58, 0xa5, 0xf4, 0x35, 0xf4, 0xa3, 0xa4, 0x6c,
	0xf3, 0x3b, 0x33, 0x32, 0x3f, 0x5a, 0xcc, 0xc3, 0x3f, 0x6b, 0x86, 0x77, 0x7a, 0xe1, 0xd2, 0xf4,
	0xb3, 0xea, 0x1e, 0xa5, 0xd0, 0x8b, 0xb9, 0x50, 0x7e, 0x77, 0xd6, 0x9d, 0x13, 0x66, 0x7e, 0xd3,
	0x31, 0x0c, 0xb8, 0xd0, 0x58, 0x7c, 0x8b, 0x32, 0xbf, 0x67, 0x5f, 0xac, 0xcf, 0xc1, 0x33, 0xe8,
	0x5b, 0x05, 0xea, 0xc1, 0xc1, 0xfa, 0xdd, 0x92, 0x9d, 0x8e, 0xee, 0xd1, 0xfb, 0xd0, 0x3d, 0x59,
	0x9d, 0x8f, 0x08, 0x3d, 0x02, 0xf8, 0x78, 0x71, 0xb9, 0x3a, 0x3f, 0xdb, 0x2c, 0x2f, 0xcf, 0x46,
	0x9d, 0xe0, 0x07, 0x81, 0xc9, 0x9b, 0x02, 0x23, 0x8d, 0xa7, 0x42, 0x73, 0x9d, 0xe1, 0x17, 0x14,
	0xfa, 0x93, 0xe1, 0x62, 0xf8, 0xf5, 0x06, 0x95, 0xa6, 0x4f, 0x01, 0xf2, 0x9b, 0x38, 0xe3, 0xc9,
	0xe6, 0x1a, 0xb7, 0xc6, 0xa0, 0xc7, 0x3c, 0x5b, 0x79, 0x8f, 0x5b, 0xfa, 0x10, 0x0e, 0xb2, 0x28,
	0xc6, 0xcc, 0x18, 0xf4, 0x98, 0x3d, 0xd0, 0x25, 0x80, 0xac, 0x1d, 0x59, 0xf6, 0xe1, 0xe2, 0xb8,
	0xd5, 0x3b, 0xdb, 0xb9, 0x14, 0x7c, 0x86, 0xa9, 0x93, 0x4c, 0xe5, 0x52, 0x28, 0x2c, 0x47, 0x6f,
	0x75, 0xea, 0xd1, 0x7b, 0x6c, 0x60, 0x0b, 0xab, 0xb4, 0x04, 0xd3, 0xf2, 0x1a, 0x45, 0x0d, 0x66,
	0x0e, 0xc1, 0x1a, 0x26, 0x6f, 0x31, 0xc3, 0x06, 0xbf, 0xff, 0x20, 0x7a, 0x0c, 0x53, 0xa7, 0xa8,
	0x45, 0x0d, 0x66, 0x30, 0xf9, 0xc0, 0x95, 0xde, 0x6f, 0xe0, 0xa8, 0xaa, 0x77, 0x83, 0xdb, 0x0e,
	0x4c, 0x9d, 0x2d, 0x95, 0xe1, 0x2b, 0xb0, 0x28, 0x1c, 0x95, 0x4f, 0xcc, 0x50, 0x5f, 0xb9, 0x86,
	0xda, 0x22, 0x15, 0x56, 0x7c, 0x77, 0x7a, 0xe3, 0x5b, 0x02, 0x7d, 0x5b, 0x6c, 0x1d, 0xc1, 0x7f,
	0x59, 0xf8, 0x5e, 0xd0, 0x7a, 0x7b, 0x41, 0x5b, 0xfc, 0xec, 0xc2, 0xd0, 0xf2, 0xad, 0x4b, 0x11,
	0xfa, 0x9d, 0xc0, 0x63, 0x47, 0x40, 0xe8, 0x0b, 0xd7, 0xcb, 0xcd, 0x59, 0x1f, 0xbf, 0xfc, 0xeb,
	0x7b, 0xd5, 0x62, 0x4a, 0x18, 0x47, 0x04, 0xdc, 0x30, 0xcd, 0x41, 0x74, 0xc3, 0xb4, 0x64, 0xcd,
	0xc0, 0x38, 0xd6, 0xef, 0x86, 0x69, 0x4e, 0xa7, 0x1b, 0xa6, 0x25, 0x67, 0x27, 0x87, 0x57, 0xc3,
	0x9d, 0xf6, 0xb8, 0x6f, 0x3e, 0xad, 0xcf, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xf3, 0xcf, 0x0a,
	0x9c, 0x6f, 0x05, 0x00, 0x00,
}