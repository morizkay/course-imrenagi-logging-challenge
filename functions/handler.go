package functions

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	requestLogger := log.With().
		Str("request_id", uuid.New().String()).
		Logger()
	ctx := requestLogger.WithContext(r.Context())

	query := r.URL.Query().Get("q")
	name := r.URL.Query().Get("name")

	requestLogger.Info().
		Str("method", r.Method).
		Str("path", r.URL.Path).
		Str("host", r.Host).
		Str("query", query).
		Str("name", name).
		Msg("request received")

	res, err := Greeting(ctx, name)
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("failed to generate greeting")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(res))
}
