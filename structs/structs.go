package structs

import "fmt"
import "encoding/json"

type Person struct {
	first string
	last  string
	age   int
	Arm
}

type Arm struct {
	strengh int
}

func (p Person) fullName() string {
	return p.first + p.last
}

func (receiver Person) GetStrengh() int {
	return receiver.Arm.strengh
}

func (receiver *Person) SetStrengh(strengh int) {
	(*receiver).Arm.strengh = strengh
}

type Me struct {
	Name        string `json:"The Name"`
	Skill       string `json:"-"`
	notExported int
}

type NormalPerson struct {
	Name string `json:"The Name"`
	Age  int
}

/*ExecStructsExample is the main function of structs package*/
func ExecStructsExample() {
	fmt.Println("|-------Beginning of structs package-------|\n")
	p1 := Person{"James", "Bond", 20, Arm{42}}
	p2 := Person{"Jenny", "White", 23, Arm{46}}
	fmt.Println(p1.first, p1.last, p1.age, p1.GetStrengh())
	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())

	p3 := &Person{"Name", "Last name", 44, Arm{2}}
	fmt.Println(p3)
	fmt.Printf("%T\n", p3)
	fmt.Println(p3.first)
	fmt.Println(p3.last)

	bs, _ := json.Marshal(Me{"Amaro", "Wisdom", 32})
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	fmt.Println(string(bs))

	var me1 NormalPerson
	fmt.Println(me1.Name)
	fmt.Println(me1.Age)

	bytes := []byte(`{"The Name":"Amaro", "Age":25}`)
	json.Unmarshal(bytes, &me1)

	fmt.Println(me1.Name, me1.Age)
	fmt.Printf("%T\n", me1)
	fmt.Println("\n|-------End of structs pacakge-------------|\n")
}
