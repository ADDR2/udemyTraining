package ifGo

import "fmt"

/*ExecIfExample is the main function of ifGo package*/
func ExecIfExample() {
	fmt.Println("|-------Beginning of ifGo package-------|\n")
	if foo := "Chocolate"; true {
		fmt.Println(foo)
	}
	fmt.Println("\n|-------End of ifGo pacakge-------------|\n")
}
