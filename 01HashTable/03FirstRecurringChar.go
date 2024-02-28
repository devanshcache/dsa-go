package main

import "fmt"

// First recurring character
// Given array x := [2,4,22,6,2,1,0]
// Output:  2

func FirstRecurringCharSolution0() {
	arr := []int{2, 1, 1, 5, 1, 2, 7}
Loop:
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				fmt.Println(arr[i])
				break Loop
			}
		}
	}
}

func FirstRecurringCharSolution1() {
	arr := []int{2, 1, 1, 5, 1, 2, 7, 0, 6}
	char := map[int]int{}

	for _, val := range arr {
		_, ok := char[val]
		if ok {
			fmt.Println(val)
			break
		}
		char[val] = 1
	}
}
