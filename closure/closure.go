package closure

import "fmt"

func getClosure() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

/*ExecClosureExample is the main function of closure package*/
func ExecClosureExample() {
	fmt.Println("|-------Beginning of closure package-------|\n")
	increment := getClosure()
	fmt.Println(increment())
	_, a := increment(), increment()
	fmt.Println(a)
	fmt.Println("\n|-------End of closure pacakge-------------|\n")
}
