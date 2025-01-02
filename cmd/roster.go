package main

import (
	"github.com/leagueify/roster/internal/config"
	"github.com/leagueify/roster/internal/integrations/server"
)

func main() {
	cfg := config.GetConfig()

	server.NewServer(cfg).Start()
}
