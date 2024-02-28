package main

import "fmt"

// this is singly linked list.

type Node struct {
	Value     int
	NextValue *Node
}
type LinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func NewLinkedList(headValue int) *LinkedList {
	head := &Node{
		Value:     headValue,
		NextValue: nil,
	}
	return &LinkedList{
		Head:   head,
		Tail:   head,
		Length: 1,
	}
}

func (l *LinkedList) Append(value int) {
	newNode := Node{
		Value:     value,
		NextValue: nil,
	}

	l.Tail.NextValue = &newNode
	l.Tail = &newNode
	l.Length++

	// if l.Length == 1 {
	// 	l.Head.NextValue = &newNode
	// 	l.Tail = &newNode
	// 	l.Length++
	// 	return
	// }

	// curr := l.Head
	// for curr.NextValue != nil {
	// 	curr = curr.NextValue
	// }

	// curr.NextValue = &newNode

	// l.Tail = &newNode
	// l.Length++
}

func (l *LinkedList) Prepend(value int) {
	newNode := Node{
		Value:     value,
		NextValue: l.Head,
	}

	l.Head = &newNode
	l.Length++
}

func (l *LinkedList) Insert(value int, index int) {
	if index >= l.Length {
		panic(fmt.Sprintln("length of linked list is", l.Length, "and you passed index of", index))
	}

	if index == 0 {
		l.Prepend(value)
		return
	}

	// just for fun
	if index == l.Length {
		fmt.Println("HEHEHE")
		l.Append(value)
		return
	}

	newNode := Node{
		Value:     value,
		NextValue: nil,
	}

	curr := l.Head
	for i := 0; i < index-1; i++ {
		curr = curr.NextValue
	}

	newNode.NextValue = curr.NextValue
	curr.NextValue = &newNode
	l.Length++
}

func (l *LinkedList) Remove(index int) {
	if index >= l.Length {
		panic(fmt.Sprintln("length of linked list is", l.Length, "and you passed index of", index))
	}

	if index == 0 {
		l.Head = l.Head.NextValue
		l.Length--
		return
	}

	currPrev := l.Head
	for i := 0; i < index-1; i++ {
		currPrev = currPrev.NextValue
	}

	currPrev.NextValue = currPrev.NextValue.NextValue
	l.Length--
}

func (l *LinkedList) Reverse() *LinkedList {
	values := []int{}

	curr := l.Head
	for curr != nil {
		values = append(values, curr.Value)
		curr = curr.NextValue
	}

	newLl := NewLinkedList(l.Tail.Value)

	for i := len(values) - 2; i >= 0; i-- {
		newLl.Append(values[i])
	}

	fmt.Println(newLl)
	return newLl
}

func (l *LinkedList) Print() {
	curr := l.Head
	for curr != nil {
		fmt.Println(curr.Value)
		curr = curr.NextValue
	}
}

func CustomLinkedList() {
	myLl := NewLinkedList(10)

	myLl.Append(20)
	myLl.Append(30)
	myLl.Append(40)
	myLl.Append(50)

	myLl.Prepend(11)
	myLl.Insert(0, 4)

	myLl.Remove(0)

	// TODO:
	// myLl.Reverse()
	myLl.Print()
	fmt.Println("-----")
	newLl := myLl.Reverse()
	newLl.Print()

	fmt.Printf("%+v\n", myLl)
}
