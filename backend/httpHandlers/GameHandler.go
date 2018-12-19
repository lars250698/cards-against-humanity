package httpHandlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/lars250698/cards-against-humanity/backend/models"
	"github.com/lars250698/cards-against-humanity/backend/state"
	"log"
	"net/http"
	"strconv"
)

var upgrader = websocket.Upgrader{}
var msg chan []byte
var gameID int
var playerID int

type PlayerUpdate struct {
	Action string
	CardID int
}

func GameHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	gameID, err = strconv.Atoi(mux.Vars(r)["gameID"])
	if err != nil {
		log.Println("Error parsing gameID: ", err)
		return
	}
	playerID, err = strconv.Atoi(r.URL.Query().Get("player"))
	if err != nil {
		log.Println("Error parsing playerID: ", err)
		return
	}
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade connection: ", err)
	}
	msg = make(chan []byte)
	go recv(c)
	go updateLoop(c, w)
}

func recv(c *websocket.Conn) {
	for {
		_, m, err := c.ReadMessage()
		if err != nil {
			log.Println("Failed to read client message: ", err)
			err := c.WriteMessage(websocket.TextMessage, []byte("err"))
			if err != nil {
				log.Println("Failed to send error message, giving up: ", err)
				return
			}
		}
		var upd PlayerUpdate
		err = json.Unmarshal(m, &upd)
		if err != nil {
			log.Println("Couldn't parse client message: ", err)
		}
		switch {
		case upd.Action == "played":
			game := <- state.Games[gameID]
			var playedCard models.WhiteCard
			// clean code 101
			playedCard, game.CurrentState.Players[playerID].Cards = game.CurrentState.Players[playerID].Cards[upd.CardID], append(game.CurrentState.Players[playerID].Cards[0:upd.CardID], game.CurrentState.Players[playerID].Cards[upd.CardID+1:]...)
			game.CurrentState.PlayedWhiteCards[game.CurrentState.Players[playerID]] = playedCard
			state.Games[gameID] <- game
			break
		case upd.Action == "voted":
			break
		}
	}
}

func updateLoop(c *websocket.Conn, w http.ResponseWriter) {
	for {
		select {
		case game := <- state.Games[gameID]:
			out, err := json.Marshal(game)
			if err != nil {
				log.Println("Failed to encode client message: ", err)
				break
			}
			_, err = fmt.Fprint(w, out)
			if err != nil {
				log.Println("Failed to send message: ", err)
			}
		}
	}
}
