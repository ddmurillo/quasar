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
	router.HandleFunc("/topsecret", ctr.Location).Methods("POST")
	router.HandleFunc("/topsecret_split", ctr.GetLocation).Methods("GET")
	router.HandleFunc("/topsecret_split/{satellite_name}", ctr.SaveLocation).Methods("POST")

	// "escucha" por el puerto 8080
	log.Fatal(http.ListenAndServe(":8080", router))

}
