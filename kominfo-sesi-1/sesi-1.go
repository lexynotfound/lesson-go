package main

import "fmt"

func main() {
	var information string = "Rehan"
	var address string = "Jalan Kemang"
	// tanda garing miring dua untuk komen
	// Menampilkan hello world
	fmt.Println("Hello World", information, address)

	informasi := "adawdaws"
	deskripsi := "Short Description"
	var (
		favoriteColor = "Green"
		favoriteShape = "rectangle"
	)

	//age := 12

	//Ini adalah contoh Type Casting

	// var score int = int32(age)

	var age uint8 = 12
	price := 100
	var score int = price
	fmt.Println("Hello World", information, address, favoriteColor, favoriteShape, "%T")
	//multiple variable declartaions
	var student1, student2, student3 string = "satu", "dua", "tiga"
	var first, second, third int
	first, second, third = 1, 2, 3
	fmt.Println(student1, student2, student3)
	fmt.Print(first, second, third)

	_, _, _, _, _, _ = deskripsi, informasi, information, price, age, score

	// _ = deskripsi syntax tersebut bermaksud untuk membuat sebuah variable agar bisa di jalankan dan jika lebih dari satu harus melakukan seperti ini _,_=deskripsi,test
}
