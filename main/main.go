package main

import "github.com/ADDR2/udemyTraining/structs"
import "fmt"

func main() {
	structs.ExecStructsExample()
	p := structs.Person{Arm: structs.Arm{}}
	p.SetStrengh(56)
	fmt.Println(p.GetStrengh())
}
