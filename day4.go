package main

import (
	"fmt"
	"strings"
)

type Directions struct {
	up_down_row    int
	left_right_col int
}

func day4_1(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	grid := []string{}

	//lines... iterates over each line and adds to the slice
	grid = append(grid, lines...)

	var (
		rows   int = len(grid)
		cols   int = len(grid[0])
		output int
	)

	konamiCode := []Directions{
		{0, 1},   //R
		{0, -1},  //L
		{1, 0},   //D
		{-1, 0},  //U
		{1, 1},   //Diag DR
		{1, -1},  //Diag DL
		{-1, 1},  //Diag UR
		{-1, -1}, //Diag UL
	}

	/* Matrix directions to maintain sanity
	   (-1, -1)   (-1,  0)   (-1, +1)
	   ( 0, -1)   ( 0,  0)   ( 0, +1)
	   (+1, -1)   (+1,  0)   (+1, +1)

	   - up_down_row = row offset (negative = move up, positive = move down, 0 = stay put)
	   - left_right_col = column offset (negative = move left, positive = move right, 0 = stay put)
	*/

	var (
		target     string = "XMAS"
		wordLength int    = 4
	)

	//r & c are row and col indexes, the starting point of the grid
	//dr, dc are how we move from that starting position (the dir offsets)
	checkDir := func(r, c, dr, dc int) bool {

		//Loop through each char in our target word "XMAS" (wordLength = 4)
		for i := 0; i < wordLength; i++ {

			//Each step move further in one direction
			newRow := r + dr*i
			newCol := c + dc*i

			//Check boundaries, if nr or nc fall outside the grid, we can't continue matching
			if newRow < 0 || newRow >= rows || newCol < 0 || newCol >= cols {
				return false
			}

			// if the char in the grid is not the 'i'th char in "XMAS" we also cannot continue
			if grid[newRow][newCol] != target[i] {
				return false
			}

		}
		//If false is never returned
		return true
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == target[0] {
				for _, d := range konamiCode {
					if checkDir(r, c, d.up_down_row, d.left_right_col) {
						output++
					}
				}
			}
		}
	}

	fmt.Println("Solution One:", output)
}

func day4_2(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	// 1) Build the grid
	var grid []string
	grid = append(grid, lines...)

	// 2) Now we can safely define rows & cols
	rows := len(grid)
	if rows == 0 {
		fmt.Println("Grid is empty!")
		return
	}
	cols := len(grid[0])

	// 3) Output variable
	output := 0

	for r := 1; r < rows-1; r++ {
		for c := 1; c < cols-1; c++ {
			if grid[r][c] == 'A' {
				uL := string(grid[r-1][c-1]) // up-left
				uR := string(grid[r-1][c+1]) // up-right
				dL := string(grid[r+1][c-1]) // down-left
				dR := string(grid[r+1][c+1]) // down-right

				//Around the horn clockwise will allow the checker to pick up any possible match that results in an 'X' shaped "MAS"
				cornerCheck := strings.Join([]string{uL, uR, dR, dL}, "")

				if cornerCheck != "MMSS" && cornerCheck != "MSSM" && cornerCheck != "SSMM" && cornerCheck != "SMMS" {
					fmt.Println("fail")
				} else {
					fmt.Println("succeed")
					output++
				}
			}
		}
	}

	fmt.Println("Solution Two:", output)
}
