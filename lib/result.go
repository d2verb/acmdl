package acmdl

import (
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Result struct {
	Title    string
	Abstract string
	Keywords string
	Url      string
}

func ParseResponse(resp *http.Response) ([]Result, error) {
	results := []Result{}
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, err
	}

	details := doc.Find(".details")
	details.Each(func(index int, s *goquery.Selection) {
		title := s.Find(".title").First().Text()
		abstract := s.Find(".abstract").First().Text()
		keywords := s.Find(".kw").First().Text()
		paperUrl, _ := s.Find(".title>a").First().Attr("href")

		results = append(results, Result{
			Title:    strings.TrimSpace(title),
			Abstract: strings.TrimSpace(abstract),
			Keywords: strings.TrimSpace(strings.Replace(keywords, "\n", " ", -1)),
			Url:      strings.TrimSpace("https://dl.acm.org/" + paperUrl),
		})
	})

	return results, nil
}
