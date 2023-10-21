package main

import "fmt"

func hashTable() {

	m := map[string]string{
		"a": "aaaa",
		"b": "bbbb",
		"c": "cccc",
		"d": "dddd",
	}

	// lookup or accessing value - O(1)
	fmt.Println(m["a"])

	// Insert or adding new value - O(1)
	m["z"] = "zzzz"

	// deleteing - O(1)
	delete(m, "b")

	fmt.Println(m)

}
