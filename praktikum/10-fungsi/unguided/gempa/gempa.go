package main

import "fmt"

const NMAX = 999

type Gempa struct {
	nama      string
	tahun     int
	magnitudo float64
}

func kategori(m float64) string {
	if m < 4 {
		return "Minor"
	} else if m < 6 {
		return "Sedang"
	} else if m < 7 {
		return "Kuat"
	} else {
		return "Besar"
	}
}

func rataMagnitudo(T [NMAX]Gempa, n int) float64 {
	jumlah := 0.0

	for i := 0; i < n; i++ {
		jumlah = jumlah + T[i].magnitudo
	}

	return jumlah / float64(n)
}

func idxTerkuat(T [NMAX]Gempa, n int) int {
	idx := 0

	for i := 1; i < n; i++ {
		if T[i].magnitudo > T[idx].magnitudo {
			idx = i
		}
	}

	return idx
}

func hitungKategori(T [NMAX]Gempa, n int, k string) int {
	jumlah := 0

	for i := 0; i < n; i++ {
		if kategori(T[i].magnitudo) == k {
			jumlah++
		}
	}

	return jumlah
}

func main() {
	var T [NMAX]Gempa
	var n int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&T[i].nama, &T[i].tahun, &T[i].magnitudo)
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%s %d %.1f %s\n",
			T[i].nama, T[i].tahun,
			T[i].magnitudo, kategori(T[i].magnitudo))
	}

	rata := rataMagnitudo(T, n)
	idx := idxTerkuat(T, n)
	besar := hitungKategori(T, n, "Besar")

	fmt.Printf("Rata-rata magnitudo: %.2f\n", rata)
	fmt.Printf("Gempa terkuat: %s %d %.1f\n",
		T[idx].nama, T[idx].tahun, T[idx].magnitudo)
	fmt.Println("Jumlah gempa kategori Besar:", besar)
}
