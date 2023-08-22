package internalhttp

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/storage"
)

var (
	ErrUserIDNotSet = errors.New("user id not set in URL")
	ErrBadUserID    = errors.New("user id wrong in URL(must be UUID mask 36 chars)")
)

func (s *Server) handlerEvent(w http.ResponseWriter, r *http.Request) {
	_, tail := shiftPath(r.URL.Path)
	apiMethod, _ := shiftPath(tail)
	userID, err := getUserID(tail)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	event := storage.Event{}
	switch apiMethod {
	case "add":
		// add event
		jsDecoder := json.NewDecoder(r.Body)
		err = jsDecoder.Decode(&event)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		eventID, err := s.app.AddEvent(r.Context(), &event, userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// return eventID in json
		data, err := json.Marshal(eventID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(data)
	case "change":
		jsDecoder := json.NewDecoder(r.Body)
		err = jsDecoder.Decode(&event)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err = s.app.ChangeEventUser(r.Context(), &event, userID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	case "delete":
		// delete event
		var eventID string
		jsDecoder := json.NewDecoder(r.Body)
		err = jsDecoder.Decode(&eventID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err = s.app.DeleteEventUser(r.Context(), eventID, userID); err != nil {
			if errors.Is(err, storage.ErrWrongEventUser) {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	case "get":
		// Get event
		var eventID string
		jsDecoder := json.NewDecoder(r.Body)
		err = jsDecoder.Decode(&eventID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		event, err := s.app.GetEventUser(r.Context(), eventID, userID)
		if err != nil {
			if errors.Is(err, storage.ErrEventIDNotFound) {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		data, err := json.Marshal(event)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(data)
	default:
		http.Error(w, "wrong URL", http.StatusBadRequest)
		return
	}
}

func (s *Server) handlerEvents(w http.ResponseWriter, r *http.Request) {
	_, tail := shiftPath(r.URL.Path)
	apiMethod, _ := shiftPath(tail)
	userID, err := getUserID(tail)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var date time.Time
	jsDecoder := json.NewDecoder(r.Body)
	err = jsDecoder.Decode(&date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var events []*storage.Event
	switch apiMethod {
	case "day":
		events, err = s.app.ListEventsDay(r.Context(), date, userID)
	case "week":
		events, err = s.app.ListEventsWeek(r.Context(), date, userID)
	case "month":
		events, err = s.app.ListEventsMonth(r.Context(), date, userID)
	default:
		http.Error(w, "wrong URL", http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(events)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(data)
}
