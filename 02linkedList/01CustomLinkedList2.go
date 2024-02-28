package main

import "fmt"

// this is doubly linked list.

type Node2 struct {
	Value     int
	PrevValue *Node2
	NextValue *Node2
}
type LinkedList2 struct {
	Head   *Node2
	Tail   *Node2
	Length int
}

func NewLinkedList2(headValue int) *LinkedList2 {
	head := &Node2{
		Value:     headValue,
		PrevValue: nil,
		NextValue: nil,
	}
	return &LinkedList2{
		Head:   head,
		Tail:   head,
		Length: 1,
	}
}

func (l *LinkedList2) Append(value int) {
	newNode := Node2{
		Value:     value,
		PrevValue: l.Tail,
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

func (l *LinkedList2) Prepend(value int) {
	newNode := Node2{
		Value:     value,
		PrevValue: nil,
		NextValue: l.Head,
	}
	l.Head.PrevValue = &newNode
	l.Head = &newNode
	l.Length++
}

func (l *LinkedList2) Insert(value int, index int) {
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

	newNode := Node2{
		Value:     value,
		PrevValue: nil,
		NextValue: nil,
	}

	curr := l.Head
	for i := 0; i < index-1; i++ {
		curr = curr.NextValue
	}
	newNode.PrevValue = curr
	newNode.NextValue = curr.NextValue
	curr.NextValue.PrevValue = &newNode
	curr.NextValue = &newNode
	l.Length++
}

func (l *LinkedList2) Remove(index int) {
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

func (l *LinkedList2) Print() {
	curr := l.Head
	for curr.NextValue != nil {
		fmt.Println(curr)
		curr = curr.NextValue
	}
	if curr.NextValue == nil {
		fmt.Println(curr)
	}
}

func CustomLinkedList2() {
	myLl := NewLinkedList2(10)

	myLl.Append(20)
	myLl.Append(30)
	myLl.Append(40)
	myLl.Append(50)

	myLl.Prepend(0)

	myLl.Insert(1000, 4)
	// TODO:
	// myLl.Remove(0)

	myLl.Print()

	fmt.Printf("%+v\n", myLl)
}
