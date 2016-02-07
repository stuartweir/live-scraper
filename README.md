# Live Scraper

Live Scraper is a service that receives a request for an Amazon ID, Scrapes Amazon's website for information based on that ID, and responds back with a JSON Representation of that Resource, or an error signifying that it failed to find any data on that Resource.

## How to Build/Run:

1. `go get ./..`
2. `cd` to directory
3. `go build`
4. `./live-scraper`
5. Make request to `localhost:8080/movie/amazon/{insert_amazon_id_here}`

## On Package Selection:

Initially, I wanted to simply use the `golang.org/x/net/html` package to parse the HTML nodes found on the Amazon page. While this was possible to an extent without doing too much work, I found myself trying to figure out how to check CSS selectors for certain pieces of data, which ultimately would have lead me down the path of rewriting a lot of the work found in the `goquery` package by PuerkitoBio on GitHub: [goquery](https://github.com/PuerkitoBio/goquery). It seemed like the most practical approach to solve this problem, and was fairly speedy.

## General Outline of My Thought Process While Designing Live-Scraper:

1. Request comes in
2. Parse request
3. Make own request to amazon.com/whatever
4. Get HTML reponse back from Amzon
5. Parse it
6. Get correct elements out. if missing/broken info, bail with an error message to the user
7. Make a struct
8. Marshall data
9. Return JSON to user

1. Talk to user (main.go)
2. Talk to amazon (parse.go)
3. handling HTML (parse.go)
4. prepping data for user (marshal.go/data.go)
