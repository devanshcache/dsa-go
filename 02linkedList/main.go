package main

/*
Linked list - stored the data anywhere in the memory like Hash Table.
first element is called HEAD and the last element is called TAIL usually(not for everyone).

A linked list item is called "node". which conatines the data and memory address of next node(Pointer).
If the last node is pointing to null or nil that means its a tail of linked list.const

Prepend and Append is O(1)
Lookup, insert and delete is O(1)

we call it Traversal instend of Iteration, why because its different .. eg: we don't know the end.

Unlike array which is stored together in memory, linked list or a all nodes are scattered all over memory
kinda like Hash Table.

Linked list is ordered.

Pointer: this is where it is in the memory / reference.
Garbage Collection: Clean unsed memory example no pointer pointing to a object.
*/

func main() {
	CustomLinkedList()
}
