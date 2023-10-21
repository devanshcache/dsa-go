package main

import "fmt"

// Arrays are two types: Static vs Dynamic

func arraySlice() {
	// Arrays / Slices in Go
	arr := []string{"a", "b", "c", "d"}

	// Look-ups / Get
	fmt.Println(arr[2]) // O(1)

	// Push / Append
	arr = append(arr, "e") // O(1) but in Go / Golang it depends - explain below

	// Insert on first place of array or slices
	arr = append([]string{"x"}, arr...) // O(n)

	// Insert a value at a given index
	arr = insertToSlice(arr, "777", len(arr)/2) // O(n)

	fmt.Println(cap(arr))
	fmt.Println(len(arr))
}

func insertToSlice(a []string, element string, i int) []string {
	return append(a[:i], append([]string{element}, a[i:]...)...)
}

/*
# Length (len):
This tells you how many elements are currently in your list or slice.

# Capacity (cap):
This tells you how many elements your list or slice can hold without needing to make more room.

When you add something to the end of your list or slice:

If there's room (length is less than capacity), it's quick and easy. O(1)
If there's no room (length is equal to capacity), you might have to make the list or slice bigger, which can take more time. O(n)
*/
