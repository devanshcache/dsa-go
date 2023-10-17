package main

import (
	"fmt"
)

type Array struct {
	Length int
	Data   map[int]string
}

func NewArray() Array {
	return Array{Length: 0, Data: map[int]string{}}
}

func (a *Array) Get(index int) any {
	val, ok := a.Data[index]
	if !ok {
		panic("out of range error index")
	}
	return val
}

func (a *Array) Push(val string) {
	a.Data[a.Length] = val
	a.Length++
}

func (a *Array) Pop() {
	delete(a.Data, a.Length-1)
	a.Length--
}

func (a *Array) Delete(index int) {
	for i := index; i < a.Length; i++ {
		a.Data[i] = a.Data[i+1]
	}
	a.Pop()
}

func customArray() {
	// myArray := Array{}
	// myArray.Data = map[int]any{1: "2", 2: 2}
	myArray := NewArray()
	myArray.Push("hello")
	myArray.Push("hi")
	// myArray.Pop()
	myArray.Push("you")
	myArray.Push("bye")
	myArray.Delete(1)
	fmt.Println(myArray)
}
