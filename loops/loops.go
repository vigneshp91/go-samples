package main

import "fmt"

func main() {

	s := make([]string, 5)
	mapvar := make(map[int]string)
	s[2] = "hello"
	mapvar[1] = "v"
	mapvar[2] = "v1"
	for _, num := range s {
		fmt.Println(num)
	}
	for k, v := range mapvar {
		fmt.Printf("%d -> %s", k, v)
	}
}
