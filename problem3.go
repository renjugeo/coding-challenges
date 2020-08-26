/*
Given an array of integers, return a new array such that each element at index i of the new array
is the product of all the numbers in the original array except the one at i
Note: can't use division

Solution:
a[i] = product of all numbers on the left * product of all nnumbers on the right

*/
package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5}

	b := make([]int, len(a))
	t := 1

	// Find the product of all numbers on the left side of the index
	for k, v := range a {
		b[k] = t
		t = v * t
	}
	t = 1
	// Find the product of all numbers on the right side of the index
	for i := len(a) - 1; i >= 0; i-- {
		b[i] = b[i] * t
		t = a[i] * t
	}

	fmt.Println(b)
}
