package main

import "fmt"

func main() {
	var nilai int

	for {
		fmt.Print("Masukkan nilai Anda: ")
		fmt.Scan(&nilai)

		if nilai == -1 {
			break
		}

		if nilai >= 60 {
			fmt.Println("Anda lulus ujian.")
		} else {
			fmt.Println("Anda tidak lulus ujian.")
		}
	}

	fmt.Println("Program selesai.")
}
