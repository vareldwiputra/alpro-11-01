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
