package entry

import (
	"net/url"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/pborman/uuid"
)

// Entry describes a newznab entry
type Entry struct {
	// describes information about the entry itself rather than its contents
	Meta EntryMeta
	// describes information about the actual content; i.e. what is within the
	// files
	Content Content
	// describes the files themselves
	File File
}

// EntryMeta describes information about a newznab entry itself
type EntryMeta struct {
	// GUID for the entry
	GUID uuid.UUID
	// describes information relating to the source of the entry
	Source EntrySource
	// describes how the entry is categorised
	Categorisation EntryCategorisation
	// describes updated and published dates for the entry
	Dates EntryDates
	// describes information relating to who created the entry and its content
	Authoring EntryAuthoring
	// comments made on the entry
	Comments []Comment
	// number of times the entry has been downloaded
	Grabs int64
}

// EntrySource describes information relating to the source of an entry
type EntrySource struct {
	// the endpoint called to get the entry
	Endpoint *url.URL
	// API key used to authenticate against the endpoint
	APIKey string
	// RSS feed the entry was obtained from
	Feed *gofeed.Feed
	// RSS item the entry was obtained from
	Item *gofeed.Item
}

// EntryCategorisation describes how an entry is categorised
type EntryCategorisation struct {
	// categories strings that were provided in the RSS feed
	RSSCategories []string
	// categories that were provided through newznab:attr elements
	// Category struct details a category's ID and human readable name
	Categories []Category
}

// EntryDates descrbies the date an entry was first published, and last updated
type EntryDates struct {
	// when the entry's file was published to usenet
	PublishedUsenet time.Time
	// when the entry was first published
	Published time.Time
	// when the entry was last updated
	Updated time.Time
}

// EntryAuthoring describes who created an entry, and who provided its content
type EntryAuthoring struct {
	// pointer to the meta field, accessed from within feed/item methods
	meta *EntryMeta
	// the NNTP poster for the NZB file
	NNTPPoster string
	// the release group for the entry
	ReleaseGroup string
}

// FeedAuthor returns the author specified in the RSS feed the entry was
// obtained from
func (a EntryAuthoring) FeedAuthor() gofeed.Person {
	if a.meta != nil {
		if a.meta.Source.Feed != nil {
			return a.meta.Source.Feed.Author
		}
	}
	return gofeed.Person{}
}

// ItemAuthor returns the author specified in the RSS item the entry was
// obtained from
func (a EntryAuthoring) ItemAuthor() gofeed.Person {
	if a.meta != nil {
		if a.meta.Source.Item != nil {
			return a.meta.Source.Item.Author
		}
	}
	return gofeed.Person{}
}
