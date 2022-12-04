package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
	"github.com/runntimeterror/scanngo-api/errs"
	"github.com/runntimeterror/scanngo-api/service"
)

type ClientController struct {
	PathPrefix string
	Logger     zerolog.Logger
	db         *sqlx.DB
	service    service.IClientService
}

func NewClientController(db *sqlx.DB, clientService service.IClientService, lgr zerolog.Logger) *ClientController {
	return &ClientController{
		"/client",
		lgr,
		db,
		clientService,
	}
}

type Route struct {
	Path    string
	Handler http.HandlerFunc
	Method  string
}

func (s *ClientController) GetRoutes() []Route {
	routes := []Route{
		{
			Path:    "",
			Handler: s.handleClientCreate,
			Method:  http.MethodPost,
		},
		{
			Path:    "",
			Handler: s.handleClientUpdate,
			Method:  http.MethodPut,
		},
		{
			Path:    "/{id}",
			Handler: s.handleClientDelete,
			Method:  http.MethodDelete,
		},
		{
			Path:    "",
			Handler: s.handleClientGetAll,
			Method:  http.MethodGet,
		},
		{
			Path:    "/{id}",
			Handler: s.handleClientGet,
			Method:  http.MethodGet,
		},
	}

	return routes
}

func (s *ClientController) GetPath() string {

	return s.PathPrefix
}

// ClientGet godoc
// @Summary      ClientGet
// @Description  ClientGet
// @Tags         Client
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Client ID"
// @Failure      404  {string} string
// @Success      200  {string} string
// @Failure      400  {string} string
// @Failure      500  {string} string
// @Router       /client/{id} [get]
func (s *ClientController) handleClientGet(w http.ResponseWriter, r *http.Request) {
	lgr := *hlog.FromRequest(r)

	vars := mux.Vars(r)
	rawCid := vars["id"]

	clientID, err := strconv.Atoi(rawCid)
	if clientID < 2 {

		errs.HTTPErrorResponse(w, lgr, errs.Error{
			StatusCode: 400,
			Err:        errors.New("Bad client id"),
		})
		return
	}

	if err != nil {

		errs.HTTPErrorResponse(w, lgr, errs.Error{
			StatusCode: 400,
			Err:        err,
		})
		return
	}

	client, err := s.service.FindByID(r.Context(), clientID)
	if err != nil {

		if err.Error() == "sql: no rows in result set" {

			errs.HTTPErrorResponse(w, lgr, errs.Error{
				StatusCode: 404,
				Err:        err,
			})
			return
		}
		errs.HTTPErrorResponse(w, lgr, errs.Error{
			StatusCode: 500,
			Err:        err,
		})
		return
	}

	json.NewEncoder(w).Encode(client)
}

// ClientGetAll godoc
// @Summary      ClientGetAll
// @Description  ClientGetAll
// @Tags         Client
// @Accept       json
// @Produce      json
// @Failure      404  {string} string
// @Success      200  {string} string
// @Failure      400  {string} string
// @Failure      500  {string} string
// @Router       /client [get]
func (s *ClientController) handleClientGetAll(w http.ResponseWriter, r *http.Request) {
	lgr := *hlog.FromRequest(r)

	clients, err := s.service.FindAll(r.Context())
	if err != nil {
		errs.HTTPErrorResponse(w, lgr, errs.Error{
			StatusCode: 500,
			Err:        err,
		})
	}

	json.NewEncoder(w).Encode(clients)
}

// ClientCreate godoc
// @Summary      ClientCreate
// @Description  ClientCreate
// @Tags         Client
// @Accept       json
// @Produce      json
// @Failure      404  {string} string
// @Success      200  {string} string
// @Failure      400  {string} string
// @Failure      500  {string} string
// @Router       /client [post]
func (s *ClientController) handleClientCreate(w http.ResponseWriter, r *http.Request) {
	lgr := *hlog.FromRequest(r)

	rb := service.Client{}
	err := json.NewDecoder(r.Body).Decode(&rb)
	defer r.Body.Close()

	if err != nil {
		errs.HTTPErrorResponse(w, lgr, errs.Error{
			StatusCode: 400,
			Err:        err,
		})
		return
	}

	tx := s.db.MustBegin()
	if err != nil {
		errs.HTTPErrorResponse(w, lgr, errs.Error{
			StatusCode: 500,
			Err:        err,
		})
		return
	}

	// Enrich input
	rb.Insert_User = 1

	clnt, err := s.service.Create(r.Context(), tx, &rb)
	if err != nil {
		tx.Rollback()
		errs.HTTPErrorResponse(w, lgr, errs.Error{
			StatusCode: 500,
			Err:        err,
		})
		return
	}

	tx.Commit()

	json.NewEncoder(w).Encode(clnt)
}

// ClientUpdate godoc
// @Summary      ClientUpdate
// @Description  ClientUpdate
// @Tags         Client
// @Accept       json
// @Produce      json
// @Failure      404  {string} string
// @Success      200  {string} string
// @Failure      400  {string} string
// @Failure      500  {string} string
// @Router       /client [put]
func (s *ClientController) handleClientUpdate(w http.ResponseWriter, r *http.Request) {
	lgr := *hlog.FromRequest(r)

	rb := service.Client{}
	err := json.NewDecoder(r.Body).Decode(&rb)
	defer r.Body.Close()

	if err != nil {
		errs.HTTPErrorResponse(w, lgr, errs.Error{
			StatusCode: 400,
			Err:        err,
		})
		return
	}

	tx := s.db.MustBegin()
	if err != nil {
		errs.HTTPErrorResponse(w, lgr, errs.Error{
			StatusCode: 500,
			Err:        err,
		})
		return
	}

	// Enrich input
	if !rb.Active {
		rb.DeactivationDate = sql.NullTime{time.Now(), true}
	}
	rb.Update_User = sql.NullInt32{1, true}

	err = s.service.Update(r.Context(), tx, &rb)
	if err != nil {
		tx.Rollback()
		errs.HTTPErrorResponse(w, lgr, errs.Error{
			StatusCode: 500,
			Err:        err,
		})
		return
	}

	tx.Commit()

	json.NewEncoder(w).Encode(map[string]bool{"response": true})
}

// ClientDelete godoc
// @Summary      ClientDelete
// @Description  ClientDelete
// @Tags         Client
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Client ID"
// @Failure      404  {string} string
// @Success      200  {string} string
// @Failure      400  {string} string
// @Failure      500  {string} string
// @Router       /client/{id} [delete]
func (s *ClientController) handleClientDelete(w http.ResponseWriter, r *http.Request) {
	lgr := *hlog.FromRequest(r)

	vars := mux.Vars(r)
	rawCid := vars["id"]

	clientID, err := strconv.Atoi(rawCid)
	if clientID < 2 {

		errs.HTTPErrorResponse(w, lgr, errs.Error{
			StatusCode: 400,
			Err:        errors.New("Bad client id"),
		})
		return
	}

	if err != nil {

		errs.HTTPErrorResponse(w, lgr, errs.Error{
			StatusCode: 400,
			Err:        err,
		})
		return
	}

	tx := s.db.MustBegin()
	if err != nil {
		errs.HTTPErrorResponse(w, lgr, errs.Error{
			StatusCode: 500,
			Err:        err,
		})
		return
	}

	err = s.service.Delete(r.Context(), tx, clientID)
	if err != nil {
		tx.Rollback()
		errs.HTTPErrorResponse(w, lgr, errs.Error{
			StatusCode: 500,
			Err:        err,
		})
		return
	}

	tx.Commit()
	json.NewEncoder(w).Encode("success")
}
