package bmecat

import (
	"encoding/xml"
	"time"
)

type Datetime struct {
	XMLName xml.Name `xml:"DATETIME"`
	Type    string   `xml:"type,attr" validate:"required"`

	Date string `xml:"DATE"`
	Time string `xml:"TIME,omitempty"`
}

func NewDatetime(typ string, t time.Time, withTime bool) *Datetime {
	d := &Datetime{
		Type: typ,
		Date: t.Format("2006-01-02"),
	}

	if withTime {
		d.Time = t.Format("15:04:05")
	}

	return d
}
