package controllers

import (
	"encoding/json"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/runntimeterror/scanngo-api/service"
)

type userController struct {
}

func NewUserController(db *sqlx.DB, userService service.IUserService) *userController {
	return &userController{
		// db,
		// userService,
	}
}

func (s *userController) RegisterRoutes(router *mux.Router) {
	userRoute := router.PathPrefix("/user").Subrouter()
	userRoute.HandleFunc("", s.handleMovieCreate).Methods("GET")
}

func (s *userController) handleMovieCreate(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"user": true})
}
