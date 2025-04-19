package game

import (
	"fmt"
	"math/rand"
	"time"

	//"github.com/surajkmr91/snake-and-food-go/snake"
	model "github.com/surajkmr91/snake-and-food-go/model"
	snake "github.com/surajkmr91/snake-and-food-go/snake"
)

type Game struct {
	Snake       *snake.Snake
	Board       *Board
	FoodManager *FoodManager
	Over        bool
}

func NewGame(boardSize int) *Game {
	board := NewBoard(boardSize, boardSize)
	snakePosition := model.Position{X: 0, Y: 0}
	snake := snake.NewSnake(snakePosition)
	var foodPositions []model.Position
	source := rand.NewSource(time.Now().UnixNano())
	// Create a new random number generator using the source
	r := rand.New(source)
	for i := 0; i < 10; i++ {
		x := r.Intn(boardSize)
		y := r.Intn(boardSize)
		foodPositions = append(foodPositions, model.Position{X: x, Y: y})
	}
	foodManager := NewFoodmanager(foodPositions)
	return &Game{
		Snake:       snake,
		Board:       board,
		FoodManager: foodManager,
		Over:        false,
	}
}

func (game *Game) Tick() {
	var initialDirection string
	fmt.Scanln(&initialDirection)

	direction := model.GetDirection(initialDirection)
	game.Snake.ChangeAndSetDirection(direction)

	if direction == model.Unkown {
		fmt.Println("Invalid direction. Please enter a valid direction.")
		var newDirection string
		fmt.Scanln(&newDirection)
		direction := model.GetDirection(newDirection)
		game.Snake.ChangeAndSetDirection(direction)
	} else {
		for {
			fmt.Println("current food position ", game.FoodManager.Position[game.FoodManager.Index])
			newHead, iscollision := game.Snake.CheckCollisionAndGetNextMove(direction)
			if iscollision || !game.Board.IsInBound(newHead) {
				game.Over = true
				fmt.Println("game over, as new head position is on body/border ", newHead)
				return
			}
			fmt.Println(" current head position ", newHead)

			toGrow := false
			if game.FoodManager.IsFoodAtCurrentPlace(newHead) {
				toGrow = true
				game.FoodManager.ConsumeFood(newHead)
			}
			game.Snake.Move(newHead, toGrow)

			time.Sleep(500 * time.Millisecond)
		}
	}

}

func (game *Game) DirectionCommand() {
	for {
		fmt.Println("is any new direction added")
		var newDirection string
		fmt.Scanln(&newDirection)
		direction := model.GetDirection(newDirection)
		game.Snake.ChangeAndSetDirection(direction)
		if game.Over {
			fmt.Println("Game Over!")
			break
		}
	}
}
