package main

import "fmt"

func main() {
	var n, lebar int

	fmt.Scan(&n)

	fmt.Scan(&lebar)
	terkecil := lebar

	for i := 1; i < n; i++ {
		fmt.Scan(&lebar)

		if lebar < terkecil {
			terkecil = lebar
		}
	}

	fmt.Println(terkecil)
}
