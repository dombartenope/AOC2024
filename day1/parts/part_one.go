package parts


func PartOne(left, right []int) int  {

	/* FIND DIFF */
	var diff int
	diffList := []int {}
	for i, v := range  left {
		k := right[i]

		// fmt.Printf("D EBUG - Index [i] = %d\t, %d\t%d ", i, v, k)

		if v > k {
			diff = v - k
		} else {
			diff = k - v
		}
		diffList = append(diffList, diff)
	}

	/* SUM THE DIFFERENCES */
	var total int
	for i := 0; i < len(diffList); i++ {
		total += diffList[i]
	}
	return total
	//Total : 2066446
}
