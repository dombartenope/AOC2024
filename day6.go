package main

import (
	"fmt"
	"strings"
)

func day6_1(input string) {
	//Same as above for basic guard movement
	lines := strings.Split(strings.TrimSpace(input), "\n")
	//Create a 2D slice of runes, one per line of input
	grid := make([][]rune, len(lines))

	for i, line := range lines {
		grid[i] = []rune(line)
		// Example grid row: [46 46 46 35 46 46 46 46 46 35 46 46 46 46 ...]
	}

	output := 0

	var (
		//Start guard at -1 for "Not found yet"
		guardStartingRow = -1
		guardStartingCol = -1
		//Guard facing direction 0 while "not found"
		guardDirection = 0
	)

	/* Getting the starting pos of the guard */
	//Iterate over rows (i) in the grid map
	for row := range grid {
		if guardStartingRow >= 0 {
			break
		}
		//Iterate over the columns (j) in the grid row
		for col := range grid[row] {
			//Check the rune in the column for guard's starting pos
			if grid[row][col] == '^' {
				//Update the starting positions for the guard
				guardStartingRow = row
				guardStartingCol = col
				break
			}
		}
	}

	//Possible directions through the maze
	directions := [][2]int{
		//Guard direction is used as index to this array
		{-1, 0}, //up
		{0, 1},  //right
		{1, 0},  //down
		{0, -1}, //left
	}

	//Coordinates that have already been seen
	seen := make(map[[2]int]bool)

	/*Moving the guard around*/
	for {
		//First we should mark the current postion as "seen"
		seen[[2]int{guardStartingRow, guardStartingCol}] = true

		//Find the next possible move based on dir that guard is facing
		//e.g. directions[0] is up, so it moves up one row in the same column
		currentDirection := directions[guardDirection]
		nextGuardRow := guardStartingRow + currentDirection[0]
		nextGuardCol := guardStartingCol + currentDirection[1]

		//If at any point we're outside of the grid, break
		if nextGuardRow < 0 || nextGuardRow >= len(grid) || nextGuardCol < 0 || nextGuardCol >= len(grid[0]) {
			break
		}

		if grid[nextGuardRow][nextGuardCol] == '#' {
			//If obstacle is encountered, increment the direction and use modulo 4
			//the %4 is needed for the final possible movement of left back to up e.g. facing left = direction[3] (3+1) = 4 which is out of index range in directions
			//using modulo we can reset back down to 4 every time like a hard reset (3+1)%4 = 0
			guardDirection = (guardDirection + 1) % 4
			currentDirection = directions[guardDirection]
			nextGuardRow = guardStartingRow + currentDirection[0]
			nextGuardCol = guardStartingCol + currentDirection[1]

			if nextGuardRow < 0 || nextGuardRow >= len(grid) || nextGuardCol < 0 || nextGuardCol >= len(grid[0]) {
				break
			}
		}
		guardStartingRow = nextGuardRow
		guardStartingCol = nextGuardCol
	}
	output = len(seen)
	fmt.Println("Solution One:", output)
}

func day6_2(input string) {
	//Same as above
	lines := strings.Split(strings.TrimSpace(input), "\n")
	//Create a 2D slice of runes, one per line of input
	grid := make([][]rune, len(lines))

	for i, line := range lines {
		grid[i] = []rune(line)
		// Example grid row: [46 46 46 35 46 46 46 46 46 35 46 46 46 46 ...]
	}

	output := 0

	var (
		//Start guard at -1 for "Not found yet"
		guardStartingRow = -1
		guardStartingCol = -1
		//Guard facing direction 0 while "not found"
		guardDirection = 0
	)
	for row := range grid {
		if guardStartingRow >= 0 {
			break
		}
		//Iterate over the columns (j) in the grid row
		for col := range grid[row] {
			//Check the rune in the column for guard's starting pos
			if grid[row][col] == '^' {
				//Update the starting positions for the guard
				guardStartingRow = row
				guardStartingCol = col
				break
			}
		}
	}

	//Possible directions through the maze
	directions := [][2]int{
		//Guard direction is used as index to this array
		{-1, 0}, //up
		{0, 1},  //right
		{1, 0},  //down
		{0, -1}, //left
	}

	/* Adjust previous code starting here */
	//"For each row and each col value within the row"
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			//We need to assume that there is no loop and then switch to true when one starts
			isLooping := false

			//"Check to see if there is an clear path"
			if grid[row][col] != '.' {
				continue
			}
			grid[row][col] = '#'
			//Make a new map for seen coordinates
			//Example visualization of this map  : `[3]int{1,2,3}: true`
			seen := make(map[[3]int]bool)
			var (
				//Current location placeholders for guard
				curRow = guardStartingRow
				curCol = guardStartingCol
				curDir = guardDirection
			)
			//Begin loop to move guard
			for {
				guardCoords := [3]int{curRow, curCol, curDir}
				//Determine whether the guard has been to these coordinates
				if seen[guardCoords] {
					isLooping = true
					break
				}
				//If not, then we can add them to the list to search for again later
				seen[guardCoords] = true

				nextGuardRow := curRow + directions[curDir][0]
				nextGuardCol := curCol + directions[curDir][1]

				//Same OoB index check as above :
				//If at any point we're outside of the grid, break
				if nextGuardRow < 0 || nextGuardRow >= len(grid) || nextGuardCol < 0 || nextGuardCol >= len(grid[0]) {
					break
				}

				//If the next possible move is an obstacle, turn right
				if grid[nextGuardRow][nextGuardCol] == '#' {
					//Same as above, handle cardinal dirs and modulos 4 so we loop back around to north
					curDir = (curDir + 1) % 4
				} else {
					curRow = nextGuardRow
					curCol = nextGuardCol
				}
			}
			if isLooping {
				//increment the count of possible ways for this to loop infinitely
				output++
			}
			grid[row][col] = '.'
		}
	}

	fmt.Println("Solution Two:", output)
}
