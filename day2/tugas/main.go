package main

import (
	"fmt"
	"time"

	guuid "github.com/google/uuid"
)


var (
	duration int =2
)

func main() {
	menu()
}

func menu() {
	input :=1
	x := true
	
	for x {

	
	fmt.Println("Selamat Datang")
	fmt.Println("1. Parkir Masuk")
	fmt.Println("2. Parkir Keluar")

	fmt.Println("pilih menu : ")
	fmt.Scan(&input)


	switch input {
	case 1:
		masuk()
	
	case 2:
		keluar()
	
	default : x = false
	}

	}
}

func masuk() {
	
	id := guuid.New()
	fmt.Printf("ID Parkir:   %s\n", id.String())

	time := time.Now()
	fmt.Println("Tanggal masuk: ", time.Day(), "/", time.Month(), "/", time.Year())
	time.Hour()
	if time.Hour() <12 {
		fmt.Println("Waktu masuk:", time.Hour(), "-", time.Minute(), "-", time.Second(), "AM")
	}else{
		fmt.Println("Waktu masuk:", time.Hour(), "-", time.Minute(), "-", time.Second(), "PM")
	}
	

}

func keluar() {
	var kendaraan, plat string

	
	fmt.Println("Masukan Jenis Kendaraan : ")
	fmt.Scan(&kendaraan)
	fmt.Println("Masukan Plat Nomor : ")
	fmt.Scan(&plat)
	time.Sleep(2 * time.Second)
	fmt.Printf("duration for %v seconds\n", duration)



} 


