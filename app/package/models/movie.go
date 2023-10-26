package movie

type Movie struct {
	IMDbID *string  `json:"imdb_id" db:"IMDb_id"`
	Title  *string  `json:"title" db:"Title"`
	Rating *float64 `json:"rating" db:"Rating"`
	Year   *string  `json:"year" db:"Year"`
	Plot   *string  `json:"plot" db:"Plot_summary"`
}
