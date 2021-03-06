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

// %v	the value in a default format
// 	when printing structs, the plus flag (%+v) adds field names
// %#v	a Go-syntax representation of the value
// %T	a Go-syntax representation of the type of the value
// %%	a literal percent sign; consumes no value

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

// %b	decimalless scientific notation with exponent a power of two,
// 	in the manner of strconv.FormatFloat with the 'b' format,
// 	e.g. -123456p-78
// %e	scientific notation, e.g. -1.234456e+78
// %E	scientific notation, e.g. -1.234456E+78
// %f	decimal point but no exponent, e.g. 123.456
// %F	synonym for %f
// %g	%e for large exponents, %f otherwise. Precision is discussed below.
// %G	%E for large exponents, %F otherwise
// %x	hexadecimal notation (with decimal power of two exponent), e.g. -0x1.23abcp+20
// %X	upper-case hexadecimal notation, e.g. -0X1.23ABCP+20
