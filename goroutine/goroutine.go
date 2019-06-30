package main

import (
	"fmt"
)

func main() {
	fmt.Println("blocking call")
	loop(5)
	fmt.Println("non blocking call")
	go loop(5)
	fmt.Println("non blocking call")
	go func(msg string) {
		fmt.Println(msg)
	}("going")
	fmt.Scanln()
	//time.Sleep(5 * time.Second)
	fmt.Println("Done!")
}
func loop(cnt int) {
	for i := 0; i < cnt; i++ {
		fmt.Println("running", i)
	}
}
