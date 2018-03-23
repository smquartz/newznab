package xml

import "fmt"

// errors used in this package
var ()

// error prefixes used in this package
const (
	ErrPrefixUnableEncodeStartElement  = "unable to encode start element: %v"
	ErrPrefixUnableEncodeRFC822UTCTime = "unable to encode RFC822 formatted UTC time"
	ErrPrefixUnableEncodeEndElement    = "unable to encode end element: %v"
	ErrPrefixUnableDecodeStartElement  = "unable to decode start element: %v"
	ErrPrefixUnableDecodeEndElement    = "unable to decode end element: %v"
)

// Error describes a parsed newznab error
type Error struct {
	// the error code
	Code int `xml:"code,attr"`
	// the error text
	Description string `xml:"description,attr"`
}

// Error implements the error interface for the Error type
func (e Error) Error() string { return fmt.Sprintf("newznab error: %s", e.Description) }
