package main

import "fmt"

func main() {
	nextInt := getFun()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
}
func getFun() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
