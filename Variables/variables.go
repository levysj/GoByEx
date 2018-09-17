package main

import "fmt"

func main() {

	//var declares 1 or more variables
	var a = "intial"
	fmt.Println(a)

	//You can declare multiple variables at once.
	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	//The := syntax is shorthand for declaring and initializing a variable,
	//e.g. for var f string = "short" in this case.
	f := "short"
	fmt.Println(f)
}
