package main

import (
	"github.com/gorilla/mux"
	"github.com/lars250698/cards-against-humanity/backend/httpHandlers"
	"github.com/lars250698/cards-against-humanity/backend/models"
	"github.com/lars250698/cards-against-humanity/backend/state"
	"net/http"
)

func main() {
	state.Games = make(chan []models.Game)
	r := mux.NewRouter()

	r.HandleFunc("/get/games", httpHandlers.GetGames)
	r.HandleFunc("/game/{gameID:[0-9]+}", httpHandlers.GameHandler)
	r.HandleFunc("/game/new", httpHandlers.NewGame)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
