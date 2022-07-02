package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	ent "github.com/ddmurillo/quasar/entities"
	repo "github.com/ddmurillo/quasar/repository"
	uc "github.com/ddmurillo/quasar/usecases"
	"github.com/gorilla/mux"
)

// determina la posicion de acuerdo a la informacion de los satellites definida y las distancias y mensajes que se recibe en el body
func Location(w http.ResponseWriter, r *http.Request) {

	// lee body
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil || err == io.EOF {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// deserrializa el mensaje en un objeto de request
	var satRequest ent.RequestLocation
	json.Unmarshal(reqBody, &satRequest)

	// valida que los datos
	if satRequest.Satellites == nil || len(satRequest.Satellites) != 3 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// array con las distancias obtenidas en el request de la peticion
	var distances []float32
	// array de mensajes obtenidos en el request de la peticion
	var messages [][]string

	// recorre el request y llena los respectivos arrays de distancias y mensajes
	for i := 0; i < len(satRequest.Satellites); i++ {
		distances = append(distances, satRequest.Satellites[i].Distance)
		messages = append(messages, satRequest.Satellites[i].Message)
	}

	// invoca la funcion que calcula la distancia de acuerdo a las posiciones de los satelites
	var x, y = uc.GetLocation(distances...)
	fmt.Println(x, y)

	// invoca la funcion que obtiene el mensaje de acuerdo a lo recibido en cada satelite
	var msg string = uc.GetMessage(messages...)
	fmt.Println(msg)

	// arma objeto de respuesta
	var response = ent.InfoLocation{Pos: ent.Position{X: x, Y: y}, Message: msg}

	// responde en formato json
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// determina la posicion de acuerdo a la informacion de los satellites definida y las distancias y mensajes que se encuentran guardadas en el repositorio
func GetLocation(w http.ResponseWriter, r *http.Request) {

	// obtiene los datos guardados en el repositorio
	var listSatelliteData []ent.InfoSatellite = repo.Get()
	fmt.Println(listSatelliteData)

	// valida que este la informacion de los 3 satelites en el repositorio
	if len(listSatelliteData) != 3 {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ent.Response{Code: 400, Message: "Información insuficiente para determinar la posición"})
		return
	}

	// array con las distancias obtenidas en el request de la peticion
	var distances []float32
	// array de mensajes obtenidos en el request de la peticion
	var messages [][]string

	// recorre el request y llena los respectivos arrays de distancias y mensajes
	for i := 0; i < len(listSatelliteData); i++ {
		distances = append(distances, listSatelliteData[i].Distance)
		messages = append(messages, listSatelliteData[i].Message)
	}

	// invoca la funcion que calcula la distancia de acuerdo a las posiciones de los satelites
	var x, y = uc.GetLocation(distances...)
	fmt.Println(x, y)

	// invoca la funcion que obtiene el mensaje de acuerdo a lo recibido en cada satelite
	var msg string = uc.GetMessage(messages...)
	fmt.Println(msg)

	// entidad de respuesta
	var response = ent.InfoLocation{Pos: ent.Position{X: x, Y: y}, Message: msg}

	// responde en formato json
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// guarda informacion de un satelite en el repositorio
func SaveLocation(w http.ResponseWriter, r *http.Request) {
	// lee el body del request
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil || err == io.EOF {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// desarializa en el objeto de info del satelite
	var satRequest ent.InfoSatellite
	json.Unmarshal(reqBody, &satRequest)

	// el nombre del satelite viene como parametro en la URL, lo lee
	params := mux.Vars(r)
	fmt.Println(params)

	// Valida que el paremetro sea valido
	if params == nil || params["satellite_name"] == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	satelliteName := params["satellite_name"]

	// entidad para guardar
	satelliteToSave := ent.InfoSatellite{Name: satelliteName, Distance: satRequest.Distance, Message: satRequest.Message}

	fmt.Println(satelliteToSave)
	// guarda
	if !repo.Save(satelliteToSave) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// respuesta en formato json
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ent.Response{Code: 200, Message: "Info guardada"})
}
