package bmecat

import "encoding/xml"

type Header struct {
	XMLName xml.Name `xml:"HEADER"`

	GeneratorInformation  string   `xml:"GENERATOR_INFO"`
	CatalogInformation    Catalog  `xml:"CATALOG"`
	BuyerInformation      *Buyer   `xml:"BUYER"`
	SkeletonAgreements    []string `xml:"AGREEMENT"`
	SupplierInformation   Supplier `xml:"SUPPLIER"`
	UserDefinedExtensions string   `xml:"USER_DEFINED_EXTENSIONS"`
}

type Catalog struct {
	XMLName xml.Name `xml:"CATALOG"`

	Language                string    `xml:"LANGUAGE" validate:"omitempty,len=3"`
	ID                      string    `xml:"CATALOG_ID" validate:"max=20"`
	Version                 string    `xml:"CATALOG_VERSION" validate:"max=7"`
	Name                    string    `xml:"CATALOG_NAME,omitempty" validate:"max=100"`
	Date                    *Datetime `xml:"DATETIME,omitempty"`
	TerritorialAvailability []string  `xml:"TERRITORY"`
	StandardCurrency        string    `xml:"CURRENCY,omitempty" validate:"omitempty,len=3"`
	MIMERoot                string    `xml:"MIME_ROOT,omitempty" validate:"max=100"`
}

type Buyer struct {
	XMLName xml.Name `xml:"BUYER"`

	ID   *TypeID `xml:"BUYER_ID,omitempty"`
	Name string  `xml:"BUYER_NAME" validate:"required,max=50"`
}

type Supplier struct {
	XMLName xml.Name `xml:"SUPPLIER"`

	ID   *TypeID `xml:"SUPPLIER_ID,omitempty"`
	Name string  `xml:"SUPPLIER_NAME" validate:"required,max=50"`
}

type TypeID struct {
	Type  string `xml:"type,attr" validate:"required"`
	Value string `xml:",innerxml" validate:"required"`
}
