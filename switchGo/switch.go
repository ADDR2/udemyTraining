package switchGo

import "fmt"

func switchOnType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown")
	}
}

/*ExecSwitchExample is the main function of switchGo package*/
func ExecSwitchExample() {
	fmt.Println("|-------Beginning of switchGo package-------|\n")
	switch "Julia" {
	case "Daniel":
		fmt.Println("Wassup Daniel")
	case "Medivh", "Julia":
		fmt.Println("Wassup Medivh or Julia")
		fallthrough
	case "Jenny":
		fmt.Println("Wassup Jenny")
	default:
		fmt.Println("Have you no friends?")
	}

	myName := "Amaro"
	switch {
	case len(myName) == 5:
		fmt.Println("Wassup friend with a name of length 5")
	case myName == "Medivh", myName == "Amaro":
		fmt.Println("Wassup Medivh or Amaro")
		fallthrough
	case myName == "Jenny":
		fmt.Println("Wassup Jenny")
	default:
		fmt.Println("Have you no friends?")
	}

	switchOnType("rfdef")
	fmt.Println("\n|-------End of switchGo pacakge-------------|\n")
}
