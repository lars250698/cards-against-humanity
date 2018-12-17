package main

import (
	"github.com/gorilla/mux"
	"github.com/lars250698/cards-against-humanity/backend/httpHandlers"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/get/games", httpHandlers.GetGames)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
