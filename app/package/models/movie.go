package movie

type Movie struct {
	IMDbID string  `json:"imdb_id" db:"IMDb_id"`
	Title  string  `json:"title" db:"Title"`
	Rating float64 `json:"rating" db:"Rating"`
	Year   int     `json:"year" db:"Year"`
}
