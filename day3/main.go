package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {

	// testInput1 := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

	in, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer in.Close()

	//Find exact cases where "mul()" contains to uni-chars of type %d separated by a comma
	p1_regex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	total := 0
	var matches [][]string

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		matches = p1_regex.FindAllStringSubmatch(scanner.Text(), -1)

		for _, match := range matches {
			if len(match) == 3 {
				num1, err1 := strconv.Atoi(match[1])
				num2, err2 := strconv.Atoi(match[2])

				if err1 == nil && err2 == nil {
					total += num1 * num2
				}

			}
		}

	}
	fmt.Println("Part One Solution:", total)
	fmt.Println("Part Two Solution:", day3_part2())

}

func day3_part2() int { // TODO add input for *os.File
	in, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer in.Close()

	p2_regex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)

	matchFound := true
	var output int

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		matches := p2_regex.FindAllStringSubmatch(scanner.Text(), -1)

		for _, match := range matches {

			if match[0] == "do()" {
				matchFound = true
				continue

			} else if match[0] == "don't()" {
				matchFound = false
				continue

			} else if !matchFound {
				continue
			}
			num1, err := strconv.Atoi(match[1])
			if err != nil {
				log.Fatalf("error: %s", err)
			}

			num2, err := strconv.Atoi(match[2])
			if err != nil {
				log.Fatalf("error: %s", err)
			}

			output += num1 * num2
		}
	}
	return output
}
