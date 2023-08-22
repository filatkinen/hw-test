package client

import (
	"context"
	"errors"
	"io"
	"time"

	pb "github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/grpcservice/calendar"
	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/grpcservice/common"
	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GrpcClientCalendar struct {
	client pb.CalendarEventsClient
	conn   *grpc.ClientConn
	addr   string
}

func NewGrpcClientCalendar(address string) *GrpcClientCalendar {
	return &GrpcClientCalendar{addr: address}
}

func (s *GrpcClientCalendar) Start() error {
	conn, err := grpc.Dial(s.addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	s.conn = conn
	s.client = pb.NewCalendarEventsClient(conn)
	return nil
}

func (s *GrpcClientCalendar) Close() error {
	return s.conn.Close()
}

func (s *GrpcClientCalendar) SetUserIDHeader(ctx context.Context, userID string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, common.UserIDHeader, userID)
}

func (s *GrpcClientCalendar) AddEvent(ctx context.Context, event *storage.Event) (string, error) {
	pbevent := common.FromEventToPB(event)
	req, err := s.client.AddEvent(ctx, pbevent)
	if err != nil {
		return "", err
	}
	return req.Id, nil
}

func (s *GrpcClientCalendar) GetEvent(ctx context.Context, eventID string) (*storage.Event, error) {
	req, err := s.client.GetEvent(ctx, &pb.EventID{Id: eventID})
	if err != nil {
		return nil, err
	}
	return common.FromPBToEvent(req), nil
}

func (s *GrpcClientCalendar) ChangeEvent(ctx context.Context, event *storage.Event) error {
	_, err := s.client.ChangeEvent(ctx, common.FromEventToPB(event))
	return err
}

func (s *GrpcClientCalendar) ChangeEventUser(ctx context.Context, event *storage.Event, userID string) error {
	ctx = s.SetUserIDHeader(ctx, userID)
	return s.ChangeEvent(ctx, event)
}

func (s *GrpcClientCalendar) DeleteEvent(ctx context.Context, eventID string) error {
	_, err := s.client.DeleteEvent(ctx, &pb.EventID{Id: eventID})
	return err
}

func (s *GrpcClientCalendar) DeleteEventUser(ctx context.Context, eventID string, userID string) error {
	ctx = s.SetUserIDHeader(ctx, userID)
	return s.DeleteEvent(ctx, eventID)
}

func (s *GrpcClientCalendar) ListEventsDay(ctx context.Context,
	date time.Time, userID string,
) ([]*storage.Event, error) {
	ctx = s.SetUserIDHeader(ctx, userID)
	var events []*storage.Event
	listStream, err := s.client.ListEventsDay(ctx, timestamppb.New(date))
	if err != nil {
		return nil, err
	}
	for {
		ev, err := listStream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, err
		}
		events = append(events, common.FromPBToEvent(ev))
	}
	return events, nil
}

func (s *GrpcClientCalendar) ListEventsWeek(ctx context.Context,
	date time.Time,
	userID string,
) ([]*storage.Event, error) {
	ctx = s.SetUserIDHeader(ctx, userID)
	var events []*storage.Event
	listStream, err := s.client.ListEventsWeek(ctx, timestamppb.New(date))
	if err != nil {
		return nil, err
	}
	for {
		ev, err := listStream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, err
		}
		events = append(events, common.FromPBToEvent(ev))
	}
	return events, nil
}

func (s *GrpcClientCalendar) ListEventsMonth(ctx context.Context,
	date time.Time,
	userID string,
) ([]*storage.Event, error) {
	ctx = s.SetUserIDHeader(ctx, userID)
	var events []*storage.Event
	listStream, err := s.client.ListEventsMonth(ctx, timestamppb.New(date))
	if err != nil {
		return nil, err
	}
	for {
		ev, err := listStream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, err
		}
		events = append(events, common.FromPBToEvent(ev))
	}
	return events, nil
}

func (s *GrpcClientCalendar) GetNoticesToSend(ctx context.Context,
	onTime time.Time,
) ([]*storage.Notice, error) {
	var notices []*storage.Notice
	listStream, err := s.client.GetNoticesToSend(ctx, timestamppb.New(onTime))
	if err != nil {
		return nil, err
	}
	for {
		notice, err := listStream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, err
		}
		notices = append(notices, common.FromPBToNotice(notice))
	}
	return notices, nil
}

func (s *GrpcClientCalendar) GetNoticesToDelete(ctx context.Context,
	onTime time.Time,
) ([]*storage.Notice, error) {
	var notices []*storage.Notice
	listStream, err := s.client.GetNoticesToDelete(ctx, timestamppb.New(onTime))
	if err != nil {
		return nil, err
	}
	for {
		notice, err := listStream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, err
		}
		notices = append(notices, common.FromPBToNotice(notice))
	}
	return notices, nil
}
