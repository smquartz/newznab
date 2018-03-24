package entry

import (
	"net/url"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/satori/go.uuid"
)

// Entry describes a newznab entry
type Entry struct {
	// describes information about the entry itself rather than its contents
	Meta Meta
	// describes the scene release
	Release Release
	// describes information about the actual content; i.e. what is within the
	// files
	Content Content
	// describes the files themselves
	File File
}

// Raw returns the raw, unparsed, entry
func (e Entry) Raw() gofeed.Item {
	return e.Meta.Source.Item
}

// Meta describes information about a newznab entry itself
type Meta struct {
	// GUID for the entry
	GUID uuid.UUID
	// describes information relating to the source of the entry
	Source Source
	// describes how the entry is categorised
	Categorisation Categorisation
	// describes updated and published dates for the entry
	Dates Dates
	// describes information relating to who created the entry and its content
	Authoring Authoring
	// comments made on the entry
	// Comments []Comment
	// number of times the entry has been downloaded
	Grabs int64
	// optionally contains a link to a corresponding NFO file
	NFO *url.URL
}

// Source describes information relating to the source of an entry
type Source struct {
	// the endpoint called to get the entry
	Endpoint *url.URL
	// API key used to authenticate against the endpoint
	APIKey string
	// RSS feed the entry was obtained from
	Feed gofeed.Feed
	// RSS item the entry was obtained from
	Item gofeed.Item
}

// Categorisation describes how an entry is categorised
type Categorisation struct {
	// categories strings that were provided in the RSS feed
	RSSCategories []string
	// categories that were provided through newznab:attr elements
	// Category struct details a category's ID and human readable name
	Categories []Category
}

// Dates descrbies the date an entry was first published, and last updated
type Dates struct {
	// when the entry's file was published to usenet
	PublishedUsenet time.Time
	// when the entry was first published
	Published time.Time
	// when the entry was last updated
	Updated time.Time
}

// Authoring describes who created an entry, and who provided its content
type Authoring struct {
	// pointer to the entry, accessed from within feed/item/release methods
	entry *Entry
	// the NNTP poster for the NZB file
	NNTPPoster string
	// the NNTP groups for the NZB file
	NNTPGroups []string
}

// FeedAuthor returns the author specified in the RSS feed the entry was
// obtained from
func (a Authoring) FeedAuthor() *gofeed.Person {
	if a.entry != nil {
		return a.entry.Meta.Source.Feed.Author
	}
	return &gofeed.Person{}
}

// ItemAuthor returns the author specified in the RSS item the entry was
// obtained from
func (a Authoring) ItemAuthor() *gofeed.Person {
	if a.entry != nil {
		return a.entry.Meta.Source.Item.Author
	}
	return &gofeed.Person{}
}

// ReleaseGroup returns the name of the "release" group
func (a Authoring) ReleaseGroup() string {
	if a.entry != nil {
		return a.entry.Release.Group
	}
	return ""
}
