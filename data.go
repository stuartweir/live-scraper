package main

// AmazonID is a struct that shares (essentially) the JSON Representation
// that will later be provided to the user.
type AmazonID struct {
	// Title is a string representation of the text content inside the element with id="aiv-content-title"
	Title string `json:"title"`

	// ReleaseYear is an integer representation of the text content inside the element with class="release-year"
	ReleaseYear int `json:"release_year"`

	// Actors is the string slice that represents the names of the actors inside the element with class="dv-meta-info size-small"
	Actors []string `json:"actors"`

	// Poster is the string representation of the URL inside the element with class="dp-meta-icon-container"
	Poster string `json:"posters"`

	// SimilarIDs is the string slice that represents the Amazon IDs of similar movies, found in the li elements with data-asin=ID
	SimilarIDs []string `json:"similar_ids"`
}
