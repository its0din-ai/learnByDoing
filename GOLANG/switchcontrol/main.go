package main

import "fmt"
import "strings"

func main() {

	var Rasa string
	fmt.Scanf("%s", &Rasa)
	rasa := strings.ToLower(Rasa)

	switch rasa {
	case "manis":
		fmt.Println("Saya Juga suka Rasa Manis")
	case "asin":
		fmt.Println("Saya Kurang Suka, kalau berlebihan")
	case "pahit":
		fmt.Println("Rasa apaan ini :(")
	default:
		fmt.Println("Saya Baru tau rasa itu")
	}
}