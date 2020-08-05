package main

import "fmt"
import "math"

type hitung interface {
	luas() float64
	keliling() float64
}

type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}
func (p persegi) keliling() float64 {
	return p.sisi = 4
}