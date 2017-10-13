package memory

import "fmt"

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
	tb = 1 << (iota * 10)
)

const metersToYards = 1.09361

/*ExecMemoryExample is the main function of memory package*/
func ExecMemoryExample() {
	fmt.Println("|-------Beginning of memory package-------|\n")
	fmt.Println("Memory address ", &a)
	fmt.Printf("In decimal %d\n", &a)
	fmt.Println(kb, mb, gb, tb)

	var meters float64
	fmt.Println("Enter meters swam: ")
	fmt.Scan(&meters)
	fmt.Println(meters, " = ", meters*metersToYards, " yards.")
	fmt.Println("\n|-------End of memory pacakge-------------|\n")
}

var a = 3
