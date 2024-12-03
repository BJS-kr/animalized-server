// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.29.0--rc1
// source: message/proto/input.proto

package message

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

type Operation_Direction int32

const (
	Operation_DIRECTION_UNSPECIFIED Operation_Direction = 0
	Operation_UP                    Operation_Direction = 1
	Operation_DOWN                  Operation_Direction = 2
	Operation_LEFT                  Operation_Direction = 3
	Operation_RIGHT                 Operation_Direction = 4
)

// Enum value maps for Operation_Direction.
var (
	Operation_Direction_name = map[int32]string{
		0: "DIRECTION_UNSPECIFIED",
		1: "UP",
		2: "DOWN",
		3: "LEFT",
		4: "RIGHT",
	}
	Operation_Direction_value = map[string]int32{
		"DIRECTION_UNSPECIFIED": 0,
		"UP":                    1,
		"DOWN":                  2,
		"LEFT":                  3,
		"RIGHT":                 4,
	}
)

func (x Operation_Direction) Enum() *Operation_Direction {
	p := new(Operation_Direction)
	*p = x
	return p
}

func (x Operation_Direction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Operation_Direction) Descriptor() protoreflect.EnumDescriptor {
	return file_message_proto_input_proto_enumTypes[0].Descriptor()
}

func (Operation_Direction) Type() protoreflect.EnumType {
	return &file_message_proto_input_proto_enumTypes[0]
}

func (x Operation_Direction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Operation_Direction.Descriptor instead.
func (Operation_Direction) EnumDescriptor() ([]byte, []int) {
	return file_message_proto_input_proto_rawDescGZIP(), []int{2, 0}
}

type Operation_OperationType int32

const (
	Operation_OPERATION_UNSPECIFIED Operation_OperationType = 0
	Operation_MOVE                  Operation_OperationType = 1
	Operation_ATTACK                Operation_OperationType = 2
	Operation_HIT                   Operation_OperationType = 3
	Operation_GAME_STATE            Operation_OperationType = 4
)

// Enum value maps for Operation_OperationType.
var (
	Operation_OperationType_name = map[int32]string{
		0: "OPERATION_UNSPECIFIED",
		1: "MOVE",
		2: "ATTACK",
		3: "HIT",
		4: "GAME_STATE",
	}
	Operation_OperationType_value = map[string]int32{
		"OPERATION_UNSPECIFIED": 0,
		"MOVE":                  1,
		"ATTACK":                2,
		"HIT":                   3,
		"GAME_STATE":            4,
	}
)

func (x Operation_OperationType) Enum() *Operation_OperationType {
	p := new(Operation_OperationType)
	*p = x
	return p
}

func (x Operation_OperationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Operation_OperationType) Descriptor() protoreflect.EnumDescriptor {
	return file_message_proto_input_proto_enumTypes[1].Descriptor()
}

func (Operation_OperationType) Type() protoreflect.EnumType {
	return &file_message_proto_input_proto_enumTypes[1]
}

func (x Operation_OperationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Operation_OperationType.Descriptor instead.
func (Operation_OperationType) EnumDescriptor() ([]byte, []int) {
	return file_message_proto_input_proto_rawDescGZIP(), []int{2, 1}
}

type RoomState_RoomStatusType int32

const (
	RoomState_ROOM_STATUS_UNSPECIFIED RoomState_RoomStatusType = 0
	RoomState_WAITING                 RoomState_RoomStatusType = 1
	RoomState_PLAYING                 RoomState_RoomStatusType = 2
)

// Enum value maps for RoomState_RoomStatusType.
var (
	RoomState_RoomStatusType_name = map[int32]string{
		0: "ROOM_STATUS_UNSPECIFIED",
		1: "WAITING",
		2: "PLAYING",
	}
	RoomState_RoomStatusType_value = map[string]int32{
		"ROOM_STATUS_UNSPECIFIED": 0,
		"WAITING":                 1,
		"PLAYING":                 2,
	}
)

func (x RoomState_RoomStatusType) Enum() *RoomState_RoomStatusType {
	p := new(RoomState_RoomStatusType)
	*p = x
	return p
}

func (x RoomState_RoomStatusType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RoomState_RoomStatusType) Descriptor() protoreflect.EnumDescriptor {
	return file_message_proto_input_proto_enumTypes[2].Descriptor()
}

func (RoomState_RoomStatusType) Type() protoreflect.EnumType {
	return &file_message_proto_input_proto_enumTypes[2]
}

func (x RoomState_RoomStatusType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RoomState_RoomStatusType.Descriptor instead.
func (RoomState_RoomStatusType) EnumDescriptor() ([]byte, []int) {
	return file_message_proto_input_proto_rawDescGZIP(), []int{4, 0}
}

type Lobby_LobbyType int32

const (
	Lobby_LOBBY_UNSPECIFIED Lobby_LobbyType = 0
	Lobby_CREATE_ROOM       Lobby_LobbyType = 1
	Lobby_JOIN_ROOM         Lobby_LobbyType = 2
	Lobby_QUIT_ROOM         Lobby_LobbyType = 3
	Lobby_STATE             Lobby_LobbyType = 4
)

// Enum value maps for Lobby_LobbyType.
var (
	Lobby_LobbyType_name = map[int32]string{
		0: "LOBBY_UNSPECIFIED",
		1: "CREATE_ROOM",
		2: "JOIN_ROOM",
		3: "QUIT_ROOM",
		4: "STATE",
	}
	Lobby_LobbyType_value = map[string]int32{
		"LOBBY_UNSPECIFIED": 0,
		"CREATE_ROOM":       1,
		"JOIN_ROOM":         2,
		"QUIT_ROOM":         3,
		"STATE":             4,
	}
)

func (x Lobby_LobbyType) Enum() *Lobby_LobbyType {
	p := new(Lobby_LobbyType)
	*p = x
	return p
}

func (x Lobby_LobbyType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Lobby_LobbyType) Descriptor() protoreflect.EnumDescriptor {
	return file_message_proto_input_proto_enumTypes[3].Descriptor()
}

func (Lobby_LobbyType) Type() protoreflect.EnumType {
	return &file_message_proto_input_proto_enumTypes[3]
}

func (x Lobby_LobbyType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Lobby_LobbyType.Descriptor instead.
func (Lobby_LobbyType) EnumDescriptor() ([]byte, []int) {
	return file_message_proto_input_proto_rawDescGZIP(), []int{5, 0}
}

type Room_RoomType int32

const (
	Room_ROOM_UNSPECIFIED Room_RoomType = 0
	Room_QUIT             Room_RoomType = 1
	Room_START            Room_RoomType = 2
	Room_STATE            Room_RoomType = 3
)

// Enum value maps for Room_RoomType.
var (
	Room_RoomType_name = map[int32]string{
		0: "ROOM_UNSPECIFIED",
		1: "QUIT",
		2: "START",
		3: "STATE",
	}
	Room_RoomType_value = map[string]int32{
		"ROOM_UNSPECIFIED": 0,
		"QUIT":             1,
		"START":            2,
		"STATE":            3,
	}
)

func (x Room_RoomType) Enum() *Room_RoomType {
	p := new(Room_RoomType)
	*p = x
	return p
}

func (x Room_RoomType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Room_RoomType) Descriptor() protoreflect.EnumDescriptor {
	return file_message_proto_input_proto_enumTypes[4].Descriptor()
}

func (Room_RoomType) Type() protoreflect.EnumType {
	return &file_message_proto_input_proto_enumTypes[4]
}

func (x Room_RoomType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Room_RoomType.Descriptor instead.
func (Room_RoomType) EnumDescriptor() ([]byte, []int) {
	return file_message_proto_input_proto_rawDescGZIP(), []int{6, 0}
}

type Input struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// Types that are assignable to Kind:
	//
	//	*Input_Init
	//	*Input_Op
	//	*Input_Lobby
	//	*Input_Room
	Kind isInput_Kind `protobuf_oneof:"kind"`
}

func (x *Input) Reset() {
	*x = Input{}
	mi := &file_message_proto_input_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Input) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Input) ProtoMessage() {}

func (x *Input) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_input_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Input.ProtoReflect.Descriptor instead.
func (*Input) Descriptor() ([]byte, []int) {
	return file_message_proto_input_proto_rawDescGZIP(), []int{0}
}

func (x *Input) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (m *Input) GetKind() isInput_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (x *Input) GetInit() *Init {
	if x, ok := x.GetKind().(*Input_Init); ok {
		return x.Init
	}
	return nil
}

func (x *Input) GetOp() *Operation {
	if x, ok := x.GetKind().(*Input_Op); ok {
		return x.Op
	}
	return nil
}

func (x *Input) GetLobby() *Lobby {
	if x, ok := x.GetKind().(*Input_Lobby); ok {
		return x.Lobby
	}
	return nil
}

func (x *Input) GetRoom() *Room {
	if x, ok := x.GetKind().(*Input_Room); ok {
		return x.Room
	}
	return nil
}

type isInput_Kind interface {
	isInput_Kind()
}

type Input_Init struct {
	Init *Init `protobuf:"bytes,2,opt,name=init,proto3,oneof"`
}

type Input_Op struct {
	Op *Operation `protobuf:"bytes,3,opt,name=op,proto3,oneof"`
}

type Input_Lobby struct {
	Lobby *Lobby `protobuf:"bytes,4,opt,name=lobby,proto3,oneof"`
}

type Input_Room struct {
	Room *Room `protobuf:"bytes,5,opt,name=room,proto3,oneof"`
}

func (*Input_Init) isInput_Kind() {}

func (*Input_Op) isInput_Kind() {}

func (*Input_Lobby) isInput_Kind() {}

func (*Input_Room) isInput_Kind() {}

type Position struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y int32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Position) Reset() {
	*x = Position{}
	mi := &file_message_proto_input_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Position) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Position) ProtoMessage() {}

func (x *Position) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_input_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Position.ProtoReflect.Descriptor instead.
func (*Position) Descriptor() ([]byte, []int) {
	return file_message_proto_input_proto_rawDescGZIP(), []int{1}
}

func (x *Position) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Position) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type Operation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type         Operation_OperationType `protobuf:"varint,1,opt,name=type,proto3,enum=input.Operation_OperationType" json:"type,omitempty"`
	Direction    Operation_Direction     `protobuf:"varint,2,opt,name=direction,proto3,enum=input.Operation_Direction" json:"direction,omitempty"`
	HitRange     *Operation_HitRange     `protobuf:"bytes,3,opt,name=hit_range,json=hitRange,proto3" json:"hit_range,omitempty"`
	GameState    *Operation_GameState    `protobuf:"bytes,4,opt,name=game_state,json=gameState,proto3" json:"game_state,omitempty"`
	TargetUserId string                  `protobuf:"bytes,5,opt,name=target_user_id,json=targetUserId,proto3" json:"target_user_id,omitempty"`
	ProjectileId int32                   `protobuf:"varint,6,opt,name=projectile_id,json=projectileId,proto3" json:"projectile_id,omitempty"`
	Context      int64                   `protobuf:"varint,7,opt,name=context,proto3" json:"context,omitempty"`
	PrevContext  int64                   `protobuf:"varint,8,opt,name=prev_context,json=prevContext,proto3" json:"prev_context,omitempty"`
}

func (x *Operation) Reset() {
	*x = Operation{}
	mi := &file_message_proto_input_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Operation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operation) ProtoMessage() {}

func (x *Operation) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_input_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operation.ProtoReflect.Descriptor instead.
func (*Operation) Descriptor() ([]byte, []int) {
	return file_message_proto_input_proto_rawDescGZIP(), []int{2}
}

func (x *Operation) GetType() Operation_OperationType {
	if x != nil {
		return x.Type
	}
	return Operation_OPERATION_UNSPECIFIED
}

func (x *Operation) GetDirection() Operation_Direction {
	if x != nil {
		return x.Direction
	}
	return Operation_DIRECTION_UNSPECIFIED
}

func (x *Operation) GetHitRange() *Operation_HitRange {
	if x != nil {
		return x.HitRange
	}
	return nil
}

func (x *Operation) GetGameState() *Operation_GameState {
	if x != nil {
		return x.GameState
	}
	return nil
}

func (x *Operation) GetTargetUserId() string {
	if x != nil {
		return x.TargetUserId
	}
	return ""
}

func (x *Operation) GetProjectileId() int32 {
	if x != nil {
		return x.ProjectileId
	}
	return 0
}

func (x *Operation) GetContext() int64 {
	if x != nil {
		return x.Context
	}
	return 0
}

func (x *Operation) GetPrevContext() int64 {
	if x != nil {
		return x.PrevContext
	}
	return 0
}

type Init struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Init) Reset() {
	*x = Init{}
	mi := &file_message_proto_input_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Init) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Init) ProtoMessage() {}

func (x *Init) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_input_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Init.ProtoReflect.Descriptor instead.
func (*Init) Descriptor() ([]byte, []int) {
	return file_message_proto_input_proto_rawDescGZIP(), []int{3}
}

type RoomState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   RoomState_RoomStatusType `protobuf:"varint,1,opt,name=status,proto3,enum=input.RoomState_RoomStatusType" json:"status,omitempty"`
	RoomName string                   `protobuf:"bytes,2,opt,name=room_name,json=roomName,proto3" json:"room_name,omitempty"`
	MaxUsers int32                    `protobuf:"varint,3,opt,name=max_users,json=maxUsers,proto3" json:"max_users,omitempty"`
	UserIds  []string                 `protobuf:"bytes,4,rep,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
}

func (x *RoomState) Reset() {
	*x = RoomState{}
	mi := &file_message_proto_input_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoomState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomState) ProtoMessage() {}

func (x *RoomState) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_input_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomState.ProtoReflect.Descriptor instead.
func (*RoomState) Descriptor() ([]byte, []int) {
	return file_message_proto_input_proto_rawDescGZIP(), []int{4}
}

func (x *RoomState) GetStatus() RoomState_RoomStatusType {
	if x != nil {
		return x.Status
	}
	return RoomState_ROOM_STATUS_UNSPECIFIED
}

func (x *RoomState) GetRoomName() string {
	if x != nil {
		return x.RoomName
	}
	return ""
}

func (x *RoomState) GetMaxUsers() int32 {
	if x != nil {
		return x.MaxUsers
	}
	return 0
}

func (x *RoomState) GetUserIds() []string {
	if x != nil {
		return x.UserIds
	}
	return nil
}

type Lobby struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type       Lobby_LobbyType `protobuf:"varint,1,opt,name=type,proto3,enum=input.Lobby_LobbyType" json:"type,omitempty"`
	RoomName   string          `protobuf:"bytes,2,opt,name=room_name,json=roomName,proto3" json:"room_name,omitempty"`
	MaxUsers   int32           `protobuf:"varint,3,opt,name=max_users,json=maxUsers,proto3" json:"max_users,omitempty"`
	RoomStates []*RoomState    `protobuf:"bytes,4,rep,name=room_states,json=roomStates,proto3" json:"room_states,omitempty"`
}

func (x *Lobby) Reset() {
	*x = Lobby{}
	mi := &file_message_proto_input_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Lobby) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lobby) ProtoMessage() {}

func (x *Lobby) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_input_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lobby.ProtoReflect.Descriptor instead.
func (*Lobby) Descriptor() ([]byte, []int) {
	return file_message_proto_input_proto_rawDescGZIP(), []int{5}
}

func (x *Lobby) GetType() Lobby_LobbyType {
	if x != nil {
		return x.Type
	}
	return Lobby_LOBBY_UNSPECIFIED
}

func (x *Lobby) GetRoomName() string {
	if x != nil {
		return x.RoomName
	}
	return ""
}

func (x *Lobby) GetMaxUsers() int32 {
	if x != nil {
		return x.MaxUsers
	}
	return 0
}

func (x *Lobby) GetRoomStates() []*RoomState {
	if x != nil {
		return x.RoomStates
	}
	return nil
}

type Room struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      Room_RoomType `protobuf:"varint,1,opt,name=type,proto3,enum=input.Room_RoomType" json:"type,omitempty"`
	RoomName  string        `protobuf:"bytes,2,opt,name=room_name,json=roomName,proto3" json:"room_name,omitempty"`
	RoomState *RoomState    `protobuf:"bytes,4,opt,name=room_state,json=roomState,proto3" json:"room_state,omitempty"`
}

func (x *Room) Reset() {
	*x = Room{}
	mi := &file_message_proto_input_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Room) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Room) ProtoMessage() {}

func (x *Room) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_input_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Room.ProtoReflect.Descriptor instead.
func (*Room) Descriptor() ([]byte, []int) {
	return file_message_proto_input_proto_rawDescGZIP(), []int{6}
}

func (x *Room) GetType() Room_RoomType {
	if x != nil {
		return x.Type
	}
	return Room_ROOM_UNSPECIFIED
}

func (x *Room) GetRoomName() string {
	if x != nil {
		return x.RoomName
	}
	return ""
}

func (x *Room) GetRoomState() *RoomState {
	if x != nil {
		return x.RoomState
	}
	return nil
}

type Operation_HitRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LeftBottom *Position `protobuf:"bytes,1,opt,name=left_bottom,json=leftBottom,proto3" json:"left_bottom,omitempty"`
	RightTop   *Position `protobuf:"bytes,2,opt,name=right_top,json=rightTop,proto3" json:"right_top,omitempty"`
}

func (x *Operation_HitRange) Reset() {
	*x = Operation_HitRange{}
	mi := &file_message_proto_input_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Operation_HitRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operation_HitRange) ProtoMessage() {}

func (x *Operation_HitRange) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_input_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operation_HitRange.ProtoReflect.Descriptor instead.
func (*Operation_HitRange) Descriptor() ([]byte, []int) {
	return file_message_proto_input_proto_rawDescGZIP(), []int{2, 0}
}

func (x *Operation_HitRange) GetLeftBottom() *Position {
	if x != nil {
		return x.LeftBottom
	}
	return nil
}

func (x *Operation_HitRange) GetRightTop() *Position {
	if x != nil {
		return x.RightTop
	}
	return nil
}

type Operation_GameState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserStates []*Operation_GameState_UserState `protobuf:"bytes,1,rep,name=user_states,json=userStates,proto3" json:"user_states,omitempty"`
}

func (x *Operation_GameState) Reset() {
	*x = Operation_GameState{}
	mi := &file_message_proto_input_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Operation_GameState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operation_GameState) ProtoMessage() {}

func (x *Operation_GameState) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_input_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operation_GameState.ProtoReflect.Descriptor instead.
func (*Operation_GameState) Descriptor() ([]byte, []int) {
	return file_message_proto_input_proto_rawDescGZIP(), []int{2, 1}
}

func (x *Operation_GameState) GetUserStates() []*Operation_GameState_UserState {
	if x != nil {
		return x.UserStates
	}
	return nil
}

type Operation_GameState_UserState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position *Position `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
	Score    int32     `protobuf:"varint,2,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *Operation_GameState_UserState) Reset() {
	*x = Operation_GameState_UserState{}
	mi := &file_message_proto_input_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Operation_GameState_UserState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operation_GameState_UserState) ProtoMessage() {}

func (x *Operation_GameState_UserState) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_input_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operation_GameState_UserState.ProtoReflect.Descriptor instead.
func (*Operation_GameState_UserState) Descriptor() ([]byte, []int) {
	return file_message_proto_input_proto_rawDescGZIP(), []int{2, 1, 0}
}

func (x *Operation_GameState_UserState) GetPosition() *Position {
	if x != nil {
		return x.Position
	}
	return nil
}

func (x *Operation_GameState_UserState) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

var File_message_proto_input_proto protoreflect.FileDescriptor

var file_message_proto_input_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x22, 0xb8, 0x01, 0x0a, 0x05, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x04, 0x69, 0x6e, 0x69, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x49, 0x6e, 0x69, 0x74,
	0x48, 0x00, 0x52, 0x04, 0x69, 0x6e, 0x69, 0x74, 0x12, 0x22, 0x0a, 0x02, 0x6f, 0x70, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x02, 0x6f, 0x70, 0x12, 0x24, 0x0a, 0x05,
	0x6c, 0x6f, 0x62, 0x62, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x2e, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x48, 0x00, 0x52, 0x05, 0x6c, 0x6f, 0x62,
	0x62, 0x79, 0x12, 0x21, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x48, 0x00, 0x52,
	0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x42, 0x06, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x22, 0x26, 0x0a,
	0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x01, 0x79, 0x22, 0xaf, 0x06, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1e, 0x2e, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x36, 0x0a, 0x09, 0x68, 0x69, 0x74, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x48, 0x69, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52,
	0x08, 0x68, 0x69, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x67, 0x61, 0x6d,
	0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x09, 0x67, 0x61, 0x6d, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x65,
	0x76, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x70, 0x72, 0x65, 0x76, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x1a, 0x6a, 0x0a, 0x08,
	0x48, 0x69, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x30, 0x0a, 0x0b, 0x6c, 0x65, 0x66, 0x74,
	0x5f, 0x62, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a,
	0x6c, 0x65, 0x66, 0x74, 0x42, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x12, 0x2c, 0x0a, 0x09, 0x72, 0x69,
	0x67, 0x68, 0x74, 0x5f, 0x74, 0x6f, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08,
	0x72, 0x69, 0x67, 0x68, 0x74, 0x54, 0x6f, 0x70, 0x1a, 0xa2, 0x01, 0x0a, 0x09, 0x47, 0x61, 0x6d,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x45, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x61,
	0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x1a, 0x4e, 0x0a,
	0x09, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x4d, 0x0a,
	0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x15, 0x44, 0x49,
	0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x55, 0x50, 0x10, 0x01, 0x12, 0x08, 0x0a,
	0x04, 0x44, 0x4f, 0x57, 0x4e, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x45, 0x46, 0x54, 0x10,
	0x03, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x49, 0x47, 0x48, 0x54, 0x10, 0x04, 0x22, 0x59, 0x0a, 0x0d,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a,
	0x15, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x4f, 0x56, 0x45,
	0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x54, 0x54, 0x41, 0x43, 0x4b, 0x10, 0x02, 0x12, 0x07,
	0x0a, 0x03, 0x48, 0x49, 0x54, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x47, 0x41, 0x4d, 0x45, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x10, 0x04, 0x22, 0x06, 0x0a, 0x04, 0x49, 0x6e, 0x69, 0x74, 0x22,
	0xe2, 0x01, 0x0a, 0x09, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x37, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e,
	0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x12, 0x19, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x22, 0x47, 0x0a, 0x0e, 0x52,
	0x6f, 0x6f, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a,
	0x17, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x41,
	0x49, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x4c, 0x41, 0x59, 0x49,
	0x4e, 0x47, 0x10, 0x02, 0x22, 0xfe, 0x01, 0x0a, 0x05, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x12, 0x2a,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x2e, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x2e, 0x4c, 0x6f, 0x62, 0x62, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f,
	0x6f, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72,
	0x6f, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x12, 0x31, 0x0a, 0x0b, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0a, 0x72, 0x6f, 0x6f,
	0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x22, 0x5c, 0x0a, 0x09, 0x4c, 0x6f, 0x62, 0x62, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x4c, 0x4f, 0x42, 0x42, 0x59, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x43,
	0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09,
	0x4a, 0x4f, 0x49, 0x4e, 0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x51,
	0x55, 0x49, 0x54, 0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x54,
	0x41, 0x54, 0x45, 0x10, 0x04, 0x22, 0xc0, 0x01, 0x0a, 0x04, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x28,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x6f, 0x6d,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6f,
	0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x0a, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x09, 0x72, 0x6f, 0x6f,
	0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x40, 0x0a, 0x08, 0x52, 0x6f, 0x6f, 0x6d, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x51, 0x55, 0x49, 0x54,
	0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x54, 0x41, 0x52, 0x54, 0x10, 0x02, 0x12, 0x09, 0x0a,
	0x05, 0x53, 0x54, 0x41, 0x54, 0x45, 0x10, 0x03, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_proto_input_proto_rawDescOnce sync.Once
	file_message_proto_input_proto_rawDescData = file_message_proto_input_proto_rawDesc
)

func file_message_proto_input_proto_rawDescGZIP() []byte {
	file_message_proto_input_proto_rawDescOnce.Do(func() {
		file_message_proto_input_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_input_proto_rawDescData)
	})
	return file_message_proto_input_proto_rawDescData
}

var file_message_proto_input_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_message_proto_input_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_message_proto_input_proto_goTypes = []any{
	(Operation_Direction)(0),              // 0: input.Operation.Direction
	(Operation_OperationType)(0),          // 1: input.Operation.OperationType
	(RoomState_RoomStatusType)(0),         // 2: input.RoomState.RoomStatusType
	(Lobby_LobbyType)(0),                  // 3: input.Lobby.LobbyType
	(Room_RoomType)(0),                    // 4: input.Room.RoomType
	(*Input)(nil),                         // 5: input.Input
	(*Position)(nil),                      // 6: input.Position
	(*Operation)(nil),                     // 7: input.Operation
	(*Init)(nil),                          // 8: input.Init
	(*RoomState)(nil),                     // 9: input.RoomState
	(*Lobby)(nil),                         // 10: input.Lobby
	(*Room)(nil),                          // 11: input.Room
	(*Operation_HitRange)(nil),            // 12: input.Operation.HitRange
	(*Operation_GameState)(nil),           // 13: input.Operation.GameState
	(*Operation_GameState_UserState)(nil), // 14: input.Operation.GameState.UserState
}
var file_message_proto_input_proto_depIdxs = []int32{
	8,  // 0: input.Input.init:type_name -> input.Init
	7,  // 1: input.Input.op:type_name -> input.Operation
	10, // 2: input.Input.lobby:type_name -> input.Lobby
	11, // 3: input.Input.room:type_name -> input.Room
	1,  // 4: input.Operation.type:type_name -> input.Operation.OperationType
	0,  // 5: input.Operation.direction:type_name -> input.Operation.Direction
	12, // 6: input.Operation.hit_range:type_name -> input.Operation.HitRange
	13, // 7: input.Operation.game_state:type_name -> input.Operation.GameState
	2,  // 8: input.RoomState.status:type_name -> input.RoomState.RoomStatusType
	3,  // 9: input.Lobby.type:type_name -> input.Lobby.LobbyType
	9,  // 10: input.Lobby.room_states:type_name -> input.RoomState
	4,  // 11: input.Room.type:type_name -> input.Room.RoomType
	9,  // 12: input.Room.room_state:type_name -> input.RoomState
	6,  // 13: input.Operation.HitRange.left_bottom:type_name -> input.Position
	6,  // 14: input.Operation.HitRange.right_top:type_name -> input.Position
	14, // 15: input.Operation.GameState.user_states:type_name -> input.Operation.GameState.UserState
	6,  // 16: input.Operation.GameState.UserState.position:type_name -> input.Position
	17, // [17:17] is the sub-list for method output_type
	17, // [17:17] is the sub-list for method input_type
	17, // [17:17] is the sub-list for extension type_name
	17, // [17:17] is the sub-list for extension extendee
	0,  // [0:17] is the sub-list for field type_name
}

func init() { file_message_proto_input_proto_init() }
func file_message_proto_input_proto_init() {
	if File_message_proto_input_proto != nil {
		return
	}
	file_message_proto_input_proto_msgTypes[0].OneofWrappers = []any{
		(*Input_Init)(nil),
		(*Input_Op)(nil),
		(*Input_Lobby)(nil),
		(*Input_Room)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_message_proto_input_proto_rawDesc,
			NumEnums:      5,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_proto_input_proto_goTypes,
		DependencyIndexes: file_message_proto_input_proto_depIdxs,
		EnumInfos:         file_message_proto_input_proto_enumTypes,
		MessageInfos:      file_message_proto_input_proto_msgTypes,
	}.Build()
	File_message_proto_input_proto = out.File
	file_message_proto_input_proto_rawDesc = nil
	file_message_proto_input_proto_goTypes = nil
	file_message_proto_input_proto_depIdxs = nil
}
