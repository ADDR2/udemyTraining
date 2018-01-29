package httpServers

import (
	"fmt"
	"net/http"
)

type id int

func (identifier id) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Any code you want in this func")
}

/*ExecHttpServersExample is the main function of httpServers package*/
func ExecHttpServersExample() {
	fmt.Println("|-------Beginning of httpServers package-------|\n")
	var identifier id
	http.ListenAndServe(":8080", identifier)
	fmt.Println("\n|-------End of httpServers pacakge-------------|\n")
}
