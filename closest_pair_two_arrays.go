package main

import (
	"fmt"
	"math"
	"sort"
)

// Given two arrays of integers and a target
// Find a pair whose sum is closest to the target

func main() {
	// Input
	array1 := []int{-1, 3, 8, 2, 9, 5}
	array2 := []int{4, 1, 2, 10, 5, 20}
	target := 30

	diff := math.MaxUint32
	pair := []int{}

	//Sort arrays
	sort.Ints(array1)
	sort.Ints(array2)

	i := 0
	j := len(array2) - 1
	for i < len(array1) && j >= 0 {
		sum := array1[i] + array2[j]
		absDiff := math.Abs(float64(target - sum))
		if int(absDiff) < diff {
			diff = int(absDiff)
			pair = []int{i, j}
		}
		if sum > target {
			j--
		} else if sum < target {
			i++
		} else {
			pair = []int{i, j}
			break
		}
	}
	fmt.Println(array1[pair[0]], array2[pair[1]])

}
