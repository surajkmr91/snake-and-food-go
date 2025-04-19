package model

type Position struct {
	X int
	Y int
}

type Direction string

const (
	Up     Direction = "UP"
	Down   Direction = "Down"
	Left   Direction = "Left"
	Right  Direction = "Right"
	Unkown Direction = "Unknown"
)

func (d Direction) GetDirectionChangeValue(direction Direction, unit int) Position {
	switch direction {
	case Up:
		return Position{0, 1 * unit}
	case Down:
		return Position{0, -1 * unit}
	case Left:
		return Position{-1 * unit, 0}
	case Right:
		return Position{1 * unit, 0}

	}

	return Position{}
}

func GetDirection(inputDirection string) Direction {
	switch inputDirection {
	case "up":
		return Up
	case "down":
		return Down
	case "left":
		return Left
	case "right":
		return Right
	}
	return Unkown
}
