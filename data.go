package main

type Amazon_ID struct {
	Title       string   `json: "title"`        // id="aiv-content-title"
	ReleaseYear int      `json: "release_year"` // span class="release-year"
	Actors      []string `json: "actors"`       // class="dv-meta-info size-small" dd -> names
	Poster      string   `json: "posters"`      // div class="dp-meta-icon-container" -> img src
	SimilarIDs  []string `json: "similar_ids"`  // ul class="shelf" -> li data-asin=ID
}
