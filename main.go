package main

import (
	"fmt"

	uc "github.com/ddmurillo/quasar/usecases"
)

func main() {

	fmt.Println(uc.GetLocation(100, 115.5, 142.7))
}
