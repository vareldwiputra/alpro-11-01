package main

import "fmt"

func main() {
	var x int
	var y int

	fmt.Scan(&x)
	fmt.Scan(&y)

	sisa := x % y

	fmt.Println(sisa)
}
