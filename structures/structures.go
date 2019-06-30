package main

import "fmt"

func main() {
	fmt.Println(Person{"vignesh", 27, ""})
	fmt.Println(Person{name: "vignesh", age: 27})
	objP := Person{"viki", 25, "chennai"}
	fmt.Println(objP.address)
	fmt.Println("location is", objP.location())
	objs := make(map[int]Person)
	objs[0] = objP
	fmt.Println(objs)
}

//Person struct
type Person struct {
	name    string
	age     int
	address string
}

func (p *Person) location() string {
	return p.address
}
