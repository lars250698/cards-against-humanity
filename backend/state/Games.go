package state

import "github.com/lars250698/cards-against-humanity/backend/models"

var Games []chan models.Game

func NewGame(game models.Game) {
	g := make(chan models.Game)
	g <- game
	Games = append(Games, g)
}
