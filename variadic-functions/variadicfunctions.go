package main
import "fmt"
func main(){
	fmt.Println(sum(1,2,3,4,5,6))
}
func sum(num ...int) int{
	var sum int
	for _,n := range num{
		sum+=n
	}
	return sum
}