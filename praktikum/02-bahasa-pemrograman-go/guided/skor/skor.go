package main

import "fmt"

func main() {
	var nama string
	var skorMatematika, skorBahasaInggris int

	//Membaca Input
	fmt.Scan(&nama)
	fmt.Scan(&skorMatematika)
	fmt.Scan(&skorBahasaInggris)

	//Menghitung total & rata-rata
	total := skorMatematika + skorBahasaInggris
	rataRata := total / 2

	//Menampilkan Output
	fmt.Println(nama)
	fmt.Println(total)
	fmt.Println(rataRata)

}
