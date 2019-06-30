package main

import (
	"errors"
	"fmt"
)

func main() {
	name, err := f1("viki")
	fmt.Println(name)
	fmt.Println(err)
}
func f1(name string) (string, error) {
	return "vignesh", errors.New("error")
}
