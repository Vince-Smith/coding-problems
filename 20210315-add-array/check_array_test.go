package main

import "testing"


func TestEmptyArray(test *testing.T) {
	_, err := checkArray([]int{}, 100)

	if err == nil {
		test.Error("Expected an error")
	}
}

func TestSingleElementArray(test *testing.T) {
	_, err := checkArray([]int{1}, 100)

	if err == nil {
		test.Error("Expected an error")
	}
}

func TestValidArrays(test *testing.T) {
	type testsCase struct {
		inputArray 	[]int
		inputSum 		int
		expected 		bool
	}

	cases := []testsCase {
		// True
		{
			inputArray: []int{10, 15, 3, 7},
			inputSum: 17,
			expected: true,
		},
		{
			inputArray: []int{1, 1, 1, 1, 1, 1},
			inputSum: 2,
			expected: true,
		},
		{
			inputArray: []int{0, 100, 20},
			inputSum: 20,
			expected: true,
		},
		{
			inputArray: []int{20, -4},
			inputSum: 16,
			expected: true,
		},
		{
			inputArray: []int{3, -4},
			inputSum: -1,
			expected: true,
		},
		// False
		{
			inputArray: []int{0, 0},
			inputSum: 1,
			expected: false,
		},
		{
			inputArray: []int{1, 2, 3, 4, 5},
			inputSum: 100,
			expected: false,
		},
	}

	for _, c := range cases {
		actual, err := checkArray(c.inputArray, c.inputSum)

		if err != nil {
			test.Error("Unexpected error")
		}

		if actual != c.expected {
			test.Errorf("Expected %t, got %t", c.expected, actual)
		}
	}
}
