package main

import "fmt"

func main() {
	var x int64 = 1

	for x < 10 {
		fmt.Println("%v", x)
		x += 1
	}

	for y := 0; y < 10; y++ {
		fmt.Println(y)
	}
}
