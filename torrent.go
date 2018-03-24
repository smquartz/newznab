package entry

import (
	"bytes"
	"net/url"

	"github.com/anacrolix/torrent/metainfo"
)

// Torrent is a File implementation that describes a torrent file
type Torrent struct {
	// embed torrent file representation
	metainfo.MetaInfo
	// whether the contents require a password to access
	passworded bool
	// URL the torrent file may be downloaded from
	downloadURL *url.URL
}

// Size returns the total size of all the files in the torrent
func (t Torrent) Size() int64 {
	info, err := t.UnmarshalInfo()
	if err != nil {
		return -1
	}
	return info.TotalLength()
}

// NumFiles returns the total number of files in the torrent
func (t Torrent) NumFiles() int {
	info, err := t.UnmarshalInfo()
	if err != nil {
		return -1
	}
	return len(info.Files)
}

// Passworded returns whether the contents of the torrent file require a password to access
func (t Torrent) Passworded() bool {
	return t.passworded
}

// SetPassworded sets whether the contents of the torrent require a password to access
func (t *Torrent) SetPassworded(b bool) {
	t.passworded = b
}

// Bytes returns the raw bytes of the torrent file
func (t Torrent) Bytes() ([]byte, error) {
	writer := bytes.NewBuffer(nil)
	err := t.Write(writer)
	if err != nil {
		return nil, err
	}
	return writer.Bytes(), nil
}

// URL returns a URL where the raw torrent file may be downloaded from
func (t Torrent) URL() *url.URL {
	return t.downloadURL
}

// SetURL sets the URL from which the raw torrent file may be downloaded
func (t *Torrent) SetURL(u *url.URL) {
	t.downloadURL = u
}
