package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Movie struct {
	Title    string
	Minutes  int
  Year int
}

type Rating struct {
	MovieID string
	RaterID string
	Rating  int
}

type ByMinutesThenTitle []Movie

func (m ByMinutesThenTitle) Len() int{ 
  return len(m) 
}
func (m ByMinutesThenTitle) Swap(i, j int){ 
  m[i], m[j] = m[j], m[i] 
}
func (m ByMinutesThenTitle) Less(i, j int) bool {
	if m[i].Minutes == m[j].Minutes {
		return m[i].Title < m[j].Title
	}
	return m[i].Minutes > m[j].Minutes
}
type ByYearThenTitle []Movie

func (m ByYearThenTitle) Len() int{ 
  return len(m) 
}
func (m ByYearThenTitle) Swap(i, j int){ 
  m[i], m[j] = m[j], m[i] 
}
func (m ByYearThenTitle) Less(i, j int) bool {
	if m[i].Year == m[j].Year {
		return m[i].Title < m[j].Title
	}
	return m[i].Year > m[j].Year
}

func main() {
	file, err := os.Open("movies.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var movies []Movie
	for _, record := range records {
		minutesStr := record[6]
		minutes,err := strconv.Atoi(minutesStr)
		if err != nil {
			continue
		}

		movie := Movie{
			Title:    record[1],
			Minutes: minutes,
		}
		movies = append(movies, movie)
	}
  var movies1 []Movie
	for _, record := range records {
		yearStr := record[2]
		year,err := strconv.Atoi(yearStr)
		if err != nil {
			continue
		}

		movie := Movie{
			Title: record[1],
			Year: year,
		}
		movies1 = append(movies1, movie)
	}

	sort.Sort(ByMinutesThenTitle(movies))
  sort.Sort(ByYearThenTitle(movies1))

  fmt.Printf("\nDuration:\n")
	for _,movie := range movies[:5] {
		fmt.Printf("%s; Duration In Minutes: %d\n",movie.Title,movie.Minutes)
	}
  fmt.Printf("\nYear of release:\n")
  for _,movie1 := range movies1[:5] {
		fmt.Printf("%s; Year of Release: %d\n",movie1.Title,movie1.Year)
	}

  uniqueRaters := make(map[string]bool)
  for _, record := range records {
    raterID := record[0] 
    uniqueRaters[raterID] = true
}
fmt.Printf("\nUnique Raters: %d\n", len(uniqueRaters))
  file1, err := os.Open("ratings.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()
	reader1 := csv.NewReader(file1)
	records1, err1 := reader1.ReadAll()
	if err1 != nil {
		fmt.Println("Error:", err1)
		return
	}
	var ratings []Rating
	for _, record := range records1 {
		rating, err := strconv.Atoi(record[2]) 
		if err != nil {
			continue
		}

		ratingObj := Rating{
			MovieID: record[0], 
			RaterID: record[1], 
			Rating:  rating,
		}
		ratings = append(ratings, ratingObj)
	}
	moviesCount := 0
	moviesMap := make(map[string]int)
	for _, rating := range ratings {
		if rating.Rating >= 7 {
			moviesMap[rating.MovieID]++
		}
	}

	for _, count := range moviesMap {
		if count >= 5 {
			moviesCount++
		}
  }
  fmt.Println("\nNumber of movies: ",moviesCount)

	ratingsFile, err := os.Open("ratings.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer ratingsFile.Close()

	moviesFile, err := os.Open("movies.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer moviesFile.Close()
	ratingsRecords, err := csv.NewReader(ratingsFile).ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	moviesRecords, err := csv.NewReader(moviesFile).ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	genreCounts := make(map[string]int)
	for _, record := range ratingsRecords {
		raterID, err := strconv.Atoi(record[0])
		if err != nil {
			continue
		}
		if raterID == 1040 {
			movieID, err := strconv.Atoi(record[1])
			if err != nil {
				continue
			}
			rating, err := strconv.Atoi(record[2])
			if err != nil {
				continue
			}
			var genre string
			for _, movieRecord := range moviesRecords {
				id, err := strconv.Atoi(movieRecord[0])
				if err != nil {
					continue
				}
				if id == movieID {
					genre = movieRecord[2]
					break
				}
			}
			if rating >= 7 {
				genreCounts[genre]++
			}
		}
	}
	var maxCount int
	var favoriteGenre string
	for genre, count := range genreCounts {
		if count > maxCount {
			maxCount = count
			favoriteGenre = genre
		}
	}

	fmt.Printf("\nFavorite movie genre of rater 1040: %s\n", favoriteGenre)
}
