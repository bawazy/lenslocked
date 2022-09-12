package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Age  int
}

func main() {
	tpl, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "sadik",
		Age:  27,
	}

	err = tpl.Execute(os.Stdout, user)

	if err != nil {
		panic(err)
	}

}
