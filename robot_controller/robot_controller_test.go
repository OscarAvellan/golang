package robot_controller

import "testing"
import "example.com/go_tutorial/robot"
import "example.com/go_tutorial/board"

func TestPlace(t *testing.T) {
	bot := robot.Robot{}
	board := board.Board{}
	board.Init()

	Place(4, 4, robot.South, &bot, &board)

	if bot.Direction() != robot.South {
		t.Errorf("Expected %v, got %v", robot.South, bot.Direction())
	}

	if bot.XCoordinate() != 4 {
		t.Errorf("Expected %v, got %v", 4, bot.XCoordinate())
	}

	if bot.YCoordinate() != 4 {
		t.Errorf("Expected %v, got %v", 4, bot.YCoordinate())
	}

	if board.Grid()[0][4] != 1 {
		t.Errorf("Expected %v, got %v", 1, board.Grid()[0][4])
	}
}

func TestPlaceInvalidPosition(t *testing.T) {
	bot := robot.Robot{}
	board := board.Board{}
	board.Init()

	Place(5, 4, robot.South, &bot, &board)

	if bot.Direction() != robot.North {
		t.Errorf("Expected %v, got %v", robot.North, bot.Direction())
	}

	if bot.XCoordinate() != 0 {
		t.Errorf("Expected %v, got %v", 0, bot.XCoordinate())
	}

	if bot.YCoordinate() != 0 {
		t.Errorf("Expected %v, got %v", 0, bot.YCoordinate())
	}

	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if board.Grid()[y][x] != 0 {
				t.Errorf("Expected grid[%v][%v] = %v, got %v", y, x, 0, bot.YCoordinate())
			}
		}
	}
}

func TestMove(t *testing.T) {
	bot := robot.Robot{}
	board := board.Board{}
	board.Init()

	Place(0, 0, robot.East, &bot, &board)
	Move(&bot, &board)

	if bot.Direction() != robot.East {
		t.Errorf("Expected direction %v, got %v", robot.East, bot.Direction())
	}

	if bot.XCoordinate() != 1 {
		t.Errorf("Expected x coordinate %v, got %v", 1, bot.XCoordinate())
	}

	if bot.YCoordinate() != 0 {
		t.Errorf("Expected y coordinate %v, got %v", 0, bot.YCoordinate())
	}

	if board.Grid()[4][1] != 1 {
		t.Errorf("Expected grid %v, got %v", 1, board.Grid()[4][1])
	}
}

func TestMoveInvalid(t *testing.T) {
	bot := robot.Robot{}
	board := board.Board{}
	board.Init()

	Place(0, 0, robot.South, &bot, &board)
	Move(&bot, &board)

	if bot.Direction() != robot.South {
		t.Errorf("Expected %v, got %v", robot.South, bot.Direction())
	}

	if bot.XCoordinate() != 0 {
		t.Errorf("Expected %v, got %v", 0, bot.XCoordinate())
	}

	if bot.YCoordinate() != 0 {
		t.Errorf("Expected %v, got %v", 0, bot.YCoordinate())
	}

	if board.Grid()[4][0] != 1 {
		t.Errorf("Expected grid %v, got %v", 1, board.Grid()[0][0])
	}
}

func TestLeft(t *testing.T) {
	bot := robot.Robot{}
	board := board.Board{}
	board.Init()

	Place(2, 3, robot.East, &bot, &board)
	Left(&bot)

	if bot.Direction() != robot.North {
		t.Errorf("Expected direction %v, got %v", robot.North, bot.Direction())
	}

	if bot.XCoordinate() != 2 {
		t.Errorf("Expected x coordinate %v, got %v", 2, bot.XCoordinate())
	}

	if bot.YCoordinate() != 3 {
		t.Errorf("Expected y coordinate %v, got %v", 3, bot.YCoordinate())
	}

	if board.Grid()[1][2] != 1 {
		t.Errorf("Expected grid %v, got %v", 1, board.Grid()[1][2])
	}
}

func TestRight(t *testing.T) {
	bot := robot.Robot{}
	board := board.Board{}
	board.Init()

	Place(2, 3, robot.East, &bot, &board)
	Right(&bot)

	if bot.Direction() != robot.South {
		t.Errorf("Expected direction %v, got %v", robot.South, bot.Direction())
	}

	if bot.XCoordinate() != 2 {
		t.Errorf("Expected x coordinate %v, got %v", 2, bot.XCoordinate())
	}

	if bot.YCoordinate() != 3 {
		t.Errorf("Expected y coordinate %v, got %v", 3, bot.YCoordinate())
	}

	if board.Grid()[1][2] != 1 {
		t.Errorf("Expected grid %v, got %v", 1, board.Grid()[1][2])
	}
}
