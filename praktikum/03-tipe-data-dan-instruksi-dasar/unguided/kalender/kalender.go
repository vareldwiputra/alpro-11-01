package main

import "fmt"

func main() {
	var tahun, bulan, minggu, hari, sisa int
	fmt.Scan(&hari)

	tahun = hari / 360
	sisa = hari % 360
	bulan = sisa / 30
	sisa = sisa % 30
	minggu = sisa / 7
	hari = sisa % 7

	fmt.Println(tahun, bulan, minggu, hari)
}
