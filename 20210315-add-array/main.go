package main

import "fmt"

func main() {
    fmt.Println(arrayChecker([]int{10, 15, 3, 7}, 17))
}

func arrayChecker(arr []int, minuend int) (bool) {
	dict := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		subtrahand := arr[i]
		difference := minuend - subtrahand

		// In Go, accessing a hashmap and finding nothing will result in a
		// "0 value" instead of nil. For integers this is just 0.
		//
		// In order to be certain you actually failed the lookup, you need to
		// look at the second value returned by the lookup.
		//
		// [key] => value, bool (where bool is whether or not the lookup was successful)
		if _, ok := dict[subtrahand]; ok {
			return true
		}

		dict[difference] = subtrahand
	}

	return false
}
