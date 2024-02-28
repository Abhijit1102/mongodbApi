package main

import (
	"fmt"
	"github.com/gorilla/mux" // Import the mux package
)

func Router() *mux.Router {
	router := mux.NewRouter()
	
	router.HandleFunc("/api/movies", controller.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
    router.HandleFunc("/api/movie/{id}", controller.MarksAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteAMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovie", controller.DeleteAllMovies).Methods("DELETE")

	return router
}
