package main

import "fmt"

func main() {

	/*
	// Pendefinisian Array dengan mendefinisikan tipe datanya dahulu
	var buah[5] string
	buah[0] = "Apel"
	buah[1] = "Jeruk"
	buah[2] = "Mangga"
	buah[3] = "Pepaya"
	buah[4] = "Pisang"

	// Pendefinisian Array secara langsung
	sayur := [5]string{"Bayam", "Kangkung", "Kubis", "Wortel", "Kacang"}

	// Pendefinisiaan Array secara langsung dengan menentukan panjangnya (2)
	hewan := [3]string{
		"Kucing",
		"Anjing",
		"Kelinci"}
	*/

	// Pendefinisiaan Array secara langsung tanpa menentukan panjangnya (...)
	bahasa :=[...]string{
		"Java",
		"C",
		"C++",
		"Python",
		"Ruby",
		"PHP",
		"Go",
		"JavaScript",
		"Rust"}

	for index, lang := range bahasa{
		fmt.Println("Indeks Ke->", index, "Bahasa Pemrograman", lang)
	}

}