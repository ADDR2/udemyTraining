package loops

import "fmt"

/*ExecForLoopsExample is the main function of loops package*/
func ExecForLoopsExample() {
	fmt.Println("|-------Beginning of loops package-------|\n")
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}

	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	i = 0
	for {
		if i >= 10 {
			break
		}
		fmt.Println(i)
		i++
	}

	i = 0
	for {
		i++
		if i&1 != 1 {
			continue
		}
		if i >= 50 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("\n|-------End of loops pacakge-------------|\n")
}
