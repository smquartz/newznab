package xml

import (
	"encoding/xml"
	"os"
	"testing"
)

func TestErrors(t *testing.T) {
	errorPath := "../samples/newznab/unauthorized.xml"
	errorFile, err := os.Open(errorPath)
	if err != nil {
		t.Errorf("unable to open test unauthorized file: %v", err)
	}

	e := new(Error)
	decoder := xml.NewDecoder(errorFile)
	err = decoder.Decode(e)
	if err != nil {
		t.Errorf("unable to decode test unauthorized XML: %v", err)
	}
	if e.Code != 100 {
		t.Errorf("Wrong error code: %d", e.Code)
	}
	if e.Description != "Incorrect user credentials" {
		t.Errorf("Wrong error description: %s", e.Description)
	}
}
