package httpServers

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
)

func init() {
	tpl = template.Must(template.New("").ParseGlob("../httpServers/views/*.gohtml"))
}

type id int

func (identifier id) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Amaro-key", "This is from Amaro")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>Any code you want in this func</h1>")
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

type hotdog int

func (handler hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog doggy")
}

type hotcat int

func (handler hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

/*ExecHTTPMuxExample  --  */
func ExecHTTPMuxExample() {
	fmt.Println("|-------Beginning of httpServers package-------|\n")
	var dog hotdog
	var cat hotcat

	mux := http.NewServeMux()
	mux.Handle("/dog/", dog) // catches: /dog, /dog/, /dog/something/else and so on
	mux.Handle("/cat", cat)  // catches: /cat

	http.ListenAndServe(":8080", mux)
	fmt.Println("\n|-------End of httpServers pacakge-------------|\n")
}

/*ExecHTTPDefaultMuxExample  --  */
func ExecHTTPDefaultMuxExample() {
	fmt.Println("|-------Beginning of httpServers package-------|\n")

	http.HandleFunc("/dog/", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "dog dog doggy")
	})

	http.HandleFunc("/cat", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "cat cat cat")
	})

	http.ListenAndServe(":8080", nil)
	fmt.Println("\n|-------End of httpServers pacakge-------------|\n")
}

func firstHandleFunc(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "First first first")
}

func secondHandleFunc(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Second second second")
}

/*ExecHTTPMuxFuncExample  --  */
func ExecHTTPMuxFuncExample() {
	fmt.Println("|-------Beginning of httpServers package-------|\n")

	mux := http.NewServeMux()
	mux.Handle("/first/", http.HandlerFunc(firstHandleFunc))  // catches: /dog, /dog/, /dog/something/else and so on
	mux.Handle("/second", http.HandlerFunc(secondHandleFunc)) // catches: /cat

	http.ListenAndServe(":8080", mux)
	fmt.Println("\n|-------End of httpServers pacakge-------------|\n")
}
