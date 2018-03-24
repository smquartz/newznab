package newznab

import (
	"os"
	"testing"

	"github.com/mmcdole/gofeed"
	uuid "github.com/satori/go.uuid"
)

func TestEntriesFromFeed(t *testing.T) {
	testPath := "samples/newznab/newznab_nzb_su.xml"
	testFile, err := os.Open(testPath)
	if err != nil {
		t.Errorf("Error opening test XML: %v", err)
	}

	parser := gofeed.NewParser()
	feed, err := parser.Parse(testFile)
	if err != nil {
		t.Errorf("Error parsing test XML: %v", err)
	}
	entries, err := entriesFromFeed(*feed)
	if err != nil {
		t.Errorf("Error parsing test feed: %v", err)
	}

	if len(entries) != 100 {
		t.Errorf("Wrong number of entries: %d", len(entries))
	}

	testEntry := entries[0]
	expectedGUID, _ := uuid.FromString("24967ef4c2e26296c65d3bbfa97aa8fe")
	if testEntry.Meta.GUID != expectedGUID {
		t.Errorf("Wrong GUID: %s", testEntry.Meta.GUID.String())
	}
	if testEntry.Release.Name != "White.Collar.S03E05.720p.HDTV.X264-DIMENSION" {
		t.Errorf("Wrong release name: %s", testEntry.Release.Name)
	}
	if testEntry.Meta.Categorisation.Categories[0] != CategoryFromCode(5000) {
		t.Errorf("Wrong category 0: %v", testEntry.Meta.Categorisation.Categories[0])
	}
	if testEntry.Meta.Categorisation.Categories[1] != CategoryFromCode(5040) {
		t.Errorf("Wrong category 1: %v", testEntry.Meta.Categorisation.Categories[1])
	}
	if testEntry.File != nil {
		if u := testEntry.File.URL(); u != nil {
			if u.String() != "http://nzb.su/getnzb/24967ef4c2e26296c65d3bbfa97aa8fe.nzb&i=37292&r=xxx" {
				t.Errorf("Wrong File URL: %s", u.String())
			}
		} else {
			t.Errorf("File URL is nil")
		}
	} else {
		t.Errorf("File field is empty")
	}
}
