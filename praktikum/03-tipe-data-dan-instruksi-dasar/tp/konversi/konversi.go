package main

import "fmt"

func main() {
	var mil float64 = 1.6
	var x float64

	fmt.Scan(&x)

	killometer := x * mil

	fmt.Println(killometer)
}
