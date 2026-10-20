# <h1 align="center">Laporan Praktikum Modul [5] - [Perulangan For Loop]</h1>
<p align="center">[Varel Dwi Putra] - [109092600024]</p>

## Dasar Teori

### A. Bahasa Pemrograman Go dan Sekuensi
Bahasa Go (Golang) merupakan bahasa pemrograman bertipe terstruktur yang mengeksekusi perintah secara runtut (sequence) dari atas ke bawah. Sekuensi adalah struktur paling dasar dalam algoritma di mana setiap instruksi dijalankan tepat satu kali secara berurutan sesuai urutan penulisannya.

### ​B. Struktur Kontrol Percabangan di Go

#### ​1. Percabangan if-else
Percabangan if-else digunakan untuk mengevaluasi suatu kondisi boolean. Jika kondisi bernilai true, blok kode di dalam if akan dieksekusi. Untuk banyak kondisi, percabangan dapat disambung menggunakan else if atau dibuat bersarang (nested if).  

#### ​2. Percabangan switch-case
Pernyataan switch menyediakan cara yang lebih bersih untuk melakukan percabangan multi-kondisi. Pada Go, switch dapat ditulis tanpa ekspresi (kondisi dimasukkan langsung pada setiap case) untuk menggantikan struktur if-else yang panjang.

## Guided

### 1. Grade,go

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var nama string
	var nilai float64
	var grade string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	nama = scanner.Text()

	fmt.Scan(&nilai)

	if nilai >= 90 && nilai <= 100 {
		grade = "A"
	} else if nilai >= 80 && nilai < 90 {
		grade = "B"
	} else if nilai >= 70 && nilai < 80 {
		grade = "C"
	} else if nilai >= 60 && nilai < 70 {
		grade = "D"
	} else {
		grade = "F"
	}

	fmt.Printf("%s mendapatkan nilai %s\n", nama, grade)
}

#### Deskripsi
Program yang dibuat untuk menghitung total tabungan Budi setelah n hari. 

package main

import (
	"fmt" // Hanya memerlukan fmt untuk keperluan input dan output
)

func main() {
	// Pendefinisian variabel secara eksplisit beserta tipe datanya
	var usia int
	var gaji int
	var keterangan string

	// Membaca dua masukan berupa angka dari pengguna (usia dan gaji)
	// fmt.Scan otomatis memisahkan input berdasarkan spasi atau baris baru (enter)
	fmt.Scan(&usia)
	fmt.Scan(&gaji)

	// Menggunakan switch tanpa ekspresi.
	// Cara kerjanya sama persis seperti deretan if - else if.
	// Program akan mengecek dari atas ke bawah, dan menjalankan case pertama yang bernilai benar (true).
	switch {
	case usia < 18:
		keterangan = "Masih sekolah"
	case usia >= 18 && usia <= 25 && gaji >= 50:
		keterangan = "Muda sukses"
	case usia >= 18 && usia <= 25 && gaji < 50:
		keterangan = "Masih belajar hidup"
	case usia >= 26 && usia <= 40 && gaji >= 100:
		keterangan = "Pekerja mapan"
	case usia >= 26 && usia <= 40 && gaji < 100:
		keterangan = "Perlu perbaikan karier"
	case usia > 40 && gaji >= 150:
		keterangan = "Profesional berpengalaman"
	case usia > 40 && gaji < 150:
		keterangan = "Perlu evaluasi finansial"
	}

#### Deskripsi
Program yang dibuat untuk mencari bilangan terbesar dari n buah bilangan bulat yang diinputkan menggunakan perulangan for. 
### 3. bioskop.go

package main

import "fmt"

func main() {
	var t int
	fmt.Print("input t (jumlah transaksi): ")
	fmt.Scan(&t)

	totalPendapatan := 0

	for i := 1; i <= t; i++ {
		var jenisFilm, hari string
		var jumlahTiket int

		fmt.Printf("input transaksi ke-%d (jenis, hari, tiket): ", i)
		fmt.Scan(&jenisFilm, &hari, &jumlahTiket)

		hargaSatuan := 0

		if jenisFilm == "Reguler" {
			if hari == "Senin" || hari == "Selasa" || hari == "Rabu" || hari == "Kamis" {
				hargaSatuan = 40000
			} else {
				hargaSatuan = 50000
			}
		} else if jenisFilm == "3D" {
			if hari == "Senin" || hari == "Selasa" || hari == "Rabu" || hari == "Kamis" {
				hargaSatuan = 60000
			} else {
				hargaSatuan = 70000
			}
		} else if jenisFilm == "IMAX" {
			if hari == "Senin" || hari == "Selasa" || hari == "Rabu" || hari == "Kamis" {
				hargaSatuan = 80000
			} else {
				hargaSatuan = 90000
			}
		}

		totalHarga := hargaSatuan * jumlahTiket

		if jumlahTiket >= 5 {
			diskon := (totalHarga * 10) / 100
			totalHarga = totalHarga - diskon
		}

		fmt.Printf("output total harga transaksi ke-%d: %d\n", i, totalHarga)

		totalPendapatan = totalPendapatan + totalHarga
	}

	fmt.Println("output total pendapatan:", totalPendapatan)
}

### Deskripsi
Program yang membaca sejumlah transaksi, lalu untuk setiap transaksi menghitung total harga berdasarkan jenis film, hari pemutaran, dan jumlah tiket yang dibeli. Jika pembelian dalam satu transaksi ≥ 5 tiket, pelanggan mendapatkan diskon 10% dari total harga transaksi tersebut.

## Unguided

### 1. prima

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

##### Output
![alt text](unguided/grade/output.png)


#### Deskripsi
Program yang dibuat untuk menampilkan seluruh faktor dari sebuah bilangan bulat positif n, kemudian menentukan apakah n merupakan bilangan prima atau bukan.

### 2. ecct

package main

import "fmt"

func main() {
	var penghasilan, pajak float64
	fmt.Scan(&penghasilan)

	if penghasilan <= 50 {
		pajak = 0.05 * penghasilan
	} else if penghasilan > 50 && penghasilan <= 100 {
		pajak = 0.05*50 + (0.1 * (penghasilan - 50))
	} else if penghasilan > 100 && penghasilan <= 200 {
		pajak = 0.05*50 + (0.1 * 50) + (0.15 * (penghasilan - 100))
	} else if penghasilan > 200 {
		pajak = 0.05*50 + (0.1 * 50) + (0.15 * 100) + (0.20 * (penghasilan - 200))
	}

	fmt.Println("Potongan =", pajak)
}


##### Output
![alt text](unguided/pajak/output.png)

#### Deskripsi
Program yang dibuat untuk menentukan apakah setiap mahasiswa baru Tel-U dalam satu kelompok bisa mendapat sertifikat ECCT atau tidak. Syarat mendapat sertifikat ECCT adalah rata-rata nilai tes ECCT minimal 3.33. Selain itu, program juga menghitung berapa banyak mahasiswa yang mendapat sertifikat. 

## Kesimpulan
Berdasarkan praktikum Modul 04, program berhasil diimplementasikan menggunakan struktur seleksi (if-else) dan switch dalam bahasa Go untuk menyelesaikan masalah sekuensi, seperti pengklasifikasian nilai, menu pilihan, dan perhitungan pajak penghasilan berjenjang berdasarkan total penghasilan yang diinputkan pengguna.  

## Referensi
1. ​Laboratorium Praktikum Informatika. (2025). Modul 04: Runtunan dan Sekuensi - Bahasa Pemrograman Go (Golang). Bandung: Fakultas Informatika, Universitas Telkom.