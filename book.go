package entry

import (
	"time"
)

// Book describes the book contained within a newznab entry
type Book struct {
	// embed common information
	MusicBook

	// title of the book
	BookTitle string
	// date the book was published
	Published time.Time
	// author of the book
	Author string
	// number of pages in the book
	Pages int
}

// Title returns the title of a Book
func (b Book) Title() string {
	return b.BookTitle
}
