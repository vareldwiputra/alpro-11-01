package main

import "fmt"

func main() {
	var a, b int
	var max, min, jumlah, banyak int

	fmt.Scan(&a)

	max = a
	min = a
	jumlah = a
	banyak = 1

	for {
		fmt.Scan(&b)
		if a == b {
			break
		}
		if b > max {
			max = b
		}
		if b < min {
			min = b
		}
		jumlah = jumlah + b
		banyak = banyak + 1

		a = b
	}

	rata := float64(jumlah) / float64(banyak)

	fmt.Printf("%d %d %.2f\n", max, min, rata)
}
