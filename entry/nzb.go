package entry

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
