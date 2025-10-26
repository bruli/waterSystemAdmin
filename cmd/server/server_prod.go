//go:build prod

package main

import (
	"errors"
	"net/http"

	"github.com/bruli/waterSystemAdmin/internal/config"
	"github.com/rs/zerolog"
)

func runServer(log zerolog.Logger, conf *config.Config, srv *http.Server) {
	log.Info().Msg("starting server at port " + conf.ServerURL())
	if err := srv.ListenAndServeTLS("/etc/letsencrypt/live/bruli.ddns.net/fullchain.pem", "/etc/letsencrypt/live/bruli.ddns.net/privkey.pem"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal().Err(err).Msg("failed to start server")
	}
}
