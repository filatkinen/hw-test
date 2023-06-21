package internalhttp

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/storage"
)

func (s *Server) handlerHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	io.WriteString(w, "Hello!\n")
}

func (s *Server) handlerEventsList(w http.ResponseWriter, r *http.Request) {
	from := time.Now().UTC().Round(0).Add(-time.Hour * 24 * 365)
	to := time.Now().UTC().Round(0).Add(time.Hour * 24 * 365)
	events, err := s.app.ListEvents(r.Context(), from, to)
	if err != nil {
		if err != nil {
			http.Error(w, "Error getting events "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
	str := strings.Builder{}
	str.WriteString("Events list:\n\n")
	for i := range events {
		str.WriteString(fmt.Sprintf("%s\n", *events[i]))
	}
	io.WriteString(w, str.String())
}

func (s *Server) handlerEventsAddSome(w http.ResponseWriter, r *http.Request) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano())) //nolint:gosec
	idUUID, err := storage.UUID()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	useridUUID, err := storage.UUID()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	event := storage.Event{
		ID:             idUUID,
		Title:          strconv.FormatInt(int64(rnd.Intn(100)), 10) + " Task",
		Description:    "",
		DateTimeStart:  time.Now().Add(time.Second * time.Duration(rnd.Intn(1000))),
		DateTimeEnd:    time.Now().Add(time.Second * time.Duration(rnd.Intn(1000))),
		DateTimeNotice: time.Now().Add(time.Second * time.Duration(rnd.Intn(1000))),
		UserID:         useridUUID,
	}
	err = s.app.CreateEvent(r.Context(), &event)
	if err != nil {
		http.Error(w, "Error creating event "+err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, "Events add\n")
}
