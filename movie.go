package entry

import (
	"strconv"

	"github.com/eefret/gomdb"
)

// Movie describes the movie contained within a newznab entry
type Movie struct {
	// embed common information
	TVMovie
	// embed common information
	MovieMusicBook
	// corresponding IMDB entry
	IMDBEntry gomdb.MovieResult
}

// Title returns the title for the movie according to IMDB
func (m Movie) Title() string {
	return m.IMDBEntry.Title
}

// ReleaseYear returns the year the movie was released according to IMDB
func (m Movie) ReleaseYear() int {
	year, _ := strconv.ParseInt(m.IMDBEntry.Year, 10, 64)
	return int(year)
}
