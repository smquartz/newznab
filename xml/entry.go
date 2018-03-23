package xml

import "github.com/mmcdole/gofeed"

// Entry represents an unparsed single XML newznab item in search results
type Entry struct {
	// embed the generic RSS item
	gofeed.Item
	// describes the information contained in newznab:attr elements
	Attributes Attributes
}
