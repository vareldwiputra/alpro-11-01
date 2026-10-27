package main

import "fmt"

func main() {
	password := ""
	kesempatan := 0
	saldo := 10000000

	for password != "1234" && kesempatan < 3 {
		fmt.Print("Masukkan password ATM: ")
		fmt.Scan(&password)

		if password != "1234" {
			fmt.Println("Password salah. Silakan coba lagi.")
			kesempatan++
		}
	}

	if password != "1234" {
		fmt.Println("ATM terblokir. Program selesai.")
	} else {
		fmt.Println("Login berhasil!")

		pilihan := 0

		for pilihan != 3 {
			fmt.Println("Menu ATM:")
			fmt.Println("1. Lihat Saldo")
			fmt.Println("2. Tarik Uang")
			fmt.Println("3. Keluar")
			fmt.Print("Pilih: ")
			fmt.Scan(&pilihan)

			if pilihan == 1 {
				fmt.Println("Saldo Anda: Rp", saldo)
			} else if pilihan == 2 {
				var tarik int

				fmt.Print("Masukkan jumlah uang yang ingin ditarik: Rp ")
				fmt.Scan(&tarik)

				if tarik > 0 && tarik <= saldo {
					saldo = saldo - tarik
					fmt.Println("Anda berhasil menarik uang sebesar Rp", tarik)
					fmt.Println("Saldo tersisa: Rp", saldo)
				} else {
					fmt.Println("Saldo tidak mencukupi.")
				}
			} else if pilihan == 3 {
				fmt.Println("Keluar dari sistem. Terima kasih!")
			}
		}
	}
}
