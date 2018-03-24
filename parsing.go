package newznab

import (
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/mmcdole/gofeed"
	uuid "github.com/satori/go.uuid"
	"github.com/smquartz/errors"
)

// parseDate attempts to parse a date string
func parseDate(date string) (time.Time, error) {
	formats := []string{time.RFC3339, time.RFC1123Z}
	var parsedTime time.Time
	var err error
	for _, format := range formats {
		if parsedTime, err = time.Parse(format, date); err == nil {
			return parsedTime, nil
		}
	}
	return parsedTime, errors.Errorf("failed to parse date %s as one of %s", date, strings.Join(formats, ", "))
}

// entriesFromFeed takes a gofeed.Feed and returns a parsed []Entry slice
func entriesFromFeed(feed gofeed.Feed) ([]Entry, error) {
	var entries []Entry
	for _, item := range feed.Items {
		if item == nil {
			continue
		}
		entry, err := entryFromItem(feed, *item)
		if err != nil {
			continue
		}
		entries = append(entries, entry)
	}
	return entries, nil
}

// entryFromItem takes a gofeed.Item and returns a parsed Entry struct
func entryFromItem(feed gofeed.Feed, item gofeed.Item) (Entry, error) {
	var newEntry Entry
	// var err error
	// var year int
	/*
		var hasNFO bool
		var tvMovie TVMovie
		var tvMovieMusic TVMovieMusic
		var musicBook MusicBook
		var tvMovieMusicBook TVMovieMusicBook
		var movieMusicBook TVMovieMusicBook
	*/

	newEntry.Meta.Source.Feed = feed
	newEntry.Meta.Source.Item = item
	newEntry.Meta.Dates.Published = *item.PublishedParsed
	newEntry.Release.Name = item.Title
	if strings.Contains(item.Enclosures[0].URL, ".nzb") {
		newEntry.File = new(NZB)
	} else if strings.Contains(item.Enclosures[0].URL, ".torrent") || strings.Contains(item.Enclosures[0].URL, "magnet:?") {
		newEntry.File = new(Torrent)
	}
	u, _ := url.Parse(item.Enclosures[0].URL)
	newEntry.File.setURL(u)

	for _, attr := range item.Extensions["newznab"]["attr"] {
		name := attr.Attrs["name"]
		value := attr.Attrs["value"]
		intValue, intErr := strconv.ParseInt(value, 10, 64)

		switch name {
		// case "size":
		case "category":
			if intErr != nil {
				continue
			}
			cat := CategoryFromCode(int(intValue))
			newEntry.Meta.Categorisation.Categories = append(newEntry.Meta.Categorisation.Categories, cat)
		case "guid":
			guid, _ := uuid.FromString(value)
			newEntry.Meta.GUID = guid
		// case "files":
		case "poster":
			newEntry.Meta.Authoring.NNTPPoster = value
		case "group":
			newEntry.Meta.Authoring.NNTPGroups = append(newEntry.Meta.Authoring.NNTPGroups, value)
		case "team":
			newEntry.Release.Group = value
		case "grabs":
			newEntry.Meta.Grabs = intValue
		case "password":
			passworded, _ := strconv.ParseBool(value)
			newEntry.File.setPassworded(passworded)
		case "comments":
			// todo
		case "usenetdate":
			newEntry.Meta.Dates.PublishedUsenet, _ = parseDate(value)
		case "info":
			newEntry.Meta.NFO, _ = url.Parse(value)
		// case "year":
		// 	year = int(intValue)
		case "season":
			if newEntry.Content == nil {
				newEntry.Content = TV{}
			}
			tv, ok := newEntry.Content.(TV)
			if !ok {
				continue
			}
			tv.Season = int(intValue)
		case "episode":
			if newEntry.Content == nil {
				newEntry.Content = TV{}
			}
			tv, ok := newEntry.Content.(TV)
			if !ok {
				continue
			}
			if intErr != nil {
				intValue, intErr = strconv.ParseInt(strings.Split(value, "/")[0], 10, 64)
			}
			if intErr != nil {
				continue
			}
			tv.Episode = int(intValue)
		case "rageid":
			if newEntry.Content == nil {
				newEntry.Content = TV{}
			}
			tv, ok := newEntry.Content.(TV)
			if !ok {
				continue
			}
			tv.TVRageID = intValue
		case "tvtitle":
			if newEntry.Content == nil {
				newEntry.Content = TV{}
			}
			tv, ok := newEntry.Content.(TV)
			if !ok {
				continue
			}
			tv.TVRageTitle = value
		}
	}

	return newEntry, nil
}
