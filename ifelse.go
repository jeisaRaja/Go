package main

import "fmt"

func main() {
	var x int64 = 20

	if x > 100 {
		fmt.Printf("%d is bigger than 100", x)
	} else if x > 50 {
		fmt.Printf("%d is bigger than 50", x)
	} else {
		fmt.Printf("%d is too small", x)
	}
}
