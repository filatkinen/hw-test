package server

import (
	"context"
	"errors"
	"log"
	"net"

	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/config/server"
	pb "github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/grpcservice/calendar"
	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/grpcservice/common"
	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/server/calendar"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ServiceCalendar struct {
	app    *calendar.App
	logger *log.Logger
}

type GrpcServerCalendar struct {
	srv    *grpc.Server
	srvApp ServiceCalendar
	Addr   string
}

func NewGrpcServerCalendar(app *calendar.App, config server.Config, log *log.Logger) *GrpcServerCalendar {
	return &GrpcServerCalendar{
		Addr: net.JoinHostPort(config.ServiceAddress, config.ServiceGrpcPort),
		srvApp: ServiceCalendar{
			app:    app,
			logger: log,
		},
	}
}

func (s *GrpcServerCalendar) Start() error {
	s.srv = grpc.NewServer(grpc.StreamInterceptor(s.eventStreamServerInterceptor),
		grpc.UnaryInterceptor(s.eventUnaryServerInterceptor))
	lis, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}
	pb.RegisterCalendarEventsServer(s.srv, &s.srvApp)
	if err := s.srv.Serve(lis); err != nil {
		if err != nil && !errors.Is(err, grpc.ErrServerStopped) {
			return err
		}
	}
	return nil
}

func (s *GrpcServerCalendar) Close() {
	s.srv.Stop()
}

func errorMessage(err error, message string, field string) error {
	errorStatus := status.New(codes.Internal, message)
	ds, er := errorStatus.WithDetails(&errdetails.BadRequest_FieldViolation{
		Field:       field,
		Description: err.Error(),
	})
	if er != nil {
		return errorStatus.Err()
	}
	return ds.Err()
}

func getUserID(ctx context.Context) (string, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		_us := md[common.UserIDHeader]
		if len(_us) > 0 {
			return _us[0], nil
		}
		return "", common.ErrUserIDNotSet
	}
	return "", common.ErrUserIDNotSet
}

func (s *ServiceCalendar) AddEvent(ctx context.Context, event *pb.Event) (*pb.EventID, error) {
	ev := common.FromPBToEvent(event)
	eventID, err := s.app.AddEvent(ctx, ev, ev.UserID)
	if err != nil {
		return nil, errorMessage(err, "Error creating event", "")
	}
	return &pb.EventID{Id: eventID}, nil
}

func (s *ServiceCalendar) GetEvent(ctx context.Context, id *pb.EventID) (*pb.Event, error) {
	event, err := s.app.GetEvent(ctx, id.Id)
	if err != nil {
		return nil, errorMessage(err, "Error getting event", id.Id)
	}
	return common.FromEventToPB(event), nil
}

func (s *ServiceCalendar) ChangeEvent(ctx context.Context, event *pb.Event) (*wrappers.BoolValue, error) {
	ev := common.FromPBToEvent(event)
	userID, err := getUserID(ctx)
	if err != nil {
		return nil, errorMessage(err, "Error changing event due empty UserID", ev.ID)
	}
	err = s.app.ChangeEventUser(ctx, ev, userID)
	if err != nil {
		return &wrappers.BoolValue{Value: false}, errorMessage(err, "Error changing event", ev.ID)
	}
	return &wrappers.BoolValue{Value: true}, nil
}

func (s *ServiceCalendar) DeleteEvent(ctx context.Context, id *pb.EventID) (*wrappers.BoolValue, error) {
	userID, err := getUserID(ctx)
	if err != nil {
		return nil, errorMessage(err, "Error deleting event due empty UserID", id.Id)
	}
	err = s.app.DeleteEventUser(ctx, id.Id, userID)
	if err != nil {
		return &wrappers.BoolValue{Value: false}, errorMessage(err, "Error deleting event", id.Id)
	}
	return &wrappers.BoolValue{Value: true}, nil
}

func (s *ServiceCalendar) ListEventsDay(timestamp *timestamppb.Timestamp,
	dayServer pb.CalendarEvents_ListEventsDayServer,
) error {
	userID, err := getUserID(dayServer.Context())
	if err != nil {
		return errorMessage(err, "Error getting list of events due empty UserID", "")
	}
	events, err := s.app.ListEventsDay(dayServer.Context(), timestamp.AsTime(), userID)
	if err != nil {
		return errorMessage(err, "Error getting list of events", "")
	}
	for i := range events {
		err := dayServer.Send(common.FromEventToPB(events[i]))
		if err != nil {
			return errorMessage(err, "Error sending event", events[i].ID)
		}
	}
	return nil
}

func (s *ServiceCalendar) ListEventsWeek(timestamp *timestamppb.Timestamp,
	weekServer pb.CalendarEvents_ListEventsWeekServer,
) error {
	userID, err := getUserID(weekServer.Context())
	if err != nil {
		return errorMessage(err, "Error getting list of events due empty UserID", "")
	}
	events, err := s.app.ListEventsWeek(weekServer.Context(), timestamp.AsTime(), userID)
	if err != nil {
		return errorMessage(err, "Error getting list of events", "")
	}
	for i := range events {
		err := weekServer.Send(common.FromEventToPB(events[i]))
		if err != nil {
			return errorMessage(err, "Error sending event", events[i].ID)
		}
	}
	return nil
}

func (s *ServiceCalendar) ListEventsMonth(timestamp *timestamppb.Timestamp,
	monthServer pb.CalendarEvents_ListEventsMonthServer,
) error {
	userID, err := getUserID(monthServer.Context())
	if err != nil {
		return errorMessage(err, "Error getting list of events due empty UserID", "")
	}
	events, err := s.app.ListEventsMonth(monthServer.Context(), timestamp.AsTime(), userID)
	if err != nil {
		return errorMessage(err, "Error getting list of events", "")
	}
	for i := range events {
		err := monthServer.Send(common.FromEventToPB(events[i]))
		if err != nil {
			return errorMessage(err, "Error sending event", events[i].ID)
		}
	}
	return nil
}

func (s *ServiceCalendar) GetNoticesToSend(_ *timestamppb.Timestamp,
	_ pb.CalendarEvents_GetNoticesToSendServer,
) error {
	return nil
}

func (s *ServiceCalendar) GetNoticesToDelete(_ *timestamppb.Timestamp,
	_ pb.CalendarEvents_GetNoticesToDeleteServer,
) error {
	return nil
}
