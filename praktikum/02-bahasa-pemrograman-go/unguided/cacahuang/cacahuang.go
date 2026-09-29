package main

import "fmt"

func main() {
	var uang int
	fmt.Scan(&uang)

	puluh := uang / 10000
	sisa := uang % 10000

	lima := sisa / 5000
	sisa = sisa % 5000

	seribu := sisa / 1000

	fmt.Println(puluh, lima, seribu)
}
