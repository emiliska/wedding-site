package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
)

func main() {
	http.HandleFunc("/", handle)
	http.HandleFunc("/_ah/health", healthCheckHandler)
	log.Print("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	content, err:= ioutil.ReadFile("https://storage.googleapis.com/eubankstocraig.com/index.html")
	if err!= nil {
		log.Fatal(err)
		return
	}
	fmt.Fprint(w, string(content))
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}
