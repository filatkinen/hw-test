package common

import (
	"errors"

	pb "github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/grpcservice/calendar"
	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/storage"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var ErrUserIDNotSet = errors.New("userID is empty in rpc metadata")

const UserIDHeader = "user_id"

func FromPBToEvent(event *pb.Event) *storage.Event {
	return &storage.Event{
		ID:             event.Id,
		Title:          event.Title,
		Description:    event.Description,
		DateTimeStart:  event.DateTimeStart.AsTime(),
		DateTimeEnd:    event.DateTimeEnd.AsTime(),
		DateTimeNotice: event.DateTimeNotice.AsTime(),
		UserID:         event.UserId,
	}
}

func FromEventToPB(event *storage.Event) *pb.Event {
	return &pb.Event{
		Id:             event.ID,
		Title:          event.Title,
		Description:    event.Description,
		DateTimeStart:  timestamppb.New(event.DateTimeStart),
		DateTimeEnd:    timestamppb.New(event.DateTimeEnd),
		DateTimeNotice: timestamppb.New(event.DateTimeNotice),
		UserId:         event.UserID,
	}
}

func FromPBToNotice(event *pb.Notice) *storage.Notice {
	return &storage.Notice{
		ID:       event.Id,
		Title:    event.Title,
		DateTime: event.DateTimeStart.AsTime(),
		UserID:   event.UserId,
	}
}

func FromNoticeToPB(event *storage.Notice) *pb.Notice {
	return &pb.Notice{
		Id:            event.ID,
		Title:         event.Title,
		DateTimeStart: timestamppb.New(event.DateTime),
		UserId:        event.UserID,
	}
}
