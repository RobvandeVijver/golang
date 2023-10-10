package movie

type Movie struct {
	IMDbID string  `json:"IMDb_id" db:"IMDb_id"`
	Title  string  `json:"Title" db:"Title"`
	Rating float64 `json:"Rating" db:"Rating"`
	Year   int     `json:"Year" db:"Year"`
}
