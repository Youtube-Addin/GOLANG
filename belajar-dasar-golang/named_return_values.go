package main

import "fmt"

// Cara 1 deklarasi ulang

// func getComplateName()(string, string, string){
// 	firstName := "Addin"
// 	middleName := "Hadi"
// 	lastName := "Rizal"

// 	return firstName, middleName, lastName
// }

// Cara 2 tanpa deklarasi ulang, tanpa pakai gini (:=)
// opsi 1
// func getComplateName()(firstName string,middleName string,lastName string){
// 	firstName = "Addin"
// 	middleName = "Hadi"
// 	lastName = "Rizal"

// 	return firstName, middleName, lastName
// }

// opsi 2 karena pas aja 3 return valuesnya string jadi langsung aja tipe datanya ditaruh diakhir
func getComplateName()(firstName ,middleName , lastName string){
	firstName = "Addin"
	middleName = "Hadi"
	lastName = "Rizal"

	return firstName, middleName, lastName
}
func main(){
	a,b,c := getComplateName()
	fmt.Println(a,b,c)
}