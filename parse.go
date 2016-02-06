// 3. handling HTML (traversing, parsing)
package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ExtractInfo(url string) (*Amazon_ID, error) {
	var amznID *Amazon_ID
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}

	amznID.Title = findTitle(doc)
	amznID.ReleaseYear = findRY(doc)
	if amznID.ReleaseYear == 0 {
		return nil, fmt.Errorf("Expected year, got ", 0)
	}
	amznID.Actors = findActors(doc)
	amznID.Poster = findPoster(doc)
	amznID.SimilarIDs = findSIDs(doc)

	return amznID, nil
}

func findTitle(doc *goquery.Document) string {
	return strings.TrimSpace(doc.Find("aiv-content-title").Before("span").Text())
}

func findRY(doc *goquery.Document) int {
	i, err := strconv.Atoi(doc.Find(".release-year").Text())
	if err != nil {
		return 0
	}
	return i
}

func findActors(doc *goquery.Document) []string {
	return strings.SplitAfter(strings.TrimSpace(doc.Find(".dv-meta-info").Find("dd").First().Text()), ",")
}

func findPoster(doc *goquery.Document) string {
	src, _ := doc.Find(".dp-meta-icon-container").Find("img").Attr("src")
	return src
}

func findSIDs(doc *goquery.Document) []string {
	var str []string
	doc.Find(".packshot").Each(func(i int, s *goquery.Selection) {
		otherID, _ := s.Attr("data-asin")
		str = append(str, otherID)
	})
	return str
}
