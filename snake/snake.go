package snake

import (
	"fmt"

	model "github.com/surajkmr91/snake-and-food-go/model"
)

type Snake struct {
	body          []model.Position // Head at index 0, tail at end
	occupied      map[model.Position]bool
	movementunit  int
	currDirection model.Direction
}

func NewSnake(startPoint model.Position) *Snake {
	occupiedMap := make(map[model.Position]bool)
	occupiedMap[startPoint] = true
	return &Snake{
		body:         []model.Position{startPoint},
		occupied:     occupiedMap,
		movementunit: 1,
	}
}

func (s *Snake) Move(position model.Position, isGrow bool) model.Position {
	s.body = append([]model.Position{position}, s.body...)
	s.occupied[position] = true

	if !isGrow {
		tail := s.body[len(s.body)-1]
		s.body = s.body[:len(s.body)-1]
		delete(s.occupied, tail)
	}
	return position
}

func (s *Snake) ChangeAndSetDirection(dir model.Direction) {
	if s.currDirection == model.Up && dir == model.Down {
		fmt.Println("Invalid direction.")
	} else if s.currDirection == model.Down && dir == model.Up {
		fmt.Println("Invalid direction. ")
	} else if s.currDirection == model.Left && dir == model.Right {
		fmt.Println("Invalid direction. ")
	} else if s.currDirection == model.Right && dir == model.Left {
		fmt.Println("Invalid direction. ")
	}
	s.currDirection = dir
}

func (s *Snake) CheckCollisionAndGetNextMove(dir model.Direction) (model.Position, bool) {
	nextPositionDelta := s.currDirection.GetDirectionChangeValue(s.currDirection, 1)
	position := model.Position{
		X: s.body[0].X + nextPositionDelta.X,
		Y: s.body[0].Y + nextPositionDelta.Y,
	}

	if s.occupied[position] {
		return position, true
	}

	return position, false
}
