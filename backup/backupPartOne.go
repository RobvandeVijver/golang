package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Movie struct {
	IMDbID string
	Title  string
	Rating float64
	Year   int
}

var db *sql.DB

func main() {
	// get localhost:8090
	baseURL := GetHost()

	// DB connection
	var err error
	db, err = sql.Open("sqlite3", "./movies.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to Homepage")
	})

	http.HandleFunc("/movies", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to Movies Page")
	})

	arguments := os.Args[1:]

	if len(arguments) > 0 {
		switch arguments[0] {
		case "help":
			printHelpMessage()

		case "add":
			addMovie()

		case "list":
			listMovies()

		case "details":
			getMovieDetails(arguments)

		case "delete":
			deleteMovie(arguments)

		default:
			fmt.Println("invalid input")
		}
	} else {
		err := http.ListenAndServe(baseURL, nil)
		if err != nil {
			panic(err)
		}
	}

}

func GetHost() string {
	host := os.Getenv("API_HOST")
	if host == "" {
		host = "localhost:8090"
	}
	return host
}

func printHelpMessage() {
	helpMessage := "Need help? Check the usecases!"
	optionAdd := "add, add a movie"
	optionList := "list, list all movies"
	optionDetails := "details, Details of a movie"
	optionDelete := "delete, delete a movie"

	fmt.Println(helpMessage)
	fmt.Println(optionAdd)
	fmt.Println(optionList)
	fmt.Println(optionDetails)
	fmt.Println(optionDelete)
}

func addMovie() {
	addCommand := flag.NewFlagSet("add", flag.ExitOnError)
	imdbID := addCommand.String("imdbid", "tt0000001", "IMDb ID of a movie")
	title := addCommand.String("title", "Carmencita", "Title of a movie")
	year := addCommand.Int("year", 1894, "Year of release of the movie")
	rating := addCommand.Float64("rating", 5.7, "IMDb-rate of a movie")
	addCommand.Parse(os.Args[2:])

	query := "INSERT INTO movies (IMDb_id, Title, Rating, Year) VALUES (?, ?, ?, ?)"
	_, err := db.Exec(query, *imdbID, *title, *rating, *year)
	if err != nil {
		fmt.Println("Error by adding the movie details:", err.Error())
		return
	}
	fmt.Printf("IMDb id: %s\n", *imdbID)
	fmt.Printf("Title: %s\n", *title)
	fmt.Printf("Rating: %.1f\n", *rating)
	fmt.Printf("Year: %d\n", *year)
}

func listMovies() {
	query := "SELECT * FROM movies"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Error database connection:", err)
		return
	}
	defer rows.Close()

	var movies []Movie
	for rows.Next() {
		var movie Movie
		if err := rows.Scan(&movie.IMDbID, &movie.Title, &movie.Rating, &movie.Year); err != nil {
			fmt.Println("Error scanning the row", err)
			return
		}
		movies = append(movies, movie)
	}

	if len(movies) == 0 {
		fmt.Println("No movies found.")
		return
	}

	for _, movie := range movies {
		fmt.Printf("%s\n", movie.Title)
	}
}

func getMovieDetails(arguments []string) {
	if len(arguments) < 2 {
		fmt.Println("Use: movie details <IMDb ID>")
		return
	}
	imdbID := arguments[2]
	query := "SELECT * FROM movies WHERE IMDb_id = ?"
	rows, err := db.Query(query, imdbID)
	if err != nil {
		fmt.Println("Error accessing movie details:", err.Error())
		return
	}
	defer rows.Close()

	if !rows.Next() {
		fmt.Printf("No movie found with the IMDb id: %s\n", imdbID)
		return
	}

	var movie Movie
	if err := rows.Scan(&movie.IMDbID, &movie.Title, &movie.Rating, &movie.Year); err != nil {
		fmt.Println("Error scanning row:", err)
		return
	}

	fmt.Printf("IMDb id: %s\n", movie.IMDbID)
	fmt.Printf("Title: %s\n", movie.Title)
	fmt.Printf("Rating: %.1f\n", movie.Rating)
	fmt.Printf("Year: %d\n", movie.Year)
}

func deleteMovie(arguments []string) {
	deleteCommand := flag.NewFlagSet("delete", flag.ExitOnError)
	imdbIDToDelete := deleteCommand.String("imdbid", "tt0000001", "IMDb ID of a movie")
	deleteCommand.Parse(arguments[1:])

	query := "DELETE FROM movies WHERE IMDb_id = ?"
	result, err := db.Exec(query, *imdbIDToDelete)
	if err != nil {
		fmt.Println("Error deleting a movie:", err.Error())
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		fmt.Println("No movie found with the IMDb id:", *imdbIDToDelete)
	} else {
		fmt.Println("Movie deleted")
	}
}
