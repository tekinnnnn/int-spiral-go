package main

import (
	"fmt"
)

type Directions struct {
	North, South, East, West byte
}

type Position struct {
	X, Y int
}

const (
	x int = 5
	y int = 5
)

var (
	directions Directions = Directions{'N', 'S', 'E', 'W'}
	direction  byte       = directions.East
	pos        Position   = Position{0, 0}
	matrix     [y][x]uint
)

func main() {
	for i := 1; i < x*y; i++ {
		switch direction {
		case directions.East:
			pos.X++
			if pos.X == x || matrix[pos.Y][pos.X] != 0 {
				pos.X--
				i--
				rotate()
				continue
			}
		case directions.South:
			pos.Y++
			if pos.Y == y || matrix[pos.Y][pos.X] != 0 {
				pos.Y--
				i--
				rotate()
				continue
			}
		case directions.West:
			pos.X--
			if pos.X < 0 || matrix[pos.Y][pos.X] != 0 {
				pos.X++
				i--
				rotate()
				continue
			}
		case directions.North:
			pos.Y--
			if pos.Y < 0 || (pos.X == 0 && pos.Y == 0) || matrix[pos.Y][pos.X] != 0 {
				pos.Y++
				i--
				rotate()
				continue
			}
		}

		fmt.Println("i:", i, "pos:", pos, "direction:", string(direction))

		matrix[pos.Y][pos.X] = uint(i)
	}

	fmt.Println(matrix)
}

func rotate() {
	switch direction {
	case directions.North:
		direction = directions.East
	case directions.East:
		direction = directions.South
	case directions.South:
		direction = directions.West
	case directions.West:
		direction = directions.North
	}
}
