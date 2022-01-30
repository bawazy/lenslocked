package main

import (
	"html/template"
	"net/http"
)

type User struct {
	Name string
	Bio  string
	Age  float64
	Data cust_details
}
type cust_details struct {
	Address string
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "",
		Bio:  "testing",
		Age:  6,
		Data: cust_details{Address: `<script>console.log(15 Lamido road yola south LGA)</script> </h1>`},
	}

	err = t.Execute(w, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}
