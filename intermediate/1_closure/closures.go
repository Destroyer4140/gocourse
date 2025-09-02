package main

import "fmt"

/*
Closure is a property of function under which callback function and varibale enclosed under that callback
function whether it's global or local method are/is enclosed till the life of function.
*/
func main() {

	sequence1 := adder()
	fmt.Println("sequence11", sequence1())
	fmt.Println("sequence12", sequence1())
	fmt.Println("sequence13", sequence1())
	fmt.Println("sequence14", sequence1())

	sequence2 := adder()
	fmt.Println("sequence21", sequence2())
	fmt.Println("sequence22", sequence2())
	fmt.Println("sequence23", sequence2())
	fmt.Println("sequence24", sequence2())

	fmt.Println("sequence15", sequence1())
	fmt.Println("sequence16", sequence1())

	fmt.Println("sequence25", sequence2())
	fmt.Println("sequence26", sequence2())

	subtractor := func() func(int) int {
		countdown := 99
		return func(x int) int {
			countdown -= x
			return countdown
		}
	}()

	fmt.Println("Using closure subtractor by 10 ->", subtractor(10))
}

func adder() func() int {
	fmt.Println("Initialisation of closure")
	i := 0
	fmt.Println("Previous value if i -> ", i)
	return func() int {
		i++
		fmt.Println("Incrased value of i -> ", i)
		return i
	}
}
