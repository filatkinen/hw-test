package internalhttp

import (
	"net/http"
)

func (s *Server) rootHandler() http.Handler {
	return s.logging(http.HandlerFunc(s.handler))
}

func (s *Server) handler(w http.ResponseWriter, r *http.Request) {
	head, _ := shiftPath(r.URL.Path)
	switch head {
	case "event":
		s.handlerEvent(w, r)
	case "events":
		s.handlerEvents(w, r)
	case "":
		w.WriteHeader(http.StatusOK)
	default:
		http.NotFound(w, r)
	}
}
