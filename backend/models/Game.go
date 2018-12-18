package models

type Game struct {
	Name string
	Password string
	CurrentState GameState
}
