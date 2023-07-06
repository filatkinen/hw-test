package server

import (
	"context"
	"log"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

type wrappedStream struct {
	grpc.ServerStream
}

func (w *wrappedStream) RecvMsg(m interface{}) error {
	return w.ServerStream.RecvMsg(m)
}

func (w *wrappedStream) SendMsg(m interface{}) error {
	return w.ServerStream.SendMsg(m)
}

func newWrappedStream(s grpc.ServerStream) grpc.ServerStream {
	return &wrappedStream{s}
}

func (s *GrpcServerCalendar) eventStreamServerInterceptor(srv interface{},
	ss grpc.ServerStream,
	_ *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	err := handler(srv, newWrappedStream(ss))
	if err != nil {
		log.Printf("RPC failed with error %v", err)
	}
	return err
}

func (s *GrpcServerCalendar) eventUnaryServerInterceptor(ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
) (interface{}, error) {
	defer s.logging(ctx, time.Now(), info.FullMethod)
	m, err := handler(ctx, req)
	return m, err
}

func (s *GrpcServerCalendar) logging(ctx context.Context, timeStart time.Time, method string) {
	timeToTakeServ := time.Since(timeStart)
	timelog := timeStart.UTC().Format("02/01/2006 15:04:05 UTC")
	var remoteIP string
	var ua string
	r, ok := peer.FromContext(ctx)
	if ok {
		remoteIP = r.Addr.String()
		remoteIP = remoteIP[0:strings.Index(remoteIP, ":")]
	}
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		_ua := md["user-agent"]
		if len(_ua) > 0 {
			ua = _ua[0]
		}
	}
	s.srvApp.logger.Printf("%s [%s] %s %s %s %s %s %s %s\n",
		remoteIP, timelog, "RPC", method, "-", "-", "-", timeToTakeServ, ua)
}
