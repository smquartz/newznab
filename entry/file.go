package entry

import (
	"net/url"
)

// File describes the "descriptor" file referred to by a newznab entry.
// This is the file that contains infomation allowing you to download the
// actual files, such as an NZB or torrent file.
type File interface {
	// returns the size of the file contents in bytes
	Size() int64
	// returns the number of files in the entry
	Files() int64
	// returns whether the contents require a password to access
	Passworded() bool
	// returns a URL from which the raw File may be downloaded from
	URL() *url.URL
	// returns the raw bytes for descriptor file
	Bytes() ([]byte, error)
}
