package xml

import (
	"encoding/xml"
	"net/url"
	"time"

	"github.com/pborman/uuid"
)

// rawAttribute describes an unparsed XML newznab:attr element
type rawAttribute struct {
	XMLName xml.Name
	Name    string `xml:"name,attr"`
	Value   string `xml:"value,attr"`
}

// Attributes describes a collection of newznab:attr elements
type Attributes interface {
	GetSize() int64
	GetCategory() int
	GetGUID() uuid.UUID
	GetFiles() int
	GetPoster() string
	GetGroups() []string
	GetTeam() string
	GetGrabs() int64
	GetPassworded() bool
	GetComments() int64
	GetUsenetDate() time.Time
	GetHasNFO() bool
	GetNFO() *url.URL
	GetYear() int
	GetTitle() string
}

// GenericAttributes describes the newznab:attr attributes that apply to all
// categories of newznab entries
type GenericAttributes struct {
	// size of the file contents in bytes
	Size int64
	// the ID of the newznab category the entry belongs to
	Category int
	// the actual GUID of the entry
	GUID uuid.UUID
	// number of files in the entry
	Files int
	// the NNTP poster of the NZB file
	Poster string
	// the NNTP group(s) for the NZB file
	Groups []string
	// the release group
	Team string
	// number of times the item has been downloaded
	Grabs int64
	// if applicable, whether the archive has a password
	Passworded *bool
	// number of comments on the entry
	Comments int64
	// date posted to usenet
	UsenetDate Time
	// whether the entry has an NFO or not
	HasNFO bool
	// nfo file url, if applicable
	NFO *url.URL
	// year of the release
	Year int
}

// GetTeam returns the release group for an entry
func (g GenericAttributes) GetTeam() string { return g.Team }

// GetGrabs returns the number of times an entry has been downloaded
func (g GenericAttributes) GetGrabs() int64 { return g.Grabs }

// GetComments returns the number of comments on an entry
func (g GenericAttributes) GetComments() int64 { return g.Comments }

// GetUsenetDate returns the date an entry was posted to usenet
func (g GenericAttributes) GetUsenetDate() time.Time { return g.UsenetDate }

// GetHasNFO returns whether an entry has a NFO
func (g GenericAttributes) GetHasNFO() bool { return g.HasNFO }

// GetNFO returns a pointer to a URL for the NFO file, if applicable
func (g GenericAttributes) GetNFO() *url.URL { return g.NFO }

// GetYear returns the year of release for an entry
func (g GenericAttributes) GetYear() int { return g.Year }

// GetSize returns the size of the file contents in bytes
func (g GenericAttributes) GetSize() int64 { return g.Size }

// GetCategory returns the ID of the category the entry belongs to
func (g GenericAttributes) GetCategory() int { return g.Category }

// GetGUID returns the actual GUID of the entry
func (g GenericAttributes) GetGUID() uuid.UUID { return g.GUID }

// GetFiles returns the number of files in the entry
func (g GenericAttributes) GetFiles() int { return g.Files }

// GetPoster returns the NNTP poster of the NZB file
func (g GenericAttributes) GetPoster() string { return g.Poster }

// GetGroups returns the NNTP group(s) for the NZB file
func (g GenericAttributes) GetGroups() []string { return g.Groups }

// GetPassworded returns whether the entry's contents require a password to
// access
func (g GenericAttributes) GetPassworded() *bool { return g.Passworded }

// TVAttributes describes the newznab:attr attributes specific to entries for
// a TV series
type TVAttributes struct {
	// embed the generic attributes
	GenericAttributes
	// embed attributes common to TV and Movie entries
	TVMovieAttributes
	// season number
	Season int
	// episode number relative to season
	Episode int
	// ID of the corresponding entry TVRage
	TVRageID int64
	// title of the series according to TVRage
	Title string
	// air date of the series according to TVRage
	Aired Time
}

// GetTitle returss the title of a TV entry
func (t TVAttributes) GetTitle() string { return t.Title }

// MovieAttributes describes the newznab:attr attributes specific to Movie
// entries
type MovieAttributes struct {
	// embed the generic attributes
	GenericAttributes
	// embed attributes common to TV and Movie entries
	TVMovieAttributes
	// embed attributes common to Movie, Music, and Book entries
	MovieMusicBookAttributes
	// ID for the corresponding IMDB entry
	IMDBID int64
	// score for the movie according to IMDB
	IMDBScore float32
	// title for the movie according to IMDB
	IMDBTitle string
	// tagline for the series according to IMDB
	IMDBTagline string
	// plot for the series according to IMDB
	IMDBPlot string
	// year of the series according to IMDB
	IMDBYear int
	// director of the series according to IMDB
	IMDBDirector string
	// actors in the series according to IMDB
	IMDBActors []string
}

// GetTitle returns the title of a Movie entry
func (m MovieAttributes) GetTitle() string { return m.IMDBTitle }

// MusicAttributes describes newznab:attr attributes specific to Music entries
type MusicAttributes struct {
	// embed the generic attributes
	GenericAttributes
	// embed attributes common to TV, Movie, Music, and Book entries
	TVMovieMusicBookAttributes
	// embed attributes common to TV, Movie, and Music entries
	TVMovieMusicAttributes
	// embed attributes common to Movie, Music, and Book entries
	MovieMusicBookAttributes
	// embed attributes common to Music and Book entries
	MusicBookAttributes
	// name of the artist
	Artist string
	// name of the album
	Album string
	// track listing
	Tracks []string
}

// GetTitle returns the name of the album for a Music entry
func (m MusicAttributes) GetTitle() string { return m.Album }

// BookAttributes describes newznab:attr attributes specific to Book entries
type BookAttributes struct {
	// embed the generic attributes
	GenericAttributes
	// embed attributes common to Music and Book entries
	MusicBookAttributes
	// embed attributes common to TV, Movie, Music, and Book entries
	TVMovieMusicBookAttributes
	// embed attributes common to Movie, Music, and Book entries
	MovieMusicBookAttributes
	// title of the book
	BookTitle string
	// date the book was published
	Published Time
	// author of the book
	Author string
	// number of pages in the book
	Pages int
}

// GetTitle returns the title of a Book entry
func (b BookAttributes) GetTitle() string { return b.BookTitle }

// TVMovieAudioAttributes describes newznab:attr attributes common to TV, Movie
// , and Audie newznab entries
type TVMovieAudioAttributes struct {
	// audio codec used
	AudioCodec string
	// audio languages available
	Languages []string
}

// TVMovieMusicBookAttributes describes newznab:attr attributes common to
// TV, Movie, Music, and Book entries
type TVMovieMusicBookAttributes struct {
	// url for a cover image
	CoverImage *url.URL
}

// TVMovieMusicAttributes describes newznab:attr attributes common to TV, Movie
// and Music entries
type TVMovieMusicAttributes struct {
	// url for a backdrop cover image
	BackdropCoverImage *url.URL
}

// MovieMusicBookAttributes describes newznab:attr attributes common to Movie,
// Music, and Book entries
type MovieMusicBookAttributes struct {
	// critics review score
	ReviewScore float32
}

// MusicBookAttributes describes newznab:attr attributes common to Music and
// Book entries
type MusicBookAttributes struct {
	// name of the publisher
	Publisher string
}

// TVMovieAttributes describes the newznab:attr attributes common to both TV
// and Movie entries
type TVMovieAttributes struct {
	// embed attributes common to TV, Movie, Audio entries
	TVMovieAudioAttributes
	// embed attributes common to TV, Movie, Music, and Book entries
	TVMovieMusicBookAttributes
	// embed attributes common to TV, Movie, and Music entries
	TVMovieMusicAttributes
	// video codec used
	VideoCodec string
	// video resolution
	Resolution string
	// framerate
	Framerate float32
	// subtitles available
	Subtitles []string
	// genre of the content
	Genre string
}
