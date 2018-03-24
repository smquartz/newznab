package newznab

import (
	"net/url"

	"github.com/smquartz/nzb"
)

// NZB is a File implementation that describes an NZB file
type NZB struct {
	// embed the NZB representation
	nzb.NZB
	// whether the contents require a password to access
	passworded bool
	// url to download the raw NZB from
	downloadURL *url.URL
}

// URL returns a URL where the raw NZB file may be downloaded from
func (n NZB) URL() *url.URL {
	return n.downloadURL
}

// setURL sets the URL where the raw NZB file may be downloaded from
func (n *NZB) setURL(u *url.URL) {
	n.downloadURL = u
}

// Passworded returns whether the contents of a NZB file require a password to
// access
func (n NZB) Passworded() bool {
	return n.passworded
}

// setPassworded sets whether the contents of a NZB file require a password to access
func (n *NZB) setPassworded(b bool) {
	n.passworded = b
}

// NumFiles returns the number of files a NZB file contains
func (n NZB) NumFiles() int {
	return len(n.Files)
}
