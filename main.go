package main

import (
	"fmt"
	"time"
)

type GliderGame struct {
	GameGrid [][]int
}

// New Game will init gris with glider pattern given in problem
func NewGame(x, y int) *GliderGame {
	g := &GliderGame{}
	g.GameGrid = make([][]int, x)
	for val := range g.GameGrid {
		g.GameGrid[val] = make([]int, y)
	}
	g.GameGrid[5][6] = 1
	g.GameGrid[6][7] = 1
	g.GameGrid[7][5] = 1
	g.GameGrid[7][6] = 1
	g.GameGrid[7][7] = 1
	print(g.GameGrid)
	return g
}

func main() {
	fmt.Println("initalizing Glider pattern in 15 x 15 grid")
	g := NewGame(15, 15)
	for {
		r := g.nextPattern()
		print(r)
		g = &GliderGame{GameGrid: r}
		time.Sleep(1 * time.Second)
	}

}

func print(g [][]int) {
	fmt.Println("______________________________________________________")
	for _, v := range g {
		for _, val := range v {
			fmt.Print(val)
			fmt.Print("|")
		}
		fmt.Println("")
	}
	fmt.Println("______________________________________________________")
}

func (g *GliderGame) nextPattern() [][]int {
	//1. Any live cell with fewer than two live neighbors dies as if caused by underpopulation.
	//2. Any live cell with two or three live neighbors lives on to the next generation.
	//3. Any live cell with more than three live neighbors dies, as if by overcrowding.
	//4. Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
	newGrid := make([][]int, len(g.GameGrid))
	for x, rows := range g.GameGrid {
		newGrid[x] = make([]int, len(rows))
		for y := range rows {
			n := neighborsCount(g, x, y)
			if g.GameGrid[x][y] == 1 && (n == 2 || n == 3) {
				newGrid[x][y] = g.GameGrid[x][y]
			} else if g.GameGrid[x][y] == 0 && n == 3 {
				newGrid[x][y] = 1
			}
		}
	}
	return newGrid
}

// neighborsCount provide number of alive cells for given x,y coordinates
func neighborsCount(g *GliderGame, x int, y int) int {
	count := 0
	rows := len(g.GameGrid)
	cols := len(g.GameGrid[0])
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			r := x + i
			c := y + j
			if r >= 0 && r < rows && c >= 0 && c < cols && g.GameGrid[r][c] == 1 {
				count++
			}
		}
	}
	return count
}
