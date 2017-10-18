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
	fmt.Println(mySlice[:])
	fmt.Println(mySlice[2])
	fmt.Printf("%c\n", "mySlice"[2])

	mySlice = append(mySlice[:2], mySlice[3:]...) // Deleting position 2
	fmt.Println(mySlice)
	mySlice[2]++
	fmt.Println(mySlice[2])

	newSlice := make([]int, 0, 5)

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	for i := 0; i < 80; i++ {
		newSlice = append(newSlice, i)
		fmt.Println("Len:", len(newSlice), "Capacity:", cap(newSlice), "Value:", newSlice[i])
	}

	for i, entry := range newSlice {
		fmt.Println(i, entry)
	}

	newSlice = append(newSlice, mySlice...)
	fmt.Println(newSlice)

	otherSlice := make([]string, 3)

	otherSlice[0] = "a"
	otherSlice[1] = "b"
	otherSlice[2] = "c"
	otherSlice = append(otherSlice, "d") // Using append because the length is 3
	fmt.Println(otherSlice)

	var class = make([][]string, 0, 10)

	student1 := make([]string, 0, 3)
	student1 = append(student1, "My")
	student1 = append(student1, "name")
	student1 = append(student1, "is")

	class = append(class, student1)

	student2 := make([]string, 0, 3)
	student2 = append(student2, "Amaro")
	student2 = append(student2, "David")
	student2 = append(student2, "Duarte")

	class = append(class, student2)

	fmt.Println(class)
	var slice []string
	fmt.Println(slice == nil)
	theSlice := []string{}
	fmt.Println(theSlice == nil)
	fmt.Println("\n|-------End of slices pacakge-------------|\n")
}
