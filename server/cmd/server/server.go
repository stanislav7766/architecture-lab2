package main

import (
	"context"
	"fmt"
	"github.com/stanislav7766/architecture-lab2/server/hostels"
	"net/http"
)

type HttpPortNumber int

// HostelApiServer configures necessary handlers and starts listening on a configured port.
type HostelApiServer struct {
	Port HttpPortNumber

	HostelsHandler hostels.HttpHandlerFunc

	server *http.Server
}

// Start will set all handlers and start listening.
// If this methods succeeds, it does not return until server is shut down.
// Returned error will never be nil.
func (s *HostelApiServer) Start() error {
	if s.HostelsHandler == nil {
		return fmt.Errorf("Hostels HTTP handler is not defined - cannot start")
	}
	if s.Port == 0 {
		return fmt.Errorf("port is not defined")
	}

	handler := new(http.ServeMux)
	handler.HandleFunc("/sendStudent", s.HostelsHandler)
	handler.HandleFunc("/getBestHostel", s.HostelsHandler)

	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: handler,
	}
	return s.server.ListenAndServe()
}

// Stops will shut down previously started HTTP server.
func (s *HostelApiServer) Stop() error {
	if s.server == nil {
		return fmt.Errorf("server was not started")
	}
	return s.server.Shutdown(context.Background())
}
