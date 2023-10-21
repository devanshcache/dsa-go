package main

import "fmt"

// a => [0, 3, 4, 31]
// b => [4, 6, 30]

func merageSortedArray() {
	aArray := []int{0, 3, 4, 31}
	bArray := []int{4, 6, 30, 31, 32}

	var mergedArrays []int // {0, 3, 4, 4, 6, 30, 31}

	i := 0
	j := 0

	for i < len(aArray) && j < len(bArray) {
		if aArray[i] < bArray[j] {
			mergedArrays = append(mergedArrays, aArray[i])
			i++
		} else {
			mergedArrays = append(mergedArrays, bArray[j])
			j++
		}
	}

	for i < len(aArray) {
		mergedArrays = append(mergedArrays, aArray[i])
		i++
	}

	for j < len(bArray) {
		mergedArrays = append(mergedArrays, bArray[j])
		j++
	}

	fmt.Println(mergedArrays)
}
