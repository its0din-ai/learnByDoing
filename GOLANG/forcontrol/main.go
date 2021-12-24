package main

import "fmt"

func main() {

	// for i:=1; i<5; i++{
	// 	fmt.Println("Saya Belajar GO")
	// }

	// z := 0
	// for z < 5{
	// 	fmt.Println("Saya Belajar Go")
	// 	z++
	// }


	title := "Golang the best language"

	// for index, letter := range title {
	// 	fmt.Println("index: ", index, " letter :", string(letter))
	// }

	for _, letter := range title {
		fmt.Println(" letter :", string(letter))
	}

}