package internalhttp

import (
	"net/http"
)

func (s *Server) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.handlerHome)
	mux.HandleFunc("/events", s.handlerEventsList)
	mux.HandleFunc("/eventsadd", s.handlerEventsAddSome)

	return s.logging(mux)
}
