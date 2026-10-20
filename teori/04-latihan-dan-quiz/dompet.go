package main

import "fmt"

func main() {
	var total, x int
	for fmt.Scan(&x); x != 0; fmt.Scan(&x) {
		total += x
	}
	fmt.Println(total)
}
