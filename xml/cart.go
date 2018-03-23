package xml

import "encoding/xml"

// CartAdd describes the response format for when calling the cart-add command
type CartAdd struct {
	// name of the XML element
	XMLName xml.Name `xml:"cartadd"`
	// unknown what this is
	ID int64 `xml:"id,attr"`
}

// CartDel describes the response format for when calling the cart-del command
type CartDel struct {
	// name of the XML element
	XMLName xml.Name `xml:"cartdel"`
	// unknown what this is
	ID int64 `xml:"id,attr"`
}
