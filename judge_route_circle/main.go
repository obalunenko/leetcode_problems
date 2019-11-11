package main

import (
	"fmt"
	"strings"
)

type Robot struct {
	position Position
}

type Position struct {
	x int
	y int
}

func (r *Robot) moveRight() {
	r.position.y++
}

func (r *Robot) moveLeft() {
	r.position.y--
}

func (r *Robot) moveUp() {
	r.position.x++
}

func (r *Robot) moveDown() {
	r.position.x--
}

func (r *Robot) Move(moves string) {
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
		default:
			fmt.Printf("Wrong movement!")
		}
	}
}

func judgeCircle(moves string) bool {
	robot := Robot{
		position: Position{
			x: 0,
			y: 0,
		},
	}

	robot.Move(moves)

	if robot.position.x == 0 && robot.position.y == 0 {
		return true
	}

	return false
}
