package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Println("masukan sisi segitiga: ")
	fmt.Scan(&a, &b, &c)

	if a == b && b == c {
		fmt.Println("segitiga sama sisi")
	} else if a == b || a == c || b == c {
		fmt.Println("segitiga sama kaki")
	} else {
		fmt.Println("segitiga sembarang")
	}
}
