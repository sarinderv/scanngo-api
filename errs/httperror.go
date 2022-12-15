package errs

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog"
)

type Error struct {
	StatusCode int
	Err        error
}

func HTTPErrorResponse(w http.ResponseWriter, lgr zerolog.Logger, err Error) {

	// log the error with stacktrace
	lgr.Error().Stack().Err(err.Err).Msg("")

	// Write Content-Type headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.StatusCode)
	json.NewEncoder(w).Encode(err.Err.Error())

}
