package router

import (
	"github.com/gorilla/mux"
	"github.com/sahilchauhan0603/society_management_backend/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/home",controller.ServeHome).Methods("GET")

	return router
}
