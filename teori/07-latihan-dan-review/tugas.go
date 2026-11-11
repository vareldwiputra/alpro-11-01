package main

import "fmt"

func main() {
	var n, m int
	var kpk int

	fmt.Scan(&n, &m)

	kpk = n

	for kpk%n != 0 || kpk%m != 0 {
		kpk = kpk + n
	}

	fmt.Println(kpk)
}
