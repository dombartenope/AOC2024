package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isSafeReport(nums []int) bool {
	var (
		diff int
		inc  int
		dec  int
	)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			diff = nums[i] - nums[i+1]
			dec++
		} else {
			diff = nums[i+1] - nums[i]
			inc++
		}

		if (diff < 1 || diff > 3) || (inc > 0 && dec > 0) {
			return false
		}
	}

	return true
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var (
		numList  []int
		count    int // Safe without modification (Solution One)
		p2_count int // Safe after applying Problem Dampener (Solution Two)
	)

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)

		for _, v := range words {
			num, err := strconv.Atoi(v)
			if err != nil {
				log.Fatalf("error: %s", err)
			}
			numList = append(numList, num)
		}

		// Check if the report is safe without modifications (Solution One)
		if isSafeReport(numList) {
			count++
			p2_count++ // Already safe, so also counts for Solution Two
		} else {
			// Try removing each level to check if the report becomes safe (Solution Two)
			for i := 0; i < len(numList); i++ {
				modifiedReport := append([]int{}, numList[:i]...)
				modifiedReport = append(modifiedReport, numList[i+1:]...)

				if isSafeReport(modifiedReport) {
					p2_count++
					break
				}
			}
		}

		// Reset numList for the next line
		numList = []int{}
	}

	fmt.Println("Solution One (Safe without modification):", count)
	fmt.Println("Solution Two (Safe with Problem Dampener):", p2_count)
}
