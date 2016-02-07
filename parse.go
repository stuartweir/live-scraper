package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ExtractInfo receives a URL, downloads the document at that URL,
// and then populates the AmazonID struct before returning.
// Note: Dependent on PuerkitoBio's goquery package.
func ExtractInfo(url string) (AmazonID, error) {
	amznID := AmazonID{
		Title:       "",
		ReleaseYear: 0,
		Actors:      []string{},
		Poster:      "",
		SimilarIDs:  []string{},
	}

	doc, err := goquery.NewDocument(url)
	if err != nil {
		return amznID, err
	}

	amznID.Title = findTitle(doc)
	amznID.ReleaseYear = findRY(doc)
	amznID.Actors = findActors(doc)
	amznID.Poster = findPoster(doc)
	amznID.SimilarIDs = findSIDs(doc)

	if err := checkID(amznID); err != nil {
		return amznID, err
	}

	return amznID, nil
}

func checkID(amznID AmazonID) error {
	if amznID.Title == "" || len(amznID.Actors) == 0 || amznID.Poster == "" {
		return fmt.Errorf("possible important values missing: title (%s); actors (%s); poster (%s)",
			amznID.Title, amznID.Actors, amznID.Poster)
	}
	return nil
}

func findTitle(doc *goquery.Document) string {
	return strings.TrimSpace(doc.Find("#aiv-content-title").Before("span").Text())
}

func findRY(doc *goquery.Document) int {
	i, err := strconv.Atoi(doc.Find(".release-year").Text())
	if err != nil {
		fmt.Printf("Bad Request: %v\n", http.StatusBadRequest)
	}
	return i
}

func findActors(doc *goquery.Document) []string {
	actors := strings.SplitAfter(strings.TrimSpace(doc.Find(".dv-meta-info").Find("dd").First().Text()), ",")
	for i, actor := range actors {
		actor = strings.TrimSpace(actor)
		actor = strings.TrimSuffix(actor, ",")
		actors[i] = actor
	}
	return actors
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
