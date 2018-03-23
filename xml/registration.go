package xml

import "encoding/xml"

// Register describes the response format for when calling the register command
type Register struct {
	// name of the XML element
	XMLName xml.Name `xml:"register"`
	// username provided to register with
	Username string `xml:"username,attr"`
	// password provided to register with
	Password string `xml:"password,attr"`
	// api key created for the user
	APIKey string `xml:"apikey,attr"`
}
