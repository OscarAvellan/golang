package main

import "example.com/go_tutorial/robot"
import "example.com/go_tutorial/board"
import "example.com/go_tutorial/robot_controller"

func main() {
	bot := robot.Robot{}
	board := board.Board{}
	board.Init()

	robot_controller.Place(0, 0, robot.East, &bot, &board)
	robot_controller.Report(&bot, &board)

	robot_controller.Move(&bot, &board)
	robot_controller.Report(&bot, &board)

	robot_controller.Left(&bot)
	robot_controller.Report(&bot, &board)

	robot_controller.Right(&bot)
	robot_controller.Report(&bot, &board)

	robot_controller.Right(&bot)
	robot_controller.Report(&bot, &board)

	robot_controller.Right(&bot)
	robot_controller.Report(&bot, &board)

	robot_controller.Right(&bot)
	robot_controller.Report(&bot, &board)

	robot_controller.Right(&bot)
	robot_controller.Report(&bot, &board)

	robot_controller.Move(&bot, &board)
	robot_controller.Report(&bot, &board)
}
