package templates

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("../templates/files/*.gohtml"))
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
		"India": "Gandhi",
		"America": "MLK",
		"Meditate": "Buddha",
		"Love": "Jesus",
		"Prophet": "Muhammad",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "map.gohtml", sages)
	if err != nil {
		log.Fatalln(err)
	}
}