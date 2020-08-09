package main

import (
	"fmt"
	"modules/calculations"
)

func main(){
	fmt.Println("Hello Indonesia")

	result := calculations.Add(10, 9)

	fmt.Println(result)


//quiz-----------------
	 result2 := calculations.Sub(14, 2)
	 fmt.Println(result2)

	 result3 := calculations.Multiple(12,2)
	 fmt.Println(result3)

//quiz2 perulangan cetak angka genap, cetak huruf vocal
title := "golang is the best language"

	for g := 0; g <= 20; g++ {
		if g%2 == 0 {
			fmt.Println("index : ", g)
		}
		}
			
	for g, letter := range title {
		letterString := string(letter)

		if letterString == "a" || letterString == "i" || letterString =="u" || letterString == "e" || letterString == "o" {
			fmt.Println("Letter : ", string(letter), "Index : ", g) 
		}
	}

	}

	




