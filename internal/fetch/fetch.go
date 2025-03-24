package fetch

import (
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func FetchTextFromURL(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	var sb strings.Builder
	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		sb.WriteString(s.Text())
		sb.WriteString("\n")
	})

	return sb.String(), nil
}
