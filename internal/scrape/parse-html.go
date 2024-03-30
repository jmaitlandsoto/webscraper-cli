package scrape

import (
	"bytes"
	"strings"

	"golang.org/x/net/html"
)

func ParseHTML(htmlContent string) (string, error) {
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
			return textContentBuffer.String(), nil
		}
	}
}
