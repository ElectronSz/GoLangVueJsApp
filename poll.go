package main

//imports
import (
	"fmt"
	"math"
)

//aa two numbers
func add(x int, y int) int {
	return x + y
}

//substract numbers
func sub(x int, y int) int {
return x - y
}

//main function
func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(add(50,55))
	fmt.Println(sub(50,20))
}
