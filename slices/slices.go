package slices

import "fmt"

/*ExecSlicesExample is the main function of slices package*/
func ExecSlicesExample() {
	fmt.Println("|-------Beginning of slices package-------|\n")
	mySlice := []int{1, 3, 5, 7, 9, 11}
	fmt.Printf("%T\n", mySlice)
	fmt.Println(mySlice)
	fmt.Println(mySlice[1:3])
	fmt.Println(mySlice[1:])
	fmt.Println(mySlice[:3])
	fmt.Println(mySlice[2])
	fmt.Printf("%c\n", "mySlice"[2])
	fmt.Println("\n|-------End of slices pacakge-------------|\n")
}
