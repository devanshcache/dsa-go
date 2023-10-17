package main

import "fmt"

// input: hello
// output: olleh

func reverceString() {
	str := "helloहह"
	//In Go, strings are immutable, which means you cannot change individual characters of a string directly. If you want to modify a string, you'll typically need to create a new string with the desired changes
	strRuns := []rune(str)
	// var out string

	for i := 0; i < len(strRuns)/2; i++ {
		temp := strRuns[i]
		strRuns[i] = strRuns[len(strRuns)-i-1]
		strRuns[len(strRuns)-i-1] = temp
	}

	fmt.Println(string(strRuns))
}

/*
Certainly! In Go, you can think of a string as a sequence of characters. You can access each character in the string as if they were stored in a box with a number (an index) on it. For example, the first character is in box number 0, the second character is in box number 1, and so on.

Because Go strings are encoded in a way that can handle many different characters and symbols from various languages, not all characters take up the same amount of space. Some characters are like small boxes, while others are like bigger boxes.

So, when you use square brackets and a number (index) to access a character in a string, you're opening the box at that position and seeing what's inside. You can use `%c` in `fmt.Printf` to print the character that's inside that box.

It's essential to be careful when you're trying to access characters at specific positions because if you ask for a box that doesn't exist (an index that is too big or too small), it can cause problems. So, always make sure you're asking for a box that's within the range of the available boxes in your string.
*/
