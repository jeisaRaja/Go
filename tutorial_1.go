package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Hey what do you say? : ")
	scanner.Scan()
	text, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	fmt.Printf("You typed: %d", 2022-text)
}
