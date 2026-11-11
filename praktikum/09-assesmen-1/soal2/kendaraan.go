package main

import "fmt"

func main() {
	var huruf string
	var m, k, t int

	for {
		fmt.Scan(&huruf)

		if huruf == "m" {
			m++
		} else if huruf == "k" {
			k++
		} else if huruf == "t" {
			t++
		} else {
			break
		}
	}

	fmt.Println("Tipe M =", m)
	fmt.Println("Tipe K =", k)
	fmt.Println("Tipe T =", t)
}
