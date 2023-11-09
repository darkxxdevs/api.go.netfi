package router

import (
	"api/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movies", controller.GetAllDBMovies).Methods("GET")
	router.HandleFunc("/api/movies/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movies/{id}", controller.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovies", controller.DeleteAllMovies).Methods("DELETE")

	return router
}
