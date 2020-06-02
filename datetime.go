package bmecat

import (
	"encoding/xml"
	"time"
)

type Datetime struct {
	XMLName xml.Name `xml:"DATETIME"`
	Type    string   `xml:"type,attr" validate:"required"`

	Date string `xml:"DATE"`
	Time string `xml:"TIME"`
}

func NewDatetime(typ string, t time.Time) *Datetime {
	return &Datetime{
		Type: typ,
		Date: t.Format("2006-01-02"),
		Time: t.Format("15:04:05"),
	}
}
