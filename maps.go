package main

import "fmt"

func main() {
	var mp map[string]int = map[string]int{
		"home":   5,
		"school": 7,
	}

	mp1 := make(map[string]int)
	mp1["home"] = 8
	fmt.Println(mp)
	fmt.Println(mp1)
}
