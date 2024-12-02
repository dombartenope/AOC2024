package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/dombartenope/AOC2024.git/parts"
)

func main() {

	leftSlice := []int{}
	rightSlice := []int{}

	//Open and read puzzle input
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer f.Close()

	/* READ & SEPARATE NUMBERS */
	//Define reader
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		//iterate over each line
		line := scanner.Text()
		words := strings.Fields(line)

		for i, v := range words {
			//Conv v to num of type int
			num, err := strconv.Atoi(v)
			if err != nil {
				log.Fatalf("error: %s", err)
			}
			//Append based on index (0 left, 1 right)
			if i > 0 {
				rightSlice = append(rightSlice, num)
			} else {
				leftSlice = append(leftSlice, num)
			}
		}
	}

	sort.Ints(leftSlice)
	sort.Ints(rightSlice)

	//PART ONE
	solutionOne := parts.PartOne(leftSlice, rightSlice)
	fmt.Printf("Solution One: %d\n", solutionOne)

	//PART TWO
	solutionTwo := parts.PartTwo(leftSlice, rightSlice)
	fmt.Printf("Solution Two: %v\n", solutionTwo)

}
