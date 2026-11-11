# <h1 align="center">Laporan Praktikum Modul [7] - [Perulangan While]</h1>
<p align="center">[Varel Dwi Putra] - [109092600024]</p>

## Dasar Teori



## Guided

### 1. Kabisat.go

package main

import "fmt"

func main() {
	var tahun int
	fmt.Scan(&tahun)

	if (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0) {
		fmt.Println("Ya")
	} else {
		fmt.Println("Tidak")
	}
}

#### Deskripsi
Program dalam bahasa Go yang dibuat untuk menentukan apakah sebuah tahun merupakan tahun kabisat atau bukan.

### 2. Segitiga.go

package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a <= 0 || b <= 0 || c <= 0 || a+c <= b || b+c <= a || a+b <= c {
		fmt.Println("Bukan Segitiga")
	} else if a == b && b == c {
		fmt.Println("Segitiga Sama Sisi")
	} else if a == b || a == c || c == b {
		fmt.Println("Segitiga Sama Kaki")
	} else {
		fmt.Println("Segitiga Sembarang")
	}
}

#### Deskripsi
Program yang menerima panjang tiga sisi segitiga (a, b, dan c), kemudian menentukan jenis segitiga tersebut dengan ketentuan berikut: 

• Jika salah satu sisi bernilai ≤ 0, atau jumlah dua sisi mana pun tidak lebih besar dari sisi lainnya (misalnya a + b <= c), tampilkan Bukan segitiga. 

• Jika ketiga sisi sama panjang, tampilkan Segitiga sama sisi. 

• Jika tepat dua sisi sama panjang, tampilkan Segitiga sama kaki. 

• Jika ketiga sisi berbeda, tampilkan Segitiga sembarang. 

### 3. Kalkulator.go

package main

import "fmt"

func main() {
	var a, b int
	var op string
	fmt.Scan(&a, &op, &b)

	switch op {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b == 0 {
			fmt.Println("Error: Undefined")
		} else {
			fmt.Printf("%.2f\n", float64(a)/float64(b))
		}
	case "%":
		if b == 0 {
			fmt.Println("Error: Undefined")
		} else {
			fmt.Println(a % b)
		}
	default:
		fmt.Println("Operator tidak valid")
	}
}

### Deskripsi
Program kalkulator sederhana dalam bahasa Go menggunakan switch. Program menerima dua bilangan bulat dan sebuah operator, lalu menampilkan hasil operasinya dengan ketentuan: 

• Jika operator / atau % digunakan dengan bilangan kedua bernilai 0, tampilkan Error: pembagian dengan nol. 

• Jika operator selain yang ada pada tabel, tampilkan Operator tidak valid. 

### 4. Ongkir.go

package main

import (
	"fmt"
	"math"
)

func main() {
	var berat float64
	var zona int
	var layanan string

	fmt.Scan(&berat, &zona, &layanan)

	if berat > 30 {
		fmt.Println("Berat melebihi batas maksimum")
		return

	}

	var tarifPerkg int

	switch zona {
	case 1:
		tarifPerkg = 10000
	case 2:
		tarifPerkg = 15000
	case 3:
		tarifPerkg = 25000
	default:
		fmt.Println("Zona tidak valid")
		return
	}

	beratBulat := int(math.Ceil(berat))
	total := beratBulat * tarifPerkg

	if layanan == "Express" {
		total += total / 2
	}

	fmt.Println(beratBulat)
	fmt.Println(total)
}

### Deskripsi
Program yang dibuat untuk menghitung ongkos kirim berdasarkan berat paket, zona tujuan, dan jenis layanan dengan ketentuan berikut: 

• Berat paket (dalam kg, bilangan desimal) dibulatkan ke atas ke kilogram terdekat, misalnya 2.3 kg dihitung 3 kg. 

• Layanan Express dikenakan biaya tambahan 50% dari ongkos kirim. Layanan Reguler tidak dikenakan biaya tambahan. 

• Jika berat paket lebih dari 30 kg, tampilkan Berat melebihi batas maksimum (30 kg). 

• Jika zona selain 1, 2, atau 3, tampilkan Zona tidak valid. 

## Unguided

### 1. ATM



##### Output


#### Deskripsi


## Kesimpulan


## Referensi
