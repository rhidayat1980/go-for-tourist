// Section 02 - Lecture 03 : Strings and rune values

package main

import "fmt"

func main() {

	// simple double quoted strings using " and "
	fmt.Println("Hello, World!")
	fmt.Println("Hello, 世界!")

	/*
		a rune is 32 bit number which represents a code-point.
		A code point is how UTF-8 (unicode) characters are encoded.
		A rune is big enough to represents any printable characters,
		even in languages like chinese with thousands of characters.
	*/
	fmt.Println('H')
	fmt.Println('e')
	fmt.Println('l')
	fmt.Println('l')
	fmt.Println('o')
	fmt.Println(',')
	fmt.Println(' ')
	fmt.Println('世')
	fmt.Println('界')
	fmt.Printf("19990 = %c\n", 19990)
	fmt.Printf("19990 = %d\n", 19990)
	fmt.Printf("19990 = %X\n", 19990)

	// escape characters
	fmt.Println("Hello", 10, "World!")
	fmt.Println('\n')
	fmt.Println("Hello, \n World!")
	fmt.Println('\\', "\\")
	fmt.Println('"', "\"")

	// a much longer string that can include just about anything.
	fmt.Println(`This is very long string in golang.
	it is called a "raw string" and can span several lines.
	NOTE: strings which are enclosed by "s are not able to do this,
	you would have to use an escape characters on the".`)

	// string concatenation using '+' operator for strings.
	fmt.Println("This is another very long string using \"." +
		"But it is getting to long for my screen." +
		"i can add more, but be sure to use the + concatenation" +
		" operator on the same line of the preceding string.")
}
