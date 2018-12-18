package models

type GameState struct {
	Players []Player
	Round int
	Czar Player
	CurrentBlackCard BlackCard
	PlayedWhiteCards map[Player]WhiteCard
	WhiteCardStack []WhiteCard
}
