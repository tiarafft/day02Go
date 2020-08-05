package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gitlab.com/martha.sutopo/restapi/model"
)

var c = model.Customers {
	model.Customer{Id: 1, Name : "Martha S"},
	model.Customer{Id: 2, Name: "Tiara F"},
}

func cust (w http.ResponseWritter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method =="POST" {
		var result, err =json.Marshal(c)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Error(w, "", http.StatusBadRequest)
	}

	
}
func main() {
	fmt.Println(c)

	http.HandleFunc("/customer", cust)
	http.ListenAndServe(":8080", nil)
}