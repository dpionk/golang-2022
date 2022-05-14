package main

import (
	"fmt"
	"sort"
)

func twoOldestAges(ages []int) [2]int {
	var twoOldestAges [2]int
	if len(ages) < 2 {
		return twoOldestAges
	} else {
		sort.Slice(ages, func(i, j int) bool {
			return ages[i] > ages[j]
		})
		twoOldestAges[0], twoOldestAges[1] = ages[1], ages[0]
		return twoOldestAges
	}
}

func main() {
	fmt.Println(twoOldestAges([]int{6, 5, 83, 5, 3, 18}))
}
