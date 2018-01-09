package templates

import (
	"log"
	"math"
	"os"
	"strings"
	"text/template"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("../templates/files/*.gohtml"))
}

/*ExecTemplatesExample is the main function of switchGo package*/
func ExecTemplatesExample() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", "three")
	if err != nil {
		log.Fatalln(err)
	}
}

/*ExecTemplatesPassingSlicesExample*/
func ExecTemplatesPassingSlicesExample() {
	sages := []string{"Gandhi", "MLK", "Buddha", "Jesus", "Muhammad"}

	err := tpl.ExecuteTemplate(os.Stdout, "slices.gohtml", sages)
	if err != nil {
		log.Fatalln(err)
	}
}

/*ExecTemplatesPassingMapsExample*/
func ExecTemplatesPassingMapsExample() {
	sages := map[string]string{
		"India":    "Gandhi",
		"America":  "MLK",
		"Meditate": "Buddha",
		"Love":     "Jesus",
		"Prophet":  "Muhammad",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "map.gohtml", sages)
	if err != nil {
		log.Fatalln(err)
	}
}

type sage struct {
	Name  string
	Motto string
}

/*ExecTemplatesPassingStructsExample*/
func ExecTemplatesPassingStructsExample() {
	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "struct.gohtml", buddha)
	if err != nil {
		log.Fatalln(err)
	}
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

type items struct {
	Wisdom    []sage
	Transport []car
}

/*ExecTemplatesPassingComplexStructExample*/
func ExecTemplatesPassingComplexStructExample() {
	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	gandhi := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	mlk := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed",
	}

	jesus := sage{
		Name:  "Jesus",
		Motto: "Love all",
	}

	f := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}

	c := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	sages := []sage{buddha, gandhi, mlk, jesus}
	cars := []car{f, c}
	data := items{
		Wisdom:    sages,
		Transport: cars,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "complex-struct.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}

var fm = template.FuncMap{
	"uc":       strings.ToUpper,
	"ft":       firstThree,
	"fdateMDY": monthDayYear,
	"fdbl":     double,
	"fsq":      square,
	"fsqrt":    sqRoot,
}

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func double(x int) int {
	return x + x
}

func square(x int) float64 {
	return math.Pow(float64(x), 2)
}

func sqRoot(x float64) float64 {
	return math.Sqrt(x)
}

/*ExecTemplatesPassingFuncsExample*/
func ExecTemplatesPassingFuncsExample() {
	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	gandhi := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	mlk := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed",
	}

	jesus := sage{
		Name:  "Jesus",
		Motto: "Love all",
	}

	f := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}

	c := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	sages := []sage{buddha, gandhi, mlk, jesus}
	cars := []car{f, c}
	data := items{
		Wisdom:    sages,
		Transport: cars,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "func.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}

/*ExecTemplatesPassingDatesExample*/
func ExecTemplatesPassingDatesExample() {

	err := tpl.ExecuteTemplate(os.Stdout, "date.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}

/*ExecTemplatesPipelineExample*/
func ExecTemplatesPipelineExample() {

	err := tpl.ExecuteTemplate(os.Stdout, "pipeline.gohtml", 3)
	if err != nil {
		log.Fatalln(err)
	}
}
