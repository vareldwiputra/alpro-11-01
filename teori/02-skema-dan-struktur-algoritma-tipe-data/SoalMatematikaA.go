package main

import "fmt"

func main() {
	var x int
	var y int
	var hasil float64

	fmt.Scan(&x)
	fmt.Scan(&y)

	hasil = 1.0/(3*float64(x)*float64(x)+10) + 10*float64(y) + 7

	fmt.Println(hasil)
}

/*PSEUDOCODE
ALGORITMA HitungFungsi

DEKLARASI
    x, y : integer
    hasil : real

DESKRIPSI
    Input(x, y)

    hasil ← 1 / (3 × x² + 10) + 10 × y + 7

    Output(hasil)

SELESAI
*/
