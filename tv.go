package entry

import "time"

// TV describes the television content contained within an entry
type TV struct {
	// embed common information
	TVMovie
	// season number
	Season int
	// episode number relative to the season
	Episode int
	// ID of the corresponding entry in TVRage
	TVRageID int64
	// title of the series according to TVRage
	TVRageTitle string
	// air date of series according to TVRage
	Aired time.Time
}

// Title returns the title of TV Content according to TVRage
func (t TV) Title() string {
	return t.TVRageTitle
}

// ReleaseYear returns the year the TV Content aired according to TVRage
func (t TV) ReleaseYear() int {
	return t.Aired.Year()
}
