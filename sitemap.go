package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

// URL represents the <url> tag in the sitemap.
type URL struct {
	Loc     string `xml:"loc"`
	LastMod string `xml:"lastmod"`
}

// URLSet represents the root <urlset> tag.
type URLSet struct {
	XMLName xml.Name `xml:"http://www.sitemaps.org/schemas/sitemap/0.9 urlset"`
	URLs    []URL    `xml:"url"`
}

func WriteSiteMap(posts []Post, outDir string) {
	var urls []URL
	currentDate := time.Now().Format("2006-01-02")
	urls = append(urls, URL{
		Loc:     "https://stewart.bot/",
		LastMod: currentDate,
	})
	for _, post := range posts {
		urls = append(urls, URL{
			// Adjust this string formatting to match how your URLs are structured
			Loc:     fmt.Sprintf("https://stewart.bot/%s", post.Metadata.Link),
			LastMod: currentDate,
		})
	}

	sitemap := URLSet{
		URLs: urls,
	}

	output, err := xml.MarshalIndent(sitemap, "", "  ")
	if err != nil {
		log.Fatalf("error marshalling xml %v", err)
	}

	xmlData := append([]byte(xml.Header), output...)

	if err := os.MkdirAll(outDir, 0755); err != nil {
		log.Fatalf("failed to create directory: %v", err)
	}

	filePath := filepath.Join(outDir, "sitemap.xml")
	if err := os.WriteFile(filePath, xmlData, 0644); err != nil {
		log.Fatalf("failed to write sitemap file: %v", err)
	}
}
