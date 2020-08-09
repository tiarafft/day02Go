package main

//acuan seleksi kondisi adalah nilai pertipe bool
//2 keyword : if else dan switch
//go tidak mendukung seleksi kondisi menggunakan tenary seperti var data = (isExist?"ada :"tidak ada)

 import "fmt"
 
func main() {
	var point = 8

	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus. nilai anda %d\n", point)
	}
}