package scraper

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/chromedp/chromedp"
	"golang.org/x/net/html"
)

func Scrape(url string) (string, error) {
	// const url = "https://github.com/chromedp/chromedp"

	// Create a new Chrome instance
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Navigate to the URL
	var siteContent string
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.InnerHTML("body", &siteContent, chromedp.ByQuery),
	)
	if err != nil {
		fmt.Printf("Failed to navigate and scrape the page: %v", err)
		return "", err
	}

	text := extractTextFromHTML(siteContent)

	fmt.Println(text)
	return text, nil
}

func extractTextFromHTML(htmlContent string) string {
	var textContentBuffer bytes.Buffer

	tokenizer := html.NewTokenizer(bytes.NewReader([]byte(htmlContent)))

	for {
		tokenType := tokenizer.Next()
		switch tokenType {

		case html.TextToken:
			text := tokenizer.Text()
			trimmed := strings.TrimSpace(string(text))
			textContentBuffer.WriteString(trimmed + " ")

		case html.ErrorToken:
			return textContentBuffer.String()
		}
	}
}
