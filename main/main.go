package main

import "fmt"
import "github.com/ADDR2/udemyTraining/stringutil"

func another() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	increment := another()
	fmt.Println(a)
	fmt.Println(stringutil.Reverse("Hello World"))
	fmt.Println(increment())
	fmt.Println(increment())
}

var a = 3
