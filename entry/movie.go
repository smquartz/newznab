package entry

// Movie describes the movie contained within a newznab entry
type Movie struct {
	// embed general information
	GenericContent
	// embed common information
	TVMovie
	// embed common information
	MovieMusicBook
	// ID for the corresponding IMDB entry
	IMDBID int64
	// score for the movie according to IMDB
	IMDBScore float32
	// title for the movie according to IMDB
	IMDBTitle string
	// tagline for the movie according to IMDB
	IMDBTagline string
	// plot for the movie according to IMDB
	IMDBPlot string
	// year of the movie according to IMDB
	IMDBYear int
	// director of the movie according to IMDB
	IMDBDirector string
	// actors in the movie according to IMDB
	IMDBActors []string
}

// Title returns the title for the movie according to IMDB
func (m Movie) Title() string {
	return m.IMDBTitle
}

// ReleaseYear returns the year the movie was released according to IMDB
func (m Movie) ReleaseYear() int {
	return m.IMDBYear
}
