# Fly.io Blog Scraper

## Overview
This Go program scrapes blog articles from [Fly.io Blog](https://fly.io/blog/) using the `colly` web scraping framework. The extracted articles, including their title, URL, and content, are saved in a JSON file.

## Features
- Extracts article titles, URLs, and content (headings and paragraphs).
- Cleans and formats the extracted text.
- Saves the scraped articles in a timestamped JSON file.
- Uses `colly` for efficient web scraping.

## Prerequisites
- Go 1.18 or later installed
- Internet connection for fetching articles

## Installation
1. Clone the repository or copy the `main.go` file.
2. Install dependencies:
   ```sh
   go mod init flyio-scraper  # If not initialized
   go get github.com/gocolly/colly/v2
   ```

## Usage
1. Run the program:
   ```sh
   go run main.go
   ```
2. The scraped articles will be saved in a file named `articles_YYYYMMDD_HHMMSS.json` in the current directory.

## Output Format
The JSON output contains an array of articles:
```json
[
  {
    "title": "Example Title",
    "url": "https://fly.io/blog/example-article",
    "content": "## Subheading\n\nParagraph text..."
  }
]
```

## Error Handling
- If a request fails, the error is logged.
- If an article lacks content, it is skipped.

## Dependencies
- [Colly](https://github.com/gocolly/colly) - A powerful Go scraping framework.

## License
This project is licensed under the MIT License.

