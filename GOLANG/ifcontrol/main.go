package main

import "fmt"

func main() {

	var score int
	fmt.Println("Nilainya Berapa?")
	fmt.Scanf("%d", &score)

	var grade string

	if score <= 50{
		grade = "F"
	}else if score <= 60{
		grade = "E"
	}else if score <= 66{
		grade = "D"
	}else if score <= 82{
		grade = "C"
	}else if score <= 90{
		grade = "B"
	}else{
		grade = "A"
	}

	fmt.Println("Grademu adalah,", grade)
}


/* 
Ada 4 Ifcontrols:
If
If - Else
Else If
Nested If (Kondisi Bersarang)
*/