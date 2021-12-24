package main

import "fmt"

func main() {
	/*
		Map adalah Tipe Data Assosiatif yang bersifat Key:Value
		Key Bersifat Unik, karena untuk pendefinisian Value
		
		Map hampir mirip dengan Slice, tetapi indeks yang dimiliki
		oleh Map dapat didefinisikan sebagai Key
	*/
	
	oktan := map[int]string{
		98: "Pertamax Turbo",
		95: "Pertamax Plus",
		92: "Pertamax",
		90: "Pertalite",
		88: "Premium",
	}

	value, isavailable := oktan[88]
	fmt.Println(isavailable)
	fmt.Println(value)

	delete(oktan, 88)
	fmt.Println("===================================")
	fmt.Println("Value dengan Key 88 sudah dihapus")
	fmt.Println("===================================")
	for key, value := range oktan {
		fmt.Println("Key : ", key, " Value : ", value)
	}
}