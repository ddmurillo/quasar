package usecases

import (
	"math"
)

//A function to apply trilateration formulas to return the (x,y) intersection point of three circles
func calcule_point(c1 ent.circle, c2 ent.circle, c3 ent.circle) (float32, float32) {
	var A float32 = 2*c2.x - 2*c1.x
	var B float32 = 2*c2.y - 2*c1.y
	var C float32 = float32(math.Pow(c1.r, 2)) - float32(math.Pow(c2.r, 2)) - float32(math.Pow(c1.x, 2)) + float32(math.Pow(c2.x, 2)) - float32(math.Pow(c1.y, 2)) + float32(math.Pow(c2.y, 2))
	var D float32 = 2*c3.x - 2*c2.x
	var E float32 = 2*c3.y - 2*c2.y
	var F float32 = float32(math.Pow(c2.r, 2)) - float32(math.Pow(c3.r, 2)) - float32(math.Pow(c2.x, 2)) + float32(math.Pow(c3.x, 2)) - float32(math.Pow(c2.y, 2)) + float32(math.Pow(c3.y, 2))
	var x float32 = (C*E - F*B) / (E*A - B*D)
	var y float32 = (C*D - A*F) / (B*D - A*E)
	return x, y
}
