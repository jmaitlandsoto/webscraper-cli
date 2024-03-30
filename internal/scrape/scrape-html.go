package scrape

import (
	"context"
	"fmt"

	"github.com/chromedp/chromedp"
)

func ScrapeHTML(url string) (string, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Navigate to the URL
	var siteHTML string
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.InnerHTML("body", &siteHTML, chromedp.ByQuery),
	)
	if err != nil {
		fmt.Printf("Failed to navigate and scrape the page: %v", err)
		return "", err
	}
	return siteHTML, nil
}
