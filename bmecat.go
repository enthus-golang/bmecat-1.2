package bmecat

import "encoding/xml"

type BMEcat struct {
	XMLName xml.Name `xml:"BMECAT"`
	Version string   `xml:"version,attr"`

	Header Header `xml:"HEADER"`
}
