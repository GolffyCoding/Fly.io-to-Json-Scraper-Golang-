package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

type Article struct {
	Title   string `json:"title"`
	URL     string `json:"url"`
	Content string `json:"content"`
}

func cleanText(text string) string {
	// Remove extra whitespace
	text = regexp.MustCompile(`\s+`).ReplaceAllString(text, " ")
	return strings.TrimSpace(text)
}

func extractArticleContent(e *colly.HTMLElement) Article {
	article := Article{
		URL:   e.Request.URL.String(),
		Title: cleanText(e.ChildText("h1")),
	}

	// Extract and format content
	var content strings.Builder
	e.ForEach("h2, h3, p", func(i int, el *colly.HTMLElement) {
		text := cleanText(el.Text)
		if text == "" {
			return
		}

		switch el.Name {
		case "h2", "h3":
			content.WriteString("\n## " + text + "\n\n")
		case "p":
			content.WriteString(text + "\n\n")
		}
	})

	article.Content = strings.TrimSpace(content.String())
	return article
}

func saveArticlesToJSON(articles []Article) error {
	timestamp := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("articles_%s.json", timestamp)

	// Convert to pretty JSON
	jsonData, err := json.MarshalIndent(articles, "", "  ")
	if err != nil {
		return err
	}

	// Save to file
	return os.WriteFile(filename, jsonData, 0644)
}

func main() {
	c := colly.NewCollector(
		colly.AllowURLRevisit(),
	)

	articles := make([]Article, 0)
	visitedURLs := make(map[string]bool)

	c.OnHTML(`a[class="opacity-0 absolute inset-0"]`, func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if !strings.HasPrefix(link, "http") {
			link = e.Request.AbsoluteURL(link)
		}

		if !visitedURLs[link] {
			visitedURLs[link] = true
			fmt.Printf("\nProcessing: %s\n", link)

			contentCollector := c.Clone()

			contentCollector.OnHTML("article", func(e *colly.HTMLElement) {
				article := extractArticleContent(e)
				articles = append(articles, article)

				// Display article info
				fmt.Printf("Title: %s\n", article.Title)

				fmt.Println("----------------------------------------")
			})

			err := contentCollector.Visit(link)
			if err != nil {
				log.Printf("Error visiting %s: %v\n", link, err)
			}
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Printf("Error: %v\n", err)
	})

	err := c.Visit("https://fly.io/blog/")
	if err != nil {
		log.Fatal(err)
	}

	// Save articles to JSON
	err = saveArticlesToJSON(articles)
	if err != nil {
		log.Printf("Error saving articles to JSON: %v\n", err)
	}

	fmt.Printf("\nScraped %d articles and saved to JSON file\n", len(articles))
}
