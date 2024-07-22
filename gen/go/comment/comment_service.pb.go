// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: comment/comment_service.proto

package comment

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SortBy int32

const (
	SortBy_SORT_BY_UNSPECIFIED SortBy = 0
	SortBy_SORT_BY_CREATED_AT  SortBy = 1
	SortBy_SORT_BY_UPDATED_AT  SortBy = 2
)

// Enum value maps for SortBy.
var (
	SortBy_name = map[int32]string{
		0: "SORT_BY_UNSPECIFIED",
		1: "SORT_BY_CREATED_AT",
		2: "SORT_BY_UPDATED_AT",
	}
	SortBy_value = map[string]int32{
		"SORT_BY_UNSPECIFIED": 0,
		"SORT_BY_CREATED_AT":  1,
		"SORT_BY_UPDATED_AT":  2,
	}
)

func (x SortBy) Enum() *SortBy {
	p := new(SortBy)
	*p = x
	return p
}

func (x SortBy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SortBy) Descriptor() protoreflect.EnumDescriptor {
	return file_comment_comment_service_proto_enumTypes[0].Descriptor()
}

func (SortBy) Type() protoreflect.EnumType {
	return &file_comment_comment_service_proto_enumTypes[0]
}

func (x SortBy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SortBy.Descriptor instead.
func (SortBy) EnumDescriptor() ([]byte, []int) {
	return file_comment_comment_service_proto_rawDescGZIP(), []int{0}
}

type SortOrder int32

const (
	SortOrder_SORT_ORDER_UNSPECIFIED SortOrder = 0
	SortOrder_SORT_ORDER_ASC         SortOrder = 1
	SortOrder_SORT_ORDER_DESC        SortOrder = 2
)

// Enum value maps for SortOrder.
var (
	SortOrder_name = map[int32]string{
		0: "SORT_ORDER_UNSPECIFIED",
		1: "SORT_ORDER_ASC",
		2: "SORT_ORDER_DESC",
	}
	SortOrder_value = map[string]int32{
		"SORT_ORDER_UNSPECIFIED": 0,
		"SORT_ORDER_ASC":         1,
		"SORT_ORDER_DESC":        2,
	}
)

func (x SortOrder) Enum() *SortOrder {
	p := new(SortOrder)
	*p = x
	return p
}

func (x SortOrder) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SortOrder) Descriptor() protoreflect.EnumDescriptor {
	return file_comment_comment_service_proto_enumTypes[1].Descriptor()
}

func (SortOrder) Type() protoreflect.EnumType {
	return &file_comment_comment_service_proto_enumTypes[1]
}

func (x SortOrder) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SortOrder.Descriptor instead.
func (SortOrder) EnumDescriptor() ([]byte, []int) {
	return file_comment_comment_service_proto_rawDescGZIP(), []int{1}
}

type CommentObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	User      *UserObject `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	PostId    string      `protobuf:"bytes,3,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	Body      string      `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	CreatedAt string      `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string      `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *CommentObject) Reset() {
	*x = CommentObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_comment_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentObject) ProtoMessage() {}

func (x *CommentObject) ProtoReflect() protoreflect.Message {
	mi := &file_comment_comment_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentObject.ProtoReflect.Descriptor instead.
func (*CommentObject) Descriptor() ([]byte, []int) {
	return file_comment_comment_service_proto_rawDescGZIP(), []int{0}
}

func (x *CommentObject) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CommentObject) GetUser() *UserObject {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *CommentObject) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

func (x *CommentObject) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *CommentObject) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *CommentObject) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type UserObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	AvatarUrl string `protobuf:"bytes,4,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
}

func (x *UserObject) Reset() {
	*x = UserObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_comment_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserObject) ProtoMessage() {}

func (x *UserObject) ProtoReflect() protoreflect.Message {
	mi := &file_comment_comment_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserObject.ProtoReflect.Descriptor instead.
func (*UserObject) Descriptor() ([]byte, []int) {
	return file_comment_comment_service_proto_rawDescGZIP(), []int{1}
}

func (x *UserObject) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserObject) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UserObject) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UserObject) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

type PaginationMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentPage  int32 `protobuf:"varint,1,opt,name=current_page,json=currentPage,proto3" json:"current_page,omitempty"`
	PageSize     int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	FirstPage    int32 `protobuf:"varint,3,opt,name=first_page,json=firstPage,proto3" json:"first_page,omitempty"`
	LastPage     int32 `protobuf:"varint,4,opt,name=last_page,json=lastPage,proto3" json:"last_page,omitempty"`
	TotalRecords int32 `protobuf:"varint,5,opt,name=total_records,json=totalRecords,proto3" json:"total_records,omitempty"`
}

func (x *PaginationMetadata) Reset() {
	*x = PaginationMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_comment_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaginationMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationMetadata) ProtoMessage() {}

func (x *PaginationMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_comment_comment_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationMetadata.ProtoReflect.Descriptor instead.
func (*PaginationMetadata) Descriptor() ([]byte, []int) {
	return file_comment_comment_service_proto_rawDescGZIP(), []int{2}
}

func (x *PaginationMetadata) GetCurrentPage() int32 {
	if x != nil {
		return x.CurrentPage
	}
	return 0
}

func (x *PaginationMetadata) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PaginationMetadata) GetFirstPage() int32 {
	if x != nil {
		return x.FirstPage
	}
	return 0
}

func (x *PaginationMetadata) GetLastPage() int32 {
	if x != nil {
		return x.LastPage
	}
	return 0
}

func (x *PaginationMetadata) GetTotalRecords() int32 {
	if x != nil {
		return x.TotalRecords
	}
	return 0
}

type GetCommentByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCommentByIDRequest) Reset() {
	*x = GetCommentByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_comment_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommentByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentByIDRequest) ProtoMessage() {}

func (x *GetCommentByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comment_comment_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentByIDRequest.ProtoReflect.Descriptor instead.
func (*GetCommentByIDRequest) Descriptor() ([]byte, []int) {
	return file_comment_comment_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetCommentByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetCommentByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comment *CommentObject `protobuf:"bytes,1,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *GetCommentByIDResponse) Reset() {
	*x = GetCommentByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_comment_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommentByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentByIDResponse) ProtoMessage() {}

func (x *GetCommentByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_comment_comment_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentByIDResponse.ProtoReflect.Descriptor instead.
func (*GetCommentByIDResponse) Descriptor() ([]byte, []int) {
	return file_comment_comment_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetCommentByIDResponse) GetComment() *CommentObject {
	if x != nil {
		return x.Comment
	}
	return nil
}

type ListPostCommentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId     string                 `protobuf:"bytes,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	Page       int32                  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	PageSize   int32                  `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	SortBy     SortBy                 `protobuf:"varint,4,opt,name=sort_by,json=sortBy,proto3,enum=comment.SortBy" json:"sort_by,omitempty"`
	SortOrder  SortOrder              `protobuf:"varint,5,opt,name=sort_order,json=sortOrder,proto3,enum=comment.SortOrder" json:"sort_order,omitempty"`
	FilterMask *fieldmaskpb.FieldMask `protobuf:"bytes,6,opt,name=filter_mask,json=filterMask,proto3" json:"filter_mask,omitempty"`
}

func (x *ListPostCommentsRequest) Reset() {
	*x = ListPostCommentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_comment_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPostCommentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPostCommentsRequest) ProtoMessage() {}

func (x *ListPostCommentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comment_comment_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPostCommentsRequest.ProtoReflect.Descriptor instead.
func (*ListPostCommentsRequest) Descriptor() ([]byte, []int) {
	return file_comment_comment_service_proto_rawDescGZIP(), []int{5}
}

func (x *ListPostCommentsRequest) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

func (x *ListPostCommentsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListPostCommentsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListPostCommentsRequest) GetSortBy() SortBy {
	if x != nil {
		return x.SortBy
	}
	return SortBy_SORT_BY_UNSPECIFIED
}

func (x *ListPostCommentsRequest) GetSortOrder() SortOrder {
	if x != nil {
		return x.SortOrder
	}
	return SortOrder_SORT_ORDER_UNSPECIFIED
}

func (x *ListPostCommentsRequest) GetFilterMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.FilterMask
	}
	return nil
}

type ListPostCommentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comments           []*CommentObject    `protobuf:"bytes,1,rep,name=comments,proto3" json:"comments,omitempty"`
	PaginationMetadata *PaginationMetadata `protobuf:"bytes,2,opt,name=pagination_metadata,json=paginationMetadata,proto3" json:"pagination_metadata,omitempty"`
}

func (x *ListPostCommentsResponse) Reset() {
	*x = ListPostCommentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_comment_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPostCommentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPostCommentsResponse) ProtoMessage() {}

func (x *ListPostCommentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_comment_comment_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPostCommentsResponse.ProtoReflect.Descriptor instead.
func (*ListPostCommentsResponse) Descriptor() ([]byte, []int) {
	return file_comment_comment_service_proto_rawDescGZIP(), []int{6}
}

func (x *ListPostCommentsResponse) GetComments() []*CommentObject {
	if x != nil {
		return x.Comments
	}
	return nil
}

func (x *ListPostCommentsResponse) GetPaginationMetadata() *PaginationMetadata {
	if x != nil {
		return x.PaginationMetadata
	}
	return nil
}

var File_comment_comment_service_proto protoreflect.FileDescriptor

var file_comment_comment_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x01, 0x0a, 0x0d, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x27, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x22, 0x77, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x22, 0xb5, 0x01, 0x0a, 0x12, 0x50, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x21, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50,
	0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x50, 0x61, 0x67, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x50, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x73, 0x22, 0x27, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x42,
	0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4a, 0x0a, 0x16, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0xfd, 0x01, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x28, 0x0a, 0x07,
	0x73, 0x6f, 0x72, 0x74, 0x5f, 0x62, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x52, 0x06,
	0x73, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x12, 0x31, 0x0a, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x09,
	0x73, 0x6f, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x0b, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x9c, 0x01, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x08, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x4c, 0x0a, 0x13, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x12, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2a, 0x51, 0x0a, 0x06, 0x53, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x12,
	0x17, 0x0a, 0x13, 0x53, 0x4f, 0x52, 0x54, 0x5f, 0x42, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x4f, 0x52, 0x54,
	0x5f, 0x42, 0x59, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x41, 0x54, 0x10, 0x01,
	0x12, 0x16, 0x0a, 0x12, 0x53, 0x4f, 0x52, 0x54, 0x5f, 0x42, 0x59, 0x5f, 0x55, 0x50, 0x44, 0x41,
	0x54, 0x45, 0x44, 0x5f, 0x41, 0x54, 0x10, 0x02, 0x2a, 0x50, 0x0a, 0x09, 0x53, 0x6f, 0x72, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x4f, 0x52, 0x54, 0x5f, 0x4f, 0x52,
	0x44, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x4f, 0x52, 0x54, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f,
	0x41, 0x53, 0x43, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x4f, 0x52, 0x54, 0x5f, 0x4f, 0x52,
	0x44, 0x45, 0x52, 0x5f, 0x44, 0x45, 0x53, 0x43, 0x10, 0x02, 0x32, 0xb5, 0x01, 0x0a, 0x07, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x51, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49,
	0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49,
	0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x10, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x20, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f,
	0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x41, 0x52, 0x55, 0x4d, 0x41, 0x4e, 0x44, 0x45, 0x53, 0x55, 0x2f, 0x75, 0x6e, 0x69, 0x63,
	0x6c, 0x75, 0x62, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_comment_comment_service_proto_rawDescOnce sync.Once
	file_comment_comment_service_proto_rawDescData = file_comment_comment_service_proto_rawDesc
)

func file_comment_comment_service_proto_rawDescGZIP() []byte {
	file_comment_comment_service_proto_rawDescOnce.Do(func() {
		file_comment_comment_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_comment_comment_service_proto_rawDescData)
	})
	return file_comment_comment_service_proto_rawDescData
}

var file_comment_comment_service_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_comment_comment_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_comment_comment_service_proto_goTypes = []any{
	(SortBy)(0),                      // 0: comment.SortBy
	(SortOrder)(0),                   // 1: comment.SortOrder
	(*CommentObject)(nil),            // 2: comment.CommentObject
	(*UserObject)(nil),               // 3: comment.UserObject
	(*PaginationMetadata)(nil),       // 4: comment.PaginationMetadata
	(*GetCommentByIDRequest)(nil),    // 5: comment.GetCommentByIDRequest
	(*GetCommentByIDResponse)(nil),   // 6: comment.GetCommentByIDResponse
	(*ListPostCommentsRequest)(nil),  // 7: comment.ListPostCommentsRequest
	(*ListPostCommentsResponse)(nil), // 8: comment.ListPostCommentsResponse
	(*fieldmaskpb.FieldMask)(nil),    // 9: google.protobuf.FieldMask
}
var file_comment_comment_service_proto_depIdxs = []int32{
	3, // 0: comment.CommentObject.user:type_name -> comment.UserObject
	2, // 1: comment.GetCommentByIDResponse.comment:type_name -> comment.CommentObject
	0, // 2: comment.ListPostCommentsRequest.sort_by:type_name -> comment.SortBy
	1, // 3: comment.ListPostCommentsRequest.sort_order:type_name -> comment.SortOrder
	9, // 4: comment.ListPostCommentsRequest.filter_mask:type_name -> google.protobuf.FieldMask
	2, // 5: comment.ListPostCommentsResponse.comments:type_name -> comment.CommentObject
	4, // 6: comment.ListPostCommentsResponse.pagination_metadata:type_name -> comment.PaginationMetadata
	5, // 7: comment.Comment.GetCommentByID:input_type -> comment.GetCommentByIDRequest
	7, // 8: comment.Comment.ListPostComments:input_type -> comment.ListPostCommentsRequest
	6, // 9: comment.Comment.GetCommentByID:output_type -> comment.GetCommentByIDResponse
	8, // 10: comment.Comment.ListPostComments:output_type -> comment.ListPostCommentsResponse
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_comment_comment_service_proto_init() }
func file_comment_comment_service_proto_init() {
	if File_comment_comment_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_comment_comment_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CommentObject); i {
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
		file_comment_comment_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*UserObject); i {
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
		file_comment_comment_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*PaginationMetadata); i {
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
		file_comment_comment_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetCommentByIDRequest); i {
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
		file_comment_comment_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetCommentByIDResponse); i {
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
		file_comment_comment_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ListPostCommentsRequest); i {
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
		file_comment_comment_service_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ListPostCommentsResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_comment_comment_service_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_comment_comment_service_proto_goTypes,
		DependencyIndexes: file_comment_comment_service_proto_depIdxs,
		EnumInfos:         file_comment_comment_service_proto_enumTypes,
		MessageInfos:      file_comment_comment_service_proto_msgTypes,
	}.Build()
	File_comment_comment_service_proto = out.File
	file_comment_comment_service_proto_rawDesc = nil
	file_comment_comment_service_proto_goTypes = nil
	file_comment_comment_service_proto_depIdxs = nil
}
