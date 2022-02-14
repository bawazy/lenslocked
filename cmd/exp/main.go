package main

import (
	"errors"
	"fmt"
)

func main() {
	err := B()
	// fmt.Println(err)
	if errors.Is(err, ErrNotFound) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	//TODO: Determine if the 'err' variable is an 'ErrNotFound'
}

// it is common for packages like databse/sql to return an error
//that is predefined like this one.

var ErrNotFound = errors.New("not Found")

func A() error {
	return ErrNotFound
}

func B() error {
	err := A()
	return fmt.Errorf("b: %w", err)
}
