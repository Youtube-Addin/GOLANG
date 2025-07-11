package main

import "fmt"

func getHello(name string) string{
	return "Hello" + name
}

func main(){

	// kalau gini doang ga bakal keluar hasilnya karena kita tidak ambil data nya dari func diatas
	// getHello("Addin")

	// cara kedua
	result := getHello("Addin")
	fmt.Println(result)

	// cara ketiga
	fmt.Println(getHello())
}