package main

import (
	"fmt"
)

func main() {
	var arr [64]string
	arr[0] = "Whatsapp"
	fmt.Println(arr[0])

	arr2 := []int{1, 2, 3, 5}
	fmt.Println(arr2)

	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}

	fmt.Printf("this is a %T", arr)
}
