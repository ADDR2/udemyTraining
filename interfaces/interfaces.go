package interfaces

import (
	"fmt"
	"math"
	"sort"
)

type Square struct {
	side float64
}

type Circle struct {
	radius float64
}

type Shape interface {
	area() float64
}

func (s Square) area() float64 {
	return s.side * s.side
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

type people []string

func (p people) Len() int           { return len(p) }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p people) Less(i, j int) bool { return p[i] < p[j] }

func step1() {
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	sort.SliceStable(studyGroup, func(i, j int) bool {
		return studyGroup[i] < studyGroup[j]
	})
	fmt.Println(studyGroup)

	studyGroup = people{"Zeno", "John", "Al", "Jenny"}
	sort.Sort(sort.Reverse(studyGroup))
	fmt.Println(studyGroup)
}

/*ExecInterfacesExample is the main function of interfaces package*/
func ExecInterfacesExample() {
	fmt.Println("|-------Beginning of interfaces package-------|\n")
	s := Square{10}
	c := Circle{5}
	info(s)
	info(c)
	step1()
	fmt.Println("\n|-------End of interfaces pacakge-------------|\n")
}
