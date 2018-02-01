package httpServers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
)

func init() {
	tpl = template.Must(template.New("").ParseGlob("../httpServers/views/*.gohtml"))
}

type id int

func (identifier id) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Any code you want in this func")
}

/*ExecHTTPServersExample is the main function of httpServers package*/
func ExecHTTPServersExample() {
	fmt.Println("|-------Beginning of httpServers package-------|\n")
	var identifier id
	http.ListenAndServe(":8080", identifier)
	fmt.Println("\n|-------End of httpServers pacakge-------------|\n")
}

type another int

var moreData bool

func (identifier another) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	if moreData {
		data := struct {
			Method        string
			URL           *url.URL
			Submissions   map[string][]string
			Header        http.Header
			Host          string
			ContentLength int64
		}{
			req.Method,
			req.URL,
			req.Form,
			req.Header,
			req.Host,
			req.ContentLength,
		}

		tpl.ExecuteTemplate(w, "form_with_details.gohtml", data)
	} else {
		tpl.ExecuteTemplate(w, "form.gohtml", req.Form)
	}
}

var tpl *template.Template

/*ExecHTTPFormExample  --  */
func ExecHTTPFormExample() {
	fmt.Println("|-------Beginning of httpServers package-------|\n")
	var identifier another
	moreData = false
	http.ListenAndServe(":8080", identifier)
	fmt.Println("\n|-------End of httpServers pacakge-------------|\n")
}

/*ExecHTTPFormDatailedExample  --  */
func ExecHTTPFormDatailedExample() {
	fmt.Println("|-------Beginning of httpServers package-------|\n")
	var identifier another
	moreData = true
	http.ListenAndServe(":8080", identifier)
	fmt.Println("\n|-------End of httpServers pacakge-------------|\n")
}
