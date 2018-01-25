package main

import (
	"fmt"
	"math/rand"
	"math"
)

// example taking parameters type int and returning an int
func add(x int, y int) int {
	return x + y
}

// When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
func add_print(x, y int) {
	var r = add(x, y)
	fmt.Println(r)
}

// Go's return values may be named. If so, they are treated as variables defined at the top of the function.
func add_named(x, y int)(r int) {
	r = add(x, y)
	return
}

func for_example() int {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	return sum
}

func if_example(x int) int{
	if x > 5 {
		return 10
	}
	return 0
}

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	add_print(7, 3)
	fmt.Println(add_named(10,11))

	//multiple ways to define variables
	var a1 = 1
	a2 := 2
	var a3 int
	a3 = 3
	fmt.Println(a1, a2, a3)

	// constant declaration example
	const Pi = 3.14

	for_result := for_example()
	if_result := if_example(6)

	println(for_result, if_result)

}
