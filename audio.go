package newznab

// Music describes the music contained within a newznab entry
type Music struct {
	// embed common information
	TVMovieMusic
	// embed common information
	MusicBook

	// name of the artist
	Artist string
	// name of the album
	Album string
	// track listing
	Tracks []string
}

// Title returns the title of an album
func (m Music) Title() string {
	return m.Album
}
