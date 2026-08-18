package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/morizkay/course-imrenagi-logging-challenge/functions"
	"github.com/morizkay/course-imrenagi-logging-challenge/logs"
	"github.com/rs/zerolog/log"
)

func main() {
	// Create a context and a cancel function
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	// Create a channel to listen for OS signals
	ch := make(chan os.Signal, 1)
	// Listen for OS signals and cancel the context
	signal.Notify(ch, os.Interrupt)
	// Listen for SIGTERM signal and cancel the context
	signal.Notify(ch, syscall.SIGTERM)

	// Start a goroutine to listen for OS signals and cancel the context
	go func() {
		oscall := <-ch
		log.Warn().Msgf("system call:%+v", oscall)
		cancel()
	}()

	// Create a new router
	r := mux.NewRouter()
	// Handle the root path
	r.HandleFunc("/", functions.Handler)

	// start: set up any of your logger configuration here if necessary
	logOutput := os.Getenv("LOG_OUTPUT") // stdout | file | both
	logFile := os.Getenv("LOG_FILE")
	if logFile == "" {
		logFile = "logs/app.log"
	}

	infoLogger, err := logs.NewLogger(logs.LoggerConfig{
		Level:      "info",
		TimeFormat: "2006-01-02 15:04:05",
		Caller:     true,
		StackTrace: true,
		Output:     logOutput,
		FilePath:   logFile,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to initialize logger")
	}
	log.Logger = infoLogger.Logger

	infoLogger.Info().Str("log_output", logOutput).Str("log_file", logFile).Msg("Welcome to the logging challenge!")
	infoLogger.Warn().Msg("This is a warning message")
	infoLogger.Error().Msg("This is an error message")
	// end: set up any of your logger configuration here

	// Create a new HTTP server
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// Start a goroutine to listen for HTTP requests
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("failed to listen and serve http server")
		}
	}()
	<-ctx.Done()

	if err := server.Shutdown(context.Background()); err != nil {
		log.Error().Err(err).Msg("failed to shutdown http server gracefully")
	}
}
