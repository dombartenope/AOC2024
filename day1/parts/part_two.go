package parts

func PartTwo(left, right []int) int {
	leftDupes := findDupes(left)
	rightDupes := findDupes(right)

	total := 0

	for k, leftValue := range leftDupes {
		if rightValue, exists := rightDupes[k]; exists {
			total += k * leftValue * rightValue
		} else {
			total += k * leftValue * 0
		}
	}

	return total
}

// Identify duplicate keys in the maps
func findDupes(list []int) map[int]int {
	dupeFrequency := make(map[int]int)
	for _, v := range list {
		dupeFrequency[v]++
	}
	return dupeFrequency
}
