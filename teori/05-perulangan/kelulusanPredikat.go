package main

import "fmt"

func main() {
	var nilaiUtama, nilaiRemed int
	fmt.Scan(&nilaiUtama, &nilaiRemed)

	if nilaiUtama >= 70 {
		if nilaiUtama >= 85 {
			fmt.Println("Lulus, Sangat Baik")
		} else {
			fmt.Println("lulus, Baik")
		}
	} else {
		if nilaiRemed >= 70 {
			fmt.Println("Remedial, Lulus Remedial")
		} else {
			fmt.Println("Remedial, Tidak Lulus")
		}
	}
}
