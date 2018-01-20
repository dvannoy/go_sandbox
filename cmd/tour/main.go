package main

import (
	"fmt"
	"math/rand"
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

}
