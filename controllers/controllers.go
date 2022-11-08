package controllers

type IController interface {
	GetRoutes() []Route
	GetPath() string
}
