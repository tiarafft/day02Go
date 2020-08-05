package main

import (
	"fmt"
	"math"
	//"strconv"
	
)

type Point struct {
	akar float64
}


type Count struct {
	No1, No2 int
}


func main() {
	
	var no , No1, No2 int
	var akar float64
	

	fmt.Println("Kalkulator")
	fmt.Println("1. Kali")
	fmt.Println("2. Bagi")
	fmt.Println("3. Tambah ")
	fmt.Println("4. Kurang ")
	fmt.Println("5. Akar")
	fmt.Println("6. Pangkat")
	fmt.Println("7. Luas Persegi")
	fmt.Println("8. Luas Lingkaran")
	fmt.Println("9. Volume Tabung")
	fmt.Println("10. Volume Balok ")
	fmt.Println("11. Volume Prisma ") 
	fmt.Print("Masukkan metode perhitungan yang diinginkan :   ")
	fmt.Scan(&no)

  
	counting := Count{No1, No2}
	v := Point{akar}
	

	switch no {
	case 1:
		fmt.Print("Masukkan angka1 = ")
		fmt.Scan(&No1)
		fmt.Print("Masukkan angka2 = ")
		fmt.Scan(&No2)
		fmt.Println("Hasil:", counting.Perkalian())

	case 2:
		fmt.Print("Masukkan angka1 = ")
		fmt.Scan(&No1)
		fmt.Print("Masukkan angka2 = ")
		fmt.Scan(&No2)
		fmt.Println("Hasil:", counting.Pembagian())

	case 3:
		fmt.Print("Masukkan angka1 = ")
		fmt.Scan(&No1)
		fmt.Print("Masukkan angka2 = ")
		fmt.Scan(&No2)
		fmt.Println("Hasil:", counting.Pertambahan())

	case 4:
		fmt.Print("Masukkan angka1 = ")
		fmt.Scan(&No1)
		fmt.Print("Masukkan angka2 = ")
		fmt.Scan(&No2)
		fmt.Println("Hasil:", counting.Pengurangan())

	case 5:
		fmt.Print("Masukkan akar = ")
		fmt.Scan(&akar)
		fmt.Println("Hasil:", v.Abs())
	case 6:
	 // fmt.Println(angka1, " pangkat ", angka2, " = " ,pangkat)
	case 7:
	 // fmt.Println("Luas persegi =  " ,persegi)
	case 8:
	 // fmt.Println("Luas lingkaran =  " ,lingkaran)
	case 9:
	 // fmt.Println("Volume tabung = " ,tabung)
	case 10:
	 // fmt.Println("Volum balok = " ,balok)
	case 11:
	 // fmt.Println("Volume prisma = " ,prisma)
  
	}
  

	
	

	
}

func (h Count) Pertambahan() int {
	return h.No1 + h.No2
}

func (h Count) Pengurangan() int {
	return h.No1 - h.No2 
}

func (h Count) Perkalian() int {
	return h.No1 * h.No2 
}

func (h Count) Pembagian() int {
	return h.No1 / h.No2
}

func (v Point) Abs() float64 {
	return math.Sqrt(v.akar)
}
