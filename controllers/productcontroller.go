package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
	"github.com/runntimeterror/scanngo-api/errs"
	"github.com/runntimeterror/scanngo-api/service"
)

type ProductController struct {
	PathPrefix string
	Logger     zerolog.Logger
	db         *sqlx.DB
	service    service.IProductService
}

func NewProductController(db *sqlx.DB, productService service.IProductService, lgr zerolog.Logger) *ProductController {
	return &ProductController{
		"/product",
		lgr,
		db,
		productService,
	}
}

func (s *ProductController) GetRoutes() []Route {
	routes := []Route{
		{
			Path:    "/search/{clientId}",
			Handler: s.handleProductSearch,
			Method:  http.MethodGet,
		},
		{
			Path:    "/{clientId}",
			Handler: s.handleProductCreate,
			Method:  http.MethodPost,
		},
		{
			Path:    "/{clientId}/{id}",
			Handler: s.handleProductUpdate,
			Method:  http.MethodPut,
		},
		{
			Path:    "/{clientId}/{id}",
			Handler: s.handleProductDelete,
			Method:  http.MethodDelete,
		},
		{
			Path:    "/{clientId}",
			Handler: s.handleProductGetAll,
			Method:  http.MethodGet,
		},
		{
			Path:    "/{clientId}/{id}",
			Handler: s.handleProductGet,
			Method:  http.MethodGet,
		},
	}

	return routes
}

func (s *ProductController) GetPath() string {

	return s.PathPrefix
}

// ProductGet godoc
// @Summary      ProductGet
// @Description  ProductGet
// @Tags         Product
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Failure      404  {string} string
// @Success      200  {string} string
// @Failure      400  {string} string
// @Failure      500  {string} string
// @Router       /Product/{id} [get]
func (s *ProductController) handleProductGet(w http.ResponseWriter, r *http.Request) {
	lgr := *hlog.FromRequest(r)
	vars := mux.Vars(r)
	rawCid := vars["clientId"]

	rawPid := vars["id"]

	if rawCid == "search" {
		s.handleProductSearch(w, r)
		return
	}

	clientID, err := strconv.Atoi(rawCid)
	if err != nil {
		errs.HTTPErrorResponse(w, lgr, errs.Error{
			StatusCode: 400,
			Err:        err,
		})
		return
	}

	if clientID < 1 {

		errs.HTTPErrorResponse(w, lgr, errs.Error{
			StatusCode: 400,
			Err:        errors.New("bad client id"),
		})
		return
	}

	productID, err := strconv.Atoi(rawPid)
	if productID < 1 {

		errs.HTTPErrorResponse(w, lgr, errs.Error{
			StatusCode: 400,
			Err:        errors.New("bad product id"),
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

	product, err := s.service.FindByID(r.Context(), productID, clientID)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			errs.HTTPErrorResponse(w, lgr, errs.Error{
				StatusCode: 404,
				Err:        err,
			})
			return
		}
		errs.HTTPErrorResponse(w, lgr, errs.Error{
			StatusCode: 500, Err: err,
		})
		return
	}

	json.NewEncoder(w).Encode(product)
}

// ProductSearch godoc
// @Summary      ProductSearch
// @Description  ProductSearch
// @Tags         Product
// @Accept       json
// @Produce      json
// @Param        name   query      string  true  "Product Name"
// @Failure      404  {string} string
// @Success      200  {string} string
// @Failure      400  {string} string
// @Failure      500  {string} string
// @Router       /Product/{id} [get]
func (s *ProductController) handleProductSearch(w http.ResponseWriter, r *http.Request) {
	lgr := *hlog.FromRequest(r)

	search := r.URL.Query().Get("code")
	vars := mux.Vars(r)
	rawCid := vars["clientId"]
	clientID, err := strconv.Atoi(rawCid)
	if err != nil {
		errs.HTTPErrorResponse(w, lgr, errs.Error{
			StatusCode: 400,
			Err:        err,
		})
		return
	}

	if clientID < 1 {

		errs.HTTPErrorResponse(w, lgr, errs.Error{
			StatusCode: 400,
			Err:        errors.New("bad client id"),
		})
		return
	}

	products, err := s.service.FindByName(r.Context(), search, clientID)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			errs.HTTPErrorResponse(w, lgr, errs.Error{
				StatusCode: 404,
				Err:        err,
			})
			return
		}
		errs.HTTPErrorResponse(w, lgr, errs.Error{
			StatusCode: 500, Err: err,
		})
		return
	}

	json.NewEncoder(w).Encode(products)
}

// ProductGetAll godoc
// @Summary      ProductGetAll
// @Description  ProductGetAll
// @Tags         Product
// @Accept       json
// @Produce      json
// @Failure      404  {string} string
// @Success      200  {string} string
// @Failure      400  {string} string
// @Failure      500  {string} string
// @Router       /Product [get]
func (s *ProductController) handleProductGetAll(w http.ResponseWriter, r *http.Request) {
	lgr := *hlog.FromRequest(r)
	vars := mux.Vars(r)
	rawCid := vars["clientId"]

	clientID, err := strconv.Atoi(rawCid)
	if err != nil {
		errs.HTTPErrorResponse(w, lgr, errs.Error{
			StatusCode: 400,
			Err:        err,
		})
		return
	}

	if clientID < 1 {

		errs.HTTPErrorResponse(w, lgr, errs.Error{
			StatusCode: 400,
			Err:        errors.New("bad client id"),
		})
		return
	}

	products, err := s.service.FindAll(r.Context(), clientID)
	if err != nil {
		errs.HTTPErrorResponse(w, lgr, errs.Error{
			StatusCode: 500,
			Err:        err,
		})
	}

	json.NewEncoder(w).Encode(products)
}

// ProductCreate godoc
// @Summary      ProductCreate
// @Description  ProductCreate
// @Tags         Product
// @Accept       json
// @Produce      json
// @Failure      404  {string} string
// @Success      200  {string} string
// @Failure      400  {string} string
// @Failure      500  {string} string
// @Router       /Product [post]
func (s *ProductController) handleProductCreate(w http.ResponseWriter, r *http.Request) {
	lgr := *hlog.FromRequest(r)

	rb := []*service.Product{}
	// err := json.NewDecoder(r.Body).Decode(&rb)
	// defer r.Body.Close()
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

	err = s.service.Create(r.Context(), tx, rb)
	if err != nil {
		tx.Rollback()
		errs.HTTPErrorResponse(w, lgr, errs.Error{
			StatusCode: 500,
			Err:        err,
		})
		return
	}
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

// ProductUpdate godoc
// @Summary      ProductUpdate
// @Description  ProductUpdate
// @Tags         Product
// @Accept       json
// @Produce      json
// @Failure      404  {string} string
// @Success      200  {string} string
// @Failure      400  {string} string
// @Failure      500  {string} string
// @Router       /Product [put]
func (s *ProductController) handleProductUpdate(w http.ResponseWriter, r *http.Request) {
	lgr := *hlog.FromRequest(r)

	vars := mux.Vars(r)

	rb := service.Product{}
	err := json.NewDecoder(r.Body).Decode(&rb)
	defer r.Body.Close()

	if err != nil {
		errs.HTTPErrorResponse(w, lgr, errs.Error{
			StatusCode: 400,
			Err:        err,
		})
		return
	}

	rawCid := vars["clientId"]

	rawPid := vars["id"]

	clientID, err := strconv.Atoi(rawCid)
	if err != nil {
		errs.HTTPErrorResponse(w, lgr, errs.Error{
			StatusCode: 400,
			Err:        err,
		})
		return
	}

	if clientID < 1 {

		errs.HTTPErrorResponse(w, lgr, errs.Error{
			StatusCode: 400,
			Err:        errors.New("bad client id"),
		})
		return
	}

	productID, err := strconv.Atoi(rawPid)
	if productID < 1 {

		errs.HTTPErrorResponse(w, lgr, errs.Error{
			StatusCode: 400,
			Err:        errors.New("bad product id"),
		})
		return
	}

	a := &sql.TxOptions{}

	tx, err := s.db.BeginTxx(r.Context(), a)
	if err != nil {
		errs.HTTPErrorResponse(w, lgr, errs.Error{
			StatusCode: 500,
			Err:        errors.Wrap(err, "failed to start transaction"),
		})
		return
	}

	if err != nil {
		errs.HTTPErrorResponse(w, lgr, errs.Error{
			StatusCode: 500,
			Err:        err,
		})
		return
	}

	s.Logger.Debug().Msg("Starting transaction")

	err = s.service.Update(r.Context(), tx, &rb, clientID)
	if err != nil {
		tx.Rollback()
		errs.HTTPErrorResponse(w, lgr, errs.Error{
			StatusCode: 500,
			Err:        err,
		})
		return
	}

	if err != nil {
		errs.HTTPErrorResponse(w, lgr, errs.Error{
			StatusCode: 500,
			Err:        err,
		})
		return
	}

	tx.Commit()

	json.NewEncoder(w).Encode(map[string]bool{"response": true})
}

// ProductDelete godoc
// @Summary      ProductDelete
// @Description  ProductDelete
// @Tags         Product
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Failure      404  {string} string
// @Success      200  {string} string
// @Failure      400  {string} string
// @Failure      500  {string} string
// @Router       /Product/{id} [delete]
func (s *ProductController) handleProductDelete(w http.ResponseWriter, r *http.Request) {
	lgr := *hlog.FromRequest(r)

	vars := mux.Vars(r)
	rawCid := vars["id"]

	productID, err := strconv.Atoi(rawCid)
	if productID < 2 {

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

	err = s.service.Delete(r.Context(), tx, productID)
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
