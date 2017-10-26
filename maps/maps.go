package maps

import "fmt"

/*ExecMapsExample is the main function of maps package*/
func ExecMapsExample() {
	fmt.Println("|-------Beginning of maps package-------|\n")
	var myMap map[string]int
	yourMap := make(map[string]int)
	ourMap := map[string]string{
		"First key":  "First value",
		"Second key": "Second value",
	}

	myMap = make(map[string]int)
	myMap["1"] = 1

	yourMap["1"] = 1
	yourMap["2"] = 2

	fmt.Println(myMap, yourMap, ourMap)
	delete(yourMap, "2")
	fmt.Println(yourMap)

	value, isIn := yourMap["1"]
	fmt.Println(value, isIn)

	myGreeting := map[int]string{
		0: "Good morning!",
		1: "Bonjour!",
		2: "Buenos dias!",
		3: "Bongiorno!",
	}
	if val, exists := myGreeting[2]; exists {
		delete(myGreeting, 2)
		fmt.Println(val, exists)
	} else {
		fmt.Println("The value doesn't exists.")
		fmt.Println(val, exists)
	}

	for key, value := range myGreeting {
		fmt.Println(key, ":", value)
	}
	fmt.Println("\n|-------End of maps pacakge-------------|\n")
}
