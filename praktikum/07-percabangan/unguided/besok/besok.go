package main

import "fmt"

func main() {
	var tanggal, bulan, tahun int
	var jumlahHari int

	fmt.Scan(&tanggal, &bulan, &tahun)

	// Menentukan jumlah hari dalam bulan
	switch bulan {
	case 1, 3, 5, 7, 8, 10, 12:
		jumlahHari = 31
	case 4, 6, 9, 11:
		jumlahHari = 30
	case 2:
		if tahun%400 == 0 || (tahun%4 == 0 && tahun%100 != 0) {
			jumlahHari = 29
		} else {
			jumlahHari = 28
		}
	default:
		jumlahHari = 0
	}

	// Memeriksa tanggal
	if tanggal < 1 || tanggal > jumlahHari {
		fmt.Println("Tanggal tidak valid")
	} else {
		fmt.Println("Valid")

		tanggal++

		// Jika sudah melewati jumlah hari
		if tanggal > jumlahHari {
			tanggal = 1
			bulan++

			if bulan > 12 {
				bulan = 1
				tahun++
			}
		}

		fmt.Printf("%d/%d/%d\n", tanggal, bulan, tahun)
	}
}
