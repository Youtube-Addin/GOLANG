package main

import "fmt"

func getFullname() (string, string) {
	return "Addin", "Rizal"
}

func main() {
	// firstName, lastName := getFullname()

	// fmt.Println(firstName, lastName)


	// Menghiraukan return value. 
	firstName, _ := getFullname()

	fmt.Println(firstName)
}
