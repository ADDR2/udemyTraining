package arrays

import "fmt"

/*ExecArraysExample is the main function of arrays package*/
func ExecArraysExample() {
	fmt.Println("|-------Beginning of arrays package-------|\n")
	var arr [20][20]string
	for i := 0; i < 20; i++ {
		for j := 0; j < 20; j++ {
			fmt.Println("Before ---->", arr[i][j])
			arr[i][j] = string(i * j)
			fmt.Println("After ---->", arr[i][j])
		}
	}

	arr2 := [...][1]string{{"a"}, {"b"}}
	fmt.Println(arr2)
	fmt.Println("\n|-------End of arrays pacakge-------------|\n")
}
