# <h1 align="center">Tugas Pendahuluan Modul [5] - [Perulangan For Loop]</h1>
<p align="center">[Varel Dwi Putra] - [109092600024]</p>

### 1. 

### 2. 

### 3. Ganjil
package main

import "fmt"

func main() {
	intNum := 5
	intOther := 10
	//var sngNum float64 = -3

	if intNum >= 5 && intOther < 11 {
		fmt.Println("Beep")
	}

}


##### Output
![alt text](tp/evaluasi/output.png)

1. false
intNum > 5
2. true
intNum >= 5 && intOther < 11
3. true
sngNum != -1 || intOther < 0
4. true
!(intNum > 3) || intNum <= 5
5. false
!(intOther >= intNum)
6. false
0 - sngNum > 0
7. true
4 / 2 == intOther / intNum
8. true
intOther % 2 == 0
9. true
intOther + 2 * intNum != 30 || !(sngNum > 0)
10. true
intOther > 0 && intNum > 0 || sngNum > 0
11. true
sngNum > 0 || (intNum >= 0 && -1 * intOther == -10)
12. true
intNum == 5
13. true
intNum > 0 || (sngNum <= 0 && intOther == 13)
14. true
!(!(!(!(intNum > 0))))

#### Deskripsi
Program yang menggunakan perulangan for untuk menampilkan semua bilangan ganjil dari 1 hingga n, kemudian menampilkan jumlah seluruh bilangan ganjil tersebut. 

### 4. Perkalian

package main

import "fmt"

func main() {
	x := 10
	y := 5
	z := 15
	result := 0

	if x > 5 {
		if y < 10 {
			result = x + y
		} else {
			result = x - y
		}
	}

	if z > 10 && x == 10 {
		result += z
	} else {
		result = z - x
	}

	if x == 10 || y > 10 {
		result += 5
	} else if y == 5 && z > 10 {
		result -= 5
	} else {
		result *= 2
	}

	if !(x < 15 && y < 10) {
		result += 10
	} else {
		_ = result - 10
	}

	fmt.Println("Nilai akhir result:", result)
}


##### Output
![alt text](tp/tracing/output.png)

​1. Nilai akhir variabel result: 35

​2. Output program: Nilai akhir result: 35

​3. Penjelasan Singkat (Tracing):

​ -Kondisi 1: 10 > 5 (true) dan 5 < 10 (true) \rightarrow result = 10 + 5 = 15.

​ -Kondisi 2: 15 > 10 \text{ \&\& } 10 == 10 (true) \rightarrow result = 15 + 15 = 30.

​ -Kondisi 3: 10 == 10 \text{ \vert{}\vert{} } 5 > 10 (true) \rightarrow result = 30 + 5 = 35.

​ -Kondisi 4: !(10 < 15 && 5 < 10) \rightarrow !(true) = false. Masuk ke else yang berisi result - 10. Karena tidak ada tanda sama dengan (-=), nilainya tidak tersimpan dan result tetap 35.

#### Deskripsi
 Program yang menggunakan perulangan for untuk menampilkan tabel perkalian dari sebuah bilangan a, mulai dari a × 1 hingga a × b. 

## Kesimpulan
Tugas pendahuluan praktikum Algoritma Pemrograman yang berfokus pada pemahaman dan penerapan struktur kontrol keputusan (if-else dan switch-case) serta logika ekspresi boolean dalam bahasa pemrograman Go.
