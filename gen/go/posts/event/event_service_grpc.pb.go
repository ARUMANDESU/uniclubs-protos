// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: posts/event/event_service.proto

package uniclubs_posts_service_v1_eventv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Event_GetEvent_FullMethodName           = "/posts.Event/GetEvent"
	Event_ListEvents_FullMethodName         = "/posts.Event/ListEvents"
	Event_CreateEvent_FullMethodName        = "/posts.Event/CreateEvent"
	Event_UpdateEvent_FullMethodName        = "/posts.Event/UpdateEvent"
	Event_DeleteEvent_FullMethodName        = "/posts.Event/DeleteEvent"
	Event_SendToReview_FullMethodName       = "/posts.Event/SendToReview"
	Event_RevokeReview_FullMethodName       = "/posts.Event/RevokeReview"
	Event_PublishEvent_FullMethodName       = "/posts.Event/PublishEvent"
	Event_UnpublishEvent_FullMethodName     = "/posts.Event/UnpublishEvent"
	Event_ApproveEvent_FullMethodName       = "/posts.Event/ApproveEvent"
	Event_RejectEvent_FullMethodName        = "/posts.Event/RejectEvent"
	Event_AddCollaborator_FullMethodName    = "/posts.Event/AddCollaborator"
	Event_RemoveCollaborator_FullMethodName = "/posts.Event/RemoveCollaborator"
	Event_HandleInviteClub_FullMethodName   = "/posts.Event/HandleInviteClub"
	Event_RevokeInviteClub_FullMethodName   = "/posts.Event/RevokeInviteClub"
	Event_AddOrganizer_FullMethodName       = "/posts.Event/AddOrganizer"
	Event_RemoveOrganizer_FullMethodName    = "/posts.Event/RemoveOrganizer"
	Event_HandleInviteUser_FullMethodName   = "/posts.Event/HandleInviteUser"
	Event_RevokeInviteUser_FullMethodName   = "/posts.Event/RevokeInviteUser"
)

// EventClient is the client API for Event service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventClient interface {
	GetEvent(ctx context.Context, in *GetEventRequest, opts ...grpc.CallOption) (*EventObject, error)
	ListEvents(ctx context.Context, in *ListEventsRequest, opts ...grpc.CallOption) (*ListEventsResponse, error)
	CreateEvent(ctx context.Context, in *CreateEventRequest, opts ...grpc.CallOption) (*EventObject, error)
	UpdateEvent(ctx context.Context, in *UpdateEventRequest, opts ...grpc.CallOption) (*EventObject, error)
	DeleteEvent(ctx context.Context, in *DeleteEventRequest, opts ...grpc.CallOption) (*EventObject, error)
	SendToReview(ctx context.Context, in *EventActionRequest, opts ...grpc.CallOption) (*EventObject, error)
	RevokeReview(ctx context.Context, in *EventActionRequest, opts ...grpc.CallOption) (*EventObject, error)
	PublishEvent(ctx context.Context, in *EventActionRequest, opts ...grpc.CallOption) (*EventObject, error)
	UnpublishEvent(ctx context.Context, in *EventActionRequest, opts ...grpc.CallOption) (*EventObject, error)
	ApproveEvent(ctx context.Context, in *EventActionRequest, opts ...grpc.CallOption) (*EventObject, error)
	RejectEvent(ctx context.Context, in *EventActionRequest, opts ...grpc.CallOption) (*EventObject, error)
	AddCollaborator(ctx context.Context, in *AddCollaboratorRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RemoveCollaborator(ctx context.Context, in *RemoveCollaboratorRequest, opts ...grpc.CallOption) (*EventObject, error)
	HandleInviteClub(ctx context.Context, in *HandleInviteClubRequest, opts ...grpc.CallOption) (*EventObject, error)
	RevokeInviteClub(ctx context.Context, in *RevokeInviteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AddOrganizer(ctx context.Context, in *AddOrganizerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RemoveOrganizer(ctx context.Context, in *RemoveOrganizerRequest, opts ...grpc.CallOption) (*EventObject, error)
	HandleInviteUser(ctx context.Context, in *HandleInviteUserRequest, opts ...grpc.CallOption) (*EventObject, error)
	RevokeInviteUser(ctx context.Context, in *RevokeInviteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type eventClient struct {
	cc grpc.ClientConnInterface
}

func NewEventClient(cc grpc.ClientConnInterface) EventClient {
	return &eventClient{cc}
}

func (c *eventClient) GetEvent(ctx context.Context, in *GetEventRequest, opts ...grpc.CallOption) (*EventObject, error) {
	out := new(EventObject)
	err := c.cc.Invoke(ctx, Event_GetEvent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventClient) ListEvents(ctx context.Context, in *ListEventsRequest, opts ...grpc.CallOption) (*ListEventsResponse, error) {
	out := new(ListEventsResponse)
	err := c.cc.Invoke(ctx, Event_ListEvents_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventClient) CreateEvent(ctx context.Context, in *CreateEventRequest, opts ...grpc.CallOption) (*EventObject, error) {
	out := new(EventObject)
	err := c.cc.Invoke(ctx, Event_CreateEvent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventClient) UpdateEvent(ctx context.Context, in *UpdateEventRequest, opts ...grpc.CallOption) (*EventObject, error) {
	out := new(EventObject)
	err := c.cc.Invoke(ctx, Event_UpdateEvent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventClient) DeleteEvent(ctx context.Context, in *DeleteEventRequest, opts ...grpc.CallOption) (*EventObject, error) {
	out := new(EventObject)
	err := c.cc.Invoke(ctx, Event_DeleteEvent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventClient) SendToReview(ctx context.Context, in *EventActionRequest, opts ...grpc.CallOption) (*EventObject, error) {
	out := new(EventObject)
	err := c.cc.Invoke(ctx, Event_SendToReview_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventClient) RevokeReview(ctx context.Context, in *EventActionRequest, opts ...grpc.CallOption) (*EventObject, error) {
	out := new(EventObject)
	err := c.cc.Invoke(ctx, Event_RevokeReview_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventClient) PublishEvent(ctx context.Context, in *EventActionRequest, opts ...grpc.CallOption) (*EventObject, error) {
	out := new(EventObject)
	err := c.cc.Invoke(ctx, Event_PublishEvent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventClient) UnpublishEvent(ctx context.Context, in *EventActionRequest, opts ...grpc.CallOption) (*EventObject, error) {
	out := new(EventObject)
	err := c.cc.Invoke(ctx, Event_UnpublishEvent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventClient) ApproveEvent(ctx context.Context, in *EventActionRequest, opts ...grpc.CallOption) (*EventObject, error) {
	out := new(EventObject)
	err := c.cc.Invoke(ctx, Event_ApproveEvent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventClient) RejectEvent(ctx context.Context, in *EventActionRequest, opts ...grpc.CallOption) (*EventObject, error) {
	out := new(EventObject)
	err := c.cc.Invoke(ctx, Event_RejectEvent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventClient) AddCollaborator(ctx context.Context, in *AddCollaboratorRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Event_AddCollaborator_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventClient) RemoveCollaborator(ctx context.Context, in *RemoveCollaboratorRequest, opts ...grpc.CallOption) (*EventObject, error) {
	out := new(EventObject)
	err := c.cc.Invoke(ctx, Event_RemoveCollaborator_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventClient) HandleInviteClub(ctx context.Context, in *HandleInviteClubRequest, opts ...grpc.CallOption) (*EventObject, error) {
	out := new(EventObject)
	err := c.cc.Invoke(ctx, Event_HandleInviteClub_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventClient) RevokeInviteClub(ctx context.Context, in *RevokeInviteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Event_RevokeInviteClub_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventClient) AddOrganizer(ctx context.Context, in *AddOrganizerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Event_AddOrganizer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventClient) RemoveOrganizer(ctx context.Context, in *RemoveOrganizerRequest, opts ...grpc.CallOption) (*EventObject, error) {
	out := new(EventObject)
	err := c.cc.Invoke(ctx, Event_RemoveOrganizer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventClient) HandleInviteUser(ctx context.Context, in *HandleInviteUserRequest, opts ...grpc.CallOption) (*EventObject, error) {
	out := new(EventObject)
	err := c.cc.Invoke(ctx, Event_HandleInviteUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventClient) RevokeInviteUser(ctx context.Context, in *RevokeInviteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Event_RevokeInviteUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventServer is the server API for Event service.
// All implementations must embed UnimplementedEventServer
// for forward compatibility
type EventServer interface {
	GetEvent(context.Context, *GetEventRequest) (*EventObject, error)
	ListEvents(context.Context, *ListEventsRequest) (*ListEventsResponse, error)
	CreateEvent(context.Context, *CreateEventRequest) (*EventObject, error)
	UpdateEvent(context.Context, *UpdateEventRequest) (*EventObject, error)
	DeleteEvent(context.Context, *DeleteEventRequest) (*EventObject, error)
	SendToReview(context.Context, *EventActionRequest) (*EventObject, error)
	RevokeReview(context.Context, *EventActionRequest) (*EventObject, error)
	PublishEvent(context.Context, *EventActionRequest) (*EventObject, error)
	UnpublishEvent(context.Context, *EventActionRequest) (*EventObject, error)
	ApproveEvent(context.Context, *EventActionRequest) (*EventObject, error)
	RejectEvent(context.Context, *EventActionRequest) (*EventObject, error)
	AddCollaborator(context.Context, *AddCollaboratorRequest) (*emptypb.Empty, error)
	RemoveCollaborator(context.Context, *RemoveCollaboratorRequest) (*EventObject, error)
	HandleInviteClub(context.Context, *HandleInviteClubRequest) (*EventObject, error)
	RevokeInviteClub(context.Context, *RevokeInviteRequest) (*emptypb.Empty, error)
	AddOrganizer(context.Context, *AddOrganizerRequest) (*emptypb.Empty, error)
	RemoveOrganizer(context.Context, *RemoveOrganizerRequest) (*EventObject, error)
	HandleInviteUser(context.Context, *HandleInviteUserRequest) (*EventObject, error)
	RevokeInviteUser(context.Context, *RevokeInviteRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedEventServer()
}

// UnimplementedEventServer must be embedded to have forward compatible implementations.
type UnimplementedEventServer struct {
}

func (UnimplementedEventServer) GetEvent(context.Context, *GetEventRequest) (*EventObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvent not implemented")
}
func (UnimplementedEventServer) ListEvents(context.Context, *ListEventsRequest) (*ListEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEvents not implemented")
}
func (UnimplementedEventServer) CreateEvent(context.Context, *CreateEventRequest) (*EventObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEvent not implemented")
}
func (UnimplementedEventServer) UpdateEvent(context.Context, *UpdateEventRequest) (*EventObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEvent not implemented")
}
func (UnimplementedEventServer) DeleteEvent(context.Context, *DeleteEventRequest) (*EventObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEvent not implemented")
}
func (UnimplementedEventServer) SendToReview(context.Context, *EventActionRequest) (*EventObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendToReview not implemented")
}
func (UnimplementedEventServer) RevokeReview(context.Context, *EventActionRequest) (*EventObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeReview not implemented")
}
func (UnimplementedEventServer) PublishEvent(context.Context, *EventActionRequest) (*EventObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishEvent not implemented")
}
func (UnimplementedEventServer) UnpublishEvent(context.Context, *EventActionRequest) (*EventObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnpublishEvent not implemented")
}
func (UnimplementedEventServer) ApproveEvent(context.Context, *EventActionRequest) (*EventObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApproveEvent not implemented")
}
func (UnimplementedEventServer) RejectEvent(context.Context, *EventActionRequest) (*EventObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RejectEvent not implemented")
}
func (UnimplementedEventServer) AddCollaborator(context.Context, *AddCollaboratorRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCollaborator not implemented")
}
func (UnimplementedEventServer) RemoveCollaborator(context.Context, *RemoveCollaboratorRequest) (*EventObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveCollaborator not implemented")
}
func (UnimplementedEventServer) HandleInviteClub(context.Context, *HandleInviteClubRequest) (*EventObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleInviteClub not implemented")
}
func (UnimplementedEventServer) RevokeInviteClub(context.Context, *RevokeInviteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeInviteClub not implemented")
}
func (UnimplementedEventServer) AddOrganizer(context.Context, *AddOrganizerRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddOrganizer not implemented")
}
func (UnimplementedEventServer) RemoveOrganizer(context.Context, *RemoveOrganizerRequest) (*EventObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveOrganizer not implemented")
}
func (UnimplementedEventServer) HandleInviteUser(context.Context, *HandleInviteUserRequest) (*EventObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleInviteUser not implemented")
}
func (UnimplementedEventServer) RevokeInviteUser(context.Context, *RevokeInviteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeInviteUser not implemented")
}
func (UnimplementedEventServer) mustEmbedUnimplementedEventServer() {}

// UnsafeEventServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventServer will
// result in compilation errors.
type UnsafeEventServer interface {
	mustEmbedUnimplementedEventServer()
}

func RegisterEventServer(s grpc.ServiceRegistrar, srv EventServer) {
	s.RegisterService(&Event_ServiceDesc, srv)
}

func _Event_GetEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServer).GetEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Event_GetEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServer).GetEvent(ctx, req.(*GetEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Event_ListEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServer).ListEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Event_ListEvents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServer).ListEvents(ctx, req.(*ListEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Event_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServer).CreateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Event_CreateEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServer).CreateEvent(ctx, req.(*CreateEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Event_UpdateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServer).UpdateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Event_UpdateEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServer).UpdateEvent(ctx, req.(*UpdateEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Event_DeleteEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServer).DeleteEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Event_DeleteEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServer).DeleteEvent(ctx, req.(*DeleteEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Event_SendToReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServer).SendToReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Event_SendToReview_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServer).SendToReview(ctx, req.(*EventActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Event_RevokeReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServer).RevokeReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Event_RevokeReview_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServer).RevokeReview(ctx, req.(*EventActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Event_PublishEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServer).PublishEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Event_PublishEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServer).PublishEvent(ctx, req.(*EventActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Event_UnpublishEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServer).UnpublishEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Event_UnpublishEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServer).UnpublishEvent(ctx, req.(*EventActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Event_ApproveEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServer).ApproveEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Event_ApproveEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServer).ApproveEvent(ctx, req.(*EventActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Event_RejectEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServer).RejectEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Event_RejectEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServer).RejectEvent(ctx, req.(*EventActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Event_AddCollaborator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCollaboratorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServer).AddCollaborator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Event_AddCollaborator_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServer).AddCollaborator(ctx, req.(*AddCollaboratorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Event_RemoveCollaborator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveCollaboratorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServer).RemoveCollaborator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Event_RemoveCollaborator_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServer).RemoveCollaborator(ctx, req.(*RemoveCollaboratorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Event_HandleInviteClub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandleInviteClubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServer).HandleInviteClub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Event_HandleInviteClub_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServer).HandleInviteClub(ctx, req.(*HandleInviteClubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Event_RevokeInviteClub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeInviteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServer).RevokeInviteClub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Event_RevokeInviteClub_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServer).RevokeInviteClub(ctx, req.(*RevokeInviteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Event_AddOrganizer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddOrganizerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServer).AddOrganizer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Event_AddOrganizer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServer).AddOrganizer(ctx, req.(*AddOrganizerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Event_RemoveOrganizer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveOrganizerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServer).RemoveOrganizer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Event_RemoveOrganizer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServer).RemoveOrganizer(ctx, req.(*RemoveOrganizerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Event_HandleInviteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandleInviteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServer).HandleInviteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Event_HandleInviteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServer).HandleInviteUser(ctx, req.(*HandleInviteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Event_RevokeInviteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeInviteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServer).RevokeInviteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Event_RevokeInviteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServer).RevokeInviteUser(ctx, req.(*RevokeInviteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Event_ServiceDesc is the grpc.ServiceDesc for Event service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Event_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "posts.Event",
	HandlerType: (*EventServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEvent",
			Handler:    _Event_GetEvent_Handler,
		},
		{
			MethodName: "ListEvents",
			Handler:    _Event_ListEvents_Handler,
		},
		{
			MethodName: "CreateEvent",
			Handler:    _Event_CreateEvent_Handler,
		},
		{
			MethodName: "UpdateEvent",
			Handler:    _Event_UpdateEvent_Handler,
		},
		{
			MethodName: "DeleteEvent",
			Handler:    _Event_DeleteEvent_Handler,
		},
		{
			MethodName: "SendToReview",
			Handler:    _Event_SendToReview_Handler,
		},
		{
			MethodName: "RevokeReview",
			Handler:    _Event_RevokeReview_Handler,
		},
		{
			MethodName: "PublishEvent",
			Handler:    _Event_PublishEvent_Handler,
		},
		{
			MethodName: "UnpublishEvent",
			Handler:    _Event_UnpublishEvent_Handler,
		},
		{
			MethodName: "ApproveEvent",
			Handler:    _Event_ApproveEvent_Handler,
		},
		{
			MethodName: "RejectEvent",
			Handler:    _Event_RejectEvent_Handler,
		},
		{
			MethodName: "AddCollaborator",
			Handler:    _Event_AddCollaborator_Handler,
		},
		{
			MethodName: "RemoveCollaborator",
			Handler:    _Event_RemoveCollaborator_Handler,
		},
		{
			MethodName: "HandleInviteClub",
			Handler:    _Event_HandleInviteClub_Handler,
		},
		{
			MethodName: "RevokeInviteClub",
			Handler:    _Event_RevokeInviteClub_Handler,
		},
		{
			MethodName: "AddOrganizer",
			Handler:    _Event_AddOrganizer_Handler,
		},
		{
			MethodName: "RemoveOrganizer",
			Handler:    _Event_RemoveOrganizer_Handler,
		},
		{
			MethodName: "HandleInviteUser",
			Handler:    _Event_HandleInviteUser_Handler,
		},
		{
			MethodName: "RevokeInviteUser",
			Handler:    _Event_RevokeInviteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "posts/event/event_service.proto",
}
