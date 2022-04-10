package robot

import "fmt"

type Direction int

const (
	North Direction = iota
	West  Direction = iota
	South Direction = iota
	East  Direction = iota
)

type Robot struct {
	direction  Direction
	x_position int
	y_position int
}

func (r *Robot) SetDirection(new_direction Direction) {
	r.direction = new_direction
}

func (r *Robot) SetPosition(x int, y int) {
	r.x_position = x
	r.y_position = y
}

func (r *Robot) RotateLeft() {
	if r.direction == East {
		r.direction = North
		return
	}

	r.direction += 1
}

func (r *Robot) RotateRight() {
	if r.direction == North {
		r.direction = East
		return
	}

	r.direction -= 1
}

func (r *Robot) Print() {
	fmt.Printf("x: %v, y: %v, direction: %v\n", r.x_position, r.y_position, r.direction)
}

func (r *Robot) XCoordinate() int {
	return r.x_position
}

func (r *Robot) YCoordinate() int {
	return r.y_position
}

func (r *Robot) Direction() Direction {
	return r.direction
}
