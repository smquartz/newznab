package xml

import (
	"encoding/xml"
	"strings"
	"time"

	"github.com/smquartz/errors"
)

type Time struct {
	time.Time
}

// MarshalXML defines how to marshal the Time struct into XML
func (t *Time) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeToken(start); err != nil {
		return errors.Wrapf(err, ErrPrefixUnableEncodeStartElement, 1, start.Name)
	}
	if err := e.EncodeToken(xml.CharData([]byte(t.UTC().Format(time.RFC822)))); err != nil {
		return errors.Wrapf(err, ErrPrefixUnableEncodeRFC822UTCTime, 1)
	}
	if err := e.EncodeToken(xml.EndElement{start.Name}); err != nil {
		return errors.Wrapf(err, ErrPrefixUnableEncodeEndElement, 1, start.Name)
	}
	return nil
}

// UnmarshalXML defines how to unmarshall XML into the Time struct
func (t *Time) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var raw string

	err := d.DecodeElement(&raw, &start)
	if err != nil {
		return err
	}
	date, err := time.Parse(time.RFC1123Z, raw)

	if err != nil {
		return err
	}

	*t = Time{Time: date}
	return nil

}

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
