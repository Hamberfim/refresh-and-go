package main

import (
	"fmt"
	"math"
	"math/rand"
)

// a function to use within main (each var has to declare it's type and the func return type)
func add(x int, y int) int {
	return x + y
}

// shortened type declaration
func prod(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("A random number is", rand.Intn(10))
	fmt.Println("Here is PI:", math.Pi)

	fmt.Printf("Now you have %g problems. \n", math.Sqrt(7)) // %g for floating-point

	fmt.Printf("The sum of %d + %d = %d.", 24, 13, add(42, 13)) // %d for int/
}
