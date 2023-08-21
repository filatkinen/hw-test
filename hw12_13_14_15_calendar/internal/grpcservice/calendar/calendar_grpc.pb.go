// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: calendar.proto

package calendar

import (
	context "context"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CalendarEventsClient is the client API for CalendarEvents service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalendarEventsClient interface {
	AddEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*EventID, error)
	GetEvent(ctx context.Context, in *EventID, opts ...grpc.CallOption) (*Event, error)
	ChangeEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*wrappers.BoolValue, error)
	DeleteEvent(ctx context.Context, in *EventID, opts ...grpc.CallOption) (*wrappers.BoolValue, error)
	ListEventsDay(ctx context.Context, in *timestamppb.Timestamp, opts ...grpc.CallOption) (CalendarEvents_ListEventsDayClient, error)
	ListEventsWeek(ctx context.Context, in *timestamppb.Timestamp, opts ...grpc.CallOption) (CalendarEvents_ListEventsWeekClient, error)
	ListEventsMonth(ctx context.Context, in *timestamppb.Timestamp, opts ...grpc.CallOption) (CalendarEvents_ListEventsMonthClient, error)
	GetNoticesToSend(ctx context.Context, in *timestamppb.Timestamp, opts ...grpc.CallOption) (CalendarEvents_GetNoticesToSendClient, error)
	GetNoticesToDelete(ctx context.Context, in *timestamppb.Timestamp, opts ...grpc.CallOption) (CalendarEvents_GetNoticesToDeleteClient, error)
}

type calendarEventsClient struct {
	cc grpc.ClientConnInterface
}

func NewCalendarEventsClient(cc grpc.ClientConnInterface) CalendarEventsClient {
	return &calendarEventsClient{cc}
}

func (c *calendarEventsClient) AddEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*EventID, error) {
	out := new(EventID)
	err := c.cc.Invoke(ctx, "/calendar.CalendarEvents/addEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarEventsClient) GetEvent(ctx context.Context, in *EventID, opts ...grpc.CallOption) (*Event, error) {
	out := new(Event)
	err := c.cc.Invoke(ctx, "/calendar.CalendarEvents/getEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarEventsClient) ChangeEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*wrappers.BoolValue, error) {
	out := new(wrappers.BoolValue)
	err := c.cc.Invoke(ctx, "/calendar.CalendarEvents/changeEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarEventsClient) DeleteEvent(ctx context.Context, in *EventID, opts ...grpc.CallOption) (*wrappers.BoolValue, error) {
	out := new(wrappers.BoolValue)
	err := c.cc.Invoke(ctx, "/calendar.CalendarEvents/deleteEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarEventsClient) ListEventsDay(ctx context.Context, in *timestamppb.Timestamp, opts ...grpc.CallOption) (CalendarEvents_ListEventsDayClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalendarEvents_ServiceDesc.Streams[0], "/calendar.CalendarEvents/listEventsDay", opts...)
	if err != nil {
		return nil, err
	}
	x := &calendarEventsListEventsDayClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalendarEvents_ListEventsDayClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type calendarEventsListEventsDayClient struct {
	grpc.ClientStream
}

func (x *calendarEventsListEventsDayClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calendarEventsClient) ListEventsWeek(ctx context.Context, in *timestamppb.Timestamp, opts ...grpc.CallOption) (CalendarEvents_ListEventsWeekClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalendarEvents_ServiceDesc.Streams[1], "/calendar.CalendarEvents/listEventsWeek", opts...)
	if err != nil {
		return nil, err
	}
	x := &calendarEventsListEventsWeekClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalendarEvents_ListEventsWeekClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type calendarEventsListEventsWeekClient struct {
	grpc.ClientStream
}

func (x *calendarEventsListEventsWeekClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calendarEventsClient) ListEventsMonth(ctx context.Context, in *timestamppb.Timestamp, opts ...grpc.CallOption) (CalendarEvents_ListEventsMonthClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalendarEvents_ServiceDesc.Streams[2], "/calendar.CalendarEvents/listEventsMonth", opts...)
	if err != nil {
		return nil, err
	}
	x := &calendarEventsListEventsMonthClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalendarEvents_ListEventsMonthClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type calendarEventsListEventsMonthClient struct {
	grpc.ClientStream
}

func (x *calendarEventsListEventsMonthClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calendarEventsClient) GetNoticesToSend(ctx context.Context, in *timestamppb.Timestamp, opts ...grpc.CallOption) (CalendarEvents_GetNoticesToSendClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalendarEvents_ServiceDesc.Streams[3], "/calendar.CalendarEvents/getNoticesToSend", opts...)
	if err != nil {
		return nil, err
	}
	x := &calendarEventsGetNoticesToSendClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalendarEvents_GetNoticesToSendClient interface {
	Recv() (*Notice, error)
	grpc.ClientStream
}

type calendarEventsGetNoticesToSendClient struct {
	grpc.ClientStream
}

func (x *calendarEventsGetNoticesToSendClient) Recv() (*Notice, error) {
	m := new(Notice)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calendarEventsClient) GetNoticesToDelete(ctx context.Context, in *timestamppb.Timestamp, opts ...grpc.CallOption) (CalendarEvents_GetNoticesToDeleteClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalendarEvents_ServiceDesc.Streams[4], "/calendar.CalendarEvents/getNoticesToDelete", opts...)
	if err != nil {
		return nil, err
	}
	x := &calendarEventsGetNoticesToDeleteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalendarEvents_GetNoticesToDeleteClient interface {
	Recv() (*Notice, error)
	grpc.ClientStream
}

type calendarEventsGetNoticesToDeleteClient struct {
	grpc.ClientStream
}

func (x *calendarEventsGetNoticesToDeleteClient) Recv() (*Notice, error) {
	m := new(Notice)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalendarEventsServer is the server API for CalendarEvents service.
// All implementations should embed UnimplementedCalendarEventsServer
// for forward compatibility
type CalendarEventsServer interface {
	AddEvent(context.Context, *Event) (*EventID, error)
	GetEvent(context.Context, *EventID) (*Event, error)
	ChangeEvent(context.Context, *Event) (*wrappers.BoolValue, error)
	DeleteEvent(context.Context, *EventID) (*wrappers.BoolValue, error)
	ListEventsDay(*timestamppb.Timestamp, CalendarEvents_ListEventsDayServer) error
	ListEventsWeek(*timestamppb.Timestamp, CalendarEvents_ListEventsWeekServer) error
	ListEventsMonth(*timestamppb.Timestamp, CalendarEvents_ListEventsMonthServer) error
	GetNoticesToSend(*timestamppb.Timestamp, CalendarEvents_GetNoticesToSendServer) error
	GetNoticesToDelete(*timestamppb.Timestamp, CalendarEvents_GetNoticesToDeleteServer) error
}

// UnimplementedCalendarEventsServer should be embedded to have forward compatible implementations.
type UnimplementedCalendarEventsServer struct {
}

func (UnimplementedCalendarEventsServer) AddEvent(context.Context, *Event) (*EventID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddEvent not implemented")
}
func (UnimplementedCalendarEventsServer) GetEvent(context.Context, *EventID) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvent not implemented")
}
func (UnimplementedCalendarEventsServer) ChangeEvent(context.Context, *Event) (*wrappers.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeEvent not implemented")
}
func (UnimplementedCalendarEventsServer) DeleteEvent(context.Context, *EventID) (*wrappers.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEvent not implemented")
}
func (UnimplementedCalendarEventsServer) ListEventsDay(*timestamppb.Timestamp, CalendarEvents_ListEventsDayServer) error {
	return status.Errorf(codes.Unimplemented, "method ListEventsDay not implemented")
}
func (UnimplementedCalendarEventsServer) ListEventsWeek(*timestamppb.Timestamp, CalendarEvents_ListEventsWeekServer) error {
	return status.Errorf(codes.Unimplemented, "method ListEventsWeek not implemented")
}
func (UnimplementedCalendarEventsServer) ListEventsMonth(*timestamppb.Timestamp, CalendarEvents_ListEventsMonthServer) error {
	return status.Errorf(codes.Unimplemented, "method ListEventsMonth not implemented")
}
func (UnimplementedCalendarEventsServer) GetNoticesToSend(*timestamppb.Timestamp, CalendarEvents_GetNoticesToSendServer) error {
	return status.Errorf(codes.Unimplemented, "method GetNoticesToSend not implemented")
}
func (UnimplementedCalendarEventsServer) GetNoticesToDelete(*timestamppb.Timestamp, CalendarEvents_GetNoticesToDeleteServer) error {
	return status.Errorf(codes.Unimplemented, "method GetNoticesToDelete not implemented")
}

// UnsafeCalendarEventsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalendarEventsServer will
// result in compilation errors.
type UnsafeCalendarEventsServer interface {
	mustEmbedUnimplementedCalendarEventsServer()
}

func RegisterCalendarEventsServer(s grpc.ServiceRegistrar, srv CalendarEventsServer) {
	s.RegisterService(&CalendarEvents_ServiceDesc, srv)
}

func _CalendarEvents_AddEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarEventsServer).AddEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendar.CalendarEvents/addEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarEventsServer).AddEvent(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarEvents_GetEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarEventsServer).GetEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendar.CalendarEvents/getEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarEventsServer).GetEvent(ctx, req.(*EventID))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarEvents_ChangeEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarEventsServer).ChangeEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendar.CalendarEvents/changeEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarEventsServer).ChangeEvent(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarEvents_DeleteEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarEventsServer).DeleteEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendar.CalendarEvents/deleteEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarEventsServer).DeleteEvent(ctx, req.(*EventID))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarEvents_ListEventsDay_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(timestamppb.Timestamp)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalendarEventsServer).ListEventsDay(m, &calendarEventsListEventsDayServer{stream})
}

type CalendarEvents_ListEventsDayServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type calendarEventsListEventsDayServer struct {
	grpc.ServerStream
}

func (x *calendarEventsListEventsDayServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func _CalendarEvents_ListEventsWeek_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(timestamppb.Timestamp)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalendarEventsServer).ListEventsWeek(m, &calendarEventsListEventsWeekServer{stream})
}

type CalendarEvents_ListEventsWeekServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type calendarEventsListEventsWeekServer struct {
	grpc.ServerStream
}

func (x *calendarEventsListEventsWeekServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func _CalendarEvents_ListEventsMonth_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(timestamppb.Timestamp)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalendarEventsServer).ListEventsMonth(m, &calendarEventsListEventsMonthServer{stream})
}

type CalendarEvents_ListEventsMonthServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type calendarEventsListEventsMonthServer struct {
	grpc.ServerStream
}

func (x *calendarEventsListEventsMonthServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func _CalendarEvents_GetNoticesToSend_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(timestamppb.Timestamp)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalendarEventsServer).GetNoticesToSend(m, &calendarEventsGetNoticesToSendServer{stream})
}

type CalendarEvents_GetNoticesToSendServer interface {
	Send(*Notice) error
	grpc.ServerStream
}

type calendarEventsGetNoticesToSendServer struct {
	grpc.ServerStream
}

func (x *calendarEventsGetNoticesToSendServer) Send(m *Notice) error {
	return x.ServerStream.SendMsg(m)
}

func _CalendarEvents_GetNoticesToDelete_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(timestamppb.Timestamp)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalendarEventsServer).GetNoticesToDelete(m, &calendarEventsGetNoticesToDeleteServer{stream})
}

type CalendarEvents_GetNoticesToDeleteServer interface {
	Send(*Notice) error
	grpc.ServerStream
}

type calendarEventsGetNoticesToDeleteServer struct {
	grpc.ServerStream
}

func (x *calendarEventsGetNoticesToDeleteServer) Send(m *Notice) error {
	return x.ServerStream.SendMsg(m)
}

// CalendarEvents_ServiceDesc is the grpc.ServiceDesc for CalendarEvents service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalendarEvents_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calendar.CalendarEvents",
	HandlerType: (*CalendarEventsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "addEvent",
			Handler:    _CalendarEvents_AddEvent_Handler,
		},
		{
			MethodName: "getEvent",
			Handler:    _CalendarEvents_GetEvent_Handler,
		},
		{
			MethodName: "changeEvent",
			Handler:    _CalendarEvents_ChangeEvent_Handler,
		},
		{
			MethodName: "deleteEvent",
			Handler:    _CalendarEvents_DeleteEvent_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "listEventsDay",
			Handler:       _CalendarEvents_ListEventsDay_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "listEventsWeek",
			Handler:       _CalendarEvents_ListEventsWeek_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "listEventsMonth",
			Handler:       _CalendarEvents_ListEventsMonth_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "getNoticesToSend",
			Handler:       _CalendarEvents_GetNoticesToSend_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "getNoticesToDelete",
			Handler:       _CalendarEvents_GetNoticesToDelete_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "calendar.proto",
}