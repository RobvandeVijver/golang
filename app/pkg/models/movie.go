package movie

type Movie struct {
	IMDbID string  `json:"imdbid"`
	Title  string  `json:"title"`
	Rating float64 `json:"rating"`
	Year   int     `json:"year"`
}
