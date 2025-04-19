package game

import (
	"log"

	model "github.com/surajkmr91/snake-and-food-go/model"
)

// import "github.com/surajkmr91/snake-and-food-go/model"

type FoodManager struct {
	Position []model.Position
	Index    int
}

func NewFoodmanager(position []model.Position) *FoodManager {
	return &FoodManager{
		Position: position,
		Index:    0,
	}
}

func (f *FoodManager) GetNextFood(snakePosition map[model.Position]bool) *model.Position {
	if f.Index >= len(f.Position) {
		return nil
	}

	for i := f.Index; i < len(f.Position); i++ {
		_, exists := snakePosition[f.Position[i]]
		if exists {
			log.Printf("Invalid Food position - %+v, since its on the snake body ", f.Position[i])
			continue
		}
		f.Index = i
		return &f.Position[i]
	}
	return nil
}

func (f *FoodManager) IsFoodAtCurrentPlace(currPos model.Position) bool {
	if f.Position[f.Index].X == currPos.X && f.Position[f.Index].Y == currPos.Y {
		return true
	}
	return false
}

func (f *FoodManager) ConsumeFood(position model.Position) {
	if f.Index >= len(f.Position) {
		return
	}
	for i := f.Index; i < len(f.Position); i++ {
		if f.Position[i].X == position.X && f.Position[i].Y == position.Y {
			f.Index = i + 1
			return
		}
	}
}
