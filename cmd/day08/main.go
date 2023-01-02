package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	grid := readGrid("input.txt")
	v := findNumVisibleTrees(grid)
	fmt.Printf("%d visible trees out of %d\n", v, len(grid)*len(grid[0]))
}

func readGrid(filename string) [][]uint8 {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var grid [][]uint8
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var row []uint8
		for _, r := range line {
			i, _ := strconv.ParseInt(string(r), 10, 8)
			row = append(row, uint8(i))
		}
		grid = append(grid, row)
	}
	return grid
}

func findNumVisibleTrees(grid [][]uint8) int {
	sum := 0
	for r := range grid {
		for c := range grid[r] {
			if !isHidden(grid, r, c) {
				sum++
			}
		}
	}
	return sum
}

func isHidden(grid [][]uint8, r int, c int) bool {
	hn := isHiddenNorth(grid, r, c)
	hs := isHiddenSouth(grid, r, c)
	hw := isHiddenWest(grid, r, c)
	he := isHiddenEast(grid, r, c)
	return hn && hs && hw && he
}

func isHiddenNorth(grid [][]uint8, r int, c int) bool {
	height := grid[r][c]
	for i := r - 1; i >= 0; i-- {
		if grid[i][c] >= height {
			return true
		}
	}
	return false
}

func isHiddenSouth(grid [][]uint8, r int, c int) bool {
	height := grid[r][c]
	for i := r + 1; i < len(grid); i++ {
		if grid[i][c] >= height {
			return true
		}
	}
	return false
}

func isHiddenWest(grid [][]uint8, r int, c int) bool {
	height := grid[r][c]
	for i := c - 1; i >= 0; i-- {
		if grid[r][i] >= height {
			return true
		}
	}
	return false
}

func isHiddenEast(grid [][]uint8, r int, c int) bool {
	height := grid[r][c]
	for i := c + 1; i < len(grid[r]); i++ {
		if grid[r][i] >= height {
			return true
		}
	}
	return false
}
