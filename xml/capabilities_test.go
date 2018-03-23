package xml

import (
	"encoding/xml"
	"os"
	"reflect"
	"testing"
)

func TestCapabilitiesUnmarshalling(t *testing.T) {
	testCapabilitiesPath := "../samples/newznab/newznab_caps.xml"
	testCapabilities, err := os.Open(testCapabilitiesPath)
	if err != nil {
		t.Errorf("Failed to open test capabilities XML file: %v", err)
	}

	capabilities := new(Capabilities)

	decoder := xml.NewDecoder(testCapabilities)
	err = decoder.Decode(capabilities)
	if err != nil {
		t.Errorf("Failed to parse test capabilities XML: %v", err)
	}

	if capabilities.Server.Version != "0.1" {
		t.Errorf("Wrong server version: %s", capabilities.Server.Version)
	}
	if capabilities.Limits.Max != 60 {
		t.Errorf("Wrong limits max: %d", capabilities.Limits.Max)
	}
	if capabilities.Limits.Default != 25 {
		t.Errorf("Wrong limits default: %d", capabilities.Limits.Default)
	}
	if capabilities.Registration.Available != true {
		t.Errorf("Wrong registration available: %v", capabilities.Registration.Available)
	}
	if capabilities.Registration.Open != false {
		t.Errorf("Wrong registration open: %v", capabilities.Registration.Open)
	}
	if capabilities.Searching.Search.Available != true {
		t.Errorf("Wrong searching search available: %v", capabilities.Searching.Search.Available)
	}
	if capabilities.Searching.TV.Available != true {
		t.Errorf("Wrong searching tv-search available: %v", capabilities.Searching.TV.Available)
	}
	if capabilities.Searching.Movie.Available != true {
		t.Errorf("Wrong searching movie-search available: %v", capabilities.Searching.Movie.Available)
	}
	if capabilities.Searching.Audio.Available != true {
		t.Errorf("Wrong searching audio-search available: %v", capabilities.Searching.Audio.Available)
	}

	expectedCategories := []CapabilitiesCategory{
		CapabilitiesCategory{ID: 5000, Name: "TV", Subcategories: []CapabilitiesCategory{
			CapabilitiesCategory{ID: 5070, Name: "Anime"},
			CapabilitiesCategory{ID: 5080, Name: "Documentary"},
			CapabilitiesCategory{ID: 5020, Name: "Foreign"},
			CapabilitiesCategory{ID: 5040, Name: "HD"},
			CapabilitiesCategory{ID: 5050, Name: "Other"},
			CapabilitiesCategory{ID: 5030, Name: "SD"},
			CapabilitiesCategory{ID: 5060, Name: "Sport"},
			CapabilitiesCategory{ID: 5010, Name: "WEB-DL"},
		}},
		CapabilitiesCategory{ID: 7000, Name: "Other", Subcategories: []CapabilitiesCategory{
			CapabilitiesCategory{ID: 7010, Name: "Misc"},
		}},
		CapabilitiesCategory{ID: 8000, Name: "Books", Subcategories: []CapabilitiesCategory{
			CapabilitiesCategory{ID: 8020, Name: "Comics"},
		}},
	}
	if !reflect.DeepEqual(capabilities.Categories.Categories, expectedCategories) {
		t.Errorf("Wrong categories: got %v", capabilities.Categories.Categories)
	}

	if len(capabilities.Groups.Groups) != 0 {
		t.Errorf("Wrong numbebr of groups: %d", len(capabilities.Groups.Groups))
	}
	if len(capabilities.Genres.Genres) != 0 {
		t.Errorf("Wrong number of genres: %d", len(capabilities.Genres.Genres))
	}
}
