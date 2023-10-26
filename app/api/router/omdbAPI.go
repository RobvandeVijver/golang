package router

import (
	"database/sql"
	"encoding/json"
	"fmt"
	movie "hz/package/models"
	"io"
	"net/http"
	"strings"
	"sync"
)

func ApiRequest(db *sql.DB) {
	imdbIDs := getIMDbIDsFromDB(db)

	// Create a WaitGroup to synchronize the goroutines
	var wg sync.WaitGroup

	// Create a channel to control the number of active goroutines
	maxGoroutines := 5
	semaphore := make(chan struct{}, maxGoroutines)

	// Use a mutex to protect the database
	var dbMutex sync.Mutex

	for _, imdbID := range imdbIDs {
		wg.Add(1)

		// Acquire a semaphore to control the number of active goroutines
		semaphore <- struct{}{}

		// Start a goroutine for each IMDb ID
		go func(imdbID string) {
			defer wg.Done()
			defer func() {
				// Release the semaphore when the goroutine is finished
				<-semaphore
			}()

			// Call the API request
			movieData, err := callAPI(imdbID)
			if err != nil {
				fmt.Printf("Error calling API for IMDb ID %s: %v\n", imdbID, err)
				return
			}

			// JSON parsing
			parsedMovie, err := parseJSON(movieData)
			if err != nil {
				fmt.Printf("Error parsing JSON for IMDb ID %s: %v\n", imdbID, err)
				return
			}

			// Write to the database within a transaction
			err = writeToDatabase(&dbMutex, db, imdbID, parsedMovie)
			if err != nil {
				fmt.Printf("Error writing to the database for IMDb ID %s: %v\n", imdbID, err)
			}
		}(imdbID)
	}

	// Wait for all IMDb data processing to complete
	wg.Wait()

	fmt.Println("Summaries added")
}

func callAPI(imdbID string) ([]byte, error) {
	apiUrl := "https://www.omdbapi.com/"
	apiKey := "c2a32020"

	url := fmt.Sprintf("%s?apikey=%s&i=%s", apiUrl, apiKey, imdbID)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Non-OK status code")
	}

	// Read the response body
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

func parseJSON(data []byte) (movie.Movie, error) {
	var movieData movie.Movie
	decoder := json.NewDecoder(strings.NewReader(string(data)))
	if err := decoder.Decode(&movieData); err != nil {
		return movie.Movie{}, err
	}
	return movieData, nil
}

func writeToDatabase(dbMutex *sync.Mutex, db *sql.DB, imdbID string, movieData movie.Movie) error {
	dbMutex.Lock()
	defer dbMutex.Unlock()

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// Get the plot summary from movieData
	summaryPlot := movieData.Plot

	// Update the database with the plot summary within the transaction
	updateQuery := "UPDATE movies SET Plot_summary = ? WHERE IMDb_id = ?"
	_, err = tx.Exec(updateQuery, summaryPlot, imdbID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func getIMDbIDsFromDB(db *sql.DB) []string {
	var imdbIDs []string
	query := "SELECT IMDb_id FROM movies"

	rows, err := db.Query(query)
	if err != nil {
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var imdbID string
		if err := rows.Scan(&imdbID); err != nil {
			return nil
		}
		imdbIDs = append(imdbIDs, imdbID)
	}

	if err := rows.Err(); err != nil {
		return nil
	}

	return imdbIDs
}
