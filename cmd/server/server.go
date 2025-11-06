package main

import (
	"fmt"
	"github.com/Muriilow/golb"
	"html/template"
	"log"
	"os"
	"net/http"
	"io"
)

func main() {
	err := run(os.Args, os.Stdout)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}


}
func run(args []string, stdout io.Writer) error {
	mux := http.NewServeMux()

	postReader := golb.FileReader{
		Dir: "posts",
	}

	postTemplate := template.Must(template.ParseFiles("post.gohtml"))
	mux.HandleFunc("GET /posts/{slug}", golb.PostHandler(postReader, postTemplate))

	indexTemplate := template.Must(template.ParseFiles("index.gohtml"))
	mux.HandleFunc("GET /", golb.IndexHandler(postReader, indexTemplate))

	err := http.ListenAndServe(":3030", mux)

	if err != nil {
		log.Fatal(err)
	}
	
	return nil
}
