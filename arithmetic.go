package main

import "fmt"

func main() {
	var number int64 = 80
	var number2 float64 = 2.5

	answer := float64(number) * number2
	fmt.Printf("The answer is %g", answer)
}
