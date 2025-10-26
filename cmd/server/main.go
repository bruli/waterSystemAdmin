package main

import (
	"context"
	"io/fs"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/bruli/waterSystemAdmin/internal/domain/password"
	"github.com/bruli/waterSystemAdmin/internal/infra/disk"

	"github.com/gorilla/sessions"

	"github.com/bruli/waterSystemAdmin/internal/domain/execution"
	"github.com/bruli/waterSystemAdmin/internal/domain/logs"
	"github.com/bruli/waterSystemAdmin/internal/domain/programs"
	"github.com/bruli/waterSystemAdmin/internal/domain/zones"
	"github.com/bruli/waterSystemAdmin/internal/infra/http/templates"

	"github.com/bruli/waterSystemAdmin/internal/config"
	"github.com/bruli/waterSystemAdmin/internal/domain/status"
	"github.com/bruli/waterSystemAdmin/internal/infra/api"
	"github.com/bruli/waterSystemAdmin/internal/infra/http/controller"
	pongo2 "github.com/flosch/pongo2/v6"
	"github.com/rs/zerolog"
)

func main() {
	log := buildLogger()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	conf, err := config.New()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load config")
		os.Exit(1)
	}

	cl := api.NewClient(conf.ApiAuthKey(), conf.ApiUrl(), 5*time.Second)
	statusRepo := api.NewStatusRepository(cl)
	logsRepo := api.NewLogRepository(cl)
	programsRepo := api.NewAllProgramsRepository(cl)
	executionRepo := api.NewExecuteZoneRepository(cl)
	activateRepo := api.NewActivateRepository(cl)
	zoneRepo := api.NewZoneRepository(cl)
	weeklyRepo := api.NewWeeklyRepository(cl)
	tempRepo := api.NewTemperatureRepository(cl)
	passwordRepo := disk.NewPasswordRepository(conf.PasswordFile())

	findZones := zones.NewFindZones(zoneRepo)
	passwordExists := password.NewExists(passwordRepo)

	tplSet := pongo2.NewSet("embed", templates.NewEmbedLoader(templates.FS, "."))

	sessionStore := sessions.NewCookieStore([]byte("clau-secreta-super-segura"))

	authMdw := controller.AuthMiddleware(sessionStore, passwordExists, log)

	// fs := http.FileServer(http.Dir("internal/infra/http/templates/static"))
	staticFiles, _ := fs.Sub(templates.FS, "static")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFiles))))

	http.HandleFunc("/", authMdw(controller.FindStatus(tplSet, status.NewFindStatus(statusRepo), log)))
	http.HandleFunc("/logs", authMdw(controller.FindLogs(tplSet, logs.NewFindLogs(logsRepo), log)))
	http.HandleFunc("/update-status", authMdw(controller.UpdateStatus(status.NewUpdateStatus(statusRepo), log)))
	http.HandleFunc("/programs", authMdw(controller.Programs(tplSet, programs.NewFindAllPrograms(programsRepo), log)))
	http.HandleFunc("/programs/new", authMdw(controller.CreateProgram(tplSet, findZones, programs.NewCreate(programsRepo), log)))
	http.HandleFunc("/programs/delete/{hour}/{type}", authMdw(controller.RemoveProgram(tplSet, programs.NewRemove(programsRepo), log)))
	http.HandleFunc("/programs/weekly/new", authMdw(controller.CreateWeeklyProgram(tplSet, findZones, programs.NewCreateWeekly(weeklyRepo), log)))
	http.HandleFunc("/programs/weekly/delete/{weekday}", authMdw(controller.RemoveWeeklyProgram(tplSet, programs.NewRemoveWeekly(weeklyRepo), log)))
	http.HandleFunc("/programs/temperature/new", authMdw(controller.CreateTemperatureProgram(tplSet, findZones, programs.NewCreateTemperature(tempRepo), log)))
	http.HandleFunc("/programs/temperature/delete/{temperature}", authMdw(controller.RemoveTemperatureProgram(tplSet, programs.NewRemoveTemperature(tempRepo), log)))
	http.HandleFunc("/zones", authMdw(controller.Zones(tplSet, findZones, log)))
	http.HandleFunc("/execution/{id}", authMdw(controller.Execution(tplSet, execution.NewExecuteZone(executionRepo), log)))
	http.HandleFunc("/deactivate", authMdw(controller.Deactivate(status.NewActivateDeactivate(activateRepo), log)))
	http.HandleFunc("/activate", authMdw(controller.Activate(status.NewActivateDeactivate(activateRepo), log)))
	http.HandleFunc("/zones/new", authMdw(controller.CreateZone(tplSet, zones.NewCreate(zoneRepo), log)))
	http.HandleFunc("/zones/{id}/delete", authMdw(controller.DeleteZone(zones.NewDelete(zoneRepo))))
	http.HandleFunc("/zones/{id}/edit", authMdw(controller.UpdateZone(tplSet, findZones, zones.NewUpdate(zoneRepo), log)))
	http.HandleFunc("/login", controller.Login(tplSet, sessionStore, password.NewCheck(passwordRepo), log))
	http.HandleFunc("/logout", controller.Logout(sessionStore, log))
	http.HandleFunc("/password", controller.Password(tplSet, password.NewCreate(passwordRepo), log))

	srv := &http.Server{
		Addr:    conf.ServerURL(),
		Handler: nil,
	}

	go runServer(log, conf, srv)

	<-ctx.Done()
	log.Info().Msg("shutting down server...")
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Error().Err(err).Msg("failed to shutdown server")
	}
}

func buildLogger() zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()
	return log
}
