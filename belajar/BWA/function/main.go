package main

import "fmt"
import "errors"

func main(){
	sentence := printMyResult("Saya Sedang belajar golang ")
	fmt.Println(sentence)

	luas, _ := calculate(10,2)
	fmt.Println("Jumlah Luas : ", luas)
	//fmt.Println("jumlah keliling : ", keliling)

	//quiz ---
	fmt.Println("Hasil QUIZ")
	scores := []int{10, 5, 8, 9, 7}
	total := sum(scores)
	fmt.Println("Jumlah Score : ", total)

	//quiz 2--
	result, err := calculation(10, 2, "+")
	if err != nil {
		fmt.Println("Terjadi Kesalahan")
		fmt.Println(err.Error())
	}

	
	fmt.Println("hasil jumlah : ", result)

	

	
}

func printMyResult(sentence string) string {
	newSentence := sentence + "jadi sress"
	return newSentence
}
//function multiple return
func calculate(panjang int, lebar int) (int, int) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)

	return luas, keliling
}

func sum(scores []int)int{
	var total int
	for _, value := range scores {
		total += value
	}
	return int(total)

}

func calculation(no1, no2 int, operation string) (int, error) {
	var result int
	var errorResult error

	switch operation {
	case "+" :
		result = no1 + no2
	case "-" :
		result = no1 - no2
	case "*" :
		result = no1 * no2
	case "/" :
		result = no1 /no2
	default :
		errorResult = errors.New("hasil tidak diketahui")
	}

	return result, errorResult
}

