package main

import (
	"log"
	"net/http"

	ctr "github.com/ddmurillo/quasar/controllers"
	"github.com/gorilla/mux"
)

func main() {
	// configura rutas de la API
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/topsecret", ctr.Location).Methods("GET")
	// "escucha" por el puerto 8080
	log.Fatal(http.ListenAndServe(":8080", router))

}
