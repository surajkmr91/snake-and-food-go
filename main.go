package main

import "github.com/surajkmr91/snake-and-food-go/game"

func main() {

	gameEngine := game.NewGame(30)
	go gameEngine.DirectionCommand()
	gameEngine.Tick()

}
