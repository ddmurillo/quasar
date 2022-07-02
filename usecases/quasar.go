package usecases

import (
	"fmt"
	"strings"

	ent "github.com/ddmurillo/quasar/entities"
)

// Configuracion de los satellites en uso
var Kenobi = [2]float64{-500, -200}
var Skywalker = [2]float64{100, -100}
var Sato = [2]float64{500, 100}

// Obtiene la posicion de la nave de acuerdo a las distancias con los satelites. El metodo se llama triletaracion (https://es.wikipedia.org/wiki/Trilateraci%C3%B3n)
func GetLocation(distances ...float32) (float32, float32) {

	// crear entradas para enviar a la funcion de trilateracion
	var circle1 = ent.Circle{X: Kenobi[0], Y: Kenobi[1], R: float64(distances[0])}
	var circle2 = ent.Circle{X: Skywalker[0], Y: Skywalker[1], R: float64(distances[1])}
	var circle3 = ent.Circle{X: Sato[0], Y: Sato[1], R: float64(distances[2])}

	// calcula la posicion con el algoritmo de triletaracion
	var cx, cy = CalculatePosition(circle1, circle2, circle3)

	return float32(cx), float32(cy)
}

// Obtiene el mensaje completo de acuerdo a los mensajes que tiene cada satelite
func GetMessage(messages ...[]string) string {
	var maxlen int = 0
	// determina el array mas grande ya que no todos los mensajes son del mismo tamaño (debido a los desfases)
	for i := 0; i < len(messages); i++ {
		if maxlen < len(messages[i]) {
			maxlen = len(messages[i])
		}
	}

	fmt.Println(maxlen)
	var resultArray = make([]string, maxlen)

	// agrega los items a un array que "concatena" todos los mensajes de los satelites
	for i := 0; i < len(messages); i++ {
		for j := 0; j < len(messages[i]); j++ {
			// descarta mensajes vacios y posiciones que ya contienen el mismo texto
			if messages[i][j] != "" && (resultArray[j] == "" || resultArray[j] == messages[i][j]) {
				resultArray[j] = messages[i][j]
			}
		}
	}

	fmt.Println(resultArray)

	// resultado del mensaje
	var result string = ""
	previousWord := ""
	// concatena los items del array que tiene todos los mensajes en un string
	for i := 0; i < len(resultArray); i++ {
		// para evitar palabras repetidas por el desfase, las compara y la omite
		if resultArray[i] != "" && previousWord != resultArray[i] {
			result += resultArray[i] + " "
			previousWord = resultArray[i]
		}
	}

	return strings.Trim(result, " ")
}
