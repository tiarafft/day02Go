package main

import (
	"net/http"
	"fmt"
	"encoding/json"
)
type Hobby struct {
	Name string
}

type User struct {
	FullName string `json:"Name"`
	Age      int
	Hobbies  []Hobby
}
func main(){
	http.HandleFunc("/getIDParkir", func(w http.ResponseWriter, r *http.Request) {
		// get new parking id from server
		fmt.Fprintln(w, "YourParkingID is: "+ ResponseFromServer)
		// fmt.Fprintln(w, data.FullName)
	})

	// example for encode from struct to string
	var hobbyStruct = []Hobby{{"Sleeping"}, {"Listening Music"}}
	var dataStruct = User{"Admin Istrator", 20, hobbyStruct}
	jsonDt, err := json.Marshal(dataStruct)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	http.HandleFunc("/marshal", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "POST" {
			// fmt.Fprintln(w, string(jsonDt))
			w.Write(jsonDt)
			return
		}

		http.Error(w, "", http.StatusBadRequest)

	})

	fmt.Println("running server...")
	http.ListenAndServe(":8888", nil)
}

/*
	Buat service parkir yang sudah dibuat menjadi 2 service terpisah, yaitu service yang berperan sebagai client dan sebagai server.
	- service client: hanya mengirim data dan menerima data
	- service server: yang menjalankan logic/menghitung, dsb
*/