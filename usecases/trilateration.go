package usecases

import (
	"math"

	ent "github.com/ddmurillo/quasar/entities"
)

//A function to apply trilateration formulas to return the (x,y) intersection point of three circles
func CalculatePosition(c1 ent.Circle, c2 ent.Circle, c3 ent.Circle) (float64, float64) {
	var A float64 = 2*c2.X - 2*c1.X
	var B float64 = 2*c2.Y - 2*c1.Y
	var C float64 = math.Pow(c1.R, 2) - math.Pow(c2.R, 2) - math.Pow(c1.X, 2) + math.Pow(c2.X, 2) - math.Pow(c1.Y, 2) + math.Pow(c2.Y, 2)
	var D float64 = 2*c3.X - 2*c2.X
	var E float64 = 2*c3.Y - 2*c2.Y
	var F float64 = math.Pow(c2.R, 2) - math.Pow(c3.R, 2) - math.Pow(c2.X, 2) + math.Pow(c3.X, 2) - math.Pow(c2.Y, 2) + math.Pow(c3.Y, 2)
	var x float64 = (C*E - F*B) / (E*A - B*D)
	var y float64 = (C*D - A*F) / (B*D - A*E)
	return x, y
}
