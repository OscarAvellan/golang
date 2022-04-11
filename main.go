package main

import (
	"os"

	"example.com/go_tutorial/board"
	"example.com/go_tutorial/commands_parser"
	"example.com/go_tutorial/robot"
	"example.com/go_tutorial/robot_controller"
)

func main() {
	bot := robot.Robot{}
	board := board.Board{}
	board.Init()

	file_name := os.Args[1]
	commands := commands_parser.Parse(file_name)
	robot_controller.CommandsExecutor(commands, &bot, &board)
}
