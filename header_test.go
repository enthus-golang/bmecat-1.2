package bmecat_test

import (
	"encoding/xml"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/enthus-golang/bmecat-1.2"
)

func TestHeader(t *testing.T) {
	header := bmecat.Header{
		GeneratorInformation: "TestEngine",
		CatalogInformation: bmecat.Catalog{
			Language:         "eng",
			ID:               "2012",
			Version:          "001.001",
			Name:             "asd",
			Date:             bmecat.NewDatetime("generation_date", time.Date(2010, 5, 7, 13, 21, 15, 0, time.UTC), true),
			StandardCurrency: "USD",
		},
		SupplierInformation: bmecat.Supplier{
			ID: &bmecat.TypeID{
				Type:  "supplier_specific",
				Value: "123123123",
			},
		},
	}

	b, err := xml.MarshalIndent(header, "", "  ")
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	assert.Equal(t, `<HEADER>
  <GENERATOR_INFO>TestEngine</GENERATOR_INFO>
  <CATALOG>
    <LANGUAGE>eng</LANGUAGE>
    <CATALOG_ID>2012</CATALOG_ID>
    <CATALOG_VERSION>001.001</CATALOG_VERSION>
    <CATALOG_NAME>asd</CATALOG_NAME>
    <DATETIME type="generation_date">
      <DATE>2010-05-07</DATE>
      <TIME>13:21:15</TIME>
    </DATETIME>
    <CURRENCY>USD</CURRENCY>
  </CATALOG>
  <SUPPLIER>
    <SUPPLIER_ID type="supplier_specific">123123123</SUPPLIER_ID>
    <SUPPLIER_NAME></SUPPLIER_NAME>
  </SUPPLIER>
  <USER_DEFINED_EXTENSIONS></USER_DEFINED_EXTENSIONS>
</HEADER>`, string(b))
}
