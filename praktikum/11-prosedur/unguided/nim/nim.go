package main

import "fmt"

const NMAX = 2002

func cariNIM(T [NMAX]int, n, x int, idx *int) {
	*idx = -1

	for i := 0; i < n; i++ {
		if T[i] == x {
			*idx = i
			break
		}
	}
}

func ubahNIM(T *[NMAX]int, idx, baru int) {
	T[idx] = baru
}

func cetak(T [NMAX]int, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(T[i])
	}
}

func main() {
	var T [NMAX]int
	var n, cari, baru, idx int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&T[i])
	}

	fmt.Scan(&cari)

	cariNIM(T, n, cari, &idx)

	if idx != -1 {
		fmt.Printf("Mahasiswa dengan NIM %d terdaftar pada posisi ke-%d\n", cari, idx+1)

		fmt.Scan(&baru)

		ubahNIM(&T, idx, baru)

		fmt.Println("Daftar NIM setelah terjadi perubahan data:")
		cetak(T, n)
	} else {
		fmt.Println("NIM tidak ditemukan dalam sistem")
	}
}
