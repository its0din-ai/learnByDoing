package main

import "fmt"

func main() {
	/*
		Slice adalah tipe data yang bersifat Reference / referensi
		Semua elemen yang dipanggil adalah referensinya (berada di memori yang sama)
		Slice bisa diubah-ubah, namun elemen yang dipanggil tetap berada di memori yang sama
		Berbeda dengan Array, yang menduplikat / copy elemen yang dipanggil
	*/

	// Pendefinisian Slice tanpa mendefinisikan jumlah elemen, berbeda dengan Array
	var koin = []string{"Bitcoin", "Ethereum", "USD Tether", "Holo Token"}

	first := koin[0]				// Pemanggilan Slice ke-0
	tiga := koin[0:3]			// Pemanggilan Slice ke-0 sampai elemen sebelum 4

	fmt.Println(first)
	fmt.Println(tiga)

	koin = append(koin, "Ripple")	// Menambahkan elemen baru pada Slice
	koin = append(koin, "Litecoin")	// Menambahkan elemen baru pada Slice

	fmt.Println("===================================")
	fmt.Println("Koin yang ada di dalam Slice sudah berubah")

	for _, Koin := range koin{
		fmt.Println(Koin)
	}
}