package httpHandlers

import (
	"encoding/json"
	"github.com/lars250698/cards-against-humanity/backend/models"
	"github.com/lars250698/cards-against-humanity/backend/state"
	"log"
	"net/http"
)

type Body struct {
	Name string
	Password string
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	var body Body
	var game models.Game
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Println("Failed to parse request body: ", err)
		return
	}
	game.Name = body.Name
	game.Password = body.Password
	game.CurrentState = models.GameState{

	}
	state.NewGame(game)
}
