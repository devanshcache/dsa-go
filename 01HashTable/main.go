package main

/*
Can have multiple name: HashTables, HashMaps, UnOrdered Maps, Dict in Python, Object in JavaScript is a type of Hash Table..

Hash Table -> Use something called hash function which convert our key/values to a hash (sha, sh256, ..)
No order
-----
If two values stores in same memory location then it becomes coalition then it stores like linked list form(or there are diff methods)
If there are many values in one address space due to coalition,
usually it slows down reading and writing with O(n/k) => O(n)

So occasionally lookup in HashTable can be O(n)

-----
when we need to iterate over the all the elements of a HashTable, it would O(n)
beacuse it would go over the hashtable memory to check if there is something stored there.
-----
In GO HashTable is called map "map[key-type]value-type"
example:
map[string]string
map[int]int
map[int]bool
map[string]int
...

*/
func main() {
	// hashTable()
	FirstRecurringCharSolution1()
}
