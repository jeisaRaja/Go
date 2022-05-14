package main

import "fmt"

func print() {
	fmt.Println("WHATSAPP")
}

func add(x, y int) int {
	z := x + y
	return z
}

// Fungsi juga dapat mereturn 2 atau lebih value
func multiply(x, y int) (int, int) {
	return x * y, x * x
}

// Kita juga bisa ngasih label terhadap variabel return, jadi di akhir fungsi hanya perlu perintah return
// defer adalah perintah untuk menunggu return dari sebuah fungsi baru defer di eksekusi
func divide(x, y int) (hasil1 int, hasil2 int) {
	defer fmt.Println("this is defer")
	hasil1 = x / y
	hasil2 = y / x
	fmt.Println("before defer")
	return
}

func main() {
	fmt.Println(add(5, 5))
	fmt.Println(multiply(5, 6))
	fmt.Println(divide(8, 4))
}
