package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/justinas/alice"
	"github.com/rs/cors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
	"github.com/runntimeterror/scanngo-api/controllers"
	"github.com/runntimeterror/scanngo-api/service"
)

type Server struct {
	router *mux.Router
	Logger zerolog.Logger
	Driver *http.Server
	DB     *sqlx.DB
}

func New(rtr *mux.Router, lgr zerolog.Logger, db *sqlx.DB) *Server {

	handler := cors.Default().Handler(rtr)
	driver := http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      handler,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	s := &Server{rtr, lgr, &driver, db}

	s.InstantiateServices()
	return s
}

func (s *Server) InstantiateServices() {
	clientService := service.NewClientService(s.DB)
	productService := service.NewProductService(s.DB, s.Logger)

	controllers := []controllers.IController{
		controllers.NewClientController(s.DB, clientService, s.Logger),
		controllers.NewProductController(s.DB, productService, s.Logger),
	}

	for _, c := range controllers {
		subRouter := s.router.PathPrefix(c.GetPath()).Subrouter()

		for _, r := range c.GetRoutes() {
			s.Logger.Info().Msg(fmt.Sprintf("Instantiating route -> %s : %s", c.GetPath()+r.Path, r.Method))
			subRouter.Handle(r.Path,
				s.loggerChain().
					Append(s.jsonContentTypeResponseHandler).
					Then(r.Handler)).
				Methods(r.Method)

			// if r.Method == http.MethodPost || r.Method == http.MethodPut {
			// 	subRouter.Headers(contentTypeHeaderKey, appJSONContentTypeHeaderVal)
			// }
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
