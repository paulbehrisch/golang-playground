package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly/v2"
	"log"
	"os"
)

// Stores information about a hacker news submission
type Submission struct {
	Title string
	URL   string
}

func main() {
	fName := "submissions.json"
	file, err := os.Create(fName)
	submissions := make([]Submission, 0, 30)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()
	c := colly.NewCollector(
		// Cache responses to prevent multiple download of pages
		// even if the collector is restarted
		colly.CacheDir("./.cache"),
	)

	// Find and visit all links
	c.OnHTML("tr[class='athing']", func(e *colly.HTMLElement) {
		submission := Submission{
			Title: e.ChildText("a.storylink"),
			URL:   e.ChildAttr("a.storylink", "href"),
		}
		submissions = append(submissions, submission)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://news.ycombinator.com/")

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	enc.Encode(submissions)
}
