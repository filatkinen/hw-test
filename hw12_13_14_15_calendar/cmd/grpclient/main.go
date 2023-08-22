package main

import (
	"context"
	"fmt"
	"time"

	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/grpcservice/client"
	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/storage"
)

func main() {
	eventID, _ := storage.UUID()
	userID, _ := storage.UUID()
	event := storage.Event{
		ID:             eventID,
		Title:          "Task1",
		Description:    "Desc1",
		DateTimeStart:  time.Date(2022, 1, 1, 1, 1, 1, 1, time.UTC),
		DateTimeEnd:    time.Now(),
		DateTimeNotice: time.Now(),
		UserID:         userID,
	}
	cl := client.NewGrpcClientCalendar("localhost:50051")
	err := cl.Start()
	if err != nil {
		fmt.Printf("error starting %s\n", err)
		return
	}
	defer cl.Close()
	fmt.Println(eventID)
	newEventID, err := cl.AddEvent(context.Background(), &event)
	if err != nil {
		fmt.Printf("error adding %s\n", err)
		return
	}
	fmt.Println(newEventID)
	eventNew, err := cl.GetEvent(context.Background(), newEventID)
	if err != nil {
		fmt.Printf("error getting %s\n", err)
		return
	}
	fmt.Printf("%v\n", *eventNew)
}
