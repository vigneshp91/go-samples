package main

import "fmt"

func main() {
	fmt.Printf("hello, world\n")
	a := 99
	fmt.Println(a)
	var b, c, d int = 4, 5, 6
	fmt.Println(c)
	fmt.Println(b)
	fmt.Println(d)
	const con = "hello constant"
	fmt.Println(con)
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	arr[3] = 55
	fmt.Println("array value in 3 is", arr[3])

	s := make([]string, 5)
	mapvar := make(map[int]string)
	s[2] = "hello"
	mapvar[1] = "v"
	mapvar[2] = "v1"
	fmt.Println(mapvar, len(mapvar))
	fmt.Println(s)
}
