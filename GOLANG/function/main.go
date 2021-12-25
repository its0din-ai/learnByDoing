package main

import "fmt"

func main() {
	// Fungsi didefinisikan dengan func
	// Fungsi publik menggunakan awalan huruf kapital
	// Fungsi privat menggunakan awalan huruf kecil

	jari := fmt.Scanln("%r") //lanjut nanti

	hasil := calcul(jari)
	fmt.Println(hasil)

}

func calcul(jari float64) float64 {
	var luas float64
	luas = 3.14 * jari * jari
	fmt.Println("Luas Lingkaran : ", luas)
	return luas
}