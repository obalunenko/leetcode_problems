package judgecircle

import (
	"strings"
)

type robot struct {
	position position
}

type position struct {
	x int
	y int
}

func (r *robot) moveRight() {
	r.position.y++
}

func (r *robot) moveLeft() {
	r.position.y--
}

func (r *robot) moveUp() {
	r.position.x++
}

func (r *robot) moveDown() {
	r.position.x--
}

func (r *robot) move(moves string) {
	for _, step := range moves {
		switch strings.ToUpper(string(step)) {
		case "R":
			r.moveRight()
		case "L":
			r.moveLeft()
		case "U":
			r.moveUp()
		case "D":
			r.moveDown()
		}
	}
}

func judgeCircle(moves string) bool {
	r := robot{
		position: position{
			x: 0,
			y: 0,
		},
	}

	r.move(moves)

	if r.position.x == 0 && r.position.y == 0 {
		return true
	}

	return false
}
