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
	var isStar bool
	length := 1 + 2*(height-1) // количество символов в каждой строке
	fmt.Println(length)

	for i := 0; i < height; i++ {
		line = ""
		countStars = 1 + 2*i
		countSpace = (length - countStars) / 2 //количество пробелов с каждой стороны
		for j := 0; j < length; j++ {
			isStar = j >= countSpace && j < countSpace+countStars
			if isStar {
				line += "*"
			} else {
				line += " "
			}
		}
		println(line)
	}

}
