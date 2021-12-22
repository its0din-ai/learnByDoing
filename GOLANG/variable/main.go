// Belajar GOLANG #0

package main

import(
	"fmt"		//Package buat Print
)

func main() {

	/* 

		variabel = string, int, float64, boolean

		string : "ini adalah string", 'ini juga string'
		int : 1,2,3,99,424,512
		float64 : 1.2, 3.14, 0.1, 6.66
		boolean : true, false
	
	*/
	
	var nama string = "Odin"	//untuk = hanya digunakan sbg assignment, harus di deklarasikan -> 'var nama string' adalah deklarasi variabel
	nick := "its-odin"			//untuk := bisa digunakan sbg deklarasi & assignment
	kode := 666					//ini integer

	fmt.Println(nama)
	fmt.Println(nick)
	fmt.Println(kode)

	kode = 404					//untuk mengubah nilai, gunakan = saja
	fmt.Println(kode)
	
	var nim string
	fmt.Println(nim)			//nilai variabel nim kosong, defaultValue=(string: kosong), (int: 0), (float64: 0.0), (boolean: false)
	nim = "L200210056"			//assignment, untuk mengisi nilai variabel
	fmt.Println(nim)

	bisa := true
	fmt.Println(bisa)			//tipe data boolean

}