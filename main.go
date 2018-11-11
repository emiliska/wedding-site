// Thank you @picat for the beautiful code <3

package main

import (
	"io/ioutil"
	"net/http"

	"google.golang.org/appengine"
)

// Home returns the home page of the application.
func Home(w http.ResponseWriter, r *http.Request) {
	file, err := ioutil.ReadFile("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(file)
}

func main() {
	http.HandleFunc("/", Home)
	appengine.Main()
}
