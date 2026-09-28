package main

import "fmt"

func main() {
	var name string

	name = "Varel Dwi Putra"

	fmt.Println("nama : ", name)
	//Variable bisa tanpa menggunakan kata kunci
	var lastName = "Putra"
	fmt.Println("nama belakang : ", lastName)
	//Variable bisa menggunakan :=
	middleName := "Dwi"
	fmt.Println("nama tenggah : ", middleName)

	firstName := "Varel"
	fmt.Println("nama depan : ", firstName)

	fmt.Println(firstName)
	fmt.Println(name)
}
