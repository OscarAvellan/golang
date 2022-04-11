# Robot Simulation
Simple CLI tool to control a robot in 5x5 board.
The robot can be controlled with the following commands:
- PLACE(X, Y, Direction)
- MOVE
- LEFT
- RIGHT
- REPORT


## Contents
- [Dependencies](#dependencies)
- [Installing dependencies](#installing-dependencies)
- [Executing app](#executing-app)
- [Running tests](#running-tests)

<br>

## Dependencies
- Go 1.18

## Installing dependencies
- Installing `Go 1.18` -> [golang installation](https://go.dev/doc/install)

## Executing app
To run the app, from the project root directory, open a terminal and run:
```bash
$ go build main.go
$ ./main commands.txt
```

## Running tests
**NOTE: There are missing tests due to lack of time.**

To run the tests, execute the following command from the root directory
```bash
$ go test example.com/go_tutorial/robot_controller
```

If you want more verbose output
```bash
$ go test example.com/go_tutorial/robot_controller -v
```
