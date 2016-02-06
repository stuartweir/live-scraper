// 4. prepping data for user (making struct + JSON)
package main

type Amazon_ID struct {
	Title       string   `json: "title"`        // h1 id="aiv-content-title"
	ReleaseYear int      `json: "release_year"` // span class="releease-year"
	Actors      []string `json: "actors"`       // dl class="dv-meta-info size-small" dd -> names, not nums
	Poster      string   `json: "posters"`      // div class="dp-meta-icon-container" -> img src
	SimilarIDs  []string `json: "similar_ids"`  // ul class="shelf cf" -> li data-asin=ID
}
