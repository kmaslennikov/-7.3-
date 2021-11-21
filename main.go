package main

import (
	"fmt"
)

func main() {
	fmt.Println("Введите высоту ёлочки:")
	var height int
	fmt.Scan(&height)
	var countStars, countSpace int
	var line string
	length := 1 + 2*(height-1) // количество символов в каждой строке
	fmt.Println(length)

	for i := 0; i < height; i++ {
    line = ""
		countStars = 1 + 2*i
		countSpace = (length - countStars) / 2 //количество пробелов с каждой стороны
		for s := 0; s < countSpace; s++ {
			line += " "
		}
		for j := 0; j < countStars; j++ {
			line += "*"
		}

		for s := 0; s <= countSpace; s++ {
			line += " "
		}
    println(line)
	}
	
}
