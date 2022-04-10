package board

import "fmt"

type Board struct {
	grid [][]int
}

func (b *Board) Init() {
	b.grid = [][]int{
		make([]int, 5),
		make([]int, 5),
		make([]int, 5),
		make([]int, 5),
		make([]int, 5)}
}

func (b *Board) Move(x int, y int) {
	b.clear()

	y = translateYCoordinate(y)
	b.grid[y][x] = 1
}

func (b *Board) Grid() [][]int {
	return b.grid
}

func (b *Board) Print() {
	for i := 0; i < 5; i++ {
		fmt.Printf("%v\n", b.grid[i])
	}
	fmt.Printf("\n")
}

func (b *Board) clear() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			b.grid[i][j] = 0
		}
	}
}

func translateYCoordinate(y int) int {
	max_grid_index := 4
	return max_grid_index - y
}
