package main

import (
	"log"
	"net/http"
	"os"
	"io"
    "bytes"
    "github.com/yuin/goldmark"
	"github.com/yuin/goldmark-highlighting/v2"
	"text/template"
)
func main() {
	mux := http.NewServeMux()

	postTemplate := template.Must(template.ParseFiles("post.gohtml"))
	mux.HandleFunc("GET /posts/{slug}", PostHandler(FileReader{}, postTemplate))

	err := http.ListenAndServe(":3030", mux)

	if err != nil {
		log.Fatal(err)
	}
}

type SlugReader interface {
	Read(slug string) (string, error)
} 

type FileReader struct {}

func (fr FileReader) Read(slug string) (string, error) {
	f, err := os.Open(slug + ".md")
	if err != nil {
		return "", err
	}

	defer f.Close()

	text, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}

	return string(text), nil
}
type PostData struct {
	Content string
	Author string 
	Title string
	Date string
}

func PostHandler(sl SlugReader, tpl *template.Template) http.HandlerFunc {
	mdRenderer := goldmark.New(
		goldmark.WithExtensions(
			highlighting.NewHighlighting(
				highlighting.WithStyle("monokai"),
			), 
		),
	)


	return func(w http.ResponseWriter, r *http.Request) {
		slug := r.PathValue("slug")
		postMarkdown, err := sl.Read(slug)

		if err != nil {
			// TODO: Handle different errors in the future
			http.Error(w, "Post not found", http.StatusNotFound)
			return
		}
		
		var buf bytes.Buffer
		err = mdRenderer.Convert([]byte(postMarkdown), &buf);
		
		if  err != nil {
			panic(err)
		}
		
		err = tpl.Execute(w, PostData {
			Content: buf.String(),
			Author: "Mur1il0w",
			Title: "My Blog",
			Date: "today",
		})
		if err != nil {
			http.Error(w, "Error executing template", http.StatusInternalServerError)
			return
		}

	}
}
