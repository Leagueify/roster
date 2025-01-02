package server

import (
	"fmt"
	"net/http"

	"github.com/leagueify/roster/handler"
	"github.com/leagueify/roster/internal/config"
	"github.com/leagueify/roster/internal/middleware"
)

type server struct {
	cfg *config.Config
}

func NewServer(cfg *config.Config) Server {
	return &server{
		cfg: cfg,
	}
}

func (s *server) Start() {
	router := http.NewServeMux()
	rosterRouter := handler.RosterRouter()

	router.Handle("/roster/", http.StripPrefix("/roster", rosterRouter))

	// define server config
	rosterServer := &http.Server{
		Addr:         fmt.Sprintf(":%d", s.cfg.Server.Port),
		Handler:      middleware.Logging(router),
		IdleTimeout:  s.cfg.Server.IdleTimeout,
		ReadTimeout:  s.cfg.Server.ReadTimeout,
		WriteTimeout: s.cfg.Server.WriteTimeout,
	}

	// show server banner
	showStartBanner()

	// start server
	if err := rosterServer.ListenAndServe(); err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
