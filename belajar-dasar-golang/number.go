package main

import "fmt"

// Error "redeclared" ini karena aku buat nama function double disini ada di file number.go ada juga di file helloworld.go
func main() {
	fmt.Println("Satu = ", 1)
	fmt.Println("Dua = ", 2)
	fmt.Println("Tiga koma lima = ", 3.5)
}