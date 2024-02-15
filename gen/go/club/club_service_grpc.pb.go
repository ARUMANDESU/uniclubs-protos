// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: club/club_service.proto

package uniclubs_club_service_v1_clubv1

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ClubClient is the client API for Club service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClubClient interface {
	CreateClub(ctx context.Context, in *CreateClubRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	HandleNewClub(ctx context.Context, in *HandleNewClubRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetClub(ctx context.Context, in *GetClubRequest, opts ...grpc.CallOption) (*ClubObject, error)
	ListClubs(ctx context.Context, in *ListClubRequest, opts ...grpc.CallOption) (*ListClubResponse, error)
	ListNotApprovedClubs(ctx context.Context, in *ListNotApprovedClubsRequest, opts ...grpc.CallOption) (*ListNotApprovedClubsResponse, error)
	RequestToJoinClub(ctx context.Context, in *RequestToJoinClubRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	HandleJoinClub(ctx context.Context, in *HandleJoinClubRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	DeactivateClub(ctx context.Context, in *DeactivateClubRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	UpdateClub(ctx context.Context, in *UpdateClubRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetUserClubs(ctx context.Context, in *GetUserClubsRequest, opts ...grpc.CallOption) (*GetUserClubsResponse, error)
	ListClubMembers(ctx context.Context, in *ListClubMembersRequest, opts ...grpc.CallOption) (*ListClubMembersResponse, error)
	LeaveClub(ctx context.Context, in *LeaveClubRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type clubClient struct {
	cc grpc.ClientConnInterface
}

func NewClubClient(cc grpc.ClientConnInterface) ClubClient {
	return &clubClient{cc}
}

func (c *clubClient) CreateClub(ctx context.Context, in *CreateClubRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/club.Club/CreateClub", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clubClient) HandleNewClub(ctx context.Context, in *HandleNewClubRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/club.Club/HandleNewClub", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clubClient) GetClub(ctx context.Context, in *GetClubRequest, opts ...grpc.CallOption) (*ClubObject, error) {
	out := new(ClubObject)
	err := c.cc.Invoke(ctx, "/club.Club/GetClub", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clubClient) ListClubs(ctx context.Context, in *ListClubRequest, opts ...grpc.CallOption) (*ListClubResponse, error) {
	out := new(ListClubResponse)
	err := c.cc.Invoke(ctx, "/club.Club/ListClubs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clubClient) ListNotApprovedClubs(ctx context.Context, in *ListNotApprovedClubsRequest, opts ...grpc.CallOption) (*ListNotApprovedClubsResponse, error) {
	out := new(ListNotApprovedClubsResponse)
	err := c.cc.Invoke(ctx, "/club.Club/ListNotApprovedClubs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clubClient) RequestToJoinClub(ctx context.Context, in *RequestToJoinClubRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/club.Club/RequestToJoinClub", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clubClient) HandleJoinClub(ctx context.Context, in *HandleJoinClubRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/club.Club/HandleJoinClub", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clubClient) DeactivateClub(ctx context.Context, in *DeactivateClubRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/club.Club/DeactivateClub", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clubClient) UpdateClub(ctx context.Context, in *UpdateClubRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/club.Club/UpdateClub", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clubClient) GetUserClubs(ctx context.Context, in *GetUserClubsRequest, opts ...grpc.CallOption) (*GetUserClubsResponse, error) {
	out := new(GetUserClubsResponse)
	err := c.cc.Invoke(ctx, "/club.Club/GetUserClubs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clubClient) ListClubMembers(ctx context.Context, in *ListClubMembersRequest, opts ...grpc.CallOption) (*ListClubMembersResponse, error) {
	out := new(ListClubMembersResponse)
	err := c.cc.Invoke(ctx, "/club.Club/ListClubMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clubClient) LeaveClub(ctx context.Context, in *LeaveClubRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/club.Club/LeaveClub", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClubServer is the server API for Club service.
// All implementations must embed UnimplementedClubServer
// for forward compatibility
type ClubServer interface {
	CreateClub(context.Context, *CreateClubRequest) (*empty.Empty, error)
	HandleNewClub(context.Context, *HandleNewClubRequest) (*empty.Empty, error)
	GetClub(context.Context, *GetClubRequest) (*ClubObject, error)
	ListClubs(context.Context, *ListClubRequest) (*ListClubResponse, error)
	ListNotApprovedClubs(context.Context, *ListNotApprovedClubsRequest) (*ListNotApprovedClubsResponse, error)
	RequestToJoinClub(context.Context, *RequestToJoinClubRequest) (*empty.Empty, error)
	HandleJoinClub(context.Context, *HandleJoinClubRequest) (*empty.Empty, error)
	DeactivateClub(context.Context, *DeactivateClubRequest) (*empty.Empty, error)
	UpdateClub(context.Context, *UpdateClubRequest) (*empty.Empty, error)
	GetUserClubs(context.Context, *GetUserClubsRequest) (*GetUserClubsResponse, error)
	ListClubMembers(context.Context, *ListClubMembersRequest) (*ListClubMembersResponse, error)
	LeaveClub(context.Context, *LeaveClubRequest) (*empty.Empty, error)
	mustEmbedUnimplementedClubServer()
}

// UnimplementedClubServer must be embedded to have forward compatible implementations.
type UnimplementedClubServer struct {
}

func (UnimplementedClubServer) CreateClub(context.Context, *CreateClubRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateClub not implemented")
}
func (UnimplementedClubServer) HandleNewClub(context.Context, *HandleNewClubRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleNewClub not implemented")
}
func (UnimplementedClubServer) GetClub(context.Context, *GetClubRequest) (*ClubObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClub not implemented")
}
func (UnimplementedClubServer) ListClubs(context.Context, *ListClubRequest) (*ListClubResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListClubs not implemented")
}
func (UnimplementedClubServer) ListNotApprovedClubs(context.Context, *ListNotApprovedClubsRequest) (*ListNotApprovedClubsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNotApprovedClubs not implemented")
}
func (UnimplementedClubServer) RequestToJoinClub(context.Context, *RequestToJoinClubRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestToJoinClub not implemented")
}
func (UnimplementedClubServer) HandleJoinClub(context.Context, *HandleJoinClubRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleJoinClub not implemented")
}
func (UnimplementedClubServer) DeactivateClub(context.Context, *DeactivateClubRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeactivateClub not implemented")
}
func (UnimplementedClubServer) UpdateClub(context.Context, *UpdateClubRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateClub not implemented")
}
func (UnimplementedClubServer) GetUserClubs(context.Context, *GetUserClubsRequest) (*GetUserClubsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserClubs not implemented")
}
func (UnimplementedClubServer) ListClubMembers(context.Context, *ListClubMembersRequest) (*ListClubMembersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListClubMembers not implemented")
}
func (UnimplementedClubServer) LeaveClub(context.Context, *LeaveClubRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveClub not implemented")
}
func (UnimplementedClubServer) mustEmbedUnimplementedClubServer() {}

// UnsafeClubServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClubServer will
// result in compilation errors.
type UnsafeClubServer interface {
	mustEmbedUnimplementedClubServer()
}

func RegisterClubServer(s grpc.ServiceRegistrar, srv ClubServer) {
	s.RegisterService(&Club_ServiceDesc, srv)
}

func _Club_CreateClub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClubServer).CreateClub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/club.Club/CreateClub",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClubServer).CreateClub(ctx, req.(*CreateClubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Club_HandleNewClub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandleNewClubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClubServer).HandleNewClub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/club.Club/HandleNewClub",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClubServer).HandleNewClub(ctx, req.(*HandleNewClubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Club_GetClub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClubServer).GetClub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/club.Club/GetClub",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClubServer).GetClub(ctx, req.(*GetClubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Club_ListClubs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClubServer).ListClubs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/club.Club/ListClubs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClubServer).ListClubs(ctx, req.(*ListClubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Club_ListNotApprovedClubs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNotApprovedClubsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClubServer).ListNotApprovedClubs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/club.Club/ListNotApprovedClubs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClubServer).ListNotApprovedClubs(ctx, req.(*ListNotApprovedClubsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Club_RequestToJoinClub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestToJoinClubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClubServer).RequestToJoinClub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/club.Club/RequestToJoinClub",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClubServer).RequestToJoinClub(ctx, req.(*RequestToJoinClubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Club_HandleJoinClub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandleJoinClubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClubServer).HandleJoinClub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/club.Club/HandleJoinClub",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClubServer).HandleJoinClub(ctx, req.(*HandleJoinClubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Club_DeactivateClub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeactivateClubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClubServer).DeactivateClub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/club.Club/DeactivateClub",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClubServer).DeactivateClub(ctx, req.(*DeactivateClubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Club_UpdateClub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateClubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClubServer).UpdateClub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/club.Club/UpdateClub",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClubServer).UpdateClub(ctx, req.(*UpdateClubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Club_GetUserClubs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserClubsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClubServer).GetUserClubs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/club.Club/GetUserClubs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClubServer).GetUserClubs(ctx, req.(*GetUserClubsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Club_ListClubMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClubMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClubServer).ListClubMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/club.Club/ListClubMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClubServer).ListClubMembers(ctx, req.(*ListClubMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Club_LeaveClub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveClubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClubServer).LeaveClub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/club.Club/LeaveClub",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClubServer).LeaveClub(ctx, req.(*LeaveClubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Club_ServiceDesc is the grpc.ServiceDesc for Club service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Club_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "club.Club",
	HandlerType: (*ClubServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateClub",
			Handler:    _Club_CreateClub_Handler,
		},
		{
			MethodName: "HandleNewClub",
			Handler:    _Club_HandleNewClub_Handler,
		},
		{
			MethodName: "GetClub",
			Handler:    _Club_GetClub_Handler,
		},
		{
			MethodName: "ListClubs",
			Handler:    _Club_ListClubs_Handler,
		},
		{
			MethodName: "ListNotApprovedClubs",
			Handler:    _Club_ListNotApprovedClubs_Handler,
		},
		{
			MethodName: "RequestToJoinClub",
			Handler:    _Club_RequestToJoinClub_Handler,
		},
		{
			MethodName: "HandleJoinClub",
			Handler:    _Club_HandleJoinClub_Handler,
		},
		{
			MethodName: "DeactivateClub",
			Handler:    _Club_DeactivateClub_Handler,
		},
		{
			MethodName: "UpdateClub",
			Handler:    _Club_UpdateClub_Handler,
		},
		{
			MethodName: "GetUserClubs",
			Handler:    _Club_GetUserClubs_Handler,
		},
		{
			MethodName: "ListClubMembers",
			Handler:    _Club_ListClubMembers_Handler,
		},
		{
			MethodName: "LeaveClub",
			Handler:    _Club_LeaveClub_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "club/club_service.proto",
}
