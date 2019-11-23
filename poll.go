package main

//imports
import (
	"fmt"
	"math"
	"math/cmplx"
)

//vars
var c, python, java bool = true, true, false
var (
	mAXInt uint64     = 1<<64 - 1
	ToBe   bool       = false
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

//aa two numbers
func add(x int, y int) int {
	return x + y
}

//substract numbers
func sub(x, y int) int {
	return x - y
}

//swap
func swap(x, y string) (string, string) {
	return y, x
}

//naked returns
func split(sum int) (x, y int) {
	x = sum + 10
	y = sum * x
	return
}

//main function
func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(add(50, 55))
	fmt.Println(sub(50, 20))

	a, b := swap("Hello", "World")
	fmt.Println(a, b)

	fmt.Println(split(10))

	var i int
	j := 10
	fmt.Println(i, c, python, java, j)

	fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Printf("Type: %T Value: %v\n", mAXInt, mAXInt)

	var k int
	var f float64
	var l bool
	var s string
	fmt.Printf("%v %v %v %q\n", k, f, l, s)

	var x, y int = 3, 4
	var g float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, g, z)

}
