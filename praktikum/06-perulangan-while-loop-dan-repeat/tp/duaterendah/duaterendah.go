package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var nama string
	var nilai int

	var nama1, nama2 string
	nilai1 := 1000000
	nilai2 := 1000000

	i := 1
	for i <= n {
		fmt.Scan(&nama, &nilai)

		if nilai < nilai1 {
			nilai2 = nilai1
			nama2 = nama1

			nilai1 = nilai
			nama1 = nama
		} else if nilai < nilai2 {
			nilai2 = nilai
			nama2 = nama
		}

		i++
	}

	fmt.Println(nama1)
	fmt.Println(nama2)
}
