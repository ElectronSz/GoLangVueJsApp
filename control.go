package main

import (
	"fmt"
	"math"
)

func loops(num int) int {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		//fmt.Println(i)
	}
	return sum
}

func loopsV1(num int) int {
	sum := 1
	for sum < 10 {
		sum += sum
	}
	return sum
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(loops(10))
	fmt.Println(loopsV1(10))
	fmt.Println(sqrt(2), sqrt(-4))
}
