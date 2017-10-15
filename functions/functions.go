package functions

import "fmt"

func get(fName, lName string) {
	fmt.Println(fName, lName)
}

func returnGet(fName, lName string) string {
	return fmt.Sprint(fName, lName)
}

func returnGetNamed(fName, lName string) (text string) {
	text = fmt.Sprint(fName, lName) // Ya esta declarado text
	return
}

func multipleGet(fName, lName string) (string, string) {
	return fmt.Sprint(fName, lName), fmt.Sprint(lName, fName)
}

func variadic(sl ...float64) float64 {
	fmt.Println(sl)
	fmt.Printf("%T\n", sl)

	var total float64
	for index, value := range sl {
		total += value
		fmt.Println(index)
	}
	return total / float64(len(sl))
}

type integer int

func (receiver integer) test() integer {
	return receiver & 1
}

func visit(numbers []int, callback func(int)) {
	for _, value := range numbers {
		callback(value)
	}
}

func hello() {
	fmt.Print("Hello ")
}

func world() {
	fmt.Println("World")
}

/*ExecFunctionsExample is the main function of functions package*/
func ExecFunctionsExample() {
	fmt.Println("|-------Beginning of functions package-------|\n")
	get("Amaro", "Duarte")

	fmt.Println(returnGet("Amaro ", "Duarte"))

	fmt.Println(returnGetNamed("Amaro ", "Duarte"))

	result1, result2 := multipleGet("Amaro ", "Duarte ")
	fmt.Println(result1, result2)

	fmt.Println(multipleGet("Amaro ", "Duarte "))

	fmt.Println(variadic(43, 56, 87, 12, 45, 57))

	var number integer = 33
	fmt.Println(number.test())

	data := []float64{43, 56, 87, 12, 45, 57}
	fmt.Println(variadic(data...))

	visit([]int{1, 2, 3, 4}, func(number int) {
		fmt.Println(number)
	})

	defer fmt.Println("\n|-------End of functions pacakge-------------|\n")
	defer world()
	defer hello()

	m := make([]string, 1, 25)
	fmt.Println(m)
	func() {
		changeMe(m)
	}()
	fmt.Println(m)
}

func changeMe(z []string) {
	z[0] = "Todd"
	fmt.Println(z)
	fmt.Println(len(z))
}
