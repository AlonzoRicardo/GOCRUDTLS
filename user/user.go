package user

import (
	"log"
	"net/http"
	"time"
)

const message = "Hello User!"

type Handlers struct {
	logger *log.Logger
}

func (h *Handlers) User(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}

func (h *Handlers) Users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}

// Logger configure logging
func (h *Handlers) Logger(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer h.logger.Printf("request processed in %s\n", time.Now().Sub(startTime))
		handler(w, r)
	}
}

// SetupRoutes To actually enable the routes
func (h *Handlers) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/user", h.Logger(h.User))
	mux.HandleFunc("/users", h.Logger(h.Users))
}

// NewHandlers Returns a Handler struct
func NewHandlers(logger *log.Logger) *Handlers {
	return &Handlers{
		logger: logger,
	}
}
