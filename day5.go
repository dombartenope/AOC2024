package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func day5_1(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	output := 0

	rules := make(map[int][]int)
	updates := [][]int{}

	//Vars to check if my splits are correct
	var (
		lineCount       int
		commaCount      int
		whitespaceCount int
	)

	for _, line := range lines {
		if strings.Contains(line, "|") {
			ruleNums := strings.Split(line, "|")
			num1, err := strconv.Atoi(ruleNums[0])
			if err != nil {
				log.Fatalf("error: %s", err)
			}

			num2, err := strconv.Atoi(ruleNums[1])
			if err != nil {
				log.Fatalf("error: %s", err)
			}

			rules[num1] = append(rules[num1], num2)
			lineCount++
		} else if strings.Contains(line, ",") {
			currentIndex := []int{}
			updateNums := strings.Split(line, ",")
			for _, v := range updateNums {
				num1, err := strconv.Atoi(v)
				if err != nil {
					log.Fatalf("error: %s", err)
				}

				currentIndex = append(currentIndex, num1)

			}
			updates = append(updates, currentIndex)
			commaCount++

		} else {
			//Should cover cases where there is a newline with no data to read
			whitespaceCount++
			continue
		}
	}
	//To double check we're reading the file correctly
	fmt.Printf("Number of lines split by pipe: %d\nNumber of lines split by comma: %d\nNumber of empty lines: %d\n", lineCount, commaCount, whitespaceCount)

	for _, updateLine := range updates {
		allow := true
		//fmt.Println(j) = [27 52 74 23 73 97 22 58 89 11 18 25 16 72 99 62 26]
		//We should be able to loop from the end back to the start
		for i := len(updateLine) - 1; i >= 0; i-- {
			// Grab each number in the updates line moving RTL
			numInLine := updateLine[i] //numInLine will be a single integer in this loop

			//Iterate over each index down to i
			for _, num := range updateLine[:i] {
				//Iterate over each number in rules map
				for _, rule := range rules[numInLine] {
					/* 	fmt.Println(num, rule) outputs the number vs the rule
					55 27
					55 89
					55 97
					55 64
					*/
					//Check if number in line == rule num
					if num == rule {
						allow = false
						break
					}
				}
				if !allow {
					break
				}
			}
		}
		if allow {
			//We can find the middle of the line by dividing Len/2
			lineLen := len(updateLine)
			middle := updateLine[lineLen/2]
			output += middle

		}
	}

	fmt.Println("Solution One:", output)
}

func day5_2(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	output := 0

	rules := make(map[int][]int)
	updates := [][]int{}

	//No pipe,comma,whitespace count needed -- same input

	for _, line := range lines {
		if strings.Contains(line, "|") {
			ruleNums := strings.Split(line, "|")
			num1, err := strconv.Atoi(ruleNums[0])
			if err != nil {
				log.Fatalf("error: %s", err)
			}

			num2, err := strconv.Atoi(ruleNums[1])
			if err != nil {
				log.Fatalf("error: %s", err)
			}

			rules[num1] = append(rules[num1], num2)
		} else if strings.Contains(line, ",") {
			currentIndex := []int{}
			updateNums := strings.Split(line, ",")
			for _, v := range updateNums {
				num1, err := strconv.Atoi(v)
				if err != nil {
					log.Fatalf("error: %s", err)
				}

				currentIndex = append(currentIndex, num1)

			}
			updates = append(updates, currentIndex)

		} else {
			//Should cover cases where there is a newline with no data to read
			continue
		}
	}

	for _, currentUpdate := range updates {
		validEntry := true
		for i := len(currentUpdate) - 1; i > 0; i-- {
			currentNum := currentUpdate[i]
			for _, num := range currentUpdate[:i] {
				for _, ruleNum := range rules[currentNum] {
					if num == ruleNum {
						validEntry = false
						break
					}
				}

				if !validEntry {
					break
				}
			}
		}

		if !validEntry {
			//Store ints sorted High to Low
			ordered := []int{}
			//Map of ints with T/F values
			numsLeft := map[int]bool{}
			//Map ints to a slice of ints
			dependencies := make(map[int][]int)

			for _, num := range currentUpdate {
				numsLeft[num] = true

				for _, dep := range rules[num] {
					dependencies[dep] = append(dependencies[dep], num)
				}
			}

			for len(numsLeft) > 0 {
				for num := range numsLeft {
					if len(dependencies[num]) == 0 {
						ordered = append(ordered, num)
						delete(numsLeft, num)

						for key, val := range dependencies {
							newList := []int{}

							for _, n := range val {
								if n != num {
									newList = append(newList, n)
								}
							}

							dependencies[key] = newList
						}
					}
				}
			}

			middle := ordered[len(ordered)/2]
			output += middle
		}
	}
	fmt.Println("Solution Two:", output)

}
