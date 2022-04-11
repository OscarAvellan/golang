package commands_parser

import (
	"bufio"
	"os"
)

func Parse(file_name string) []string {
	file, _ := os.Open(file_name)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var commands []string

	for scanner.Scan() {
		commands = append(commands, scanner.Text())
	}

	return commands
}
