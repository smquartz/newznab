package xml

import "encoding/xml"

// CommentAdd describes the response format for when calling the comment-add
// command
type CommentAdd struct {
	// name of the XML element
	XMLName xml.Name `xml:"commentadd"`
	// id of the created comment
	ID int64 `xml:"id,attr"`
}
