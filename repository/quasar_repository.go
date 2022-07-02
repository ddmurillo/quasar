package repository

import (
	"fmt"
	"strings"

	ent "github.com/ddmurillo/quasar/entities"
)

// arreglo que guarda la informacion recibida de los satelies
var data []ent.InfoSatellite

// funcion que guarda la informacion en el arreglo
func Save(dataToSave ent.InfoSatellite) bool {
	fmt.Println(strings.ToLower(dataToSave.Name))

	// aplica solo para los tres satelites configurados
	if !(strings.ToLower(dataToSave.Name) == "sato" ||
		strings.ToLower(dataToSave.Name) == "skywalker" ||
		strings.ToLower(dataToSave.Name) == "kenobi") {
		return false
	}

	// variable de contral para determinar si el satelite ya esta guardado en el array
	fouded := false

	// reccore la info del array para actualizar la info del staelite si existe
	for i := 0; i < len(data); i++ {
		if data[i].Name == dataToSave.Name {
			data[i] = dataToSave
			fmt.Println(data)
			fouded = true
			break
		}
	}

	fmt.Print(fouded)
	// si el satelite aun no esta en el array lo agrega
	if !fouded {
		data = append(data, dataToSave)
	}

	return true
}

// Retorna la info guardada en el array
func Get() []ent.InfoSatellite {
	return data
}
