package main

import "fmt"

func main() {
	var umur int8
	var suhu float32

	//umur = 10
	//suhu = 36.3

	fmt.Scan(&umur)
	fmt.Scan(&suhu)

	fmt.Println("umur : ", umur)
	fmt.Println("suhu : ", suhu)
	fmt.Println("alamat umur ", &umur)
	fmt.Println("alamat suhu ", &suhu)
}
