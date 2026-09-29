package main

import "fmt"

func main() {
	var pi float64 = 3.14
	var r float64
	var luas float64

	fmt.Scan(&r)
	fmt.Scan(&luas)

	luas = pi * r * r

	fmt.Println(luas)

}
