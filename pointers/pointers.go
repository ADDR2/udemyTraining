package pointers

import "fmt"

func zero(x *int) {
	*x = 0
}

/*ExecPointersExample is the main function of pointers package*/
func ExecPointersExample() {
	fmt.Println("|-------Beginning of pointers package-------|\n")
	a := 43
	fmt.Println(a)
	fmt.Printf("%p\n", &a) // is the same as fmt.Println(&a)

	var b = &a // is the same as var b *int = &a
	fmt.Println(b)
	fmt.Println(*b)

	*b = 30
	fmt.Println(a)

	x := 5
	zero(&x)
	fmt.Println(x)
	fmt.Println("\n|-------End of pointers pacakge-------------|\n")
}
