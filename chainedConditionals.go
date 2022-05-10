package main

import "fmt"

func main() {
	var x int64 = 20
	var y int64 = 80
	var z int64 = 60
	test := x < y && x < z

	fmt.Printf("%t", test)

}
