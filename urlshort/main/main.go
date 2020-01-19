package main

import (
	"fmt"
	"io/ioutil"
	"flag"
	"net/http"

	"github.com/nandotheessen/Gophercises/urlshort"
)

var yamlfile string

func init() {
	flag.StringVar(&yamlfile, "pathfile", "", "provade a yaml file composed of sequence of path & url mappings")
	flag.Parse()
}

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)
	var yaml []byte
	var err error
	// Build the YAMLHandler using the mapHandler as the
	// fallback
	if yamlfile != "" {
		yaml, err = ioutil.ReadFile(yamlfile)
		if err != nil {
			fmt.Println(err)
		}
		
	} else {
		yamlstring := `
		- path: /urlshort
		  url: https://github.com/gophercises/urlshort
		- path: /urlshort-final
		  url: https://github.com/gophercises/urlshort/tree/solution
		`
		yaml = []byte(yamlstring)
	}
	
	yamlHandler, err := urlshort.YAMLHandler(yaml, mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}