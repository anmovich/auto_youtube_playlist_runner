package server

import (
	"net/http"
	"time"
	handlers "ypp/internal/transport/http"

)

type Server struct{
	HttpServer handlers.Handler
}

func(s *Server) Run(port_site string) error{
	srv := &http.Server{
    Addr:         ":" + port_site,
    Handler:      s.HttpServer.InitRoute(),
    ReadTimeout:  10 * time.Second,
    WriteTimeout: 10 * time.Second,
    IdleTimeout:  60 * time.Second,
}
	return srv.ListenAndServe()
}
