/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-gogo.
// source: k8s.io/kubernetes/pkg/apis/policy/v1beta1/generated.proto
// DO NOT EDIT!

/*
	Package v1beta1 is a generated protocol buffer package.

	It is generated from these files:
		k8s.io/kubernetes/pkg/apis/policy/v1beta1/generated.proto

	It has these top-level messages:
		Eviction
		PodDisruptionBudget
		PodDisruptionBudgetList
		PodDisruptionBudgetSpec
		PodDisruptionBudgetStatus
*/
package v1beta1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import k8s_io_apimachinery_pkg_apis_meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

import k8s_io_apimachinery_pkg_util_intstr "k8s.io/apimachinery/pkg/util/intstr"

import strings "strings"
import reflect "reflect"
import github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

func (m *Eviction) Reset()                    { *m = Eviction{} }
func (*Eviction) ProtoMessage()               {}
func (*Eviction) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *PodDisruptionBudget) Reset()                    { *m = PodDisruptionBudget{} }
func (*PodDisruptionBudget) ProtoMessage()               {}
func (*PodDisruptionBudget) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{1} }

func (m *PodDisruptionBudgetList) Reset()                    { *m = PodDisruptionBudgetList{} }
func (*PodDisruptionBudgetList) ProtoMessage()               {}
func (*PodDisruptionBudgetList) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{2} }

func (m *PodDisruptionBudgetSpec) Reset()                    { *m = PodDisruptionBudgetSpec{} }
func (*PodDisruptionBudgetSpec) ProtoMessage()               {}
func (*PodDisruptionBudgetSpec) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{3} }

func (m *PodDisruptionBudgetStatus) Reset()      { *m = PodDisruptionBudgetStatus{} }
func (*PodDisruptionBudgetStatus) ProtoMessage() {}
func (*PodDisruptionBudgetStatus) Descriptor() ([]byte, []int) {
	return fileDescriptorGenerated, []int{4}
}

func init() {
	proto.RegisterType((*Eviction)(nil), "k8s.io.client-go.pkg.apis.policy.v1beta1.Eviction")
	proto.RegisterType((*PodDisruptionBudget)(nil), "k8s.io.client-go.pkg.apis.policy.v1beta1.PodDisruptionBudget")
	proto.RegisterType((*PodDisruptionBudgetList)(nil), "k8s.io.client-go.pkg.apis.policy.v1beta1.PodDisruptionBudgetList")
	proto.RegisterType((*PodDisruptionBudgetSpec)(nil), "k8s.io.client-go.pkg.apis.policy.v1beta1.PodDisruptionBudgetSpec")
	proto.RegisterType((*PodDisruptionBudgetStatus)(nil), "k8s.io.client-go.pkg.apis.policy.v1beta1.PodDisruptionBudgetStatus")
}
func (m *Eviction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Eviction) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ObjectMeta.Size()))
	n1, err := m.ObjectMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if m.DeleteOptions != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.DeleteOptions.Size()))
		n2, err := m.DeleteOptions.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *PodDisruptionBudget) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PodDisruptionBudget) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ObjectMeta.Size()))
	n3, err := m.ObjectMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	dAtA[i] = 0x12
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Spec.Size()))
	n4, err := m.Spec.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	dAtA[i] = 0x1a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Status.Size()))
	n5, err := m.Status.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n5
	return i, nil
}

func (m *PodDisruptionBudgetList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PodDisruptionBudgetList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ListMeta.Size()))
	n6, err := m.ListMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n6
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			dAtA[i] = 0x12
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *PodDisruptionBudgetSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PodDisruptionBudgetSpec) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.MinAvailable != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.MinAvailable.Size()))
		n7, err := m.MinAvailable.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	if m.Selector != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Selector.Size()))
		n8, err := m.Selector.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n8
	}
	if m.MaxUnavailable != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.MaxUnavailable.Size()))
		n9, err := m.MaxUnavailable.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n9
	}
	return i, nil
}

func (m *PodDisruptionBudgetStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PodDisruptionBudgetStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ObservedGeneration))
	if len(m.DisruptedPods) > 0 {
		for k := range m.DisruptedPods {
			dAtA[i] = 0x12
			i++
			v := m.DisruptedPods[k]
			msgSize := 0
			if (&v) != nil {
				msgSize = (&v).Size()
				msgSize += 1 + sovGenerated(uint64(msgSize))
			}
			mapSize := 1 + len(k) + sovGenerated(uint64(len(k))) + msgSize
			i = encodeVarintGenerated(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintGenerated(dAtA, i, uint64((&v).Size()))
			n10, err := (&v).MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n10
		}
	}
	dAtA[i] = 0x18
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.PodDisruptionsAllowed))
	dAtA[i] = 0x20
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.CurrentHealthy))
	dAtA[i] = 0x28
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.DesiredHealthy))
	dAtA[i] = 0x30
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ExpectedPods))
	return i, nil
}

func encodeFixed64Generated(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Generated(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Eviction) Size() (n int) {
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if m.DeleteOptions != nil {
		l = m.DeleteOptions.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func (m *PodDisruptionBudget) Size() (n int) {
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Spec.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Status.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *PodDisruptionBudgetList) Size() (n int) {
	var l int
	_ = l
	l = m.ListMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *PodDisruptionBudgetSpec) Size() (n int) {
	var l int
	_ = l
	if m.MinAvailable != nil {
		l = m.MinAvailable.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.Selector != nil {
		l = m.Selector.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.MaxUnavailable != nil {
		l = m.MaxUnavailable.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func (m *PodDisruptionBudgetStatus) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovGenerated(uint64(m.ObservedGeneration))
	if len(m.DisruptedPods) > 0 {
		for k, v := range m.DisruptedPods {
			_ = k
			_ = v
			l = v.Size()
			mapEntrySize := 1 + len(k) + sovGenerated(uint64(len(k))) + 1 + l + sovGenerated(uint64(l))
			n += mapEntrySize + 1 + sovGenerated(uint64(mapEntrySize))
		}
	}
	n += 1 + sovGenerated(uint64(m.PodDisruptionsAllowed))
	n += 1 + sovGenerated(uint64(m.CurrentHealthy))
	n += 1 + sovGenerated(uint64(m.DesiredHealthy))
	n += 1 + sovGenerated(uint64(m.ExpectedPods))
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Eviction) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Eviction{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(this.ObjectMeta.String(), "ObjectMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`DeleteOptions:` + strings.Replace(fmt.Sprintf("%v", this.DeleteOptions), "DeleteOptions", "k8s_io_apimachinery_pkg_apis_meta_v1.DeleteOptions", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *PodDisruptionBudget) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PodDisruptionBudget{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(this.ObjectMeta.String(), "ObjectMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Spec:` + strings.Replace(strings.Replace(this.Spec.String(), "PodDisruptionBudgetSpec", "PodDisruptionBudgetSpec", 1), `&`, ``, 1) + `,`,
		`Status:` + strings.Replace(strings.Replace(this.Status.String(), "PodDisruptionBudgetStatus", "PodDisruptionBudgetStatus", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *PodDisruptionBudgetList) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PodDisruptionBudgetList{`,
		`ListMeta:` + strings.Replace(strings.Replace(this.ListMeta.String(), "ListMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ListMeta", 1), `&`, ``, 1) + `,`,
		`Items:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Items), "PodDisruptionBudget", "PodDisruptionBudget", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *PodDisruptionBudgetSpec) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PodDisruptionBudgetSpec{`,
		`MinAvailable:` + strings.Replace(fmt.Sprintf("%v", this.MinAvailable), "IntOrString", "k8s_io_apimachinery_pkg_util_intstr.IntOrString", 1) + `,`,
		`Selector:` + strings.Replace(fmt.Sprintf("%v", this.Selector), "LabelSelector", "k8s_io_apimachinery_pkg_apis_meta_v1.LabelSelector", 1) + `,`,
		`MaxUnavailable:` + strings.Replace(fmt.Sprintf("%v", this.MaxUnavailable), "IntOrString", "k8s_io_apimachinery_pkg_util_intstr.IntOrString", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *PodDisruptionBudgetStatus) String() string {
	if this == nil {
		return "nil"
	}
	keysForDisruptedPods := make([]string, 0, len(this.DisruptedPods))
	for k := range this.DisruptedPods {
		keysForDisruptedPods = append(keysForDisruptedPods, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForDisruptedPods)
	mapStringForDisruptedPods := "map[string]k8s_io_apimachinery_pkg_apis_meta_v1.Time{"
	for _, k := range keysForDisruptedPods {
		mapStringForDisruptedPods += fmt.Sprintf("%v: %v,", k, this.DisruptedPods[k])
	}
	mapStringForDisruptedPods += "}"
	s := strings.Join([]string{`&PodDisruptionBudgetStatus{`,
		`ObservedGeneration:` + fmt.Sprintf("%v", this.ObservedGeneration) + `,`,
		`DisruptedPods:` + mapStringForDisruptedPods + `,`,
		`PodDisruptionsAllowed:` + fmt.Sprintf("%v", this.PodDisruptionsAllowed) + `,`,
		`CurrentHealthy:` + fmt.Sprintf("%v", this.CurrentHealthy) + `,`,
		`DesiredHealthy:` + fmt.Sprintf("%v", this.DesiredHealthy) + `,`,
		`ExpectedPods:` + fmt.Sprintf("%v", this.ExpectedPods) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Eviction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Eviction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Eviction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeleteOptions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DeleteOptions == nil {
				m.DeleteOptions = &k8s_io_apimachinery_pkg_apis_meta_v1.DeleteOptions{}
			}
			if err := m.DeleteOptions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PodDisruptionBudget) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PodDisruptionBudget: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PodDisruptionBudget: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PodDisruptionBudgetList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PodDisruptionBudgetList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PodDisruptionBudgetList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ListMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, PodDisruptionBudget{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PodDisruptionBudgetSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PodDisruptionBudgetSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PodDisruptionBudgetSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinAvailable", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MinAvailable == nil {
				m.MinAvailable = &k8s_io_apimachinery_pkg_util_intstr.IntOrString{}
			}
			if err := m.MinAvailable.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Selector", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Selector == nil {
				m.Selector = &k8s_io_apimachinery_pkg_apis_meta_v1.LabelSelector{}
			}
			if err := m.Selector.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxUnavailable", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MaxUnavailable == nil {
				m.MaxUnavailable = &k8s_io_apimachinery_pkg_util_intstr.IntOrString{}
			}
			if err := m.MaxUnavailable.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PodDisruptionBudgetStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PodDisruptionBudgetStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PodDisruptionBudgetStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObservedGeneration", wireType)
			}
			m.ObservedGeneration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ObservedGeneration |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisruptedPods", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapkey := int(stringLenmapkey)
			if intStringLenmapkey < 0 {
				return ErrInvalidLengthGenerated
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(dAtA[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			if m.DisruptedPods == nil {
				m.DisruptedPods = make(map[string]k8s_io_apimachinery_pkg_apis_meta_v1.Time)
			}
			if iNdEx < postIndex {
				var valuekey uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					valuekey |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				var mapmsglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					mapmsglen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if mapmsglen < 0 {
					return ErrInvalidLengthGenerated
				}
				postmsgIndex := iNdEx + mapmsglen
				if mapmsglen < 0 {
					return ErrInvalidLengthGenerated
				}
				if postmsgIndex > l {
					return io.ErrUnexpectedEOF
				}
				mapvalue := &k8s_io_apimachinery_pkg_apis_meta_v1.Time{}
				if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
					return err
				}
				iNdEx = postmsgIndex
				m.DisruptedPods[mapkey] = *mapvalue
			} else {
				var mapvalue k8s_io_apimachinery_pkg_apis_meta_v1.Time
				m.DisruptedPods[mapkey] = mapvalue
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PodDisruptionsAllowed", wireType)
			}
			m.PodDisruptionsAllowed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PodDisruptionsAllowed |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentHealthy", wireType)
			}
			m.CurrentHealthy = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurrentHealthy |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DesiredHealthy", wireType)
			}
			m.DesiredHealthy = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DesiredHealthy |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpectedPods", wireType)
			}
			m.ExpectedPods = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExpectedPods |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipGenerated(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("k8s.io/client-go/pkg/apis/policy/v1beta1/generated.proto", fileDescriptorGenerated)
}

var fileDescriptorGenerated = []byte{
	// 835 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0x4f, 0x6f, 0x1b, 0x45,
	0x18, 0xc6, 0xbd, 0x71, 0x1c, 0xc2, 0xd4, 0xb6, 0xc2, 0x40, 0x21, 0x58, 0xc2, 0x46, 0x3e, 0xb5,
	0x88, 0xce, 0x92, 0x16, 0xa1, 0xc0, 0xa1, 0xa2, 0x8b, 0x23, 0x5a, 0xd4, 0x28, 0xd5, 0x04, 0x84,
	0x84, 0x40, 0x62, 0xbc, 0xfb, 0x76, 0x3d, 0x78, 0xf6, 0x8f, 0x66, 0x66, 0x4d, 0x7c, 0xe3, 0x23,
	0x70, 0xe0, 0x43, 0x45, 0x42, 0x42, 0x3d, 0x22, 0x84, 0x2c, 0x62, 0x3e, 0x03, 0x77, 0xb4, 0xb3,
	0xe3, 0x78, 0x37, 0x76, 0x54, 0x43, 0x11, 0xb7, 0x9d, 0x3f, 0xbf, 0xe7, 0x99, 0xf7, 0x79, 0x67,
	0x07, 0x7d, 0x38, 0x3e, 0x54, 0x84, 0x27, 0xee, 0x38, 0x1b, 0x82, 0x8c, 0x41, 0x83, 0x72, 0xd3,
	0x71, 0xe8, 0xb2, 0x94, 0x2b, 0x37, 0x4d, 0x04, 0xf7, 0xa7, 0xee, 0xe4, 0x60, 0x08, 0x9a, 0x1d,
	0xb8, 0x21, 0xc4, 0x20, 0x99, 0x86, 0x80, 0xa4, 0x32, 0xd1, 0x09, 0xbe, 0x5d, 0xa0, 0x64, 0x89,
	0x92, 0x74, 0x1c, 0x92, 0x1c, 0x25, 0x05, 0x4a, 0x2c, 0xda, 0xb9, 0x13, 0x72, 0x3d, 0xca, 0x86,
	0xc4, 0x4f, 0x22, 0x37, 0x4c, 0xc2, 0xc4, 0x35, 0x0a, 0xc3, 0xec, 0xa9, 0x19, 0x99, 0x81, 0xf9,
	0x2a, 0x94, 0x3b, 0xef, 0xdb, 0x43, 0xb1, 0x94, 0x47, 0xcc, 0x1f, 0xf1, 0x18, 0xe4, 0x74, 0x79,
	0xac, 0x08, 0x34, 0x73, 0x27, 0x2b, 0xe7, 0xe9, 0xb8, 0xd7, 0x51, 0x32, 0x8b, 0x35, 0x8f, 0x60,
	0x05, 0xf8, 0xe0, 0x79, 0x80, 0xf2, 0x47, 0x10, 0xb1, 0x15, 0xee, 0xde, 0x75, 0x5c, 0xa6, 0xb9,
	0x70, 0x79, 0xac, 0x95, 0x96, 0x2b, 0x50, 0xa9, 0x26, 0x05, 0x72, 0x02, 0x72, 0x59, 0x10, 0x9c,
	0xb1, 0x28, 0x15, 0xb0, 0xae, 0xa6, 0x87, 0xa5, 0xf6, 0xdc, 0x61, 0x61, 0x28, 0x21, 0x64, 0x3a,
	0x29, 0xb1, 0x2c, 0xe5, 0x12, 0x42, 0xae, 0xb4, 0x64, 0x9a, 0x27, 0xb1, 0x3b, 0x39, 0x60, 0x22,
	0x1d, 0xad, 0x76, 0xab, 0xf3, 0xee, 0xb5, 0x8d, 0x5e, 0xe3, 0xdb, 0xff, 0xcd, 0x41, 0xbb, 0x47,
	0x13, 0xee, 0xe7, 0x9a, 0xf8, 0x5b, 0xb4, 0x9b, 0x67, 0x1e, 0x30, 0xcd, 0xf6, 0x9d, 0xb7, 0x9d,
	0x5b, 0x37, 0xee, 0xbe, 0x47, 0x6c, 0xef, 0xcb, 0x11, 0x2c, 0xbb, 0x9f, 0xef, 0x26, 0x93, 0x03,
	0x72, 0x32, 0xfc, 0x0e, 0x7c, 0x7d, 0x0c, 0x9a, 0x79, 0xf8, 0x7c, 0xd6, 0xab, 0xcd, 0x67, 0x3d,
	0xb4, 0x9c, 0xa3, 0x97, 0xaa, 0x58, 0xa0, 0x56, 0x00, 0x02, 0x34, 0x9c, 0xa4, 0xb9, 0xa3, 0xda,
	0xdf, 0x32, 0x36, 0xf7, 0x36, 0xb3, 0x19, 0x94, 0x51, 0xef, 0x95, 0xf9, 0xac, 0xd7, 0xaa, 0x4c,
	0xd1, 0xaa, 0x78, 0xff, 0xe7, 0x2d, 0xf4, 0xea, 0x93, 0x24, 0x18, 0x70, 0x25, 0x33, 0x33, 0xe5,
	0x65, 0x41, 0x08, 0xfa, 0x7f, 0xa8, 0x33, 0x40, 0xdb, 0x2a, 0x05, 0xdf, 0x96, 0xe7, 0x91, 0x8d,
	0xff, 0x20, 0xb2, 0xe6, 0xbc, 0xa7, 0x29, 0xf8, 0x5e, 0xd3, 0xfa, 0x6d, 0xe7, 0x23, 0x6a, 0xd4,
	0xb1, 0x40, 0x3b, 0x4a, 0x33, 0x9d, 0xa9, 0xfd, 0xba, 0xf1, 0x19, 0xbc, 0xa0, 0x8f, 0xd1, 0xf2,
	0xda, 0xd6, 0x69, 0xa7, 0x18, 0x53, 0xeb, 0xd1, 0xff, 0xdd, 0x41, 0x6f, 0xac, 0xa1, 0x1e, 0x73,
	0xa5, 0xf1, 0xd7, 0x2b, 0x89, 0x92, 0xcd, 0x12, 0xcd, 0x69, 0x93, 0xe7, 0x9e, 0x75, 0xdd, 0x5d,
	0xcc, 0x94, 0xd2, 0xf4, 0x51, 0x83, 0x6b, 0x88, 0xf2, 0xdb, 0x52, 0xbf, 0x75, 0xe3, 0xee, 0xfd,
	0x17, 0x2b, 0xd3, 0x6b, 0x59, 0xab, 0xc6, 0xa3, 0x5c, 0x94, 0x16, 0xda, 0xfd, 0x5f, 0xb6, 0xd6,
	0x96, 0x97, 0xc7, 0x8d, 0x9f, 0xa2, 0x66, 0xc4, 0xe3, 0x07, 0x13, 0xc6, 0x05, 0x1b, 0x0a, 0x78,
	0xee, 0xa5, 0xc9, 0xdf, 0x07, 0x52, 0xbc, 0x0f, 0xe4, 0x51, 0xac, 0x4f, 0xe4, 0xa9, 0x96, 0x3c,
	0x0e, 0xbd, 0xbd, 0xf9, 0xac, 0xd7, 0x3c, 0x2e, 0x29, 0xd1, 0x8a, 0x2e, 0xfe, 0x06, 0xed, 0x2a,
	0x10, 0xe0, 0xeb, 0x44, 0xfe, 0xb3, 0x3f, 0xe3, 0x31, 0x1b, 0x82, 0x38, 0xb5, 0xa8, 0xd7, 0xcc,
	0x73, 0x5c, 0x8c, 0xe8, 0xa5, 0x24, 0x16, 0xa8, 0x1d, 0xb1, 0xb3, 0x2f, 0x62, 0x76, 0x59, 0x48,
	0xfd, 0x5f, 0x16, 0x82, 0xe7, 0xb3, 0x5e, 0xfb, 0xb8, 0xa2, 0x45, 0xaf, 0x68, 0xf7, 0xff, 0xda,
	0x46, 0x6f, 0x5e, 0x7b, 0xcb, 0xf0, 0x67, 0x08, 0x27, 0x43, 0xf3, 0x40, 0x06, 0x9f, 0x16, 0x6f,
	0x12, 0x4f, 0x62, 0x13, 0x6c, 0xdd, 0xeb, 0xd8, 0x06, 0xe1, 0x93, 0x95, 0x1d, 0x74, 0x0d, 0x85,
	0x7f, 0x72, 0x50, 0x2b, 0x28, 0x6c, 0x20, 0x78, 0x92, 0x04, 0x8b, 0x8b, 0xf2, 0xe5, 0x7f, 0xf1,
	0x3f, 0x90, 0x41, 0x59, 0xf9, 0x28, 0xd6, 0x72, 0xea, 0xdd, 0xb4, 0x07, 0x6c, 0x55, 0xd6, 0x68,
	0xf5, 0x10, 0xf8, 0x18, 0xe1, 0xe0, 0x52, 0x52, 0x3d, 0x10, 0x22, 0xf9, 0x1e, 0x02, 0x13, 0x79,
	0xc3, 0x7b, 0xcb, 0x2a, 0xdc, 0xac, 0xf8, 0x2e, 0x36, 0xd1, 0x35, 0x20, 0xbe, 0x8f, 0xda, 0x7e,
	0x26, 0x25, 0xc4, 0xfa, 0x21, 0x30, 0xa1, 0x47, 0xd3, 0xfd, 0x6d, 0x23, 0xf5, 0xba, 0x95, 0x6a,
	0x7f, 0x52, 0x59, 0xa5, 0x57, 0x76, 0xe7, 0x7c, 0x00, 0x8a, 0x4b, 0x08, 0x16, 0x7c, 0xa3, 0xca,
	0x0f, 0x2a, 0xab, 0xf4, 0xca, 0x6e, 0x7c, 0x88, 0x9a, 0x70, 0x96, 0x82, 0xbf, 0xc8, 0x78, 0xc7,
	0xd0, 0xaf, 0x59, 0xba, 0x79, 0x54, 0x5a, 0xa3, 0x95, 0x9d, 0x1d, 0x81, 0xf0, 0x6a, 0x88, 0x78,
	0x0f, 0xd5, 0xc7, 0x30, 0x35, 0x2d, 0x7f, 0x99, 0xe6, 0x9f, 0xf8, 0x63, 0xd4, 0x98, 0x30, 0x91,
	0x81, 0xbd, 0xfb, 0xef, 0x6c, 0x76, 0xf7, 0x3f, 0xe7, 0x11, 0xd0, 0x02, 0xfc, 0x68, 0xeb, 0xd0,
	0xf1, 0x6e, 0x9f, 0x5f, 0x74, 0x6b, 0xcf, 0x2e, 0xba, 0xb5, 0x5f, 0x2f, 0xba, 0xb5, 0x1f, 0xe6,
	0x5d, 0xe7, 0x7c, 0xde, 0x75, 0x9e, 0xcd, 0xbb, 0xce, 0x1f, 0xf3, 0xae, 0xf3, 0xe3, 0x9f, 0xdd,
	0xda, 0x57, 0x2f, 0xd9, 0xa6, 0xff, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x4d, 0x19, 0x94, 0x39, 0x15,
	0x09, 0x00, 0x00,
}
