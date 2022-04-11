package robot_controller

import (
	"strconv"
	"strings"

	"example.com/go_tutorial/board"
	"example.com/go_tutorial/robot"
)

func CommandsExecutor(commands []string, r *robot.Robot, b *board.Board) {
	for _, line := range commands {
		executeCommand(line, r, b)
	}
}

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

func executeCommand(command string, r *robot.Robot, b *board.Board) {
	command_array := strings.Split(command, " ")

	switch command_array[0] {
	case "PLACE":
		x, _ := strconv.Atoi(command_array[1])
		y, _ := strconv.Atoi(command_array[2])
		direction := getDirection(command_array[3])
		Place(x, y, direction, r, b)
	case "MOVE":
		Move(r, b)
	case "LEFT":
		Left(r)
	case "RIGHT":
		Right(r)
	case "REPORT":
		Report(r, b)
	}
}

func getDirection(direction string) robot.Direction {
	switch direction {
	case "NORTH":
		return robot.North
	case "WEST":
		return robot.West
	case "SOUTH":
		return robot.South
	case "EAST":
		return robot.East
	}

	return robot.North
}
