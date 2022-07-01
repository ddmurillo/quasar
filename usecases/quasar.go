package usecases

import (
	ent "github.com/ddmurillo/quasar/entities"
)

var Kenobi = [2]float64{-500, -200}
var Skywalker = [2]float64{100, -100}
var Sato = [2]float64{500, 100}

func GetLocation(distances ...float32) (float32, float32) {
	// crear entradas para enviar a la funcion de trilateracion
	var circle1 = ent.Circle{X: Kenobi[0], Y: Kenobi[1], R: float64(distances[0])}
	var circle2 = ent.Circle{X: Skywalker[0], Y: Skywalker[1], R: float64(distances[1])}
	var circle3 = ent.Circle{X: Sato[0], Y: Sato[1], R: float64(distances[2])}

	var cx, cy = CalculatePosition(circle1, circle2, circle3)

	return float32(cx), float32(cy)
}

func GetMessages(messages ...[]string) string {
	var msg string = ""
	var maxlen int = 0
	for i := 0; i < len(messages); i++ {
		if len(messages[i]) > maxlen {
			maxlen = len(messages[i])
		}
	}

	// maxima longitud del mensaje
	//var msg

	return msg
}
