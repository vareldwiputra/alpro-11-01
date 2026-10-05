package main

import "fmt"

func main() {
	var r float64
	const pi float64 = 22.0 / 7.0

	fmt.Scan(&r)
	fmt.Println("Masukan nilai r: ")
	fmt.Println("Luas Permukaan Bola:", 4*pi*r*r)
}
