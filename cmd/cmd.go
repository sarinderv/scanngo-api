package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"github.com/runntimeterror/scanngo-api/config"
	_ "github.com/runntimeterror/scanngo-api/docs"
	"github.com/runntimeterror/scanngo-api/server"
	httpSwagger "github.com/swaggo/http-swagger"
)

type flags struct {
	loglvl string
}

func Run(args []string) (err error) {

	logger := NewLogger()

	logger.Info().Msg("Initiating Application")

	var cfg config.Config

	err = cleanenv.ReadConfig("config.yaml", &cfg)
	if err != nil {
		logger.Fatal().Err(err).Msg("Error reqding config file")
	}

	fmt.Println(&cfg)

	db, err := ConnectDatabase(&cfg.Database)
	if err != nil {
		logger.Fatal().Err(err).Msg("sqldb error")
	}

	router := mux.NewRouter()

	s := server.New(router, logger, db)
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), //The url pointing to API definition
	)).Methods(http.MethodGet)

	return s.Driver.ListenAndServe()
}

func ConnectDatabase(cfg *config.Database) (*sqlx.DB, error) {

	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB))

	return db, err
}

func NewLogger() zerolog.Logger {
	lgr := zerolog.New(os.Stdout).Level(zerolog.InfoLevel)
	lgr = lgr.With().Timestamp().Logger().Output(zerolog.ConsoleWriter{Out: os.Stderr})

	return lgr
}
