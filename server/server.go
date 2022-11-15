package server

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/justinas/alice"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
	"github.com/scanngo-api/controllers"
	"github.com/scanngo-api/service"
)

type Server struct {
	router *mux.Router
	Logger zerolog.Logger
	Driver *http.Server
	DB     *sqlx.DB
}

func New(rtr *mux.Router, lgr zerolog.Logger, db *sqlx.DB) *Server {

	driver := http.Server{
		Addr:         "127.0.0.1:8080",
		Handler:      rtr,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	s := &Server{rtr, lgr, &driver, db}
	rtr.Use(s.jsonContentTypeResponseHandler)
	s.InstantiateServices()
	return s
}

func (s *Server) InstantiateServices() {
	clientService := service.NewClientService(s.DB)
	productService := service.NewProductService(s.DB)

	controllers := []controllers.IController{
		controllers.NewClientController(s.DB, clientService, s.Logger),
		controllers.NewProductController(s.DB, productService, s.Logger),
	}

	for _, c := range controllers {
		subRouter := s.router.PathPrefix(c.GetPath()).Subrouter()

		for _, r := range c.GetRoutes() {

			subRouter.Handle(r.Path,
				s.loggerChain().
					Append(s.jsonContentTypeResponseHandler).
					Then(r.Handler)).
				Methods(r.Method)

			if r.Method == http.MethodPost || r.Method == http.MethodPut {
				subRouter.Headers(contentTypeHeaderKey, appJSONContentTypeHeaderVal)
			}
		}

	}

}

func (s *Server) jsonContentTypeResponseHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add(contentTypeHeaderKey, appJSONContentTypeHeaderVal)
			h.ServeHTTP(w, r) // call original
		})
}

func (s *Server) loggerChain() alice.Chain {
	ac := alice.New(hlog.NewHandler(s.Logger),
		hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
			hlog.FromRequest(r).Info().
				Str("method", r.Method).
				Stringer("url", r.URL).
				Int("status", status).
				Int("size", size).
				Dur("duration", duration).
				Msg("request logged")
		}),
		hlog.RemoteAddrHandler("remote_ip"),
		hlog.UserAgentHandler("user_agent"),
		hlog.RefererHandler("referer"),
		hlog.RequestIDHandler("request_id", "Request-Id"),
	)

	return ac
}
