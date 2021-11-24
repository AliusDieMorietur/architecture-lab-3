package main

import (
	"context"
	"fmt"
	"net/http"
)

type ForumApiServer struct {
	Port        int
	HandlerFunc http.HandlerFunc
	server      *http.Server
}

func (s *ForumApiServer) Start() error {
	if s.HandlerFunc == nil {
		return fmt.Errorf("channels HTTP handler is not defined - cannot start")
	}
	if s.Port == 0 {
		return fmt.Errorf("port is not defined")
	}

	handler := new(http.ServeMux)
	handler.HandleFunc("/", s.HandlerFunc)

	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: handler,
	}

	return s.server.ListenAndServe()
}

func (s *ForumApiServer) Stop() error {
	if s.server == nil {
		return fmt.Errorf("server was not started")
	}
	return s.server.Shutdown(context.Background())
}
