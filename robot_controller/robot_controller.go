package robot_controller

import "example.com/go_tutorial/robot"
import "example.com/go_tutorial/board"

func Place(x int, y int, direction robot.Direction, r *robot.Robot, b *board.Board) {
	if canMove(x, y) {
		r.SetPosition(x, y)
		r.SetDirection(direction)
		b.Move(x, y)
	}
}

func Move(r *robot.Robot, b *board.Board) {
	new_x_position := r.XCoordinate()
	new_y_position := r.YCoordinate()

	switch r.Direction() {
	case robot.North:
		new_y_position += 1
	case robot.South:
		new_y_position -= 1
	case robot.West:
		new_x_position -= 1
	case robot.East:
		new_x_position += 1
	}

	if canMove(new_x_position, new_y_position) {
		r.SetPosition(new_x_position, new_y_position)
		b.Move(new_x_position, new_y_position)
	}
}

func Left(r *robot.Robot) {
	r.RotateLeft()
}

func Right(r *robot.Robot) {
	r.RotateRight()
}

func Report(r *robot.Robot, b *board.Board) {
	r.Print()
	b.Print()
}

func canMove(x, y int) bool {
	if x < 0 || x > 4 {
		return false
	}

	if y < 0 || y > 4 {
		return false
	}

	return true
}
