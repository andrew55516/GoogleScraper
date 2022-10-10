package ScrapeAndParse

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
)

func googleResultParsing(res *http.Response, rank int) ([]searchResult, error) {
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}
	results := []searchResult{}
	sel := doc.Find("div.g")
	rank++
	for i := range sel.Nodes {
		item := sel.Eq(i)
		linkTag := item.Find("a")
		link, _ := linkTag.Attr("href")
		titleTag := item.Find("h3.r")
		descTag := item.Find("span.st")
		desc := descTag.Text()
		title := titleTag.Text()
		link = strings.Trim(link, " ")
		if link != "" && link != "#" && !strings.HasPrefix(link, "/") {
		}
		result := searchResult{
			rank,
			link,
			title,
			desc,
		}
		results = append(results, result)
	}
	return results, nil
}
