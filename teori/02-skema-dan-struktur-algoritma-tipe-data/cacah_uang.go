package main

import "fmt"

func main() {
	var uang int

	fmt.Scan(&uang)

	fmt.Println(uang/10000, "lembar")
	uang = uang % 10000

	fmt.Println(uang/500, "lembar")
	uang = uang % 500

	fmt.Println(uang/1000, "lembar")
	uang = uang % 1000
}
