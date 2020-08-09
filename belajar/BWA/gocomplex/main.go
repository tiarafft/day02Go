package main

import "fmt"

func main(){
 //belajar array
/**
 languages := [...]string{"go", "ruby", "java", "c#", "goblog"}
	
	
	for index, lang := range languages {
		fmt.Println("Index :", index, "Languages : ", lang)

	}
//belajar slice
	var gamingConsole []string
	gamingConsole = append(gamingConsole, "Playstation4")
	gamingConsole = append(gamingConsole, "Nintendo")
	gamingConsole = append(gamingConsole, "Xbox")

	for _, console := range gamingConsole {
		fmt.Println(console)
	}
//belajar map
	myMap := map[string]string{"ruby": "isbeautiful", "go": "issuperfast", "java " : "hype"}


	for key, value :=range myMap {
		fmt.Println("Key :", key, "value :", value)
	}
//map delete
	fmt.Println("=======delete=======")
	delete(myMap, "ruby")
	
	for key, value :=range myMap {
		fmt.Println("Key :", key, "value :", value)
	}

	//mengecek apakah datanya ada
	 value, isAvailable := myMap["Python"]
	 fmt.Println(isAvailable)
	 fmt.Println(value)

	//SLICE OF MAP
	students := []map[string]string{
		{"name" : "agung", "score" : "A"},
		{"name" : "tiara", "score" : "A"},
		{"name" : "fauzi", "score" : "B"},
	}
	for _, students := range students {
		fmt.Println(students["name"], "-", students["score"])
	} **/

	//quiz
	//hitung rata2
	Scores := []int{100, 80, 75, 92, 70, 93, 88, 67}
	var goodScores []int
	for _, score := range Scores {
		if score >= 90 {
			goodScores = append(goodScores, score)
		}
	}
	fmt.Println("rata- ratanya adalah : ", averageScores(Scores))
	fmt.Println("Score Terbaik adalah: ", goodScores)

	
}
	
	func averageScores(Scores []int) float64 {
		var total int 
		for _, value := range Scores {
			total += value
		}
		return float64(total) / float64(len(Scores))
	}



