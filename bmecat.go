package bmecat

import "encoding/xml"

type BMEcat struct {
	XMLName xml.Name `xml:"BMECAT"`
	Version string   `xml:"version,attr"`

	Header Header `xml:"HEADER"`
}

type BMEcatNewCatalog struct {
	BMEcat

	NewCatalog TNewCatalog `xml:"T_NEW_CATALOG"`
}

func NewCatalog(header Header, catalog TNewCatalog) *BMEcatNewCatalog {
	return &BMEcatNewCatalog{
		BMEcat: BMEcat{
			Version: "1.2",
			Header: header,
		},
		NewCatalog: catalog,
	}
}
