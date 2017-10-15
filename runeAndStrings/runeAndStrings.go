package runeAndStrings

import "fmt"

/*ExecRuneAndStringsExample is the main function of runeAndStrings package*/
func ExecRuneAndStringsExample() {
	fmt.Println("|-------Beginning of runeAndStrings package-------|\n")
	fmt.Println([]byte("Hello"))
	for i := 50; i <= 140; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}
	fmt.Println("\n|-------End of runeAndStrings pacakge-------------|\n")
}
