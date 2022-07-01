package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	ent "github.com/ddmurillo/quasar/entities"
	uc "github.com/ddmurillo/quasar/usecases"
)

func Location(w http.ResponseWriter, r *http.Request) {
	var x, y = uc.GetLocation(100, 115.5, 142.7)
	fmt.Println(x, y)

	fmt.Println(uc.GetMessages([]string{"esto", "", "message"}, []string{"un", "", "message"}))

	var response = ent.InfoLocation{Pos: ent.Position{X: x, Y: y}, Message: ""}

	json.NewEncoder(w).Encode(response)
}
