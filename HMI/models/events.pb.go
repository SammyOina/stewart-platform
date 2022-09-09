// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.11.4
// source: models/events.proto

package models

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IMUEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Yaw   float32 `protobuf:"fixed32,1,opt,name=yaw,proto3" json:"yaw,omitempty"`
	Pitch float32 `protobuf:"fixed32,2,opt,name=pitch,proto3" json:"pitch,omitempty"`
	Roll  float32 `protobuf:"fixed32,3,opt,name=roll,proto3" json:"roll,omitempty"`
}

func (x *IMUEvent) Reset() {
	*x = IMUEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IMUEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IMUEvent) ProtoMessage() {}

func (x *IMUEvent) ProtoReflect() protoreflect.Message {
	mi := &file_models_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IMUEvent.ProtoReflect.Descriptor instead.
func (*IMUEvent) Descriptor() ([]byte, []int) {
	return file_models_events_proto_rawDescGZIP(), []int{0}
}

func (x *IMUEvent) GetYaw() float32 {
	if x != nil {
		return x.Yaw
	}
	return 0
}

func (x *IMUEvent) GetPitch() float32 {
	if x != nil {
		return x.Pitch
	}
	return 0
}

func (x *IMUEvent) GetRoll() float32 {
	if x != nil {
		return x.Roll
	}
	return 0
}

type PitotEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IntakePitot      float32 `protobuf:"fixed32,1,opt,name=intakePitot,proto3" json:"intakePitot,omitempty"`
	TestSectionPitot float32 `protobuf:"fixed32,2,opt,name=testSectionPitot,proto3" json:"testSectionPitot,omitempty"`
	DiffuserPitot    float32 `protobuf:"fixed32,3,opt,name=diffuserPitot,proto3" json:"diffuserPitot,omitempty"`
}

func (x *PitotEvent) Reset() {
	*x = PitotEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PitotEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PitotEvent) ProtoMessage() {}

func (x *PitotEvent) ProtoReflect() protoreflect.Message {
	mi := &file_models_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PitotEvent.ProtoReflect.Descriptor instead.
func (*PitotEvent) Descriptor() ([]byte, []int) {
	return file_models_events_proto_rawDescGZIP(), []int{1}
}

func (x *PitotEvent) GetIntakePitot() float32 {
	if x != nil {
		return x.IntakePitot
	}
	return 0
}

func (x *PitotEvent) GetTestSectionPitot() float32 {
	if x != nil {
		return x.TestSectionPitot
	}
	return 0
}

func (x *PitotEvent) GetDiffuserPitot() float32 {
	if x != nil {
		return x.DiffuserPitot
	}
	return 0
}

type StrainEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Strain1 float32 `protobuf:"fixed32,1,opt,name=strain1,proto3" json:"strain1,omitempty"`
	Strain2 float32 `protobuf:"fixed32,2,opt,name=strain2,proto3" json:"strain2,omitempty"`
	Strain3 float32 `protobuf:"fixed32,3,opt,name=strain3,proto3" json:"strain3,omitempty"`
	Strain4 float32 `protobuf:"fixed32,4,opt,name=strain4,proto3" json:"strain4,omitempty"`
	Strain5 float32 `protobuf:"fixed32,5,opt,name=strain5,proto3" json:"strain5,omitempty"`
	Strain6 float32 `protobuf:"fixed32,6,opt,name=strain6,proto3" json:"strain6,omitempty"`
}

func (x *StrainEvent) Reset() {
	*x = StrainEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_events_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StrainEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StrainEvent) ProtoMessage() {}

func (x *StrainEvent) ProtoReflect() protoreflect.Message {
	mi := &file_models_events_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StrainEvent.ProtoReflect.Descriptor instead.
func (*StrainEvent) Descriptor() ([]byte, []int) {
	return file_models_events_proto_rawDescGZIP(), []int{2}
}

func (x *StrainEvent) GetStrain1() float32 {
	if x != nil {
		return x.Strain1
	}
	return 0
}

func (x *StrainEvent) GetStrain2() float32 {
	if x != nil {
		return x.Strain2
	}
	return 0
}

func (x *StrainEvent) GetStrain3() float32 {
	if x != nil {
		return x.Strain3
	}
	return 0
}

func (x *StrainEvent) GetStrain4() float32 {
	if x != nil {
		return x.Strain4
	}
	return 0
}

func (x *StrainEvent) GetStrain5() float32 {
	if x != nil {
		return x.Strain5
	}
	return 0
}

func (x *StrainEvent) GetStrain6() float32 {
	if x != nil {
		return x.Strain6
	}
	return 0
}

type ServoPositionEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Servo1 float32 `protobuf:"fixed32,1,opt,name=servo1,proto3" json:"servo1,omitempty"`
	Servo2 float32 `protobuf:"fixed32,2,opt,name=servo2,proto3" json:"servo2,omitempty"`
	Servo3 float32 `protobuf:"fixed32,3,opt,name=servo3,proto3" json:"servo3,omitempty"`
	Servo4 float32 `protobuf:"fixed32,4,opt,name=servo4,proto3" json:"servo4,omitempty"`
	Servo5 float32 `protobuf:"fixed32,5,opt,name=servo5,proto3" json:"servo5,omitempty"`
	Servo6 float32 `protobuf:"fixed32,6,opt,name=servo6,proto3" json:"servo6,omitempty"`
}

func (x *ServoPositionEvent) Reset() {
	*x = ServoPositionEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_events_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServoPositionEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServoPositionEvent) ProtoMessage() {}

func (x *ServoPositionEvent) ProtoReflect() protoreflect.Message {
	mi := &file_models_events_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServoPositionEvent.ProtoReflect.Descriptor instead.
func (*ServoPositionEvent) Descriptor() ([]byte, []int) {
	return file_models_events_proto_rawDescGZIP(), []int{3}
}

func (x *ServoPositionEvent) GetServo1() float32 {
	if x != nil {
		return x.Servo1
	}
	return 0
}

func (x *ServoPositionEvent) GetServo2() float32 {
	if x != nil {
		return x.Servo2
	}
	return 0
}

func (x *ServoPositionEvent) GetServo3() float32 {
	if x != nil {
		return x.Servo3
	}
	return 0
}

func (x *ServoPositionEvent) GetServo4() float32 {
	if x != nil {
		return x.Servo4
	}
	return 0
}

func (x *ServoPositionEvent) GetServo5() float32 {
	if x != nil {
		return x.Servo5
	}
	return 0
}

func (x *ServoPositionEvent) GetServo6() float32 {
	if x != nil {
		return x.Servo6
	}
	return 0
}

type SensorEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Event:
	//	*SensorEvent_IMUEvent
	//	*SensorEvent_StrainEvent
	//	*SensorEvent_PitotEvent
	Event isSensorEvent_Event `protobuf_oneof:"event"`
}

func (x *SensorEvent) Reset() {
	*x = SensorEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_events_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SensorEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SensorEvent) ProtoMessage() {}

func (x *SensorEvent) ProtoReflect() protoreflect.Message {
	mi := &file_models_events_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SensorEvent.ProtoReflect.Descriptor instead.
func (*SensorEvent) Descriptor() ([]byte, []int) {
	return file_models_events_proto_rawDescGZIP(), []int{4}
}

func (m *SensorEvent) GetEvent() isSensorEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *SensorEvent) GetIMUEvent() *IMUEvent {
	if x, ok := x.GetEvent().(*SensorEvent_IMUEvent); ok {
		return x.IMUEvent
	}
	return nil
}

func (x *SensorEvent) GetStrainEvent() *StrainEvent {
	if x, ok := x.GetEvent().(*SensorEvent_StrainEvent); ok {
		return x.StrainEvent
	}
	return nil
}

func (x *SensorEvent) GetPitotEvent() *PitotEvent {
	if x, ok := x.GetEvent().(*SensorEvent_PitotEvent); ok {
		return x.PitotEvent
	}
	return nil
}

type isSensorEvent_Event interface {
	isSensorEvent_Event()
}

type SensorEvent_IMUEvent struct {
	IMUEvent *IMUEvent `protobuf:"bytes,1,opt,name=iMUEvent,proto3,oneof"`
}

type SensorEvent_StrainEvent struct {
	StrainEvent *StrainEvent `protobuf:"bytes,2,opt,name=strainEvent,proto3,oneof"`
}

type SensorEvent_PitotEvent struct {
	PitotEvent *PitotEvent `protobuf:"bytes,3,opt,name=pitotEvent,proto3,oneof"`
}

func (*SensorEvent_IMUEvent) isSensorEvent_Event() {}

func (*SensorEvent_StrainEvent) isSensorEvent_Event() {}

func (*SensorEvent_PitotEvent) isSensorEvent_Event() {}

var File_models_events_proto protoreflect.FileDescriptor

var file_models_events_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x08, 0x49, 0x4d, 0x55, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x79, 0x61, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03,
	0x79, 0x61, 0x77, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x69, 0x74, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x05, 0x70, 0x69, 0x74, 0x63, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x6c, 0x22, 0x80, 0x01,
	0x0a, 0x0a, 0x50, 0x69, 0x74, 0x6f, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b,
	0x69, 0x6e, 0x74, 0x61, 0x6b, 0x65, 0x50, 0x69, 0x74, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0b, 0x69, 0x6e, 0x74, 0x61, 0x6b, 0x65, 0x50, 0x69, 0x74, 0x6f, 0x74, 0x12, 0x2a,
	0x0a, 0x10, 0x74, 0x65, 0x73, 0x74, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x69, 0x74,
	0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x10, 0x74, 0x65, 0x73, 0x74, 0x53, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x69, 0x74, 0x6f, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x64, 0x69,
	0x66, 0x66, 0x75, 0x73, 0x65, 0x72, 0x50, 0x69, 0x74, 0x6f, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0d, 0x64, 0x69, 0x66, 0x66, 0x75, 0x73, 0x65, 0x72, 0x50, 0x69, 0x74, 0x6f, 0x74,
	0x22, 0xa9, 0x01, 0x0a, 0x0b, 0x53, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x07, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x31, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x73, 0x74, 0x72,
	0x61, 0x69, 0x6e, 0x32, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x33, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x33, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x34, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x07, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x34, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x72, 0x61,
	0x69, 0x6e, 0x35, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x73, 0x74, 0x72, 0x61, 0x69,
	0x6e, 0x35, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x36, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x07, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x36, 0x22, 0xa4, 0x01, 0x0a,
	0x12, 0x53, 0x65, 0x72, 0x76, 0x6f, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x6f, 0x31, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x6f, 0x31, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x65, 0x72, 0x76, 0x6f, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x73, 0x65, 0x72,
	0x76, 0x6f, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x6f, 0x33, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x6f, 0x33, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x65, 0x72, 0x76, 0x6f, 0x34, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x73, 0x65, 0x72,
	0x76, 0x6f, 0x34, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x6f, 0x35, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x6f, 0x35, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x65, 0x72, 0x76, 0x6f, 0x36, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x73, 0x65, 0x72,
	0x76, 0x6f, 0x36, 0x22, 0xa0, 0x01, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x08, 0x69, 0x4d, 0x55, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x49, 0x4d, 0x55, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x48, 0x00, 0x52, 0x08, 0x69, 0x4d, 0x55, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x0b,
	0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48,
	0x00, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x2d,
	0x0a, 0x0a, 0x70, 0x69, 0x74, 0x6f, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x50, 0x69, 0x74, 0x6f, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48,
	0x00, 0x52, 0x0a, 0x70, 0x69, 0x74, 0x6f, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x07, 0x0a,
	0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_events_proto_rawDescOnce sync.Once
	file_models_events_proto_rawDescData = file_models_events_proto_rawDesc
)

func file_models_events_proto_rawDescGZIP() []byte {
	file_models_events_proto_rawDescOnce.Do(func() {
		file_models_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_events_proto_rawDescData)
	})
	return file_models_events_proto_rawDescData
}

var file_models_events_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_models_events_proto_goTypes = []interface{}{
	(*IMUEvent)(nil),           // 0: IMUEvent
	(*PitotEvent)(nil),         // 1: PitotEvent
	(*StrainEvent)(nil),        // 2: StrainEvent
	(*ServoPositionEvent)(nil), // 3: ServoPositionEvent
	(*SensorEvent)(nil),        // 4: SensorEvent
}
var file_models_events_proto_depIdxs = []int32{
	0, // 0: SensorEvent.iMUEvent:type_name -> IMUEvent
	2, // 1: SensorEvent.strainEvent:type_name -> StrainEvent
	1, // 2: SensorEvent.pitotEvent:type_name -> PitotEvent
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_models_events_proto_init() }
func file_models_events_proto_init() {
	if File_models_events_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IMUEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_models_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PitotEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_models_events_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StrainEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_models_events_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServoPositionEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_models_events_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SensorEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_models_events_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*SensorEvent_IMUEvent)(nil),
		(*SensorEvent_StrainEvent)(nil),
		(*SensorEvent_PitotEvent)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_models_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_events_proto_goTypes,
		DependencyIndexes: file_models_events_proto_depIdxs,
		MessageInfos:      file_models_events_proto_msgTypes,
	}.Build()
	File_models_events_proto = out.File
	file_models_events_proto_rawDesc = nil
	file_models_events_proto_goTypes = nil
	file_models_events_proto_depIdxs = nil
}
