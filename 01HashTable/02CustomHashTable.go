package main

import "fmt"

type MyHashTable struct {
	Data []interface{}
	Size int
}

func NewMyHashTable(size int) *MyHashTable {
	return &MyHashTable{
		Data: make([]any, size),
		Size: size,
	}
}

func CustomHashTable() {
	myhashTable := NewMyHashTable(10)

	myhashTable.Data = append(myhashTable.Data, [2]any{1, "one"})

	fmt.Println(myhashTable.Data)
}
