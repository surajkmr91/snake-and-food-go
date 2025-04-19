package game

import model "github.com/surajkmr91/snake-and-food-go/model"

type Board struct {
	Size int
	Rows int
	Cols int
}

func NewBoard(rows, cols int) *Board {
	return &Board{
		Rows: rows,
		Cols: cols,
	}
}

func (b *Board) IsInBound(p model.Position) bool {
	// check if the position is within the bounds of the board
	return p.X >= 0 && p.X < b.Rows &&
		p.Y >= 0 && p.Y < b.Cols
}

// 	// map to store start and end positions of snakes and ladders
// 	// so that they dont overlap
// 	snakeLadderMap := make(map[string]bool)

// 	for i := 0; i < int(noOfSnakes); i++ {
// 		for {
// 			start := rand.Int63n(b.size) + 1
// 			end := rand.Int63n(b.size) + 1
// 			if end >= start || start == size {
// 				continue
// 			}
// 			if _, ok := snakeLadderMap[fmt.Sprintf("%d:%d", start, end)]; !ok {
// 				b.snakes = append(b.snakes, InitSnake(start, end))
// 				snakeLadderMap[fmt.Sprintf("%d:%d", start, end)] = true
// 				break
// 			}
// 		}
// 	}

// 	for i := 0; i < int(noOfLadders); i++ {
// 		for {
// 			start := rand.Int63n(b.size) + 1
// 			end := rand.Int63n(b.size) + 1
// 			if start >= end || start == 1 {
// 				continue
// 			}
// 			if _, ok := snakeLadderMap[fmt.Sprintf("%d:%d", start, end)]; !ok {
// 				b.ladder = append(b.ladder, InitLadder(start, end))
// 				snakeLadderMap[fmt.Sprintf("%d:%d", start, end)] = true
// 				break
// 			}
// 		}
// 	}
// 	return b
// }

// func (b *Board) isLadder(pos int64) (bool, int64) {
// 	for _, val := range b.ladder {
// 		if val.startPoint == pos {
// 			return true, val.endPoint
// 		}
// 	}
// 	return false, -1
// }

// func (b *Board) isSnake(pos int64) (bool, int64) {
// 	for _, val := range b.snakes {
// 		if val.startPoint == pos {
// 			return true, val.endPoint
// 		}
// 	}
// 	return false, -1
// }

// func (b *Board) GetEndValue() int64 {
// 	return b.size
// }

// func (b *Board) GetNewPosition(pos int64) int64 {
// 	if ok, val := b.isLadder(pos); ok {
// 		fmt.Printf("is ladder. start: %d, end: %d\n", pos, val)
// 		return val
// 	}

// 	if ok, val := b.isSnake(pos); ok {
// 		fmt.Printf("is snake. start: %d, end: %d\n", pos, val)
// 		return val
// 	}

// 	return pos
// }
