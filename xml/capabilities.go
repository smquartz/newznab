package xml

import (
	"encoding/xml"
	"strconv"
	"strings"

	"github.com/smquartz/errors"
)

// Capabilities describes the "capabilities" of a newznab indexer; it describes
// the output of the t=caps command, which returns information such as details
// of what is indexed, supported functions, etc
type Capabilities struct {
	// describes information about the indexer itself, rather than what
	// it indexes
	Server CapabilitiesServer `xml:"server"`
	// describes the limits imposed on searches
	Limits CapabilitiesLimits `xml:"limits"`
	// describes how long NZBs/whatever else are retained by an indexer before
	// being deleted
	Retention CapabilitiesRetention `xml:"retention"`
	// describes whether registration is available, and whether it is open, for
	// a given indexer
	Registration CapabilitiesRegistration `xml:"registration"`
	// describes what kinds of searches may be executed against an indexer
	Searching CapabilitiesSearching `xml:"searching"`
	// describes the newznab categories indexed by an indexer
	Categories []CapabilitiesCategory `xml:"categories>category"`
	// describes the usenet groups indexed by an indexer
	Groups []CapabilitiesGroup `xml:"groups>group"`
	// describes the genres indexed by an indexer
	Genres []CapabilitiesGenre `xml:"genres>genre"`
}

// SearchCapabilities describes whether a particular kind of search is supported
type SearchCapabilities struct {
	Available bool `xml:"available,attr"`
}

// UnmarshalXML enables the unmarshalling of XML into SearchCapabilities
func (sc *SearchCapabilities) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var search struct {
		Available string `xml:"available,attr"`
	}
	err := d.DecodeElement(&search, &start)
	if err != nil {
		return errors.Wrapf(err, "unable to parse searching element", 1)
	}

	sc.Available, err = parseYesNoBool(search.Available)
	if err != nil {
		return errors.Wrapf(err, "unable to parse available attribute: %s", 1, search.Available)
	}

	return nil
}

// CapabilitiesServer describes an indexer itself
type CapabilitiesServer struct {
	// the version of the newznab protoc implemented by the server
	Version string `xml:"version,attr"`
	// the title of the indexer
	Title string `xml:"title,attr"`
	// a tagline for the indexer
	Strapline string `xml:"strapline,attr"`
	// a contact email address for the indexer
	Email string `xml:"email,attr"`
	// a website for the indexer
	URL string `xml:"url,attr"`
	// a logo or other image for the indexer
	Image string `xml:"image,attr"`
}

// CapabilitiesLimits describes the limits an indexer imposes on searches
type CapabilitiesLimits struct {
	// describes the maximum number of items returned in a search
	Max int `xml:"max,attr"`
	// describes the default number of items returned in a search
	Default int `xml:"default,attr"`
}

// CapabilitiesRetention describes the how long an indexer retains content for
type CapabilitiesRetention struct {
	Days int `xml:"days,int"`
}

// CapabilitiesRegistration describes whether registration is available for an
// indexer, and whether it is currently open
type CapabilitiesRegistration struct {
	Available bool `xml:"available,attr"`
	Open      bool `xml:"open,attr"`
}

func parseYesNoBool(str string) (bool, error) {
	str2 := strings.ToLower(str)
	str2 = strings.Replace(str2, "yes", "true", -1)
	str2 = strings.Replace(str2, "no", "false", -1)
	return strconv.ParseBool(str2)
}

// UnmarshalXML enables the unmarshalling of XML into CapabilitiesRegistration
func (cr *CapabilitiesRegistration) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var reg struct {
		Available string `xml:"available,attr"`
		Open      string `xml:"open,attr"`
	}
	err := d.DecodeElement(&reg, &start)
	if err != nil {
		return errors.Wrapf(err, "unable to parse registration element", 1)
	}

	cr.Available, err = parseYesNoBool(reg.Available)
	if err != nil {
		return errors.Wrapf(err, "unable to parse available attribute: %s", 1, reg.Available)
	}

	cr.Open, err = parseYesNoBool(reg.Open)
	if err != nil {
		return errors.Wrapf(err, "unable to parse open attribute: %s", 1, reg.Open)
	}

	return nil
}

// CapabilitiesSearching describes what kinds of searches may be executed
// against an indexer
type CapabilitiesSearching struct {
	General SearchCapabilities `xml:"search"`
	TV      SearchCapabilities `xml:"tv-search"`
	Movie   SearchCapabilities `xml:"movie-search"`
	Audio   SearchCapabilities `xml:"audio-search"`
}

// CapabilitiesCategory describes an individual category indexed by an indexer
type CapabilitiesCategory struct {
	ID            int                    `xml:"id,attr"`
	Name          string                 `xml:"name,attr"`
	Subcategories []CapabilitiesCategory `xml:"subcat"`
}

// CapabilitiesGroup describes an individual usenet group indexed by an indexer
type CapabilitiesGroup struct {
	Name        string `xml:"name,attr"`
	Description string `xml:"description,attr"`
	LastUpdate  Time   `xml:"lastupdate,attr"`
}

// CapabilitiesGenre describes an individual genre that is indexed by an
// indexer
type CapabilitiesGenre struct {
	ID         string `xml:"id,attr"`
	Name       string `xml:"name,attr"`
	CategoryID int    `xml:"categoryid,attr"`
}
