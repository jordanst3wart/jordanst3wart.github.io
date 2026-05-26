package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"
	"time"

	"github.com/a-h/templ"
	"github.com/gosimple/slug"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"go.abhg.dev/goldmark/frontmatter"
	"gopkg.in/yaml.v3"

	highlighting "github.com/yuin/goldmark-highlighting/v2" // TODO fork and maintain
)

func Unsafe(html string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		_, err = io.WriteString(w, html)
		return
	})
}

func main() {
	posts := GetWritings()

	rootPath := "public"
	if fileExists(rootPath) {
		if err := os.RemoveAll(rootPath); err != nil {
			log.Fatalf("failed to remove rootPath %v directory", rootPath)
		}
	}
	if err := os.Mkdir(rootPath, 0o755); err != nil {
		log.Fatalf("failed to create output directory: %v", err)
	}

	name := path.Join(rootPath, "index.html")
	file, err := os.Create(name)
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}

	var links []string
	for index, post := range posts {
		// supports number 1 to 9
		if index > 8 {
			break
		}
		links = append(links, post.Metadata.Link)
	}
	err = indexPage(posts, links).Render(context.Background(), file)
	if err != nil {
		log.Fatalf("failed to write index page: %v", err)
	}

	// Create a page for each post.
	for _, post := range posts {
		// Create the output directory.
		dir := path.Join(rootPath, "writing", slug.Make(post.Metadata.Title))
		if err := os.MkdirAll(dir, 0o755); err != nil && err != os.ErrExist {
			log.Fatalf("failed to create dir %q: %v", dir, err)
		}

		// Create the output file.
		name := path.Join(dir, "index.html")
		file, err := os.Create(name)
		if err != nil {
			log.Fatalf("failed to create output file: %v", err)
		}

		// Create an unsafe component containing raw HTML.
		content := Unsafe(post.Content)

		err = contentPage(post.Metadata.Title, post.Metadata.Date.Format("January 2, 2006"), content).Render(context.Background(), file)
		if err != nil {
			log.Fatalf("failed to write output file: %v", err)
		}
	}
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true // File exists
	}
	if errors.Is(err, os.ErrNotExist) {
		return false // File does not exist
	}
	// File may or may not exist (e.g., permission denied)
	return false
}

type Metadata struct {
	Title string
	Tags  []string   // might not need tags
	Date  CustomDate // `yaml:"date"`
	Link  string
}

type Post struct {
	Content  string
	Metadata Metadata
}

func GetWritings() []Post {
	var posts []Post
	files, err := os.ReadDir("writing")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if strings.HasSuffix(file.Name(), ".md") {
			fmt.Printf("reading file: %s\n", file.Name())
			source, err := os.ReadFile(path.Join("writing", file.Name()))
			if err != nil {
				log.Fatalf("error using file %v", err)
			}
			post, err := ParseMarkdownWriting(source)
			if err != nil {
				log.Fatalf("error parsing post %v", err)
			}
			post.Metadata.Link = path.Join("writing", slug.Make(post.Metadata.Title), "/")
			posts = append(posts, post)
		}
	}
	return posts
}

func ParseMarkdownWriting(source []byte) (Post, error) {
	md := goldmark.New(goldmark.WithExtensions(&frontmatter.Extender{},
		highlighting.NewHighlighting(
			highlighting.WithStyle("catppuccin-macchiato"), // TODO change...
		),
	))
	parserCtx := parser.NewContext()
	var buf bytes.Buffer
	if err := md.Convert(source, &buf, parser.WithContext(parserCtx)); err != nil {
		// handle error
		log.Fatalf("convert failed...%v", err)
	}
	data := frontmatter.Get(parserCtx)
	if data == nil {
		log.Fatalf("no front matter") //return errors.New("no front matter")
	}
	var metadata Metadata
	if err := data.Decode(&metadata); err != nil {
		log.Fatalf("decode failed %v", err) //	return err
	}
	return Post{Content: buf.String(), Metadata: metadata}, nil
}

// parsing the date in a different format.. :/
type CustomDate struct {
	time.Time
}

func (d *CustomDate) UnmarshalYAML(value *yaml.Node) error {
	// Custom layout: Change this to match your YAML's date format
	const layout = "2006-01-02"

	t, err := time.Parse(layout, strings.TrimSpace(value.Value))
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}
