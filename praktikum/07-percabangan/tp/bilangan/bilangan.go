package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n == 0 {
		fmt.Println("Nol")
	} else {
		if n%2 == 0 {
			fmt.Print("Genap, ")
		} else {
			fmt.Print("Ganjil, ")
		}

		if n > 0 {
			fmt.Println("Positif")
		} else {
			fmt.Println("Negatif")
		}
	}
}
