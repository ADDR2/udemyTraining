package structs

import "fmt"

type Person struct {
	first string
	last  string
	age   int
	Arm
}

type Arm struct {
	strengh int
}

func (receiver Person) GetStrengh() int {
	return receiver.Arm.strengh
}

func (receiver *Person) SetStrengh(strengh int) {
	(*receiver).Arm.strengh = strengh
}

/*ExecStructsExample is the main function of structs package*/
func ExecStructsExample() {
	fmt.Println("|-------Beginning of structs package-------|\n")
	p1 := Person{"James", "Bond", 20, Arm{42}}
	fmt.Println(p1.first, p1.last, p1.age, p1.GetStrengh())
	fmt.Println("\n|-------End of structs pacakge-------------|\n")
}
