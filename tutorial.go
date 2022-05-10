package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	var name string = "Jeisa Raja"
	fmt.Println(name)
	number := "nulls"
	fmt.Printf("%T", number)
}

// := digunakan untuk membuat variabel tanpa data types (mirip seperti python, go yang akan meng assign data types nya)
// 'Go ruh' digunakan untuk menjalankan program go
// 'Go build' digunakan untuk compile program go

// %s	the uninterpreted bytes of the string or slice
// %q	a double-quoted string safely escaped with Go syntax
// %x	base 16, lower-case, two characters per byte
// %X	base 16, upper-case, two characters per byte

// %b	base 2
// %c	the character represented by the corresponding Unicode code point
// %d	base 10
// %o	base 8
// %O	base 8 with 0o prefix
// %q	a single-quoted character literal safely escaped with Go syntax.
// %x	base 16, with lower-case letters for a-f
// %X	base 16, with upper-case letters for A-F
// %U	Unicode format: U+1234; same as "U+%04X"
