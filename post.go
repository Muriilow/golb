package golb

import (
	"fmt"
	"net/http"
	"os"
	"io"
    "bytes"
    "github.com/yuin/goldmark"
	"github.com/yuin/goldmark-highlighting/v2"
	"github.com/adrg/frontmatter"
	"html/template"
	"strings"
	"path/filepath"
)

type SlugReader interface {
	Read(slug string) (string, error)
} 

type FileReader struct { 
	Dir string
}

type MetadataQuerier interface {
	Query() ([]PostMetadata, error)
}

type PostMetadata struct {
	Slug string
	Description string `json:"description"`
	Author string `json:"author"`
	Title string `json:"title"`
	Date string `json:"date"`
}

type PostData struct {
	Content template.HTML
	Description string `json:"description"`
	Author string `json:"author"`
	Title string `json:"title"`
	Date string `json:"date"`

}

type IndexData struct {
	Posts []PostMetadata
}


func (fr FileReader) Read(slug string) (string, error) {
	slugPath := filepath.Join(fr.Dir, slug+".md")

	f, err := os.Open(slugPath)
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

func (fr FileReader) Query() ([]PostMetadata, error) {
	postsPath := filepath.Join(fr.Dir, "*.md")

	filenames, err := filepath.Glob(postsPath)

	if err != nil {
		return nil, fmt.Errorf("Querying for files: %w", err)
	}

	var posts []PostMetadata

	for _, filename := range filenames {
		f, err := os.Open(filename)

		if err != nil {
			return nil, fmt.Errorf("Opening file %s: %w", filename, err)
		}

		defer f.Close()

		var post PostMetadata
		_, err = frontmatter.Parse(f, &post)
		
		if err != nil {
			return nil, fmt.Errorf("Error parsing frontmatter for file %s: %w", filename, err)
		}

		post.Slug = strings.TrimSuffix(filepath.Base(filename), ".md")
		posts = append(posts, post)
	}

	return posts, nil
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
		
		var post PostData
		remainingMd, err := frontmatter.Parse(strings.NewReader(postMarkdown), &post)
		
		if err != nil {
			http.Error(w, "Error paarsing frontmatter", http.StatusInternalServerError)
		}

		var buf bytes.Buffer
		err = mdRenderer.Convert([]byte(remainingMd), &buf);
		
		if  err != nil {
			panic(err)
		}

		post.Content = template.HTML(buf.String())	
		err = tpl.Execute(w, post)

		if err != nil {
			http.Error(w, "Error executing template", http.StatusInternalServerError)
			return
		}
	}
}

func IndexHandler(mq MetadataQuerier, tpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		posts, err := mq.Query()

		if err != nil {
			http.Error(w, "Error querying posts", http.StatusInternalServerError)
			return
		}
		
		data := IndexData{
			Posts: posts,
		}

		err = tpl.Execute(w, data)

		if err != nil {
			http.Error(w, "Error executing template", http.StatusInternalServerError)
			return
		}
	}
}
