package main

import "fmt"

func main() {
	x, y := addSub(5, 3)
	fmt.Println(x)
	fmt.Println(y)
}
func addSub(a int, b int) (int, int) {
	val := a + b
	val1 := a - b
	return val, val1
}
