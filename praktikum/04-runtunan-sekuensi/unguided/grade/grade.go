package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var nama string
	var nilai int
	var grade string

	fmt.Printf("Masukan Nama: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	nama = scanner.Text()

	fmt.Printf("Masukan Nilai: ")
	fmt.Scan(&nilai)

	switch {
	case nilai >= 90 && nilai <= 100:
		grade = "A"
	case nilai >= 80 && nilai < 90:
		grade = "B"
	case nilai >= 70 && nilai < 80:
		grade = "C"
	case nilai >= 60 && nilai < 70:
		grade = "D"
	case nilai >= 1 && nilai < 60:
		grade = "F"
	}

	fmt.Printf("%s Mendapatankan Nilai %s\n", nama, grade)
}
