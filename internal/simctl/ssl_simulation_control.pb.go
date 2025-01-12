// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.28.3
// source: ssl_simulation_control.proto

package simctl

import (
	referee "github.com/RoboCup-SSL/ssl-simulation-controller/internal/referee"
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

// Teleport the ball to a new location and optionally set it to some velocity
type TeleportBall struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// x-coordinate [m]
	X *float32 `protobuf:"fixed32,1,opt,name=x" json:"x,omitempty"`
	// y-coordinate [m]
	Y *float32 `protobuf:"fixed32,2,opt,name=y" json:"y,omitempty"`
	// z-coordinate (height) [m]
	Z *float32 `protobuf:"fixed32,3,opt,name=z" json:"z,omitempty"`
	// Velocity in x-direction [m/s]
	Vx *float32 `protobuf:"fixed32,4,opt,name=vx" json:"vx,omitempty"`
	// Velocity in y-direction [m/s]
	Vy *float32 `protobuf:"fixed32,5,opt,name=vy" json:"vy,omitempty"`
	// Velocity in z-direction [m/s]
	Vz *float32 `protobuf:"fixed32,6,opt,name=vz" json:"vz,omitempty"`
	// Teleport the ball safely to the target, for example by
	// moving robots out of the way in case of collision and set speed of robots close-by to zero
	TeleportSafely *bool `protobuf:"varint,7,opt,name=teleport_safely,json=teleportSafely,def=0" json:"teleport_safely,omitempty"`
	// Adapt the angular ball velocity such that the ball is rolling
	Roll          *bool `protobuf:"varint,8,opt,name=roll,def=0" json:"roll,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

// Default values for TeleportBall fields.
const (
	Default_TeleportBall_TeleportSafely = bool(false)
	Default_TeleportBall_Roll           = bool(false)
)

func (x *TeleportBall) Reset() {
	*x = TeleportBall{}
	mi := &file_ssl_simulation_control_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TeleportBall) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeleportBall) ProtoMessage() {}

func (x *TeleportBall) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_simulation_control_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeleportBall.ProtoReflect.Descriptor instead.
func (*TeleportBall) Descriptor() ([]byte, []int) {
	return file_ssl_simulation_control_proto_rawDescGZIP(), []int{0}
}

func (x *TeleportBall) GetX() float32 {
	if x != nil && x.X != nil {
		return *x.X
	}
	return 0
}

func (x *TeleportBall) GetY() float32 {
	if x != nil && x.Y != nil {
		return *x.Y
	}
	return 0
}

func (x *TeleportBall) GetZ() float32 {
	if x != nil && x.Z != nil {
		return *x.Z
	}
	return 0
}

func (x *TeleportBall) GetVx() float32 {
	if x != nil && x.Vx != nil {
		return *x.Vx
	}
	return 0
}

func (x *TeleportBall) GetVy() float32 {
	if x != nil && x.Vy != nil {
		return *x.Vy
	}
	return 0
}

func (x *TeleportBall) GetVz() float32 {
	if x != nil && x.Vz != nil {
		return *x.Vz
	}
	return 0
}

func (x *TeleportBall) GetTeleportSafely() bool {
	if x != nil && x.TeleportSafely != nil {
		return *x.TeleportSafely
	}
	return Default_TeleportBall_TeleportSafely
}

func (x *TeleportBall) GetRoll() bool {
	if x != nil && x.Roll != nil {
		return *x.Roll
	}
	return Default_TeleportBall_Roll
}

// Teleport a robot to some location and give it a velocity
type TeleportRobot struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Robot id to teleport
	Id *referee.RobotId `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	// x-coordinate [m]
	X *float32 `protobuf:"fixed32,2,opt,name=x" json:"x,omitempty"`
	// y-coordinate [m]
	Y *float32 `protobuf:"fixed32,3,opt,name=y" json:"y,omitempty"`
	// Orientation [rad], measured from the x-axis counter-clockwise
	Orientation *float32 `protobuf:"fixed32,4,opt,name=orientation" json:"orientation,omitempty"`
	// Global velocity [m/s] towards x-axis
	VX *float32 `protobuf:"fixed32,5,opt,name=v_x,json=vX,def=0" json:"v_x,omitempty"`
	// Global velocity [m/s] towards y-axis
	VY *float32 `protobuf:"fixed32,6,opt,name=v_y,json=vY,def=0" json:"v_y,omitempty"`
	// Angular velocity [rad/s]
	VAngular *float32 `protobuf:"fixed32,7,opt,name=v_angular,json=vAngular,def=0" json:"v_angular,omitempty"`
	// Robot should be present on the field?
	// true -> robot will be added, if it does not exist yet
	// false -> robot will be removed, if it is present
	Present       *bool `protobuf:"varint,8,opt,name=present" json:"present,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

// Default values for TeleportRobot fields.
const (
	Default_TeleportRobot_VX       = float32(0)
	Default_TeleportRobot_VY       = float32(0)
	Default_TeleportRobot_VAngular = float32(0)
)

func (x *TeleportRobot) Reset() {
	*x = TeleportRobot{}
	mi := &file_ssl_simulation_control_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TeleportRobot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeleportRobot) ProtoMessage() {}

func (x *TeleportRobot) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_simulation_control_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeleportRobot.ProtoReflect.Descriptor instead.
func (*TeleportRobot) Descriptor() ([]byte, []int) {
	return file_ssl_simulation_control_proto_rawDescGZIP(), []int{1}
}

func (x *TeleportRobot) GetId() *referee.RobotId {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *TeleportRobot) GetX() float32 {
	if x != nil && x.X != nil {
		return *x.X
	}
	return 0
}

func (x *TeleportRobot) GetY() float32 {
	if x != nil && x.Y != nil {
		return *x.Y
	}
	return 0
}

func (x *TeleportRobot) GetOrientation() float32 {
	if x != nil && x.Orientation != nil {
		return *x.Orientation
	}
	return 0
}

func (x *TeleportRobot) GetVX() float32 {
	if x != nil && x.VX != nil {
		return *x.VX
	}
	return Default_TeleportRobot_VX
}

func (x *TeleportRobot) GetVY() float32 {
	if x != nil && x.VY != nil {
		return *x.VY
	}
	return Default_TeleportRobot_VY
}

func (x *TeleportRobot) GetVAngular() float32 {
	if x != nil && x.VAngular != nil {
		return *x.VAngular
	}
	return Default_TeleportRobot_VAngular
}

func (x *TeleportRobot) GetPresent() bool {
	if x != nil && x.Present != nil {
		return *x.Present
	}
	return false
}

// Control the simulation
type SimulatorControl struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Teleport the ball
	TeleportBall *TeleportBall `protobuf:"bytes,1,opt,name=teleport_ball,json=teleportBall" json:"teleport_ball,omitempty"`
	// Teleport robots
	TeleportRobot []*TeleportRobot `protobuf:"bytes,2,rep,name=teleport_robot,json=teleportRobot" json:"teleport_robot,omitempty"`
	// Change the simulation speed
	SimulationSpeed *float32 `protobuf:"fixed32,3,opt,name=simulation_speed,json=simulationSpeed" json:"simulation_speed,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *SimulatorControl) Reset() {
	*x = SimulatorControl{}
	mi := &file_ssl_simulation_control_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SimulatorControl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimulatorControl) ProtoMessage() {}

func (x *SimulatorControl) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_simulation_control_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimulatorControl.ProtoReflect.Descriptor instead.
func (*SimulatorControl) Descriptor() ([]byte, []int) {
	return file_ssl_simulation_control_proto_rawDescGZIP(), []int{2}
}

func (x *SimulatorControl) GetTeleportBall() *TeleportBall {
	if x != nil {
		return x.TeleportBall
	}
	return nil
}

func (x *SimulatorControl) GetTeleportRobot() []*TeleportRobot {
	if x != nil {
		return x.TeleportRobot
	}
	return nil
}

func (x *SimulatorControl) GetSimulationSpeed() float32 {
	if x != nil && x.SimulationSpeed != nil {
		return *x.SimulationSpeed
	}
	return 0
}

// Command from the connected client to the simulator
type SimulatorCommand struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Control the simulation
	Control *SimulatorControl `protobuf:"bytes,1,opt,name=control" json:"control,omitempty"`
	// Configure the simulation
	Config        *SimulatorConfig `protobuf:"bytes,2,opt,name=config" json:"config,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SimulatorCommand) Reset() {
	*x = SimulatorCommand{}
	mi := &file_ssl_simulation_control_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SimulatorCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimulatorCommand) ProtoMessage() {}

func (x *SimulatorCommand) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_simulation_control_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimulatorCommand.ProtoReflect.Descriptor instead.
func (*SimulatorCommand) Descriptor() ([]byte, []int) {
	return file_ssl_simulation_control_proto_rawDescGZIP(), []int{3}
}

func (x *SimulatorCommand) GetControl() *SimulatorControl {
	if x != nil {
		return x.Control
	}
	return nil
}

func (x *SimulatorCommand) GetConfig() *SimulatorConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

// Response of the simulator to the connected client
type SimulatorResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// List of errors, like using unsupported features
	Errors        []*SimulatorError `protobuf:"bytes,1,rep,name=errors" json:"errors,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SimulatorResponse) Reset() {
	*x = SimulatorResponse{}
	mi := &file_ssl_simulation_control_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SimulatorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimulatorResponse) ProtoMessage() {}

func (x *SimulatorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_simulation_control_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimulatorResponse.ProtoReflect.Descriptor instead.
func (*SimulatorResponse) Descriptor() ([]byte, []int) {
	return file_ssl_simulation_control_proto_rawDescGZIP(), []int{4}
}

func (x *SimulatorResponse) GetErrors() []*SimulatorError {
	if x != nil {
		return x.Errors
	}
	return nil
}

var File_ssl_simulation_control_proto protoreflect.FileDescriptor

var file_ssl_simulation_control_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x73, 0x73, 0x6c, 0x5f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d,
	0x65, 0x64, 0x75, 0x2e, 0x74, 0x69, 0x67, 0x65, 0x72, 0x73, 0x2e, 0x73, 0x75, 0x6d, 0x61, 0x74,
	0x72, 0x61, 0x2e, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x13, 0x73,
	0x73, 0x6c, 0x5f, 0x67, 0x63, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x73, 0x73, 0x6c, 0x5f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1a, 0x73, 0x73, 0x6c, 0x5f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x01, 0x0a, 0x0c,
	0x54, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x61, 0x6c, 0x6c, 0x12, 0x0c, 0x0a, 0x01,
	0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x7a, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x01, 0x7a, 0x12, 0x0e, 0x0a, 0x02, 0x76, 0x78, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x02, 0x76, 0x78, 0x12, 0x0e, 0x0a, 0x02, 0x76, 0x79, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x02, 0x76, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x76, 0x7a, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x02, 0x76, 0x7a, 0x12, 0x2e, 0x0a, 0x0f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x5f, 0x73, 0x61, 0x66, 0x65, 0x6c, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x3a,
	0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x52, 0x0e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x53, 0x61, 0x66, 0x65, 0x6c, 0x79, 0x12, 0x19, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x6c, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c,
	0x6c, 0x22, 0xc9, 0x01, 0x0a, 0x0d, 0x54, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x6f,
	0x62, 0x6f, 0x74, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32,
	0x08, 0x2e, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x49, 0x64, 0x52, 0x02, 0x69, 0x64, 0x12, 0x0c, 0x0a,
	0x01, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x72, 0x69,
	0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b,
	0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x03, 0x76,
	0x5f, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x3a, 0x01, 0x30, 0x52, 0x02, 0x76, 0x58, 0x12,
	0x12, 0x0a, 0x03, 0x76, 0x5f, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x3a, 0x01, 0x30, 0x52,
	0x02, 0x76, 0x59, 0x12, 0x1e, 0x0a, 0x09, 0x76, 0x5f, 0x61, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x3a, 0x01, 0x30, 0x52, 0x08, 0x76, 0x41, 0x6e, 0x67, 0x75,
	0x6c, 0x61, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x22, 0xe4, 0x01,
	0x0a, 0x10, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x12, 0x50, 0x0a, 0x0d, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x62,
	0x61, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x65, 0x64, 0x75, 0x2e,
	0x74, 0x69, 0x67, 0x65, 0x72, 0x73, 0x2e, 0x73, 0x75, 0x6d, 0x61, 0x74, 0x72, 0x61, 0x2e, 0x73,
	0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x65, 0x6c, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x42, 0x61, 0x6c, 0x6c, 0x52, 0x0c, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x42, 0x61, 0x6c, 0x6c, 0x12, 0x53, 0x0a, 0x0e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x5f, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x65,
	0x64, 0x75, 0x2e, 0x74, 0x69, 0x67, 0x65, 0x72, 0x73, 0x2e, 0x73, 0x75, 0x6d, 0x61, 0x74, 0x72,
	0x61, 0x2e, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x65, 0x6c,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x52, 0x0d, 0x74, 0x65, 0x6c, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x69, 0x6d,
	0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x0f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x70, 0x65, 0x65, 0x64, 0x22, 0x87, 0x01, 0x0a, 0x10, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74,
	0x6f, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x49, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x65, 0x64, 0x75,
	0x2e, 0x74, 0x69, 0x67, 0x65, 0x72, 0x73, 0x2e, 0x73, 0x75, 0x6d, 0x61, 0x74, 0x72, 0x61, 0x2e,
	0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x69, 0x6d, 0x75, 0x6c,
	0x61, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x12, 0x28, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x5a,
	0x0a, 0x11, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x65, 0x64, 0x75, 0x2e, 0x74, 0x69, 0x67, 0x65, 0x72, 0x73,
	0x2e, 0x73, 0x75, 0x6d, 0x61, 0x74, 0x72, 0x61, 0x2e, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x6f, 0x62, 0x6f, 0x43, 0x75, 0x70,
	0x2d, 0x53, 0x53, 0x4c, 0x2f, 0x73, 0x73, 0x6c, 0x2d, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x69, 0x6d, 0x63, 0x74, 0x6c,
}

var (
	file_ssl_simulation_control_proto_rawDescOnce sync.Once
	file_ssl_simulation_control_proto_rawDescData = file_ssl_simulation_control_proto_rawDesc
)

func file_ssl_simulation_control_proto_rawDescGZIP() []byte {
	file_ssl_simulation_control_proto_rawDescOnce.Do(func() {
		file_ssl_simulation_control_proto_rawDescData = protoimpl.X.CompressGZIP(file_ssl_simulation_control_proto_rawDescData)
	})
	return file_ssl_simulation_control_proto_rawDescData
}

var file_ssl_simulation_control_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_ssl_simulation_control_proto_goTypes = []any{
	(*TeleportBall)(nil),      // 0: edu.tigers.sumatra.simulation.TeleportBall
	(*TeleportRobot)(nil),     // 1: edu.tigers.sumatra.simulation.TeleportRobot
	(*SimulatorControl)(nil),  // 2: edu.tigers.sumatra.simulation.SimulatorControl
	(*SimulatorCommand)(nil),  // 3: edu.tigers.sumatra.simulation.SimulatorCommand
	(*SimulatorResponse)(nil), // 4: edu.tigers.sumatra.simulation.SimulatorResponse
	(*referee.RobotId)(nil),   // 5: RobotId
	(*SimulatorConfig)(nil),   // 6: SimulatorConfig
	(*SimulatorError)(nil),    // 7: edu.tigers.sumatra.simulation.SimulatorError
}
var file_ssl_simulation_control_proto_depIdxs = []int32{
	5, // 0: edu.tigers.sumatra.simulation.TeleportRobot.id:type_name -> RobotId
	0, // 1: edu.tigers.sumatra.simulation.SimulatorControl.teleport_ball:type_name -> edu.tigers.sumatra.simulation.TeleportBall
	1, // 2: edu.tigers.sumatra.simulation.SimulatorControl.teleport_robot:type_name -> edu.tigers.sumatra.simulation.TeleportRobot
	2, // 3: edu.tigers.sumatra.simulation.SimulatorCommand.control:type_name -> edu.tigers.sumatra.simulation.SimulatorControl
	6, // 4: edu.tigers.sumatra.simulation.SimulatorCommand.config:type_name -> SimulatorConfig
	7, // 5: edu.tigers.sumatra.simulation.SimulatorResponse.errors:type_name -> edu.tigers.sumatra.simulation.SimulatorError
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_ssl_simulation_control_proto_init() }
func file_ssl_simulation_control_proto_init() {
	if File_ssl_simulation_control_proto != nil {
		return
	}
	file_ssl_simulation_config_proto_init()
	file_ssl_simulation_error_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ssl_simulation_control_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ssl_simulation_control_proto_goTypes,
		DependencyIndexes: file_ssl_simulation_control_proto_depIdxs,
		MessageInfos:      file_ssl_simulation_control_proto_msgTypes,
	}.Build()
	File_ssl_simulation_control_proto = out.File
	file_ssl_simulation_control_proto_rawDesc = nil
	file_ssl_simulation_control_proto_goTypes = nil
	file_ssl_simulation_control_proto_depIdxs = nil
}
