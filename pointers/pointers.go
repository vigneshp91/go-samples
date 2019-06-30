package main

import "fmt"

func main() {
	i := 10
	incValue(i)
	fmt.Println(i)
	incPts(&i)
	fmt.Println(i)

}
func incValue(i int) {
	i++
}
func incPts(i *int) {
	*i++
}
