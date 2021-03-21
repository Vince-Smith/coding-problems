package main

import (
	"errors"
	"fmt"
)

func main() {
    fmt.Println(checkArray([]int{10, 15, 3, 7}, 17))
}

func checkArray(array []int, minuend int) (bool, error) {
	if len(array) < 2 {
		return false, errors.New("Array must contain at least two elements")
	}

	dict := make(map[int]int)
	for _, subtrahand := range array {
		difference := minuend - subtrahand

		if _, ok := dict[subtrahand]; ok {
			return true, nil
		}

		dict[difference] = subtrahand
	}

	return false, nil
}


