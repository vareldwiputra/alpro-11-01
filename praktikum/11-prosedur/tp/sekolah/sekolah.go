package main

import "fmt"

const NMAX = 1001

type tabMurid [NMAX]string

func cariMurid(T tabMurid, n int, nama string, posisi *int, ketemu *bool) {
	*ketemu = false
	*posisi = -1

	for i := 0; i < n; i++ {
		if T[i] == nama {
			*ketemu = true
			*posisi = i + 1
			break
		}
	}
}

func hitungHuruf(T tabMurid, n int, h byte, jumlah *int) {
	*jumlah = 0

	for i := 0; i < n; i++ {
		if T[i][0] == h {
			*jumlah = *jumlah + 1
		}
	}
}

func main() {
	var T tabMurid
	var n int
	var nama string
	var posisi, jumlah int
	var ketemu bool

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&T[i])
	}

	fmt.Scan(&nama)

	cariMurid(T, n, nama, &posisi, &ketemu)

	if ketemu {
		fmt.Printf("Murid terdaftar dan berada di urutan absen ke-%d\n", posisi)
	} else {
		fmt.Println("Murid tidak terdaftar")
	}

	hitungHuruf(T, n, nama[0], &jumlah)

	fmt.Printf("Jumlah murid dengan huruf awal %c: %d\n", nama[0], jumlah)
}
