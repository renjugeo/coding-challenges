package main

import "fmt"

func main() {
	fmt.Println(pairExists([]int{10,15,3,7}, 17))
        fmt.Println(pairExists([]int{1,2,3,4,5,6,7}, 13))
        fmt.Println(pairExists([]int{1,2,3,4,5,6,7}, 14))
}

func pairExists(input []int, s int) bool {
	var hash = make(map[int]bool)
	for _, j := range input {
		diff := s - j
		_, ok := hash[diff]
		if !ok {
			hash[j] = true
		} else {
			return true
		}
	}
	return false
}
