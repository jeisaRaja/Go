package main

import "fmt"

func main() {
	ans := 10
	switch {
	case ans < 20:
		fmt.Println("smaller than 20")
	case ans > 20:
		fmt.Println("bigger than 20")
	}

	switch ans {
	case 10:
		fmt.Println("its a 10")
	case 20:
		fmt.Println("its a 20")
	}

}
