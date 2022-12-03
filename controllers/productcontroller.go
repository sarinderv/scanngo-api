package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
	"github.com/scanngo-api/errs"
)

type ProductController struct {
	PathPrefix string
	Logger     zerolog.Logger
	db         *sqlx.DB
	// service    service.IProductService
}

func NewProductController(db *sqlx.DB, lgr zerolog.Logger) *ProductController {
	return &ProductController{
		"/Product",
		lgr,
		db,
	}
}

func (s *ProductController) GetRoutes() []Route {
	routes := []Route{
		{
			Path:    "",
			Handler: s.handleProductCreate,
			Method:  http.MethodPost,
		},
		{
			Path:    "",
			Handler: s.handleProductUpdate,
			Method:  http.MethodPut,
		},
		{
			Path:    "/{id}",
			Handler: s.handleProductDelete,
			Method:  http.MethodDelete,
		},
		{
			Path:    "",
			Handler: s.handleProductGetAll,
			Method:  http.MethodGet,
		},
		{
			Path:    "/{id}",
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
	rawCid := vars["id"]

	ProductID, err := strconv.Atoi(rawCid)
	if ProductID < 2 {

		errs.HTTPErrorResponse(w, lgr, errs.Error{
			StatusCode: 400,
			Err:        errors.New("Bad Product id"),
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
	var Product []int
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
	json.NewEncoder(w).Encode(Product)
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
	// lgr := *hlog.FromRequest(r)

	var Products []int

	json.NewEncoder(w).Encode(Products)
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
	// lgr := *hlog.FromRequest(r)

	// rb := service.Product{}
	// err := json.NewDecoder(r.Body).Decode(&rb)
	// defer r.Body.Close()

	// if err != nil {
	// 	errs.HTTPErrorResponse(w, lgr, errs.Error{
	// 		StatusCode: 400,
	// 		Err:        err,
	// 	})
	// 	return
	// }

	// tx := s.db.MustBegin()
	// if err != nil {
	// 	errs.HTTPErrorResponse(w, lgr, errs.Error{
	// 		StatusCode: 500,
	// 		Err:        err,
	// 	})
	// 	return
	// }

	// // Enrich input
	// rb.Insert_User = 1

	// clnt, err := s.service.Create(r.Context(), tx, &rb)
	// if err != nil {
	// 	tx.Rollback()
	// 	errs.HTTPErrorResponse(w, lgr, errs.Error{
	// 		StatusCode: 500,
	// 		Err:        err,
	// 	})
	// 	return
	// }

	// tx.Commit()

	json.NewEncoder(w).Encode("clnt")
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
	// lgr := *hlog.FromRequest(r)

	// rb := service.Product{}
	// err := json.NewDecoder(r.Body).Decode(&rb)
	// defer r.Body.Close()

	// if err != nil {
	// 	errs.HTTPErrorResponse(w, lgr, errs.Error{
	// 		StatusCode: 400,
	// 		Err:        err,
	// 	})
	// 	return
	// }

	// tx := s.db.MustBegin()
	// if err != nil {
	// 	errs.HTTPErrorResponse(w, lgr, errs.Error{
	// 		StatusCode: 500,
	// 		Err:        err,
	// 	})
	// 	return
	// }

	// // Enrich input
	// if !rb.Active {
	// 	rb.DeactivationDate = sql.NullTime{time.Now(), true}
	// }
	// rb.Update_User = sql.NullInt32{1, true}

	// err = s.service.Update(r.Context(), tx, &rb)
	// if err != nil {
	// 	tx.Rollback()
	// 	errs.HTTPErrorResponse(w, lgr, errs.Error{
	// 		StatusCode: 500,
	// 		Err:        err,
	// 	})
	// 	return
	// }

	// tx.Commit()

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
	// lgr := *hlog.FromRequest(r)

	// vars := mux.Vars(r)
	// rawCid := vars["id"]

	// ProductID, err := strconv.Atoi(rawCid)
	// if ProductID < 2 {

	// 	errs.HTTPErrorResponse(w, lgr, errs.Error{
	// 		StatusCode: 400,
	// 		Err:        errors.New("Bad Product id"),
	// 	})
	// 	return
	// }

	// if err != nil {

	// 	errs.HTTPErrorResponse(w, lgr, errs.Error{
	// 		StatusCode: 400,
	// 		Err:        err,
	// 	})
	// 	return
	// }

	// tx := s.db.MustBegin()
	// if err != nil {
	// 	errs.HTTPErrorResponse(w, lgr, errs.Error{
	// 		StatusCode: 500,
	// 		Err:        err,
	// 	})
	// 	return
	// }

	// err = s.service.Delete(r.Context(), tx, ProductID)
	// if err != nil {
	// 	tx.Rollback()
	// 	errs.HTTPErrorResponse(w, lgr, errs.Error{
	// 		StatusCode: 500,
	// 		Err:        err,
	// 	})
	// 	return
	// }

	// tx.Commit()
	json.NewEncoder(w).Encode("success")
}
