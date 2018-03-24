package entry

// Content describes the content of an newznab entry; that is, what is actually
// in the files
type Content interface {
	// title of the content
	Title() string
	// year the content was released
	ReleaseYear() int
}
